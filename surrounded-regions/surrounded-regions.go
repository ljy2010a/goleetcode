package leetcode
import (
	"fmt"
)

func solve(board [][]byte) {
	h := len(board)
	if h == 0 {
		return
	}
	w := len(board[0])
	xb := byte('X')
	ob := byte('O')

	dx, dy := [4]int{1, -1, 0, 0}, [4]int{0, 0, 1, -1}

	var append_check = func(ps [][2]int, nx, ny int) bool {
		exist := false
		for _, p := range ps {
			if p[0] == ny && p[1] == nx {
				exist = true
			}
		}
		return exist
	}

	o_arr := [][2]int{}
	o_index := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if board[y][x] != ob {
				continue
			}

			if append_check(o_arr, y, x) {
				continue
			}

			change_arr := [][2]int{}

			o_arr = append(o_arr, [2]int{y, x})

			for {

				if o_index >= len(o_arr) {
					break
				}

				cury, curx := o_arr[o_index][0], o_arr[o_index][1]

				fmt.Printf("pos %v,%v \n", curx, cury)

				edge_flag := true
				for z := 0; z < 4; z++ {
					nx := curx + dx[z]
					ny := cury + dy[z]
					fmt.Printf("nx : %v  ny : %v \n", nx, ny)
					if nx >= h || ny >= w || nx < 0 || ny < 0 {
						fmt.Printf("to the edge \n")
						edge_flag = false
						continue
					}

					if board[ny][nx] == ob {
						if !append_check(o_arr, ny, nx) {
							o_arr = append(o_arr, [2]int{ny, nx})
						}
					}


				}

				if edge_flag {
					for _, p := range change_arr {
						x := p[0]
						y := p[1]
						board[x][y] = xb
					}
				}
				o_index ++
			}
		}
	}
}
