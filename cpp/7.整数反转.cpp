/*
 * @lc app=leetcode.cn id=7 lang=cpp
 *
 * [7] 整数反转
 */

#include <iostream>
#include <algorithm>

using namespace std;

// @lc code=start
class Solution {
   public:
    int reverse(int x) {
        bool is_negative = false;
        if (x < 0) {
            is_negative = true;
        }
        string num = to_string(abs(x));
        std::reverse(num.begin(), num.end());
        cout << num << endl;
        long result = is_negative ? -stol(num) : stol(num);
        if (result > INT_MAX || result < INT_MIN) {
            return 0;
        } else {
            return int(result);
        }

        // 正常解法-
        // int rev = 0;
        // while (x != 0) {
        //     int pop = x % 10;
        //     x /= 10;
        //     if (rev > INT_MAX/10 || (rev == INT_MAX / 10 && pop > 7)) return 0;
        //     if (rev < INT_MIN/10 || (rev == INT_MIN / 10 && pop < -8)) return 0;
        //     rev = rev * 10 + pop;
        // }
        // return rev;
    }
};
// @lc code=end
