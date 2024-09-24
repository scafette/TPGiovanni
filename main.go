package main

import (
	"fmt"
)

func Ft_coin(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
func Ft_missing(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return total - sum
}
func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sortIntervals(intervals)
	end := intervals[0][1]
	count := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}
	return count
}
func Ft_profit(prices []int) int {
	minPrice := int(^uint(0) >> 1) // Equivalent to math.MaxInt32
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
func Ft_max_substring(s string) int {
	charIndex := make(map[rune]int)
	start, maxLength := 0, 0
	for i, c := range s {
		if lastSeenIndex, found := charIndex[c]; found && lastSeenIndex >= start {
			start = lastSeenIndex + 1
		}
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
		charIndex[c] = i
	}
	return maxLength
}
func Ft_min_window(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	tFreq := make(map[rune]int)
	windowFreq := make(map[rune]int)
	for _, c := range t {
		tFreq[c]++
	}
	left, right, start, minLen, valid := 0, 0, 0, len(s)+1, 0
	for right < len(s) {
		char := rune(s[right])
		right++
		if tFreq[char] > 0 {
			windowFreq[char]++
			if windowFreq[char] == tFreq[char] {
				valid++
			}
		}
		for valid == len(tFreq) {
			if right-left < minLen {
				start = left
				minLen = right - left
			}
			char = rune(s[left])
			left++
			if tFreq[char] > 0 {
				if windowFreq[char] == tFreq[char] {
					valid--
				}
				windowFreq[char]--
			}
		}
	}
	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func sortIntervals(intervals [][]int) {
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][1] > intervals[j][1] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
}
func main() {
	// Test Ft_coin
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // Résultat : 3
	fmt.Println(Ft_coin([]int{2}, 3))        // Résultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // Résultat : 0
	// Test Ft_missing
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // Résultat : 0
	fmt.Println(Ft_missing([]int{0, 1}))                      // Résultat : 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // Résultat : 8
	// Test Ft_non_overlap
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // Résultat : 1
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // Résultat : 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // Résultat : 2
	// Test Ft_profit
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4})) // Résultat : 5
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))    // Résultat : 0
	// Test Ft_max_substring
	fmt.Println(Ft_max_substring("abcabcbb")) // Résultat : 3
	fmt.Println(Ft_max_substring("bbbbb"))    // Résultat : 1
	// Test Ft_min_window
	fmt.Println(Ft_min_window("ADOBECODEBANC", "ABC")) // Résultat : "BANC"
	fmt.Println(Ft_min_window("a", "aa"))              // Résultat : ""
}
