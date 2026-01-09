package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var ErrChannelNotFound = errors.New("chan not found")

const undefined = "undefined"

type Pipeline struct {
	size       int
	channels   map[string]chan string
	muChannels sync.RWMutex
	modifiers  []modifierEntry
	muxers     []muxEntry
	splitters  []splitterEntry
}

type modifierEntry struct {
	function func(ctx context.Context, input chan string, output chan string) error
	input    string
	output   string
}

type muxEntry struct {
	function func(ctx context.Context, inputs []chan string, output chan string) error
	inputs   []string
	output   string
}

type splitterEntry struct {
	function func(ctx context.Context, input chan string, outputs []chan string) error
	input    string
	outputs  []string
}

func New(size int) *Pipeline {
	return &Pipeline{
		size:       size,
		channels:   make(map[string]chan string),
		muChannels: sync.RWMutex{},
		modifiers:  make([]modifierEntry, 0),
		muxers:     make([]muxEntry, 0),
		splitters:  make([]splitterEntry, 0),
	}
}

func (p *Pipeline) getOrCreateChannel(name string) chan string {
	p.muChannels.Lock()
	defer p.muChannels.Unlock()

	if ch, exists := p.channels[name]; exists {
		return ch
	}

	ch := make(chan string, p.size)
	p.channels[name] = ch

	return ch
}

func (p *Pipeline) getChannel(name string) (chan string, bool) {
	p.muChannels.RLock()
	defer p.muChannels.RUnlock()

	ch, exists := p.channels[name]

	return ch, exists
}

func (p *Pipeline) closeAllChannels() {
	p.muChannels.Lock()
	defer p.muChannels.Unlock()

	for _, ch := range p.channels {
		close(ch)
	}
}

func (p *Pipeline) Send(input string, data string) error {
	ch, exists := p.getChannel(input)
	if !exists {
		return ErrChannelNotFound
	}

	ch <- data

	return nil
}

func (p *Pipeline) Recv(output string) (string, error) {
	ch, exists := p.getChannel(output)
	if !exists {
		return "", ErrChannelNotFound
	}

	value, ok := <-ch
	if !ok {
		return undefined, nil
	}

	return value, nil
}

func (p *Pipeline) RegisterDecorator(
	function func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	p.getOrCreateChannel(input)
	p.getOrCreateChannel(output)

	p.modifiers = append(p.modifiers, modifierEntry{
		function: function,
		input:    input,
		output:   output,
	})
}

func (p *Pipeline) RegisterMultiplexer(
	function func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, inputName := range inputs {
		p.getOrCreateChannel(inputName)
	}

	p.getOrCreateChannel(output)

	p.muxers = append(p.muxers, muxEntry{
		function: function,
		inputs:   inputs,
		output:   output,
	})
}

func (p *Pipeline) RegisterSeparator(
	function func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	p.getOrCreateChannel(input)

	for _, outputName := range outputs {
		p.getOrCreateChannel(outputName)
	}

	p.splitters = append(p.splitters, splitterEntry{
		function: function,
		input:    input,
		outputs:  outputs,
	})
}

func (p *Pipeline) Run(ctx context.Context) error {
	defer p.closeAllChannels()

	group, groupCtx := errgroup.WithContext(ctx)

	for _, modifier := range p.modifiers {
		group.Go(func() error {
			inputChan := p.getOrCreateChannel(modifier.input)
			outputChan := p.getOrCreateChannel(modifier.output)

			return modifier.function(groupCtx, inputChan, outputChan)
		})
	}

	for _, mux := range p.muxers {
		group.Go(func() error {
			inputChans := make([]chan string, len(mux.inputs))
			for i, inputName := range mux.inputs {
				inputChans[i] = p.getOrCreateChannel(inputName)
			}

			outputChan := p.getOrCreateChannel(mux.output)

			return mux.function(groupCtx, inputChans, outputChan)
		})
	}

	for _, splitter := range p.splitters {
		group.Go(func() error {
			inputChan := p.getOrCreateChannel(splitter.input)

			outputChans := make([]chan string, len(splitter.outputs))
			for i, outputName := range splitter.outputs {
				outputChans[i] = p.getOrCreateChannel(outputName)
			}

			return splitter.function(groupCtx, inputChan, outputChans)
		})
	}

	if err := group.Wait(); err != nil {
		return fmt.Errorf("conveyer run failed: %w", err)
	}

	return nil
}
