/*
 * @lc app=leetcode.cn id=214 lang=cpp
 *
 * [214] 最短回文串
 */

// @lc code=start
class Solution {
   public:
    // 暴力算法
    // 超时
    // string shortestPalindrome(string s) {
    //     int n=s.size();
    //     if(n==1){
    //         return s;
    //     }
    //     string rev(s);
    //     reverse(rev.begin(),rev.end());
    //     int j=0;
    //     for(int i=0;i<n;i++){
    //         if(s.substr(0,n-i)==rev.substr(i)){
    //             return rev.substr(0,i)+s;
    //         }
    //     }
    //     return "";
    // }

    string shortestPalindrome(string s) {
        int n = s.size();
        int i = 0;
        for (int j = n - 1; j >= 0; j--) {
            if (s[i] == s[j]) i++;
        }
        if (i == n) return s;
        string remain_rev = s.substr(i, n);
        reverse(remain_rev.begin(), remain_rev.end());
        return remain_rev + shortestPalindrome(s.substr(0, i)) + s.substr(i);
    }
};
// @lc code=end
