/*
 * @lc app=leetcode.cn id=118 lang=cpp
 *
 * [118] 杨辉三角
 */

// @lc code=start
class Solution {
   public:
    vector<vector<int>> generate(int numRows) {
        if (numRows == 1) {
            return {{1}};
        }
        if (numRows == 2) {
            return {{1}, {1, 1}};
        }
        vector<vector<int>> nums{{1}, {1, 1}};
        vector<int> temp;
        for (int i = 2; i < numRows; i++) {
            temp.clear();
            temp.push_back(1);
            for (int j = 0; j < i - 1; j++) {
                temp.push_back(nums[i - 1][j] + nums[i - 1][j + 1]);
            }
            temp.push_back(1);
            nums.push_back(temp);
        }
        return nums;
    }
};
// @lc code=end
