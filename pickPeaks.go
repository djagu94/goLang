package kata

type PosPeaks struct {
  Pos   []int
  Peaks []int
}

func PickPeaks(array []int) PosPeaks {
  var posPeak PosPeaks
  posPeak.Pos = []int{}
  posPeak.Peaks = []int{}
  
  if len(array) < 2 {
    return posPeak
  }
  
  prevVal := array[0]
  plateauStart := 1
  
  for i := 1; i < len(array) - 1; i++ {
    
    
    
    if prevVal < array[i] && array[i+1] < array[i] {
      posPeak.Pos = append(posPeak.Pos, plateauStart)
      posPeak.Peaks = append(posPeak.Peaks, array[i])
    }
    
    //In case of plateau starting
    if array[i] != array[i+1]{
      prevVal = array[i]
      plateauStart = i+1
    }
  }
  
  return posPeak
}