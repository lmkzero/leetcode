/*
 * @lc app=leetcode.cn id=518 lang=cpp
 *
 * [518] 零钱兑换 II
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int change(int amount, vector<int>& coins) {
        int n = coins.size();
        // 初始化 dp 表
        vector<int> dp(amount + 1, 0);
        dp[0] = 1;
        // 状态转移
        for (int i = 1; i <= n; i++) {
            for (int a = 1; a <= amount; a++) {
                if (coins[i - 1] > a) {
                    // 若超过目标金额，则不选硬币 i
                    continue;
                }
                // 不选和选硬币 i 这两种方案之和
                dp[a] = dp[a] + dp[a - coins[i - 1]];
            }
        }
        return dp[amount];
    }
};
// @lc code=end
