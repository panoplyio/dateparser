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

// 10:24:05 am
// var _ = Parser.Add().
    // MatchFmt("(HH24) : (MINS) : (SECS) ampm").
    // Match(HH24, TimeSep, MINSEC, TimeSep, MINSEC, AmPm).
    // Handle(func (d *Date, ts []*Token) int {
    //     d.hour = ts[0].V
    //     d.minute = ts[2].V
    //     d.second = ts[4].V
    //     return 6
    // })

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
    Match("2006-01-02").
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
    Match("02-01-2006").
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
    Match("01-02-2006").
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
    Match("02-01-06").
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
    Match("+-0700").
    Handle(func (d *Date, ts []*Token) bool {
        d.tzoffset = ts[0].V + ts[1].V
        return true
    })

var _ = Parser.Add().
    Match("+-07:00").
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


// --- matchers

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








