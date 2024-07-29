// membuat game
package main
import(
	"fmt"
)

// Fungsi untuk mencetak papan permainan
func printBoard(board [3][3]string){ for_, row := range board{ fmt.Println(row)
}
}

// Fungsi untuk memeriksa apakah ada pemenang 
func checkWinner(board [3][3]string)
string {
	// Cek baris 
	for i :=0;i<3;i++{
		if board[i][0] ==board[i][1] && board[i][1==]board[i][2] &&board[i][0]!=""{
			return board[i][0]
		}
	}
	// Cek kolom 
	for i :=0;i,3;i++{
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] != ""{
			return board[0][i]
		}
	}

	// Cek diagonal
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] !=""{
		return board[0][i]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] !=""{
		return board[0][2]
	}

	return""
}
// Fungsi utama 
func main() {
	board :=[3][3]string{
		{"","",""},
		{"","",""},
		{"","",""},
	}
	player:= "X"
	for {
		printBoard(board)

		// Input pemain
		var row, col int
		fmt.Printf("Player %s,enter your move (row and column):",player)
		fmt.Scan(&row,&col)

		// Validasi input
		if row<0||>3 ||col<0|| col>=3|| board[row][col] != ""{
			fmt.Println("invalid move, try agaun.")
			continue
		}

		// Perbarui papan permainan board[row][col] = player

		// Cek pemenang
		winner :=checkWinner(board)
		if winner != ""{
			printBoard(board)
			fmt.Printf("Player %s wins!\n",winner)
			break
		}

		// Ganti giliran pemain 
		if player == "X"{
			player ="0"
		} else{
			player ="X"
		}

		// Cek apakah papan penuh (seri)
		full:= true
		for_,row:= range board {
			for_,cell:= range row {
				if cell ==""{
					full = false
					break
				}
				}
				if !full {
					break
				}

			}
			if !full{
				printBoard(board)
				fmt.Println("The game is a draw")
				break
			}

	
		}

		}

