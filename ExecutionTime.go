package appointment_recurrence

import (
	"log"
	"time"
	)


func (this *ExecutionTimeModel) calcNextExecutionTimeForWeekly() {
	if this.RecurrencePattern != E_RecurrencePatternWeekly {
		log.Println("Haftalık hesapla tanımlı değil")
		return
	}

	var tDay time.Time
	if this.StartDate.After(time.Now()) {
		tDay = time.Date(this.StartDate.Year(),
			this.StartDate.Month(),
			this.StartDate.Day(), 0, 0, 0, 0,tr_Location)
	} else {
		tDay = time.Now()
	}

	if this.EndDate.Before(time.Now()) && !this.NoEndDate {
		log.Println("Bitiş tarihi geçti.", this.EndDate)
		return
	}

timeLoop:
	for i := 1; i <= 7; i++ {
		switch tDay.Weekday() {
		case time.Sunday:
			if this.OnSunday {
				break timeLoop
			}
		case time.Monday:
			if this.OnMonday {
				break timeLoop
			}
		case time.Tuesday:
			if this.OnTuesday {
				break timeLoop
			}
		case time.Wednesday:
			if this.OnWednesday {
				break timeLoop
			}
		case time.Thursday:
			if this.OnThursday {
				break timeLoop
			}
		case time.Friday:
			if this.OnFriday {
				break timeLoop
			}
		case time.Saturday:
			if this.OnSaturday {
				break timeLoop
			}
		}
		tDay = tDay.AddDate(0, 0, 1)
	}

	tDay = time.Date(
		tDay.Year(),
		tDay.Month(),
		tDay.Day(),
		int(this.StartTimeHour),
		int(this.StartTimeMinute),
		0,
		0,
		tr_Location)

	if tDay.Before(time.Now()) {
		tDay = tDay.AddDate(0, 0, 7)
	}
	this.NextExecutingTime = tDay
}

func (this *ExecutionTimeModel) calcNextExecutionTimeForDaily() {
	if this.RecurrencePattern != E_RecurrencePatternDaily {
		log.Println("Haftalık hesapla tanımlı değil")
		return
	}

	var tDay time.Time
	if this.StartDate.After(time.Now()) {
		log.Println("Başlangıç tarihi henüz gelmemiş", this.StartDate)
		tDay = time.Date(this.StartDate.Year(),
			this.StartDate.Month(),
			this.StartDate.Day(), 0, 0, 0, 0, tr_Location)
	} else {
		tDay = time.Now()
	}

	if this.EndDate.Before(time.Now()) && !this.NoEndDate {
		log.Println("Bitiş tarihi geçti.", this.EndDate)
		return
	}

	tDay = time.Date(
		tDay.Year(),
		tDay.Month(),
		tDay.Day(),
		int(this.StartTimeHour),
		int(this.StartTimeMinute),
		0,
		0,
		tr_Location)

	if tDay.Before(time.Now()) {
		tDay = tDay.AddDate(0, 0, 1)
	}
	this.NextExecutingTime = tDay
}

