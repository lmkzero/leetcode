/*
 * @lc app=leetcode.cn id=215 lang=cpp
 *
 * [215] 数组中的第K个最大元素
 */

#include <stdlib.h>
#include <queue>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int findKthLargest(vector<int>& nums, int k) {
        priority_queue<int, vector<int>, greater<int>> minQ;
        for (int x : nums) {
            minQ.push(x);
            if (minQ.size() > k) {
                minQ.pop();
            }
        }
        return minQ.top();
    }
};
// @lc code=end
