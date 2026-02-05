/*
 * @lc app=leetcode.cn id=93 lang=cpp
 *
 * [93] 复原 IP 地址
 */

#include <string>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
   public:
    vector<string> restoreIpAddresses(string s) {
        dfs(s, 0);
        return ans;
    }

   private:
    vector<string> ans;
    vector<string> temp;

    void dfs(string s, int i) {
        int n = s.size();
        if (i >= n && temp.size() == 4) {
            ans.push_back(temp[0] + "." + temp[1] + "." + temp[2] + "." + temp[3]);
            return;
        }
        if (i >= n || temp.size() >= 4) {
            return;
        }
        int x = 0;
        for (int j = i; j < n && j < i + 3; j++) {
            x = x * 10 + s[j] - '0';
            if (x > 255 || (j > i && s[i] == '0')) {
                break;
            }
            temp.push_back(s.substr(i, j - i + 1));
            dfs(s, j + 1);
            temp.pop_back();
        }
    }
};
// @lc code=end
