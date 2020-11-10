/*
 * @lc app=leetcode.cn id=49 lang=cpp
 *
 * [49] 字母异位词分组
 */

// @lc code=start
class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        unordered_map<string,vector<string>> group;
        for(const auto &s:strs){
            string key=s;
            sort(key.begin(),key.end());
            group[key].push_back(s);
        }
        vector<vector<string>> result;
        for(auto it=group.cbegin();it!=group.cend();it++){
            result.insert(result.end(),it->second);
        }
        return result;
    }
};
// @lc code=end

