package process

import (
	"log"
	"regexp"
	"strings"

	"util"
)

func FormatUrl(rawUrl string) (url string) {
	url = strings.ToLower(rawUrl)
	if !strings.HasPrefix(url, "http:") {
		url = "http://" + url
	}
	//println(urlStr)
	return url
}

func CheckUrl(urlStr string) bool {
	reg := `^((ht|f)tps?):\/\/[\w\-]+(\.[\w\-]+)+([\w\-\.,@?^=%&:\/~\+#]*[\w\-\@?^=%&\/~\+#])?$`
	match, err := regexp.MatchString(reg, urlStr)
	if err != nil {
		log.Fatalln(err)
	}
	return match
}

func CacheUrl(dirPath string, suffix string) (urls []string) {
	urls = make([]string, 0, 128)
	files, err := util.WalkDir(dirPath, suffix)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(files)

	for _, v := range files {
		//urls = util.ReadFile(v)
		urls = append(urls, util.ReadFile(v)...) // slice merge ...表示切片解构
	}
	return urls
}

// 进行URL合并和追加处理
func AppendUrl(urlQueue chan string, appendUrl string, urlList []string) {
	for _, v := range urlList {
		urlQueue <- mergeUrl(appendUrl, v)
	}
}

func mergeUrl(url1 string, url2 string) (url string) {
	if (strings.HasSuffix(url1, "/") && !strings.HasPrefix(url2, "/")) || (!strings.HasSuffix(url1, "/") && strings.HasPrefix(url2, "/")) {
		url = url1 + url2
	}
	if !strings.HasSuffix(url1, "/") && !strings.HasPrefix(url2, "/") {
		url = url1 + "/" + url2
	}
	if strings.HasSuffix(url1, "/") && strings.HasPrefix(url2, "/") {
		url = strings.TrimRight(url1, "/") + url2
	}
	return url
}
