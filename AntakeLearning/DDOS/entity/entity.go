package entity

import "fmt"

// IPPort 封装一下 ip:port
type IPPort struct {
	IP   string
	Port int
}

//ToString tostring
func (item *IPPort) ToString() string {
	return fmt.Sprintf("%s:%d", item.IP, item.Port)
}
