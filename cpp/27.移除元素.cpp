/*
 * @lc app=leetcode.cn id=27 lang=cpp
 *
 * [27] 移除元素
 */

// @lc code=start
class Solution {
   public:
    int removeElement(vector<int>& nums, int val) {
        int len = nums.size();
        int n = 0;
        for (int i = 0, j = len - 1;; i++) {
            if (i < len - n) {
                if (nums[i] == val) {
                    nums[i] = nums[j];
                    i--;
                    j--;
                    n++;
                }
            } else {
                break;
            }
        }
        return len - n;
    }
};
// @lc code=end
