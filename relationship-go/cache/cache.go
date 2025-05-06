package cache

import (
	"maps"

	"github.com/young2j/relationship/relationship-go/data"
)

func GetCacheData(modeData map[string][]string) map[string][]string {
	_hash := maps.Clone(modeData)
	for sel, inputNames := range data.Input {
		if names, ok := _hash[sel]; ok {
			_hash[sel] = append(names, inputNames...)
		} else {
			_hash[sel] = inputNames
		}
	}

	for sel, sortNames := range data.Sort {
		if names, ok := _hash[sel]; ok {
			_hash[sel] = append(names, sortNames...)
		} else {
			_hash[sel] = sortNames
		}
	}

	cacheData := make(map[string][]string)
	for sel, names := range _hash {
		for _, name := range names {
			if sels, ok := cacheData[name]; ok {
				cacheData[name] = append(sels, sel)
			} else {
				cacheData[name] = []string{sel}
			}
		}
	}

	return cacheData
}
