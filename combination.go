package conmbinnation

//SampleIndexs 获取所有组合的元素索引
func SampleIndexs(total, n int) [][]int {
	var result [][]int
	indexs := make([]int, total)
	for i := 0; i < n; i++ {
		indexs[i] = 1
	}

	tmpindex := make([]int, total)
	copy(tmpindex, indexs)
	result = append(result, tmpindex)
	for {
		find := false
		for i := 0; i < total-1; i++ {
			if indexs[i]-indexs[i+1] == 1 {
				indexs[i], indexs[i+1] = 0, 1
				moveOneToleft(indexs[:i])
				tmpindex := make([]int, total)
				copy(tmpindex, indexs)
				result = append(result, tmpindex)
				find = true
				break
			}
		}
		if !find {
			break
		}
	}
	return result
}

func moveOneToleft(leftindexs []int) {
	sum := 0
	for _, v := range leftindexs {
		if v == 1 {
			sum++
		}
	}
	for i := range leftindexs {
		if i < sum {
			leftindexs[i] = 1
		} else {
			leftindexs[i] = 0
		}
	}
}
