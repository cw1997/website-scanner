package util

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

// 只能读取结尾带空行的文件，否则会漏读最后一行数据
func ReadFileAllLine(file string) (content []string) {
	f, err := os.Open(file)
	if err != nil {
		//panic(err)
		log.Fatalln(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
		content = append(content, line)
	}
	return content
}

func ReadFile(file string) (content []string) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(string(b))
	// check "\r\n" or "\n"
	content = strings.Split(string(b), "\r\n")
	if len(content) > 0 {
		return content
	} else {
		content = strings.Split(string(b), "\n")
		return content
	}
}
