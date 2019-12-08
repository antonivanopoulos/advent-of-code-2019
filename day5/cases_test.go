package main

type IntcodeTest struct {
	input    int
	program  []int
	expected []int
	output   string
}

var intcodeTests = []IntcodeTest{
	{1, []int{3, 0, 4, 0, 99}, []int{1, 0, 4, 0, 99}, "Output: 1"},
	{0, []int{1002, 4, 3, 4, 33}, []int{1002, 4, 3, 4, 99}, ""},
}
