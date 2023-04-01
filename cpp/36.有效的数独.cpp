/*
 * @lc app=leetcode.cn id=36 lang=cpp
 *
 * [36] 有效的数独
 */

// @lc code=start
class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        // unordered_set<char> row;
        // unordered_set<char> column;
        // unordered_set<char> cube;
        // for (int i = 0; i < 9; i++) {
        //     for (int j = 0; j < 9; j++) {
        //         if (board[i][j] != '.' && row.find(board[i][j])!=row.end()) {
        //             return false;
        //         }
        //         else {
        //             row.insert(board[i][j]);
        //         }
        //         if (board[j][i] != '.' && column.find(board[j][i])!=column.end()) {
        //             return false;
        //         }
        //         else {
        //             column.insert(board[j][i]);
        //         }
        //         int row_index = 3 * (i / 3) + j / 3;
        //         int column_index = 3 * (i % 3) + j % 3;
        //         if (board[row_index][column_index] != '.' && cube.find(board[row_index][column_index])!=cube.end()) {
        //             cout << row_index << "," << column_index << endl;
        //             return false;
        //         }
        //         else {
        //             cube.insert(board[row_index][column_index]);
        //         }
        //     }
        // }
        // return true;

        if (board.empty() || board[0].empty()) return false;
            int m = board.size(), n = board[0].size();
            vector<vector<bool> > rowFlag(m, vector<bool>(n, false));
            vector<vector<bool> > colFlag(m, vector<bool>(n, false));
            vector<vector<bool> > cellFlag(m, vector<bool>(n, false));
            for (int i = 0; i < m; ++i) {
                for (int j = 0; j < n; ++j) {
                    if (board[i][j] >= '1' && board[i][j] <= '9') {
                        int c = board[i][j] - '1';
                        if (rowFlag[i][c] || colFlag[c][j] || cellFlag[3 * (i / 3) + j / 3][c]) 
                            return false;
                        rowFlag[i][c] = true;
                        colFlag[c][j] = true;
                        cellFlag[3 * (i / 3) + j / 3][c] = true;
                    }
                }
            }
            return true;
    }
};
// @lc code=end

