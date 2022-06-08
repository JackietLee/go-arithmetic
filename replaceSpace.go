package main

func replaceSpace(str string) string {
	count := 0
	for i := range str {
		if ' ' == str[i] {
			count++
		}
	}
	bytes := make([]byte, len(str)+2*count)
	index1, index2 := 0, 0
	for ; index1 < len(str); index1++ {
		if str[index1] != ' ' {
			bytes[index2] = str[index1]
			index2++
		} else {
			// 相当于空格替换成"%20"
			bytes[index2] = '%'
			bytes[index2+1] = '2'
			bytes[index2+2] = '0'
			index2 += 3
		}
	}
	return string(bytes)
}
