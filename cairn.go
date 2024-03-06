///////////////////////////////////////////////////////////////////////////////////////
//                   Cairn: A personal programming language in Go.                   //
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //
//                            Version 0.0.0 (2024-03-05).                            //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////
//                              Part 1: Global Variables                             //
///////////////////////////////////////////////////////////////////////////////////////

// 1.1: Memory Variables
/////////////////////////

// Commands is a map of symbols to built-in command functions.
var Commands = make(map[string]func() error)

// Functions is a map of symbols to program-defined functions.
var Functions = make(map[string][]any)

// Queue is a first-in-first-out queue of parsed atoms.
var Queue = make([]any, 0, 63356)

// Registers is a fixed array of stored register values.
var Registers [8]uint8

// Stack is a last-in-first-out stack of stored memory values.
var Stack = make([]uint8, 0, 65536)

// 1.2: System Variables
/////////////////////////

// ExitFunc is the function used to exit the program.
var ExitFunc = os.Exit

// Library is the embedded standard library file.
//
//go:embed library.cairn
var Library string

// Stdin is the default input Reader.
var Stdin = bufio.NewReader(os.Stdin)

// Stdout is the default output Writer.
var Stdout = bufio.NewWriter(os.Stdout)

// VersionDate is the date of the current Cairn version.
var VersionDate = "2024-03-05"

// VersionNums is the SemVer number of the current Cairn version.
var VersionNums = "0.0.0"

///////////////////////////////////////////////////////////////////////////////////////
//                            Part 2: Collection Functions                           //
///////////////////////////////////////////////////////////////////////////////////////

// 2.1: Queue Functions
////////////////////////

// Dequeue removes and returns the first atom in the Queue.
func Dequeue() (any, error) {
	if len(Queue) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	a := Queue[0]
	Queue = Queue[1:]
	return a, nil
}

// DequeueTo removes and returns all atoms before an Atom in the Queue.
func DequeueTo(a any) ([]any, error) {
	var as []any
	for {
		a2, err := Dequeue()
		if err != nil {
			return nil, err
		}

		as = append(as, a2)
		if as[len(as)-1] == a {
			break
		}
	}

	return as[:len(as)-1], nil
}

// Enqueue appends an atom to the end of the Queue.
func Enqueue(a any) {
	Queue = append(Queue, a)
}

// EnqueueAll appends an atom slice to the end of the Queue.
func EnqueueAll(as []any) {
	Queue = append(Queue, as...)
}

// 2.2: Register Functions
///////////////////////////

// GetRegister returns the value of a given register.
func GetRegister(i uint8) (uint8, error) {
	if i > 7 {
		return 0, fmt.Errorf("register %d does not exist", i)
	}

	return Registers[int(i)], nil
}

// SetRegister sets the value of a given register.
func SetRegister(i, u uint8) error {
	if i > 7 {
		return fmt.Errorf("register %d does not exist", i)
	}

	Registers[int(i)] = u
	return nil
}

// 2.3: Stack Functions
////////////////////////

// Pop removes and returns the top item on the Stack.
func Pop() (uint8, error) {
	if len(Stack) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	u := Stack[len(Stack)-1]
	Stack = Stack[:len(Stack)-1]
	return u, nil
}

// PopN removes and returns the top N items on the Stack.
func PopN(i int) ([]uint8, error) {
	var us []uint8
	for len(us) < i {
		u, err := Pop()
		if err != nil {
			return nil, err
		}

		us = append(us, u)
	}

	return us, nil
}

// Push appends an integer to the top of the Stack.
func Push(u uint8) {
	Stack = append(Stack, u)
}

// PushAll appends an integer slice to the top of the Stack.
func PushAll(us []uint8) {
	Stack = append(Stack, us...)
}

///////////////////////////////////////////////////////////////////////////////////////
//                       Part 3: Parsing & Evaluation Functions                      //
///////////////////////////////////////////////////////////////////////////////////////

