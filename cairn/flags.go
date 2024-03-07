package cairn

import "flag"

// Flags is a container for parsed command-line flags.
type Flags struct {
	Command string
}

// ParseFlags returns a parsed Flags from an argument slice.
func ParseFlags(ss []string) (*Flags, error) {
	f := flag.NewFlagSet("cairn", flag.ContinueOnError)
	c := f.String("c", "", "eval string")
	return &Flags{*c}, f.Parse(ss)
}
