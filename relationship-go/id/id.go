/*
标识符 - 唯一性关系链，将【关系链】转换成【中文表述】
*/

package id

import (
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/young2j/relationship/relationship-go/data"
	"github.com/young2j/relationship/relationship-go/helpers"
)

// 逆转ID
func ReverseID(id string, sex int) []string {
	if id == "" {
		return []string{""}
	}
	// 映射关系
	hashMap := map[string][]string{
		"f":  {"d", "s"},
		"m":  {"d", "s"},
		"h":  {"w", ""},
		"w":  {"", "h"},
		"s":  {"m", "f"},
		"d":  {"m", "f"},
		"lb": {"os", "ob"},
		"ob": {"ls", "lb"},
		"xb": {"xs", "xb"},
		"ls": {"os", "ob"},
		"os": {"ls", "lb"},
		"xs": {"xs", "xb"},
	}
	// 年纪判断
	age := ""
	if strings.HasSuffix(id, "&o") {
		age = "&l"
	} else if strings.HasSuffix(id, "&l") {
		age = "&o"
	}
	reId := regexp.MustCompile(`&[ol\d+]`)
	id = reId.ReplaceAllString(id, "")
	// 性别判断
	if sex < 0 {
		if strings.HasPrefix(id, "w") {
			sex = 1
		} else if strings.HasPrefix(id, "h") {
			sex = 0
		}
	}

	doing := func(sex int) string {
		sid := "," + strconv.Itoa(sex) + "," + id
		sid = regexp.MustCompile(`,[fhs]|,[olx]b`).ReplaceAllString(sid, ",1")
		sid = regexp.MustCompile(`,[mwd]|,[olx]s`).ReplaceAllString(sid, ",0")
		sid = sid[:len(sid)-2]
		sidArr := strings.Split(sid, ",")
		slices.Reverse(sidArr)

		idParts := strings.Split(id, ",")
		slices.Reverse(idParts)

		var rID []string
		for i, idPart := range idParts {
			idx, _ := strconv.Atoi(sidArr[i])
			if part, ok := hashMap[idPart]; ok {
				rID = append(rID, part[idx])
			}
		}

		rid := strings.Join(rID, ",")
		gen := GetGenByID(rid)

		return rid + helpers.Ternary(gen == 0, age, "")
	}

	if sex < 0 {
		return []string{doing(1), doing(0)}
	} else {
		return []string{doing(sex)}
	}
}

// ID列表去重
func FilterID(arr []string) []string {
	repl := func(item string) string {
		replItem := regexp.MustCompile(`[ol]([s|b])`).ReplaceAllString(item, "x$1")
		replItem = regexp.MustCompile(`&[ol]`).ReplaceAllString(replItem, "")
		return replItem
	}

	sameList := make([]string, 0, len(arr))
	for _, item := range arr {
		if repl(item) == item {
			sameList = append(sameList, item)
		}
	}

	filtered := make([]string, 0, len(arr))
	seen := make(map[string]struct{})
	for _, item := range arr {
		replItem := repl(item)
		if slices.Contains(sameList, item) || (item != replItem && !slices.Contains(sameList, replItem)) {
			if _, ok := seen[item]; !ok {
				seen[item] = struct{}{}
				filtered = append(filtered, item)
			}
		}
	}

	return filtered
}

// 通过ID获取世代数
func GetGenByID(id string) int {
	gMap := map[string]int{
		"f": 1,
		"m": 1,
		"s": -1,
		"d": -1,
	}
	gen := 0
	regex := regexp.MustCompile(`&[ol\d]+`)
	for _, sub := range strings.Split(id, ",") {
		s := regex.ReplaceAllString(sub, "")
		if val, ok := gMap[s]; ok {
			gen += val
		}
	}

	return gen
}

// 通过ID获取关系称呼
func GetItemsByID(id string, modeData map[string][]string) []string {
	items := make([]string, 0, 4)

	getData := func(key string) []string {
		ids := make([]string, 0, 2)
		k1 := regexp.MustCompile(`(,[sd])(,[wh])?$`).ReplaceAllString(key, "$1&o$2")
		k2 := regexp.MustCompile(`(,[sd])(,[wh])?$`).ReplaceAllString(key, "$1&l$2")

		if _, ok1 := modeData[k1]; ok1 {
			if _, ok2 := modeData[k2]; ok2 {
				ids = []string{k1, k2}
			}
		} else if _, ok := modeData[key]; ok {
			ids = []string{key}
		}

		result := make([]string, 0, len(ids))
		for _, id := range FilterID(ids) {
			if val, ok := modeData[id]; ok && len(val) > 0 {
				result = append(result, val[0])
			}
		}
		return result
	}

	// 对排序进行处理
	regex := regexp.MustCompile(`&([\d]+)(,[hw])?$`)
	group := regex.FindStringSubmatch(id)
	if len(group) > 0 {
		num := group[1]
		zh := helpers.Number2Zh(num)
		id = regexp.MustCompile(`&\d+`).ReplaceAllString(id, "")

		if names, ok := data.Sort[id]; ok {
			item := strings.ReplaceAll(names[0], "几", zh)
			items = append(items, item)
		} else if names, ok := modeData[id]; ok {
			gen := GetGenByID(id)
			item := ""
			if gen < 3 && !regexp.MustCompile(`[hw],`).MatchString(id) {
				for _, name := range names {
					if item == "" && strings.Contains(name, "几") {
						item = strings.ReplaceAll(name, "几", zh)
					}
				}
				if item == "" {
					item = names[0]
					reg := regexp.MustCompile(`^[大小]`)
					if reg.MatchString(item) {
						item = reg.ReplaceAllString(item, zh)
					} else {
						item = zh + item
					}
				}
			}
			items = append(items, item)
		}
	}

	// 直接匹配称呼
	if len(items) == 0 {
		id = regexp.MustCompile(`&\d+`).ReplaceAllString(id, "")
		items = getData(id)
	}

	// 忽略年龄条件查找
	if len(items) == 0 {
		id = regexp.MustCompile(`&[ol]`).ReplaceAllString(id, "")
		items = getData(id)
	}

	// 忽略年龄条件查找
	if len(items) == 0 {
		id = regexp.MustCompile(`[ol](b|s)`).ReplaceAllString(id, "x$1")
		items = getData(id)
	}

	// 缩小访问查找
	if len(items) == 0 {
		lID := strings.ReplaceAll(id, "x", "l")
		oID := strings.ReplaceAll(id, "x", "o")
		items = slices.Concat(getData(oID), getData(lID))
	}

	return items
}

// 通过ID获取关系链
func GetChainByID(id string, modeData map[string][]string, sex int) string {
	data := make(map[string][]string, len(modeData))
	maps.Copy(data, modeData)
	data["xb"] = []string{"兄弟"}
	data["xs"] = []string{"姐妹"}

	parts := make([]string, 0, 4)
	for _, sign := range strings.Split(id, ",") {
		sign := regexp.MustCompile(`&[ol\d]+`).ReplaceAllString(sign, "")
		if names, ok := data[sign]; ok && len(names) > 0 {
			parts = append(parts, names[0])
		}
	}

	item := strings.Join(parts, "的")

	if sex > -1 {
		key := strconv.Itoa(sex) + "," + id
		if _, ok := data[key]; ok {
			if sex == 0 {
				item = "(女性)" + item
			} else if sex == 1 {
				item = "(男性)" + item
			}
		}
	}

	return item
}
