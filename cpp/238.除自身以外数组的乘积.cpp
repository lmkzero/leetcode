/*
 * @lc app=leetcode.cn id=238 lang=cpp
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        int len=nums.size();
        vector<int> ans(len);
        int k=1;
        for(int i=0;i<len;i++)
        {
            ans[i]=k;
            k*=nums[i];
        }
        k=1;
        for(int i=len-1;i>=0;i--)
        {
            ans[i]*=k;
            k*=nums[i];
        }
        return ans;
    }
};
// @lc code=end

