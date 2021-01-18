// Problem: https://leetcode.com/problems/four-divisors/
// 4 tricks here
// 1. For numbers less than 5, there are no case for 4 divisor present. It starts with 6 (1,2,3,6)
// 2. Run the loop to the square root of its number, as we do check for prime condition
// 3. if root * root == number, it means it has only 3 divisors as it is a perfect square
// 4. Find the two remaining divisors and 1 as well as the value to the sum 
func sumFourDivisors(nums []int) int {
    divCount := 0
    sum := 0 
    tmp := 0 
    root := 0
    for _, val := range nums {
        root = int(math.Sqrt(float64(val)))
        for i := 2; i <= root ; i++ {
             if root * root == val { break }
            if val%i == 0 && val > 5  {
              tmp += i
              divCount += 2
              tmp += val/i
                fmt.Println(val, ":", tmp, divCount,  int(math.Sqrt(float64(val))))
            }   
            if divCount > 5 {
                break
            }
        }
        if divCount == 2 {
            sum = sum + tmp + val + 1
        }
        tmp = 0
        divCount = 0
    }
    return sum
}