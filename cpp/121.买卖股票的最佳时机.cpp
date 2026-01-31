/*
 * @lc app=leetcode.cn id=121 lang=cpp
 *
 * [121] 买卖股票的最佳时机
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int maxProfit(vector<int> &prices) {
        if (prices.size() <= 1) {
            return 0;
        }
        int max = 0;
        int slow = 0;
        int fast = 0;
        for (; fast < prices.size();) {
            int cur = prices[fast] - prices[slow];
            if (cur > max) {
                max = cur;
            }
            if (cur < 0) {
                slow = fast;
            }
            fast++;
        }
        return max;
    }
};
// @lc code=end
