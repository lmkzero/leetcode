/*
 * @lc app=leetcode.cn id=78 lang=cpp
 *
 * [78] 子集
 */

#include <iostream>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
private:
    vector<vector<int>> res;
    vector<int> subset;

public:
    vector<vector<int>> subsets(vector<int>& nums) {
        dfs(nums, 0);
        return res;
    }

    void dfs(vector<int>& nums, int n) {
        res.push_back(subset);
        for (int i = n; i < nums.size(); i++) {
            subset.push_back(nums[i]);
            dfs(nums, i+1);
            subset.pop_back();
        }
    }
};
// @lc code=end

