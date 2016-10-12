package dateparser

import (
    "strings"
)

type HandleFn func(d *Date, ts []*Token) bool

type Pattern struct {
    children []*Pattern
    Matchers []MatchFn
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
    if len(ts) < len(p.Matchers) {
        return 0
    }

    for i, match := range p.Matchers {
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
        return len(p.Matchers)
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
        {"15", HH24},
        {"03", HH12},
        {"07", HH12}, // used for timezone offsets (07:00)
        {"04", MINS},
        {"05", SECS},
        {"00", SECS}, // used for timezone offsets (07:00)
        {"pm", AmPm},
        {"-", Sign},

        // uncaptured.
        {"hours", HoursName},
        {"mins", MinsName},
        {"secs", SecsName},
        {"/", DateSep},
        {":", TimeSep},
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

    p.Matchers = matchers
    return p
}

func (p *Pattern) Handle(fn HandleFn) *Pattern {
    p.HandleFn = fn
    return p
}



// --- MATCHERS


// 59-mins or 59-seconds
var MINS = Match([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11",
    "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23",
    "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35",
    "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47",
    "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59"})

var SECS = MINS

// 12-hours
var HH12 = Match([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11",
    "12"})

// 24-hours
var HH24 = Match([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11",
    "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"})

// 31-days
var DD = Match([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11",
    "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23",
    "24", "25", "26", "27", "28", "29", "30", "31"})

// weekday names
var Weekday = Match([]string{"Sunday", "Monday", "Tuesday", "Wednesday",
    "Thursday", "Friday", "Saturday", "Sun", "Mon", "Tue", "Tues", "Wed",
    "Thu", "Thur", "Thurs", "Fri", "Sat"})

// 12-months
var MM = Match([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11",
    "12"})

// month names
var MonthName = Match([]string{"January", "February", "March", "April", "May",
    "June", "July", "August", "September", "October", "Novemeber", "December",
    "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", 
    "Nov", "Dec"})

// month: either name or MM
var Month = Match([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11",
    "12", "January", "February", "March", "April", "May",
    "June", "July", "August", "September", "October", "Novemeber", "December",
    "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", 
    "Nov", "Dec"})

// 4-digit year
var YYYY = &Matcher{
    Match: func(t *Token) bool {
        return t.IsNumber() && t.IsLen(4)
    },
}

// 2-digit year
var YY = &Matcher{
    Match: func(t *Token) bool {
        return t.IsNumber() && t.IsLen(2)
    },
}

// 4-digits timezone offset: HHMM
var TimezoneOffset = &Matcher{
    Match: func(t *Token) bool {
        return t.IsNumber() && t.IsLen(4) &&
            HH12.Match(&Token{t.V[:2], t.T}) &&
            MINS.Match(&Token{t.V[2:4], t.T})
    },
}

// named timezone
var Timezone = Match([]string{"ACDT", "ACST", "ACT", "ACWDT", "ACWST", "ADDT",
    "ADT", "AEDT", "AEST", "AFT", "AHDT", "AHST", "AKDT", "AKST", "AMST", "AMT",
    "ANT", "APT", "ARST", "ART", "AST", "AWDT", "AWST", "AWT", "AZOMT", "AZOST",
    "AZOT", "BDST", "BDT", "BEAT", "BEAUT", "BMT", "BNT", "BORT", "BORTST",
    "BOST", "BOT", "BRST", "BRT", "BST", "BTT", "BURT", "CANT", "CAPT", "CAST",
    "CAT", "CAWT", "CCT", "CDDT", "CDT", "CEMT", "CEST", "CET", "CGST", "CGT",
    "CHADT", "CHAST", "CHDT", "CHOST", "CHOT", "CHUT", "CKHST", "CKT", "CLST",
    "CLT", "CMT", "COST", "COT", "CPT", "CST", "CUT", "CVST", "CVT", "CWT",
    "CXT", "ChST", "DACT", "DMT", "EASST", "EAST", "EAT", "ECT", "EDDT", "EDT",
    "EEST", "EET", "EGST", "EGT", "EHDT", "EMT", "EPT", "EST", "EWT", "FFMT",
    "FJST", "FJT", "FKST", "FKT", "FMT", "FNST", "FNT", "GALT", "GAMT", "GBGT",
    "GFT", "GHST", "GILT", "GMT", "GST", "GYT", "HDT", "HKST", "HKT", "HMT",
    "HOVST", "HOVT", "HST", "ICT", "IDDT", "IDT", "IHST", "IMT", "IOT", "IRDT",
    "IRST", "ISST", "IST", "JAVT", "JCST", "JDT", "JMT", "JST", "JWST", "KART",
    "KDT", "KMT", "KOST", "KST", "KWAT", "LHDT", "LHST", "LINT", "LKT", "LMT",
    "LRT", "LST", "MADMT", "MADST", "MADT", "MALST", "MALT", "MART", "MDDT",
    "MDST", "MDT", "MHT", "MIST", "MMT", "MOST", "MOT", "MPT", "MSD", "MSK",
    "MST", "MUST", "MUT", "MVT", "MWT", "MYT", "NCST", "NCT", "NDDT", "NDT",
    "NEGT", "NEST", "NET", "NFST", "NFT", "NMT", "NPT", "NRT", "NST", "NUT",
    "NWT", "NZDT", "NZMT", "NZST", "PDDT", "PDT", "PEST", "PET", "PGT", "PHOT",
    "PHST", "PHT", "PKST", "PKT", "PLMT", "PMDT", "PMMT", "PMST", "PMT", "PNT",
    "PONT", "PPMT", "PPT", "PST", "PWT", "PYST", "PYT", "QMT", "RET", "RMT",
    "SAST", "SBT", "SCT", "SDMT", "SDT", "SET", "SGT", "SJMT", "SMT", "SRT",
    "SST", "SWAT", "TAHT", "TBMT", "TKT", "TLT", "TMT", "TOST", "TOT", "TVT",
    "ULAST", "ULAT", "UYHST", "UYST", "UYT", "VET", "VUST", "VUT", "WAKT",
    "WARST", "WART", "WAST", "WAT", "WEMT", "WEST", "WET", "WFT", "WGST", "WGT",
    "WIB", "WIT", "WITA", "WMT", "WSDT", "WSST", "XJT", "YDDT", "YDT", "YPT",
    "YST", "YWT"})

// formatting
var HoursName = Match([]string{"h", "hour", "hours"})
var MinsName = Match([]string{"m", "min", "mins", "minute", "minutes"})
var SecsName = Match([]string{"s", "sec", "secs", "second", "seconds"})
var DateSep = Match([]string{"-", "/", "."})
var TimeSep = Match([]string{":"})

var AmPm = Match([]string{"am", "pm"})
var Sign = Match([]string{"-", "+"})

