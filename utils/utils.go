package utils

import (
	"log"
)


func CreatePointers(content *[]string) ([]interface{}) {
	point 		:= make([]interface{}, len(*content))
	for i		:= range *content {
		point[i] = &(*content)[i]
	}
	
	return point
}

func ConvetIntoMap(slices *[][]string, columns []string) []map[string]string {
	newMaps 	:= []map[string]string{}

	for _, data := range *slices {
		newMap 	:= map[string]string{}
		for i, colName := range columns {
			newMap[colName] = data[i]
		}
		newMaps = append(newMaps, newMap)
	}
	return newMaps
}
