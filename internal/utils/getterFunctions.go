package utils

import "fmt"

func GetString(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok {
		return val.(string)
	}
	return ""
}

func GetInt(data map[string]interface{}, key string) int {
	if value, ok := data[key]; ok {
		switch v := value.(type) {
		case float64:
			return int(v)
		case int:
			return v
		default:
			fmt.Printf("unexpected type %T for key %s\n", v, key)
		}
	}
	return 0
}

func GetIntSlice(data map[string]interface{}, key string) []int {
	if val, ok := data[key]; ok {
		return val.([]int)
	}
	return []int{}
}