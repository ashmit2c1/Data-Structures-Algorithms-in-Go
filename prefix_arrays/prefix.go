package prefix_arrays

import (
	"math"
)

// PAIR STRUCTURE
type pair struct {
	first  int
	second int
}

// STATIC SUM RANGE QUERY BRUTE FORCE
func calcfunction1(left int, right int, arr []int) int {
	sum := 0
	for i := left; i < right; i++ {
		sum += arr[i]
	}
	return sum
}
func staticrange1(arr []int, queries []pair) []int {
	var ans []int
	for i := 0; i < len(queries); i++ {
		left := queries[i].first - 1
		right := queries[i].second - 1
		answer := calcfunction1(left, right, arr)
		ans = append(ans, answer)
	}
	return ans
}

// STATIC SUM RANGE QUERY OPTIMSIED
func calcfunction2(left int, right int, pref []int) int {
	if left == 0 {
		return pref[right]
	} else {
		return pref[right] - pref[left-1]
	}
}
func staticrange2(arr []int, queries []pair) []int {
	var ans []int
	pref := arr
	for i := 1; i < len(pref); i++ {
		pref[i] += pref[i-1]
	}
	for i := 0; i < len(queries); i++ {
		left := queries[i].first - 1
		right := queries[i].second - 1
		answer := calcfunction2(left, right, pref)
		ans = append(ans, answer)
	}
	return ans
}

// FIND THE NUMBER OF SUB-ARRAYS WHOSE SUM IS EQUAL TO X BRUTE FORCE
func findsub1(arr []int, target int) int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := 0; j < len(arr); j++ {
			sum += arr[j]
			if sum == target {
				cnt++
			}
		}
	}
	return cnt
}

// FIND THE NUMBER OF SUB-ARRAYS WHOSE SUM IS EQUAL TO X OPTIMISED
func findsub2(arr []int, target int) int {
	cnt := 0
	mp := make(map[int]int)
	currentsum := 0
	mp[0] = 1
	for i := 0; i < len(arr); i++ {
		currentsum += arr[i]
		_, ok := mp[currentsum-target]
		if ok == true {
			cnt += mp[currentsum-target]
		}
		mp[currentsum]++
	}
	return cnt
}

// FIND THE INDEX WHERE THE PREFIX SUM AND SUFFIX SUM ARE MINMUM
func solvefunction(arr []int) int {
	prefix := arr
	suffix := arr
	for i := 1; i < len(arr); i++ {
		prefix[i] += prefix[i-1]
	}
	for i := len(suffix) - 2; i >= 0; i-- {
		suffix[i] += suffix[i+1]
	}
	index := -1
	minsum := math.MaxInt32
	for i := 0; i < len(arr); i++ {
		sum := prefix[i] + suffix[i]
		if sum < minsum {
			minsum = sum
			index = i
		}
	}
	return index
}

// FIND THE INDEX WHERE THE SUFFIX AND PREFIX SUM ARE MINIMUM
func findmin(arr []int) int {
	minel := math.MaxInt32
	index := -1
	for i := 0; i < len(arr); i++ {
		el := arr[i]
		if el < minel {
			minel = el
			index = i
		}
	}
	return index
}

// FIND THE COUNT OF SUB-ARRAYS WHOSE SUM IS EQUAL TO THE LENGTH OF THE SUBARRAY
func findsub3(arr []int) int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			len := j - i + 1
			if sum == len {
				cnt++
			}
		}
	}
	return cnt
}

// FIND THE COUNT OF SUBARRAYS WHOSE SUM IS EQUAL TO THE LENGTH OF THE SUBARRAY
func findsub4(arr []int) int {
	cnt := 0
	mp := make(map[int]int)
	current := 0
	mp[0] = 1
	for i := 0; i < len(arr); i++ {
		current += arr[i]
		_, ok := mp[current-(i+1)]
		if ok == true {
			cnt += mp[current-(i+1)]
		}
		mp[current]++
	}
	return cnt
}

