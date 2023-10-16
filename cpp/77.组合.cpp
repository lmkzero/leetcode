/*
 * @lc app=leetcode.cn id=77 lang=cpp
 *
 * [77] 组合
 */

#include <iostream>
#include <vector>

using namespace std;

// @lc code=start
#include <vector>
class Solution {
private:
    vector<vector<int>> res;
    vector<int> subset;
public:
    vector<vector<int>> combine(int n, int k) {
        dfs(1, n, k);
        return res;
    }
    
    void dfs(int start, int end, int k) {
        if (subset.size() == k) {
            res.push_back(subset);
            return;
        }
        for (int i = start; i<= end; i++) {
            subset.push_back(i);
            dfs(i+1, end, k);
            subset.pop_back();
        }
    }
};
// @lc code=end

