package main

import "fmt"

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC", "123123"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(list1, list2))
}

func findRestaurant(list1 []string, list2 []string) []string {
	hash := map[string]int{}
	for index, item := range list1 {
		hash[item] = index
	}
	result := make([]string, 0, len(hash))
	index := -1

	for i, k := range list2 {
		if j, exist := hash[k]; exist {
			if i+j > index && index != -1 {
				continue
			}
			if i+j == index {
				result = append(result, k)
				continue
			}
			index = i + j
			result = []string{k}
		}
	}

	return result
}
