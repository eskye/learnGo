package main

import "fmt"

func main() {
	slice := []string{"apple", "banana", "orange", "date"}
	result := convertToMap(slice)
	fmt.Println(result)
	for k, v := range result {
		fmt.Printf("Key = %v, value = %v\n", k, v)
	}

}

func convertToMap(items []string) map[string]float64 {
	result := make(map[string]float64)

	elementValue := 100 / float64(len(items))
	for _, fruit := range items {
		result[fruit] = elementValue
	}
	return result
}
