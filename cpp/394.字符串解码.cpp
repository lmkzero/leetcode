/*
 * @lc app=leetcode.cn id=394 lang=cpp
 *
 * [394] 字符串解码
 */

#include <stack>
#include <string>

using namespace std;

// @lc code=start
class Solution {
   public:
    string decodeString(string s) {
        stack<int> intQ;
        stack<string> stringQ;
        int num = 0;
        string res = "";
        for (char c : s) {
            if (c >= '0' && c <= '9') {
                num = num * 10 + int(c - '0');
                continue;
            }
            if (c == '[') {
                intQ.push(num);
                stringQ.push(res);
                num = 0;
                res = "";
                continue;
            }
            if (c == ']') {
                int n = intQ.top();
                string t = "";
                for (int i = 0; i < n; i++) {
                    t += res;
                }
                res = stringQ.top() + t;
                intQ.pop();
                stringQ.pop();
                continue;
            }
            res += c;
        }
        return res;
    }
};
// @lc code=end
