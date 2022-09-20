package manipulator

import (
	"fmt"
	s "strings"
)

func StringM(stg string) {
	var p = fmt.Println
	p("Contains:  ", s.Contains(stg, "es"))
	p("Count:     ", s.Count(stg, "t"))
	p("HasPrefix: ", s.HasPrefix(stg, "te"))
	p("HasSuffix: ", s.HasSuffix(stg, "st"))
	p("Index:     ", s.Index(stg, "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower(stg))
	p("ToUpper:   ", s.ToUpper(stg))
}
