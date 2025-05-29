/*
 * @lc app=leetcode.cn id=239 lang=cpp
 *
 * [239] 滑动窗口最大值
 */

#include <queue>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        priority_queue<pair<int, int>> q;
        int n = nums.size();
        for (int i = 0; i < k - 1; i++) {
            q.push({nums[i], -i});
        }
        vector<int> ans;
        for (int i = k - 1; i < n; i++) {
            q.push({nums[i], -i});
            while (-q.top().second <= i - k) {
                q.pop();
            }
            ans.emplace_back(q.top().first);
        }
        return ans;
    }
};
// @lc code=end
