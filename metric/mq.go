package metric

import (
	"github.com/assembla/cony"
	"github.com/streadway/amqp"
	"grandhelmsman/filecoin-monitor/utils"
)

var (
	rabbitPub *cony.Publisher
)

func initRabbit() {
	rabbitPub = utils.GetRabbitPublisher(base.MQUrl, options.Exchange, options.RouteKey)
}

func sendToRabbit(data []byte) error {
	return rabbitPub.PublishWithRoutingKey(amqp.Publishing{Body: data}, options.RouteKey)
}
