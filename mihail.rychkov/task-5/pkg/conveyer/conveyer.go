package conveyer;

import "context";
import "errors"

type Conveyer[T any] struct {
	channelCapacity int;
	pipes map[string] chan T;
	nodes []func(c context.Context) error;
}

var ErrorChannelNotFound = errors.New("chan not found");
var ErrorClosedChanelEmpty = errors.New("requested channel was closed and is empty");

func New[T any](channelCapacity int) Conveyer[T] {
	return Conveyer[T]{channelCapacity, make(map[string] chan T), []func(c context.Context) error{}};
}

func (obj *Conveyer[T]) reserveChannel(name string) chan T {
	ch, exists := obj.pipes[name];
	if (exists) {
		return ch;
	}
	ch = make(chan T, obj.channelCapacity);
	obj.pipes[name] = ch;
	return ch;
}

func (obj *Conveyer[T]) Run(c context.Context) error {
	for _, fn := range(obj.nodes) {
		go fn(c);
	}
	return nil;
}
func (obj *Conveyer[T]) Send(inChName string, data T) error {
	ch, exists := obj.pipes[inChName];
	if (!exists) {
		return ErrorChannelNotFound;
	}
	ch <- data;
	return nil;
}
func (obj *Conveyer[T]) Recv(outChName string) (T, error) {
	ch, exists := obj.pipes[outChName];
	if (!exists) {
		var res T;
		return res, ErrorChannelNotFound;
	}
	res, ok := <- ch;
	if (!ok) {
		return res, ErrorClosedChanelEmpty;
	}
	return res, nil;
}

func (obj *Conveyer[T]) RegisterDecorator(
		fn func(c context.Context, input chan T, output chan T) error,
		input string, output string,
	) {
	in := obj.reserveChannel(input);
	out := obj.reserveChannel(output);
	obj.nodes = append(obj.nodes, func(c context.Context) error {
		return fn(c, in, out);
	});
}
