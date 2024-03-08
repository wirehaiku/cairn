package cairn

// Library is a string containing Cairn standard library definitions.
const Library = `
// Library Helpers //

def . // No-op for use as a code separator.
	nop
end

def rclr // Clear registers 0 and 1.
	0 0 set . 0 1 set
end

// Stack Functions //

def dup // (a -- a a) Duplicate the top integer.
	0 set . 0 get 0 get . rclr
end

def drop // (a b -- a) Delete the top integer.
	0 set . rclr
end

def swap // (a b -- b a) Swap the top two integers.
	0 set 1 set . 0 get 1 get . rclr
end

// Operator Functions //

def != // (a b -- b) Return true if a != b.
	== f?
end

def <= // (a b -- c) Return true if a <= b.
	> f?
end

def >= // (a b -- c) Return true if a >= b.
	< f?
end

// Logic Functions //

def f? // (a -- b) Return true if a is false.
	0 ==
end

def t? // (a -- b) Return true if a is true.
	0 >
end

def and // (a b -- c) Return true if a and b are true.
	t? . swap t? . + 2 ==
end

def or // (a b -- c) Return true if a or b are true.
	t? . swap t? . + t?
end

def xor // (a b -- c) Return true if only a or b are true.
	t? . swap t? . + 1 ==
end
`
