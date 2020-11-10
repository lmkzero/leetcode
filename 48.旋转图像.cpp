/*
 * @lc app=leetcode id=48 lang=cpp
 *
 * [48] Rotate Image
 */

// @lc code=start
class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        // const int n = matrix.size();

        // for (int i = 0; i < n; ++i)  // 沿着副对角线反转
        //     for (int j = 0; j < n - i; ++j)
        //         swap(matrix[i][j], matrix[n - 1 - j][n - 1 - i]);

        // for (int i = 0; i < n / 2; ++i) // 沿着水平中线反转
        //     for (int j = 0; j < n; ++j)
        //         swap(matrix[i][j], matrix[n - 1 - i][j]);

        for(int i=0; i<matrix.size(); i++)
        {
            for(int j=i; j<matrix[0].size(); j++)
                swap(matrix[i][j], matrix[j][i]);
        }
        
        for(int i=0; i<matrix.size(); i++)
        {
            for(int j=0, k=matrix[0].size()-1; j < k; j++, k--)
                swap(matrix[i][j], matrix[i][k]);
        }
    }
};
// @lc code=end

