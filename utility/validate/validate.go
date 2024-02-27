package validate

import (
	_ "github.com/gogf/gf/v2/frame/g"
)

// InSlice 元素是否存在于切片中
func InSlice[K comparable](slice []K, key K) bool {
	for _, v := range slice {
		if v == key {
			return true
		}
	}
	return false
}
