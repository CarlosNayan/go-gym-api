package utils

import (
	"errors"
	"time"
)

type Moment struct {
	date time.Time
}

func NewMoment(date ...interface{}) (*Moment, error) {
	var t time.Time
	if len(date) == 0 {
		t = time.Now()
	} else {
		switch v := date[0].(type) {
		case nil:
			t = time.Now()
		case time.Time:
			t = v
		case string:
			var err error
			t, err = time.Parse(time.RFC3339, v)
			if err != nil {
				return nil, errors.New("data inválida")
			}
		case int64:
			t = time.Unix(v, 0)
		default:
			return nil, errors.New("tipo de data inválido")
		}
	}
	return &Moment{date: t}, nil
}

func Time(date time.Time) *Moment {
	return &Moment{date: date}
}

func (m *Moment) StartOf(unit string) *Moment {
	date := m.date.UTC()
	switch unit {
	case "day":
		date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	case "month":
		date = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	case "year":
		date = time.Date(date.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	case "week":
		weekday := int(date.Weekday())
		date = date.AddDate(0, 0, -weekday)
		date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	}
	return &Moment{date: date}
}

func (m *Moment) EndOf(unit string) *Moment {
	date := m.date.UTC()
	switch unit {
	case "day":
		date = time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, time.UTC)
	case "month":
		date = time.Date(date.Year(), date.Month()+1, 0, 23, 59, 59, 999999999, time.UTC)
	case "year":
		date = time.Date(date.Year()+1, 1, 0, 23, 59, 59, 999999999, time.UTC)
	case "week":
		weekday := int(date.Weekday())
		date = date.AddDate(0, 0, 6-weekday)
		date = time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, time.UTC)
	}
	return &Moment{date: date}
}

func (m *Moment) Format(formatString ...string) string {
	defaultFormat := time.RFC3339
	if len(formatString) > 0 {
		return m.date.Format(formatString[0])
	}
	return m.date.Format(defaultFormat)
}

func (m *Moment) Add(amount int, unit string) *Moment {
	date := m.date
	switch unit {
	case "days":
		date = date.AddDate(0, 0, amount)
	case "months":
		date = date.AddDate(0, amount, 0)
	case "years":
		date = date.AddDate(amount, 0, 0)
	case "hours":
		date = date.Add(time.Duration(amount) * time.Hour)
	case "minutes":
		date = date.Add(time.Duration(amount) * time.Minute)
	}
	return &Moment{date: date}
}

func (m *Moment) Subtract(amount int, unit string) *Moment {
	return m.Add(-amount, unit)
}

func (m *Moment) UtcOffset(offset int) *Moment {
	date := m.date.Add(time.Duration(offset) * time.Minute)
	return &Moment{date: date}
}

func (m *Moment) ToDate() time.Time {
	return m.date
}

func (m *Moment) IsBefore(other *Moment) bool {
	return m.date.Before(other.date)
}

func (m *Moment) IsAfter(other *Moment) bool {
	return m.date.After(other.date)
}

func (m *Moment) IsSame(other *Moment) bool {
	return m.date.Equal(other.date)
}

func (m *Moment) Diff(other *Moment, unit string) int {
	diff := m.date.Sub(other.date)
	switch unit {
	case "minutes":
		return int(diff.Minutes())
	case "hours":
		return int(diff.Hours())
	case "days":
		return int(diff.Hours() / 24)
	case "months":
		years := other.date.Year() - m.date.Year()
		months := int(other.date.Month()) - int(m.date.Month())
		return years*12 + months
	case "years":
		return other.date.Year() - m.date.Year()
	default:
		return 0
	}
}

func (m *Moment) Weekday(day *int) (*Moment, int) {
	if day == nil {
		return nil, int(m.date.Weekday())
	}
	if *day < 0 || *day > 6 {
		panic("O dia deve estar entre 0 (domingo) e 6 (sábado)")
	}
	currentDay := int(m.date.Weekday())
	diff := *day - currentDay
	newDate := m.date.AddDate(0, 0, diff)
	return &Moment{date: newDate}, -1
}
