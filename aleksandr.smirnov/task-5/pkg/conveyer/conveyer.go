package conveyer

import (
	"context"
	"errors"
	"sync"
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
