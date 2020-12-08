package util

import (
	"DDOS/entity"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// File memcached server txt，default location in root path
var File os.File

//时间必须是golang诞生的时间
const (
	filePath     = "logs"
	fileNameLast = ".txt"
	baseFormat   = "20060102"
)

//Load 初始化文件配置
func Load() {
	checkPath(filePath)
	fileName := filePath + string(os.PathSeparator) + time.Now().Format(baseFormat) + fileNameLast
	fileTmp, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		fmt.Println(err)
		return
	}
	File = *fileTmp
	fmt.Println("memcached server文件初始化完成")
}

//checkPath 检查文件夹是否存在，不存在就创建
func checkPath(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(path, os.ModePerm)
		}
	}
}

// AddServer 给日志里面增加一行Server xxx.xxx.xxx.xxx:xxxxx
func AddServer(server string) {
	File.WriteString(server + "\n")
	fmt.Printf("%v--success\n", server)
}

// Close 手动关闭文件
func Close() {
	File.Close()
}

// ReadServer 从文件里面读数据
func ReadServer() []entity.IPPort {
	resultArr := []entity.IPPort{}
	fileName := filePath + string(os.PathSeparator) + time.Now().Format(baseFormat) + fileNameLast
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Printf("ReadServer error:%v\n", err)
		return resultArr
	}
	//从文件一行一行读数据，放在实体类中，放在数组中
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		//解析
		ipPortArr := strings.Split(line, ":")
		if len(ipPortArr) == 2 {
			targetPort, _ := strconv.Atoi(removeBlankAndR(ipPortArr[1]))
			resultArr = append(resultArr, entity.IPPort{
				IP:   removeBlankAndR(ipPortArr[0]),
				Port: targetPort,
			})
		}
	}
	return resultArr
}

func removeBlankAndR(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}
