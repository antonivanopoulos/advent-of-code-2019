package main

type ModuleFuelTest struct {
	input    int
	expected int
}

var moduleFuelTests = []ModuleFuelTest{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}
