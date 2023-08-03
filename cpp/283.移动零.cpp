/*
 * @lc app=leetcode.cn id=283 lang=cpp
 *
 * [283] 移动零
 */

// @lc code=start

#include <iostream>
#include <vector>

using namespace std;

class Solution {
   public:
    void moveZeroes(vector<int>& nums) {
        if (nums.size() <= 1) {
            return;
        }
        int slow = 0;
        int fast = 0;
        for (; fast < nums.size();) {
            if (nums[fast] != 0) {
                nums[slow] = nums[fast];
                slow++;
            }
            fast++;
        }
        for (; slow < nums.size(); slow++) {
            nums[slow] = 0;
        }
    }
};
// @lc code=end
