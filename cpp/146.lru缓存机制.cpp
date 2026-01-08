/*
 * @lc app=leetcode.cn id=146 lang=cpp
 *
 * [146] LRU缓存机制
 */

#include <list>

using namespace std;

// @lc code=start
class LRUCache {
   private:
    list<pair<int, int>> l;
    unordered_map<int, list<pair<int, int>>::iterator> mp;
    int cap;

   public:
    LRUCache(int capacity) { cap = capacity; }

    int get(int key) {
        if (mp.find(key) != mp.end()) {
            int res = (*mp[key]).second;
            l.erase(mp[key]);
            l.push_front(make_pair(key, res));
            mp[key] = l.begin();
            return res;
        }
        return -1;
    }

    void put(int key, int value) {
        if (mp.find(key) != mp.end()) {
            l.erase(mp[key]);
            l.push_front(make_pair(key, value));
            mp[key] = l.begin();
        } else {
            if (l.size() == cap) {
                mp.erase(l.back().first);
                l.pop_back();
            }
            l.push_front(make_pair(key, value));
            mp[key] = l.begin();
        }
    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
// @lc code=end
