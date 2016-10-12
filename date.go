package dateparser

import (
    // "fmt"
    "time"
)

type Date struct {
    year int
    month int
    day int
    hour int
    minute int
    second int
    weekday int
    tz *time.Location
}

func (d *Date) AddDefault(t *time.Time) *Date {
    if d.year == 0 {
        d.year = t.Year()
    }
    if d.month == 0 {
        d.month = int(t.Month())
    }
    if d.day == 0 {
        d.day = t.Day()
    }
    if d.hour == 0 {
        d.hour = t.Hour()
    }
    if d.minute == 0 {
        d.minute = t.Minute()
    }
    if d.second == 0 {
        d.second = t.Second()
    }
    if d.tz == nil {
        d.tz = t.Location()
    }

    return d
}

func (d *Date) ToTime() time.Time {
    loc := d.tz
    if loc == nil {
        loc = time.UTC
    }

    return time.Date(d.year, time.Month(d.month), d.day, d.hour, d.minute, d.second, 0, loc)
}

func (d *Date) String() string {
    return d.ToTime().String()
}