package dateparser

import (
    // "fmt"
)

type HandleFn func(d *Date, ts []*Token) int

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

    p.Handle(func(d *Date, ts []*Token) int {
        for _, ptn := range p.children {
            if n := ptn.parse(d, ts); n > 0 {
                return n
            }
        }

        return 0
    })

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
    if p.Matches != nil {
        if len(ts) < len(p.Matches) {
            return 0
        }

        for i, match := range p.Matches {
            if !match(ts[i]) {
                return 0 // unmatched.
            }
        }
    }

    return p.HandleFn(d, ts)
}

func (p *Pattern) Match(matches ...MatchFn) *Pattern {
    p.Matches = matches
    return p
}

func (p *Pattern) MatchFmt(s string) *Pattern {
    return p.Match(MatchFmt(s)...)
}

func (p *Pattern) Handle(fn HandleFn) *Pattern {
    p.HandleFn = fn
    return p
}

