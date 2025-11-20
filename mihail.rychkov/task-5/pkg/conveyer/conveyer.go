package conveyer;

import "context";
import "errors";
import "golang.org/x/sync/errgroup";

type Conveyer[T any] struct {
	channelCapacity int;
	pipes map[string] chan T;
	nodes []func(c context.Context) error;
}

var ErrorChannelNotFound = errors.New("chan not found");
var ErrorClosedChanelEmpty = errors.New("requested channel was closed and is empty");

func NewConveyer[T any](channelCapacity int) Conveyer[T] {
	return Conveyer[T]{channelCapacity, make(map[string] chan T), []func(ctx context.Context) error{}};
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

func (obj *Conveyer[T]) Run(ctx context.Context) error {
	defer func() {
		for _, ch := range(obj.pipes) {
			close(ch);
		}
	}();

	group, ctx := errgroup.WithContext(ctx);
	for _, fn := range(obj.nodes) {
		group.Go(func() error { return fn(ctx); });
	}
	return group.Wait();
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
func (obj *Conveyer[T]) RegisterMultiplexer(
		fn func(c context.Context, input []chan T, output chan T) error,
		input []string, output string,
	) {
	inputs := make([]chan T, len(input));
	for idx, name := range(input) {
		inputs[idx] = obj.reserveChannel(name);
	}
	out := obj.reserveChannel(output);
	obj.nodes = append(obj.nodes, func(c context.Context) error {
		return fn(c, inputs, out);
	});
}
func (obj *Conveyer[T]) RegisterSeparator(
		fn func(c context.Context, input chan T, output []chan T) error,
		input string, output []string,
	) {
	in := obj.reserveChannel(input);
	outputs := make([]chan T, len(output));
	for idx, name := range(output) {
		outputs[idx] = obj.reserveChannel(name);
	}
	obj.nodes = append(obj.nodes, func(c context.Context) error {
		return fn(c, in, outputs);
	});
}
