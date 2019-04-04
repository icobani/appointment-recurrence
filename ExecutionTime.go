package appointment_recurrence

import (
	"log"
	"time"
)

func (this *AppointmentRecurrence) calcNextExecutionTimeForWeekly() (time.Time, *ErrorStruct) {
	if this.ETimeModel.RecurrencePattern != E_RecurrencePatternWeekly {
		// log.Println("Haftalık hesapla tanımlı değil")
		errTime, _ := time.Parse("2006-01-02", "0001-01-01")
		return errTime, &ErrorStruct{
			"Haftalık hesapla tanımlı değil",
		}
	}

	var tDay time.Time
	if this.ETimeModel.StartDate.After(time.Now()) {
		tDay = time.Date(this.ETimeModel.StartDate.Year(),
			this.ETimeModel.StartDate.Month(),
			this.ETimeModel.StartDate.Day(), 0, 0, 0, 0, tr_Location)
	} else {
		tDay = time.Now()
	}

	if this.ETimeModel.EndDate.Before(time.Now()) && !this.ETimeModel.NoEndDate {
		// log.Println("Bitiş tarihi geçti.", this.ETimeModel.EndDate)
		errTime, _ := time.Parse("2006-01-02", "0001-01-01")
		return errTime, &ErrorStruct{
			"Bitiş tarihi geçti",
		}
	}

timeLoop:
	for i := 1; i <= 7; i++ {
		switch tDay.Weekday() {
		case time.Sunday:
			if this.ETimeModel.OnSunday {
				break timeLoop
			}
		case time.Monday:
			if this.ETimeModel.OnMonday {
				break timeLoop
			}
		case time.Tuesday:
			if this.ETimeModel.OnTuesday {
				break timeLoop
			}
		case time.Wednesday:
			if this.ETimeModel.OnWednesday {
				break timeLoop
			}
		case time.Thursday:
			if this.ETimeModel.OnThursday {
				break timeLoop
			}
		case time.Friday:
			if this.ETimeModel.OnFriday {
				break timeLoop
			}
		case time.Saturday:
			if this.ETimeModel.OnSaturday {
				break timeLoop
			}
		}
		tDay = tDay.AddDate(0, 0, 1)
	}

	tDay = time.Date(
		tDay.Year(),
		tDay.Month(),
		tDay.Day(),
		int(this.ETimeModel.StartTimeHour),
		int(this.ETimeModel.StartTimeMinute),
		0,
		0,
		tr_Location)

	if tDay.Before(time.Now()) {
		tDay = tDay.AddDate(0, 0, 7)
	}
	this.ETimeModel.NextExecutingTime = tDay

	return tDay, nil
}

func (this *AppointmentRecurrence) calcNextExecutionTimeForDaily() (time.Time, *ErrorStruct) {
	if this.ETimeModel.RecurrencePattern != E_RecurrencePatternDaily {
		// log.Println("Haftalık hesapla tanımlı değil")
		errTime, _ := time.Parse("2006-01-02", "0001-01-01")
		return errTime, &ErrorStruct{
			"Haftalık hesapla tanımlı değil",
		}
	}

	var tDay time.Time
	if this.ETimeModel.StartDate.After(time.Now()) {
		log.Println("Başlangıç tarihi henüz gelmemiş", this.ETimeModel.StartDate)
		tDay = time.Date(this.ETimeModel.StartDate.Year(),
			this.ETimeModel.StartDate.Month(),
			this.ETimeModel.StartDate.Day(), 0, 0, 0, 0, tr_Location)
	} else {
		tDay = time.Now()
	}

	if this.ETimeModel.EndDate.Before(time.Now()) && !this.ETimeModel.NoEndDate {
		// log.Println("Bitiş tarihi geçti.", this.ETimeModel.EndDate)
		errTime, _ := time.Parse("2006-01-02", "0001-01-01")
		return errTime, &ErrorStruct{
			"Bitiş tarihi geçti",
		}
	}

	tDay = time.Date(
		tDay.Year(),
		tDay.Month(),
		tDay.Day(),
		int(this.ETimeModel.StartTimeHour),
		int(this.ETimeModel.StartTimeMinute),
		0,
		0,
		tr_Location)

	if tDay.Before(time.Now()) {
		tDay = tDay.AddDate(0, 0, 1)
	}
	this.ETimeModel.NextExecutingTime = tDay

	return tDay, nil
}

