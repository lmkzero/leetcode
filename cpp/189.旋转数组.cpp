/*
 * @lc app=leetcode.cn id=189 lang=cpp
 *
 * [189] 旋转数组
 */

// @lc code=start
class Solution {
   public:
    void rotate(vector<int>& nums, int k) {
        // 超时
        // int len=nums.size();
        // for(int i=0;i<k;i++){
        //     int last=nums[len-1];
        //     for(int j=len-1;j>0;j--){
        //         nums[j]=nums[j-1];
        //     }
        //     nums[0]=last;
        // }

        k = k % nums.size();
        reverse(nums.begin(), nums.begin() + nums.size() - k);
        reverse(nums.begin() + nums.size() - k, nums.end());
        reverse(nums.begin(), nums.end());
    }
};
// @lc code=end
