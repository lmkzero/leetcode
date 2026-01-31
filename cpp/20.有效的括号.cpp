/*
 * @lc app=leetcode.cn id=20 lang=cpp
 *
 * [20] 有效的括号
 */

#include <stack>
#include <unordered_map>

using namespace std;

// @lc code=start
class Solution {
   private:
    unordered_map<char, char> pairs = {{')', '('}, {']', '['}, {'}', '{'}};

   public:
    bool isValid(string s) {
        if (s.length() % 2 != 0) {
            return false;
        }
        stack<char> stack;
        for (auto c : s) {
            if (pairs.count(c) == 0) {
                stack.push(c);
                continue;
            }
            if (stack.empty() || stack.top() != pairs.at(c)) {
                return false;
            }
            stack.pop();
        }
        return stack.empty();
    }
};
// @lc code=end
