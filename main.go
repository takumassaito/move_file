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
const src = "C:/Users/Oniichan/Desktop/pngandsketch"
const dst = "C:/Users/Oniichan/Desktop/画像まとめ/imgur 2022~"
//初めのループだとファイルパスが "/"で区切られてるためループカウントでキャッチする
var loop_count = 0

func main() {
	//dstパスのディレクトリをサブディレクトリまですべて作成
	err := os.MkdirAll(dst,0777)
	if err != nil{
		fmt.Println(err)
		return
	}
	filepath.Walk(src, list_files)
}

func list_files(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	rename_result := rename_files(path)

	loop_count++

	if rename_result == nil{
		return nil
	}

	return err
}

func rename_files(path string) error {
	var file_splited []string
	if loop_count == 0 {
		file_splited = strings.Split(path, "/")
		file_index := len(file_splited) - 1
		// fmt.Println(dst + "/" + file_splited[file_index]) テスト用
		dir_or_file, _ := os.Stat(path)
		if !dir_or_file.IsDir() {
			dst_filepath := "/"+file_splited[file_index]
			os.Rename(src + dst_filepath,  dst + dst_filepath)
		}
		return nil
	}

	//ファイルpathを分割し[]stringで返す
	file_splited = strings.Split(path, "\\")
	//配列の最後のインデックス番号を取得。配列の最後に入っているファイルを取り出すための準備
	file_index := len(file_splited) - 1
	// fmt.Println(dst + "/" + file_splited[file_index]) テスト用
	//ファイルもしくはディレクトリが存在していればfileinfoを返す。なければerrorを返す
	dir_or_file, _ := os.Stat(path)
	//pathから読み取ったものがディレクトリかどうかを判断
	//ディレクトリなら次のファイルへ。ファイルならRenameで移動させる。
	if !dir_or_file.IsDir() {
		dst_filepath := "/"+file_splited[file_index]
		os.Rename(src + dst_filepath, dst + dst_filepath)
	}

	return nil
}
