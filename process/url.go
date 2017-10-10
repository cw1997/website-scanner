package process

import (
	"log"
	"regexp"
	"strings"

	"net/url"

	"github.com/cw1997/website-scanner/util"
)

func GetHost(urlStr string) string {
	u, err := url.Parse(FormatUrl(urlStr))
	if err != nil {
		log.Fatalln("解析URL出错", err)
	}
	return u.Host
}

func FormatUrl(rawUrl string) (url string) {
	url = strings.ToLower(rawUrl)
	url = strings.TrimRight(url, "/")
	if !strings.HasPrefix(url, "http:") {
		url = "http://" + url
	}
	//println(urlStr)
	lastSlash := strings.LastIndex(url, "/")
	if lastSlash > len("http://") {
		lastPath := url[lastSlash:]
		if strings.Contains(lastPath, ".") {
			url = url[:lastSlash]
		}
	}
	url = url + "/"
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
		lines := util.ReadFile(v)
		for _, line := range lines {
			line = strings.Trim(line, "/")
			urls = append(urls, line)
		}
	}
	return urls
}

// 进行URL合并和追加处理
func AppendUrl(urlQueue chan string, appendUrl string, urlList []string) {
	var mergedUrl string
	for _, v := range urlList {
		mergedUrl = mergeUrl(appendUrl, v)
		urlQueue <- mergedUrl
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
	//url = url1 + url2
	//log.Println("merge", url1, "\t", url2, "\t", url)
	return url
}
