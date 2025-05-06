package pairs

import (
	"regexp"
	"slices"
	"strings"

	"github.com/young2j/relationship/relationship-go/data"
	"github.com/young2j/relationship/relationship-go/selector"
)

// 通过ID获取关系合称
func GetPairsByID(id1 string, id2 string) []string {

	reId := regexp.MustCompile(`&\d+`)
	id1 = reId.ReplaceAllString(id1, "")
	id2 = reId.ReplaceAllString(id2, "")

	reIdX := regexp.MustCompile(`([ol])([bs])`)
	id1X := reIdX.ReplaceAllString(id1, "x$2")
	id2X := reIdX.ReplaceAllString(id2, "x$2")

	reIdR := regexp.MustCompile(`&[ol]`)
	id1R := reIdR.ReplaceAllString(id1, "")
	id2R := reIdR.ReplaceAllString(id2, "")

	replEveryItem := func(selectorList []string) []string {
		result := make([]string, 0, len(selectorList))
		reCleaned := regexp.MustCompile(`&[ol\d]+`)
		for _, selector := range selectorList {
			cleaned := reCleaned.ReplaceAllString(selector, "")
			replaced := reIdX.ReplaceAllString(cleaned, "x$2")
			result = append(result, replaced)
		}
		return result
	}

	result := make([]string, 0, len(data.Pair))
	resultX := make([]string, 0, len(data.Pair))
	resultR := make([]string, 0, len(data.Pair))

	for k, v := range data.Pair {
		selectors := strings.Split(k, "#")
		if len(selectors) > 1 {
			list1 := selector.Selector2ID(selectors[0], nil)
			list2 := selector.Selector2ID(selectors[1], nil)
			list1R := replEveryItem(list1)
			list2R := replEveryItem(list2)

			if (slices.Contains(list1, id1) && slices.Contains(list2, id2)) ||
				(slices.Contains(list1, id2) && slices.Contains(list2, id1)) {
				result = append(result, v[0])
			}

			if (slices.Contains(list1R, id1X) && slices.Contains(list2R, id2X)) ||
				(slices.Contains(list1R, id2X) && slices.Contains(list2R, id1X)) {
				resultX = append(resultX, v[0])
			}

			if (slices.Contains(list1R, id1R) && slices.Contains(list2R, id2R)) ||
				(slices.Contains(list1R, id2R) && slices.Contains(list2R, id1R)) {
				resultR = append(resultR, v[0])
			}
		}
	}

	if len(result) == 0 {
		result = resultX
	}

	if len(result) == 0 {
		result = resultR
	}

	return result
}
