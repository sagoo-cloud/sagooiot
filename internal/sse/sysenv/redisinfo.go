package sysenv

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"strings"
)

// GetRedisInfo 获取Redis运行信息
func GetRedisInfo() (data map[string][]byte) {
	ctx := gctx.GetInitCtx()
	conn, _ := g.Redis().Conn(ctx)
	defer func(conn gredis.Conn, ctx context.Context) {
		if err := conn.Close(ctx); err != nil {
			fmt.Println(err)
		}
	}(conn, ctx)
	info, _ := conn.Do(ctx, "INFO")

	var result = make(map[string][]map[string]interface{})
	scanner := bufio.NewScanner(strings.NewReader(info.String()))
	var key string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			key = strings.TrimSpace(strings.Split(line, "#")[1])
			result[key] = make([]map[string]interface{}, 0)
		} else if len(line) != 0 {
			kv := strings.Split(line, ":")
			m := make(map[string]interface{})

			//判断指标值是否有多个
			if strings.Contains(kv[1], ",") {
				sunValueList := strings.Split(kv[1], ",")
				sunValue := make(map[string]interface{})
				for _, s := range sunValueList {
					skv := strings.Split(s, "=")
					sunValue[skv[0]] = skv[1]
				}
				m[kv[0]] = sunValue
			} else {
				m[kv[0]] = kv[1]

			}

			result[key] = append(result[key], m)
		}
	}

	var res = make(map[string][]byte)
	for k, vList := range result {
		var value = make(map[string]interface{})
		for _, v := range vList {
			for k1, v1 := range v {
				value[k1] = v1
			}
		}
		res[k], _ = json.Marshal(value)
	}

	return res
}