// 3.1: Parsing Functions
//////////////////////////

// Clean returns an uppercase program string without comments.
func Clean(s string) string {
	var ss []string
	for _, s := range strings.Split(s, "\n") {
		s = strings.SplitN(s, "//", 2)[0]
		s = strings.TrimSpace(s)
		ss = append(ss, strings.ToUpper(s))
	}

	return strings.Join(ss, "\n")
}

// Tokenise returns a token slice from a clean program string.
func Tokenise(s string) []string {
	return strings.Fields(s)
}

// 3.2: Evaluation Functions
/////////////////////////////

// Atomise returns an atom from a token string.
func Atomise(s string) any {
	if u, err := strconv.ParseUint(s, 10, 8); err == nil {
		return uint8(u)
	}

	return s
}

// AtomiseAll returns an atom slice from a token slice.
func AtomiseAll(ss []string) []any {
	var as []any
	for _, s := range ss {
		as = append(as, Atomise(s))
	}

	return as
}

// Evaluate evaluates an atom.
func Evaluate(a any) error {
	switch a := a.(type) {
	case uint8:
		Push(a)
		return nil

	case string:
		if c, ok := Commands[a]; ok {
			return c()
		}

		if as, ok := Functions[a]; ok {
			return EvaluateAll(as)
		}

		return fmt.Errorf("symbol %q is not defined", a)

	case func() error:
		return a()

	case []any:
		return EvaluateAll(a)

	default:
		return fmt.Errorf(`cannot evaluate atom type "%T"`, a)
	}
}

// EvaluateAll evaluates an atom slice.
func EvaluateAll(as []any) error {
	for _, a := range as {
		if err := Evaluate(a); err != nil {
			return err
		}
	}

	return nil
}

// 3.3: Standard Library Functions
///////////////////////////////////

// Import evaluates the contents of a file.
func Import(p string) error {
	bs, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("cannot read file %q", p)
	}

	ss := Tokenise(Clean(string(bs)))
	as := AtomiseAll(ss)
	return EvaluateAll(as)
}

///////////////////////////////////////////////////////////////////////////////////////
//                           Part 4: Input/Output Functions                          //
///////////////////////////////////////////////////////////////////////////////////////

// 4.1: Standard IO Functions
//////////////////////////////

// Input returns an ASCII character as an integer.
func Input() (uint8, error) {
	r, err := Stdin.ReadByte()
	if err != nil {
		return 0, fmt.Errorf("cannot read STDIN")
	}

	return uint8(r), nil
}

// Output writes an integer as an ASCII character to Stdout.
func Output(u uint8) error {
	if err := Stdout.WriteByte(u); err != nil {
		return fmt.Errorf("cannot write STDOUT")
	}

	if err := Stdout.Flush(); err != nil {
		return fmt.Errorf("cannot write STDOUT")
	}

	return nil
}

// 4.2: Command-Line Functions
///////////////////////////////

// Flags is a container for parsed command-line flags.
type Flags struct {
	Command string
}

// ParseFlags returns a parsed Flags from an argument slice.
func ParseFlags(ss []string) (*Flags, error) {
	fs := flag.NewFlagSet("cairn", flag.ContinueOnError)
	fc := fs.String("c", "", "eval single string")
	return &Flags{*fc}, fs.Parse(ss)
}

///////////////////////////////////////////////////////////////////////////////////////
//                             Part 5: Command Functions                             //
///////////////////////////////////////////////////////////////////////////////////////

// 5.1: Command Helper Functions
/////////////////////////////////

// Bool returns a boolean as an integer.
func Bool(b bool) uint8 {
	if b {
		return 1
	}

	return 0
}

// 5.2: Integer Commands
/////////////////////////

// ADD (a b → c) returns a + b.
func ADD() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(us[1] + us[0])
	return nil
}

// SUB (a b → c) returns a - b.
func SUB() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(us[1] - us[0])
	return nil
}

