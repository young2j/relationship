/*
* 通过表达式获取配置
 */

package options

type OptionMode string

const (
	OptionModeDefault   OptionMode = "default"
	OptionModeGuangDong OptionMode = "guang_dong"
	OptionModeNorth     OptionMode = "north"
)

type OptionType string

const (
	OptionTypeDefault OptionType = "default"
	OptionTypeChain   OptionType = "chain"
	OptionTypePair    OptionType = "pair"
)

type ParameterOptions struct {
	// 目标对象：目标对象的称谓汉字表达，称谓间用‘的’字分隔
	Text string
	// 相对对象：相对对象的称谓汉字表达，称谓间用‘的’字分隔，空表示自己
	Target string
	// 本人性别：0表示女性,1表示男性
	Sex int
	// 转换类型：'default'计算称谓,'chain'计算关系链,'pair'计算关系合称
	Type OptionType
	// 称呼方式：true对方称呼我,false我称呼对方
	Reverse bool
	// 模式选择：使用setMode方法定制不同地区模式，在此选择自定义模式
	Mode OptionMode
	// 最短关系：计算两者之间的最短关系
	Optimal bool
}

func Default() *ParameterOptions {
	return &ParameterOptions{
		Sex:  -1,
		Type: OptionTypeDefault,
		Mode: OptionModeDefault,
	}
}

func (opts *ParameterOptions) SetText(text string) *ParameterOptions {
	opts.Text = text
	return opts
}
func (opts *ParameterOptions) SetTarget(target string) *ParameterOptions {
	opts.Target = target
	return opts
}
func (opts *ParameterOptions) SetSex(sex int) *ParameterOptions {
	opts.Sex = sex
	return opts
}
func (opts *ParameterOptions) SetType(typ OptionType) *ParameterOptions {
	opts.Type = typ
	return opts
}
func (opts *ParameterOptions) SetReverse(reverse bool) *ParameterOptions {
	opts.Reverse = reverse
	return opts
}
func (opts *ParameterOptions) SetMode(mode OptionMode) *ParameterOptions {
	opts.Mode = mode
	return opts
}
func (opts *ParameterOptions) SetOptimal(optimal bool) *ParameterOptions {
	opts.Optimal = optimal
	return opts
}
