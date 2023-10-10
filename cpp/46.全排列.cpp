/*
 * @lc app=leetcode.cn id=46 lang=cpp
 *
 * [46] 全排列
 */

#include <iostream>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    void permute(vector<int> &nums, vector<vector<int>> &res, int start, int len) {
        if (start == len) {
            res.emplace_back(nums);
            return;
        }
        for (int i = start; i < len; i++) {
            swap(nums[i], nums[start]);
            permute(nums, res, start + 1, len);
            swap(nums[i], nums[start]);
        }
    }
    vector<vector<int>> permute(vector<int> &nums) {
        vector<vector<int>> res;
        permute(nums, res, 0, nums.size());
        return res;
    }
};
// @lc code=end
