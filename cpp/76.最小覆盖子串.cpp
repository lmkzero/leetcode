/*
 * @lc app=leetcode.cn id=76 lang=cpp
 *
 * [76] 最小覆盖子串
 */

#include <climits>
#include <string>
#include <unordered_map>

using namespace std;

// @lc code=start
class Solution {
   public:
    string minWindow(string s, string t) {
        unordered_map<char, int> chars, window;
        for (char c : t) {
            chars[c]++;
        }
        int left = 0, right = 0;
        int valid = 0;
        int begin = 0, len = INT_MAX;
        while (right < s.size()) {
            char c = s[right];
            right++;
            if (chars.count(c)) {
                window[c]++;
                if (window[c] == chars[c]) {
                    valid++;
                }
            }
            while (valid == chars.size()) {
                if (right - left < len) {
                    begin = left;
                    len = right - left;
                }
                char moveC = s[left];
                if (chars.count(moveC)) {
                    if (window[moveC] == chars[moveC]) {
                        valid--;
                    }
                    window[moveC]--;
                    left++;
                    continue;
                }
                window[moveC]--;
                left++;
            }
        }
        return len == INT_MAX ? "" : s.substr(begin, len);
    }
};
// @lc code=end
