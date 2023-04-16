/*
 * @lc app=leetcode.cn id=34 lang=cpp
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
class Solution {
   public:
    vector<int> searchRange(vector<int>& nums, int target) {
        if (nums.empty()) return {-1, -1};
        int first = findFirst(nums, target);
        int last = findLast(nums, target);
        return {first, last};
    }
    int findFirst(vector<int>& nums, int target) {
        int f = 0, l = nums.size() - 1;
        while (f <= l) {
            int m = (f + l) / 2;
            if (nums[m] < target)
                f = m + 1;
            else
                l = m - 1;
        }
        if (f < nums.size() && nums[f] == target) return f;
        return -1;
    }
    int findLast(vector<int>& nums, int target) {
        int f = 0, l = nums.size() - 1;
        while (f <= l) {
            int m = (f + l) / 2;
            if (nums[m] <= target)
                f = m + 1;
            else
                l = m - 1;
        }
        if (l >= 0 && nums[l] == target) return l;
        return -1;
    }
};
// @lc code=end
