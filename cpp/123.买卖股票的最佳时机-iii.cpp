/*
 * @lc app=leetcode.cn id=123 lang=cpp
 *
 * [123] 买卖股票的最佳时机 III
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int maxProfit(vector<int>& prices) {
        int maxK = 2, n = prices.size();
        vector<vector<vector<int>>> dp(n, vector<vector<int>>(maxK + 1, vector<int>(2)));
        for (int i = 0; i < n; i++) {
            for (int k = maxK; k >= 1; k--) {
                if (i - 1 == -1) {
                    // 处理 base case
                    dp[i][k][0] = 0;
                    dp[i][k][1] = -prices[i];
                    continue;
                }
                dp[i][k][0] = max(dp[i - 1][k][0], dp[i - 1][k][1] + prices[i]);
                dp[i][k][1] = max(dp[i - 1][k][1], dp[i - 1][k - 1][0] - prices[i]);
            }
        }
        return dp[n - 1][maxK][0];
    }
};
// @lc code=end
