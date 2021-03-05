package kata

func Tribonacci(signature [3]float64, n int) []float64 {
  var i = 0
  output := []float64{}
  for ; i < 3 && i < n; i++ {
    output = append(output, signature[i])
  }
  
  for ; i < n; i++ {
    output = append(output, GetSum(output[i-3:i]))
  }
  
  return output
}

func GetSum(input []float64) float64 {
  return input[0] + input[1] + input[2]
}