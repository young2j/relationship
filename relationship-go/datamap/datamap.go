package datamap

/*
完整关系链数据 - 合并各类关系链数据
*/

import (
	"maps"
	"regexp"
	"strings"

	"github.com/young2j/relationship/relationship-go/data"
	"github.com/young2j/relationship/relationship-go/helpers"
	"github.com/young2j/relationship/relationship-go/selector"
)

// selectorNameListMap 是从选择器字符串到名称列表的映射
type selectorNameListMap map[string][]string

// 分支 - 前缀处理
func getPrefixMap() map[string]selectorNameListMap {
	prefixMap := make(map[string]selectorNameListMap)
	for tag, dictValues := range data.Prefix {
		prefixMap[tag] = make(selectorNameListMap)
		for sel, nameList := range dictValues {
			for _, s := range selector.ExpandSelector(sel) {
				prefixMap[tag][s] = nameList
			}
		}
	}
	return prefixMap
}

// 分支 - 节点处理
func getBranchMap() selectorNameListMap {
	branchMap := make(selectorNameListMap)
	for sel, nameList := range data.Branch {
		for _, s := range selector.ExpandSelector(sel) {
			branchMap[s] = nameList
		}
	}
	return branchMap
}

// 分支 - 合并
func getMap(prefixMap map[string]selectorNameListMap, branchMap selectorNameListMap) selectorNameListMap {
	mapResult := make(selectorNameListMap)

	re := regexp.MustCompile(`(\{.*?\})`)
	for branchSel, branchNameList := range branchMap {
		// 提取标签
		matches := re.FindStringSubmatch(branchSel)
		if len(matches) < 1 {
			continue
		}
		tag := matches[0]

		for prefixTagSel, prefixNameList := range prefixMap[tag] {
			newBranchSel := strings.Replace(branchSel, tag, prefixTagSel, 1)

			// 过滤条件
			isFilter := strings.Contains(newBranchSel, "h,h") ||
				strings.Contains(newBranchSel, "w,w") ||
				strings.Contains(newBranchSel, "w,h") ||
				strings.Contains(newBranchSel, "h,w")

			if !isFilter {
				newNameList := make([]string, 0, len(prefixNameList)*len(branchNameList))

				for _, prefixName := range prefixNameList {
					for _, branchName := range branchNameList {
						if strings.Contains(branchName, "?") {
							newNameList = append(newNameList, strings.ReplaceAll(branchName, "?", prefixName))
						} else {
							newNameList = append(newNameList, prefixName+branchName)
						}
					}
				}

				if _, exists := mapResult[newBranchSel]; !exists {
					if multipleData, ok := data.Multiple[newBranchSel]; ok {
						mapResult[newBranchSel] = multipleData
					} else {
						mapResult[newBranchSel] = []string{}
					}
				}

				mapResult[newBranchSel] = append(newNameList, mapResult[newBranchSel]...)
			}
		}
	}

	return mapResult
}

func mateMap(map_ selectorNameListMap) {
	mateMap_ := selectorNameListMap{
		"w": {"妻", "内", "岳", "岳家", "丈人"},
		"h": {"夫", "外", "公", "婆家", "婆婆"},
	}

	// 构建名称集合
	nameSet := make(map[string]struct{})
	for _, nameList := range map_ {
		for _, name := range nameList {
			nameSet[name] = struct{}{}
		}
	}

	// 复制原映射
	mapCopy := maps.Clone(map_)

	reFm := regexp.MustCompile(`^[fm]`)
	reBro := regexp.MustCompile(`^[olx][bs]$|^[olx][bs],[^mf]`)
	for sel, nameList := range mapCopy {
		// 只对长辈或者兄弟辈匹配
		if reFm.MatchString(sel) || reBro.MatchString(sel) {
			for mateSel, mateNameList := range mateMap_ {
				newSel := mateSel + "," + sel

				if strings.ContainsAny(sel, "fm") {
					// 替换年龄标记
					newSelX := helpers.ReplaceAllStringAndSpace(newSel, `,[ol]([sb])(,[wh])?$`, ",x$1$2")
					newSelX = helpers.ReplaceAllStringAndSpace(newSelX, `(,[sd])&[ol](,[wh])?$`, "$1$2")

					// 不扩大解释年龄
					if newSelX != newSel {
						if _, exists := map_[newSelX]; exists {
							continue
						}
					}
				}

				if _, exists := map_[newSel]; !exists {
					map_[newSel] = []string{}
				}

				for _, prefix := range mateNameList {
					for _, name := range nameList {
						newName := prefix + name
						// 配偶组合的称呼不得与原有称呼冲突
						if _, ok := nameSet[newName]; !ok {
							map_[newSel] = append(map_[newSel], newName)
						}
					}
				}
			}
		}
	}
}

func InitMap() selectorNameListMap {
	// 复制 multiple 数据
	map_ := maps.Clone(data.Multiple)

	// 拷贝基础关系映射
	maps.Copy(map_, getMap(getPrefixMap(), getBranchMap()))

	// 版权彩蛋
	map_["o"] = []string{"passer-by.com", "\u4f5c\u8005", "young2j", "\u0067\u006F\u4F5C\u8005"}

	// 添加主要关系
	for mainSel, mainNameList := range data.Main {
		if _, ok := map_[mainSel]; !ok {
			map_[mainSel] = mainNameList
		} else {
			map_[mainSel] = append(mainNameList, map_[mainSel]...)
		}
	}

	// 处理配偶关系
	mateMap(map_)

	return map_
}
