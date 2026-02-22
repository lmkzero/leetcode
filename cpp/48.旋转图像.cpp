/*
 * @lc app=leetcode id=48 lang=cpp
 *
 * [48] Rotate Image
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    void rotate(vector<vector<int>>& matrix) {
        // 上下翻转
        for (int i = 0; i < matrix.size(); i++) {
            for (int j = i; j < matrix[0].size(); j++) {
                swap(matrix[i][j], matrix[j][i]);
            }
        }
        // 主对角线翻转
        for (int i = 0; i < matrix.size(); i++) {
            for (int j = 0, k = matrix[0].size() - 1; j < k; j++, k--) {
                swap(matrix[i][j], matrix[i][k]);
            }
        }
    }
};
// @lc code=end
