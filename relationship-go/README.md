# 简介

这是中国亲戚关系计算器的golang实现版，用于家庭称谓、亲戚称呼、辈分关系的计算。

原始版本为js代码实现，查看更多信息、了解具体算法原理等，请查看原仓库说明 —— [https://github.com/mumuy/relationship](https://github.com/mumuy/relationship)

# 安装

```shell
go get github.com/young2j/relationship/relationship-go@latest
```

# 使用

1. 通用方法: 唯一的计算方法 `Relationship`.

   * 选项模式 `Relationship(options map[string]any | options.ParameterOptions | *options.ParameterOptions)`

   参数 `options`结构为：

   ```golang
   options := map[string]any{
    "text":    "",        // 目标对象：目标对象的称谓汉字表达，称谓间用‘的’字分隔
    "target":  "",        // 相对对象：相对对象的称谓汉字表达，称谓间用‘的’字分隔，空表示自己
    "sex":     -1,        // 本人性别：0表示女性,1表示男性
    "type":    "default", // 转换类型：'default'计算称谓,'chain'计算关系链,'pair'计算关系合称
    "reverse": false,     // 称呼方式：True对方称呼我,False我称呼对方
    "mode":    "default", // 模式选择：使用set_mode方法定制不同地区模式，在此选择自定义模式
    "optimal": false,     // 最短关系：计算两者之间的最短关系
   }

   ```

   代码示例：

   ```golang
    import (
        "github.com/young2j/relationship/relationship-go"
        "github.com/young2j/relationship/relationship-go/options"
    )

    func main() {
        rel := relationship.NewRelationship()

        // 如：我应该叫外婆的哥哥什么？
        rel.Relationship(options.ParameterOptions{
            Text: "妈妈的妈妈的哥哥",
        })
        // => [舅外公]

        // 如：七舅姥爷应该叫我什么？
        rel.Relationship(options.ParameterOptions{
            Text:    "七舅姥爷",
            Reverse: true,
            Sex:     1,
        })
        // => [甥外孙]

        // 如：舅公是什么亲戚
        rel.Relationship(options.ParameterOptions{
            Text: "舅公",
            Type: options.OptionTypeChain,
        })
        // => [妈妈的妈妈的兄弟 老公的妈妈的兄弟 爸爸的妈妈的兄弟]

        // 如：舅妈如何称呼外婆？
        rel.Relationship(options.ParameterOptions{
            Text:   "外婆",
            Target: "舅妈",
            Sex:    1,
        })
        // => [婆婆]

        // 如：外婆和奶奶之间是什么关系？
        rel.Relationship(options.ParameterOptions{
            Text:   "外婆",
            Target: "奶奶",
            Type:   options.OptionTypePair,
        })
        // => [儿女亲家]
    }
   ```

* 语句模式 `Relationship(exptession string)`

  参数 `exptession`句式可以为：`xxx是xxx的什么人`、`xxx叫xxx什么`、`xxx如何称呼xxx`等.
  代码示例：

  ```golang
  // 如：舅妈如何称呼外婆？
  rel.Relationship("舅妈如何称呼外婆？")
  // => [婆婆]

  // 如：外婆和奶奶之间是什么关系？
  rel.Relationship("外婆和奶奶之间是什么关系？")
  // => [儿女亲家]
  ```

2. 内部方法：获取当前数据表 `rel.GetData()`, 获取模式数据 `rel.GetModeData(mode options.OptionMode)`.
3. 内部方法：获取当前数据量 `rel.DataCount()`.
4. 内部方法：

   * 用户自定义模式 `rel.SetMode(mode options.OptionMode)`
   * 用户自定义模式+数据 `rel.SetModeWithExtraData(mode options.OptionMode, extraData map[string][]string)`.

   代码示例(可参考数据表格式对数据进行覆盖)：

   ```golang
   // ...
   rel.SetMode(options.OptionModeGuangDong)
   // ...
   ```
   ```golang
   // 关系解析语法
   // 【关系链】f:父,m:母,h:夫,w:妻,s:子,d:女,xb:兄弟,ob:兄,lb:弟,xs:姐妹,os:姐,ls:妹
   // 【修饰符】 1:男性,0:女性,&o:年长,&l:年幼,#:隔断,[a|b]:并列
   rel.SetModeWithExtraData("northern", map[string][]string{
    "m,f":      {"姥爷"},
    "m,m":      {"姥姥"},
    "m,xb,s&o": {"表哥"},
    "m,xb,s&l": {"表弟"},
   })
   ```
