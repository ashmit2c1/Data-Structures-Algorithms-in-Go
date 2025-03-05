package recursion

import (
	"fmt"
	"math"
)

// FIND THE SUM OF FIRST N NATURAL NUMBERS
func findsumrecursive(n int) int {
	if n == 0 {
		return 0
	}
	return n + findsumrecursive(n-1)

}
func findsum(n int) int {
	ans := findsumrecursive(n)
	return ans
}

// FIND THE FACTORIAL OF A GIVEN NUNBER
func findfact1(n int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}
	return ans
}

// FIND THE FACTORIAL OF A GIVEN NUMBER USING RECURSION
func findfactrecursive(n int) int {
	if n == 1 {
		return 1
	}
	return n * findfactrecursive(n-1)

}
func findfact2(n int) int {
	ans := findfactrecursive(n)
	return ans
}

// FIND THE FIBONACCI SEQUENCE USING RECURSION
func findfibrec(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return findfibrec(n-1) + findfibrec(n-2)
}
func findfib(n int) int {
	ans := findfibrec(n)
	return ans
}

// PRINT THE PATTERN
// ****
// ***
// **
// *
func printpatternrec(n int) {
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
func printpattern(n int) {
	printpatternrec(n)
}

// LINEAAR SEARCH USING RECURSION
func linear_search_rec(arr []int, index int, target int) int {
	if index == len(arr) {
		return -1
	}
	if arr[index] == target {
		return index
	}
	return linear_search_rec(arr, index+1, target)
}
func linear_search(arr []int, target int) int {
	ans := linear_search_rec(arr, 0, target)
	return ans
}
func checkrec(s string, start int, end int) bool {
	if start == end {
		return true
	}
	if s[start] != s[end] {
		return false
	}
	return checkrec(s, start+1, end-1)
}

// CHECK IF THE GIVEN STRING IS PALINDROME OR NOT
func check(s string) bool {
	ans := checkrec(s, 0, len(s)-1)
	return ans
}

// FIND ALL THE SUBSETS OF THE ARRAY BRUTE FORCE
func findsubsets1(arr []int) [][]int {
	var ans [][]int
	for i := 0; i < len(arr); i++ {
		var temp []int
		for j := i; j < len(arr); j++ {
			temp = append(temp, arr[j])
			ans = append(ans, temp)
		}
	}
	return ans
}

// FIND SUBSETS IN ARRAY USING RECURSION / BACKTRAACKING
func findsubsetfunction(ans *[][]int, current *[]int, arr []int, index int) {
	(*ans) = append((*ans), *(current))
	for i := index; i < len(arr); i++ {
		(*current) = append((*current), arr[i])
		findsubsetfunction(ans, current, arr, i+1)
		(*current) = (*current)[:len((*current))-1]
	}
}
func findsubsets2(arr []int) [][]int {
	var ans [][]int
	var current []int
	findsubsetfunction(&ans, &current, arr, 0)
	return ans
}

// FIND ONLY UNIQUE SUBSETS IN ARRAY
func gensubset2(current *[]int, ans *[][]int, index int, arr []int) {
	(*ans) = append((*ans), (*current))
	for i := index; i < len(arr); i++ {
		if i > index && arr[i-1] == arr[i] {
			continue
		}
		(*current) = append((*current), arr[i])
		gensubset2(current, ans, i+1, arr)
		(*current) = (*current)[:len((*current))-1]
	}
}
func findsubset(arr []int) [][]int {
	var current []int
	var ans [][]int
	gensubset2(&current, &ans, 0, arr)
	return ans
}

// FIND A POWER B USING RECURSION
func findansfun(a int, b int) int {
	if b == 1 {
		return a
	}
	return a * findansfun(a, b-1)
}
func findans(a int, b int) int {
	ans := findansfun(a, b)
	return ans
}

// FND A POWER B USING RECURSION AND IN LOG N TIME
func findans2(a int, b int) int {
	if b == 1 {
		return a
	}
	ans := findans2(a, b/2)
	ans *= ans
	if b%2 == 0 {
		return ans
	}
	return ans * a
}

// GENERATE THE BINARY STRING FOR THE NUMBER N
func findstring(n int) string {
	if n == 0 {
		return "0"
	}
	ans := findstring(n / 2)
	if n%2 == 1 {
		ans += "1"
	} else {
		ans += "1"
	}
	return ans
}
func genstring(n int) string {
	ans := findstring(n)
	return ans
}

// GENERATE ALL BINARY STRINGS OF SIZE N WITHOUT ADJACENT ZEROES
func genstringfun(n int, ans *[]string, current string) {
	if len(current) == n {
		(*ans) = append((*ans), current)
	}
	genstringfun(n, ans, current+"1")
	if len(current) == 0 || current[len(current)-1] != '0' {
		genstringfun(n, ans, current+"0")
	}
}
func validstring(n int) []string {
	var ans []string
	genstringfun(n, &ans, "")
	return ans
}

// GIVEN A MATRIX FIND THE MAX PATH SUM
func findmaxpathfunc(i int, j int, mat [][]int) int {
	if i < 0 || j < 0 || i > len(mat) || j > len(mat[0]) {
		return math.MinInt32
	}
	if i == 0 && j == 0 {
		return mat[0][0]
	}
	goingup := findmaxpathfunc(i-1, j, mat)
	goingleft := findmaxpathfunc(i, j-1, mat)
	var pathtotake int
	if goingleft > goingup {
		pathtotake = goingleft
	} else {
		pathtotake = goingup
	}
	return mat[i][j] + pathtotake
}
func findmaxpath(mat [][]int) int {
	r := len(mat)
	c := len(mat[0])
	ans := findmaxpathfunc(r-1, c-1, mat)
	return ans
}

// GIVEN A MATRIX FIND THE MIN PATH SUM
func findminpathfunc(i int, j int, mat [][]int) int {
	if i < 0 || j < 0 || i > len(mat) || j > len(mat[0]) {
		return math.MaxInt32
	}
	goingup := findminpathfunc(i-1, j, mat)
	goingleft := findminpathfunc(i, j-1, mat)
	var pathtotake int
	if goingup < goingleft {
		pathtotake = goingup
	} else {
		pathtotake = goingleft
	}
	return mat[i][j] + pathtotake

}
func findminpath(mat [][]int) int {
	r := len(mat)
	c := len(mat[0])
	ans := findminpathfunc(r-1, c-1, mat)
	return ans
}

// SUDOKU SOLVER RECURSION
func isvalid(num int, row int, col int, board [][]int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
		if board[row][i] == num {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}
func solve(board [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					check := isvalid(num, row, col, board)
					if check == true {
						board[row][col] = num
						if solve(board) == true {
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}
