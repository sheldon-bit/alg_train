package main

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//示例 1：
//输入：strs = ["flower","flow","flight"]
//输出："fl"

//示例 2：
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。

//提示：
//1 <= strs.length <= 200
//0 <= strs[i].length <= 200
//strs[i] 仅由小写英文字母组成

func GongZ(strs []string) string {
	strMap := map[string]int{}
	for _, i := range strs {
		strMap[i] = len(i)
	}
	strs2 := []string{}
	for i, j := range strMap {
		if j == 0 {
			return ""
		}
		if len(strs2) == 0 {
			strs2 = append(strs2, i)
			continue
		}
		if strMap[strs2[0]] >= j {
			strs2 = append([]string{i}, strs2...)
		} else {
			strs2 = append(strs2, i)
		}
	}
	strslen := len(strs2)
	card := strs2[0]
	i := 0
	for {
		if i+1 < strslen {
			lenj := len(card)
			str3 := strs2[i+1]
			for {
				if card != str3[:len(card)] {
					lenj = lenj - 1
					card = card[:lenj]
				} else {
					i++
					break
				}

			}

		} else {
			break
		}
	}
	return card

}

func main() {
	//strs := []string{"flower", "flow", "flight"}
	strs := []string{"dog", "racecar", "car"}
	GongZ(strs)
}
