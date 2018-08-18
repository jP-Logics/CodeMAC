package metrics

import (
	"CodeMAC/configuration"
	"CodeMAC/errors"
	"CodeMAC/helper"
	"strings"
)

//Counters map contains counters of all given configurations
var Counters map[string]int

// init is to initialize required data. This method works automatically
func init() {
	Counters = make(map[string]int, 0)
	Counters["lines"] = 0
}

// Measure access the given string and parses it and maintains the counters
func Measure(c configuration.Configuration, str string) {
	if c == nil {
		panic(errors.PanicNoConfiguration)
	}
	lines := strings.Split(str, "\n")
	Counters["lines"] = len(lines)
	strs := strings.FieldsFunc(str, helper.Split)
	for _, v := range strs {
		val := c.Get(strings.TrimSpace(v)).Value("type")
		hasToCount := c.Get(strings.TrimSpace(v)).Value("counter")
		if val != nil && val == "keyword" {
			if hasToCount != nil && hasToCount.(bool) {
				count, ok := Counters[v]
				if ok {
					count++
					Counters[v] = count
				} else {
					Counters[v] = 1
				}
			}
		}
	}
}

// if Function then caluclate: Function size/Number of lines of code
//
