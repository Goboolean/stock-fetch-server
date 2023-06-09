package buycycle

import "time"



var (
	location = "Asia/Seoul"
	startHour = 9
	startMinute = 30
	endHour = 15
	endMinute = 30
)


func IsMarketOn() bool {

	location, _ := time.LoadLocation(location)
	now := time.Now().In(location)

	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return false
	}

	start := time.Date(now.Year(), now.Month(), now.Day(), startHour, startMinute, 0, 0, now.Location())
	end := time.Date(now.Year(), now.Month(), now.Day(), endHour, endMinute, 0, 0, now.Location())
	
	if !(now.After(start) && now.Before(end)) {
		return false
	}

	return true
}