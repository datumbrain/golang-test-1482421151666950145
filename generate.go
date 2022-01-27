package main

import (
	regen "github.com/zach-klippenstein/goregen"
	"math/rand"
	"time"
)

// generate function, takes boolean flag and generates random correct strings if the parameter is true and random
// incorrect strings if the flag is false.
// difficulty: medium
// estimated time: 1 hour and 30 minutes
// actual time: 1 hour and 25 minutes
func generate(flag bool) string {
	if flag {
		// this pattern helps to generate the random string of defined format e.g "3-ab-48-caba-56-haha"
		pattern := "^((^([0-9]{1,4})-[a-z]{1,4})(((-)([0-9]{1,4})-([a-z]{1,4})){1,5}))$"

		// Using a random seed (e.g. time-based) so we can generate random strings
		x := rand.NewSource(time.Now().UnixNano())
		y := rand.New(x)
		generator, _ := regen.NewGenerator(pattern, &regen.GeneratorArgs{ // this function will generate the random string of the given pattern (regular expression)
			RngSource: rand.NewSource(int64(y.Intn(10))),
		})

		str := generator.Generate()
		return str
	} else {
		// this is a wrong pattern to generate the random wrong string
		pattern := "^((^([0-9])-[0-9]{1,4})((([0-9]{1,4})([a-z]{1,4})){1,5}))$"

		// Using a random seed (e.g. time-based) so we can generate random strings
		x := rand.NewSource(time.Now().UnixNano())
		y := rand.New(x)
		generator, _ := regen.NewGenerator(pattern, &regen.GeneratorArgs{ // this function will generate the random string of the given pattern (regular expression)
			RngSource: rand.NewSource(int64(y.Intn(10))),
		})

		str := generator.Generate()
		return str
	}
}
