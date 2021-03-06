package BinarySearch

func SearchR(data []int, target int) int {
	return searchR(data, 0, len(data)-1, target)
}

//递归式实现
func searchR(data []int, l int, r int, target int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if data[mid] == target {
		return mid
	} else if data[mid] > target {
		return searchR(data, l, mid-1, target)
	} else {
		return searchR(data, mid+1, r, target)
	}
}

//非递归式实现
func Search(data []int, target int) int {
	l := 0
	r := len(data) - 1
	for l <= r {
		mid := l + (r-l)/2
		if data[mid] == target {
			return mid
		} else if data[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

//>target的最小索引
func upper(data []int, target int) int {
	l := 0
	r := len(data)
	if l < r {
		mid := l + (r-l)/2
		if data[mid] < target {
			l = mid + 1
		} else if data[mid] > target {
			r = mid
		}
	}
	return l
}
