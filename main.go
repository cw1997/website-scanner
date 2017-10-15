package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/cw1997/website-scanner/process"
	"github.com/cw1997/website-scanner/util"
)

func main() {
	urlStr := flag.String("url", "127.0.0.1/scantest/conn.asp", "please input a url")
	method := flag.String("method", "HEAD", "please input a method")
	headers := flag.String("header", "header.txt", "please input a header filename")
	threadNum := flag.Int("thread", 10, "please input thread number")
	directoriesPath := flag.String("path", "字典", "please input directories path")
	suffix := flag.String("extname", "txt", "please input directories file extname")

	startTime := time.Now().Unix()
	scan(*urlStr, *method, *headers, *threadNum, *directoriesPath, *suffix)
	endTime := time.Now().Unix()
	fmt.Println("spend time:", endTime-startTime, "s")
}

func scan(urlStr string, methodStr string, headers string, threadNum int, directoriesPath string, suffix string) {
	urlStr = process.FormatUrl(urlStr)
	log.Println("scan", urlStr)
	if !process.CheckUrl(urlStr) {
		log.Fatalln("check url fail")
	}

	log.Println("threadNum:", threadNum)

	urlList := process.CacheUrl(directoriesPath, suffix)
	log.Println("urlList(before remove duplicate1)", len(urlList))
	urlList = util.RemoveDuplicate1(urlList)
	log.Println("urlList(after remove duplicate1)", len(urlList))

	headerMap := process.CacheHeader(headers)
	log.Println("headerList", len(headerMap))

	methodStr = strings.ToUpper(methodStr)

	urlQueue := make(chan string)
	resultQueue := make(chan string)

	result := make(map[string]string)

	go func(resultQueue chan string, result map[string]string) {
		for urlInfo := range resultQueue {
			ret := strings.SplitN(urlInfo, " ", 2) // 必须使用splitn且n=2，因为status可能含有空格
			result[ret[0]] = ret[1]
		}
	}(resultQueue, result)

	for i := 0; i < threadNum; i++ {
		go scanThread(urlQueue, resultQueue, methodStr, headerMap, urlList)
	}

	process.AppendUrl(urlQueue, urlStr, urlList)

	process.OutputHtml(urlStr, result)

}

//扫描thread
func scanThread(urlQueue chan string, resultQueue chan string, methodStr string, headers map[string]string, urlList []string) {
	// get url from urlQueue channel
	for urlStr := range urlQueue {
		//log.Println(urlStr)

		client := &http.Client{}

		req, err := http.NewRequest(methodStr, urlStr, nil)
		if err != nil {
			log.Println(" 扫描URL： " + urlStr + " 出错，可能为Bad Request。\n" + err.Error())
			continue
		}

		for k, v := range headers {
			req.Header.Set(k, v)
		}
		//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}

		//defer resp.Body.Close()

		/*body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(body))*/
		//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
		/*stdout := os.Stdout
		_, err = io.Copy(stdout, response.Body)*/

		//返回的状态码
		StatusCode := resp.StatusCode
		status := resp.Status
		result := urlStr + " " + status
		log.Println(result)
		if StatusCode != 404 {
			resultQueue <- result
		}

		//TODO: 这里本来要实现扫描到最后一个path为文件夹的情况下继续append到channel里等待下一轮扫描，但是运行时出现漏扫和死锁现象。
		/*u, err := url.Parse(urlStr)
		//log.Println(u.Path, strings.Contains(u.Path, "."))
		if err != nil {
			log.Println("解析URL: " + urlStr + " 失败")
			return
		}
		//|| strings.HasPrefix(status, "30")
		if !strings.Contains(u.Path+u.Fragment, ".") && (StatusCode == 200 || StatusCode == 403) {
			log.Println("lenqueue:", len(urlQueue), "urlList", len(urlList), urlStr)
			process.AppendUrl(urlQueue, process.FormatUrl(urlStr), urlList)
		}*/
	}
}
