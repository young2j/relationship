/*
* 词义扩展 - 对称呼词义进行扩展
 */

package rule

import "regexp"

type ReplaceItem struct {
	Exp *regexp.Regexp
	Arr []string
}

var Replace = []ReplaceItem{
	{ // 规则1
		Exp: regexp.MustCompile(`^从表`),
		Arr: []string{"从父姑表", "从父舅表", "从父姨表", "从母姑表", "从母舅表", "从母叔表"},
	},
	{ // 规则2
		Exp: regexp.MustCompile(`^表表`),
		Arr: []string{"姑表叔表", "姑表姑表", "姑表舅表", "姑表姨表", "舅表叔表", "舅表姑表", "舅表舅表", "舅表姨表"},
	},
	{ // 规则3
		Exp: regexp.MustCompile(`^([夫妻内外]?)表`),
		Arr: []string{"$1 姑表", "$1 舅表"},
	},
	{ // 规则4
		Exp: regexp.MustCompile(`^([姑舅])表([^伯叔]+)`),
		Arr: []string{"$1 表伯$2", "$1 表叔$2"},
	},
	{ // 规则5
		Exp: regexp.MustCompile(`^姻`),
		Arr: []string{"姑姻", "姨姻", "姊妹姻", "女姻"},
	},
	{ // 规则6
		Exp: regexp.MustCompile(`^眷`),
		Arr: []string{"叔眷", "舅眷", "兄弟眷", "男眷"},
	},
	{ // 规则7
		Exp: regexp.MustCompile(`^亲家`),
		Arr: []string{"姊妹姻", "兄弟眷"},
	},
	{ // 规则8
		Exp: regexp.MustCompile(`^([堂表姨]?)([曾高天烈太远鼻]?)(祖?)([伯叔姑舅姨])`),
		Arr: []string{"$1$4$2$3"},
	},
	{ // 规则9
		Exp: regexp.MustCompile(`^([曾高天烈太远鼻]?)祖?王姑`),
		Arr: []string{"姑$1 祖母"},
	},
	{ // 规则10
		Exp: regexp.MustCompile(`^([曾玄元来晜仍云耳])([侄甥])`),
		Arr: []string{"$2$1"},
	},
	{ // 规则11
		Exp: regexp.MustCompile(`^外表([伯叔姑舅姨])`),
		Arr: []string{"姑表$1 外", "舅表$1 外"},
	},
	{ // 规则12
		Exp: regexp.MustCompile(`([堂表姨]?)外甥`),
		Arr: []string{"$1 甥"},
	},
	{ // 规则13
		Exp: regexp.MustCompile(`^([舅叔])([曾玄外]*)孙`),
		Arr: []string{"$1 侄$2 孙"},
	},
	{ // 规则14
		Exp: regexp.MustCompile(`^([姨姑])([曾玄外]*)孙`),
		Arr: []string{"$1 甥$2 孙"},
	},
	{ // 规则15
		Exp: regexp.MustCompile(`([孙甥侄])$`),
		Arr: []string{"$1 男", "$1 女"},
	},
	{ // 规则16
		Exp: regexp.MustCompile(`([姑舅姨叔])([孙外]*)([男女])$`),
		Arr: []string{"$1 表侄$2$3", "$1 表甥$2$3"},
	},
	{ // 规则17
		Exp: regexp.MustCompile(`(.+)父母$`),
		Arr: []string{"$1 父", "$1 母"},
	},
	{ // 规则18
		Exp: regexp.MustCompile(`(.+)公婆$`),
		Arr: []string{"$1 公", "$1 婆"},
	},
	{ // 规则19
		Exp: regexp.MustCompile(`祖$`),
		Arr: []string{"祖父"},
	},
	{ // 规则20
		Exp: regexp.MustCompile(`(新妇|媳)$`),
		Arr: []string{"媳妇"},
	},
	{ // 规则21
		Exp: regexp.MustCompile(`嫂$`),
		Arr: []string{"兄妇"},
	},
	{ // 规则22
		Exp: regexp.MustCompile(`女儿$`),
		Arr: []string{"女"},
	},
}
