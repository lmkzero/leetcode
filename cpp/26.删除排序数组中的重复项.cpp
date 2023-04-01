/*
 * @lc app=leetcode.cn id=26 lang=cpp
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int length=nums.size();
        if(length<=0){
            return 0;
        }
        int len=0;
        int index=len+1;
        while(index<length){
            if(nums[index]==nums[len]){
                index++;
            }else{
                len++;
                nums[len]=nums[index];
                index++;
            }
        }
        return len+1;
    }
};
// @lc code=end

