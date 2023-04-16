/*
 * @lc app=leetcode.cn id=205 lang=cpp
 *
 * [205] 同构字符串
 */

// @lc code=start
class Solution {
   public:
    bool isIsomorphic(string s, string t) {
        if (s.size() != t.size()) {
            return false;
        }
        if (s.size() == 0 && t.size() == 0) {
            return true;
        }
        // 双map，正反映射
        // "ab" "aa"
        // a->a，单map在检查第二个字符b时会直接将b->a加入map，忽略了正反映射的一致性
        unordered_map<char, char> smap;
        unordered_map<char, char> tmap;
        for (int i = 0; i < s.size(); i++) {
            char sc = s[i];
            char tc = t[i];
            if (smap.count(sc)) {
                if (smap[sc] != tc) {
                    return false;
                }
            } else if (tmap.count(tc)) {
                if (tmap[tc] != sc) {
                    return false;
                }
            } else {
                smap[sc] = tc;
                tmap[tc] = sc;
            }
        }
        return true;
    }
};
// @lc code=end
