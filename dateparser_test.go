package dateparser

import (
    "fmt"
    "testing"
    assert "github.com/stretchr/testify/require"
)

var tests = []struct{s string; e string}{
    {   // date command reversed
        "Thu Sep 25 10:36:28 BRST 2003",
        "&{2003 Sep 25 10 36 28 Thu BRST }",
    },
    {   // date command reversed
        "Thu Sep 25 10:36:28 BRST 2003",
        "&{2003 Sep 25 10 36 28 Thu BRST }",
    },
    {   // date command stripped timezone
        "Thu Sep 25 10:36:28 2003",
        "&{2003 Sep 25 10 36 28 Thu  }",
    },
    {   // date command stripped year
        "Thu Sep 25 10:36:28",
        "&{ Sep 25 10 36 28 Thu  }",
    },
    {   // date command stripped day
        "Thu Sep 10:36:28",
        "&{ Sep  10 36 28 Thu  }",
    },
    {   // date command stripped month
        "Thu 10:36:28",
        "&{   10 36 28 Thu  }",
    },
    {   // date command stripped weekday
        "Sep 10:36:28",
        "&{ Sep  10 36 28   }",
    },
    {   // date command only time
        "10:36:28",
        "&{   10 36 28   }",
    },
    {   // date command only time stripped seconds
        "10:36",
        "&{   10 36    }",
    },
    {   // date command only date
        "Thu Sep 25 2003",
        "&{2003 Sep 25    Thu  }",
    },
    {   // date command only date stripped weekday
        "Sep 25 2003",
        "&{2003 Sep 25      }",
    },
    {   // date command only date stripped day
        "Sep 2003",
        "&{2003 Sep       }",
    },
    {   // month only
        "Sep",
        "&{ Sep       }",
    },
    {   // year only
        "2003",
        "&{2003        }",
    },
    {   // date command R format
        "Thu, 25 Sep 2003 10:49:41 -0300",
        "&{2003 Sep 25 10 49 41 Thu  -0300}",
    },
    {   // ISO-8601
        "2003-09-25T10:49:41-03:00", // TODO: 41.5
        "&{2003 09 25 10 49 41   -0300}",
    },
    {   // ISO-8601 stripped timezone
        "2003-09-25T10:49:41",
        "&{2003 09 25 10 49 41   }",
    },
    {   // ISO-8601 stripped seconds
        "2003-09-25T10:49",
        "&{2003 09 25 10 49    }",
    },
    {   // ISO-8601 stripped minutes
        "2003-09-25T10",
        "&{2003 09 25 10     }",
    },
    {   // ISO-8601 stripped hours
        "2003-09-25",
        "&{2003 09 25      }",
    },
    {   // Python Logger
        "2003-09-25 10:49:41,502",
        "&{2003 09 25 10 49 41   }",
    },
    {   // dashed date
        "2003-09-25",
        "&{2003 09 25      }",
    },
    {   // dashed date with named month
        "2003-Sep-25",
        "&{2003 Sep 25      }",
    },
    {
        "Sep-25-2003",
        "&{2003 Sep 25      }",
    },
    {
        "09-25-2003",
        "&{2003 09 25      }",
    },
    {
        "25-09-2003",
        "&{2003 09 25      }",
    },
    {
        "10-09-2003",
        "&{2003 09 10      }",
    },
    {
        "10-09-03",
        "&{03 09 10      }",
    },
    {
        "2003.09.25",
        "&{2003 09 25      }",
    },
    {
        "2003.Sep.25",
        "&{2003 Sep 25      }",
    },
    {
        "25.Sep.2003",
        "&{2003 Sep 25      }",
    },
    {
        "Sep.25.2003",
        "&{2003 Sep 25      }",
    },
    {
        "09.25.2003",
        "&{2003 09 25      }",
    },
    {
        "25.09.2003",
        "&{2003 09 25      }",
    },
    {
        "10.09.2003",
        "&{2003 09 10      }",
    },
    {
        "10.09.03",
        "&{03 09 10      }",
    },
    {
        "2003/09/25",
        "&{2003 09 25      }",
    },
    {
        "2003/Sep/25",
        "&{2003 Sep 25      }",
    },
    {
        "25/Sep/2003",
        "&{2003 Sep 25      }",
    },
    {
        "Sep/25/2003",
        "&{2003 Sep 25      }",
    },
    {
        "09/25/2003",
        "&{2003 09 25      }",
    },
    {
        "25/09/2003",
        "&{2003 09 25      }",
    },
    {
        "10/09/2003",
        "&{2003 09 10      }",
    },
    {
        "10/09/03",
        "&{03 09 10      }",
    },
    {
        "2003 09 25",
        "&{2003 09 25      }",
    },
    {
        "2003 Sep 25",
        "&{2003 Sep 25      }",
    },
    {
        "25 Sep 2003",
        "&{2003 Sep 25      }",
    },
    {
        "Sep 25 2003",
        "&{2003 Sep 25      }",
    },
    {
        "09 25 2003",
        "&{2003 09 25      }",
    },
    {
        "25 09 2003",
        "&{2003 09 25      }",
    },
    {
        "10 09 2003",
        "&{2003 09 10      }",
    },
    {
        "10 09 03",
        "&{03 09 10      }",
    },
    {
        "25 09 03",
        "&{03 09 25      }",
    },
    // {
    //     "03 25 Sep",
    //     "&{03 09 25      }",
    // },
    {
        "10h36m28s",
        "&{   10 36 28   }",
    },
    {
        "10h36m",
        "&{   10 36    }",
    },
    {
        "10h",
        "&{   10     }",
    },
    {
        "10:00 am",
        "&{   10 00    }",
    },
}

func TestParse(t *testing.T) {
    for _, test := range tests {
        v := fmt.Sprint(Parse([]byte(test.s)))
        fmt.Println("Parse=", test.s, "Result=", v)
        assert.Equal(t, test.e, v)
    }
}

