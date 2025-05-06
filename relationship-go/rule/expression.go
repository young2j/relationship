/*
* 自然语言表达 - 将【语言表达】转换成【参数调用】
 */
package rule

import (
	"regexp"
	"strings"

	"github.com/young2j/relationship/relationship-go/helpers"
	"github.com/young2j/relationship/relationship-go/options"
)

type ExpressionItem struct {
	Exp *regexp.Regexp
	Opt func(group []string) *options.ParameterOptions
}

var Expression = []ExpressionItem{
	{
		Exp: regexp.MustCompile(`^(\S+?)(应该|得)?(称呼|叫|喊)(\S+?)(什么|啥)？?$`),
		Opt: func(group []string) *options.ParameterOptions {
			return options.Default().
				SetText(group[4]).
				SetTarget(group[1])
		},
	},
	{
		Exp: regexp.MustCompile(`^(\S+?)(应该|得)?(如何|怎么)(称呼|叫|喊)(\S+?)？?$`),
		Opt: func(group []string) *options.ParameterOptions {
			return options.Default().
				SetText(group[5]).
				SetTarget(group[1])
		},
	},
	{
		Exp: regexp.MustCompile(`^(\S+?)是(\S+?)的(谁|什么|什么亲戚|啥亲戚|什么人)？?$`),
		Opt: func(group []string) *options.ParameterOptions {
			typ := helpers.Ternary(strings.Contains(group[1], "的"), options.OptionTypeDefault, options.OptionTypeChain)
			return options.Default().
				SetText(group[1]).
				SetTarget(group[2]).
				SetType(typ)
		},
	},
	{
		Exp: regexp.MustCompile(`^(\S+?)(应该|得)?管(\S+?)叫(什么|啥)？?$`),
		Opt: func(group []string) *options.ParameterOptions {
			return options.Default().
				SetText(group[3]).
				SetTarget(group[1])
		},
	},
	{
		Exp: regexp.MustCompile(`^(\S+?)(和|与)(\S+?)(之间)?是(什么|啥)关系？?$`),
		Opt: func(group []string) *options.ParameterOptions {
			return options.Default().
				SetText(group[1]).
				SetTarget(group[3]).
				SetType(options.OptionTypePair)
		},
	},
	{
		Exp: regexp.MustCompile(`^(\S+?)对于(\S+?)是(什么|啥)关系？?$`),
		Opt: func(group []string) *options.ParameterOptions {
			return options.Default().
				SetText(group[1]).
				SetTarget(group[2]).
				SetType(options.OptionTypeChain)
		},
	},
	{
		Exp: regexp.MustCompile(`^(\S+?)(指的)?是(什么|啥)(意思|关系|亲戚关系|辈分|人)？?$`),
		Opt: func(group []string) *options.ParameterOptions {
			return options.Default().
				SetText(group[1]).
				SetType(options.OptionTypeChain)

		},
	},
	{
		Exp: regexp.MustCompile(`^(\S+?)(应该|得)?(称呼|叫|喊)(什么|啥)？?$`),
		Opt: func(group []string) *options.ParameterOptions {
			return options.Default().
				SetText(group[1])
		},
	},
	{
		Exp: regexp.MustCompile(`^(\S+?)(应该|得)?(如何|怎么)(称呼|叫|喊)？?$`),
		Opt: func(group []string) *options.ParameterOptions {
			return options.Default().
				SetText(group[1])
		},
	},
	{
		Exp: regexp.MustCompile(`^(\S+?)是(谁|什么|什么亲戚|啥亲戚|什么人|什么辈分|啥辈分)？?$`),
		Opt: func(group []string) *options.ParameterOptions {
			typ := helpers.Ternary(strings.Contains(group[1], "的"), options.OptionTypeDefault, options.OptionTypeChain)
			return options.Default().
				SetText(group[1]).
				SetType(typ)
		},
	},
}
