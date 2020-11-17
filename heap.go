package main

import (
    "fmt"
)

func swap(array []int, firstIndex int, secondIndex int){
    fmt.Println("Before swap", array[firstIndex], array[secondIndex])
    temp := array[secondIndex]
    array[secondIndex] = array[firstIndex]
    array[firstIndex] = temp
    fmt.Println("After swap", array[firstIndex], array[secondIndex])
    
}

func Heapify(arr []int, idx int){
    left := 2*idx + 1
    right := 2 *idx + 2
    largest := idx
    if left <= len(arr) && right <= len(arr){
    if arr[left] > arr[largest]{
        largest = left
    }
    if arr[right] > arr[largest]{
        largest = right
    }
    if largest!= idx{
        swap(arr, idx, largest)
    } else {
        return
    }
    Heapify(arr, largest)
   }
}

func DeleteMax(arr []int) int{
    fmt.Println("Array", arr)
    swap(arr, 0, len(arr)-1)
    max := arr[0]
    arr = arr[:len(arr)-1]
    fmt.Println("After", arr)
    return max
}


func BuildHeap(arr []int){
    n := int(0)
    for n <= int(len(arr)/2){
        fmt.Println("The value of n", n)
        Heapify(arr, n)
        n++
    }
}

func main(){
    fmt.Println("Hello world")
    data := []int{10,20,5,1,0,50,60}
    // Heapify(data, 0)
    // fmt.Println("Data", data)
    // for len(data) >= 0{
    //     Heapify(data, 0)
    //     fmt.Println("Deleted element",DeleteMax(data))
    // }
    BuildHeap(data)
    fmt.Println("After build heap procedure", data)

}