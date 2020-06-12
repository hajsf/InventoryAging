// Another option is to usea imath package
package methods

func MinOf(vars ...int64) int64 {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func MaxOf(vars ...int64) int64 {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func AbsOf(num int64) int64 {
	if num < 0 {
		num = -1 * num
	}
	return num
}
