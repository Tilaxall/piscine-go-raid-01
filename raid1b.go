package student

import (
	"github.com/01-edu/z01"
)

func Raid1b(x, y int) {
	if x <= 0 || y <= 0 {
	} else {
		if x == 1 && y != 1 {
			for j := 1; j <= y; j++ {
				if j == y {
					z01.PrintRune(10)
				}
				if j == 1 {
					z01.PrintRune('/')
				} else if j == x {
					z01.PrintRune(92)
				} else {
					z01.PrintRune(10)
					z01.PrintRune('*')

				}
			}
		} else if x != 1 && y == 1 {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('/')
				} else if i == x {
					z01.PrintRune(92)
				} else {
					z01.PrintRune('*')
				}
			}
		} else if x == 1 && y == 1 {
			z01.PrintRune('/')
		} else {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('/')
				} else if i == x {
					z01.PrintRune(92)
				} else {
					z01.PrintRune('*')
				}
			}
			z01.PrintRune('\n')
			for j := 1; j <= y-2; j++ {
				for k := 1; k <= x; k++ {
					if k == 1 || k == x {
						z01.PrintRune('*')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune(10)
			}
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune(92)
				} else if i == x {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('*')
				}

			}
		}
	}
	z01.PrintRune(10)
}
