package main

import (
	"strconv"
	"strings"
)

func genPool() []string {
	var collector []string
	c := "116.{GROUP1}.{GROUP2}.0/24"
	for i := 0; i < 255; i++ {
		cc := strings.Replace(c, "{GROUP1}", strconv.Itoa(i), -1)
		for j := 0; j < 255; j++ {
			collector = append(collector, strings.Replace(cc, "{GROUP2}", strconv.Itoa(j), -1))
		}
	}

	return collector
}
