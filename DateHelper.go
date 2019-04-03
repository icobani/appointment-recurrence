package appointment_recurrence

import "time"

var secondsTROfUTC = int((3 * time.Hour).Seconds())
var tr_Location = time.FixedZone("Turkey Time", secondsTROfUTC)

func NextFirstDay(year int, month time.Month, day time.Weekday, hour int, minute int, AddMonth int) time.Time {
	tDay := time.Date(
		year,
		month,
		1,
		hour,
		minute,
		0,
		0,
		tr_Location)

	for i := 1; i < 8; i++ {
		if tDay.Weekday() == day {
			break
		}
		tDay = tDay.AddDate(0, 0, 1)
	}

	if tDay.Before(time.Now()) {

		tDay = time.Date(
			year,
			month,
			1,
			hour,
			minute,
			0,
			0,
			tr_Location).AddDate(0, AddMonth, 0)
		for i := 1; i < 8; i++ {
			if tDay.Weekday() == day {
				break
			}
			tDay = tDay.AddDate(0, 0, 1)
		}
	}
	return tDay
}

func FirstWeekday(year int, month time.Month, hour int, minute int, AddMonth int) time.Time {

	tDay := time.Date(
		year,
		month,
		1,
		0,
		0,
		0,
		0,
		tr_Location)

	tDay = time.Date(
		tDay.Year(),
		tDay.Month(),
		1,
		hour,
		minute,
		0,
		0,
		tr_Location)

	tDay = time.Date(
		tDay.Year(),
		tDay.Month(),
		(8-int(tDay.Weekday()))%7+1,
		hour,
		minute,
		0,
		0,
		tr_Location)

	if tDay.Before(time.Now()) {

		tDay = time.Date(
			tDay.Year(),
			tDay.Month(),
			1,
			hour,
			minute,
			0,
			0,
			tr_Location).AddDate(0, AddMonth, 0)

		tDay = time.Date(
			tDay.Year(),
			tDay.Month(),
			(8-int(tDay.Weekday()))%7+1,
			hour,
			minute,
			0,
			0,
			tr_Location)
	}
	return tDay
}

func FirstWeekendDay(year int, month time.Month, hour int, minute int, AddMonth int) time.Time {

	return FirstWeekday(year, month, hour, minute, AddMonth).AddDate(0, 0, 5)
}
