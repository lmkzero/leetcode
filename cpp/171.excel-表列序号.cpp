/*
 * @lc app=leetcode.cn id=171 lang=cpp
 *
 * [171] Excel 表列序号
 */

// @lc code=start
class Solution {
public:
    int titleToNumber(string columnTitle) {
        int res = 0;
        reverse(columnTitle.begin(), columnTitle.end());
        for (int i = 0; i < columnTitle.length(); i++) {
            res += (columnTitle[i] - 'A' + 1) * pow(26, i);
        }
        return res;
    }
};
// @lc code=end

