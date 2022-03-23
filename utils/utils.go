package utils

import "sort"

func ConvetIntoMap(slices [][]string, columns []string) []map[string]string {
	newMaps 	:= []map[string]string{}

	for _, data := range slices {
		newMap 	:= map[string]string{}
		for i, colName := range columns {
			newMap[colName] = data[i]
		}
		newMaps = append(newMaps, newMap)
	}
	_, exist := newMaps[0]["id"]
	if exist {
  		sort.Slice(newMaps, func(i, j int) bool { return newMaps[i]["id"] < newMaps[j]["id"]})
	}
	return newMaps
}
