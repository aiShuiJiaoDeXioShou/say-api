package dao

import "say/src/modle"

func getPage(pages ...int) int {
	var page int
	if len(pages) == 0 {
		page = 0
	} else {
		page = pages[0]
	}
	return page
}

func getAttributes(arr []modle.UserInfo) []int {
	arrs := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		userInfo := arr[i]
		arrs[i] = userInfo.ID
	}
	return arrs
}