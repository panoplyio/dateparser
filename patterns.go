package dateparser

import (
    "strings"
)

type HandleFn func(d *Date, ts []*Token) bool

type Pattern struct {
    children []*Pattern
    Matches []MatchFn
    HandleFn HandleFn
}

func NewPattern() *Pattern {
    return &Pattern{}
}

func (p *Pattern) Add() *Pattern {
    if p.children == nil {
        p.children = []*Pattern{}
    }

    ptn := &Pattern{}
    p.children = append(p.children, ptn)

    return ptn
}

func (p *Pattern) Parse(b []byte) *Date {
    tokens := (&Timelex{b, 0}).All()
    date := &Date{}

    idx := 0
    for idx < len(tokens) {
        n := p.parse(date, tokens[idx:])
        if n == 0 {
            n = 1 // unparsed. skip the first token and continue.
        }

        idx += n
    }

    return date
}

func (p *Pattern) parse(d *Date, ts []*Token) int {
    if len(ts) < len(p.Matches) {
        return 0
    }

    for i, match := range p.Matches {
        if !match(ts[i]) {
            return 0 // unmatched.
        }
    }

    if len(p.children) > 0 {
        for _, ptn := range p.children {
            if n := ptn.parse(d, ts); n > 0 {
                return n
            }
        }

        return 0
    }

    if p.HandleFn(d, ts) {
        return len(p.Matches)
    } else {
        return 0
    }
}

func (p *Pattern) Match(s string) *Pattern {
    allmatchers := []struct{prefix string; matcher *Matcher}{
        {"2006", YYYY},
        {"06", YY},
        {"01", Month},
        {"Jan", MonthName},
        {"02", DD},
        {"Mon", Weekday},
        {"MST", Timezone},
        {"0700", TimezoneOffset},
        {"07", HH12},
        {"15", HH24},
        {"03", HH12},
        {"04", MINS},
        {"05", SECS},
        {"00", SECS},
        {"hours", HoursName},
        {"mins", MinsName},
        {"secs", SecsName},
        {"pm", AmPm},
        {"/", DateSep},
        {"-", DateSep},
        {":", TimeSep},
        {"+-", Sign},
    }

    matchers := []MatchFn{}
    for len(s) > 0 {
        if s[0] == ' ' {
            s = s[1:] // skip spaces
            continue
        }

        found := false
        for _, m := range allmatchers {
            if strings.HasPrefix(s, m.prefix) {
                found = true
                matchers = append(matchers, m.matcher.Match)
                s = s[len(m.prefix):]
                break
            }
        }

        if !found {
            panic("Unrecognized format:" + s)
        }
    }

    p.Matches = matchers
    return p
}

func (p *Pattern) Handle(fn HandleFn) *Pattern {
    p.HandleFn = fn
    return p
}

