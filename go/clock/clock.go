package clock

import "fmt"

const TestVersion = 2
const minutesInHour = 60
const hoursInDay = 24
const minutesInDay = minutesInHour * hoursInDay

type Clock struct {
	h, m int
}

func Time(hours, minutes int) Clock {
	formattedHours, formattedMinutes := normalizeTime(hours, minutes)

	return Clock{formattedHours, formattedMinutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(a int) Clock {
	newHour, newMinute := normalizeTime(c.h, c.m+a)
	return Clock{newHour, newMinute}
}

func normalizeTime(hours, minutes int) (formattedHours, formattedMinutes int) {
	totalMinutes := hours*minutesInHour + minutes
	for totalMinutes < 0 {
		totalMinutes += minutesInDay
	}
	formattedHours = totalMinutes / minutesInHour % hoursInDay
	formattedMinutes = totalMinutes % minutesInHour
	return
}
