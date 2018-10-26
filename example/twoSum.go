package main

import "fmt"

/*func twoSum(nums []int, target int) []int {
	for i:= 0; i< len(nums); i++ {
		for j:= 0;j< i;j++ {
			if nums[i] + nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	panic("err no")

}*/



/*func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i,v := range nums {
		m[v] = i
	}
	for _, q := range nums {
		x, _ := m[q]
		y, ok := m[target - q]
		if ok && x != y {
			return []int{x, y}
		}
	}
	panic("err")

} */


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i,v := range nums {
		m[v] = i
		x, _ := m[v]
		y, ok := m[target - v]
		if ok && x != y {
			return []int{y, x}
		}

	}
	panic("err")

}


func main()  {

	a := []int{3, 2, 4}
	b := 6

	m := twoSum(a, b)
	fmt.Println(m)
}
