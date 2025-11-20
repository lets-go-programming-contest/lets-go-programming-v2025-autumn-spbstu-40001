package conveyer;

import "testing";
import "context";

func TestSingleNodeConveyer(t *testing.T) {
	var c = New[int](5);
	foo := func(c context.Context, in, out chan int) error {
		for i := range(in) {
			out <- i*2;
		}
		close(out);
		return nil;
	};
	c.RegisterDecorator(foo, "in", "out");
	c.Send("in", 5);
	c.Run(context.TODO());
	r, _ := c.Recv("out");
	t.Logf("recieved %d", r);
	t.Error("allo");
}
