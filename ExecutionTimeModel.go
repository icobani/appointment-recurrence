package appointment_recurrence

import "time"

type enumRecurrencePattern string

var (
	E_RecurrencePatternDaily   enumRecurrencePattern = "Daily"
	E_RecurrencePatternWeekly  enumRecurrencePattern = "Weekly"
	E_RecurrencePatternMonthly enumRecurrencePattern = "Monthly"
	E_RecurrencePatternYearly  enumRecurrencePattern = "Yearly"
	E_RecurrencePatternOneTime enumRecurrencePattern = "OneTime"
)

type enumOn string

var (
	E_OnDay    enumOn = "Day"
	E_OnFirst  enumOn = "First"
	E_OnSecond enumOn = "Second"
	E_OnThird  enumOn = "Third"
	E_OnFourth enumOn = "Fourth"
)

type enumOnValue string

var (
	E_OnValueDay        enumOnValue = "Day"
	E_OnValueWeekday    enumOnValue = "Weekday"
	E_OnValueWeekendday enumOnValue = "Weekend day"
	E_OnValueSunday     enumOnValue = "Sunday"
	E_OnValueMonday     enumOnValue = "Monday"
	E_OnValueTuesday    enumOnValue = "Tuesday"
	E_OnValueWednesday  enumOnValue = "Wednesday"
	E_OnValueThursday   enumOnValue = "Thursday"
	E_OnValueFriday     enumOnValue = "Friday"
	E_OnValueSaturday   enumOnValue = "Saturday"
)

type ExecutionTimeModel struct {
	RecurrencePattern       enumRecurrencePattern `json:"recurrence_pattern"`
	Every                   int8                  `json:"every"`
	StartDate               time.Time             `json:"start_date"`
	EndDate                 time.Time             `json:"end_date"`
	NoEndDate               bool                  `json:"no_end_date"`
	StartTimeHour           int8                  `json:"start_time_hour"`
	StartTimeMinute         int8                  `json:"start_time_minute"`
	OnSunday                bool                  `json:"on_sunday"`
	OnMonday                bool                  `json:"on_monday"`
	OnTuesday               bool                  `json:"on_tuesday"`
	OnWednesday             bool                  `json:"on_wednesday"`
	OnThursday              bool                  `json:"on_thursday"`
	OnFriday                bool                  `json:"on_friday"`
	OnSaturday              bool                  `json:"on_saturday"`
	On                      enumOn                `json:"on"`
	OnValue                 enumOnValue           `json:"on_value"`
	EveryMonthValue         int8                  `json:"every_month_value"`
	NextExecutingTime       time.Time             `json:"next_executing_time"`
	LastExecutingTime       time.Time             `json:"last_executing_time"`
	Active                  bool                  `json:"active"`
}

