package main

import "fmt"

func cubeRoot(num int, p int) float64 {
	s := 0
	e := num
	cubicRoot := 0.0
	for e>= s{
		m := s + (e-s)/2

		if m*m*m == num{
			cubicRoot = float64(m)
			return float64(m);
		} else {
		   e = m-1
		}
	
		for i:=0; i<p; i++ {
			incr:= 0.001;
			for cubicRoot*cubicRoot* cubicRoot <=float64(num){
				cubicRoot += incr;
			}

			cubicRoot -= incr;
			incr /=10;
		}
}

return cubeRoot
} 

func main() {
	var num int
	fmt.Scanln(&num)
	if(num==0){
		num = 64
		fmt.Printf("Invalid value ")
	}
	cube := cubeRoot(num, 3)

	fmt.Printf("\nThe cube root of %0.3f\n\n",num,cube)

	fmt.Println(`Thank You Google 🙏`)
}
