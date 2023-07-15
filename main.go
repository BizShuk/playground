package main

import (
	"fmt"
	"os"

	"sample.shuk/tools"
)

func main() {
	fmt.Println(tools.COUNTER2_A)
	fmt.Println(tools.COUNTER2_B)
	fmt.Println(tools.COUNTER2_C)
	fmt.Println(tools.COUNTER2_D)
	x, _ := os.Stat(".")
	fmt.Println(x.IsDir())

	fmt.Println(tools.SampleIntList())
}

/*
//  Wrap a struct method
func main() {
	v1 := &app{30}
	f1 := wrapper(v1.Run)

	fmt.Println(f1(30))
}

func wrapper(v1 func(int) int) func(int) int {
	return v1
}

type app struct {
	p1 int
}

func (a *app) Run(v1 int) int {
	return a.p1 + v1
}
*/
/*
	6. phone validator XXX-XXX-XXXX, if invalid, ERROR out log
	8. improve import process with goroutine/channel


	cmd:
	4. create cmd for import data
	5. export json file to csv file
	7. add cli output with -column flag





	[x]2. read csv file ()
	[x]3. writing json object into json file
*/
