package main

func GetAllHourglasses(matrix [][]int) [][]int {

	result := [][]int{}

	for i, line := range matrix {

		if i > len(line)-3 {
			break
		}

		for j := range line {

			if j > len(line)-3 {
				break
			}

			result = append(result, GetHourglass(matrix, i, j))
		}
	}

	return result
}

func GetHourglass(matrix [][]int, coordinate1 int, coordinate2 int) []int {
	result := []int{}

	for i := 0; i < 3; i += 1 {
		result = append(result, matrix[coordinate1][coordinate2+i])
	}

	result = append(result, matrix[coordinate1+1][coordinate2+1])

	for i := 0; i < 3; i += 1 {
		result = append(result, matrix[coordinate1+2][coordinate2+i])
	}

	return result
}

func GetHourglassSum(hourglass []int) int {

	result := 0

	for _, value := range hourglass {
		result += value
	}

	return result
}

func GetHighestHourglassSum(hourglasses [][]int) int {
	highestHourglassSum := -64

	for _, hourglass := range hourglasses {
		sumValue := GetHourglassSum(hourglass)

		if sumValue > highestHourglassSum {
			highestHourglassSum = sumValue
		}
	}

	return highestHourglassSum
}

func OptimizedHighestHourglassSum(matrix [][]int) int {
	highestHourglassSum := -64

	for i := 0; i < len(matrix)-3; i += 1 {

		for j := 0; j < len(matrix[i])-3; j += 1 {

			sum := matrix[i][j] + matrix[i][j+1] + matrix[i][j+2] +
				matrix[i+1][j+1] +
				matrix[i+2][j] + matrix[i+2][j+1] + matrix[i+2][j+2]

			if highestHourglassSum < sum {
				highestHourglassSum = sum
			}
		}
	}

	return highestHourglassSum
}
