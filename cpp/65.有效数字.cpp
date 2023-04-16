/*
 * @lc app=leetcode.cn id=65 lang=cpp
 *
 * [65] 有效数字
 */

// @lc code=start
class Solution {
   public:
    bool isNumber(string s) { return isNumber(s.c_str()); }

   private:
    bool isNumber(char const *s) {
        char *endptr;
        strtod(s, &endptr);
        if (endptr == s) {
            return false;
        }
        for (; *endptr; endptr++) {
            if (!isspace(*endptr)) {
                return false;
            }
        }
        return true;
    }
};
// @lc code=end
