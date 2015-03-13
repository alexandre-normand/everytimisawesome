package main

import (
	"fmt"
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/slack"
	_ "github.com/danryan/hal/store/memory"
	"os"
)

var canYouHandler = hal.Hear("(Tim|tim).*can you.*", func(res *hal.Response) error {
	return res.Send("on it!")
})

var karmaHandler = hal.Hear(".*(\\w+)(\\+\\+|\\-\\-).*", func(res *hal.Response) error {
	var format string
	if res.Match[2] == "++" {
		format = "%s just gained a level (%s: %d)"
	} else {
		format = "%s just lost a life (%s: %d)"
	}

	thing := res.Match[1]

	return res.Reply(fmt.Sprintf(format, thing, thing, 1))
})

var echoHandler = hal.Respond(`echo (.+)`, func(res *hal.Response) error {
	return res.Reply(res.Match[1])
})

func run() int {
	robot, err := hal.NewRobot()

	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	robot.Handle(
		canYouHandler,
		echoHandler,
		karmaHandler,
	)

	if err := robot.Run(); err != nil {
		hal.Logger.Error(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
