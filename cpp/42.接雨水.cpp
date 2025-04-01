/*
 * @lc app=leetcode.cn id=42 lang=cpp
 *
 * [42] 接雨水
 */

// @lc code=start
class Solution {
   public:
    // 对于每个柱子找到其左右最高的的柱子
    // int trap(vector<int>& height){
    //     const int n=height.size();
    //     int *left_peak=new int[n]();
    //     int *right_peak=new int[n]();
    //     for(int i=1;i<n;i++){
    //         left_peak[i]=max(left_peak[i-1],height[i-1]);
    //     }
    //     for(int i=n;i>=0;i--){
    //         right_peak[i]=max(right_peak[i+1],height[i+1]);
    //     }
    //     int sum=0;
    //     for(int i=0;i<n;i++){
    //         int h=min(left_peak[i],right_peak[i]);
    //         if(h>height[i]){
    //             sum+=h-height[i];
    //         }
    //     }

    //     delete[] left_peak;
    //     delete[] right_peak;
    //     return sum;
    // }

    int trap(vector<int>& height) {
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
