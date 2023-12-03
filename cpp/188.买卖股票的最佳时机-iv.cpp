/*
 * @lc app=leetcode.cn id=188 lang=cpp
 *
 * [188] 买卖股票的最佳时机 IV
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int maxProfit(int k, vector<int>& prices) {
        int maxK = k;
        int n = prices.size();
        if (n <= 0) {
            return 0;
        }
        if (maxK > n / 2) {
            return maxProfitII(prices);
        }
        vector<vector<vector<int>>> dp(n, vector<vector<int>>(maxK + 1, vector<int>(2)));
        for (int i = 0; i < n; i++) {
            dp[i][0][1] = INT_MIN;
            dp[i][0][0] = 0;
        }

        for (int i = 0; i < n; i++) {
            for (int k = maxK; k >= 1; k--) {
                if (i - 1 == -1) {
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

    int maxProfitII(vector<int>& prices) {
        int n = prices.size();
        vector<vector<int>> dp(n, vector<int>(2, 0));
        for (int i = 0; i < n; i++) {
            if (i == 0) {
                dp[i][0] = 0;
                dp[i][1] = -prices[i];
                continue;
            }
            dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
            dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i]);
        }
        return dp[n - 1][0];
    }
};
// @lc code=end
