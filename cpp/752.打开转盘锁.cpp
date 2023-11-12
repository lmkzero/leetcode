/*
 * @lc app=leetcode.cn id=752 lang=cpp
 *
 * [752] 打开转盘锁
 */

#include <queue>
#include <string>
#include <unordered_set>

using namespace std;

// @lc code=start
class Solution {
   public:
    int openLock(vector<string>& deadends, string target) {
        unordered_set<string> deads;
        for (string s : deadends) {
            deads.insert(s);
        }
        unordered_set<string> visited;
        queue<string> q;
        int step = 0;
        q.push("0000");
        visited.insert("0000");
        while (!q.empty()) {
            int sz = q.size();
            for (int i = 0; i < sz; i++) {
                string cur = q.front();
                q.pop();
                if (deads.count(cur)) {
                    continue;
                }
                if (cur == target) {
                    return step;
                }
                for (int j = 0; j < 4; j++) {
                    string up = cur;
                    up[j] = plusOne(up[j]);     // 写时拷贝
                    if (!visited.count(up)) {
                        q.push(up);
                        visited.insert(up);
                    }
                    string down = cur;
                    down[j] = minusOne(down[j]);
                    if (!visited.count(down)) {
                        q.push(down);
                        visited.insert(down);
                    }
                }
            }
            step++;
        }
        return -1;
    }

    // 将 s[j] 向上拨动一次
    char plusOne(char x) {
       return  (x == '9' ? '0' : x + 1);
    }

    // 将 s[i] 向下拨动一次
    char minusOne(char x) {
        return (x == '0' ? '9' : x - 1);
    }
};
// @lc code=end
