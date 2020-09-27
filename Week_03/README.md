学习笔记

### 递归要点
* 不要人肉递归
* 找到最近最简方法，将其拆解成可重复解决的问题（重复子问题）
* 数学归纳思维
### 递归代码模板
```cassandraql
func recursion(level, param1, param2,...){
    //recursion terminator
    if level > MAX_LEVEL{
        return
    }

    //process logic in current level
    process(level, data...)

    //drill down
    resusion(level+1, p1,...)

    //reverse the current level status if needed
}
```

