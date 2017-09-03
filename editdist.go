// Package editdist implements various algorithms that calculate the edit
// distance (difference) of two strings. For more information see:
// https://en.wikipedia.org/wiki/Edit_distance. Unless otherwise noted, the
// lower the score, the better and a score of zero indicates an exact match.
package editdist

import "unicode/utf8"

// minimum is a helper function used by many of the algorithms. It returns the
// smallest value of the three given.
func minimum(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

// Levenshtein calculates the edit distance using the algorithm considered by
// Vladimir Levenshtein in 1965. For more information, see:
// https://en.wikipedia.org/wiki/Levenshtein_distance.
func Levenshtein(a, b string) int {
	n, m := len(a), len(b)
	// If a string is zero-length, then the distance is the length of the other.
	if n == 0 {
		return m
	} else if m == 0 {
		return n
	}

	// If the last characters match, there is no cost.
	cost := 1
	ra, _ := utf8.DecodeLastRuneInString(a)
	rb, _ := utf8.DecodeLastRuneInString(b)
	if ra == rb {
		cost = 0
	}

	// Return the minimum of deletion, insertion, and substitution.
	return minimum(
		Levenshtein(a[:n-1], b)+1,
		Levenshtein(a, b[:m-1])+1,
		Levenshtein(a[:n-1], b[:m-1])+cost)
}

// WagnerFischer calculates the edit distance using the algorithm considered by
// Robert A. Wagner and Michael J. Fischer in 1974. For more information, see:
// https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm.
func WagnerFischer(a, b string) int {
	n, m := len(a), len(b)
	// Setup an 2 dimensional array where d[i][j] is the edit distance between
	// the first i characters of a and the first j characters of b.
	d := [][]int{}
	for x := 0; x <= n; x++ {
		d = append(d, make([]int, m+1))
	}

	// Mark default values for doing deletion and insertion for every character.
	for i := 1; i <= n; i++ {
		d[i][0] = i
	}
	for j := 1; j <= m; j++ {
		d[0][j] = j
	}

	i := 1
	for _, ra := range a {
		j := 1
		for _, rb := range b {
			cost := 1
			if ra == rb {
				cost = 0
			}
			d[i][j] = minimum(
				d[i-1][j]+1,      // deletion
				d[i][j-1]+1,      // insertion
				d[i-1][j-1]+cost, //substitution
			)
			j++
		}
		i++
	}
	return d[n][m]
}

// Hjelmqvist calculates the edit distance using the algorithm considered by
// Sten Hjelmqvist in 2012. It's similar to WagnerFischer but uses only two rows
// in the table so it's less memory intensive. For more information, see:
// https://www.codeproject.com/Articles/13525/Fast-memory-efficient-Levenshtein-algorithm.
func Hjelmqvist(a, b string) int {
	n, m := len(a), len(b)
	if n == 0 {
		return m
	} else if m == 0 {
		return n
	}

	// We only need two arrays for this algorithm because we are really just
	// updating the values as we move.
	v0 := make([]int, m+1)
	v1 := make([]int, m+1)

	// Mark default values for doing deletion and insertion for every character.
	for i := 0; i <= m; i++ {
		v0[i] = i
	}

	i := 1
	for _, ra := range a {
		j := 1
		for _, rb := range b {
			cost := 1
			if ra == rb {
				cost = 0
			}
			v1[j] = minimum(
				v1[j-1]+1,    // deletion
				v0[j]+1,      // insertion
				v0[j-1]+cost, //substitution
			)
			j++
		}
		i++
		// Move v1 to v0 so we can work on a new row.
		v0, v1 = v1, v0
	}
	return v0[m]
}
