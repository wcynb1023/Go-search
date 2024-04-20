package main

import "strconv"

func Atoi_list(str string) []int {
	arr := Get_index_list(str)
	intSlice1 := make([]int, len(arr))
	for j, val := range arr {
		intSlice1[j], _ = strconv.Atoi(val)
	}
	return intSlice1
}

func merge(arr []int, brr []int) []int {
	len_arr := len(arr)
	len_brr := len(brr)
	var ans []int
	var i = 0
	var j = 0
	for i < len_arr && j < len_brr {
		if arr[i] == brr[j] {
			ans = append(ans, arr[i])
			i += 1
			j += 1
		} else if arr[i] < brr[j] {
			i += 1
		} else if arr[i] > brr[j] {
			j += 1
		}
	}
	return ans
}

func merge_list(str []string) []int {
	ans := Atoi_list(str[0])

	for i := 1; i < len(str); i++ {
		ans = merge(ans, Atoi_list(str[i]))
	}

	return ans
}
