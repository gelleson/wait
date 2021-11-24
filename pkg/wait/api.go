package wait

import (
	"context"
	"errors"
	"fmt"
	"github.com/gelleson/wait/pkg/wait/provider"
)

type Waiter interface {
	// Wait - func will wait until deadline exceed
	Wait(ctx context.Context, url string, waitInSeconds uint, attempts uint) error
}

func GetWaiter(typex string) (Waiter, error) {
	switch typex {
	case "tcp":
		return &provider.WaiterTCP{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("%s doesn't support", typex))
	}
}
