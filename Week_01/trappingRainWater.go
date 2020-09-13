
func trap(height []int) int {
    /*双指针法
    方法类似盛最多水的容器
    左右向中间逼近，矮的一方向高的一方走
    不断记录最高点
    矮于最高点的地方都能盛水
    复杂度：O(n) 
    */
    left, right := 0, len(height) - 1
    res := 0
    maxleft, maxright := 0, 0 //记录左右两边的最高点
    for left < right {
        if height[left] <= height[right] {//从矮的一边向高的一边走
            if height[left] >= maxleft { 
                maxleft = height[left] //更新最高点的位置
            } else {
                res += maxleft - height[left]  //当前值与最高点的差值即接到的雨水
            }
            left++
        } else {
            if height[right] >= maxright {
                maxright = height[right]
            } else {
                res += maxright - height[right]
            }
            right--
        }
    }
    return res
}