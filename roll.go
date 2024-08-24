package main

import (
	"fmt"
	// "log"
	"math/rand/v2"
	// "os"
	"strconv"
	"strings"
)

var die_options = map[int]bool{
	2: true,
	4: true,
	6: true,
	8: true,
	10: true,
	20: true,
}

func make_rolls(rolls []string) {
	if len(rolls) > 0 {
		i := 1
		for _, roll := range rolls {
			roll_args := strings.Split(roll, ",")
			die_arg, err1 := strconv.Atoi(roll_args[0])
			if err1 != nil { err(fmt.Sprintf("Unable to parse arg %d in '%v'", die_arg, roll))}
			num_rolls, err2 := strconv.Atoi(roll_args[1])
			if err2 != nil { err(fmt.Sprintf("Unable to parse arg %d in '%v'", num_rolls, roll))}

			if !die_options[die_arg] {
				err(fmt.Sprintf("Invalid die '%d'", die_arg));
			}

			for j := 0; j < num_rolls; j++ {
				fmt.Printf("%d: D%d rolled %d\n", i, die_arg, rand.IntN(die_arg) + 1)
				i++
			}
		}
	} else {
		err("Missing args")
	}
}
