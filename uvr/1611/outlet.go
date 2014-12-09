package uvr1611

import (
    "github.com/brutella/gouvr/uvr"
    "fmt"
)

type Outlet struct {
    Enabled bool
}

// OutletsFromValue returns a list of 13 outlets which
// represent the outlet A1 to A13
func OutletsFromValue(value uvr.Value) []Outlet {
    return []Outlet{
        Outlet{value.Low & 0x01 == 0x01},
        Outlet{value.Low & 0x02 == 0x02},
        Outlet{value.Low & 0x04 == 0x04},
        Outlet{value.Low & 0x08 == 0x08},
        Outlet{value.Low & 0x10 == 0x10},
        Outlet{value.Low & 0x20 == 0x20},
        Outlet{value.Low & 0x40 == 0x40},
        Outlet{value.Low & 0x80 == 0x80},
        Outlet{value.High & 0x01 == 0x01},
        Outlet{value.High & 0x02 == 0x02},
        Outlet{value.High & 0x04 == 0x04},
        Outlet{value.High & 0x08 == 0x08},
        Outlet{value.High & 0x10 == 0x10},
    }
}

func OutletsStringFromValue(value uvr.Value) string {
    var str string
    for i, v := range OutletsFromValue(value) {
        str += fmt.Sprintf("A%d %t\n", i+1, v.Enabled)
    }
    
    return str
}