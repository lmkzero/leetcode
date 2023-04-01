/*
 * @lc app=leetcode.cn id=35 lang=cpp
 *
 * [35] 搜索插入位置
 */

// @lc code=start
class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int l = 0, r= nums.size()-1;
        while (l < r)
        {
            int mid = l + r >> 1;
            if (nums[mid] < target) l = mid + 1;
            else r  = mid;
        }
        if (l == nums.size()-1 && target > nums[l]) l++;
        return l;
    }
};
// @lc code=end

