package main

import (
	"fmt"
	"os"

	"github.com/wirehaiku/cairn/cairn"
)

func die(s string, vs ...any) {
	s = fmt.Sprintf(s, vs...)
	fmt.Printf("Error: %s.\n", s)
	os.Exit(1)
}

func try(err error) {
	if err != nil {
		die(err.Error())
	}
}

func main() {
	c := cairn.NewCairn(os.Stdin, os.Stdout)
	f, err := cairn.ParseFlags(os.Args[1:])
	try(err)
	try(c.Execute(cairn.Library))

	if f.Command != "" {
		try(c.Execute(f.Command))

	} else if len(f.Files) != 0 {

		for _, p := range f.Files {
			bs, err := os.ReadFile(p)
			if err != nil {
				die("cannot execute file %q", p)
			}

			try(c.Execute(string(bs)))
		}

	} else {
		c.WriteString("Cairn version 0.0.0 (2024-03-05).\n")

		for {
			c.WriteString(">>> ")
			s := c.ReadString('\n')

			if err := c.Execute(s); err != nil {
				c.WriteString("Error: %s.\n\n", err.Error())

			} else if !c.Stack.Empty() {
				c.WriteString("[ %s ]\n", c.Stack.String())
			}
		}
	}
}
