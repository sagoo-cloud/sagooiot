package param

import "fmt"

func ConvertMapToSlice(data []map[string]interface{}) []map[string]interface{} {
	var transformed []map[string]interface{}
	for _, item := range data {
		transformedItem := make(map[string]interface{})
		for key, valueMap := range item {
			if subMap, ok := valueMap.(map[string]interface{}); ok {
				for subKey, subValue := range subMap {
					if val, ok := subValue.(float64); ok {
						transformedKey := fmt.Sprintf("%s_%s", key, subKey)
						transformedItem[transformedKey] = val
					}
				}
			}
		}
		transformed = append(transformed, transformedItem)
	}
	return transformed
}
