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

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

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

func WriteFile(filename string, writeString string) (size int) {
	var err error
	var f *os.File
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		if err != nil {
			log.Fatalln(filename+" 文件已经存在", err.Error())
		}
	} else {
		f, err = os.Create(filename) //创建文件
		if err != nil {
			log.Fatalln(filename+" 创建文件失败", err.Error())
		}
	}
	n, err := io.WriteString(f, writeString) //写入文件(字符串)
	if err != nil {
		log.Fatalln(filename+" 写入文件失败", err.Error())
	}
	return n
}
