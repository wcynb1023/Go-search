package main

import (
	"strconv"
	"strings"
	"sync"
)

func Atoi_list(str string) []int {
	str = strings.ToLower(str)
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
	i, j := 0, 0
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

	if len(ans) == 0 {
		ans = append(arr, brr...)
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

func merge_list_dc(str []string, l int, r int) []int {

	if l == r {
		return Atoi_list(str[l])
	}
	mid := (l + r) / 2
	var l_list, r_list []int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		l_list = merge_list_dc(str, l, mid)
	}()
	go func() {
		defer wg.Done()
		r_list = merge_list_dc(str, mid+1, r)
	}()

	wg.Wait()
	return merge(l_list, r_list)

}
