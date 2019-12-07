package main

type OrbitTest struct {
	input    []string
	expected float64
}

var orbitTests = []OrbitTest{
	{[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}, 42},
}

var orbitTwoTests = []OrbitTest{
	{[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}, 4},
}