// FIND THE MAX GCD THAT CAN BE OBTAINED BY REMOVING ONE ELEMENT OF THE ARRAY
func calcf(a int, b int) int {
	for b != 0 {
		rem := a % b
		a = b
		b = rem
	}
	return a
}
func findf(arr []int) int {
	ans := arr[0]
	for i := 1; i < len(arr); i++ {
		ans = calcf(ans, arr[i])
	}
	return ans
}
func findmaxgcd(arr []int) int {
	ans := math.MinInt32
	for i := 0; i < len(arr); i++ {
		var newarr []int
		for j := 0; j < len(arr); j++ {
			if i != j {
				newarr = append(newarr, arr[j])
			}
		}
		gcd := findf(newarr)
		if gcd > ans {
			ans = gcd
		}
	}
	return ans
}

// FIND THE MAX GCD THAT CAN BE OBTAINED BY REMOVING ONE ELEMENT FROM THE ARAY OPTIMSISED
func findmaxgcd2(arr []int) int {
	n := len(arr)
	prefix := make([]int, n)
	suffix := make([]int, n)
	for i := 0; i < n; i++ {
		prefix[i] = 0
		suffix[i] = 0
	}
	ans := math.MinInt32
	for i := 1; i < len(arr); i++ {
		prefix[i] = calcf(prefix[i-1], arr[i])
	}
	for i := n - 2; i >= 0; i-- {
		suffix[i] = calcf(suffix[i+1], arr[i])
	}
	for i := 0; i < n; i++ {
		var gcdex int
		if i == 0 {
			gcdex = suffix[i+1]
		}
		if i == n-1 {
			gcdex = prefix[i-1]
		} else {
			gcdex = calcf(prefix[i-1], suffix[i+1])
		}
		if gcdex > ans {
			ans = gcdex
		}
	}
	return ans
}

// FIND THE MAX POPULATION IN GIVEN YEAR
func maxpop(logs [][]int) int {
	pop := make([]int, 101)
	for i := 0; i < len(pop); i++ {
		pop[i] = 0
	}
	year := 0
	current := 0
	maxpop := math.MinInt32
	for i := 0; i < len(pop); i++ {
		current += pop[i]
		if current > maxpop {
			maxpop = current
			year = i
		}
	}
	year += 1950
	return year
}

// GIVEN AN ARRAY OF N INTEGERS, FIND THE COUNT OF SUB-ARRAYS THAT ARE DIVISIBLE BY N
func findcnt(arr []int) int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum%len(arr) == 0 {
				cnt++
			}
		}
	}
	return cnt
}

// GIVEN AN ARRAY OF N INTEGERS, FIND THE COUNT OF SUB-ARRAYS THAT ARE DIVISIBLE BY N OPTIMSOIED WAY
func findcnt2(arr []int) int {
	cnt := 0
	n := len(arr)
	mp := make(map[int]int)
	current := 0
	mp[0] = 1
	for i := 0; i < len(arr); i++ {
		current += arr[i]
		mod := (current%n + n) % n
		_, ok := mp[mod]
		if ok == true {
			cnt += mp[mod]
		}
		mp[mod]++
	}
	return cnt
}

// COUNT THE NUMBER OF NICE SUBARRAYS, A NICE SUBARRAY IS THE ONE THAT HAS EXACTLY K ODD NUMBERS
func cntnice(arr []int, target int) int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			arr[i] = 0
		} else {
			arr[i] = 1
		}
	}
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := 0; j < len(arr); j++ {
			sum += arr[j]
			if sum == target {
				cnt++
			}
		}
	}
	return cnt
}

// COUNT THE NUMBER OF NICE SUBARRAYS A NICE SUBARRAY IS THE ONE THAT HAS K EXACTLY
func cntnice2(arr []int, k int) int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			arr[i] = 0
		} else {
			arr[i] = 1
		}
	}
	mp := make(map[int]int)
	current := 0
	mp[0] = 1
	for i := 0; i < len(arr); i++ {
		current += arr[i]
		_, ok := mp[current-k]
		if ok == true {
			cnt += mp[current-k]
		}
		mp[current]++
	}
	return cnt
}

// FIND THE MAX LENGTH SUB-ARRAY WHICH HAS EQUAL NUMBER OF 1 AND 0
func findmax(arr []int) int {
	ans := math.MinInt32
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			if arr[j] == 0 {
				arr[j] = -1
			}
			sum += arr[j]
			if sum == 0 {
				newlen := j - i + 1
				if newlen > ans {
					ans = newlen
				}
			}
		}
	}
	return ans
}

