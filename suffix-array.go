package datastructures

import (
	"sort"
)

// SuffixArray represents a suffix array data structure.
type SuffixArray struct {
	text                 string
	suffixArr            []int
	hasComputedSuffixArr bool
	lcpArr               []int
	hasComputedLCPArr    bool
	size                 int
}

// NewSuffixArray returns a new suffix array data structure.
func NewSuffixArray(text string) *SuffixArray {
	sa := &SuffixArray{
		suffixArr: []int{},
		text:      text,
		size:      len(text),
	}
	return sa
}

// GetSuffixArray returns the suffix array.
func (sa *SuffixArray) GetSuffixArray() []int {
	sa.buildSuffixArray()
	return sa.suffixArr
}

// GetSize returns the size of the suffix array.
func (sa *SuffixArray) GetSize() int {
	return sa.size
}

// buildSuffixArray is a helper method to build suffix array.
func (sa *SuffixArray) buildSuffixArray() {
	if sa.hasComputedSuffixArr {
		return
	}
	suffixes := []string{}
	suffixIndexMap := map[string]int{}
	for i := 0; i < len(sa.text); i++ {
		suffixes = append(suffixes, sa.text[i:])
		suffixIndexMap[sa.text[i:]] = i
	}
	sort.Strings(suffixes)

	for _, suffix := range suffixes {
		sa.suffixArr = append(sa.suffixArr, suffixIndexMap[suffix])
	}
	sa.hasComputedSuffixArr = true
}

// GetLCPArray returns the Longest Common Prefix array.
func (sa *SuffixArray) GetLCPArray() []int {
	sa.buildSuffixArray()
	sa.buildLCPArray()
	return sa.lcpArr
}

// buildLCPArray is a helper method to build an lcp array from the
// suffix array.
//
// LCP stands for: Longest Common Prefix
func (sa *SuffixArray) buildLCPArray() {
	if sa.hasComputedLCPArr {
		return
	}
	lcpArr := []int{}
	lcpArr = append(lcpArr, 0)
	for k, v := range sa.suffixArr {
		k2 := k + 1
		if k2 > sa.GetSize()-1 {
			break
		}
		currentSuffix := sa.text[v:]
		nextSuffix := sa.text[sa.suffixArr[k2]:]

		l := len(currentSuffix)
		for l >= 0 {
			if len(nextSuffix) < l {
				l--
				continue
			}
			if nextSuffix[:l] == currentSuffix[:l] {
				lcpArr = append(lcpArr, l)
				break
			}
			l--
		}
	}
	sa.lcpArr = lcpArr
	sa.hasComputedLCPArr = true
}

// UniqueSubstrings returns the number of unique substrings in the text.
//
// Formula:
//     n := lenght(text)
//
func (sa *SuffixArray) UniqueSubstrings() int {
	n := len(sa.text)
	lcpSum := 0
	lcpArr := sa.GetLCPArray()
	for _, v := range lcpArr {
		lcpSum += v
	}
	uniqueSubstrings := (n * (n + 1) / 2) - lcpSum
	return uniqueSubstrings
}
