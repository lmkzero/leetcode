/*
 * @lc app=leetcode.cn id=509 lang=cpp
 *
 * [509] 斐波那契数
 */

// @lc code=start
class Solution {
   public:
    // 递归
    // int fib(int N) {
    //     if(N==0||N==1){
    //         return N;
    //     }
    //     return fib(N-1)+fib(N-2);
    // }

    // 备忘录法
    // int fib(int N) {
    //     if(N==0||N==1){
    //         return N;
    //     }
    //     vector<int> memo(N+1,0);
    //     return helper(memo,N);
    // }
    // int helper(vector<int> &memo,int N){
    //     if(N==0||N==1){
    //         return N;
    //     }
    //     if(memo[N]!=0){
    //         return memo[N];
    //     }
    //     return helper(memo,N-1)+helper(memo,N-2);
    // }

    // DP table
    // int fib(int N){
    //     vector<int> dp(N+2,0);
    //     dp[0]=0;
    //     dp[1]=1;
    //     for(int i=2;i<=N;i++){
    //         dp[i]=dp[i-1]+dp[i-2];
    //     }
    //     return dp[N];
    // }

    // DP
    // 空间复杂度O(1)
    int fib(int N) {
        if (N == 0 || N == 1) {
            return N;
        }
        int pre = 0;
        int cur = 1;
        for (int i = 2; i <= N; i++) {
            int sum = pre + cur;
            pre = cur;
            cur = sum;
        }
        return cur;
    }
};
// @lc code=end
