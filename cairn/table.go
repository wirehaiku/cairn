package cairn

// Table is an addressable map of stored integers.
type Table struct {
	Integers map[int]int
}

// NewTable returns a pointer to a new Table.
func NewTable(im map[int]int) *Table {
	return &Table{im}
}

// Clear removes all entries from the Table.
func (t *Table) Clear() {
	t.Integers = make(map[int]int)
}

// Delete removes an entry from the Table.
func (t *Table) Delete(i int) {
	delete(t.Integers, i)
}

// Empty returns true if the Table has no entries.
func (t *Table) Empty() bool {
	return len(t.Integers) == 0
}

// Get returns the value of an entry in the Table.
func (t *Table) Get(i int) int {
	return t.Integers[i]
}

// Has returns true if the Table contains an entry.
func (t *Table) Has(i int) bool {
	_, ok := t.Integers[i]
	return ok
}

// Len returns the number of entries in the Table.
func (t *Table) Len() int {
	return len(t.Integers)
}

// Set sets the value of an entry in the Table.
func (t *Table) Set(i, v int) {
	t.Integers[i] = v
}
