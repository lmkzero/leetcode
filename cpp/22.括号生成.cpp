/*
 * @lc app=leetcode.cn id=22 lang=cpp
 *
 * [22] 括号生成
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    vector<string> generateParenthesis(int n) {
        dfs(n, 0, 0, "");
        return ans;
    }

   private:
    vector<string> ans;
    void dfs(int n, int left, int right, string s) {
        if (left > n || right > n || left < right) {
            return;
        }
        if (left == n && right == n) {
            ans.push_back(s);
            return;
        }
        dfs(n, left + 1, right, s + "(");
        dfs(n, left, right + 1, s + ")");
    }
};
// @lc code=end
