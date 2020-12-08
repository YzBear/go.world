package util

import (
	"strconv"
)

//ScanMemcacheServer 扫描指定ip 段的memcached server
func ScanMemcacheServer() {
	//初始化文件配置
	Load()
	// for i := 197; i < 255; i++ {
	// 	for j := 1; j < 255; j++ {
	// 		server := "154.203." + strconv.Itoa(i) + "." + strconv.Itoa(j) + ":11211"
	// 		ConnectAndSave(server)
	// 	}
	// 	time.Sleep(time.Millisecond * 100)
	// }
	for j := 1; j < 255; j++ {
		server := "154.203.197." + strconv.Itoa(j) + ":11211"
		ConnectAndSave(server)
	}
	//最后关闭文件
	defer Close()
}
