/*
 * @lc app=leetcode.cn id=334 lang=cpp
 *
 * [334] 递增的三元子序列
 */

// @lc code=start
class Solution {
   public:
    bool increasingTriplet(vector<int>& nums) {
        int len = nums.size();
        int x1 = INT_MAX;
        int x2 = INT_MAX;
        for (int i = 0; i < len; i++) {
            if (nums[i] <= x1)
                x1 = nums[i];
            else if (nums[i] <= x2)
                x2 = nums[i];
            else
                return true;
        }
        return false;
    }
};
// @lc code=end
