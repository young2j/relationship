package main

import (
	"fmt"
	"regexp"

	"github.com/young2j/relationship/relationship-go"
)

func main() {
	// 示例1：基本的组匹配替换
	// ! 组$后面除非是组或者标点符号，否则需要空格才能识别
	text := "hello world"
	re := regexp.MustCompile(`(hello)\s+(world)`)
	result := re.ReplaceAllString(text, "$2 and $1 !")
	fmt.Println(result) // 输出: world and hello !
	result_0 := re.ReplaceAllString(text, "$2and$1 !")
	fmt.Printf("result_0: %v\n", result_0) // 输出: hello !
	result_1 := re.ReplaceAllString(text, "$2 and$1 !")
	fmt.Printf("result_1: %v\n", result_1) // 输出: world andhello !
	result_2 := re.ReplaceAllString(text, "$2and$1!")
	fmt.Printf("result_2: %v\n", result_2) // 输出: hello!
	result_3 := re.ReplaceAllString(text, "$2and$1a")
	fmt.Printf("result_3: %v\n", result_3) // 输出:
	result_4 := re.ReplaceAllString(text, "$2$1")
	fmt.Printf("result_4: %v\n", result_4) // 输出:worldhello
	result_5 := re.ReplaceAllString(text, "$2 $1")
	fmt.Printf("result_5: %v\n", result_5) // 输出:world hello
	result_6 := re.ReplaceAllString(text, "and$2 $1")
	fmt.Printf("result_6: %v\n", result_6) // 输出:andworld hello

	// 示例2：使用命名捕获组
	text2 := "John Doe"
	re2 := regexp.MustCompile(`(?P<first>\w+)\s+(?P<last>\w+)`)
	result2 := re2.ReplaceAllString(text2, "${last}, ${first}")
	fmt.Println(result2) // 输出: Doe, John

	// 示例3：使用函数进行替换
	re3 := regexp.MustCompile(`(\w+)@(\w+)\.com`)
	text3 := "contact@example.com"
	result3 := re3.ReplaceAllStringFunc(text3, func(s string) string {
		matches := re3.FindStringSubmatch(s)
		return matches[2] + "/" + matches[1]
	})
	fmt.Println(result3) // 输出: example/contact

	re4 := regexp.MustCompile(`^(\S+?)(应该|得)?(称呼|叫|喊)(\S+?)(什么|啥)？?$`)
	group := re4.FindStringSubmatch("张三什么？")
	fmt.Printf("group: %v\n", group == nil)

	fmt.Printf("(-50 / 10): %v\n", (-50 / 10))

	rel := relationship.NewRelationship()
	res := rel.Relationship(map[string]any{"target": "我", "text": "外婆的哥哥"})
	// res := rel.Relationship("我应该叫外婆的哥哥什么？")
	fmt.Printf("result: %v\n", res)

	res = rel.Relationship("外婆的哥哥")
	fmt.Printf("res: %v\n", res)

}
