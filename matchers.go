package dateparser

import (
    "strings"
)

type MatchFn func(*Token) bool

func MatchFmt(s string) []MatchFn {
    comps := strings.Split(s, " ")
    matchers := make([]MatchFn, len(comps))

    for i, comp := range comps {
        switch comp {
        case "2006":
            matchers[i] = YYYY
        case "06":
            matchers[i] = YY
        case "01":
            matchers[i] = Month
        case "Jan":
            matchers[i] = MonthName
        case "02":
            matchers[i] = DD
        case "Mon":
            matchers[i] = Weekday
        case "MST":
            matchers[i] = Timezone
        case "0700":
            matchers[i] = TZOffset
        case "07":
            matchers[i] = HH12
        case "15":
            matchers[i] = HH24
        case "03":
            matchers[i] = HH12
        case "04":
            matchers[i] = MINSEC
        case "05":
            matchers[i] = MINSEC
        case "00":
            matchers[i] = MINSEC
        case "hours":
            matchers[i] = HoursName
        case "mins":
            matchers[i] = MinsName
        case "secs":
            matchers[i] = SecsName
        case "pm":
            matchers[i] = AmPm
        case "/":
            matchers[i] = DateSep
        case "-":
            matchers[i] = DateSep
        case ":":
            matchers[i] = TimeSep
        case "+-":
            matchers[i] = Sign
        default:
            panic("Unknown format: " + comp)
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

