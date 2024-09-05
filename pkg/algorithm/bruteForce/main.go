package bruteforce


// findElement method given array and k element
func FindElement(arr[10]int,k int) bool {
	for _,element:=range arr{
		if element == k {
			return true
		} 
	} 
	return false
}
