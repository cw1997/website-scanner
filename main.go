package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	//接受URL
	//校验URL
	//缓存代理
	//输出结果
	//ScanThread("HEAD")
	//CacheUrl()
	println(CheckUrl("www/s/3?a=1#b=2"))
	println(CheckUrl("www.me/s/3?a=1#b=2"))
	println(CheckUrl("changwei.me/s/3?a=1#b=2"))
	println(CheckUrl("..me/s/3?a=1#b=2"))
	println(CheckUrl("http://c.me/s/3?a=1#b=2"))
	println(CheckUrl("https://c.me/s/3?a=1#b=2"))
	println(CheckUrl("https//c.me/s/3?a=1#b=2"))
	println(CheckUrl("https:/c.me/s/3?a=1#b=2"))
}
func CheckUrl(urlStr string) bool {
	urlStr = strings.ToLower(urlStr)
	if !strings.HasPrefix(urlStr, "http:") {
		urlStr = "http://" + urlStr
	}
	println(urlStr)
	reg := `^((ht|f)tps?):\/\/[\w\-]+(\.[\w\-]+)+([\w\-\.,@?^=%&:\/~\+#]*[\w\-\@?^=%&\/~\+#])?$`
	match, err := regexp.MatchString(reg, urlStr)
	if err != nil {
		log.Fatalln(err)
	}
	return match
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

func CacheUrl() {
	var c int

	files, err := WalkDir("directories", "txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(files)

	for i := 0; i < len(files); i = i + 1 {
		f, err := os.Open(files[i])
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
			c = c + 1
		}
	}
	println(c)
}
func ScanThread(methodStr string) {
	client := &http.Client{}

	req, err := http.NewRequest(methodStr, "https://www.baidu.com", nil)
	if err != nil {
		log.Fatalln(" 程序运行出现致命错误，请检查您是否输入了不合法的URL。\n" + err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

	//返回的状态码
	status := resp.StatusCode
	fmt.Println(status)
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
}

/*func ScanThread(methodStr string) {
	client := &http.Client{}

	//生成要访问的url
	url := "http://www.baidu.com"

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)

	//返回的状态码
	status := response.StatusCode

	fmt.Println(status)
}*/
