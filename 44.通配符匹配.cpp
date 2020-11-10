/*
 * @lc app=leetcode.cn id=44 lang=cpp
 *
 * [44] 通配符匹配
 */

// @lc code=start
class Solution {
// 递归
// public:
//     bool isMatch(string s, string p) {
//         return isMatch(s.c_str(),p.c_str());
//     }
// private:
//     bool isMatch(const char *s,const char *p){
//         if(*p=='*'){
//             while(*p=='*'){
//                 p++;
//             }
//             if(*p=='\0'){
//                 return true;
//             }
//             while(*s!='\0'&&isMatch(s,p)){
//                 s++;
//             }
//             return *s!='\0';
//         }else if(*p=='\0'||*s=='\0'){
//             return *p==*s;
//         }else if(*p==*s||*p=='?'){
//             return isMatch(s++,p++);
//         }else{
//             return false;
//         }
//     }


// 迭代
public:
    bool isMatch(string s, string p) {
        return isMatch(s.c_str(),p.c_str());
    }
private:
    bool isMatch(const char *s,const char *p){
        bool star=false;
        const char *str,*ptr;
        for(str=s,ptr=p;*str!='\0';str++,ptr++){
            switch(*ptr){
                case '?':
                    break;
                case '*':
                    star=true;
                    s=str;
                    p=ptr;
                    while(*p=='*'){
                        p++;
                    }
                    if(*p=='\0'){
                        return true;
                    }
                    str=s-1;
                    ptr=p-1;
                    break;
                default:
                    if(*str!=*ptr){
                        if(!star){
                            return false;
                        }
                        s++;
                        str=s-1;
                        ptr=p-1;
                    }
            }
        }
        // str已至末尾，ptr末尾是*  
        while(*ptr=='*'){
            ptr++;
        }
        return *ptr=='\0';
    }
};
// @lc code=end

