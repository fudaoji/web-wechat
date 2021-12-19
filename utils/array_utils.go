package utils

//函数Contains用于判断数组$array中是否包含元素$search
func Contains(array []interface{}, search interface{}) (exists bool, index int) {
	exists = false
	index = -1

	for i, v := range array {
		if search == v {
			index = i
			exists = true
			return
		}
	}
	return
}
