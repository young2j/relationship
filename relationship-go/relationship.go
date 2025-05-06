package relationship

import (
	"regexp"
	"strconv"

	"github.com/young2j/relationship/relationship-go/cache"
	"github.com/young2j/relationship/relationship-go/helpers"
	"github.com/young2j/relationship/relationship-go/id"
	"github.com/young2j/relationship/relationship-go/id/pairs"
	"github.com/young2j/relationship/relationship-go/modedata"
	"github.com/young2j/relationship/relationship-go/options"
	"github.com/young2j/relationship/relationship-go/options/optionsbuild"
	"github.com/young2j/relationship/relationship-go/selector"
)

// 从其他模块导入的类型和函数
type OptionType int
type ParameterOptions struct {
	Mode    options.OptionMode
	Text    string
	Target  string
	Type    OptionType
	Sex     int
	Reverse bool
	Optimal bool
}

type SelectorOptions struct {
	From    string
	To      string
	Sex     int
	Optimal bool
}

type MergeData struct {
	Selector string
	Sex      int
}

// Relationship 表示亲戚关系的类
type Relationship struct {
	data      map[string][]string
	cacheData map[string][]string
	mode      options.OptionMode
}

// NewRelationship 创建一个新的Relationship实例
func NewRelationship() *Relationship {
	return NewRelationshipWithMode(options.OptionModeDefault)
}

// NewRelationship 创建一个新的指定区域模式Relationship实例
func NewRelationshipWithMode(mode options.OptionMode) *Relationship {
	r := &Relationship{
		data:      make(map[string][]string),
		cacheData: make(map[string][]string),
		mode:      mode,
	}

	r.SetMode(mode)

	return r
}

// Relationship 计算亲戚关系
// parameter 可以是
//   - 字符串
//   - map[string]any
//   - options.ParameterOptions
//   - *options.ParameterOptions
func (r *Relationship) Relationship(parameter any) (result []string) {
	opts := options.Default()

	switch p := parameter.(type) {
	case options.ParameterOptions:
		opts = &p
	case *options.ParameterOptions:
		opts = p
	case string:
		opts = optionsbuild.FromString(p)
	case map[string]any:
		opts = optionsbuild.FromMapping(p)
	}

	if opts.Mode != r.mode {
		r.SetMode(opts.Mode)
	}

	fromSelectors := selector.GetSelectors(opts.Text, r.cacheData)
	toSelectors := selector.GetSelectors(opts.Target, r.cacheData)
	if len(toSelectors) == 0 {
		toSelectors = []string{""}
	}

	for _, fromSelector := range fromSelectors {
		for _, toSelector := range toSelectors {
			mergeDatas := selector.MergeSelector(selector.SelectorOptions{
				From:    fromSelector,
				To:      toSelector,
				Sex:     opts.Sex,
				Optimal: opts.Optimal,
			})

			idsDuplicated := make(map[string]struct{})
			for _, mergeData := range mergeDatas {
				ids := selector.Selector2ID(mergeData.Selector, &mergeData.Sex)

				for _, id_ := range ids {
					if _, ok := idsDuplicated[id_]; ok {
						continue
					}
					idsDuplicated[id_] = struct{}{}

					temps := []string{id_}
					sex := mergeData.Sex

					if opts.Reverse {
						temps = id.ReverseID(id_, sex)
						re := regexp.MustCompile(`([fhs1](&[ol\d]+)?|[olx]b)$`)
						if re.MatchString(id_) {
							sex = 1
						} else {
							sex = 0
						}
					}

					if opts.Type == options.OptionTypeChain {
						for _, idStr := range temps {
							item := id.GetChainByID(idStr, r.data, mergeData.Sex)
							if item != "" {
								result = append(result, item)
							}
						}
					} else if opts.Type == options.OptionTypePair {
						temps = id.ReverseID(id_, mergeData.Sex)
						for _, rId := range temps {
							pairs := pairs.GetPairsByID(id_, rId)
							result = append(result, pairs...)
						}
					} else {
						for _, idStr := range temps {
							items := id.GetItemsByID(idStr, r.data)
							if len(items) == 0 {
								items = id.GetItemsByID(strconv.Itoa(sex)+","+idStr, r.data)
							}
							result = append(result, items...)
						}
					}
				}
			}
		}
	}

	// 字典序排列且去重
	result = helpers.RemoveDuplicates(result)

	return
}

// DataCount 返回数据数量
func (r *Relationship) DataCount() int {
	return len(r.data)
}

// SetMode 设置模式
func (r *Relationship) SetMode(mode options.OptionMode) {
	r.SetModeWithExtraData(mode, nil)
}

// SetModeWithExtraData 设置模式
func (r *Relationship) SetModeWithExtraData(mode options.OptionMode, extraData map[string][]string) {
	r.data = modedata.GetModeData(mode)

	if extraData != nil {
		for k, v := range extraData {
			if dat, ok := r.data[k]; ok {
				r.data[k] = append(v, dat...)
			} else {
				r.data[k] = v
			}
		}
	}

	r.cacheData = cache.GetCacheData(r.data)
}

// GetData 获取数据
func (r *Relationship) GetData() map[string][]string {
	return r.GetModeData(r.mode)
}

// GetModeData 获取模式数据
func (r *Relationship) GetModeData(mode options.OptionMode) map[string][]string {
	if mode == r.mode {
		return r.data
	}

	return modedata.GetModeData(mode)
}
