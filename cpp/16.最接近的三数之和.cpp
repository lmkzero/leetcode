/*
 * @lc app=leetcode.cn id=16 lang=cpp
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
        int result = 0;
        int min_gap = INT_MAX;

        sort(nums.begin(), nums.end());

        for (auto a = nums.begin(); a != prev(nums.end(), 2); a++) {
            auto b = next(a);
            auto c = prev(nums.end());
            while (b < c) {
                const int sum = *a + *b + *c;
                const int gap = abs(sum - target);
                if (gap < min_gap) {
                    result = sum;
                    min_gap = gap;
                }
                if (sum < target) {
                    b++;
                }
                else {
                    c--;
                }
            }
        }
        return result;
    }
};
// @lc code=end

