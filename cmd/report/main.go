package main

import (
	"encoding/json"
	"fmt"
	"github.com/assembla/cony"
	"grandhelmsman/filecoin-monitor/cmd/report/db"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	dbUrl             = "user=root password=root dbname=zdz host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	rabbitUrl         = "amqp://root:root@localhost/"
	rabbitExTrace     = "zdz.exchange.trace"
	rabbitQueueTrace  = "zdz.queue.trace"
	rabbitExMetric    = "zdz.exchange.metric"
	rabbitQueueMetric = "zdz.queue.metric"
)

func main() {
	model.InitBaseOptions(&model.BaseOptions{
		LogErr: func(err error) {
			fmt.Printf("error: %v\n", err)
		},
		LogInfo: func(info string) {
			fmt.Printf("info: %v\n", info)
		},
	})
	db.Init(dbUrl)

	cnsTrace := utils.GetRabbitConsumer(rabbitUrl, rabbitExTrace, "*", rabbitQueueTrace)
	go loopRabbit(cnsTrace, handleTrace)

	cnsMetric := utils.GetRabbitConsumer(rabbitUrl, rabbitExMetric, "*", rabbitQueueMetric)
	go loopRabbit(cnsMetric, handleMetric)

	exitC := make(chan os.Signal)
	signal.Notify(exitC, os.Interrupt, syscall.SIGTERM)
	<-exitC
	os.Exit(0)
}

func loopRabbit(cns *cony.Consumer, handler func([]byte) error) {
	for {
		select {
		case msg := <-cns.Deliveries():
			if err := handler(msg.Body); err != nil {
				panic(err)
			}
			if err := msg.Ack(false); err != nil {
				utils.Error(fmt.Errorf("Consumer ack error: %v\n", err))
			}
		case err := <-cns.Errors():
			utils.Error(fmt.Errorf("Consumer error: %v\n", err))
			time.Sleep(time.Second * 5)
		}
	}
}

func handleTrace(data []byte) error {
	span := &model.Span{}
	if err := json.Unmarshal(data, span); err != nil {
		return err
	}

	return db.InsertSpan(span)
}

func handleMetric(data []byte) error {
	ms := make([]*model.Metric, 0, 0)
	if err := json.Unmarshal(data, &ms); err != nil {
		return err
	}

	return db.InsertMetrics(ms)
}
