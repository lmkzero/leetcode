/*
 * @lc app=leetcode.cn id=704 lang=cpp
 *
 * [704] 二分查找
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int search(vector<int>& nums, int target) {
        if (nums.size() == 0) {
            return -1;
        }
        int begin = 0;
        int end = nums.size() - 1;
        for (; begin <= end;) {
            int mid = (begin + end) / 2;
            if (nums[mid] == target) {
                return mid;
            }
            if (nums[mid] > target) {
                end = mid - 1;
                continue;
            }
            begin = mid + 1;
        }
        return -1;
    }
};
// @lc code=end
