package repositories

import (
	"errors"
	"sort"
)

type TaskLogic interface {
	TaskOne(nums []int, target int) ([]int, error)
	TaskTwo(nums []int) [][]int
	TaskThree(s string, words []string) []int
}

type taskLogic struct{}

func TaskLogicRepository() *taskLogic {
	return &taskLogic{}
}

//Task One
func (uc *taskLogic) TaskOne(nums []int, target int) ([]int, error) {
	temp := make(map[int]int)
	for i, num := range nums {
		result := target - num
		if index, err := temp[result]; err {
			return []int{index, i}, nil
		}
		temp[num] = i
	}
	return nil, errors.New("task one no solution")
}

//Task Two
func (uc *taskLogic) TaskTwo(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

//Untuk Task Three Tidak Selesai
func (uc *taskLogic) TaskThree(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	wordLength := len(words[0])
	totalWords := len(words)
	substringLength := wordLength * totalWords
	wordMap := make(map[string]int)
	result := []int{}

	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i <= len(s)-substringLength; i++ {
		substr := s[i : i+substringLength]
		foundWords := make(map[string]int)
		valid := true

		for j := 0; j < totalWords; j++ {
			start := j * wordLength
			word := substr[start : start+wordLength]
			if count, ok := wordMap[word]; ok {
				foundWords[word]++
				if foundWords[word] > count {
					valid = false
					break
				}
			} else {
				valid = false
				break
			}
		}

		if valid {
			result = append(result, i)
		}
	}

	return result
}
