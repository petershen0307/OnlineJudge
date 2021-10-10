package wordsearch

func exist(board [][]byte, word string) bool {
	for l1, l1v := range board {
		for l2 := range l1v {
			if dfs(board, word, l1, l2) == true {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, l1, l2 int) bool {
	if len(word) <= 0 {
		return true
	}
	if l1 < 0 || l1 >= len(board) || l2 < 0 || l2 >= len(board[0]) || board[l1][l2] != word[0] {
		return false
	}
	char := word[0]
	board[l1][l2] = byte('*')
	result := dfs(board, word[1:], l1-1, l2) || dfs(board, word[1:], l1+1, l2) || dfs(board, word[1:], l1, l2-1) || dfs(board, word[1:], l1, l2+1)
	board[l1][l2] = char
	return result
}
