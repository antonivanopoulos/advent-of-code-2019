package main

type ManhattenWiringDistanceTest struct {
	input    []string
	expected float64
}

var manhattentWiringDistanceTests = []ManhattenWiringDistanceTest{
	{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 159},
	{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 135},
}

type ManhattenWiringDelayTest struct {
	input    []string
	expected int
}

var manhattentWiringDelayTests = []ManhattenWiringDelayTest{
	{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 610},
	{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 410},
}
