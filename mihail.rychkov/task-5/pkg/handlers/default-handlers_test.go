package handlers;

import "testing";
import "context";
import "github.com/stretchr/testify/assert";
import "github.com/Rychmick/task-5/pkg/conveyer"

func TestDecoratorConveyer(t *testing.T) {
	var conv = conveyer.New[string](5);
	conv.RegisterDecorator(PrefixDecoratorFunc, "in", "mid");
	conv.RegisterDecorator(PrefixDecoratorFunc, "mid", "out");

	err := conv.Send("in", "allo");
	assert.Nil(t, err, "no error expected");
	err = conv.Send("in", "does it works?");
	assert.Nil(t, err, "no error expected");
	err = conv.Run(context.Background());
	assert.Nil(t, err, "no error expected");

	res, err := conv.Recv("out");
	assert.Nil(t, err, "no error expected");
	assert.Equal(t, "decorated: decorated: allo", res);
	res, err = conv.Recv("out");
	assert.Nil(t, err, "no error expected");
	assert.Equal(t, "decorated: decorated: does it works?", res);
}
func TestDecoratorFailConveyer(t *testing.T) {
	var conv = conveyer.New[string](5);
	conv.RegisterDecorator(PrefixDecoratorFunc, "in", "mid");
	conv.RegisterDecorator(PrefixDecoratorFunc, "mid", "out");

	err := conv.Send("in", "text no decorator contains");
	assert.Nil(t, err, "no error expected");
	err = conv.Run(context.Background());
	assert.ErrorIs(t, err, ErrorNoDecorator, "expected decorator cancel message");

	res, err := conv.Recv("out");
	assert.Nil(t, err, "no error expected");
	assert.Equal(t, "undefined", res);
}
func TestMuxConveyer(t *testing.T) {
	var conv = conveyer.New[string](5);
	conv.RegisterSeparator(SeparatorFunc, "in", []string{"out1", "out2", "out3"});

	err := conv.Send("in", "1");
	assert.Nil(t, err, "no error expected");
	err = conv.Send("in", "2");
	assert.Nil(t, err, "no error expected");
	err = conv.Send("in", "3");
	assert.Nil(t, err, "no error expected");
	err = conv.Send("in", "4");
	assert.Nil(t, err, "no error expected");
	err = conv.Run(context.Background());
	assert.Nil(t, err, "no error expected");

	res, err := conv.Recv("out1");
	assert.Nil(t, err, "no error expected");
	assert.Equal(t, "1", res);
	res, err = conv.Recv("out1");
	assert.Nil(t, err, "no error expected");
	assert.Equal(t, "4", res);
	res, err = conv.Recv("out2");
	assert.Nil(t, err, "no error expected");
	assert.Equal(t, "2", res);
	res, err = conv.Recv("out3");
	assert.Nil(t, err, "no error expected");
	assert.Equal(t, "3", res);
}
