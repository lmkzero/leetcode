/*
 * @lc app=leetcode.cn id=289 lang=cpp
 *
 * [289] 生命游戏
 */

// @lc code=start
class Solution {
public:
    void gameOfLife(vector<vector<int>>& board) {
        int m = board.size();
        int n = board[0].size();
        // clock wise, start from upper-left corner
        int dx[] = {-1, -1, -1, 0, 1, 1, 1, 0};
        int dy[] = {-1, 0, 1, 1, 1, 0, -1, -1};

        // encode
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                int live = 0; // number of live cells
                for (int k = 0; k < 8; ++k) {
                    int x = i + dx[k];
                    int y = j + dy[k];
                    if (x > -1 && x < m && y > -1 && y < n &&
                            (board[x][y] == 1 || board[x][y] == 2)) {
                        ++live;
                    }
                }
                if (board[i][j] == 0 && live == 3) {
                    board[i][j] = 3;
                } else if (board[i][j] == 1 && (live < 2 || live > 3)) {
                    board[i][j] = 2;
                }
            }
        }

        //decode
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                board[i][j] %= 2;
            }
        }
    }
};
// @lc code=end

