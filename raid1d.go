package main

import (
	"github.com/01-edu/z01"
)

func Raid1d(x, y int) {
	if x <= 0 || y <= 0 {
	} else {
		if x == 1 && y != 1 {
			for j := 1; j <= y; j++ {
				if j == y {
					z01.PrintRune(10)
				}
				if j == 1 {
					z01.PrintRune('A')
				} else if j == x {
					z01.PrintRune('A')
				} else {
					z01.PrintRune(10)
					z01.PrintRune('B')

				}
			}
		} else if x != 1 && y == 1 {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('A')
				} else if i == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
		} else if x == 1 && y == 1 {
			z01.PrintRune('A')
		} else {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('A')
				} else if i == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
			for j := 1; j <= y-2; j++ {
				for k := 1; k <= x; k++ {
					if k == 1 || k == x {
						z01.PrintRune('B')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune(10)
			}
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('A')
				} else if i == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}

			}
		}
	}
	z01.PrintRune(10)
}
func main() {

	//Raid1d(-5, 3)
	//Raid1d(8, 1)
	//Raid1d(1, 1)
	Raid1d(5, 4)
	//Raid1d(1, 5)
	//Raid1d(9, 9)
}
