package dateparser

var Parser = &Pattern{}
func Parse(b []byte) *Date {
    return Parser.Parse(b)
}

var _ = Parser.Add().
    Match("Mon").
    Handle(func (d *Date, ts []*Token) bool {
        d.weekday = ts[0].V
        return true
    })

var _ = Parser.Add().
    Match("MST").
    Handle(func (d *Date, ts []*Token) bool {
        d.tz = ts[0].V
        return true
    })

var _ = Parser.Add().
    Match("15 hours").
    Handle(func (d *Date, ts []*Token) bool {
        d.hour = ts[0].V
        return true
    })

var _ = Parser.Add().
    Match("03 pm").
    Handle(func (d *Date, ts []*Token) bool {
        d.hour = ts[0].V
        return true
    })

var _ = Parser.Add().
    Match("04 mins").
    Handle(func (d *Date, ts []*Token) bool {
        d.minute = ts[0].V
        return true
    })

var _ = Parser.Add().
    Match("05 secs").
    Handle(func (d *Date, ts []*Token) bool {
        d.second = ts[0].V
        return true
    })

// 10:24:05
var _ = Parser.Add().
    Match("15:04:05").
    Handle(func (d *Date, ts []*Token) bool {
        d.hour = ts[0].V
        d.minute = ts[2].V
        d.second = ts[4].V
        return true
    })

var _ = Parser.Add().
    Match("15:04").
    Handle(func (d *Date, ts []*Token) bool {
        if d.hour == "" {
            d.hour = ts[0].V
            d.minute = ts[2].V
        } else {
            d.minute = ts[0].V
            d.second = ts[2].V
        }

        return true
    })

var _ = Parser.Add().
    Match("15:04 pm").
    Handle(func (d *Date, ts []*Token) bool {
        if d.hour == "" {
            d.hour = ts[0].V
            d.minute = ts[2].V
        } else {
            d.minute = ts[0].V
            d.second = ts[2].V
        }

        return true
    })

var _ = Parser.Add().
    Match("2006/01/02").
    Handle(func (d *Date, ts []*Token) bool {
        d.year = ts[0].V
        d.month = ts[2].V
        d.day = ts[4].V
        return true
    })

var _ = Parser.Add().
    Match("2006 01 02").
    Handle(func (d *Date, ts []*Token) bool {
        d.year = ts[0].V
        d.month = ts[1].V
        d.day = ts[2].V
        return true
    })

var _ = Parser.Add().
    Match("02/01/2006").
    Handle(func (d *Date, ts []*Token) bool {
        d.day = ts[0].V
        d.month = ts[2].V
        d.year = ts[4].V
        return true
    })

var _ = Parser.Add().
    Match("02 01 2006").
    Handle(func (d *Date, ts []*Token) bool {
        d.day = ts[0].V
        d.month = ts[1].V
        d.year = ts[2].V
        return true
    })

var _ = Parser.Add().
    Match("01/02/2006").
    Handle(func (d *Date, ts []*Token) bool {
        d.month = ts[0].V
        d.day = ts[2].V
        d.year = ts[4].V
        return true
    })

var _ = Parser.Add().
    Match("01 02 2006").
    Handle(func (d *Date, ts []*Token) bool {
        d.month = ts[0].V
        d.day = ts[1].V
        d.year = ts[2].V
        return true
    })

var _ = Parser.Add().
    Match("02/01/06").
    Handle(func (d *Date, ts []*Token) bool {
        d.day = ts[0].V
        d.month = ts[2].V
        d.year = ts[4].V
        return true
    })

var _ = Parser.Add().
    Match("Jan").
    Handle(func (d *Date, ts []*Token) bool {
        d.month = ts[0].V
        return true
    })

var _ = Parser.Add().
    Match("-0700").
    Handle(func (d *Date, ts []*Token) bool {
        d.tzoffset = ts[0].V + ts[1].V
        return true
    })

var _ = Parser.Add().
    Match("-07:00").
    Handle(func (d *Date, ts []*Token) bool {
        d.tzoffset = ts[0].V + ts[1].V + ts[3].V
        return true
    })

var _ = Parser.Add().
    Match("2006").
    Handle(func (d *Date, ts []*Token) bool {
        if d.year != "" {
            return false
        }

        d.year = ts[0].V
        return true
    })

var _ = Parser.Add().
    Match("02").
    Handle(func (d *Date, ts []*Token) bool {
        if d.day != "" {
            return false
        }

        d.day = ts[0].V
        return true
    })

var _ = Parser.Add().
    Match("01").
    Handle(func (d *Date, ts []*Token) bool {
        if d.month != "" {
            return false
        }

        d.month = ts[0].V
        return true
    })

var _ = Parser.Add().
    Match("06").
    Handle(func (d *Date, ts []*Token) bool {
        if d.year != "" {
            return false
        }

        d.year = ts[0].V
        return true
    })

var _ = Parser.Add().
    Match("15").
    Handle(func (d *Date, ts []*Token) bool {
        if d.hour != "" {
            return false
        }

        d.hour = ts[0].V
        return true
    })








