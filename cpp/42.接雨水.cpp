/*
 * @lc app=leetcode.cn id=42 lang=cpp
 *
 * [42] 接雨水
 */

// @lc code=start
class Solution {
   public:
    // 双层遍历(有问题)
    //  int trap(vector<int>& height) {
    //      if(height.size()<2){
    //          return 0;
    //      }
    //      int left,right;
    //      int water=0;
    //      for(left=1,right=2;left<height.size()-2;left++){
    //          if(height[left]>height[right]){
    //              right++;
    //          }else{
    //              water=height[left]*(right-left-1)+water;
    //              for(int i=left+1;i<right;i++){
    //                  water=water-height[i];
    //              }
    //              left=right;
    //          }
    //      }
    //      return water;
    //  }

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
