package binaryGap

import "strconv"

func BinaryGap(N int) int {
    binary := strconv.FormatInt(int64(N), 2)
    longestSequence := 0
    currentSequence := 0
    previousChar := 'x' //x means the previous char was either a beginning of string or a '0' not preceded by any '1'

    for _, currentChar := range binary {
        if currentChar == '1' {
            if currentSequence > longestSequence {
                longestSequence = currentSequence
            }
            currentSequence = 0
            previousChar = currentChar
        }else if previousChar != 'x' {
            previousChar = currentChar
                currentSequence++
        }
    }

    return longestSequence
}