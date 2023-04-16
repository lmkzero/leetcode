/*
 * @lc app=leetcode.cn id=219 lang=cpp
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
class Solution {
   public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        unordered_map<int, int> map;
        int len = nums.size();
        for (int i = 0; i < len; i++) {
            if (map.find(nums[i]) == map.end()) {
                map.emplace(nums[i], i);
            } else {
                int gap = i - map.at(nums[i]);
                if (abs(gap) <= k) {
                    return true;
                } else {
                    map.at(nums[i]) = i;
                }
            }
        }
        return false;
    }
};
// @lc code=end
