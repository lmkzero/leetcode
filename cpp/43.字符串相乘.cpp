/*
 * @lc app=leetcode.cn id=43 lang=cpp
 *
 * [43] 字符串相乘
 */

// @lc code=start
class Solution {
public:
    string multiply(string num1, string num2) {
        if(num1=="0"||num2=="0"){
            return "0";
        }
        reverse(num1.begin(),num1.end());
        reverse(num2.begin(),num2.end());
        int len1=num1.size();
        int len2=num2.size();
        vector<int> ans(len1+len2,0);
        for(int i=0;i<len1;i++){
            for(int j=0;j<len2;j++){
                ans[i+j]+=(num1[i]-'0')*(num2[j]-'0');
            }
        }
        for(int i=0;i<len1+len2;i++){
            if(ans[i]>9){
                int t=ans[i];
                ans[i]=t%10;
                ans[i+1]+=(t/10);
            }
        }
        int pos=(ans[len1+len2-1]==0?len1+len2-2:len1+len2-1);
        string aa="";
        for(;pos>=0;pos--){
            aa+=(ans[pos]+'0');
        }
        return aa;
    }
};
// @lc code=end