// MOD (a b → c) returns a % b.
func MOD() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(us[1] % us[0])
	return nil
}

// GTE (a b → c) returns a >= b.
func GTE() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(Bool(us[1] >= us[0]))
	return nil
}

// LTE (a b → c) returns a <= b.
func LTE() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(Bool(us[1] <= us[0]))
	return nil
}

// 5.3: Memory Commands
////////////////////////

// CLR (... → _) clears the stack.
func CLR() error {
	Stack = make([]uint8, 0, 65536)
	return nil
}

// DUP (a → a a) duplicates the top stack item.
func DUP() error {
	u, err := Pop()
	if err != nil {
		return err
	}

	PushAll([]uint8{u, u})
	return nil
}

// DRP (a b → a) deletes the top stack item.
func DRP() error {
	_, err := Pop()
	if err != nil {
		return err
	}

	return nil
}

// SWP (a b → b a) swaps the top two stack items.
func SWP() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	PushAll([]uint8{us[0], us[1]})
	return nil
}

// GET (a → b) returns the value of register a.
func GET() error {
	u, err := Pop()
	if err != nil {
		return err
	}

	u, err = GetRegister(u)
	if err != nil {
		return err
	}

	Push(u)
	return nil
}

// SET (a b → _) sets a to register b.
func SET() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	return SetRegister(us[0], us[1])
}

// 5.4: Logic Commands
///////////////////////

// EQU (a b → c) returns true if a equals b.
func EQU() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(Bool(us[0] == us[1]))
	return nil
}

// NEQ (a b → c) returns true if a does not equal b.
func NEQ() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(Bool(us[0] != us[1]))
	return nil
}

// AND (a b → c) returns true if both a and b are true.
func AND() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(Bool(us[0] != 0 && us[1] != 0))
	return nil
}

// ORR (a b → c) returns true if either a or b are true.
func ORR() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(Bool(us[0] != 0 || us[1] != 0))
	return nil
}

// XOR (a b → c) returns true if only a or only b is true.
func XOR() error {
	us, err := PopN(2)
	if err != nil {
		return err
	}

	Push(Bool(us[0] != 0 && us[1] == 0 || us[0] == 0 && us[1] != 0))
	return nil
}

// NOT (a → b) returns true if a is false.
func NOT() error {
	u, err := Pop()
	if err != nil {
		return err
	}

	Push(Bool(u == 0))
	return nil
}

// 5.5: System Commands
////////////////////////

// INN (_ → a) returns an input ASCII character as an integer.
func INN() error {
	u, err := Input()
	if err != nil {
		return err
	}

	Push(u)
	return nil
}

// OUT (a → _) writes a as an ASCII character to output.
func OUT() error {
	u, err := Pop()
	if err != nil {
		return err
	}

	return Output(u)
}

// BYE (_ → _) exits the program successfully.
func BYE() error {
	ExitFunc(0)
	return nil
}

// DIE (a → _) exits the program with error code a.
func DIE() error {
	u, err := Pop()
	if err != nil {
		return err
	}

	ExitFunc(int(u))
	return nil
}

// 5.6: Flow Control Commands
//////////////////////////////

// IFT (a → _) evaluates code if a is true.
func IFT() error {
	u, err := Pop()
	if err != nil {
		return err
	}

	as, err := DequeueTo("END")
	if err != nil {
		return err
	}

	if u != 0 {
		return EvaluateAll(as)
	}

	return nil
}

// IFF (a → _) evaluates code if a is false.
func IFF() error {
	u, err := Pop()
	if err != nil {
		return err
	}

	as, err := DequeueTo("END")
	if err != nil {
		return err
	}

	if u == 0 {
		return EvaluateAll(as)
	}

	return nil
}

