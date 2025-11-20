package handlers_test

import (
	"context"
	"testing"
	"time"

	"github.com/Rychmick/task-5/pkg/conveyer"
	"github.com/Rychmick/task-5/pkg/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertGoodResult(t *testing.T, err *error, conv *conveyer.StringConveyer, inName, outName, send, expected string) {
	t.Helper()

	if *err != nil {
		return
	}

	*err = conv.Send(inName, send)
	if *err != nil {
		return
	}

	var res string

	res, *err = conv.Recv(outName)
	if *err != nil {
		return
	}

	assert.Equal(t, expected, res)
}

func TestDecoratorConveyer(t *testing.T) {
	t.Parallel()

	conv := conveyer.New(5)

	conv.RegisterDecorator(handlers.PrefixDecoratorFunc, "in", "mid")
	conv.RegisterDecorator(handlers.PrefixDecoratorFunc, "mid", "out")

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*1)

	var runErr error

	go func() {
		assertGoodResult(t, &runErr, &conv, "in", "out", "1", "decorated: "+"decorated: "+"1")
		assertGoodResult(t, &runErr, &conv, "in", "out", "2", "decorated: "+"decorated: "+"2")

		cancelFunc()
	}()

	err := conv.Run(ctx)
	require.NoError(t, err)
	require.NoError(t, runErr)
}

func TestDecoratorFailConveyer(t *testing.T) {
	t.Parallel()

	conv := conveyer.New(5)

	conv.RegisterDecorator(handlers.PrefixDecoratorFunc, "in", "mid")
	conv.RegisterDecorator(handlers.PrefixDecoratorFunc, "mid", "out")

	err := conv.Send("in", "text no decorator contains")
	require.NoError(t, err)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*1)
	defer cancelFunc()

	err = conv.Run(ctx)
	require.ErrorIs(t, err, handlers.ErrNoDecorator, "expected decorator cancel message")

	res, err := conv.Recv("out")
	require.NoError(t, err)
	assert.Equal(t, "undefined", res)
}

func TestDMuxConveyer(t *testing.T) {
	t.Parallel()

	conv := conveyer.New(5)

	conv.RegisterSeparator(handlers.SeparatorFunc, "in", []string{"out1", "out2", "out3"})

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*1)

	var runErr error

	go func() {
		assertGoodResult(t, &runErr, &conv, "in", "out1", "1", "1")
		assertGoodResult(t, &runErr, &conv, "in", "out2", "2", "2")
		assertGoodResult(t, &runErr, &conv, "in", "out3", "3", "3")
		assertGoodResult(t, &runErr, &conv, "in", "out1", "4", "4")

		cancelFunc()
	}()

	err := conv.Run(ctx)
	require.NoError(t, err)
	require.NoError(t, runErr)
}

func TestMuxConveyer(t *testing.T) {
	t.Parallel()

	conv := conveyer.New(5)

	conv.RegisterMultiplexer(handlers.MultiplexerFunc, []string{"in1", "in2", "in3"}, "out")

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*1)

	var runErr error

	go func() {
		assertGoodResult(t, &runErr, &conv, "in1", "out", "1", "1")
		assertGoodResult(t, &runErr, &conv, "in2", "out", "2", "2")
		assertGoodResult(t, &runErr, &conv, "in3", "out", "3", "3")
		assertGoodResult(t, &runErr, &conv, "in1", "out", "4", "4")

		if runErr == nil {
			runErr = conv.Send("in2", "text no multiplexer 5")
		}

		cancelFunc()
	}()

	err := conv.Run(ctx)
	require.NoError(t, err)
	require.NoError(t, runErr)

	res, err := conv.Recv("out")
	require.NoError(t, err)
	assert.Equal(t, "undefined", res)
}
