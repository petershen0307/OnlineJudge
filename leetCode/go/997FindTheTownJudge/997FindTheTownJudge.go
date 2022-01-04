package findthetownjudge

func findJudge(n int, trust [][]int) int {
	trustedCount := make([][]int, n+1)
	trustCount := make([]int, n+1)
	for _, oneLine := range trust {
		trustedCount[oneLine[1]] = append(trustedCount[oneLine[1]], oneLine[0])
		trustCount[oneLine[0]]++
	}
	// find judge
	judge := -1
	// The town judge trusts nobody.
	// the judge trust count should be 0
	for person, count := range trustCount {
		// skip index 0
		if person == 0 {
			continue
		}
		// judge candidate
		// Everybody (except for the town judge) trusts the town judge.
		if count == 0 && len(trustedCount[person]) == n-1 {
			// find second judge
			if judge != -1 {
				judge = -1
				break
			}
			judge = person
		}
	}
	return judge
}
