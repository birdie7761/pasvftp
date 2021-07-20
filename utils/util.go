package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func JoinHostPort(host string, port uint) string {
	return net.JoinHostPort(host, fmt.Sprintf("%d", port))
}

var publicIp string

func GetLocalIP() (ipv4 string, err error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}

	err = errors.New("ERR_NO_LOCAL_IP_FOUND")
	return
}

func PublicIp() (ip string, err error) {
	if os.Getenv("ENV") == "dev" {
		return "127.0.0.1", nil
	}

	if len(publicIp) > 0 {
		return publicIp, nil
	}

	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("resp status is err:" + resp.Status)
	}

	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	IP := net.ParseIP(string(buff))
	if IP == nil {
		return "", errors.New("ip is invalid:" + string(buff))
	}
	ip = IP.String()
	return
}
