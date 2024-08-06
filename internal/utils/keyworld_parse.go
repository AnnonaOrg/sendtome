package utils

import (
	"strings"
)

func KeyworldListParse(keyworldListStr string) []string {
	if len(keyworldListStr) <= 0 {
		return nil
	}

	retArr := make([]string, 0)
	listText := keyworldListStr
	list := strings.Split(listText, "|")
	for _, v := range list {
		vc := v
		if len(vc) > 0 {
			retArr = append(retArr, vc)
		}
	}
	return retArr
}

//解析 k>>kc|k1>>kc1 到 map
func KeyworldListParseToMap(keyworldText string) map[string]string {
	keyworldList := KeyworldListParse(keyworldText)
	retMap := make(map[string]string, 0)
	for _, v := range keyworldList {
		if len(v) > 0 {
			if strings.Contains(v, ">>") {
				if vk, vv, isOk := strings.Cut(v, ">>"); isOk {
					if len(vk) > 0 {
						retMap[vk] = vv
					}
				}
			} else {
				retMap[v] = ""
			}
		}
	}
	return retMap
}

// 检查屏蔽关键词，关键词，存在屏蔽词(keyworldFilter) 返回false，存在订阅关键词(keyworldList)或无订阅关键词词 返回true
func FeedKeyworldCheck(msgText, feedKeyworldFilter, feedKeyworldList string) (retText, retFilter string, retBool bool) {
	retBool = true

	keyworldFilter := KeyworldListParse(feedKeyworldFilter)
	for _, v := range keyworldFilter {
		vc := v
		if strings.Contains(strings.ToLower(msgText), strings.ToLower(vc)) {
			retFilter = vc
			retBool = false
			return
		}
	}

	keyworldList := KeyworldListParse(feedKeyworldList)
	if len(keyworldList) <= 0 {
		retText = "无订阅词限定"
		retBool = true
		return
	} else {
		retBool = false
	}
	for _, v := range keyworldList {
		vc := v
		if strings.Contains(msgText, vc) {
			retText = vc
			retBool = true
			return
		}
	}
	return
}

// 文案级关键词替换
func FeedKeyworldReplace(msgText, feedKeyworldReplace string) string {
	if len(msgText) <= 0 {
		return ""
	}
	retText := msgText
	keyworldListMap := KeyworldListParseToMap(feedKeyworldReplace)
	for k, v := range keyworldListMap {
		if len(k) > 0 {
			retText = strings.ReplaceAll(retText, k, v)
		}
	}
	return retText
}
