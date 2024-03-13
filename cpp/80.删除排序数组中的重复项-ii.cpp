/*
 * @lc app=leetcode.cn id=80 lang=cpp
 *
 * [80] 删除排序数组中的重复项 II
 */

#include <iostream>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int removeDuplicates(vector<int>& nums) {
        if (nums.empty()) {
            return 0;
        }
        int L = 0;
        for (int R = 0; R < nums.size(); ++R) {
            // nums[L-2] < nums[R]
            // 如果重复数字个数不大于2，那么这个判断肯定是成立的
            // 如果重复个数大于2，那么nums[L-2] < nums[R]就是不成立的
            // 即多余重复项都被忽视了，直到遇到新的数字或者数字结束
            if (L < 2 || nums[L - 2] < nums[R]) {
                nums[L] = nums[R];
                ++L;
            }
        }
        return L;
    }
};
// @lc code=end
