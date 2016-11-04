package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type ByName []os.FileInfo

func (fi ByName) Len() int {
	return len(fi)
}
func (fi ByName) Swap(i, j int) {
	fi[i], fi[j] = fi[j], fi[i]
}
func (fi ByName) Less(i, j int) bool {
	return fi[j].ModTime().Unix() < fi[i].ModTime().Unix()
}

func IsDirectory(name string) (isDir bool, err error) {
	fInfo, err := os.Stat(name) // FileInfo型が返る。
	if err != nil {
		return false, err
	}
	return fInfo.IsDir(), nil
}

func main() {
	var (
		searchFile  string
		forceDelete bool
		limitDay    int
	)

	flag.StringVar(&searchFile, "f", "", "search file pattern. ex) /your/path/*.log")
	flag.BoolVar(&forceDelete, "d", false, "force delete")
	flag.IntVar(&limitDay, "l", 0, "limited day. default 1week(7days)")
	flag.Parse()

	if searchFile == "" {
		fmt.Printf("\x1b[31m[Error]\x1b[0m Search file pattern is empty.\n\n")
		fmt.Printf("[Usage]\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	// 期限日が指定されていなければ一週間を設定
	if limitDay == 0 {
		limitDay = 7
	}

	var dirName, filePattern = filepath.Split(searchFile)
	if dirName == "" {
		var cDir, _ = os.Getwd()
		dirName = cDir + "/"
	}

	var isDir, _ = IsDirectory(dirName + filePattern)

	if isDir == true {
		dirName = dirName + filePattern
		filePattern = ""
	}

	fileInfos, err := ioutil.ReadDir(dirName)

	if err != nil {
		fmt.Errorf("Directory cannot read %s\n", err)
		os.Exit(1)
	}

	// 現在時刻
	now := time.Now()
	// 保護する期間
	limit := now.AddDate(0, 0, -(limitDay))

	if forceDelete == false {
		fmt.Printf("now: %s\n", now.Format("2006-01-02 15:04:05"))
		fmt.Printf("limit: %s\n\n", limit.Format("2006-01-02 15:04:05"))
		fmt.Printf("Delete Files: \n")
	}

	sort.Sort(ByName(fileInfos))
	for _, fileInfo := range fileInfos {
		var findName = (fileInfo).Name()
		var matched = true
		if filePattern != "" {
			matched, _ = filepath.Match(filePattern, findName)
		}
		if matched == true {
			if limit.After((fileInfo).ModTime()) {
				if forceDelete == true { // 削除フラグが立っていれば削除
					if err := os.Remove(dirName + findName); err != nil {
						fmt.Println(err)
					}
					fmt.Printf("Delteted: %s [modified: %s]\n", findName, (fileInfo).ModTime().Format("2006-01-02 15:04:05"))
				} else { // 表示だけ
					fmt.Printf("%s [modified: %s]\n", findName, (fileInfo).ModTime().Format("2006-01-02 15:04:05"))
				}
			}
		}
	}
}
