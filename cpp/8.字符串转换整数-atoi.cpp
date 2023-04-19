/*
 * @lc app=leetcode.cn id=8 lang=cpp
 *
 * [8] 字符串转换整数 (atoi)
 */

#include <iostream>

using namespace std;

// @lc code=start
class Solution {
   public:
    int myAtoi(string str) {
        if (str.empty()) return 0;
        int i = 0;
        bool negative = false;
        long res = 0;
        while (str[i] == ' ') i++;
        if (str[i] == '-') {
            i++;
            negative = true;
        } else if (str[i] == '+') {
            i++;
        }
        for (; i < str.size(); i++) {
            if (str[i] >= '0' && str[i] <= '9') {
                res = res * 10 + (str[i] - '0');
                if (res > INT_MAX && !negative)
                    return INT_MAX;
                else if (res > INT_MAX && negative)
                    return INT_MIN;
            } else  // 遇到非法字符直接退出，停止解析
                break;
        }
        return negative ? -res : res;
    }
};
// @lc code=end
