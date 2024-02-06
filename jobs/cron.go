package jobs

import (
	"github.com/go-co-op/gocron/v2"
)

type Cron struct{}

func (c *Cron) RunSchedule() {
	s, err := gocron.NewScheduler()
	if err != nil {
		return
	}

	new(RemoveUnusedFile).Perform(s)

	new(ImportGoogleBook).Perform(s)

	s.Start()
}
