package main

import "fmt"

func main() {
	fmt.Println("Hello 구구단, 세계")
	fmt.Println("guguDan")

	for i := 2; i <= 9; i++ { // 2부터 9까지 반복
		fmt.Print("\n")                   // 줄바꿈
		fmt.Printf(" **  %d 단 **\t\n", i) // 2, 3, 4, 5, 6, 7, 8, 9까지 출력

		for j := 1; j <= 9; j++ { // 1부터 9까지 반복
			fmt.Printf("%d * %d = %d\n", i, j, i*j) // 2 * 1 = 2, 2 * 2 = 4, 2 * 3 = 6, 2 * 4 = 8, 2 * 5 = 10, 2 * 6 = 12, 2 * 7 = 14, 2 * 8 = 16, 2 * 9 = 18

		}
	}
}
