package dateparser

type Timelex struct {
    buffer []byte
    idx int
}

func (t *Timelex) All() []*Token {
    tokens := []*Token{}
    for token := t.Next(); !token.IsEOF() ; token = t.Next() {
        tokens = append(tokens, token)
    }
    return tokens
}

func (t *Timelex) Next() *Token {
    token := []byte{}
    state := ""

    if t.idx >= len(t.buffer) {
        return &Token{"", "EOF"}
    }

L:
    for t.idx < len(t.buffer) {
        b := t.buffer[t.idx]

        switch state {
        case "":
            // uninitialized state; determine if we're reading a word or a 
            // number
            token = append(token, b)
            if IsWord(b) {
                state = "a"
            } else if IsDigit(b) {
                state = "0"
            } else if b == ' ' {
                token = []byte{} // reset the token
            } else {
                state = "?"
            }
        case "?":
            break L
        case "a":
            if !IsWord(b) {
                break L
            }

            token = append(token, b)
        case "0":
            if !IsDigit(b) {
                break L
            }

            token = append(token, b)
        }

        t.idx += 1
    }

    return &Token{string(token), state}
}


func IsWord(b byte) bool {
    return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

func IsDigit(b byte) bool {
    return b >= '0' && b <= '9'
}

type Token struct {
    V string
    T string
}

func (t *Token) IsEOF() bool {
    return t.T == "EOF"
}

func (t *Token) IsNumber() bool {
    return t.T == "0"
}

func (t *Token) IsLen(vs ...int) bool {
    if t.IsEOF() {
        return false
    }

    l := len(t.V)
    for _, v := range vs {
        if l == v {
            return true
        }
    }
    return false
}