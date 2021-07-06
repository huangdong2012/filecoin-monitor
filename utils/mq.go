package utils

import (
	"fmt"
	"github.com/assembla/cony"
)

func GetRabbitPublisher(url, ex, key string) *cony.Publisher {
	cli := cony.NewClient(
		cony.URL(url),
		cony.Backoff(cony.DefaultBackoff),
	)

	exc := cony.Exchange{
		Name:       ex,
		Kind:       "topic",
		AutoDelete: false,
	}
	cli.Declare([]cony.Declaration{
		cony.DeclareExchange(exc),
	})

	publisher := cony.NewPublisher(exc.Name, key)
	cli.Publish(publisher)
	go loopRabbit(cli)

	return publisher
}

func loopRabbit(cli *cony.Client) {
	for cli.Loop() {
		select {
		case err := <-cli.Errors():
			Error(fmt.Errorf("Rabbit Client error: %v\n", err))
		case blocked := <-cli.Blocking():
			Error(fmt.Errorf("Rabbit Client is blocked %v\n", blocked))
		}
	}
}
