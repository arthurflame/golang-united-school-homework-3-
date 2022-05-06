package homework

func average(input [15]float32) (result float32) {
	var sum float32
	var elements float32

	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			sum += input[i]
			elements++
		}
	}
	return sum / elements
}
