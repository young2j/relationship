package selector

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/young2j/relationship/relationship-go/helpers"
	"github.com/young2j/relationship/relationship-go/id"
	"github.com/young2j/relationship/relationship-go/rule"
)

// SelectorOptions 选择器选项
type SelectorOptions struct {
	From    string
	To      string
	Sex     int
	Optimal bool
}

// SelectorResult 选择器结果
type SelectorResult struct {
	Selector string
	Sex      int
}

// 获得最简关系表达
func GetOptimal(options SelectorOptions) SelectorOptions {
	from := options.From
	to := options.To
	sex := options.Sex

	fromChain := strings.Split(from, ",")
	toChain := strings.Split(to, ",")

	getSex := func(chain string) int {
		matched, _ := regexp.MatchString(`^([fhs1](&[ol\d]+)?|[olx]b)(&[ol\d]+)?`, chain)
		if matched {
			return 1
		}
		return 0
	}

	group1 := func(patterns []*regexp.Regexp, s string) string {
		for _, pat := range patterns {
			matches := pat.FindStringSubmatch(s)
			if len(matches) > 1 {
				return matches[1]
			}
		}
		return ""
	}

	re1 := regexp.MustCompile(`^[xol][bs]|^[sd]`)
	re2 := regexp.MustCompile(`&([ol\d]+)`)
	re3 := regexp.MustCompile(`^[xol]([bs])`)
	re4 := regexp.MustCompile(`([ol])[bs]`)
	re5 := regexp.MustCompile(`^[xol]b|^s`)
	re6 := regexp.MustCompile(`^[xol]s|^d`)
	re7 := regexp.MustCompile(`([xol])[bs]`)

	for i := 0; i < min(len(fromChain), len(toChain)); i++ {
		if fromChain[i] == toChain[i] {
			from = strings.Join(fromChain[i+1:], ",")
			to = strings.Join(toChain[i+1:], ",")
			sex = getSex(fromChain[i])
			continue
		} else {
			if id.GetGenByID(fromChain[i]) == id.GetGenByID(toChain[i]) &&
				re1.MatchString(fromChain[i]) &&
				re1.MatchString(toChain[i]) {

				formType := re3.ReplaceAllString(re2.ReplaceAllString(fromChain[i], ""), "$1")
				toType := re3.ReplaceAllString(re2.ReplaceAllString(toChain[i], ""), "$1")

				if formType != toType {
					break
				}

				patterns := []*regexp.Regexp{re2, re4}
				fromAttr := group1(patterns, fromChain[i])
				toAttr := group1(patterns, toChain[i])

				if fromAttr != "" && toAttr != "" {
					fromNum, fromIsDigit := strconv.Atoi(fromAttr)
					toNum, toIsDigit := strconv.Atoi(toAttr)

					if fromIsDigit == nil && toIsDigit == nil {
						if fromNum > toNum {
							fromChain[i] = re6.ReplaceAllString(re5.ReplaceAllString(fromChain[i], "lb"), "ls")
						} else if fromNum < toNum {
							fromChain[i] = re6.ReplaceAllString(re5.ReplaceAllString(fromChain[i], "ob"), "os")
						}
					} else if (fromIsDigit == nil && toAttr == "o") || (fromAttr == "l" && toIsDigit == nil) {
						fromChain[i] = re6.ReplaceAllString(re5.ReplaceAllString(fromChain[i], "lb"), "ls")

					} else if (fromIsDigit == nil && toAttr == "l") || (fromAttr == "o" && toIsDigit == nil) {
						fromChain[i] = re6.ReplaceAllString(re5.ReplaceAllString(fromChain[i], "ob"), "os")

					}

					from = strings.Join(fromChain[i:], ",")
					to = strings.Join(toChain[i+1:], ",")
					sex = getSex(toChain[i])

				} else if options.Optimal {
					fromAttrMatch := re7.FindStringSubmatch(fromChain[i])
					fromAttr := ""
					if len(fromAttrMatch) > 1 {
						fromAttr = fromAttrMatch[1]
					}

					toAttrMatch := re7.FindStringSubmatch(toChain[i])
					toAttr := ""
					if len(toAttrMatch) > 1 {
						toAttr = toAttrMatch[1]
					}

					if fromAttr == "x" || toAttr == "x" {
						from = strings.Join(fromChain[i+1:], ",")
						to = strings.Join(toChain[i+1:], ",")
						sex = getSex(fromChain[i])
						continue
					}
				}
			}
			break
		}
	}

	return SelectorOptions{
		From: from,
		To:   to,
		Sex:  sex,
	}
}

