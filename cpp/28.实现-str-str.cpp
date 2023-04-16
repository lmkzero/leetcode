/*
 * @lc app=leetcode.cn id=28 lang=cpp
 *
 * [28] 实现 strStr()
 */

// @lc code=start
class Solution {
   public:
    int strStr(string haystack, string needle) {
        int len1 = haystack.length();
        int len2 = needle.length();
        int res = -1;
        if (len2 == 0) return 0;
        if (len1 < len2) return -1;
        for (int i = 0; i <= len1 - len2; i++) {
            // compare()
            //  param：起始位置，长度，子串
            res = haystack.compare(i, len2, needle);
            if (res == 0) return i;
        }
        return -1;
    }
};
// @lc code=end
