/*
模式数据
*/

package modedata

import (
	"maps"

	"github.com/young2j/relationship/relationship-go/datamap"
	"github.com/young2j/relationship/relationship-go/locale"
	"github.com/young2j/relationship/relationship-go/options"
)

// 获取模式数据
func GetModeData(mode options.OptionMode) map[string][]string {
	data := maps.Clone(datamap.InitMap())

	if mode == options.OptionModeGuangDong {
		for k, v := range locale.Guangdong.Data {
			if dat, ok := data[k]; ok {
				data[k] = append(v, dat...)
			} else {
				data[k] = v
			}
		}

	} else if mode == options.OptionModeNorth {
		for k, v := range locale.North.Data {
			if dat, ok := data[k]; ok {
				data[k] = append(v, dat...)
			} else {
				data[k] = v
			}
		}
	}

	return data
}
