package cleaner

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func IsDirectory(name string) (isDir bool, err error) {
	fInfo, err := os.Stat(name) // FileInfo型が返る。
	if err != nil {
		return false, err
	}
	return fInfo.IsDir(), nil
}

func Parse() (string, bool, int) {
	var (
		pattern string
		delete  bool
		limit   int
	)

	flag.StringVar(&pattern, "f", "", "search file pattern. e.g. /your/path/*.log")
	flag.BoolVar(&delete, "d", false, "force delete")
	flag.IntVar(&limit, "l", 0, "limited day. default 1week(7days)")
	flag.Parse()

	if pattern == "" {
		fmt.Printf("\x1b[31m[Error]\x1b[0m Search file pattern is empty.\n\n")
		fmt.Printf("[Usage]\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	// 期限日が指定されていなければ一週間を設定
	if limit == 0 {
		limit = 7
	}

	return pattern, delete, limit
}

func SearchFiles(dirName string, pattern string, limit int) []os.FileInfo {
	// 現在時刻
	now := time.Now()

	// 保護する期間
	limitDay := now.AddDate(0, 0, -(limit))

	// ディレクトリチェック
	var isDir, err = IsDirectory(dirName + pattern)
	if isDir == true {
		dirName = dirName + pattern
		pattern = ""
	}

	// ディレクトリ読み込み
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Printf("Directory can not read \n")
		os.Exit(1)
	}

	// ファイルを検索
	matchedFiles := make([]os.FileInfo, len(fileInfos))
	for _, fileInfo := range fileInfos {
		var findName = (fileInfo).Name()
		var matched = true
		if pattern != "" {
			matched, _ = filepath.Match(pattern, findName)
		}
		if matched == true {
			if limitDay.After((fileInfo).ModTime()) {
				fmt.Println(findName)
				_ = append(matchedFiles, fileInfo)
			}
		}
	}

	return matchedFiles
}

func DeleteFiles(dirName string, fileInfos []os.FileInfo) error {
	for _, fileInfo := range fileInfos {
		var findName = (fileInfo).Name()
		if err := os.Remove(dirName + findName); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Delteted: %s [modified: %s]\n", findName, (fileInfo).ModTime().Format("2006-01-02 15:04:05"))
		}
	}

	return nil
}
