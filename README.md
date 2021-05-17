# 静海
by jialanli 2021 持续更新中...

so surprise

致力于打造快速、便捷的通用工具库
常用高频功能、其它功能直接开箱使用

-->欢迎提出问题或建议，持续改进和补充

go get -u "github.com/jialanli/lacia"

or

cd ~/go/src/github.com
git clone https://github.com/jialanli/lacia.git

## 内容介绍

eg：需要按多个分隔符切割字符串：

方式1：    

    lacia.SplitByManyStr("ab+c*de+f/gh")

方式2：

    lacia.SplitByManyStrWith("ab+c*de+f/gh", []string{`*`, `+`, `/`}))   
参数2为指定的要去除的自定义字符集合, []string , 非常便捷,按需使用即可~

eg：获取参数2出现次数、判断参数1中是否含有参数2等：

    if lacia.ExistsInListInt([]int{0, 1, 2}, 5) == -1 {}

    ...多用途方式详见单元测试

各功能均有明确说明，使用效果也可见对应文件的单元测试呦!


持续改进和更新！
作者的博客：https://lan6193.blog.csdn.net/  欢迎围观、感谢支持

一直在路上，随时补充中...期待新功能发布！