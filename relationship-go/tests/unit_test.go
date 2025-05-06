package tests

import (
	"reflect"
	"testing"

	"github.com/young2j/relationship/relationship-go"
)

// SameElementsSlice checks if two slices contain the same elements, regardless of order.
func SameElementsSlice[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	m := make(map[T]int)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}

	return true
}

func TestCall(t *testing.T) {
	rel := relationship.NewRelationship()

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "爸爸的哥哥的弟弟的儿子"})
		wantResult := []string{"堂哥", "堂弟", "哥哥", "弟弟", "自己"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "爸爸的妹妹的女儿的老公"})
		wantResult := []string{"姑表姐夫", "姑表妹夫"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "表姐的哥哥", "sex": 1})
		wantResult := []string{"姑表哥", "舅表哥"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "表姐的妹妹", "sex": 1})
		wantResult := []string{"姑表姐", "姑表妹", "舅表姐", "舅表妹"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "弟弟的表哥", "sex": 1})
		wantResult := []string{"姑表哥", "姑表弟", "舅表哥", "舅表弟"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "姐姐的老公的姐姐的老公"})
		wantResult := []string{"姊妹姻姊妹壻"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "哥哥的弟弟的爸爸的儿子", "sex": 1})
		wantResult := []string{"哥哥", "弟弟", "自己"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "爱人", "sex": 1})
		wantResult := []string{"老婆"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "爱人的爱人", "sex": 1})
		wantResult := []string{"自己"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "老婆的爱人", "sex": 1})
		wantResult := []string{"自己"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "老婆的老公", "sex": 1})
		wantResult := []string{"自己"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "老婆的外孙的姥爷", "sex": 1})
		wantResult := []string{"自己"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "老公的老婆的儿子的爸爸", "sex": 0})
		wantResult := []string{"老公"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "堂兄弟的孩子", "sex": 1})
		wantResult := []string{"堂侄", "堂侄女"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "儿子的爸爸的妈妈", "sex": 1})
		wantResult := []string{"妈妈"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "岳母的配偶的孩子的爸爸"})
		wantResult := []string{"岳父"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("call", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "爸爸的妈妈的姐姐的儿子"})
		wantResult := []string{"姨伯父", "姨叔父"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})
}

func TestTarget(t *testing.T) {
	rel := relationship.NewRelationship()

	t.Run("target", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "我", "target": "爸爸"})
		wantResult := []string{"儿子", "女儿"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("target", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "老公的父母", "target": "孩子"})
		wantResult := []string{"爷爷", "奶奶"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("target", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "爱人", "target": "娘家侄子"})
		wantResult := []string{"姑丈"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})
}

func TestReverse(t *testing.T) {
	rel := relationship.NewRelationship()

	t.Run("reverse", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "爸爸的舅舅", "sex": 0, "reverse": true})
		wantResult := []string{"甥孙女"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("reverse", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "岳母", "target": "女儿", "reverse": true})
		wantResult := []string{"外孙女"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("reverse", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "姑妈", "target": "舅妈", "reverse": true})
		wantResult := []string{"兄弟眷兄弟妇"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("reverse", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "舅妈", "target": "女儿", "reverse": true})
		wantResult := []string{"姑甥孙女", "姑甥外孙女"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("reverse", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "外婆", "target": "女婿", "reverse": true})
		wantResult := []string{"外曾孙女婿", "外曾外孙女婿"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})
}

func TestFilter(t *testing.T) {
	rel := relationship.NewRelationship()

	t.Run("filter", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "内侄"})
		wantResult := []string{"舅侄", "舅侄女"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})
}

func TestTypeChain(t *testing.T) {
	rel := relationship.NewRelationship()

	t.Run("typeChain", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "舅爷爷", "type": "chain"})
		wantResult := []string{"爸爸的妈妈的兄弟"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("typeChain", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "妻儿", "type": "chain"})
		wantResult := []string{"老婆", "儿子", "女儿"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("typeChain", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "父母", "target": "老公", "type": "chain"})
		wantResult := []string{"老婆的爸爸", "老婆的妈妈"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})
}

func TestTypePair(t *testing.T) {
	rel := relationship.NewRelationship()

	t.Run("typePair", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "舅妈", "target": "哥哥", "type": "pair"})
		wantResult := []string{"舅甥"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("typePair", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "舅妈", "target": "外婆", "type": "pair"})
		wantResult := []string{"婆媳"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("typePair", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "舅妈", "target": "二舅", "type": "pair"})
		wantResult := []string{"伯媳", "叔嫂", "夫妻"}
		if !SameElementsSlice(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("typePair", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "堂哥", "target": "叔叔", "type": "pair"})
		wantResult := []string{"伯侄", "叔侄", "父子"}
		if !SameElementsSlice(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})
}

func TestAge(t *testing.T) {
	rel := relationship.NewRelationship()

	t.Run("age", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "妈妈的二哥"})
		wantResult := []string{"二舅"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("age", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "爸爸的二哥"})
		wantResult := []string{"二伯"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("age", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "二舅妈", "target": "三舅"})
		wantResult := []string{"二嫂"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("age", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "爸爸的二爸"})
		wantResult := []string{"二爷爷"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("age", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "姑姑", "target": "叔叔", "optimal": true})
		wantResult := []string{"姐姐", "妹妹"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("age", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "大舅", "target": "二舅的儿子"})
		wantResult := []string{"伯父"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("age", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "二舅妈", "target": "二舅", "type": "pair"})
		wantResult := []string{"夫妻"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("age", func(t *testing.T) {
		gotResult := rel.Relationship(map[string]any{"text": "二舅妈", "target": "大舅", "type": "pair"})
		wantResult := []string{"伯媳"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})
}

func TestExpression(t *testing.T) {
	rel := relationship.NewRelationship()

	t.Run("expression", func(t *testing.T) {
		gotResult := rel.Relationship("外婆和奶奶之间是什么关系？")
		wantResult := []string{"儿女亲家"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("expression", func(t *testing.T) {
		gotResult := rel.Relationship("妈妈应该如何称呼姑姑")
		wantResult := []string{"姑子"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("expression", func(t *testing.T) {
		gotResult := rel.Relationship("姑奶奶是什么关系")
		wantResult := []string{"爸爸的爸爸的姐妹"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("expression", func(t *testing.T) {
		gotResult := rel.Relationship("姑奶奶和爸爸是什么关系")
		wantResult := []string{"姑侄"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("expression", func(t *testing.T) {
		gotResult := rel.Relationship("我应该叫外婆的哥哥什么？")
		wantResult := []string{"舅外公"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("expression", func(t *testing.T) {
		gotResult := rel.Relationship("七舅姥爷应该叫我什么？")
		wantResult := []string{"甥外孙", "甥外孙女"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("expression", func(t *testing.T) {
		gotResult := rel.Relationship("舅公是什么关系？")
		wantResult := []string{"爸爸的妈妈的兄弟", "妈妈的妈妈的兄弟", "老公的妈妈的兄弟"}
		if !SameElementsSlice(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("expression", func(t *testing.T) {
		gotResult := rel.Relationship("舅妈如何称呼外婆？")
		wantResult := []string{"婆婆"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})

	t.Run("expression", func(t *testing.T) {
		gotResult := rel.Relationship("外婆和奶奶之间是什么关系？")
		wantResult := []string{"儿女亲家"}
		if !reflect.DeepEqual(gotResult, wantResult) {
			t.Errorf("got %v, want %v", gotResult, wantResult)
		}
	})
}
