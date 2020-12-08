package util

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

// ConnectAndTestIsActive 测试连接与存活
func ConnectAndTestIsActive(server string) (canConnect bool, mc *memcache.Client) {
	mc = memcache.New(server)
	if mc == nil {
		fmt.Printf("memcache new failed:%v\n", mc)
		return false, nil
	}
	err := mc.Ping()
	if err != nil {
		fmt.Printf("memcache new failed:%v\n", err)
		return false, nil
	}
	return true, mc
}

//ConnectAndSave 连接并保存在文件中
func ConnectAndSave(server string) {
	canConnect, _ := ConnectAndTestIsActive(server)
	if canConnect {
		AddServer(server)
	}
}