// FIND THE MAX LENGTH OF THE SUBARRAY WHICH HAS EQUAL NUMBER OF 1 AND  0 OPTIMISED
func findmax2(arr []int) int {
	current := 0
	maxlen := math.MinInt32
	mp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		current += arr[i]
		if current == 0 {
			newlen := i + 1
			if newlen > maxlen {
				maxlen = newlen
			}
		}
		_, ok := mp[current]
		if ok == true {
			newlen := i - mp[current]
			if newlen > maxlen {
				maxlen = newlen
			}
		}
		mp[current] = i
	}
	return maxlen
}

// CREATE A PREFIX 2D ARRAY
func create2dpref(mat [][]int) [][]int {
	pref := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		pref[i] = make([]int, len(mat[i]))
	}
	for i := 0; i < len(mat); i++ {
		pref[i] = make([]int, len(mat[i]))
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			pref[i][j] = mat[i][j]
			if i > 0 {
				pref[i][j] += pref[i-1][j]
			}
			if j > 0 {
				pref[i][j] += pref[i][j-1]
			}
			if i > 0 && j > 0 {
				pref[i][j] -= pref[i-1][j-1]
			}
		}
	}
	return pref
}

// HANDLE QUERIES IN 2d prefix array
func handle(coords []pair, pref [][]int) int {
	x1 := coords[0].first
	y1 := coords[0].second
	x2 := coords[1].first
	y2 := coords[1].second

	ans := pref[x2][y2]
	if x1 > 0 {
		ans -= pref[x1-1][y2]
	}
	if y1 > 0 {
		ans -= pref[x2][y1-1]
	}
	if x1 > 0 && y1 > 0 {
		ans += pref[x1-1][y1-1]
	}
	return ans
}

// FOREST QUERIES, GIVEN A A FOREST FIND THE NUMBER OF TREES PRESENT IN THE QUERIED RECTANGLE
func findforest(pref [][]int, x1 int, y1 int, x2 int, y2 int) int {
	ans := pref[x2][y2]
	if x1 > 0 {
		ans -= pref[x1-1][y2]
	}
	if y1 > 0 {
		ans -= pref[x2][y1-1]
	}
	if x1 > 0 && y1 > 0 {
		ans += pref[x1-1][y1-1]
	}
	return ans
}
func findans(forest [][]int, queries [][]pair) []int {
	pref := make([][]int, len(forest))
	for i := 0; i < len(forest); i++ {
		pref[i] = make([]int, len(forest[i]))
	}
	for i := 0; i < len(pref); i++ {
		for j := 0; j < len(pref[i]); j++ {
			pref[i][j] = forest[i][j]
			if i > 0 {
				pref[i][j] += pref[i-1][j]
			}
			if j > 0 {
				pref[i][j] += pref[i][j-1]
			}
			if i > 0 && j > 0 {
				pref[i][j] -= pref[i][j]
			}
		}
	}
	var ans []int
	for i := 0; i < len(queries); i++ {
		x1 := queries[i][0].first
		y1 := queries[i][0].second
		x2 := queries[i][1].first
		y2 := queries[i][1].second
		answer := findforest(pref, x1, y1, x2, y2)
		ans = append(ans, answer)
	}
	return ans
}

// DIFFERENCE ARRAYS
func operation(arr []int, left int, right int, val int) {
	for i := left; i < right; i++ {
		arr[i] += val
	}
}
func diffarr(arr []int, queries []pair, change []int) {
	for i := 0; i < len(queries); i++ {
		left := queries[i].first - 1
		right := queries[i].second - 1
		val := change[i]
		operation(arr, left, right, val)
	}
}

// DIFFERENCE ARRAYS
func operation2(left int, right int, val int, temp []int) {
	temp[left] += val
	n := len(temp)
	if right+1 < n {
		temp[right+1] -= val
	}
}
func diffarr1(arr []int, queries []pair, change []int) {
	n := len(arr)
	temp := make([]int, n)
	for i := 0; i < len(temp); i++ {
		temp[i] = 0
	}
	for i := 0; i < len(queries); i++ {
		left := queries[i].first - 1
		right := queries[i].second - 1
		val := change[i]
		operation2(left, right, val, temp)
	}
	for i := 1; i < len(temp); i++ {
		temp[i] += temp[i-1]
	}
	for i := 0; i < len(arr); i++ {
		arr[i] += temp[i]
	}
}
