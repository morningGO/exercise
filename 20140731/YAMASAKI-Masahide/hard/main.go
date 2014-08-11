package main

import (
	"net"

	"github.com/kr/pretty"
)

func GetInterfaces() *map[string]net.IP {
	result := map[string]net.IP{}
	/* Hard:自分のマシンのIPアドレスを取得し画面表示 */
	interfaces, _ := net.Interfaces()
	//fmt.Printf("%#v\n", interfaces)
	for _, i := range interfaces {
		ips, _ := i.Addrs()
		for _, ip := range ips {
			//pretty.Printf("%# v\n", ip)
			switch ip := ip.(type) {
			default:
				pretty.Printf("unexpected type %t\n", ip)
			case *net.IPNet:
				result[i.Name] = ip.IP
			case *net.IPAddr:
				result[i.Name] = ip.IP
				/*
					//↓の書き方はコンパイルエラーに
					case *net.IPNet, *net.IPAddr:
						printIP(i.Name, &ip.IP)  // > ip.IP undefined (type net.Addr has no field or method IP)
				*/
			}
		}
	}
	return &result
}

func main() {
	for key, ip := range *GetInterfaces() {
		if !ip.IsLoopback() {
			pretty.Printf("%s:%s\n", key, ip)
		}
	}
}
