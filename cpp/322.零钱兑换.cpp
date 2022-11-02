/*
 * @lc app=leetcode.cn id=322 lang=cpp
 *
 * [322] 零钱兑换
 */

// @lc code=start
class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        // dp
        // 自底向上
        vector<int> dp(amount+1,amount+1);
        dp[0]=0;
        for(int i=0;i<dp.size();i++){
            for(int coin:coins){
                if(i-coin<0){
                    continue;
                }else{
                    dp[i]=min(dp[i],dp[i-coin]+1);
                }
            }
        }
        if(dp[amount]==amount+1){
            return -1;
        }else{
            return dp[amount];
        }

        // vector<int> memo(amount+1,0);
        // return dp(coins,amount,memo);
    }


    
    // 依然超时
    // int dp(vector<int> &coins, int amount,vector<int> &memo){
    //     if(amount==0){
    //         return 0;
    //     }
    //     if(amount<0){
    //         return -1;
    //     }
    //     if(memo[amount]!=0){
    //         return memo[amount];
    //     }
    //     int res=INT_MAX;
    //     for(int coin:coins){
    //         int sub=dp(coins,amount-coin,memo);
    //         if(sub==-1){
    //             continue;
    //         }
    //         res=min(res,sub+1);
    //     }
    //     if(res!=INT_MAX){
    //         memo[amount]=res;
    //         return res;
    //     }else{
    //         return -1;
    //     }
    // }

    // 暴力递归
    // 超时
    // int dp(vector<int> &coins,int amount){
    //     if(amount==0){
    //         return 0;
    //     }
    //     if(amount<0){
    //         return -1;
    //     }
    //     int res=INT_MAX;
    //     for(int coin:coins){
    //         int sub=dp(coins,amount-coin);
    //         if(sub==-1){
    //             continue;
    //         }
    //         res=min(res,sub+1);
    //     }
    //     if(res!=INT_MAX){
    //         return res;
    //     }else{
    //         return -1;
    //     }
    // }
};
// @lc code=end

