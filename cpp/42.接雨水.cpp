/*
 * @lc app=leetcode.cn id=42 lang=cpp
 *
 * [42] 接雨水
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int trap(vector<int>& height) {
        int ans = 0;
        int left = 0, right = height.size() - 1;
        int leftMax = 0, rightMax = 0;
        while (left < right) {
            leftMax = max(leftMax, height[left]);
            rightMax = max(rightMax, height[right]);
            if (height[left] < height[right]) {
                ans += leftMax - height[left];
                ++left;
            } else {
                ans += rightMax - height[right];
                --right;
            }
        }
        return ans;
    }

    int trapDP(vector<int>& height) {
        const int n = height.size();
        int peak_index = 0;  // 最高的柱子，将数组分为两半
        for (int i = 0; i < n; i++) {
            if (height[i] > height[peak_index]) {
                peak_index = i;
            }
        }

        int water = 0;
        for (int i = 0, left_peak = 0; i < peak_index; i++) {
            if (height[i] > left_peak) {
                left_peak = height[i];
            } else {
                water += left_peak - height[i];
            }
        }
        for (int i = n - 1, right_peak = 0; i > peak_index; i--) {
            if (height[i] > right_peak) {
                right_peak = height[i];
            } else {
                water += right_peak - height[i];
            }
        }
        return water;
    }
};
// @lc code=end