func (this *AppointmentRecurrence) calcNextExecutionTimeForMonthly() (time.Time, *ErrorStruct) {
	if this.ETimeModel.RecurrencePattern != E_RecurrencePatternMonthly {
		// log.Println("Aylık hesapla tanımlı değil")
		errTime, _ := time.Parse("2006-01-02", "0001-01-01")
		return errTime, &ErrorStruct{
			"Aylık hesapla tanımlı değil",
		}
	}

	var tDay time.Time
	if this.ETimeModel.StartDate.After(time.Now()) {
		log.Println("Başlangıç tarihi henüz gelmemiş", this.ETimeModel.StartDate)
		tDay = time.Date(this.ETimeModel.StartDate.Year(),
			this.ETimeModel.StartDate.Month(),
			this.ETimeModel.StartDate.Day(), 0, 0, 0, 0, tr_Location)
	} else {
		tDay = time.Now()
	}

	if this.ETimeModel.EndDate.Before(time.Now()) && !this.ETimeModel.NoEndDate {
		// log.Println("Bitiş tarihi geçti.", this.ETimeModel.EndDate)
		errTime, _ := time.Parse("2006-01-02", "0001-01-01")
		return errTime, &ErrorStruct{
			"Bitiş tarihi geçti",
		}
	}

	switch this.ETimeModel.On {
	case E_OnDay:
		tDay = time.Date(
			tDay.Year(),
			tDay.Month(),
			int(this.ETimeModel.Every),
			int(this.ETimeModel.StartTimeHour),
			int(this.ETimeModel.StartTimeMinute),
			0,
			0,
			tr_Location)

		if tDay.Before(time.Now()) {
			tDay = tDay.AddDate(0, 1, 0)
		}

		break
	case E_OnFirst:
		switch this.ETimeModel.OnValue {
		case E_OnValueDay: // Ayın günü
			tDay = time.Date(
				tDay.Year(),
				tDay.Month(),
				1,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				0,
				0,
				tr_Location)
			if tDay.Before(time.Now()) {
				tDay = tDay.AddDate(0, 1, 0)
			}
			break
		case E_OnValueWeekday:

			tDay = FirstWeekday(
				tDay.Year(),
				tDay.Month(),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1)

			break
		case E_OnValueWeekendday:

			tDay = FirstWeekendDay(
				tDay.Year(),
				tDay.Month(),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1)

			break
		case E_OnValueSunday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Sunday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Monday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Tuesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Wednesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Thursday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Friday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Saturday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1)

			break
		}
		break
	case E_OnSecond:
		switch this.ETimeModel.OnValue {
		case E_OnValueDay: // Ayın ilk iş günü
			tDay = time.Date(
				tDay.Year(),
				tDay.Month(),
				2,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				0,
				0,
				tr_Location)
			if tDay.Before(time.Now()) {
				tDay = tDay.AddDate(0, 1, 0)
			}
			break
		case E_OnValueWeekday:

			tDay = FirstWeekday(
				tDay.Year(),
				tDay.Month(),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueWeekendday:

			tDay = FirstWeekendDay(
				tDay.Year(),
				tDay.Month(),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueSunday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Sunday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Monday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Tuesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Wednesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Thursday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Friday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Saturday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		}
		break
	case E_OnThird:
		switch this.ETimeModel.OnValue {
		case E_OnValueDay: // Ayın ilk iş günü
			tDay = time.Date(
				tDay.Year(),
				tDay.Month(),
				3,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				0,
				0,
				tr_Location)
			if tDay.Before(time.Now()) {
				tDay = tDay.AddDate(0, 1, 0)
			}
			break
		case E_OnValueWeekday:

			tDay = FirstWeekday(
				tDay.Year(),
				tDay.Month(),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueWeekendday:

			tDay = FirstWeekendDay(
				tDay.Year(),
				tDay.Month(),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueSunday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Sunday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Monday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Tuesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Wednesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Thursday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Friday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Saturday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		}
		break
	case E_OnFourth:
		switch this.ETimeModel.OnValue {
		case E_OnValueDay: // Ayın ilk iş günü
			tDay = time.Date(
				tDay.Year(),
				tDay.Month(),
				4,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				0,
				0,
				tr_Location)
			if tDay.Before(time.Now()) {
				tDay = tDay.AddDate(0, 1, 0)
			}
			break
		case E_OnValueWeekday:

			tDay = FirstWeekday(
				tDay.Year(),
				tDay.Month(),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueWeekendday:

			tDay = FirstWeekendDay(
				tDay.Year(),
				tDay.Month(),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueSunday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Sunday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Monday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Tuesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Wednesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Thursday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Friday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Saturday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		}
		break
	}

	this.ETimeModel.NextExecutingTime = tDay

	return tDay, nil
}

func (this *AppointmentRecurrence) calcNextExecutionTimeForYearly() (time.Time, *ErrorStruct) {
	if this.ETimeModel.RecurrencePattern != E_RecurrencePatternYearly {
		// log.Println("Yıllık hesapla tanımlı değil")
		errTime, _ := time.Parse("2006-01-02", "0001-01-01")
		return errTime, &ErrorStruct{
			"Yıllık hesapla tanımlı değil",
		}
	}

	var tDay time.Time
	if this.ETimeModel.StartDate.After(time.Now()) {
		log.Println("Başlangıç tarihi henüz gelmemiş", this.ETimeModel.StartDate)
		tDay = time.Date(this.ETimeModel.StartDate.Year(),
			this.ETimeModel.StartDate.Month(),
			this.ETimeModel.StartDate.Day(), 0, 0, 0, 0, tr_Location)
	} else {
		tDay = time.Now()
	}

	if this.ETimeModel.EndDate.Before(time.Now()) && !this.ETimeModel.NoEndDate {
		// log.Println("Bitiş tarihi geçti.", this.ETimeModel.EndDate)
		errTime, _ := time.Parse("2006-01-02", "0001-01-01")
		return errTime, &ErrorStruct{
			"Bitiş tarihi geçti",
		}
	}

	switch this.ETimeModel.On {
	case E_OnDay:

		tDay = time.Date(
			tDay.Year(),
			time.Month(this.ETimeModel.EveryMonthValue),
			int(this.ETimeModel.Every),
			int(this.ETimeModel.StartTimeHour),
			int(this.ETimeModel.StartTimeMinute),
			0,
			0,
			tDay.Location(),
		)

		if tDay.Before(time.Now()) {
			tDay = tDay.AddDate(1, 0, 0)
		}

		break
	case E_OnFirst:
		switch this.ETimeModel.OnValue {
		case E_OnValueDay: // Ayın ilk günü
			tDay = time.Date(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				1,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				0,
				0,
				tr_Location)
			if tDay.Before(time.Now()) {
				tDay = tDay.AddDate(1, 0, 0)
			}
			break
		case E_OnValueWeekday: // ilk iş günü

			tDay = FirstWeekday(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12)

			break
		case E_OnValueWeekendday: // ilk Haftasonu

			tDay = FirstWeekendDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12)

			break
		case E_OnValueSunday: // ilk pazarı

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Sunday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Monday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Tuesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Wednesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Thursday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Friday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Saturday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12)

			break
		}
		break
	case E_OnSecond:
		switch this.ETimeModel.OnValue {
		case E_OnValueDay: // Ayın ikinci günü
			tDay = time.Date(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				2,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				0,
				0,
				tr_Location)
			if tDay.Before(time.Now()) {
				tDay = tDay.AddDate(1, 0, 0)
			}
			break
		case E_OnValueWeekday: // ikinci iş günü

			tDay = FirstWeekday(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueWeekendday: // ilk Haftasonu

			tDay = FirstWeekendDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueSunday: // ilk pazarı

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Sunday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Monday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Tuesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Wednesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Thursday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Friday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Saturday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		}
		break
	case E_OnThird:
		switch this.ETimeModel.OnValue {
		case E_OnValueDay: // Ayın ikinci günü
			tDay = time.Date(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				3,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				0,
				0,
				tr_Location)
			if tDay.Before(time.Now()) {
				tDay = tDay.AddDate(1, 0, 0)
			}
			break
		case E_OnValueWeekday: // ikinci iş günü

			tDay = FirstWeekday(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueWeekendday: // ilk Haftasonu

			tDay = FirstWeekendDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueSunday: // ilk pazarı

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Sunday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Monday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Tuesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Wednesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Thursday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Friday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Saturday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		}
		break
	case E_OnFourth:
		switch this.ETimeModel.OnValue {
		case E_OnValueDay: // Ayın ikinci günü
			tDay = time.Date(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				4,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				0,
				0,
				tr_Location)
			if tDay.Before(time.Now()) {
				tDay = tDay.AddDate(1, 0, 0)
			}
			break
		case E_OnValueWeekday: // ikinci iş günü

			tDay = FirstWeekday(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueWeekendday: // ilk Haftasonu

			tDay = FirstWeekendDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueSunday: // ilk pazarı

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Sunday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Monday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Tuesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Wednesday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Thursday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Friday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.ETimeModel.EveryMonthValue),
				time.Saturday,
				int(this.ETimeModel.StartTimeHour),
				int(this.ETimeModel.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		}
		break
	}

	this.ETimeModel.NextExecutingTime = tDay

	return tDay, nil
}

func (this *AppointmentRecurrence) Calc() (time.Time, *ErrorStruct) {
	switch this.ETimeModel.RecurrencePattern {
	case E_RecurrencePatternDaily:
		return this.calcNextExecutionTimeForDaily()
	case E_RecurrencePatternWeekly:
		return this.calcNextExecutionTimeForWeekly()
	case E_RecurrencePatternMonthly:
		return this.calcNextExecutionTimeForMonthly()
	case E_RecurrencePatternYearly:
		return this.calcNextExecutionTimeForYearly()
	default:
		errTime, _ := time.Parse("2006-01-02", "0001-01-01")
		return errTime, &ErrorStruct{
			"Bitiş tarihi geçti",
		}
	}
}
