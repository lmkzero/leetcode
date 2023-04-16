/*
 * @lc app=leetcode.cn id=5 lang=cpp
 *
 * [5] 最长回文子串
 */

// @lc code=start
class Solution {
   public:
    string longestPalindrome(string s) {
        // https://www.jianshu.com/p/a7741619dd58
        const int n = s.size();
        if (n <= 1) {
            return s;
        }
        bool f[n][n];
        fill_n(&f[0][0], n * n, false);
        size_t max_len = 1;
        size_t start = 0;
        for (size_t i = 0; i < n; i++) {
            f[i][i] = true;
            for (size_t j = 0; j < i; j++) {
                f[j][i] = (s[j] == s[i] && (i - j < 2 || f[j + 1][i - 1]));
                if (f[j][i] && max_len < (i - j + 1)) {
                    max_len = i - j + 1;
                    start = j;
                }
            }
        }
        return s.substr(start, max_len);
    }
};
// @lc code=end