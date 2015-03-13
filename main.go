package main

import (
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/shell"
	_ "github.com/danryan/hal/adapter/slack"
	_ "github.com/danryan/hal/store/memory"
	"os"
)

var canYouHandler = hal.Hear("(Tim|tim).*can you.*", func(res *hal.Response) error {
	return res.Send("on it!")
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
