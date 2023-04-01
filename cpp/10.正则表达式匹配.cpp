/*
 * @lc app=leetcode.cn id=10 lang=cpp
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
class Solution {
public:
    bool isMatch(string s, string p) {
        return isMatch(s.c_str(),p.c_str());
    }
private:
    bool isMatch(const char *s,const char *p){
        if(*p=='\0'){
            return *s=='\0';
        }
        if(*(p+1)!='*'){
            if(*p==*s||(*p=='.'&&*s!='\0')){
                return isMatch(s+1,p+1);
            }else{
                return false;
            }
        }else{
            while(*p==*s||(*p=='.'&&*s!='\0')){
                if(isMatch(s,p+2)){
                    return true;
                }
                s++;
            }
            return isMatch(s,p+2);
        }
    }
};
// @lc code=end

