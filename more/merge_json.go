package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// mergeJSONStrings 合并两个JSON字符串
// 规则：以第一个JSON的字段值为准，第二个JSON补充第一个不存在的字段
// 输出结果是map直接序列化的，无法保证字段顺序
func mergeJSONStrings(jsonStr1, jsonStr2 string) (string, error) {
	// 1. 解析第一个JSON为map（保留原有字段值）
	var map1 map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr1), &map1); err != nil {
		return "", fmt.Errorf("解析第一个JSON失败: %v", err)
	}
	// 2. 解析第二个JSON为map（补充新字段）
	var map2 map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr2), &map2); err != nil {
		return "", fmt.Errorf("解析第二个JSON失败: %v", err)
	}
	// 3. 合并map：遍历第二个map，补充第一个map中不存在的字段
	for key, value := range map2 {
		if _, exists := map1[key]; !exists {
			map1[key] = value
		}
	}
	// 4. 序列化为格式化的JSON字符串（添加缩进，符合示例格式）
	mergedJSON, err := json.MarshalIndent(map1, "", "\t")
	if err != nil {
		return "", fmt.Errorf("序列化合并后的JSON失败: %v", err)
	}
	// 替换JSON默认的null为空对象（示例中data是{}而非null）
	result := strings.ReplaceAll(string(mergedJSON), "null", "{}")
	return result, nil
}

// mergeJSON 递归合并两个JSON对象（map[string]interface{}）
// 规则：以map1为主，map2补充字段；嵌套对象递归合并
func mergeJSON(map1, map2 map[string]interface{}) map[string]interface{} {
	// 遍历map2的所有字段
	for key, value2 := range map2 {
		value1, exists := map1[key]
		// map1中无该字段，直接追加map2的字段
		if !exists {
			map1[key] = value2
			continue
		}
		// 若map1和map2中该字段都是对象（map类型），则递归合并
		m1Obj, ok1 := value1.(map[string]interface{})
		m2Obj, ok2 := value2.(map[string]interface{})
		if ok1 && ok2 {
			map1[key] = mergeJSON(m1Obj, m2Obj)
		}
		// 若类型不一致或不是对象，保留map1的原值（不覆盖）
	}
	return map1
}
