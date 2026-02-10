/*
 * @lc app=leetcode.cn id=31 lang=cpp
 *
 * [31] 下一个排列
 */

#include <algorithm>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    void nextPermutation(vector<int>& nums) {
        int length = nums.size();
        // From right to left, find the first digit(partitionNumber)
        // which violates the increase trend
        int p = length - 2;
        while (p > -1 && nums[p] >= nums[p + 1]) {
            p--;
        }
        // If not found, which means current sequence is already the largest
        // permutation, then rearrange to the first permutation and return false
        if (p == -1) {
            reverse(nums.begin(), nums.end());
            return;
        }
        // From right to left, find the first digit which is greater
        // than the partition number, call it changeNumber
        int c = length - 1;
        while (c > 0 && nums[c] <= nums[p]) {
            c--;
        }
        // Swap the partitionNumber and changeNumber
        swap(nums[p], nums[c]);
        // Reverse all the digits on the right of partitionNumber
        reverse(nums.begin() + p + 1, nums.end());
    }
};
// @lc code=end
