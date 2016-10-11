package dateparser

import (
    "strings"
)

type MatchFn func(*Token) bool

func MatchFmt(s string) []MatchFn {
    allmatchers := []struct{prefix string; matcher MatchFn}{
        {"2006", YYYY},
        {"06", YY},
        {"01", Month},
        {"Jan", MonthName},
        {"02", DD},
        {"Mon", Weekday},
        {"MST", Timezone},
        {"0700", TZOffset},
        {"07", HH12},
        {"15", HH24},
        {"03", HH12},
        {"04", MINSEC},
        {"05", MINSEC},
        {"00", MINSEC},
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
                matchers = append(matchers, m.matcher)
                s = s[len(m.prefix):]
                break
            }
        }

        if !found {
            panic("Unrecognized format:" + s)
        }
    }

    return matchers
}

func Match(vs []string) MatchFn {
    // create a map of the values for fast lookups
    m := map[string]bool{}
    for _, v := range vs {
        m[v] = true
    }

    return func(t *Token) bool {
        return m[t.V]
    }
}

func MatchOr(matchers ...MatchFn) MatchFn {
    return func(t *Token) bool {
        for _, matcher := range matchers {
            if matcher(t) {
                return true
            }
        }
        return false
    }
}

func MatchAnd(matchers ...MatchFn) MatchFn {
    return func(t *Token) bool {
        for _, matcher := range matchers {
            if !matcher(t) {
                return false
            }
        }
        return true
    }
}

func MatchNumberLen(n int) MatchFn {
    return func(t *Token) bool {
        return t.IsNumber() && t.IsLen(n)
    }
}

func MatchSub(start, end int, matcher MatchFn) MatchFn {
    return func(t *Token) bool {
        if len(t.V) < end {
            return false
        }

        return matcher(&Token{t.V[start:end], t.T})
    }
}

