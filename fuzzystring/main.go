package main

import "fmt"

// StringPair consists of source and target strings
type StringPair struct {
	Source string
	Target string
}

func main() {
	strs := []StringPair{
		StringPair{"oliver", "olivia"},
		StringPair{"oliver", "olive"},
		StringPair{"oliver", "oirlve"},
		StringPair{"oliver", "revilo"},
	}

	for _, s := range strs {
		fmt.Println(s.Source, s.Target, LevenshteinDistanceRecur(s.Source, s.Target), LevenshteinDistanceOptimized(s.Source, s.Target))
	}
}

// LevenshteinDistanceOptimized has time complexity of len(s)*len(t) and needs space of len(t)
func LevenshteinDistanceOptimized(source, target string) int {
	if source == target {
		return 0
	}

	if len(source) == 0 {
		return len(target)
	}

	if len(target) == 0 {
		return len(source)
	}

	row := make([]int, len(target)+1)

	for j := 0; j <= len(target); j++ {
		row[j] = j
	}

	for i := 1; i <= len(source); i++ {
		last := i
		for j := 1; j <= len(target); j++ {
			cost := 1
			if source[i-1] == target[j-1] {
				cost = 0
			}
			temp := MinInt(row[j]+1, last+1, row[j-1]+cost)
			row[j-1] = last
			last = temp
		}
		row[len(target)] = last
	}

	return row[len(target)]
}

// LevenshteinDistance has time complexity of len(s)*len(t) and needs space of len(s)*len(t)
func LevenshteinDistance(source, target string) int {
	f := make([][]int, len(source)+1)
	for i := 0; i <= len(source); i++ {
		f[i] = make([]int, len(target)+1)
	}

	for i := 1; i <= len(source); i++ {
		f[i][0] = i
	}

	for j := 1; j <= len(target); j++ {
		f[0][j] = j
	}

	for i := 1; i <= len(source); i++ {
		for j := 1; j <= len(target); j++ {
			//fmt.Println(i, j)
			cost := 1
			if source[i-1] == target[j-1] {
				cost = 0
			}
			f[i][j] = MinInt(f[i-1][j]+1, f[i][j-1]+1, f[i-1][j-1]+cost)
		}
	}

	return f[len(source)][len(target)]
}

// LevenshteinDistanceRecur has time complexity of 3^(len(s)+len(t))
func LevenshteinDistanceRecur(source, target string) int {

	if len(source) == 0 {
		return len(target)
	}

	if len(target) == 0 {
		return len(source)
	}

	cost := 1
	if source[len(source)-1] == target[len(target)-1] {
		cost = 0
	}

	return MinInt(
		LevenshteinDistanceRecur(source[:len(source)-1], target)+1,
		LevenshteinDistanceRecur(source, target[:len(target)-1])+1,
		LevenshteinDistanceRecur(source[:len(source)-1], target[:len(target)-1])+cost)
}

// MinInt returns the minimum one among multiple integer values
func MinInt(v1 int, v ...int) int {
	min := v1
	for _, v := range v {
		if v < min {
			min = v
		}
	}

	return min
}
