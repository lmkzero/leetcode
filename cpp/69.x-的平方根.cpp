/*
 * @lc app=leetcode.cn id=69 lang=cpp
 *
 * [69] x 的平方根
 */

// @lc code=start
class Solution {
   public:
    int mySqrt(int x) {
        int min = 0;
        int max = x;
        int ans = -1;
        while (min <= max) {
            int m = (max + min) / 2;
            if ((long long)m * m <= x) {
                ans = m;
                min = m + 1;
                continue;
            }
            max = m - 1;
        }
        return ans;
    }
};
// @lc code=end
