package homework

func average(input [15]float32) (result float32) {
	var sum float32
	var i int
	for _, value := range input {
		if value == 0 {
			continue
		}
		sum += value
		i += 1
	}
	result = sum / float32(i)
	return
}
