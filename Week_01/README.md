学习笔记

### 学习方法

* 三分看视频理解、七分练习：（1.5~2.0倍数播放、难点反复观看
* 摒弃坏习惯，不要死磕
   * 五毒神掌
   * 勤看高手代码（主要是国际版）
   * 最佳方式：5分钟想不出来，直接看题解。理解并背下来。

#### 五毒神掌

* 5-10 分钟读题和思考，如果没有思路，看题解，默写代码
* 马上自己写，提交LeetCode，多种解法，体会优化
* 每题至少做5遍，每遍有一定时间间隔，勤看国际版高票回答。把答案背下来。
* 面试前一周恢复性训练



##### 盛最多水的容器

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。 说明：你不能倾斜容器，且 n 的值至少为 2。

解法一：类似于冒泡排序，找出这些数中面积最大的数

解法二：双指针

​	两边向中间逼近，要找出面积最大，移动边小的一方。

```
func maxArea(height []int) int {
    //两路夹逼
    max := 0
    area := 0
    i := 0
    j := len(height) - 1
    for ; i < j; {
        if height[i] >= height[j] {
            area = (j - i) * height[j]
            j--
        } else {
            area = (j - i) * height[i]
            i++
        }
        if area > max {
            max = area
        }
    }
    return max
}
```

##### 接雨水

给定 *n* 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。 

解法一：暴力法：找到当前节点左右两边的最大值，然后用小的一个减去当前节点的值

- 时间复杂度： O(n^2)。数组中的每个元素都需要向左向右扫描。
- 空间复杂度 O(1) 的额外空间。

解法二：动态编程

* 找到数组中从下标 i 到最左端最高的条形块高度left_max。

* 找到数组中从下标 i 到最右端最高的条形块高度right_max。

* 扫描数组 height 并更新答案

* 累加 min(max_left[i],max_right[i])−height[i] 到 ans 上

  *时间复杂度：O(n)。*

  *空间复杂度：O(n)额外空间。*  

解法三：栈

时间复杂度：O(n)O(n)。
空间复杂度：O(n)。

解法四：双指针

- 时间复杂度：O(n)。
- 空间复杂度：O(1)

```
func trap(height []int) int {
    /*双指针法
    使用类似盛最多水的容器的方法
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
```

