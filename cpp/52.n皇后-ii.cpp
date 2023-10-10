/*
 * @lc app=leetcode.cn id=52 lang=cpp
 *
 * [52] N皇后 II
 */

#include <vector>
#include <string>

using namespace std;

// @lc code=start
class Solution {
public:
    int totalNQueens(int n) {
        int res = 0;
        vector<string> board(n, string(n, '.'));
        backtrack(board, 0, res);
        return res;
    }
    void backtrack(vector<string> &board, int row, int &res) {
        if (row == board.size()) {
            res++;
            return;
        }

        int n = board[row].size();
        for (int col = 0; col < n; ++col) {
            if (!isValid(board, row, col)) {
                continue;
            }

            board[row][col] = 'Q';
            backtrack(board, row + 1, res);
            board[row][col] = '.';
        }
    }

    bool isValid(vector<string> &board, int row, int col) {
        int n = board.size();
        // 检查同列
        for (int i = 0; i < row; i++) {
            if (board[i][col] == 'Q') {
                return false;
            }
        }
        // 右上
        for (int i = row - 1, j = col + 1; i >= 0 && j < n; i--, j++) {
            if (board[i][j] == 'Q') {
                return false;
            }
        }
        // 左上
        for (int i = row - 1, j = col - 1; i >= 0 && j >= 0; i--, j--) {
            if (board[i][j] == 'Q') {
                return false;
            }
        }
        return true;
    }
};
// @lc code=end

