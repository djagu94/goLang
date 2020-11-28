package oddOccurrencesInArray

func FindOddOccurencies(A []int) int {
    result := 0

    for _, value := range A {
        result ^= value
    }

    return result
}