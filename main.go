package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

type Status struct {
	Tracing []int
	Values  []int
}

var version = "MISSING build version [git hash]"

var input = make([]int, 0)
var circle bool
var length int
var queue = make([]*Status, 0)

func puzzle() {

	// stat.Values = input
	queue = append(queue, &Status{Values: input})
	for {
		// fmt.Println(queue)
		stats := pop(&queue)
		// fmt.Println(stats)
		for i := 0; i < length; i++ {
			// tracing
			values := make([]int, length)
			copy(values, stats.Values)
			newStat := &Status{
				Tracing: stats.Tracing,
				Values:  values,
			}
			newStat.Next(i)
			queue = append(queue, newStat)
		}
	}

}

func pop(v *[]*Status) *Status {
	stat := (*v)[0]
	// fmt.Println(stat)
	(*v) = (*v)[1:]
	return stat
}

func (stat *Status) Next(next int) {
	stat.Tracing = append(stat.Tracing, next)
	if next == 0 {
		if circle {
			stat.Values[length-1] += 1
		}
		stat.Values[0] += 1
		stat.Values[1] += 1
	} else if next == (length - 1) {
		if circle {
			stat.Values[0] += 1
		}
		stat.Values[length-1] += 1
		stat.Values[length-2] += 1
	} else {
		stat.Values[next-1] += 1
		stat.Values[next] += 1
		stat.Values[next+1] += 1
	}
	for index := range stat.Values {
		if stat.Values[index] > 3 {
			stat.Values[index] -= 3
		}
	}
	// judge
	tag := true
	for _, value := range stat.Values {
		if value != 3 {
			tag = false
			break
		}
	}
	if tag {
		for index := range stat.Tracing {
			stat.Tracing[index] += 1
		}
		fmt.Println(stat.Tracing)
		os.Exit(0)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "genshin-daoqi-puzzle"
	app.Usage = "genshin daoqi puzzle"
	app.Version = fmt.Sprintf("Git:[%s] (%s)", strings.ToUpper(version), runtime.Version())
	app.Flags = []cli.Flag{
		&cli.IntSliceFlag{
			Name:     "input",
			Usage:    "input values, it should be like {1,2,3}",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "circle",
			Usage: "circle or not",
			Value: false,
		},
	}
	app.Action = func(c *cli.Context) error {
		input = c.IntSlice("input")
		length = len(input)
		circle = c.Bool("circle")
		puzzle()
		return nil
	}
	app.Run(os.Args)
}
