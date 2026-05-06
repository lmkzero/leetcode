/*
 * @lc app=leetcode.cn id=624 lang=cpp
 *
 * [624] 数组列表中的最大距离
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int maxDistance(vector<vector<int>>& arrays) {
        int res = 0;
        int n = arrays[0].size();
        int min_val = arrays[0][0];
        int max_val = arrays[0][arrays[0].size() - 1];
        for (int i = 1; i < arrays.size(); i++) {
            n = arrays[i].size();
            res = max(res, max(abs(arrays[i][n - 1] - min_val), abs(max_val - arrays[i][0])));
            min_val = min(min_val, arrays[i][0]);
            max_val = max(max_val, arrays[i][n - 1]);
        }
        return res;
    }
};
// @lc code=end
