package main

type PossiblePasswordTest struct {
	input    string
	expected bool
}

var possiblePasswordTests = []PossiblePasswordTest{
	{"111111", false},
	{"223450", false},
	{"123789", false},
	{"112233", true},
	{"123444", false},
	{"111122", true},
	{"112222", true},
}
