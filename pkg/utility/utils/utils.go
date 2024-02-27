package utils

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

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

// GetLocalIP 获取服务器内网IP
func GetLocalIP() (string, error) {
	var localIP string
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	// 遍历所有网卡
	for _, i := range interfaces {
		if i.Flags&net.FlagUp == 0 {
			continue // 网卡未开启
		}
		if i.Flags&net.FlagLoopback != 0 {
			continue // 网卡为loopback地址
		}
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}
		// 遍历网卡上的所有地址
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // 不是ipv4地址
			}
			localIP = ip.String()
			return localIP, nil
		}
	}
	return "", fmt.Errorf("no local IP address found")
}

// GetPublicIP 获取公网IP
func GetPublicIP() (ip string, err error) {
	resp, err := http.Get("https://ifconfig.co/ip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	ip = string(body)
	// 去除空格
	ip = strings.Replace(ip, " ", "", -1)
	// 去除换行符
	ip = strings.Replace(ip, "\n", "", -1)

	return
}

// GetUserAgent 获取user-agent
func GetUserAgent(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).Header.Get("User-Agent")
}

// GetCityByIp 获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := g.Client().GetBytes(context.TODO(), url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.Get("code").Int() == 0 {
		city := ""
		if strings.EqualFold(json.Get("pro").String(), json.Get("city").String()) {
			city = fmt.Sprintf("%s", json.Get("pro").String())
		} else {
			city = fmt.Sprintf("%s %s", json.Get("pro").String(), json.Get("city").String())
		}

		return city
	} else {
		return ""
	}
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

// FileSize 字节的单位转换 保留两位小数
func FileSize(fileSize int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "EB"}
	var size = float64(fileSize)
	var i int
	for i = 0; size > 1024; i++ {
		size /= 1024
	}
	return fmt.Sprintf("%.2f %s", size, units[i])
}

type fileInfo struct {
	name string
	size int64
}

// WalkDir 获取目录下文件的名称和大小
func WalkDir(dirname string) ([]fileInfo, error) {
	var fileInfos []fileInfo
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileInfos = append(fileInfos, fileInfo{name: path, size: info.Size()})
		}
		return nil
	})

	return fileInfos, err
}

// DirSize 获取目录下所有文件大小
func DirSize(dirname string) string {
	var (
		s        int64
		files, _ = WalkDir(dirname)
	)
	for _, n := range files {
		s += n.size
	}
	return FileSize(s)
}

func ConvertToStringSlice(data []interface{}) []string {
	result := make([]string, len(data))
	for i, v := range data {
		str, ok := v.(string)
		if !ok {
			// 如果类型断言失败，可以在此处进行相应的错误处理
			// 这里简单地将该元素转换为空字符串
			str = ""
		}
		result[i] = str
	}
	return result
}

// 删除文件
func DeleteFile(name string) error {
	//判断文件是否存在
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	//打开文件
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	//删除文件
	err = os.Remove(name)
	if err != nil {
		return err
	}

	return nil
}

// ReverseReadLines 从文件末尾开始逆序高效地读取行。
func ReverseReadLines(name string) ([]string, error) {
	// 打开文件。
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			g.Log().Error(context.Background(), err)
		}
	}(file)

	// 获取文件大小。
	fs, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fs.Size()

	// 定义合理的缓冲区大小进行读取。这个值可以根据预期的文件大小和行长度进行调整。
	const bufferSize = 1024 * 20

	// 创建一个切片用于保存行。
	var lines []string

	// 从文件末尾开始分块读取。
	for offset := fileSize; offset > 0; offset -= bufferSize {
		// 计算读取的大小。
		size := GetMin(bufferSize, offset)
		buffer := make([]byte, size)

		// 定位并读取这一块。
		_, err := file.ReadAt(buffer, offset-size)
		if err != nil {
			return nil, err
		}

		// 将缓冲区分割成行，并以逆序添加它们。
		scanner := bufio.NewScanner(bufio.NewReaderSize(file, int(size)))
		var chunkLines []string
		for scanner.Scan() {
			chunkLines = append([]string{scanner.Text()}, chunkLines...)
		}
		// 将这一块的行添加到总行中。
		lines = append(chunkLines, lines...)
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return lines, nil
}

func ReverseRead(name string, lineNum uint) ([]string, error) {
	//打开文件
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			g.Log().Error(context.Background(), err)
		}
	}(file)
	//获取文件大小
	fs, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fs.Size()

	var offset int64 = -1   //偏移量，初始化为-1，若为0则会读到EOF
	char := make([]byte, 1) //用于读取单个字节
	lineStr := ""           //存放一行的数据
	buff := make([]string, 0, 100)
	for (-offset) <= fileSize {
		//通过Seek函数从末尾移动游标然后每次读取一个字节
		_, err := file.Seek(offset, io.SeekEnd)
		if err != nil {
			return nil, err
		}
		_, err = file.Read(char)
		if err != nil {
			return buff, err
		}
		if char[0] == '\n' {
			offset--  //windows跳过'\r'
			lineNum-- //到此读取完一行
			buff = append(buff, lineStr)
			lineStr = ""
			if lineNum == 0 {
				return buff, nil
			}
		} else {
			lineStr = string(char) + lineStr
		}
		offset--
	}
	buff = append(buff, lineStr)
	//使用mahonia解码
	for i := 0; i < len(buff); i++ {
		buff[i], err = gcharset.ToUTF8("UTF-8", buff[i])
	}
	return buff, nil
}

