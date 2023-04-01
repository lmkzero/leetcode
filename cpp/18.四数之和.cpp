/*
 * @lc app=leetcode.cn id=18 lang=cpp
 *
 * [18] 四数之和
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        vector<vector<int>> result;
        if (nums.size() < 4)
            return result;
        sort(nums.begin(), nums.end());

        unordered_map<int, vector<pair<int, int>>> cache;
        for (size_t a = 0; a < nums.size(); a++) {
            for (size_t b = a + 1; b < nums.size(); b++) {
                //cache存储形式：（sum,[<value,value>,<value,value>]）
                cache[nums[a] + nums[b]].push_back(pair<int, int>(a, b));
            }
        }

        for (int c = 0; c < nums.size(); c++) {
            for (size_t d = c + 1; d < nums.size(); d++) {
                const int key = target - nums[c] - nums[d];
                if (cache.find(key) == cache.end())
                    continue;
                //获取vector，其中pair元素能和当前pair元素和等于target
                const auto &vec = cache[key];
                for (size_t k = 0; k < vec.size(); k++) {
                    //<nums[c],num[d]>和vec[k]进行比较
                    //排序数组，只需要比较大小即可确定位置
                    if (c <= vec[k].second)
                        continue;
                    result.push_back({ nums[vec[k].first],nums[vec[k].second],nums[c],nums[d] });
                }
            }
        }
        // 去除重复
        // https://devdocs.io/cpp/algorithm/unique
        sort(result.begin(), result.end());
        result.erase(unique(result.begin(), result.end()), result.end());
        return result;
    }
};
// @lc code=end

