package fmtPrt

import (
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
)

func Test01(t *testing.T) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}

		}
	}
}

func GetOutBoundIP()(ip string, err error)  {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func Test02(t *testing.T) {
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}

