package dateparser

import (
    // "fmt"
    "time"
)

var Parser = &Pattern{}
func Parse(b []byte, def *time.Time) *Date {
    return Parser.Parse(b, def)
}

var _ = Parser.Add().
    Match("Mon").
    Handle(func (d *Date, ts []*Token) bool {
        d.weekday = ts[0].Number()
        return true
    })

var _ = Parser.Add().
    Match("MST").
    Handle(func (d *Date, ts []*Token) bool {
        loc, err := time.LoadLocation(ts[0].V)
        if err != nil {
            return false
        }

        d.tz = loc
        return true
    })

var _ = Parser.Add().
    Match("15 hours").
    Handle(func (d *Date, ts []*Token) bool {
        d.hour = ts[0].Number()
        return true
    })

var _ = Parser.Add().
    Match("03 pm").
    Handle(func (d *Date, ts []*Token) bool {
        d.hour = ts[0].Number()
        return true
    })

var _ = Parser.Add().
    Match("04 mins").
    Handle(func (d *Date, ts []*Token) bool {
        d.minute = ts[0].Number()
        return true
    })

var _ = Parser.Add().
    Match("05 secs").
    Handle(func (d *Date, ts []*Token) bool {
        d.second = ts[0].Number()
        return true
    })

// 10:24:05
var _ = Parser.Add().
    Match("15:04:05").
    Handle(func (d *Date, ts []*Token) bool {
        d.hour = ts[0].Number()
        d.minute = ts[2].Number()
        d.second = ts[4].Number()
        return true
    })

var _ = Parser.Add().
    Match("15:04").
    Handle(func (d *Date, ts []*Token) bool {
        if d.hour == 0 {
            d.hour = ts[0].Number()
            d.minute = ts[2].Number()
        } else {
            d.minute = ts[0].Number()
            d.second = ts[2].Number()
        }

        return true
    })

var _ = Parser.Add().
    Match("15:04 pm").
    Handle(func (d *Date, ts []*Token) bool {
        if d.hour == 0 {
            d.hour = ts[0].Number()
            d.minute = ts[2].Number()
        } else {
            d.minute = ts[0].Number()
            d.second = ts[2].Number()
        }

        return true
    })

var _ = Parser.Add().
    Match("2006/01/02").
    Handle(func (d *Date, ts []*Token) bool {
        d.year = ts[0].Number()
        d.month = ts[2].Number()
        d.day = ts[4].Number()
        return true
    })

var _ = Parser.Add().
    Match("2006 01 02").
    Handle(func (d *Date, ts []*Token) bool {
        d.year = ts[0].Number()
        d.month = ts[1].Number()
        d.day = ts[2].Number()
        return true
    })

var _ = Parser.Add().
    Match("02/01/2006").
    Handle(func (d *Date, ts []*Token) bool {
        d.day = ts[0].Number()
        d.month = ts[2].Number()
        d.year = ts[4].Number()
        return true
    })

var _ = Parser.Add().
    Match("02 01 2006").
    Handle(func (d *Date, ts []*Token) bool {
        d.day = ts[0].Number()
        d.month = ts[1].Number()
        d.year = ts[2].Number()
        return true
    })

var _ = Parser.Add().
    Match("01/02/2006").
    Handle(func (d *Date, ts []*Token) bool {
        d.month = ts[0].Number()
        d.day = ts[2].Number()
        d.year = ts[4].Number()
        return true
    })

var _ = Parser.Add().
    Match("01 02 2006").
    Handle(func (d *Date, ts []*Token) bool {
        d.month = ts[0].Number()
        d.day = ts[1].Number()
        d.year = ts[2].Number()
        return true
    })

var _ = Parser.Add().
    Match("02/01/06").
    Handle(func (d *Date, ts []*Token) bool {
        d.day = ts[0].Number()
        d.month = ts[2].Number()
        d.year = ts[4].NumberYear()
        return true
    })

var _ = Parser.Add().
    Match("Jan").
    Handle(func (d *Date, ts []*Token) bool {
        d.month = ts[0].Number()
        return true
    })

var _ = Parser.Add().
    Match("-0700").
    Handle(func (d *Date, ts []*Token) bool {

        // d.tzoffset = ts[0].Number() + ts[1].V
        return true
    })

var _ = Parser.Add().
    Match("-07:00").
    Handle(func (d *Date, ts []*Token) bool {
        offset := ts[1].Number() * 3600 + ts[3].Number() * 60

        if ts[0].V == "-" {
            offset *= -1
        }

        d.tz = time.FixedZone("", offset)
        return true
    })

var _ = Parser.Add().
    Match("2006").
    Handle(func (d *Date, ts []*Token) bool {
        if d.year != 0 {
            return false
        }

        d.year = ts[0].Number()
        return true
    })

var _ = Parser.Add().
    Match("02").
    Handle(func (d *Date, ts []*Token) bool {
        if d.day != 0 {
            return false
        }

        d.day = ts[0].Number()
        return true
    })

var _ = Parser.Add().
    Match("01").
    Handle(func (d *Date, ts []*Token) bool {
        if d.month != 0 {
            return false
        }

        d.month = ts[0].Number()
        return true
    })

var _ = Parser.Add().
    Match("06").
    Handle(func (d *Date, ts []*Token) bool {
        if d.year != 0 {
            return false
        }

        d.year = ts[0].NumberYear()
        return true
    })

var _ = Parser.Add().
    Match("15").
    Handle(func (d *Date, ts []*Token) bool {
        if d.hour != 0 {
            return false
        }

        d.hour = ts[0].Number()
        return true
    })








