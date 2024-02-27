package utils

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// IpInBlackListRange 判断IP是否在黑名单
func IpInBlackListRange(ip string, ipList []string) (result bool) {
	result = false
	cidr := strings.Join(ipList[:], ",")
	cidrs, err := ParseIpRange(cidr)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	return isInRangeList(ip, cidrs)
}

// isInRange 判断IP是否在指定的范围
// 支持单个IP，支持多个IP，多IP时需要用“,”隔开
// 支持IP段，如192.168.0.1/24
// 支持IP范围，格式如：192.168.1.xx-192.168.1.xx
func isInRange(ip, cidr string) bool {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	return ipnet.Contains(net.ParseIP(ip))
}

// isInRangeList 判断一个 IP 地址是否在多个 CIDR 范围内
func isInRangeList(ip string, cidrs []string) bool {
	for _, cidr := range cidrs {
		if isInRange(ip, cidr) {
			return true
		}
	}
	return false
}

// ParseIpRange 用于将以 "," 分割的多个 CIDR 范围字符串解析为 CIDR 数组。
// 如果一个 CIDR 范围是以 "-" 分割的两个 IP 地址，那么我们会使用 binaryToInt 和 intToIP 函数将它们转换为整数并再次转换为 CIDR 字符串，
// 并将其加入到 CIDR 数组中。
func ParseIpRange(cidr string) ([]string, error) {
	var cidrs []string

	parts := strings.Split(cidr, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if strings.Contains(part, "-") {
			ips := strings.Split(part, "-")
			startIP := net.ParseIP(ips[0]).To4()
			endIP := net.ParseIP(ips[1]).To4()
			if startIP == nil || endIP == nil {
				return nil, fmt.Errorf("无效IP范围: %s", part)
			}
			start := binaryToInt(startIP)
			end := binaryToInt(endIP)
			for i := start; i <= end; i++ {
				startIP := net.ParseIP(intToIP(i).String())
				if startIP == nil {
					return nil, fmt.Errorf("无效IP地址: %s", intToIP(i).String())
				}

				cidrs = append(cidrs, intToIP(i).String()+"/32")
			}
		} else {
			startIP := net.ParseIP(part)
			if startIP == nil {
				return nil, fmt.Errorf("无效IP地址: %s", part)
			}
			cidrs = append(cidrs, part)
		}
	}
	return cidrs, nil
}

func binaryToInt(ip net.IP) int {
	return int(ip[0])<<24 | int(ip[1])<<16 | int(ip[2])<<8 | int(ip[3])
}

func intToIP(n int) net.IP {
	b := make([]byte, 4)
	b[0] = byte(n >> 24)
	b[1] = byte(n >> 16)
	b[2] = byte(n >> 8)
	b[3] = byte(n)
	return net.IPv4(b[0], b[1], b[2], b[3])
}
