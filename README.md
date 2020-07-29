# ExpressionParser

显示的标示变量类型 str($var1)  int($var2)

处理过程

1 编译rule

    1.1 检查语法
    
    1.2 词法处理
    
    1.3 逆波兰 -> 得到中间代码 ，变量表
    

2 带入变量

    2.1 校验变量类型和完备性
    
    2.2 带入逆波兰公式得到结果
    


公式示例：

* num{$age} > 20 && num{$age} < 100

* num{$var1} - num{$var2} == 30

* str{$name} == 'Mary' && ( num{$age} > 20 && num{$age} < 100 ) || num{$number} == 1234567890

