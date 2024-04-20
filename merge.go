package main

func merge(arr []int , brr []int) []int{
	len_arr := len(arr)
	len_brr := len(brr)
	var ans []int
	var i = 0 
	var j = 0
	for (i < len_arr && j < len_brr)  {
		if(arr[i] == brr[j]){
			ans = append(ans , arr[i])
			i+=1
			j+=1
		}else if(arr[i] < brr[j]){
			i+=1
		}else if(arr[i] > brr[j]){
			j+=1
		}
	}
	return ans
}
