/*
 * @lc app=leetcode.cn id=41 lang=cpp
 *
 * [41] 缺失的第一个正数
 */

#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int firstMissingPositive(vector<int>& nums) {
        bucket_sort(nums);
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] != (i + 1)) {
                return i + 1;
            }
        }
        return nums.size() + 1;
    }

   private:
    void bucket_sort(vector<int>& A) {
        int n = A.size();
        for (int i = 0; i < n; i++) {
            while (A[i] != i + 1) {
                if (A[i] <= 0 || A[i] > n || A[i] == A[A[i] - 1]) {
                    break;
                }
                swap(A[i], A[A[i] - 1]);
            }
        }
    }
};
// @lc code=end
