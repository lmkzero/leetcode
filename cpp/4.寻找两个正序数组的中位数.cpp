/*
 * @lc app=leetcode.cn id=4 lang=cpp
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int len1=nums1.size();
        int len2=nums2.size();
        if(len1==0){
            return findMedianSortedArrays(nums2);
        }else if(len2==0){
            return findMedianSortedArrays(nums1);
        }else{
            vector<int> nums;
            int i=0,j=0;
            while(i<len1&&j<len2){
                if(nums1[i]<nums2[j]){
                    nums.push_back(nums1[i]);
                    i++;
                }else{
                    nums.push_back(nums2[j]);
                    j++;
                }
            }
            while(i<len1){
                nums.push_back(nums1[i]);
                i++;
            }
            while(j<len2){
                nums.push_back(nums2[j]);
                j++;
            }
            return findMedianSortedArrays(nums);
        }
    }
    double findMedianSortedArrays(vector<int> &nums){
        int len=nums.size();
        if(len%2==0){
            return (nums[len/2-1]+nums[len/2])*1.0/2;
        }else{
            return nums[len/2]*1.0;
        }
    }
};
// @lc code=end

