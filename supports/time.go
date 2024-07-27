package supports

import "time"

type TimeSupport struct{}

func (t *TimeSupport) FirstDayOfMonth(targetTime time.Time) time.Time {
	currentYear, currentMonth, _ := targetTime.Date()
	currentLocation := targetTime.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)

	return firstOfMonth
}
