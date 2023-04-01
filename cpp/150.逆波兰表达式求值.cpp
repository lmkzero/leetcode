/*
 * @lc app=leetcode.cn id=150 lang=cpp
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        int x,y;
        auto token=tokens.back();
        tokens.pop_back();
        if(is_opterator(token)){
            y=evalRPN(tokens);
            x=evalRPN(tokens);
            if(token[0]=='+'){
                x+=y;
            }else if(token[0]=='-'){
                x-=y;
            }else if(token[0]=='*'){
                x*=y;
            }else{
                x/=y;
            }
        }else{
            size_t i;
            x=stoi(token,&i);
        }
        return x;
    }
private:
    bool is_opterator(const string &op){
        return op.size()==1&&string("+-*/").find(op)!=string::npos;
    }
};
// @lc code=end

