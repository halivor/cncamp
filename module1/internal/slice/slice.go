package slice

import "fmt"

func Run() {
	strSlice := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Println(strSlice)
	for idx, val := range strSlice {
		strSlice[idx] = replace(val)
	}
	fmt.Println(strSlice)
}

func replace(val string) string {
	valueMap := map[string]string{
		"stupid": "smart",
		"weak":   "strong",
	}
	if data, ok := valueMap[val]; ok {
		return data
	}
	return val
}
