// Problem: https://leetcode.com/problems/available-captures-for-rook/

func numRookCaptures(board [][]byte) int {
	kill := 0
	// find rook
	var posX, posY int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				posX, posY = i, j
				break
			}
		}
	}

	// attack in east
	for i, j := posX+1, posY; i < 8; i++ {
		if board[i][j] == 'p' {
			kill++
			break
		}
		if board[i][j] == 'B' {
			break
		}
	}

	// attack in west
	for i, j := posX-1, posY; i >= 0; i-- {
		if board[i][j] == 'p' {
			kill++
			break
		}
		if board[i][j] == 'B' {
			break
		}
	}

	// attack in south
	for i, j := posX, posY+1; j < 8; j++ {
		if board[i][j] == 'p' {
			kill++
			break
		}
		if board[i][j] == 'B' {
			break
		}
	}

	// attack in north
	for i, j := posX, posY-1; j >= 0; j-- {
		if board[i][j] == 'p' {
			kill++
			break
		}
		if board[i][j] == 'B' {
			break
		}
	}

	return kill
}