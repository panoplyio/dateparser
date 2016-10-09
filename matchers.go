package dateparser

type MatchFn func(*Token) bool

// Utilities

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

