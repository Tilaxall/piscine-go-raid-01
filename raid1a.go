package main

import (
	"github.com/01-edu/z01"
)

func main() {
	Raid1a(5, 3)
	Raid1a(5, 1)
	Raid1a(1, 1)
}

func Raid1a(x, y int) {
	if x == 1 && y != 1 {
		for j := 1; j <= x; j++ {
			if j == 1 || j == y {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
	} else if x != 1 && y == 1 {
		for i := 1; i <= x; i++ {
			if i == 1 || i == x {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
	} else if x == 1 && y == 1 {
		z01.PrintRune('o')
	} else {
		for i := 1; i <= x; i++ {
			if i == 1 || i == x {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune('\n')
		for j := 1; j <= x; j++ {
			if j == 1 || j == x {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		for i := 1; i <= x; i++ {
			if i == 1 || i == x {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
	}
	z01.PrintRune('\n')
}
