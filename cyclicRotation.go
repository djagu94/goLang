package cyclicRotation

func ShiftKTimes(A []int, K int) []int {
    if len(A) < 2 {
        return A
    }

    for i := 0; i < K; i++ {
        ShiftOnce(A)
    }

    return A
}

func ShiftOnce(A []int){
    prev := A[0]
    for i := 1; i < len(A); i++ {       
        A[i], prev = prev, A[i]
    }
    A[0] = prev
}