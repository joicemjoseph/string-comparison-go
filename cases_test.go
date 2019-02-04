package exam

var testCases = []struct {
	A        string
	B        string
	Expected int
}{
	{ // empty strands
		"",
		"",
		1,
	},
	{ // empty strands
		"This is some statement",
		"is",
		2,
	},
}
