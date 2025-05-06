/*
* 关系链缩写规则 - 将【关系链】缩减成最简表达
*
* 关系数据语法说明：
* 【关系符】 f:父; m:母; h:夫; w:妻; s:子; d:女; xb:兄弟; ob:兄; lb:弟; xs:姐妹; os:姐; ls:妹
* 【修饰符】 1:男性; 0:女性; &o:年长; &l:年幼; &\d+:排序 //:隔断; [a|b]:并列
 */

package rule

import "regexp"

type FilterItem struct {
	Exp *regexp.Regexp
	Str string
}

var Filter = []FilterItem{
	// 年龄
	{ // 有大有小，增加不确定性
		Exp: regexp.MustCompile(`^(.+)&o([^#]+)&l`),
		Str: "$1$2",
	},
	{ // 有大有小，增加不确定性
		Exp: regexp.MustCompile(`^(.+)&l([^#]+)&o`),
		Str: "$1$2",
	},
	{ // 有大有小，增加不确定性
		Exp: regexp.MustCompile(`^(.+)?,o([sb](&\d+)?)([^#]+)&l`),
		Str: "$1,x$2$4",
	},
	{ // 有大有小，增加不确定性
		Exp: regexp.MustCompile(`^(.+)?,l([sb](&\d+)?)([^#]+)&o`),
		Str: "$1,x$2$4",
	},
	// 兄弟姐妹
	{ // 哥哥姐姐的哥哥姐姐还是自己的哥哥姐姐(年龄判断)
		Exp: regexp.MustCompile(`(,o[sb](&\d+)?)+(,o[sb])(&\d+)?`),
		Str: "$3$4",
	},
	{ // 弟弟妹妹的弟弟妹妹还是自己的弟弟妹妹(年龄判断)
		Exp: regexp.MustCompile(`(,l[sb](&\d+)?)+(,l[sb])(&\d+)?`),
		Str: "$3$4",
	},
	{ // 如果自己是男性,兄弟姐妹的兄弟就是自己的兄弟或自己
		Exp: regexp.MustCompile(`^(.*,[fh1])(,[olx][sb](&\d+)?)+,[olx]b(.*)$`),
		Str: "$1,xb$4#$1$4",
	},
	{ // 如果自己是女性,兄弟姐妹的姐妹就是自己的姐妹或自己
		Exp: regexp.MustCompile(`^(.*,[mw0])(,[olx][sb](&\d+)?)+,[olx]s(.*)$`),
		Str: "$1,xs$4#$1$4",
	},
	{ // 如果自己是男性,兄弟姐妹的姐妹就是自己的姐妹
		Exp: regexp.MustCompile(`(,[fh1])(,[olx][sb](&\d+)?)+,[olx]s`),
		Str: "$1,xs",
	},
	{ // 如果自己是女性,兄弟姐妹的兄弟就是自己的兄弟
		Exp: regexp.MustCompile(`(,[mw0])(,[olx][sb](&\d+)?)+,[olx]b`),
		Str: "$1,xb",
	},
	{ // 不知道性别，兄弟姐妹的兄弟是兄弟或自己
		Exp: regexp.MustCompile(`^,[olx][sb](&\d+)?,[olx]b(.+)?$`),
		Str: ",xb$2#$2",
	},
	{ // 不知道性别，兄弟姐妹的姐妹是姐妹或自己
		Exp: regexp.MustCompile(`^,[olx][sb](&\d+)?,[olx]s(.+)?$`),
		Str: ",xs$2#$2",
	},
	{ // 将复合称谓拆分
		Exp: regexp.MustCompile(`^,x([sb])$`),
		Str: ",o$1#,l$1",
	},
	// 父母
	{ // 母亲的丈夫是自己的父亲
		Exp: regexp.MustCompile(`m,h`),
		Str: "f",
	},
	{ // 父亲的妻子是自己的母亲
		Exp: regexp.MustCompile(`f,w`),
		Str: "m",
	},
	{ // 兄弟的父母就是自己的父母
		Exp: regexp.MustCompile(`,[xol][sb](&\d+)?(,[mf])`),
		Str: "$2",
	},
	// 父母的子女
	{ // 父母的女儿排序判断姐妹的排序
		Exp: regexp.MustCompile(`,[mf],d(&\d+)`),
		Str: ",xs$1",
	},
	{ // 父母的女儿年龄判断是姐姐还是妹妹
		Exp: regexp.MustCompile(`,[mf],d&([ol])`),
		Str: ",$1 s",
	},
	{ // 父母的儿子排序判断兄弟的排序
		Exp: regexp.MustCompile(`,[mf],s(&\d+)`),
		Str: ",xb$1",
	},
	{ // 父母的儿子年龄判断是哥哥还是弟弟
		Exp: regexp.MustCompile(`,[mf],s&([ol])`),
		Str: ",$1 b",
	},
	{ // 如果自己是男性,父母的儿子是兄弟或者自己
		Exp: regexp.MustCompile(`^(.*,[fh1]|[xol]b(&\d+)?),[mf],s(.*)$`),
		Str: "$1,xb$3#$1$3",
	},
	{ // 如果自己是女性,父母的女儿是姐妹或者自己
		Exp: regexp.MustCompile(`^(.*,[mw0]|[xol]s(&\d+)?),[mf],d(.*)$`),
		Str: "$1,xs$3#$1$3",
	},
	{ // 如果自己是女性,父母的儿子是自己的兄弟
		Exp: regexp.MustCompile(`(,[mw0]|[xol]s(&\d+)?),[mf],s`),
		Str: "$1,xb",
	},
	{ // 如果自己是男性,父母的女儿是自己的姐妹
		Exp: regexp.MustCompile(`(,[fh1]|[xol]b(&\d+)?),[mf],d`),
		Str: "$1,xs",
	},
	{ // 父母的儿子是兄弟或自己
		Exp: regexp.MustCompile(`^,[mf],s(.+)?$`),
		Str: ",xb$1#,1$1",
	},
	{ // 父母的女儿是姐妹或者自己
		Exp: regexp.MustCompile(`^,[mf],d(.+)?$`),
		Str: ",xs$1#,0$1",
	},
	// 孩子
	{ // 孩子的兄弟是自己的儿子(排序判断)
		Exp: regexp.MustCompile(`,[ds](&[ol\d]+)?,[olx]b(&\d+)`),
		Str: ",s$2",
	},
	{ // 孩子的姐妹是自己的女儿(排序判断)
		Exp: regexp.MustCompile(`,[ds](&[ol\d]+)?,[olx]s(&\d+)`),
		Str: ",d$2",
	},
	{ // 孩子的兄弟是自己的儿子(年龄判断)
		Exp: regexp.MustCompile(`,[ds]&o,ob`),
		Str: ",s&o",
	},
	{ // 孩子的姐妹是自己的女儿(年龄判断)
		Exp: regexp.MustCompile(`,[ds]&o,os`),
		Str: ",d&o",
	},
	{ // 孩子的兄弟是自己的儿子(年龄判断)
		Exp: regexp.MustCompile(`,[ds]&l,lb`),
		Str: ",s&l",
	},
	{ // 孩子的姐妹是自己的女儿(年龄判断)
		Exp: regexp.MustCompile(`,[ds]&l,ls`),
		Str: ",d&l",
	},
	{ // 孩子的姐妹是自己的女儿
		Exp: regexp.MustCompile(`,[ds](&[ol\d]+)?,[olx]s(&\d+)?`),
		Str: ",d$2",
	},
	{ // 孩子的兄弟是自己的儿子
		Exp: regexp.MustCompile(`,[ds](&[ol\d]+)?,[olx]b(&\d+)?`),
		Str: ",s$2",
	},
	// 夫妻
	{ // 自己是女性，女儿或儿子的妈妈是自己
		Exp: regexp.MustCompile(`(,[mwd0]|,[olx]s)(&[ol\d]+)?,[ds](&[ol\d]+)?,m`),
		Str: "$1",
	},
	{ // 自己是女性，女儿或儿子的爸爸是自己的丈夫
		Exp: regexp.MustCompile(`(,[mwd0]|,[olx]s)(&[ol\d]+)?,[ds](&[ol\d]+)?,f`),
		Str: "$1,h",
	},
	{ // 自己是男性，女儿或儿子的爸爸是自己
		Exp: regexp.MustCompile(`(,[fhs1]|,[olx]b)(&[ol\d]+)?,[ds](&[ol\d]+)?,f`),
		Str: "$1",
	},
	{ // 自己是男性，女儿或儿子的妈妈是自己的妻子
		Exp: regexp.MustCompile(`(,[fhs1]|,[olx]b)(&[ol\d]+)?,[ds](&[ol\d]+)?,m`),
		Str: "$1,w",
	},
	{ // 不知道性别，子女的妈妈是妻子或自己
		Exp: regexp.MustCompile(`^,[ds],m(.+)?$`),
		Str: ",w$1#$1",
	},
	{ // 不知道性别，子女的爸爸是丈夫或自己
		Exp: regexp.MustCompile(`^,[ds],f(.+)?$`),
		Str: ",h$1#$1",
	},
	{ // 夫妻的孩子就是自己的孩子
		Exp: regexp.MustCompile(`,[wh](,[ds])`),
		Str: "$1",
	},
	{ // 夫妻的对方是自己
		Exp: regexp.MustCompile(`,w,h|,h,w`),
		Str: "",
	},
	{ // 并列关系处理1 X[A|B|C]Y = XAY#X[B|C]Y
		Exp: regexp.MustCompile(`(.+)?\[([^\|]+?)\|([^\[\]]*\|[^\[\]]*)\](.+)?`),
		Str: "$1$2$4#$1[$3]$4",
	},
	{ // 并列关系处理2 X[A|B]Y = XAY#XBY
		Exp: regexp.MustCompile(`(.+)?\[([^\[\]\|]+?)\|([^\[\]\|]+?)?\](.+)?`),
		Str: "$1$2$4#$1$3$4",
	},
}
