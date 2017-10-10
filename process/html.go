package process

import (
	"fmt"

	"log"

	"time"

	"github.com/cw1997/website-scanner/util"
)

func OutputHtml(urlStr string, result map[string]string) {
	var html string
	html += fmt.Sprintf(`<h1> %s 的扫描结果</h1>`, urlStr)
	html += "<ol>"
	for url, status := range result {
		html += fmt.Sprintf(`<li><a href="%s">%s</a><span style="color: gray"> %s</span></li>`, url, url, status)
	}
	html += "</ol>"
	html += fmt.Sprintf(`<p>总共 <strong>%d</strong> 个扫描结果</p>`, len(result))
	html += fmt.Sprintf(`<p>生成时间 <strong>%s</strong> </p>`, time.Now().Format("2006-01-02 15:04:05"))
	html += `Powered by website-scanner <a href="https://github.com/cw1997/website-scanner">https://github.com/cw1997/website-scanner</a>`
	html += `<hr>`
	filename := GetHost(urlStr) + ".html"
	util.WriteFile(filename, html)
	log.Println("扫描结果已输出至:" + filename)
}
