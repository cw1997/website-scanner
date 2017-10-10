package util

import (
	"log"
)

//print slice
func Dump(param []string) {
	for k, v := range param {
		log.Println("k:", k, "v", v)
	}
}

func RemoveDuplicate1(arr []string) (resArr []string) {
	//resArr := make([]string, 0)
	tmpMap := make(map[string]interface{})
	for _, val := range arr {
		if _, ok := tmpMap[val]; !ok {
			resArr = append(resArr, val)
			tmpMap[val] = struct{}{}
		}
	}
	return resArr
}

/*func RemoveDuplicate(list []string, ret chan []string) {
	var x []string = []string{}
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	//return x
	ret <- x
}*/

/*func RemoveDuplicateMultiThread(list []string) (ret []string) {
	listQueue := make(chan []string)
	var listList [4][]string
	listLen := len(list)
	sliceLen := int(listLen / 4)
	lastSliceLen := listLen % 4
	var start, end int
	for i := 0; i < 4-1; i++ {
		start = i * sliceLen
		end = (i + 1) * sliceLen
		listList[i] = list[start:end]
	}
	listList[4-1] = list[:lastSliceLen]
	for i := 0; i < 4-1; i++ {
		go RemoveDuplicate(listList[i], listQueue)
	}
	ret = <-listQueue
	ret = append(ret, ret...)
	return ret
}*/
