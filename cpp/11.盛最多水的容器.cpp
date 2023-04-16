/*
 * @lc app=leetcode.cn id=11 lang=cpp
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
class Solution {
   public:
    int maxArea(vector<int>& height) {
        int len = height.size();
        if (len <= 0) {
            return 0;
        }
        int start = 0;
        int end = len - 1;
        int result = 0;
        while (start < end) {
            int temp = (end - start) * (height[start] <= height[end] ? height[start] : height[end]);
            if (temp > result) {
                result = temp;
            }
            if (height[start] <= height[end]) {
                start++;
            } else {
                end--;
            }
        }
        return result;
    }
};
// @lc code=end
