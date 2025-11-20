package handlers;

import "time";
import "testing";
import "context";
import "github.com/stretchr/testify/assert";
import "github.com/Rychmick/task-5/pkg/conveyer"

func assertGoodResult(t *testing.T, conv *conveyer.StringConveyer, inName, outName, send, expected string) {
	err := conv.Send(inName, send);
	assert.Nil(t, err, "no error expected");
	res, err := conv.Recv(outName);
	assert.Nil(t, err, "no error expected");
	assert.Equal(t, expected, res);
}

func TestDecoratorConveyer(t *testing.T) {
	var conv = conveyer.New(5);
	conv.RegisterDecorator(PrefixDecoratorFunc, "in", "mid");
	conv.RegisterDecorator(PrefixDecoratorFunc, "mid", "out");

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second * 1);
	go func() {
		assertGoodResult(t, &conv, "in", "out", "1", "decorated: decorated: 1");
		assertGoodResult(t, &conv, "in", "out", "2", "decorated: decorated: 2");

		cancelFunc();
	}();
	err := conv.Run(ctx);
	assert.Nil(t, err, "no error expected");
}
func TestDecoratorFailConveyer(t *testing.T) {
	var conv = conveyer.New(5);
	conv.RegisterDecorator(PrefixDecoratorFunc, "in", "mid");
	conv.RegisterDecorator(PrefixDecoratorFunc, "mid", "out");

	err := conv.Send("in", "text no decorator contains");
	assert.Nil(t, err, "no error expected");
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second * 1);
	defer cancelFunc();
	err = conv.Run(ctx);
	assert.ErrorIs(t, err, ErrorNoDecorator, "expected decorator cancel message");

	res, err := conv.Recv("out");
	assert.Nil(t, err, "no error expected");
	assert.Equal(t, "undefined", res);
}
func TestDMuxConveyer(t *testing.T) {
	var conv = conveyer.New(5);
	conv.RegisterSeparator(SeparatorFunc, "in", []string{"out1", "out2", "out3"});

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second * 1);
	go func() {
		assertGoodResult(t, &conv, "in", "out1", "1", "1");
		assertGoodResult(t, &conv, "in", "out2", "2", "2");
		assertGoodResult(t, &conv, "in", "out3", "3", "3");
		assertGoodResult(t, &conv, "in", "out1", "4", "4");

		cancelFunc();
	}();
	err := conv.Run(ctx);
	assert.Nil(t, err, "no error expected");
}
func TestMuxConveyer(t *testing.T) {
	var conv = conveyer.New(5);
	conv.RegisterMultiplexer(MultiplexerFunc, []string{"in1", "in2", "in3"}, "out");

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second * 1);
	go func() {
		assertGoodResult(t, &conv, "in1", "out", "1", "1");
		assertGoodResult(t, &conv, "in2", "out", "2", "2");
		assertGoodResult(t, &conv, "in3", "out", "3", "3");
		assertGoodResult(t, &conv, "in1", "out", "4", "4");

		cancelFunc();
	}();
	err := conv.Run(ctx);
	assert.Nil(t, err, "no error expected");
}
