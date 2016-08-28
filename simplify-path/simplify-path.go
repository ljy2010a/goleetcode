package leetcode
import (
	"strings"
	"fmt"
)


func simplifyPath(path string) string {

	paths := strings.Split(path, "/")
	fmt.Printf("%v : len : %v \n ", paths, len(paths))
	if len(paths) == 0 {
		return "/"
	}

	tpath := []string{}
	i := 0
	for {
		if i >= (len(paths) ) {
			break
		}

		switch paths[i] {
		case "":
			break
		case ".":
			break
		case "..":
			if len(tpath) > 0 {
				tpath = tpath[:len(tpath) - 1]
			}

		default:
			tpath = append(tpath, paths[i])

		}
		i++

	}
	return "/" + strings.Join(tpath, "/")
}