/*
 * @lc app=leetcode.cn id=135 lang=cpp
 *
 * [135] 分发糖果
 */

// @lc code=start
class Solution {
public:
    int candy(vector<int>& ratings) {
        int len=ratings.size();
        vector<int> candy(len,1);
        for(int i=1;i<len;i++){
            if(ratings[i]>ratings[i-1]){
                candy[i]=candy[i-1]+1;
            }
        }
        for(int j=len-1;j>0;j--){
            if(ratings[j-1]>ratings[j]){
                candy[j-1]=max(candy[j]+1,candy[j-1]);
            }
        }
        int sum=0;
        for(int k=0;k<len;k++){
            sum=sum+candy[k];
        }
        return sum;
    }
};
// @lc code=end

