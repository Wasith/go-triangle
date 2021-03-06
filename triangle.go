package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	MAXIMUM_SIZE       = 3
	MAXIMUM_ARGRUMENTS = MAXIMUM_SIZE + 1
	MAXIMUM_VALUE      = 200
	MINNIMUM_VALUE     = 1
)

func main() {
	if ShowHelpIfArgsMissMatch() {
		return
	}

	var sides [MAXIMUM_SIZE]int
	var isInteger bool
	if sides, isInteger = IsAllArgsAreInt(); !isInteger {
		fmt.Println("[!] Please enter all 3 parameters with integer value.")
		return
	}

	if IsLessThanMinimum(sides) {
		fmt.Println("[!] All 3 argruments must be positive value.")
		return
	}

	if IsAnyValueMoreThanMaximumRange(sides, MAXIMUM_VALUE) {
		fmt.Println("[!] There is some value is out of range.")
		return
	}

	if SumOfTwoSidesOfTriangleMustBeMoreThanTheLeftSide(sides) {
		fmt.Println("[!] Not a triangle.")
		return
	}

	if IsAnyRightTriangle(sides) {
		fmt.Println("[A] Right Triangle.")
		return
	}

	if IsEquilateral(sides) {
		fmt.Println("[A] Equilateral.")
		return
	}

	if IsAnyIsosceles(sides) {
		fmt.Println("[A] Isosceles.")
		return
	}

	if IsAnyScalene(sides) {
		fmt.Println("[A] Scalene.")
		return
	}
}

func ShowHelpIfArgsMissMatch() bool {
	if len(os.Args) != MAXIMUM_ARGRUMENTS {
		fmt.Println("[I] Please enter 3 parameters for 3 sides of triangle.")
		fmt.Println("    eg. " + os.Args[0] + " 3 4 5")
		return true
	}
	return false
}

func IsAllArgsAreInt() (sides [MAXIMUM_SIZE]int, isTriangle bool) {
	for i := 1; i < MAXIMUM_SIZE+1; i++ {
		side, err := strconv.Atoi(os.Args[i])
		if err != nil {
			isTriangle = false
			return
		} else {
			sides[i-1] = side
		}
	}
	isTriangle = true
	return
}

func IsLessThanMinimum(sides [MAXIMUM_SIZE]int) bool {
	for _, sideWidth := range sides {
		if sideWidth < MINNIMUM_VALUE {
			return true
		}
	}
	return false
}

func IsAnyValueMoreThanMaximumRange(sides [MAXIMUM_SIZE]int, maximumValueValidable int) bool {
	for _, sideWidth := range sides {
		if sideWidth > maximumValueValidable {
			return true
		}
	}
	return false
}

func SumOfTwoSidesOfTriangleMustBeMoreThanTheLeftSide(sides [MAXIMUM_SIZE]int) bool {
	if IsSumOfABIsLessThanC(sides[0], sides[1], sides[2]) ||
		IsSumOfABIsLessThanC(sides[0], sides[2], sides[1]) ||
		IsSumOfABIsLessThanC(sides[1], sides[2], sides[0]) {
		return true
	}
	return false
}

func IsSumOfABIsLessThanC(a int, b int, c int) bool {
	if a+b <= c {
		return true
	}
	return false
}

func IsAnyRightTriangle(sides [MAXIMUM_SIZE]int) bool {
	aSqrt := PowerTwo(sides[0])
	bSqrt := PowerTwo(sides[1])
	cSqrt := PowerTwo(sides[2])

	if IsAPlusBEqualToC(aSqrt, bSqrt, cSqrt) ||
		IsAPlusBEqualToC(aSqrt, cSqrt, bSqrt) ||
		IsAPlusBEqualToC(bSqrt, cSqrt, aSqrt) {
		return true
	}
	return false
}

func IsAPlusBEqualToC(a int, b int, c int) bool {
	return a+b == c
}

func PowerTwo(number int) int {
	return number * number
}

func IsEquilateral(sides [MAXIMUM_SIZE]int) bool {
	for i := 1; i < len(sides); i++ {
		if sides[0] != sides[i] {
			return false
		}
	}
	return true
}

func IsAnyIsosceles(sides [MAXIMUM_SIZE]int) bool {
	if IsIsosceles(sides[0], sides[1], sides[2]) ||
		IsIsosceles(sides[0], sides[2], sides[1]) ||
		IsIsosceles(sides[1], sides[2], sides[0]) {
		return true
	}
	return false
}

func IsIsosceles(a int, b int, c int) bool {
	return a == b && a+b > c
}

func IsAnyScalene(sides [MAXIMUM_SIZE]int) bool {
	if IsScalene(sides[0], sides[1], sides[2]) {
		return true
	}
	return false
}

func IsScalene(a int, b int, c int) bool {
	return a != b && a != c && b != c
}
