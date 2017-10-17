# Selpg
## 提示信息
    USAGE: slepg -sstart_page -eend_page [-f|-llines_per_page] [-ddest] [in_filename]
### 测试(初始状态下每页4行)
    测试1: ./selpg -s1 -e1 test.txt
    输出：
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!

------------------------------------------------
    测试2：./selpg -s1 -e1 < test.txt
    输出：
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!

------------------------------------------------
    测试3：./selpg -s2 -e4 test.txt>output.txt
    结果： 写入output.txt
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    
------------------------------------------------
    测试4：./selpg -s4 -e4 test.txt>error.txt
    输出：
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    没有错误所以error.txt为空
    
------------------------------------------------
    测试5：./slepg -s5 -e5 test.txt>output.txt>error.txt
    结果：写入output.txt
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    hello,pan sir!
    error.txt为空
    
------------------------------------------------
    测试6：./slepg -s6 -e6 -l1 test.txt
    输出：
    hello,pan sir!
    
------------------------------------------------
    测试7：./slepg -s1 -e1 -f test2.txt
    输出：
    hello,pan sir!
    

# slepg
