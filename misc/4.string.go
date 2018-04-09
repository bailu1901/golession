package lesson

import s "strings"

import (
	"fmt"
)

var p = fmt.Println

func StringTest() {

	p("Contains:	", s.Contains("test", "es"))
	p("Count:		", s.Count("test", "t"))
	p("HasPrefix:	", s.HasPrefix("test", "te"))
	p("HasSuffix:	", s.HasSuffix("test", "st"))
	p("Join:		", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat:		", s.Repeat("ab", 5))
	p("Replace:		", s.Replace("test", "t", "1", -1))
	p("Replace:		", s.Replace("test", "t", "1", 0))
	p("Split:		", s.Split("1,2,3,4", ","))
	p("ToLower:		", s.ToLower("ABCD"))
	p("ToUpper:		", s.ToUpper("abcD"))
}
