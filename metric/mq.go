package metric

import (
	"grandhelmsman/filecoin-monitor/utils"
	"github.com/assembla/cony"
	"github.com/streadway/amqp"
)

var (
	rabbitPub *cony.Publisher
)

func initRabbit(url string) {
	rabbitPub = utils.GetRabbitPublisher(url, options.Exchange, options.RouteKey)
}

func sendToRabbit(data []byte) error {
	return rabbitPub.PublishWithRoutingKey(amqp.Publishing{Body: data}, options.RouteKey)
}

