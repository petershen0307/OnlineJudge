func numJewelsInStones(J string, S string) int {
    stoneTypeCount := make(map[rune] int)
    for _, runeA := range S {
        stoneTypeCount[runeA]++
    }
    jewels := 0
    for _, runeA := range J {
        jewels += stoneTypeCount[runeA]
    }
    return jewels
}
