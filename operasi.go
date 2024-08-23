package main

import (
  "fmt"
)

// 1. Append
func appendSlices(slice1, slice2 []int) []int {
  result := make([]int, len(slice1)+len(slice2))
  for i := 0; i < len(slice1); i++ {
    result[i] = slice1[i]
  }
  for i := 0; i < len(slice2); i++ {
    result[len(slice1)+i] = slice2[i]
  }
  return result
}

// 2. Concat
func concatSlices(slices [][]int) []int {
  var totalLength int
  for _, slice := range slices {
    totalLength += len(slice)
  }
  result := make([]int, totalLength)
  var currentIndex int
  for _, slice := range slices {
    for _, val := range slice {
      result[currentIndex] = val
      currentIndex++
    }
  }
  return result
}

// 3. Filter
func filterSlice(slice []int, predicate func(int) bool) []int {
  var result []int
  for _, v := range slice {
    if predicate(v) {
      result = append(result, v) 
    }
  }
  return result
}

// 4. Map
func mapSlice(slice []int, f func(int) int) []int {
  result := make([]int, len(slice))
  for i, v := range slice {
    result[i] = f(v)
  }
  return result
}

// 5. Reverse
func reverseSlice(slice []int) []int {
  result := make([]int, len(slice))
  for i, v := range slice {
    result[len(slice)-1-i] = v
  }
  return result
}

func main() {
  // 1. Append 
  slice1 := []int{1, 2, 3}
  slice2 := []int{4, 5, 6}
  fmt.Println("Append:", appendSlices(slice1, slice2))

  // 2. Concat 
  slices := [][]int{{1, 2}, {3, 4}, {5, 6}}
  fmt.Println("Concat:", concatSlices(slices))

  // 3. Filter 
  slice := []int{1, 2, 3, 4, 5, 6}
  isEven := func(x int) bool { return x%2 == 0 }
  fmt.Println("Filter:", filterSlice(slice, isEven))

  // 4. Map 
  double := func(x int) int { return x * 2 }
  fmt.Println("Map :", mapSlice(slice1, double))

  // 5. Reverse 
  slice3 := []int{1, 2, 3, 4}
  fmt.Println("Reverse:", reverseSlice(slice3))
}
