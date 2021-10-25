package main

import "fmt"

func Search(words []string) [][]string {
	var groups []string
	var groupDone []int
	var results [][]string
	if len(words) == 0 {
		return nil
	}

	for i := 0; i < len(words); i++ {
		if !isAlreadyExists(groupDone, i) {
			groups, groupDone = searchSeconds(words, i, groupDone)
			results = append(results, groups)
		}
	}

	return results
}

func isAlreadyExists(group []int, value int) bool {
	for i := 0; i < len(group); i++ {
		if group[i] == value {
			return true
		}
	}

	return false
}

func searchSeconds(words []string, no int, group []int) ([]string, []int) {
	var result []string
	result = append(result, words[no])
	//group = append(group, no)

	for i := 0; i < len(words); i++ {
		if i != no {
			if isAnagram(words[no], words[i]) {
				result = append(result, words[i])
				group = append(group, i)
			}
		}
	}

	return result, group
}

func isAnagram(wordsA, wordsB string) bool {
	hash := make(map[string]int)

	for _, r := range wordsA {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}

	for _, r := range wordsB {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}

	var isAnagram bool = true
	for _, value := range hash {
		if value%2 != 0 {
			isAnagram = false
			return isAnagram
		}

	}

	return isAnagram
}

func main() {
	words := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	result := Search(words)
	fmt.Println(result)
}
