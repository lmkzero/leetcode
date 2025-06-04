/*
 * @lc app=leetcode.cn id=73 lang=cpp
 *
 * [73] 矩阵置零
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    void setZeroes(vector<vector<int>>& matrix) {
        int h = matrix.size();
        int w = matrix[0].size();
        vector<bool> row(h, false);
        vector<bool> column(w, false);
        for (int i = 0; i < h; i++) {
            for (int j = 0; j < w; j++) {
                if (matrix[i][j] == 0) {
                    row[i] = true;
                    column[j] = true;
                }
            }
        }
        for (int i = 0; i < h; i++) {
            if (row[i] == true) {
                for (int j = 0; j < w; j++) {
                    matrix[i][j] = 0;
                }
            }
        }
        for (int i = 0; i < w; i++) {
            if (column[i] == true) {
                for (int j = 0; j < h; j++) {
                    matrix[j][i] = 0;
                }
            }
        }
    }
};
// @lc code=end
