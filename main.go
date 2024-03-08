package main

import (
	"fmt"
	"os"

	"github.com/wirehaiku/cairn/cairn"
)

func try(err error) {
	if err != nil {
		fmt.Printf("Error: %s.\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	c := cairn.NewCairn(os.Stdin, os.Stdout)
	f, err := cairn.ParseFlags(os.Args[1:])
	try(err)

	if f.Command != "" {
		try(c.EvaluateString(cairn.Library + f.Command))

	} else {
		c.WriteString("Cairn version 0.0.0 (2024-03-05).\n")
		try(c.EvaluateString(cairn.Library))

		for {
			c.WriteString(">>> ")
			s := c.ReadString('\n')

			if err := c.EvaluateString(s); err != nil {
				c.WriteString("Error: %s.\n\n", err.Error())

			} else if !c.Stack.Empty() {
				c.WriteString("[ %s ]\n", c.Stack.String())
			}
		}
	}
}
