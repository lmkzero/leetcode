/*
 * @lc app=leetcode.cn id=128 lang=cpp
 *
 * [128] 最长连续序列
 */

#include <unordered_set>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    int longestConsecutive(vector<int>& nums) {
        unordered_set<int> my_set;
        for (auto i : nums) my_set.insert(i);
        int longest = 0;
        for (auto i : nums) {
            int length = 1;
            for (int j = i - 1; my_set.find(j) != my_set.end(); j--) {
                // 针对数组中的一个连续序列，这里一定能完整获取到全部元素
                // 所以，这里直接对hash表中的元素进行删除
                my_set.erase(j);
                length++;
            }
            for (int j = i + 1; my_set.find(j) != my_set.end(); j++) {
                my_set.erase(j);
                length++;
            }
            longest = longest >= length ? longest : length;
        }
        return longest;
    }
};
// @lc code=end
