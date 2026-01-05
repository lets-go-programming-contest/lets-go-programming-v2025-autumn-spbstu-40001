package conveyer

import (
	"context"
	"errors"
	"sync"
)

var ErrChannelNotFound = errors.New("chan not found")

const undefined = "undefined"

type Pipeline struct {
	size      int
	channels  map[string]chan string
	mu        sync.RWMutex
	modifiers []modifierEntry
	muxers    []muxEntry
	splitters []splitterEntry
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
		size:      size,
		channels:  make(map[string]chan string),
		mu:        sync.RWMutex{},
		modifiers: make([]modifierEntry, 0),
		muxers:    make([]muxEntry, 0),
		splitters: make([]splitterEntry, 0),
	}
}
