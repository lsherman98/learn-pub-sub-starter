package main

import (
	"fmt"

	"github.com/lsherman98/learn-pub-sub-starter/internal/gamelogic"
	"github.com/lsherman98/learn-pub-sub-starter/internal/pubsub"
	"github.com/lsherman98/learn-pub-sub-starter/internal/routing"
)

func handlerLogs() func(gamelog routing.GameLog) pubsub.Acktype {
	return func(gamelog routing.GameLog) pubsub.Acktype {
		defer fmt.Print("> ")

		err := gamelogic.WriteLog(gamelog)
		if err != nil {
			fmt.Printf("error writing log: %v\n", err)
			return pubsub.NackRequeue
		}
		return pubsub.Ack
	}
}
