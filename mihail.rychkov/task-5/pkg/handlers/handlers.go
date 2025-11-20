package handlers;

import "context";
import "strings";
import "errors";

var ErrorNoDecorator = errors.New("can't be decorated");
var ErrorEmptyChannelList = errors.New("channels slice is empty");

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select
		{
		case str, ok := <- input:
			if (!ok) {
				return nil;
			}
			if (strings.Contains(str, "no decorator")) {
				return ErrorNoDecorator;
			}
			select {
			case output <- "decorated: " + str:
			case <- ctx.Done():
				return nil;
			}
		case <- ctx.Done():
			return nil;
		}
	}
	return nil;
}
func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if (len(outputs) == 0) {
		return ErrorEmptyChannelList;
	}

	var idx int;
	for {
		select
		{
		case str, ok := <- input:
			if (!ok) {
				return nil;
			}
			select {
			case outputs[idx] <- str:
			case <- ctx.Done():
				return nil;
			}
			idx = (idx + 1) % len(outputs);
		case <- ctx.Done():
			return nil;
		}
	}
	return nil;
}
