package app

import (
	"context"
	"github.com/gelleson/wait/pkg/env"
	"github.com/gelleson/wait/pkg/wait"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"sync"
)

var links []string

var root = &cobra.Command{
	Use: "wait",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup

		wg.Add(len(links))

		for _, link := range links {
			go func(l string) {
				defer wg.Done()

				n := strings.SplitN(l, "://", 2)

				if len(n) != 2 {
					log.Fatalln("is not valid url to resource")
				}

				waiter, err := wait.GetWaiter(n[0])

				if err != nil {
					log.Fatalln(err)
				}

				if err := waiter.Wait(context.Background(), n[1], 10, 10); err != nil {
					log.Fatalln(err)
				}

			}(link)
		}

		wg.Wait()
	},
}

func Execute() error {
	return root.Execute()
}

func init() {
	arrayString, err := env.GetArrayString("WAIT_HOSTS", ",")

	if err != nil {
		log.Fatalln(err)
	}

	root.PersistentFlags().StringArrayVar(&links, "hosts", arrayString, "tcp://localhost:5432,tcp://localhost:5433")
}
