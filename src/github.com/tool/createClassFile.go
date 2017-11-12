package main

import (
	"os"
	"bufio"
	"strings"
)

func main() {
	copyName := os.Args[1]
	pkgName := os.Args[2]
	createFile := os.Args[3]

	// 新規ファイルを作成
	newFile, err := os.Create(pkgName + "/" + createFile)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	// コピー元ファイルの中身を読込
	cpFile, err := os.Open(copyName)
	if err != nil {
		panic(err)
	}
	defer cpFile.Close()

	// コピー元ファイルを一行ずつ読込
	lines := make([]byte, 0)
	className := strings.Replace(createFile, ".java", "", 1)
	scanner := bufio.NewScanner(cpFile)
	for scanner.Scan() {
		line := scanner.Text()

		// クラス名とパッケージを置換
		if strings.Contains(line, "sample") {
			line = strings.Replace(line, "sample", className, 1) + "\n"
		} else {
			line = line + "\n"
		}
		lines = append(lines, []byte(line)...)
	}
	newFile.Write(lines)
}

