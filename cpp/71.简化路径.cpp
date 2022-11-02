/*
 * @lc app=leetcode.cn id=71 lang=cpp
 *
 * [71] 简化路径
 */

// @lc code=start
class Solution {
public:
    string simplifyPath(string path) {
        vector<string> dirs;
        for(auto i=path.begin();i!=path.end();){
            i++;
            auto j=find(i,path.end(),'/');
            auto dir=string(i,j);
            if(!dir.empty()&&dir!="."){
                if(dir==".."){
                    if(!dirs.empty()){
                        dirs.pop_back();
                    }
                }else{
                        dirs.push_back(dir);
                }
            }
            i=j;
        }
        stringstream out;
        if(dirs.empty()){
            out<<"/";
        }else{
            for(auto dir:dirs){
                out<<'/'<<dir;
            }
        }
        return out.str();
    }
};
// @lc code=end

