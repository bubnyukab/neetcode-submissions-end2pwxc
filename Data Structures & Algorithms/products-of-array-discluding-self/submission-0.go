func productExceptSelf(nums []int) []int {
    answer := []int{}

    for i := range nums {
        product := 1
        for j := range nums {
            if i == j {
                continue
            }
            product *= nums[j]
        }
    
        answer = append(answer, product)
        product = 1
    }

    return answer
}