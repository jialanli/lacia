# lacia

## 简介
致力于打造快速、便捷的通用工具库
常用的功能直接开箱使用、标准库未直接提供的的功能直接开箱使用

-->欢迎提出问题或建议，作者会不断改进；欢迎Star

## 内容介绍

eg: 需要按多个分隔符切割字符串
只需go get -u "github.com/jialanli/lacia"

eg1：lacia.SplitByManyStr("ab+c*de+f/gh")  按默认的常用字符分割，返回分割好的[]string  

eg2：lacia.SplitByManyStrWith("ab+c*de+f/gh", []string{`*`, `+`, `/`}))   参数2为指定的要去除的字符集合, []string , 非常便捷。
按需使用即可~

各函数上均有明确说明，使用效果也可见对应文件的单元测试呦!

当前包含的方向:常用类型工具如string、日期、数组、各种转换等

一直在路上，随时补充中...期待新功能发布！