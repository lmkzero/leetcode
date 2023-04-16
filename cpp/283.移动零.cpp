/*
 * @lc app=leetcode.cn id=283 lang=cpp
 *
 * [283] 移动零
 */

// @lc code=start
class Solution {
   public:
    void moveZeroes(vector<int>& nums) {
        int temp = 0;
        for (int i = 0, j = 0; i < nums.size(); ++i) {
            if (nums[i] != 0) {
                temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
                j++;
            }
        }
    }
};
// @lc code=end
