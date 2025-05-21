package main


import "fmt"


func main() {
	//TEST 1
	sequence:= adder() //since the function value is only returned here (which is another function value), the i:=0 does not
			   //happern again, meaning that i value does not get reset. and the previous i value is stored in memory. 
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2:= adder()
	fmt.Println(sequence2())
	
	//TEST 2
	subtracter := func() func(int) int{
		countdown:= 99			//we are keeping the countdown altered value in memory. Every time we run this.
		return func(x int) int{
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(1))

}

func adder () func() int {
	i := 0
	fmt.Println("Previous value of i:", i)
	return func() int{
		i++
		fmt.Println("added 1 to i")
		return i
	}
	

}
