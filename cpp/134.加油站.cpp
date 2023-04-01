/*
 * @lc app=leetcode.cn id=134 lang=cpp
 *
 * [134] 加油站
 */

// @lc code=start
class Solution {
public:
    int canCompleteCircuit(vector<int>& gas, vector<int>& cost) {
        int len=gas.size();
        int left_gas=0;
        for(int i=0;i<len;i++){
            if(gas[i]<cost[i]){
                continue;
            }else{
                // 检查后续
                bool is_ok=true;
                left_gas=gas[i]-cost[i];
                for(int j=i+1;j<i+len;j++){
                    if(j<len){
                        left_gas=left_gas+gas[j]-cost[j];
                    }else{
                        left_gas=left_gas+gas[j-len]-cost[j-len];
                    }
                    if(left_gas<0){
                        is_ok=false;
                        continue;
                    }
                }
                if(is_ok==true){
                    return i;
                }
            }
        }
        return -1;
    }
};
// @lc code=end