// 中文获取选择器
func GetSelectors(s string, cacheData map[string][]string) (result []string) {
	// 基础字符串替换
	// 将所有"之"替换为"的"
	s = strings.ReplaceAll(s, "之", "的")
	// 将"吾之xxx"或"吾xxx"模式替换为"xxx"
	s = helpers.ReplaceAllStringAndSpace(s, `吾之?(.+)`, "$1")
	// 将"我的xxx"或"我xxx"模式替换为"xxx"
	s = helpers.ReplaceAllStringAndSpace(s, `我的?(.+)`, "$1")

	// 惯用口语标准化
	if regexp.MustCompile(`[^娘婆岳亲]家的?(孩子|儿子|女儿)`).MatchString(s) {
		s = helpers.ReplaceAllStringAndSpace(s, `家的?`, "的")
	}

	s = helpers.ReplaceAllStringAndSpace(s, `(舅|姑)+(爸|父|丈|妈|母)?家的?(哥|姐|弟|妹)+`, "$1 表$3")
	s = helpers.ReplaceAllStringAndSpace(s, `(舅|姑)+(爸|父|丈|妈|母)?家的?`, "$1 表")
	s = helpers.ReplaceAllStringAndSpace(s, `(伯|叔)+(父|母)?家的?(哥|姐|弟|妹)+`, "堂$3")
	s = helpers.ReplaceAllStringAndSpace(s, `(伯|叔)+(父|母)?家的?`, "堂")
	s = helpers.ReplaceAllStringAndSpace(s, `姨+(爸|父|丈|妈|母)?家的?(哥|姐|弟|妹)+`, "姨$2")
	s = helpers.ReplaceAllStringAndSpace(s, `姨+(爸|父|丈|妈|母)?家的?`, "姨")

	lists := strings.Split(s, "的")
	isMatch := true

	var (
		getList  func(string)
		keywords []string
	)
	getList = func(name string) {
		if name == "" {
			return
		}

		// 词义扩展
		for _, item := range rule.Replace {
			for _, word := range item.Arr {
				name1 := helpers.ReplaceAllStringAndSpace2(name, item.Exp, word)
				if name1 != name {
					keywords = append(keywords, name1)
					getList(name1)
				}
			}
		}

		// 同义词替换
		for word, similar := range rule.Similar {
			name1 := strings.ReplaceAll(name, word, similar)
			if name1 != name {
				keywords = append(keywords, name1)
			}

			name2 := strings.ReplaceAll(name, similar, word)
			if name2 != name {
				keywords = append(keywords, name2)
			}
		}

	}

	c := 0
	for c < len(lists) {
		// 当前匹配词
		name := lists[c]
		c += 1
		// 当前匹配词可能性
		items := make([]string, 0, 2)
		keywords = []string{name}
		getList(name)

		// 通过关键词找关系
		var itemsMap = [3][]string{}
		for _, kw := range keywords {
			kw = regexp.MustCompile(`^[尕幺细满碎晚末尾幼]`).ReplaceAllString(kw, "小")
			match := regexp.MustCompile(`^[大小]|^[一二三四五六七八九十]+`).FindString(kw)

			if match != "" { // 匹配排序
				xName := regexp.MustCompile(match).ReplaceAllString(kw, "几")
				rName := regexp.MustCompile(match).ReplaceAllString(kw, "")
				num := helpers.Zh2Number(match)

				for index, item := range []string{xName, rName, kw} {
					ids := cacheData[item]
					for _, i := range ids {
						id_ := regexp.MustCompile(`(,[hw])$`).ReplaceAllString(i, fmt.Sprintf("&%v$1", num))
						id_ = regexp.MustCompile(`([^hw]+)$`).ReplaceAllString(id_, fmt.Sprintf("$1&%v", num))
						// 直系祖辈不参与排序
						if !regexp.MustCompile(`^[mf,]+$`).MatchString(i) &&
							!regexp.MustCompile(`^[从世]`).MatchString(kw) {
							itemsMap[index] = append(itemsMap[index], id_)
						}
					}
				}
			}

			if val, ok := cacheData[kw]; ok {
				items = append(items, val...)
			}
		}

		// 如找不到结果，检查是否存在称呼的排行问题
		for _, itemsX := range itemsMap {
			if len(items) == 0 {
				items = itemsX
			}
		}

		// 完全匹配不到结果
		if len(items) == 0 {
			isMatch = false
		}

		res := make([]string, 0, 1)
		if len(result) == 0 {
			result = append(result, "")
		}

		for _, a := range result {
			for _, b := range items {
				if b != "" {
					b = "," + b
				}
				res = append(res, a+b)
			}
		}
		result = res
	}

	if isMatch {
		result = id.FilterID(result)
		return
	}

	return
}

