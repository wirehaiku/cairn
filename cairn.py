'''
cairn.py: A basic prototype of Cairn in Python.
'''

import readline

### Global Variables

QUEUE = []
STACK = []
RGSTR = [0, 0, 0, 0, 0, 0, 0, 0]
COMMS = {}
FUNCS = {}

### Parsing Functions

def parse(string):
    return '\n'.join(
        line.split('//', 1)[0] for line in string.splitlines()
    ).upper().split()

### Evaluation Functions

def atomise(string):
    try:
        return int(string)
    except ValueError:
        return string

def evaluate(atom):
    print(f"eval = {repr(atom)}")

    if isinstance(atom, int):
        atom = max(0, min(255, atom))
        STACK.append(atom)

    elif isinstance(atom, str):
        if atom in COMMS:
            COMMS[atom]()

        elif atom in FUNCS:
            for elem in FUNCS[atom]:
                evaluate(elem)

### Command Functions

PUSH  = STACK.append
POP   = STACK.pop
BOOL  = lambda x: {True: 1, False: 0}[x]

def DEQTO(atom):
    global QUEUE
    i = QUEUE.index(atom)
    print(f"deqto = {i} q={QUEUE} atoms={QUEUE[:i]} newqueue={QUEUE[i+1:]}")
    atoms = QUEUE[:i]
    QUEUE = QUEUE[i+1:]
    return atoms

# integer commands
COMMS["ADD"] = lambda: PUSH(POP() + POP())
COMMS["SUB"] = lambda: PUSH(POP(-2) - POP())
COMMS["MOD"] = lambda: PUSH(POP(-2) % POP())
COMMS["GTE"] = lambda: PUSH(BOOL(POP(-2) >= POP()))
COMMS["LTE"] = lambda: PUSH(BOOL(POP(-2) <= POP()))

# memory commands
COMMS["CLR"] = lambda: STACK.clear()

# flow control commands

def DEF():
    name = QUEUE.pop(0)
    FUNCS[name] = DEQTO("END")

COMMS.update({
    "DEF": DEF,
})

### Main Runtime Functions

def main():
    global QUEUE, STACK, RGSTR, COMMS, FUNCS

    running = True
    print("Cairn prototype script.")

    while running:
        try:
            line = input(">>> ")
        except EOFError:
            line = ""
        except KeyboardInterrupt:
            print()
            raise SystemExit

        QUEUE.extend([atomise(elem) for elem in parse(line)])
        print(f"queue = {QUEUE}")

        while len(QUEUE) > 0:
            try:
                evaluate(QUEUE.pop(0))
            except Exception as exc:
                print(f"Error: {exc}.")

        print(f"stack = {STACK}")
        print(f"rgstr = {RGSTR}")
        if FUNCS:
            print(f"funcs = {FUNCS}")

if __name__ == "__main__":
    main()
