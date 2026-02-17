/*
 * @lc app=leetcode.cn id=695 lang=cpp
 *
 * [695] 岛屿的最大面积
 */

#include <algorithm>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int maxAreaOfIsland(vector<vector<int>>& grid) {
        int ans = 0;
        for (int i = 0; i < grid.size(); i++) {
            for (int j = 0; j < grid[i].size(); j++) {
                ans = max(ans, dfs(grid, i, j));
            }
        }
        return ans;
    }

   private:
    vector<vector<int>> dirs = {
        {-1, 0},
        {0, 1},
        {1, 0},
        {0, -1},
    };
    int dfs(vector<vector<int>>& grid, int i, int j) {
        if (i < 0 || i >= grid.size() || j < 0 || j >= grid[0].size() || grid[i][j] == 0) {
            return 0;
        }
        int area = 1;
        grid[i][j] = 0;
        for (auto dir : dirs) {
            int x = i + dir[0];
            int y = j + dir[1];
            area += dfs(grid, x, y);
        }
        return area;
    }
};
// @lc code=end
