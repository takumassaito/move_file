package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

/*var src = os.Args[1]
var dest = os.Args[2]*/
const dst = "C:/Users/Oniichan/Desktop/pngandsketch"
const src = "C:/Users/Oniichan/Desktop/画像まとめ/imgur 2022~"
const maked_dir = dst + "/dstDirectory"

var loop_count = 0

func main() {
	err := os.MkdirAll(maked_dir, 0777)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ディレクトリはすでに存在します")
	}
	filepath.Walk(src, list_files)
}

func list_files(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	files, err := filepath.Glob(path + "/*")

	fmt.Println(files)

	if err != nil {
		return err
	}

	rename_result := rename_files(path, maked_dir)

	loop_count++

	if rename_result {
		return nil
	}

	return err
}

func rename_files(path, dst string) bool {
	var file_splited []string
	if loop_count == 0 {
		file_splited = strings.Split(path, "/")
		file_index := len(file_splited) - 1
		// fmt.Println(dst + "/" + file_splited[file_index]) テスト用
		dir_or_file, _ := os.Stat(path)
		if !dir_or_file.IsDir() {
			os.Rename(src, dst+"/"+file_splited[file_index])
		}
		return true
	}

	file_splited = strings.Split(path, "\\")
	file_index := len(file_splited) - 1
	// fmt.Println(dst + "/" + file_splited[file_index]) テスト用
	dir_or_file, _ := os.Stat(path)
	if !dir_or_file.IsDir() {
		os.Rename(src, dst+"/"+file_splited[file_index])
	}

	return true
}
