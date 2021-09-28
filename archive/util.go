package archive

import "os"

func isDir(name string) bool {
	fi, err := os.Stat(name)

	return (err == nil || os.IsExist(err)) && fi.IsDir()
}
