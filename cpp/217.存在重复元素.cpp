/*
 * @lc app=leetcode.cn id=217 lang=cpp
 *
 * [217] 存在重复元素
 */

// @lc code=start
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        unordered_map<int,int> map;
        for(auto i:nums){
            if(map.find(i)==map.end()){
                map.emplace(i,1);
            }else{
                return true;
            }
        }
        return false;
    }
};
// @lc code=end

