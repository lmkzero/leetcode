/*
 * @lc app=leetcode.cn id=12 lang=cpp
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
class Solution {
   public:
    string intToRoman(int num) {
        const int radix[] = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
        const string symbol[] = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
        string roman;
        for (size_t i = 0; num > 0; i++) {
            int count = num / radix[i];
            num = num % radix[i];
            for (; count > 0; count--) {
                roman = roman + symbol[i];
            }
        }
        return roman;
    }
};
// @lc code=end
