package main

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	middle := 0

	for numbers[left] >= numbers[right] {
		if right-left == 1 {
			middle = right
			break
		}
		middle = (left + right) / 2
		if numbers[left] == numbers[middle] && numbers[right] == numbers[middle] {
			return inOrder(numbers, left, right)
		}
		if numbers[middle] >= numbers[left] {
			left = middle
		} else if numbers[middle] <= numbers[right] {
			right = middle
		}
	}
	return numbers[middle]
}

func inOrder(numbers []int, left, right int) (result int) {
	result = numbers[left]
	for i := left + 1; i < right; i++ {
		if numbers[i] < result {
			result = numbers[i]

		}
	}
	return
}
