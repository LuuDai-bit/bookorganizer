package jobs

import (
	"book-organizer/models"

	"github.com/go-co-op/gocron/v2"
)

type Cron struct{}

func (c *Cron) RunSchedule() {
	s, err := gocron.NewScheduler()
	if err != nil {
		return
	}

	fileModel := new(models.FileModel)
	_, err = s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(23, 30, 0),
			),
		),
		gocron.NewTask(fileModel.RemoveUnusedFile),
	)

	if err != nil {
		return
	}

	s.Start()
}