// FOR (_ → _) evaluates code until register a is false.
func FOR() error {
	a, err := Dequeue()
	if err != nil {
		return err
	}

	as, err := DequeueTo("END")
	if err != nil {
		return err
	}

	switch a := a.(type) {
	case uint8:
		for {
			if err := EvaluateAll(as); err != nil {
				return err
			}

			if Registers[a] == uint8(0) {
				break
			}
		}

		return nil

	default:
		return fmt.Errorf(`"%v" is not an integer`, a)
	}
}

// DEF (_ → _) sets a symbol to a named function.
func DEF() error {
	a, err := Dequeue()
	if err != nil {
		return err
	}

	switch a := a.(type) {
	case string:
		as, err := DequeueTo("END")
		if err != nil {
			return err
		}

		Functions[a] = as
		return nil

	default:
		return fmt.Errorf(`"%v" is not a symbol`, a)
	}
}

// TST (a → _) prints a error message if a is false.
func TST() error {
	u, err := Pop()
	if err != nil {
		return err
	}

	as, err := DequeueTo("END")
	if err != nil {
		return err
	}

	if u == 0 {
		var ss []string
		for _, a := range as {
			ss = append(ss, fmt.Sprintf("%v", a))
		}

		s := strings.Join(ss, " ")
		fmt.Fprintf(Stdout, "ASSERT: %s.\n", s)
		Stdout.Flush()
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////////////
//                               Part 6: Main Functions                              //
///////////////////////////////////////////////////////////////////////////////////////

// 6.1: Main Helper Functions
//////////////////////////////

// dump returns the Stack's and Registers's contents as a string.
func dump() string {
	var ss []string

	for i, u := range Registers {
		if u != 0 {
			ss = append(ss, fmt.Sprintf("R%d=%d", i, u))
		}
	}

	if len(ss) > 0 {
		return fmt.Sprintf("%v %s", Stack, strings.Join(ss, " "))
	}

	return fmt.Sprintf("%v", Stack)
}

// print prints a formatted string to Stdout.
func print(s string, vs ...any) {
	fmt.Fprintf(Stdout, s, vs...)
	Stdout.Flush()
}

// prompt prints a prompt and returns an input string.
func prompt(s string) string {
	fmt.Fprint(Stdout, s)
	Stdout.Flush()
	s, _ = Stdin.ReadString('\n')
	return s
}

// once evaluates a program string once.
func once(s string) error {
	ss := Tokenise(Clean(s))
	as := AtomiseAll(ss)
	EnqueueAll(as)

	for len(Queue) > 0 {
		a, err := Dequeue()
		if err != nil {
			return err
		}

		if err := Evaluate(a); err != nil {
			return err
		}
	}

	return nil
}

// 6.2: Main Boot Functions
////////////////////////////

// init initialises the main Cairn program.
func init() {
	Commands = map[string]func() error{
		"ADD": ADD, "SUB": SUB, "MOD": MOD, "GTE": GTE, "LTE": LTE,
		"CLR": CLR, "DUP": DUP, "DRP": DRP, "SWP": SWP, "GET": GET, "SET": SET,
		"EQU": EQU, "NEQ": NEQ, "AND": AND, "ORR": ORR, "XOR": XOR, "NOT": NOT,
		"INN": INN, "OUT": OUT, "BYE": BYE, "DIE": DIE,
		"IFT": IFT, "IFF": IFF, "FOR": FOR, "DEF": DEF,
	}
}

// main executes the main Cairn program.
func main() {
	fs, err := ParseFlags(os.Args[1:])
	if err != nil {
		print(err.Error())
		ExitFunc(1)
	}

	switch {
	case fs.Command != "":
		if err := once(Library + fs.Command); err != nil {
			print("Error: %s.\n", err.Error())
		}

	default:
		print("Cairn version %s (%s).\n", VersionNums, VersionDate)
		if err := once(Library); err != nil {
			print("Error: %s.\n", err.Error())
			ExitFunc(1)
		}

		for {
			s := prompt(">>> ")
			if err := once(s); err != nil {
				print("Error: %s.\n\n", err.Error())

			} else {
				print("%s\n", dump())
			}
		}
	}
}
