package main

import (
    "fmt"
    "net"
)

// 获取本机网卡IP
func getLocalIP() (ipv4 string, err error) {
    // var (
    // 	addrs []net.Addr
    // 	addr net.Addr
    // 	ipNet *net.IPNet // IP地址
    // 	isIpNet bool
    // )
    // 获取所有网卡
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return
    }
    // 取第一个非lo的网卡IP
    for _, addr := range addrs {
        // 这个网络地址是IP地址: ipv4, ipv6
        ipNet, isIpNet := addr.(*net.IPNet)
        if isIpNet && !ipNet.IP.IsLoopback() {
            // 跳过IPV6
            if ipNet.IP.To4() != nil {
                ipv4 = ipNet.IP.String() // 192.168.1.1
                return
            }
        }
    }

    return
}

func main() {
    ipv4, _ := getLocalIP()
    fmt.Printf("ipv4 is <%s>", ipv4)
}

