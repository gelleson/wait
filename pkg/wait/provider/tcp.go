package provider

import (
	"context"
	"log"
	"net"
	"time"
)

type WaiterTCP struct{}

func (w WaiterTCP) Wait(ctx context.Context, url string, waitInSeconds uint, attempts uint) error {

	errChain := make(chan error, 1)

	go func() {
		var e error

		for i := 0; i < int(attempts)+1; i++ {
			if err := w.connect(url, time.Second*time.Duration(waitInSeconds)); err != nil {
				log.Printf("connection error: %s next attempt will be in %d attempt %d\n", err.Error(), waitInSeconds, int(attempts)-i)
				e = err
			}

			if e == nil {
				errChain <- nil
				return
			}

			time.Sleep(time.Second * time.Duration(waitInSeconds))
		}

		errChain <- e
	}()

	for {
		select {
		case err := <-errChain:
			return err
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return err
			}

			return nil
		}
	}
}

func (w WaiterTCP) connect(url string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", url, timeout)
	if err != nil {
		return err
	}

	conn.Close()

	return nil
}
