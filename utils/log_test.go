package utils

import (
	"errors"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestLog1(t *testing.T) {
	log, _ := CreateLog("logs", "spans", logrus.DebugLevel, true)
	logger := log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	})
	logger.Info("A group of walrus emerges from the ocean")
	logger.WithField("err", errors.New("hello error")).Error("test error")

	time.Sleep(time.Second)
}
