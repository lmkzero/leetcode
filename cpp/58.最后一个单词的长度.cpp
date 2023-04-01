/*
 * @lc app=leetcode.cn id=58 lang=cpp
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
class Solution {
public:
    int lengthOfLastWord(string s) {
        auto first=find_if(s.rbegin(),s.rend(),::isalpha);
        auto last=find_if_not(first,s.rend(),::isalpha);
        return distance(first,last);
    }
};
// @lc code=end

