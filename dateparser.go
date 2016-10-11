package dateparser

var Parser = &Pattern{}
func Parse(b []byte) *Date {
    return Parser.Parse(b)
}

var _ = Parser.Add().
    MatchFmt("Mon").
    Handle(func (d *Date, ts []*Token) int {
        d.weekday = ts[0].V
        return 1
    })

var _ = Parser.Add().
    MatchFmt("MST").
    Handle(func (d *Date, ts []*Token) int {
        d.tz = ts[0].V
        return 1
    })

var _ = Parser.Add().
    MatchFmt("15 hours").
    Handle(func (d *Date, ts []*Token) int {
        d.hour = ts[0].V
        return 2
    })

var _ = Parser.Add().
    MatchFmt("03 pm").
    Handle(func (d *Date, ts []*Token) int {
        d.hour = ts[0].V
        return 2
    })

var _ = Parser.Add().
    MatchFmt("04 mins").
    Handle(func (d *Date, ts []*Token) int {
        d.minute = ts[0].V
        return 2
    })

var _ = Parser.Add().
    MatchFmt("05 secs").
    Handle(func (d *Date, ts []*Token) int {
        d.second = ts[0].V
        return 2
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
    MatchFmt("15:04:05").
    Handle(func (d *Date, ts []*Token) int {
        d.hour = ts[0].V
        d.minute = ts[2].V
        d.second = ts[4].V
        return 5
    })

var _ = Parser.Add().
    MatchFmt("15:04").
    Handle(func (d *Date, ts []*Token) int {
        if d.hour == "" {
            d.hour = ts[0].V
            d.minute = ts[2].V
        } else {
            d.minute = ts[0].V
            d.second = ts[2].V
        }

        return 3
    })

var _ = Parser.Add().
    MatchFmt("15 : 04 pm").
    Handle(func (d *Date, ts []*Token) int {
        if d.hour == "" {
            d.hour = ts[0].V
            d.minute = ts[2].V
        } else {
            d.minute = ts[0].V
            d.second = ts[2].V
        }

        return 4
    })

var _ = Parser.Add().
    MatchFmt("2006 - 01 - 02").
    Handle(func (d *Date, ts []*Token) int {
        d.year = ts[0].V
        d.month = ts[2].V
        d.day = ts[4].V
        return 5
    })

var _ = Parser.Add().
    MatchFmt("2006 01 02").
    Handle(func (d *Date, ts []*Token) int {
        d.year = ts[0].V
        d.month = ts[1].V
        d.day = ts[2].V
        return 3
    })

var _ = Parser.Add().
    MatchFmt("02 - 01 - 2006").
    Handle(func (d *Date, ts []*Token) int {
        d.day = ts[0].V
        d.month = ts[2].V
        d.year = ts[4].V
        return 5
    })

var _ = Parser.Add().
    MatchFmt("02 01 2006").
    Handle(func (d *Date, ts []*Token) int {
        d.day = ts[0].V
        d.month = ts[1].V
        d.year = ts[2].V
        return 3
    })

var _ = Parser.Add().
    MatchFmt("01 - 02 - 2006").
    Handle(func (d *Date, ts []*Token) int {
        d.month = ts[0].V
        d.day = ts[2].V
        d.year = ts[4].V
        return 5
    })

var _ = Parser.Add().
    MatchFmt("01 02 2006").
    Handle(func (d *Date, ts []*Token) int {
        d.month = ts[0].V
        d.day = ts[1].V
        d.year = ts[2].V
        return 3
    })

var _ = Parser.Add().
    MatchFmt("02 - 01 - 06").
    Match(DD, DateSep, Month, DateSep, YY).
    Handle(func (d *Date, ts []*Token) int {
        d.day = ts[0].V
        d.month = ts[2].V
        d.year = ts[4].V
        return 5
    })

var _ = Parser.Add().
    MatchFmt("Jan").
    Handle(func (d *Date, ts []*Token) int {
        d.month = ts[0].V
        return 1
    })

var _ = Parser.Add().
    MatchFmt("+- 0700").
    Handle(func (d *Date, ts []*Token) int {
        d.tzoffset = ts[0].V + ts[1].V
        return 2
    })

var _ = Parser.Add().
    // MatchFmt("-0700")
    MatchFmt("+- 07 : 00").
    Handle(func (d *Date, ts []*Token) int {
        d.tzoffset = ts[0].V + ts[1].V + ts[3].V
        return 4
    })

var _ = Parser.Add().
    // MatchFmt("20060102150405")
    Match(YYYYMMDDHHMMSS).
    Handle(func (d *Date, ts []*Token) int {
        if d.year != "" || d.month != "" || d.day != "" {
            return 0
        }

        d.year = ts[0].V[:4]
        d.month = ts[0].V[4:6]
        d.day = ts[0].V[6:8]
        d.hour = ts[0].V[8:10]
        d.minute = ts[0].V[10:12]
        d.second = ts[0].V[12:14]
        return 1
    })

var _ = Parser.Add().
    Match(YYYYMMDDHHMM).
    Handle(func (d *Date, ts []*Token) int {
        if d.year != "" || d.month != "" || d.day != "" {
            return 0
        }

        d.year = ts[0].V[:4]
        d.month = ts[0].V[4:6]
        d.day = ts[0].V[6:8]
        d.hour = ts[0].V[8:10]
        d.minute = ts[0].V[10:12]
        return 1
    })

var _ = Parser.Add().
    Match(YYYYMMDD).
    Handle(func (d *Date, ts []*Token) int {
        if d.year != "" || d.month != "" || d.day != "" {
            return 0
        }

        d.year = ts[0].V[:4]
        d.month = ts[0].V[4:6]
        d.day = ts[0].V[6:8]
        return 1
    })

var _ = Parser.Add().
    MatchFmt("2006").
    Handle(func (d *Date, ts []*Token) int {
        if d.year != "" {
            return 0
        }

        d.year = ts[0].V
        return 1
    })

var _ = Parser.Add().
    Match(HHMMSS).
    Handle(func (d *Date, ts []*Token) int {
        if d.hour != "" || d.minute != "" || d.second != "" {
            return 0
        }

        d.hour = ts[0].V[:2]
        d.minute = ts[0].V[2:4]
        d.second = ts[0].V[4:6]
        return 1
    })

var _ = Parser.Add().
    Match(HHMM).
    Handle(func (d *Date, ts []*Token) int {
        if d.hour != "" || d.minute != "" {
            return 0
        }

        d.hour = ts[0].V[:2]
        d.minute = ts[0].V[2:4]
        return 1
    })

var _ = Parser.Add().
    MatchFmt("02").
    Handle(func (d *Date, ts []*Token) int {
        if d.day != "" {
            return 0
        }

        d.day = ts[0].V
        return 1
    })

var _ = Parser.Add().
    MatchFmt("01").
    Handle(func (d *Date, ts []*Token) int {
        if d.month != "" {
            return 0
        }

        d.month = ts[0].V
        return 1
    })

var _ = Parser.Add().
    MatchFmt("06").
    Handle(func (d *Date, ts []*Token) int {
        if d.year != "" {
            return 0
        }

        d.year = ts[0].V
        return 1
    })

var _ = Parser.Add().
    MatchFmt("15").
    Handle(func (d *Date, ts []*Token) int {
        if d.hour != "" {
            return 0
        }

        d.hour = ts[0].V
        return 1
    })


// matchers
var HH12 = Match([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11",
    "12"})

var HH24 = MatchOr(HH12, Match([]string{"13", "14", "15", "16", "17", "18", 
    "19", "20", "21", "22", "23"}))

var HoursName = Match([]string{"h", "hour", "hours"})
var AmPm = Match([]string{"am", "pm"})

var MinsName = Match([]string{"m", "min", "mins", "minute", "minutes"})
var SecsName = Match([]string{"s", "sec", "secs", "second", "seconds"})

var Sep = Match([]string{" "})
var DateSep = Match([]string{"-", "/", "."})
var TimeSep = Match([]string{":"})
var Sign = Match([]string{"-", "+"})

var MM = Match([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11",
    "12"})

var MonthName = Match([]string{"January", "February", "March", "April", "May",
    "June", "July", "August", "September", "October", "Novemeber", "December",
    "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", 
    "Nov", "Dec"})

var Month = MatchOr(MonthName, MM)

var DD = Match([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11",
    "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23",
    "24", "25", "26", "27", "28", "29", "30", "31"})

var MINSEC = Match([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11",
    "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23",
    "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35",
    "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47",
    "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59"})

var Weekday = Match([]string{"Sunday", "Monday", "Tuesday", "Wednesday",
    "Thursday", "Friday", "Saturday", "Sun", "Mon", "Tue", "Tues", "Wed",
    "Thu", "Thur", "Thurs", "Fri", "Sat"})

var Whitespace = Match([]string{" ", "\n", "\t", "\r"})

var YYYY = MatchNumberLen(4)
var YYYYMMDD = MatchAnd(MatchSub(0, 4, YYYY),
    MatchSub(4, 6, MM),
    MatchSub(6, 8, DD))

var HHMM = MatchAnd(MatchSub(0, 2, HH24), 
    MatchSub(2, 4, MINSEC))

var HHMMSS = MatchAnd(MatchSub(0, 4, HHMM),
    MatchSub(4, 6, MINSEC))

var YYYYMMDDHHMMSS = MatchAnd(MatchSub(0, 8, YYYYMMDD),
    MatchSub(8, 14, HHMMSS))

var YYYYMMDDHHMM = MatchAnd(MatchSub(0, 8, YYYYMMDD),
    MatchSub(8, 12, HHMM))

// var HHMMSS = 
var YY = MatchNumberLen(2)
var TZOffset = MatchNumberLen(4)

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








