package io

import (
	"os"
)

func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil || !os.IsNotExist(err)
}

func IsDir(d string) bool {
	fi, err := os.Stat(d)
	return err == nil && fi.IsDir()
}

func IsFile(f string) bool {
	fi, err := os.Stat(f)
	return err == nil && !fi.IsDir()
}
