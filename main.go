package main

import (
	"./cleaner"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var searchFile, forceDelete, limitDay = cleaner.Parse()

	var dirName, filePattern = filepath.Split(searchFile)
	if dirName == "" {
		var cDir, _ = os.Getwd()
		dirName = cDir + "/"
	}

	// ファイル検索
	matchedFileInfos := cleaner.SearchFiles(dirName, filePattern, limitDay)
	if len(matchedFileInfos) == 0 {
		fmt.Printf("File patter is not match. [%s]\n", dirName+filePattern)
		os.Exit(0)
	}

	var timeFormat = "2006-01-02 15:04:05"
	if forceDelete == false {
		// 結果出力
		fmt.Printf("now: %s\n", time.Now().Format(timeFormat))
		fmt.Printf("limit: %s\n\n", time.Now().AddDate(0, 0, -(limitDay)).Format(timeFormat))
		fmt.Printf("Delete Files: \n")
		for _, fileInfo := range matchedFileInfos {
			var findName = (fileInfo).Name()
			fmt.Printf("%s [modified: %s]\n", findName, (fileInfo).ModTime().Format(timeFormat))
		}
	} else {
		var err = cleaner.DeleteFiles(dirName, matchedFileInfos)
		if err != nil {
			fmt.Println("something error.")
			os.Exit(1)
		} else {
			fmt.Println("complete.")
		}
	}

}
