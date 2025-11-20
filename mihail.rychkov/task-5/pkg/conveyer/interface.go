package conveyer;

type Interface[T any] interface {
	RegisterDecorator(
		fn func(
			ctx context.Context,
			input chan T,
			output chan T,
		) error,
		inputChannelName string,
		outputChannelName string,
	);
	RegisterMultiplexer(
		fn func(
			ctx context.Context,
			inputs []chan T,
			output chan T,
		) error,
		inputChannelsNames []string,
		outputChannelName string,
	);
	RegisterSeparator(
		fn func(
			ctx context.Context,
			input chan T,
			outputs []chan T,
		) error,
		inputChannelName string,
		outputChannelsNames []string,
	);
	Run(c context.Context) error;
	Send(inputChannelName string, data T) error;
	Recv(outputChannelName string) (T, error);
}
