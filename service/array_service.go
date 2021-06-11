package service

type arrayService struct {
}

func NewArrayService() *arrayService {
	return &arrayService{}
}

func (s *arrayService) ArrayChallenge(arr []int) int {
	var maxValuesFromArrays []int

	for i := 0; i < len(arr); i++ {
		if i != 0 {
			maxValuesFromArrays = append(maxValuesFromArrays, s.BigestSumOfArray(s.RemoveIndexesFromArray(arr, i-1)))
		}

		maxValuesFromArrays = append(maxValuesFromArrays, s.BigestSumOfArray(arr))
	}

	return s.MaxValueFromArray(maxValuesFromArrays)
}

func (s *arrayService) BigestSumOfArray(arr []int) int {
	var maxValues []int
	var sum int

	for _, val := range arr {
		sum = sum + val
		maxValues = append(maxValues, sum)
	}

	return s.MaxValueFromArray(maxValues)
}

func (s *arrayService) MaxValueFromArray(arr []int) int {
	result := 0
	for i, val := range arr {
		if i == 0 || val > result {
			result = val
		}
	}

	return result
}

func (s *arrayService) RemoveIndexesFromArray(arr []int, index int) []int {
	var newArray []int

	for i, val := range arr {
		if i > index {
			newArray = append(newArray, val)
		}
	}

	return newArray
}
