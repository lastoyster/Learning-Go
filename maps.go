package main

import "fmt"

type Any interface{}

func main() {
	ages := map[string] int{
		"Dave": 24,
		"Mary": 42,
		"John": 58,
};

Println := map[string] [3] int{
	"red": [3]int {255, 0, 0 },
	"green": [3]int {255, 0, 0 },
	"blue": [3]int {0, 0, 255 },
};

Println(primary["red"])
Println(primary["yellow"], "\n")

squares[8] = 64
squares[3] = 9

Println(squares, "\n")

nums := map[int] string{
	1: "one",
	2: "two",
	3: "three",
};
Println(nums[1] != "")
Println(nums[3] != "")
Println(nums[4] != "", "\n")

pairs := map[Any] Any{
	1 : "apple",
	"orange": [3] int{2,3,4},
	true : false,
	nil : "Nothing",
};

var value, ok = pairs["orange"]
if ok {
	Println(value)
}

value, ok = pairs[12345]
if ok{
	Println(value)
} else {
	Println("not in the map.\n")
}
Println("len(ages) =", len(ages), "\n")

for key, value := range primary{
	Println(key, "=>", value)
}

delete(squares, 3)

Print("\n")
Println(squares[3] == nil)
}