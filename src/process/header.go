package process

import (
	"strings"
	"util"
)

func CacheHeader(filename string) (headers map[string]string) {
	headerText := util.ReadFile(filename)
	headers = make(map[string]string)
	//println(len(headerText))
	for _, v := range headerText {
		//println(v)
		header := strings.SplitN(v, ":", 2)
		if len(header) < 2 {
			//log.Println("err:", v)
		} else {
			//log.Println(" k:", header[0], " v:", header[1])
			headers[header[0]] = header[1]
		}
	}
	return headers
}

/*func checkHeader()  {

}*/
