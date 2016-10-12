package dateparser

type MatchFn func(*Token) bool

type Matcher struct {
    Fmt string
    Match MatchFn
}

func Match(vs []string) *Matcher {
    // create a map of the values for fast lookups
    m := map[string]bool{}
    for _, v := range vs {
        m[v] = true
    }

    return &Matcher {
        Match: func(t *Token) bool {
            return m[t.V]
        },
    }
}