// 合并选择器，查找两个对象之间的关系
func MergeSelector(param SelectorOptions) (results []SelectorResult) {
	sex := param.Sex

	if sex < 0 {
		fromSex := -1
		toSex := -1
		reW := regexp.MustCompile(`^,[w1]`)
		reH := regexp.MustCompile(`^,[h0]`)

		if reW.MatchString(param.From) {
			fromSex = 1
		} else if reH.MatchString(param.From) {
			fromSex = 0
		}
		if reW.MatchString(param.To) {
			toSex = 1
		} else if reH.MatchString(param.To) {
			toSex = 0
		}

		if fromSex == -1 && toSex > -1 {
			sex = toSex
		} else if fromSex > -1 && toSex == -1 {
			sex = fromSex
		} else if fromSex == toSex {
			sex = fromSex
		} else {
			return
		}
	}

	fromIds := Selector2ID(param.From, &sex)
	toIds := Selector2ID(param.To, &sex)

	if len(fromIds) == 0 || len(toIds) == 0 {
		return
	}

	for _, fromId := range fromIds {
		for _, toId := range toIds {
			resultSex := sex
			selector := "," + toId

			if regexp.MustCompile(`,([fhs1](&[ol\d]+)?|[olx]b)(&[ol\d]+)?$`).MatchString(selector) {
				resultSex = 1
			}
			if regexp.MustCompile(`,([mwd0](&[ol\d]+)?|[olx]s)(&[ol\d]+)?$`).MatchString(selector) {
				resultSex = 0
			}

			if fromId != "" && toId != "" {
				isOptimal := param.Optimal
				regex := regexp.MustCompile(`&\d+`)
				if regex.MatchString(fromId) || regex.MatchString(toId) {
					isOptimal = true
				}

				if isOptimal {
					optimalOptions := GetOptimal(SelectorOptions{
						From:    fromId,
						To:      toId,
						Sex:     sex,
						Optimal: param.Optimal,
					})
					fromId = optimalOptions.From
					toId = optimalOptions.To
					sex = optimalOptions.Sex
				}
			}

			toRids := []string{""}
			if toId != "" {
				toRids = id.ReverseID(toId, sex)
			}

			for _, toR := range toRids {
				selector := ""
				if toR != "" {
					selector += "," + toR
				}
				if fromId != "" {
					selector += "," + fromId
				}
				results = append(results, SelectorResult{
					Selector: selector,
					Sex:      resultSex,
				})
			}
		}
	}

	return
}

// 扩展选择器，将【关系链】缩减成最简表达形式
//
//	selector: 关系选择器字符串
//	return: 扩展后的选择器列表
func ExpandSelector(selector string) (result []string) {
	hashMap := make(map[string]bool)

	var getSelector func(string)
	getSelector = func(selector string) {
		if _, ok := hashMap[selector]; !ok {
			hashMap[selector] = true
			s := ""

			for {
				s = selector
				for _, item := range rule.Filter {
					selector = helpers.ReplaceAllStringAndSpace2(selector, item.Exp, item.Str)
					if strings.Contains(selector, "#") {
						for _, part := range strings.Split(selector, "#") {
							getSelector(part)
						}
						return
					}
				}

				if s == selector {
					break
				}
			}

			// 同志关系去除
			if regexp.MustCompile(`,[mwd0](&[ol\d+])?,w|,[hfs1](&[ol\d]+)?,h`).MatchString(selector) {
				return
			}

			result = append(result, selector)
		}
	}

	getSelector(selector)

	return
}

// 选择器转ID
func Selector2ID(sel string, sex *int) (result []string) {
	if !strings.HasPrefix(sel, ",") {
		sel = "," + sel
	}

	// 性别判断
	var sexVal int
	if sex != nil {
		if *sex < 0 {
			if regexp.MustCompile(`^,[w1]`).MatchString(sel) {
				sexVal = 1
				sex = &sexVal
			} else if regexp.MustCompile(`^,[h0]`).MatchString(sel) {
				sexVal = 0
				sex = &sexVal
			}
		} else if *sex == 1 && regexp.MustCompile(`^,[h0]`).MatchString(sel) {
			return
		} else if *sex == 0 && regexp.MustCompile(`^,[w1]`).MatchString(sel) {
			return
		}

		if *sex > -1 && !strings.Contains(sel, ",1") && !strings.Contains(sel, ",0") {
			sel = "," + strconv.Itoa(*sex) + sel
		}
	}

	// 同志关系去除
	if regexp.MustCompile(`,[mwd0](&[ol\d]+)?,w|,[hfs1](&[ol\d]+)?,h`).MatchString(sel) {
		return
	}

	expanded := ExpandSelector(sel)
	for _, s := range expanded {
		clean := regexp.MustCompile(`,[01]`).ReplaceAllString(s, "")
		if len(clean) > 1 {
			// 去前面逗号和性别信息
			clean = clean[1:]
		} else {
			clean = ""
		}
		result = append(result, clean)
	}

	result = id.FilterID(result)

	return
}
