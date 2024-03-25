/*
 * @lc app=leetcode.cn id=55 lang=cpp
 *
 * [55] 跳跃游戏
 */

#include <algorithm>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    bool canJump(vector<int>& nums) {
        int length = nums.size();
        int rightMost = 0;
        for (int i = 0; i < length; i++) {
            if (i <= rightMost) {
                rightMost = max(rightMost, nums[i] + i);
                if (rightMost >= length - 1) {
                    return true;
                }
            }
        }
        return false;
    }
};
// @lc code=end
