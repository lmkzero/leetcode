/*
 * @lc app=leetcode.cn id=169 lang=cpp
 *
 * [169] 多数元素
 */

// @lc code=start
class Solution {
   public:
    int majorityElement(vector<int>& nums) {
        int res;
        unordered_map<int, int> map;
        for (auto i : nums) {
            if (map.find(i) == map.end()) {
                map.emplace(i, 1);
            } else {
                map.at(i) = map.at(i) + 1;
            }
        }
        for (auto i : map) {
            if (i.second > nums.size() / 2) {
                res = i.first;
            }
        }
        return res;
    }
};
// @lc code=end
