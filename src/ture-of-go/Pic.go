package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	mp := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		mp[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			// 青いグラデーションの画像
			mp[i][j] = uint8((i + j) / 2)
		}
	}
	return mp
}

func main() {
	pic.Show(Pic)
}
