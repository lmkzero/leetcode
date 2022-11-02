/*
 * @lc app=leetcode.cn id=88 lang=cpp
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        int i1=m-1;
        int i2=n-1;
        int cur=m+n-1;
        while(i1>=0&&i2>=0){
            nums1[cur--]=nums1[i1]>=nums2[i2]?nums1[i1--]:nums2[i2--];
        }
        while(i2>=0){
            nums1[cur--]=nums2[i2--];
        }
    }
};
// @lc code=end

