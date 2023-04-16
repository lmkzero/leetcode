/*
 * @lc app=leetcode.cn id=69 lang=cpp
 *
 * [69] x 的平方根
 */

// @lc code=start
class Solution {
   public:
    int mySqrt(int x) {
        // i*i溢出
        /* int res=0;
        for(int i=0;i<=x;i++){
            if(i*i==x){
                res=i;
                break;
            }else if(i*i<x&&(i+1)*(i+1)>x){
                res=i;
                break;
            }
        }
        return res; */
        if (x == 1) return 1;
        int min = 0;
        int max = x;
        while (max - min > 1) {
            int m = (max + min) / 2;
            if (x / m < m)
                max = m;
            else
                min = m;
        }
        return min;
    }
};
// @lc code=end
