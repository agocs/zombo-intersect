package zombo

//zombofyLinear finds the intersection between any (reasonable) number of ordered string slices.
//It does this by setting the first string in the slice to the intersection, and intersecting it
//with each subsequent slice.
func zombofyLinear(input [][]string) []string {
	intersection := []string{}

	for idx, next := range input {
		if idx == 0 {
			intersection = next
		}
		intersection = intersect(intersection, next)
	}
	return intersection
}

//intersect calculates the deduplicated intersection between two ordered string slices
func intersect(a []string, b []string) []string {
	intersection := []string{}
	idxA := 0
	idxB := 0

	stopA := len(a)
	stopB := len(b)

	for {
		if idxA == stopA || idxB == stopB {
			return intersection
		}
		if a[idxA] == b[idxB] {
			candidate := a[idxA]
			if tail(intersection) != candidate {
				intersection = append(intersection, candidate)
			}
			idxA++
			idxB++
			continue
		} else if a[idxA] > b[idxB] {
			idxB++
		} else {
			idxA++
		}
	}
}

//tail returns the last string in a string slice, or emptystring if the slice is empty.
func tail(a []string) string {
	if len(a) == 0 {
		return ""
	}
	return a[len(a)-1]
}
