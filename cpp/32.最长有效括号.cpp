/*
 * @lc app=leetcode.cn id=32 lang=cpp
 *
 * [32] 最长有效括号
 */

// @lc code=start
class Solution {
public:
    int longestValidParentheses(string s) {
        int max_len=0;
        int last=-1;
        stack<int> lefts;
        for(int i=0;i<s.size();i++){
            if(s[i]=='('){
                lefts.push(i);
            }else{
                if(lefts.empty()){
                    last=i;
                }else{
                    lefts.pop();
                    if(lefts.empty()){
                        max_len=max(max_len,i-last);
                    }else{
                        max_len=max(max_len,i-lefts.top());
                    }
                }
            }
        }
        return max_len;
    }
};
// @lc code=end