// 文件一块一块的读取
func ReadBlock(filePath string) {
	start1 := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		g.Log().Error(context.Background(), err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			g.Log().Error(context.Background(), err)
		}
	}(file)
	// 设置每次读取字节数
	buffer := make([]byte, 1024*4)
	for {
		n, err := file.Read(buffer)
		// 控制条件,根据实际调整
		if err != nil && err != io.EOF {
			g.Log().Error(context.Background(), err)
		}
		if n == 0 {
			break
		}
		// 如下代码打印出每次读取的文件块(字节数)
		//fmt.Println(string(buffer[:n]))
	}
	fmt.Println("readBolck spend : ", time.Now().Sub(start1))
}

const (
	kilobyte = 1024
	megabyte = 1024 * kilobyte
	gigabyte = 1024 * megabyte
	terabyte = 1024 * gigabyte
	petabyte = 1024 * terabyte
)

// FormatSize 格式化文件大小。
func FormatSize(size int64) string {
	switch {
	case size < kilobyte:
		return strconv.Itoa(int(size)) + "B"
	case size < megabyte:
		return fmt.Sprintf("%.2fK", float64(size)/kilobyte)
	case size < gigabyte:
		return fmt.Sprintf("%.2fM", float64(size)/megabyte)
	case size < terabyte:
		return fmt.Sprintf("%.2fG", float64(size)/gigabyte)
	case size < petabyte:
		return fmt.Sprintf("%.2fT", float64(size)/terabyte)
	default:
		return fmt.Sprintf("%.2fP", float64(size)/petabyte)
	}
}

func ValidatePassword(password string, minimumLength int, requireComplexity int, requireDigit int, requireLowercase int, requireUppercase int) (flag bool, err error) {
	//初始化返回结果
	flag = true
	//判断密码长度
	if len(password) < minimumLength {
		err = gerror.New(fmt.Sprintf(g.I18n().T(context.TODO(), "{#utilsValidatePwLen}"), minimumLength))
		flag = false
		return
	}
	//是否有复杂字符
	if requireComplexity == 1 && !hasComplexCharacters(password) {
		err = gerror.New(g.I18n().T(context.TODO(), "{#utilsValidatePwChar}"))
		flag = false
		return
	}
	//是否有数字
	if requireDigit == 1 && !hasDigit(password) {
		err = gerror.New(g.I18n().T(context.TODO(), "{#utilsValidatePwDigit}"))
		flag = false
		return
	}
	//是否有小写字母
	if requireLowercase == 1 && !hasLowercaseLetter(password) {
		err = gerror.New(g.I18n().T(context.TODO(), "{#utilsValidatePwLower}"))
		flag = false
		return
	}
	//是否有大写字母
	if requireUppercase == 1 && !hasUppercaseLetter(password) {
		err = gerror.New(g.I18n().T(context.TODO(), "{#utilsValidatePwUpper}"))
		flag = false
		return
	}

	return
}

// hasComplexCharacters：检查字符串中是否有复杂字符（特殊字符）
func hasComplexCharacters(str string) bool {
	specialCharacters := "!@#$%^&*()_+-=[]{}|;:,.<>?~"

	for _, char := range str {
		if contains(specialCharacters, string(char)) {
			return true
		}
	}

	return false
}

// hasDigit：检查字符串中是否有数字
func hasDigit(str string) bool {
	for _, char := range str {
		if isDigit(string(char)) {
			return true
		}
	}

	return false
}

// hasLowercaseLetter：检查字符串中是否有小写字母
func hasLowercaseLetter(str string) bool {
	for _, char := range str {
		if isLowercase(string(char)) {
			return true
		}
	}

	return false
}

// hasUppercaseLetter：检查字符串中是否有大写字母
func hasUppercaseLetter(str string) bool {
	for _, char := range str {
		if isUppercase(string(char)) {
			return true
		}
	}

	return false
}

// isDigit：判断字符是否是数字
func isDigit(c string) bool {
	return c >= "0" && c <= "9"
}

// isLowercase：判断字符是否是小写字母
func isLowercase(c string) bool {
	return c >= "a" && c <= "z"
}

// isUppercase：判断字符是否是大写字母
func isUppercase(c string) bool {
	return c >= "A" && c <= "Z"
}

// contains：判断字符串是否包含指定字符
func contains(str, char string) bool {
	for _, c := range str {
		if string(c) == char {
			return true
		}
	}

	return false
}

// GetMin 返回两个整数中的较小值
func GetMin(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
