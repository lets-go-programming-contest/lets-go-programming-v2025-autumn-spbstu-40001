package handlers;

import "context";
import "strings";
import "errors";
import "fmt";

var ErrorNoDecorator = errors.New("can't be decorated");
var ErrorEmptyChannelList = errors.New("channels slice is empty");

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	defer close(output);
	for str := range(input) {
		if (strings.Contains(str, "no decorator")) {
			return ErrorNoDecorator;
		}
		output <- "decorated: " + str;
	}
	return nil;
}
func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if (len(outputs) == 0) {
		return ErrorEmptyChannelList;
	}

	defer func(){
		for i, ch := range(outputs) {
			close(ch);
			fmt.Println("closed", i);
		}
	}();

	var idx int;
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
