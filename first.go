package main //ezen meg itt köll mindig

//ez a math meg a const a konstansoknál kell

//"math"

//const s string = "constant"

func main() {
	//HELLO WORD
	//fmt.Println("hello word")

	//VALUES---------------------------------
	//fmt.Println("go" + "lang")
	//fmt.Println("1+1 =", 1+1)
	//fmt.Println("7.0/3.0 = ", 7.0/3.0)

	//fmt.Println(true && false)
	//fmt.Println(true || false)
	//fmt.Println(!true)

	//VARIABLES-------------------------------
	//var a = "initial"
	//fmt.Println(a)

	//var b, c int = 1, 2
	//fmt.Println(b, c)

	//var d = true
	//fmt.Println(d)

	//var e int
	//fmt.Println(e)

	//f := "apple"
	//fmt.Println(f)

	//CONSTANS-------------------------------
	//fmt.Println(s)

	//const n = 500000000

	//const d = 3e20 / n
	//fmt.Println(d)

	//fmt.Println(int64(d))
	//fmt.Println(math.Sin(n))

	//FOR LOOPS------------------------------
	//i := 1
	//for i <= 3 {
	//fmt.Println(i)
	//i += 1
	//}

	//for j := 0; j < 3; j++ {
	//	fmt.Println(j)
	//}

	//for i := range 3 {
	//	fmt.Println("range", i)
	//}

	//for {
	//	fmt.Println("loop")
	//	break
	//}

	//for n := range 6 {
	//	if n%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(n)
	//}

	//IF/ELSE--------------------------------------
	//if 7%2 == 0 {
	//	fmt.Println("7 is even")
	//} else {
	//	fmt.Println("7 is odd")
	//}

	//if 8%4 == 0 {
	//	fmt.Println("8 is divisible by 4")
	//}

	//if 8%2 == 0 || 7%2 == 0 {
	//	fmt.Println("either 8 or 7 are even")
	//}

	//if num := 9; num < 0 {
	//	fmt.Println(num, "is negative")
	//} else if num < 10 {
	//	fmt.Println(num, "has 1 digit")
	//} else {
	//	fmt.Println(num, "has multiple digits")
	//}

	//SWITCH--------------------------------------------------------------
	//i := 2
	//fmt.Print("Write ", i, " as ")
	//switch i {
	//case 1:
	//	fmt.Println("one")
	//case 2:
	//	fmt.Println("two")
	//case 3:
	//	fmt.Println("three")
	//}

	//switch time.Now().Weekday() {
	//case time.Saturday, time.Sunday:
	//	fmt.Println("it's weekend")
	//default:
	//	fmt.Println("its weekday")
	//}

	//t := time.Now()
	//switch {
	//case t.Hour() < 12:
	//	fmt.Println("It's before noon")
	//default:
	//	fmt.Println("It's after noon")
	//}

	//whatAmI := func(i interface{}) {
	//	switch t := i.(type) {
	////	case bool:
	//		fmt.Println("I'm a bool")
	//	case int:
	//		fmt.Println("I'm an int")
	//	default:
	//		fmt.Printf("Don't know type %T\n", t)
	//	}
	//}
	//whatAmI(true)
	//whatAmI(1)
	//whatAmI("hey")

	//ARRAYS--------------------------------------------------------
	//var a [5]int
	//fmt.Println("emp: ", a)

	//a[4] = 100
	//fmt.Println("set: ", a)
	//fmt.Println("get: ", a[4])

	//fmt.Println("len: ", len(a))

	//b := [5]int{1, 2, 3, 4, 5}
	//fmt.Println("dcl: ", b)

	//b = [...]int{1, 2, 3, 4, 5}
	//fmt.Println("dcl: ", b)

	//b = [...]int{100, 3: 400, 500}
	//fmt.Println("idx: ", b)

	//var twoD [2][3]int
	//for i := 0; i < 2; i++ {
	//	twoD = [2][3]int{
	//		{1, 2, 3},
	//		{1, 2, 4},
	//	}
	//	fmt.Println("2d: ", twoD)
	//}

	//SLICES (MOST POPULAR)
}
