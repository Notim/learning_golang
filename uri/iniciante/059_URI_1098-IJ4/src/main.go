package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0.0; i <= 2.0 ; i = i + 0.2 {
		for j := 1.0; j != 4; j++ {
			ifi := strings.Split(fmt.Sprintf("%.1f",i), ".")
			ise := strings.Split(fmt.Sprintf("%.1f",j + i), ".")
			if ifi[1] != "0" && ise[1] != "0"{
				fmt.Printf("I=%.1f J=%.1f\n", i, j + i)
			} else {
				fmt.Printf("I=%s J=%s\n", ifi[0],ise[0])
			}

		}
	}
}