func (this *ExecutionTimeModel) calcNextExecutionTimeForMonthly() {
	if this.RecurrencePattern != E_RecurrencePatternMonthly {
		log.Println("Aylık hesapla tanımlı değil")
		return
	}

	var tDay time.Time
	if this.StartDate.After(time.Now()) {
		log.Println("Başlangıç tarihi henüz gelmemiş", this.StartDate)
		tDay = time.Date(this.StartDate.Year(),
			this.StartDate.Month(),
			this.StartDate.Day(), 0, 0, 0, 0, tr_Location)
	} else {
		tDay = time.Now()
	}

	if this.EndDate.Before(time.Now()) && !this.NoEndDate {
		log.Println("Bitiş tarihi geçti.", this.EndDate)
		return
	}

	switch this.On {
	case E_OnDay:
		tDay = time.Date(
			tDay.Year(),
			tDay.Month(),
			int(this.Every),
			int(this.StartTimeHour),
			int(this.StartTimeMinute),
			0,
			0,
			tr_Location)

		if tDay.Before(time.Now()) {
			tDay = tDay.AddDate(0, 1, 0)
		}

		break
	case E_OnFirst:
		switch this.OnValue {
		case E_OnValueDay: // Ayın günü
			tDay = time.Date(
				tDay.Year(),
				tDay.Month(),
				1,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
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
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1)

			break
		case E_OnValueWeekendday:

			tDay = FirstWeekendDay(
				tDay.Year(),
				tDay.Month(),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1)

			break
		case E_OnValueSunday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Sunday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Monday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Tuesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Wednesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Thursday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Friday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Saturday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1)

			break
		}
		break
	case E_OnSecond:
		switch this.OnValue {
		case E_OnValueDay: // Ayın ilk iş günü
			tDay = time.Date(
				tDay.Year(),
				tDay.Month(),
				2,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
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
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueWeekendday:

			tDay = FirstWeekendDay(
				tDay.Year(),
				tDay.Month(),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueSunday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Sunday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Monday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Tuesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Wednesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Thursday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Friday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Saturday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7)

			break
		}
		break
	case E_OnThird:
		switch this.OnValue {
		case E_OnValueDay: // Ayın ilk iş günü
			tDay = time.Date(
				tDay.Year(),
				tDay.Month(),
				3,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
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
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueWeekendday:

			tDay = FirstWeekendDay(
				tDay.Year(),
				tDay.Month(),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueSunday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Sunday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Monday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Tuesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Wednesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Thursday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Friday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Saturday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*2)

			break
		}
		break
	case E_OnFourth:
		switch this.OnValue {
		case E_OnValueDay: // Ayın ilk iş günü
			tDay = time.Date(
				tDay.Year(),
				tDay.Month(),
				4,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
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
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueWeekendday:

			tDay = FirstWeekendDay(
				tDay.Year(),
				tDay.Month(),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueSunday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Sunday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Monday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Tuesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Wednesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Thursday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Friday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				tDay.Month(),
				time.Saturday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				1).AddDate(0, 0, 7*3)

			break
		}
		break
	}

	this.NextExecutingTime = tDay
}

func (this *ExecutionTimeModel) calcNextExecutionTimeForYearly() {
	if this.RecurrencePattern != E_RecurrencePatternYearly {
		log.Println("Yıllık hesapla tanımlı değil")
		return
	}

	var tDay time.Time
	if this.StartDate.After(time.Now()) {
		log.Println("Başlangıç tarihi henüz gelmemiş", this.StartDate)
		tDay = time.Date(this.StartDate.Year(),
			this.StartDate.Month(),
			this.StartDate.Day(), 0, 0, 0, 0, tr_Location)
	} else {
		tDay = time.Now()
	}

	if this.EndDate.Before(time.Now()) && !this.NoEndDate {
		log.Println("Bitiş tarihi geçti.", this.EndDate)
		return
	}

	switch this.On {
	case E_OnDay:

		tDay = time.Date(
			tDay.Year(),
			time.Month(this.EveryMonthValue),
			int(this.Every),
			int(this.StartTimeHour),
			int(this.StartTimeMinute),
			0,
			0,
			tDay.Location(),
		)

		if tDay.Before(time.Now()) {
			tDay = tDay.AddDate(1, 0, 0)
		}

		break
	case E_OnFirst:
		switch this.OnValue {
		case E_OnValueDay: // Ayın ilk günü
			tDay = time.Date(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				1,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
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
				time.Month(this.EveryMonthValue),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12)

			break
		case E_OnValueWeekendday: // ilk Haftasonu

			tDay = FirstWeekendDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12)

			break
		case E_OnValueSunday: // ilk pazarı

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Sunday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Monday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Tuesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Wednesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Thursday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Friday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Saturday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12)

			break
		}
		break
	case E_OnSecond:
		switch this.OnValue {
		case E_OnValueDay: // Ayın ikinci günü
			tDay = time.Date(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				2,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
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
				time.Month(this.EveryMonthValue),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueWeekendday: // ilk Haftasonu

			tDay = FirstWeekendDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueSunday: // ilk pazarı

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Sunday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Monday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Tuesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Wednesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Thursday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Friday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Saturday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*1)

			break
		}
		break
	case E_OnThird:
		switch this.OnValue {
		case E_OnValueDay: // Ayın ikinci günü
			tDay = time.Date(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				3,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
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
				time.Month(this.EveryMonthValue),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueWeekendday: // ilk Haftasonu

			tDay = FirstWeekendDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueSunday: // ilk pazarı

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Sunday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Monday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Tuesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Wednesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Thursday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Friday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Saturday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*2)

			break
		}
		break
	case E_OnFourth:
		switch this.OnValue {
		case E_OnValueDay: // Ayın ikinci günü
			tDay = time.Date(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				4,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
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
				time.Month(this.EveryMonthValue),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueWeekendday: // ilk Haftasonu

			tDay = FirstWeekendDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueSunday: // ilk pazarı

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Sunday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueMonday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Monday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueTuesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Tuesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueWednesday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Wednesday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueThursday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Thursday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueFriday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Friday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		case E_OnValueSaturday:

			tDay = NextFirstDay(
				tDay.Year(),
				time.Month(this.EveryMonthValue),
				time.Saturday,
				int(this.StartTimeHour),
				int(this.StartTimeMinute),
				12).AddDate(0, 0, 7*3)

			break
		}
		break
	}

	this.NextExecutingTime = tDay
}

func (this *ExecutionTimeModel) Calc() {
	switch this.RecurrencePattern {
	case E_RecurrencePatternDaily:
		this.calcNextExecutionTimeForDaily()
		break
	case E_RecurrencePatternWeekly:
		this.calcNextExecutionTimeForWeekly()
		break
	case E_RecurrencePatternMonthly:
		this.calcNextExecutionTimeForMonthly()
		break
	case E_RecurrencePatternYearly:
		this.calcNextExecutionTimeForYearly()
		break
	}
}

