package lcs

import (
	"testing"

	"hash/fnv"
	"math/rand"
)

var chars = []rune{
	'a', 'b', 'c',
}

var result1 string
var result2 string
var result3 string
var result4 []int
var result5 []int

func generateString(length int) string {
	s := make([]rune, 0)

	for i := 0; i < length; i++ {
		c := chars[rand.Intn(len(chars))]
		s = append(s, c)
	}

	return string(s)
}

func LCS1(a, b string) (int, string) {
	arunes := []rune(a)
	brunes := []rune(b)
	aLen := len(arunes)
	bLen := len(brunes)
	lengths := make([][]int, aLen+1)
	for i := 0; i <= aLen; i++ {
		lengths[i] = make([]int, bLen+1)
	}

	// row 0 and column 0 are initialized to 0 already
	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if arunes[i] == brunes[j] {
				lengths[i+1][j+1] = lengths[i][j] + 1
			} else if lengths[i+1][j] > lengths[i][j+1] {
				lengths[i+1][j+1] = lengths[i+1][j]
			} else {
				lengths[i+1][j+1] = lengths[i][j+1]
			}
		}
	}

	// read the substring out from the matrix
	s := make([]rune, 0, lengths[aLen][bLen])
	for x, y := aLen, bLen; x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			x--
		} else if lengths[x][y] == lengths[x][y-1] {
			y--
		} else {
			s = append(s, arunes[x-1])
			x--
			y--
		}
	}
	// reverse string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return len(s), string(s)
}

func Max(more ...int) int {
	maxNum := more[0]
	for _, elem := range more {
		if maxNum < elem {
			maxNum = elem
		}
	}
	return maxNum
}

func LCS2(str1, str2 string) (int, string) {
	len1 := len(str1)
	len2 := len(str2)

	table := make([][]int, len1+1)
	for i := range table {
		table[i] = make([]int, len2+1)
	}

	i, j := 0, 0
	for i = 0; i <= len1; i++ {
		for j = 0; j <= len2; j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				table[i][j] = Max(table[i-1][j], table[i][j-1])
			}
		}
	}
	return table[len1][len2], Back(table, str1, str2, len1-1, len2-1)
}

//http://en.wikipedia.org/wiki/Longest_common_subsequence_problem
func Back(table [][]int, str1, str2 string, i, j int) string {
	if i == 0 || j == 0 {
		return ""
	} else if str1[i] == str2[j] {
		return Back(table, str1, str2, i-1, j-1) + string(str1[i])
	} else {
		if table[i][j-1] > table[i-1][j] {
			return Back(table, str1, str2, i, j-1)
		} else {
			return Back(table, str1, str2, i-1, j)
		}
	}
}

func LCSMatrix(a, b []int) [][]int {

	lengths := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		lengths[i] = make([]int, len(b)+1)
	}

	// row 0 and column 0 are initialized to 0 already
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				lengths[i+1][j+1] = lengths[i][j] + 1
			} else if lengths[i+1][j] > lengths[i][j+1] {
				lengths[i+1][j+1] = lengths[i+1][j]
			} else {
				lengths[i+1][j+1] = lengths[i][j+1]
			}
		}
	}

	return lengths
}

type Operation int

const (
	Noop Operation = iota
	Add  Operation = iota
	Del  Operation = iota
)

type Diff struct {
	Op  Operation
	key int
}

func LCSDiff(a, b []int, lengths [][]int) []Diff {

	// Read the substring out from the matrix
	s := make([]Diff, 0, lengths[len(a)][len(b)])

	// Read out diff in reverse order
	for x, y := len(a), len(b); x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			s = append(s, Diff{Del, a[x-1]})
			x--
		} else if lengths[x][y] == lengths[x][y-1] {
			s = append(s, Diff{Add, a[y-1]})
			y--
		} else {
			s = append(s, Diff{Noop, a[x-1]})
			x--
			y--
		}
	}

	// Reverse our answer
	ReverseDiffSlice(s)

	return s
}

func ReverseDiffSlice(s []Diff) []Diff {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func ReverseIntSlice(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

type Interface interface {
	Keys() []int
}

func LCS3(a, b Interface) (int, []int) {
	aKeys := a.Keys()
	bKeys := b.Keys()

	lengths := make([][]int, len(aKeys)+1)
	for i := 0; i <= len(aKeys); i++ {
		lengths[i] = make([]int, len(bKeys)+1)
	}

	// row 0 and column 0 are initialized to 0 already
	for i := 0; i < len(aKeys); i++ {
		for j := 0; j < len(bKeys); j++ {
			if aKeys[i] == bKeys[j] {
				lengths[i+1][j+1] = lengths[i][j] + 1
			} else if lengths[i+1][j] > lengths[i][j+1] {
				lengths[i+1][j+1] = lengths[i+1][j]
			} else {
				lengths[i+1][j+1] = lengths[i][j+1]
			}
		}
	}

	// read the substring out from the matrix
	s := make([]int, 0, lengths[len(aKeys)][len(bKeys)])
	for x, y := len(aKeys), len(bKeys); x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			x--
		} else if lengths[x][y] == lengths[x][y-1] {
			y--
		} else {
			s = append(s, aKeys[x-1])
			x--
			y--
		}
	}

	ReverseIntSlice(s)

	return len(s), s
}

type RuneSlice []rune

func (p RuneSlice) Keys() []int {
	s := make([]int, 0, len(p))

	for _, r := range p {
		s = append(s, int(r))
	}
	return s
}

type StringSlice []string

func (p StringSlice) Keys() []int {
	f := fnv.New32a()

	hash := make([]int, 0, len(p))
	for i := range p {
		f.Reset()
		f.Write([]byte(p[i]))
		hash = append(hash, int(f.Sum32()))
	}
	return hash
}

func LCS3String(a, b string) (int, string) {

	_, runes := LCS3(RuneSlice(a), RuneSlice(b))
	str := make([]rune, 0, len(runes))
	for _, r := range runes {
		str = append(str, rune(r))
	}

	return len(str), string(str)
}

func BenchmarkLCS1(b *testing.B) {
	str1 := generateString(b.N)
	str2 := generateString(b.N)

	b.ResetTimer()
	_, result1 = LCS1(str1, str2)
}

func BenchmarkLCS2(b *testing.B) {
	str1 := generateString(b.N)
	str2 := generateString(b.N)

	b.ResetTimer()
	_, result2 = LCS2(str1, str2)
}

func BenchmarkLCS3(b *testing.B) {
	str1 := generateString(b.N)
	str2 := generateString(b.N)

	b.ResetTimer()
	_, result3 = LCS3String(str1, str2)
}

func BenchmarkLCS4(b *testing.B) {
	s1 := make(StringSlice, 0, b.N)
	s2 := make(StringSlice, 0, b.N)

	for i := 0; i < b.N; i++ {
		s1 = append(s1, generateString(78))
	}
	for i := 0; i < b.N; i++ {
		s2 = append(s2, generateString(78))
	}

	b.ResetTimer()
	_, result4 = LCS3(s1, s2)
}

func BenchmarkLCS5(b *testing.B) {
	s1 := make(StringSlice, 0, b.N)

	for i := 0; i < b.N; i++ {
		s1 = append(s1, generateString(78))
	}

	b.ResetTimer()
	_, result5 = LCS3(s1, s1)
}
