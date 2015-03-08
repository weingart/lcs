package lcs

func LCS(a, b []rune) (int, []rune) {

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

	// read the substring out from the matrix
	s := make([]rune, 0, lengths[len(a)][len(b)])
	for x, y := len(a), len(b); x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			x--
		} else if lengths[x][y] == lengths[x][y-1] {
			y--
		} else {
			s = append(s, a[x-1])
			x--
			y--
		}
	}

	// reverse string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return len(s), s
}
