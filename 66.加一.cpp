/*
 * @lc app=leetcode.cn id=66 lang=cpp
 *
 * [66] 加一
 */

// @lc code=start
class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        int len=digits.size();
        int c=0;
        int i;
        for(i=len-1;i>=0;i--){
            digits[i]=digits[i]+1;
            if(digits[i]>9){
                c=digits[i]/10;
                digits[i]=digits[i]%10;
            }
            else
                break;
        }
        if(c>0&&i==-1){
            digits.insert(digits.begin(),1);
        }
        return digits;
    }
};
// @lc code=end

