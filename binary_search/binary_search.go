package binary_search

import "math"

// IMPLEMENTATION OF BINARY SEARCH
func binaryseacrh(arr []int, element int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == element {
			return mid
		}
		if arr[mid] < element {
			start = mid + 1
		}
		if arr[mid] > element {
			end = mid - 1
		}
	}
	return -1
}

// IMPLEMENTING RECURSIVE BINARY SEARCh
func recurisvebinarysearch(arr []int, element int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if arr[mid] == element {
		return mid
	}
	if arr[mid] > element {
		return recurisvebinarysearch(arr, element, start, mid-1)
	}
	return recurisvebinarysearch(arr, element, mid+1, end)
}

// FIND THE SQUARE ROOT OF A NUMBER USING BINARY SEARC H
func findRoot(x int) int {
	start := 0
	end := x
	ans := -1
	for start <= end {
		mid := start + (end-start)/2
		if mid <= x/mid {
			ans = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return ans
}

// FIND THE FIRST AND LAST OCCURENCE OF ELEMENT
func findfirstlast(arr []int, element int) []int {
	var ans []int
	start := 0
	end := len(arr) - 1
	first := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == element {
			first = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	start = 0
	end = len(arr) - 1
	second := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == element {
			second = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	ans = append(ans, first)
	ans = append(ans, second)
	return ans
}

// FIND THE INDEX AT WHICH WE CAN ADD THE ELEMENT
func findposition(arr []int, element int) int {
	index := -1
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := start + (end-start)/2
		if element <= arr[mid] {
			index = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return index
}

// CHECK WHETHER AN ELEMENT IS PRESENT IN 2D ARRAY OR NOT
func checkfor2d(matrix [][]int, element int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		start := 0
		end := cols - 1
		for start <= end {
			mid := start + (end-start)/2
			if matrix[i][mid] == element {
				return true
			}
			if element > matrix[i][mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}

// FIND THE PEAK ELEMENT IN MOUNTAIN ARRAY LINEAR TIME
func findpeakindex1(arr []int) int {
	for i := 1; i < len(arr)-1; i++ {
		if (arr[i] > arr[i+1]) && (arr[i] > arr[i-1]) {
			return i
		}
	}
	return -1
}

// FIND THE PEAK ELEMENT IN MOUTNAIN ARRAY BINARY SEARC H
func findpeakindex2(arr []int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := start + (end-start)/2
		if (arr[mid] > arr[mid+1]) && (arr[mid] > arr[mid-1]) {
			return mid
		}
		if arr[mid] > arr[mid-1] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// SEARCH FOR AN ELEMENT IN ROTATED SORTED ARRAY
func findinrotated(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] >= arr[start] {
			if target >= arr[start] && target < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target >= arr[mid] && target < arr[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

// FIND THE MINIUM ELEMENT IN ROTATED SORTED ARRAY
func findMin(arr []int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] > arr[end] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return arr[start]
}

// FIND THE KTH POSITIVE MISSING NUMBER
func findmissing(arr []int, k int) int {
	start := 0
	end := len(arr) - 1
	missing := -1
	for start <= end {
		mid := start + (end-start)/2
		missing = arr[mid] - (mid + 1)
		if k > missing {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return missing
}

// BOOK ALLOCATION PROBLEM
func ispossible(mid int, pages []int, n int, m int) bool {
	scnt := 0
	pcnt := 0
	for i := 0; i < n; i++ {
		if pages[i] > mid {
			return false
		}
		if pcnt+pages[i] <= mid {
			pcnt += pages[i]
		} else {
			scnt++
			if scnt > m {
				return false
			}
			pcnt = pages[i]
		}
	}
	return true
}
func findpages(pages []int, n int, m int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += pages[i]
	}
	maxpages := math.MinInt32
	for i := 0; i < n; i++ {
		page := pages[i]
		if page > maxpages {
			maxpages = page
		}
	}
	start := maxpages
	end := sum
	ans := -1
	for start <= end {
		mid := start + (end-start)/2
		check := ispossible(mid, pages, n, m)
		if check == true {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

// CAPACITY TO SHIP PAKCAGES IN D DAYS
func calctime(mid int, weights []int) int {
	time := 0
	current := 0
	for i := 0; i < len(weights); i++ {
		if current+weights[i] > mid {
			time += 1
			current = 0
		}
	}
	return time
}
func capacity(weights []int, days int) int {
	sum := 0
	maxweight := 0
	for i := 0; i < len(weights); i++ {
		weight := weights[i]
		sum += weight
		if weight > maxweight {
			maxweight = weight
		}
	}
	start := maxweight
	end := sum
	for start < end {
		mid := start + (end-start)/2
		time := calctime(mid, weights)
		if time > days {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

// FIND THE LOWER BOUND IN ARRAY USING BINARY SEARCH
func lowerbound(arr []int, element int) int {
	start := 0
	end := len(arr) - 1
	ans := -1
	for start <= end {
		mid := start + (end-start)/2
		if element >= arr[mid] {
			ans = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return ans
}

// FIND THE UPPER BOUND IN ARRAY USING BINARY SEARCH
func upperbound(arr []int, element int) int {
	start := 0
	end := len(arr) - 1
	ans := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] > element {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

// KOKO EATING BANANAS BRUTE FORCE
func kokoeating1(piles []int, hours int) int {
	maxpile := math.MinInt32
	for i := 0; i < len(piles); i++ {
		pile := piles[i]
		if pile > maxpile {
			maxpile = pile
		}
	}
	ans := maxpile
	for i := 0; i < maxpile; i++ {
		current := 0
		for j := 0; j < len(piles); j++ {
			current += piles[j] / i
		}
		if current <= hours {
			ans = i
		}
	}
	return ans
}

// KOKO EATING BANAS BINARY SEARCH
func caneat(mid int, piles []int, hours int) bool {
	current := 0
	for i := 0; i < len(piles); i++ {
		pile := piles[i]
		current += pile / mid
	}
	if current <= hours {
		return true
	}
	return false
}
func kokoeating2(piles []int, hours int) int {
	maxpile := math.MinInt32
	for i := 0; i < len(piles); i++ {
		pile := piles[i]
		if pile > maxpile {
			maxpile = pile
		}
	}
	start := 1
	end := maxpile
	ans := maxpile
	for start < end {
		mid := start + (end-start)/2
		check := caneat(mid, piles, hours)
		if check == true {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
