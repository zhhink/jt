package jsont

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/zhhink/common/file"

	"github.com/zhhink/convert"
)

// JSONT struct
type JSONT struct {
	JSONFileName string
}

// GetValueOfMap get value of key
func GetValueOfMap(m map[string]interface{}, key string) (out interface{}, err error) {
	keys := strings.Split(key, ".")
	for i, k := range keys {
		v, exist := m[k]

		if exist == false {
			return nil, fmt.Errorf("key error, %s is not in json", k)
		}

		if i == len(keys)-1 {
			out = v
			break
		}
		if reflect.TypeOf(v).Kind() != reflect.Map {
			return nil, fmt.Errorf("type error, %s type cannot be marshal", reflect.TypeOf(m).Name())
		}
		m = v.(map[string]interface{})
	}
	return out, nil
}

// FilterItemsFromJSONStr filter json keys
func FilterItemsFromJSONStr(jsonStr string, outputItems string) (out []interface{}, err error) {
	itemsList := strings.Split(outputItems, ",")
	jsonMap, error := convert.StringToMap(jsonStr)
	if error != nil {
		return nil, error
	}
	for _, item := range itemsList {
		value, err := GetValueOfMap(jsonMap, item)
		if err != nil {
			return nil, err
		}
		out = append(out, value)
	}
	return out, nil
}

// FilterItemsFromJSONFile filter json keys of json file
func (j *JSONT) FilterItemsFromJSONFile(outputItems string) {
	var line string
	jsonF := file.Open(j.JSONFileName)
	line = jsonF.ReadLine()
	var waitGroutp = sync.WaitGroup{}
	for {
		if line == "" {
			break
		}
		waitGroutp.Add(1)
		go func(jsonStr string) {
			var out string
			items, err := FilterItemsFromJSONStr(jsonStr, outputItems)
			if err != nil {
				return
			}
			for _, item := range items {
				if reflect.TypeOf(item).Kind() == reflect.Map {
					s, err := convert.MapToJSONString(item)
					if err != nil {
						return
					}
					out = fmt.Sprintf("%s\t%s", out, s)
				} else {
					out = fmt.Sprintf("%s\t%v", out, item)
				}
			}
			fmt.Println(strings.Trim(out, "\t"))
			waitGroutp.Done()
		}(line)
		line = jsonF.ReadLine()
	}
	waitGroutp.Wait()
}
