package tests

import (
	"testing"

	"github.com/young2j/relationship/relationship-go"
	"github.com/young2j/relationship/relationship-go/options"
)

// BenchmarkSimpleQueries用于测试简单查询性能
func BenchmarkSimpleQueries(b *testing.B) {
	r := relationship.NewRelationship()
	for range b.N {
		r.Relationship(map[string]any{"text": "爸爸的哥哥的弟弟的儿子"})
	}
}

// BenchmarkComplexQueries用于测试复杂查询性能
func BenchmarkComplexQueries(b *testing.B) {
	r := relationship.NewRelationship()
	for range b.N {
		r.Relationship(map[string]any{"text": "姐姐的老公的姐姐的老公"})
		r.Relationship(map[string]any{"text": "爸爸的妈妈的姐姐的儿子"})
		r.Relationship(map[string]any{"text": "老婆的外孙的姥爷", "sex": 1})
	}
}

// BenchmarkReverseQueries用于测试反向关系查询性能
func BenchmarkReverseQueries(b *testing.B) {
	r := relationship.NewRelationship()

	for range b.N {
		r.Relationship(map[string]any{"text": "爸爸的舅舅", "sex": 0, "reverse": true})
		r.Relationship(map[string]any{"text": "岳母", "target": "女儿", "reverse": true})
	}
}

// BenchmarkExpressionStringQueries用于测试表达式查询性能
func BenchmarkExpressionStringQueries(b *testing.B) {
	r := relationship.NewRelationship()

	for range b.N {
		r.Relationship("外婆和奶奶之间是什么关系？")
		r.Relationship("我应该叫外婆的哥哥什么？")
	}
}

// BenchmarkExpressionMapQueries用于测试表达式查询性能
func BenchmarkExpressionMapQueries(b *testing.B) {
	r := relationship.NewRelationship()

	for range b.N {
		r.Relationship(map[string]any{"text": "外婆", "target": "奶奶", "type":"pair"})
		r.Relationship(map[string]any{"text": "我", "target": "外婆的哥哥"})
	}
}

// BenchmarkExpressionOptsQueries用于测试表达式查询性能
func BenchmarkExpressionOptsQueries(b *testing.B) {
	r := relationship.NewRelationship()

	for range b.N {
		r.Relationship(options.Default().SetText("外婆").SetTarget("奶奶").SetType("pair"))
		r.Relationship(options.Default().SetText("外婆的哥哥").SetTarget("我"))
	}
}
