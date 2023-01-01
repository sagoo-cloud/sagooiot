package utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/russross/blackfriday/v2"
	"io"
	"net"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

// MarkdownToHtml 解析markdown为html
func MarkdownToHtml(mdContent string) string {
	return string(blackfriday.Run([]byte(mdContent)))
}

// EncryptPassword 密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

// GetDomain 获取当前请求接口域名
func GetDomain(ctx context.Context) string {
	r := g.RequestFromCtx(ctx)
	pathInfo, err := gurl.ParseURL(r.GetUrl(), -1)
	if err != nil {
		g.Log().Error(ctx, err)
		return ""
	}
	return fmt.Sprintf("%s://%s:%s/", pathInfo["scheme"], pathInfo["host"], pathInfo["port"])
}

// GetClientIp 获取客户端IP
func GetClientIp(ctx context.Context) string {
	return g.RequestFromCtx(ctx).GetClientIp()
}

// GetUserAgent 获取user-agent
func GetUserAgent(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).Header.Get("User-Agent")
}

// GetLocalIP 服务端ip
func GetLocalIP() (ip string, err error) {
	var addrs []net.Addr
	addrs, err = net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// GetCityByIp 获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := g.Client().GetBytes(context.TODO(), url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.Get("code").Int() == 0 {
		city := fmt.Sprintf("%s %s", json.Get("pro").String(), json.Get("city").String())
		return city
	} else {
		return ""
	}
}

// 写入文件
func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	n, _ := f.Seek(0, io.SeekEnd)
	_, err = f.WriteAt([]byte(content), n)
	defer f.Close()
	return err
}

// 文件或文件夹是否存在
func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

// 解析路径获取文件名称及后缀
func ParseFilePath(pathStr string) (fileName string, fileType string) {
	fileNameWithSuffix := path.Base(pathStr)
	fileType = path.Ext(fileNameWithSuffix)
	fileName = strings.TrimSuffix(fileNameWithSuffix, fileType)
	return
}

func RemoveRepeatedElementAndEmpty(arr []int) []int {
	newArr := make([]int, 0)
	for _, item := range arr {
		repeat := false
		if len(newArr) > 0 {
			for _, v := range newArr {
				if v == item {
					repeat = true
					break
				}
			}
		}
		if repeat {
			continue
		}
		newArr = append(newArr, item)
	}
	return newArr
}

// RemoveDuplicationMap 数组去重
func RemoveDuplicationMap(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	return arr[:j]
}

// Decimal 保留两位小数
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// InArray 判断字符串是否存在数组中
func InArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}
