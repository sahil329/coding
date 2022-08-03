package main

/**
  Insert value in slice at given index
*/
func insert(a *[]int, index int, value int) {
    if len(*a) == index { // nil or empty slice or after last element
        *a = append(*a, value)
        return 
    }
    *a = append((*a)[:index+1], (*a)[index:]...) // index < len(a)
    (*a)[index] = value
}

/**
  Implementation of bisect.bisect_left python pkg.
*/
func /*(this *MyCalendar)*/ bisectLeft(a *[]int, val int) int {
    left := 0
    right := len(*a)
    
    for left<right {
        mid := left + (right-left)/2
        if (*a)[mid] < val{
            left = mid + 1
        }else{
            right = mid
        }
    }
    return left
}

/**
  Implementation of bisect.bisect_right python pkg.
*/
func bisectRight(a *[]int, val int) int {
    left := 0
    right := len(*a)
    
    for left<right {
        mid := left + (right-left)/2
        if (*a)[mid] <= val{
            left = mid + 1
        }else{
            right = mid
        }
    }
    return left
}
