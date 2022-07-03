package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ToCR 转为MacOS换行符
func ToCR(s string) string {
	d := strings.ReplaceAll(s, "\r\n", "\r")
	d = strings.ReplaceAll(d, "\n", "\r")
	return d
}

// ToLF 转为Linux换行符
func ToLF(s string) string {
	d := strings.ReplaceAll(s, "\r\n", "\n")
	d = strings.ReplaceAll(d, "\r", "\n")
	return d
}

// ToCRLF 转为Windows换行符
func ToCRLF(s string) string {
	d := ToLF(s)
	d = strings.ReplaceAll(d, "\n", "\r\n")
	return d
}

// procString 处理字符串的换行符
func procString(t, s string) (string, error) {
	d := ""
	switch t {
	case "cr":
		d = ToCR(s)
	case "lf":
		d = ToLF(s)
	case "crlf":
		d = ToCRLF(s)
	default:
		return "", errors.New("未指定换行符")
	}
	return d, nil
}

// procFile 处理文件的换行符
func procFile(t, s string) error {
	// 判断文件是否存在
	f, err := os.Stat(s)
	if err != nil {
		return err
	}
	if f.Size() > 4*1024*1024 {
		return fmt.Errorf("%s 文件过大", s)
	}
	// Todo:判断文件是否为文本文件
	bs, err := ioutil.ReadFile(s)
	if err != nil {
		return err
	}
	d, err := procString(t, string(bs))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(s, []byte(d), 0666)
	return err
}

// procDir 处理文件夹下文件的换行符
func procDir(t, s string) error {
	fs, err := ioutil.ReadDir(s)
	if err != nil {
		return err
	}
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		err = procFile(t, filepath.Join(s, f.Name()))
	}
	return err
}
