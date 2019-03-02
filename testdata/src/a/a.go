package a

import (
	"context"
	"fmt"
)

func main() {
	log := test{}
	log.Debugf(context.Background(), "test %d", 1) // want "detect log.Debugf use debug code"
}

type test struct{}

func (t test) Debugf(ctx context.Context, s string, args ...interface{}) {
	fmt.Println(ctx, s, args)
}
