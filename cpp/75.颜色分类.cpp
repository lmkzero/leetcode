/*
 * @lc app=leetcode.cn id=75 lang=cpp
 *
 * [75] 颜色分类
 */

// @lc code=start
class Solution {
   public:
    void sortColors(vector<int>& nums) {
        int red = 0;
        int blue = nums.size() - 1;
        for (int i = 0; i < blue + 1;) {
            if (nums[i] == 0) {
                swap(nums[i++], nums[red++]);
            } else if (nums[i] == 2) {
                swap(nums[i], nums[blue--]);
            } else {
                i++;
            }
        }
    }
};
// @lc code=end
