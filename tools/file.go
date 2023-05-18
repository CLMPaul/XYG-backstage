package tools

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func LocalPath2Public(paths []string, host string) []string {
	var ret []string
	for _, path_ := range paths {
		sub := strings.Split(path_, "/")
		ret = append(ret, "https://"+path.Join("xueyigou.cn:30001", "image", sub[1], sub[2]))
	}
	fmt.Println(ret)
	return ret
}

func PublicPath2Local(paths []string) []string {
	var ret []string
	if len(paths) == 0 {
		return nil
	}
	for _, path_ := range paths {
		name := path.Base(path_)
		ret = append(ret, path.Join("./public/pictures", name))
	}
	return ret
}
