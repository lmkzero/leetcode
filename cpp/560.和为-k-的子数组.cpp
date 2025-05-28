/*
 * @lc app=leetcode.cn id=560 lang=cpp
 *
 * [560] 和为 K 的子数组
 */

#include <unordered_map>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int subarraySum(vector<int>& nums, int k) {
        unordered_map<int, int> cnt{{0, 1}};
        int ans = 0;
        int sum = 0;
        for (int num : nums) {
            sum += num;
            ans += cnt[sum - k];
            cnt[sum]++;
        }
        return ans;
    }
};
// @lc code=end
