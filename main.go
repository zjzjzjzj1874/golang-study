package main

import (
	"fmt"
	"sort"
)

//5
//b a
//c a
//d c
//e c
//f d
//c

type Trees struct {
	Value string
	Child []*Trees
}

func main() {
	n := 0
	fmt.Scan(&n)

	node, parent := "", ""
	nodes := make([]*Trees, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %s", &node, &parent)

		child := &Trees{Value: node}
		findP := false
		for j, tree := range nodes {
			if tree.Value == parent { // 找到了父节点，将子节点加入
				findP = true
				nodes[j].Child = append(nodes[j].Child, child)
			}
		}
		if !findP { // 没找到父节点，新建父节点
			nodes = append(nodes, &Trees{
				Value: parent,
				Child: []*Trees{{Value: node}},
			})
		}
	}
	tm := make(map[string]*Trees)
	for _, treeNode := range nodes {
		tm[treeNode.Value] = treeNode
	}

	for s1, t1 := range tm {
		for _, t2 := range tm {
			// 深度优先，这里可能不止一层
			for i, child := range t2.Child {
				if t1.Value == child.Value {
					t2.Child[i].Child = append(t2.Child[i].Child, t1)
					delete(tm, s1)
				}
			}
		}
	}

	target := ""
	fmt.Scan(&target)

	// 先将树合并
	ans := make([]string, 0, len(nodes))
	var printTree func(node, target *Trees, canAdd bool)
	printTree = func(node, target *Trees, canAdd bool) {
		if node == nil {
			return
		}
		if canAdd {
			ans = append(ans, node.Value)
			for i := range node.Child {
				printTree(node.Child[i], target, true)
			}
		} else {
			if node.Value == target.Value {
				for i := range node.Child {
					printTree(node.Child[i], target, true)
				}
			} else {
				for i := range node.Child {
					printTree(node.Child[i], target, false)
				}
			}
		}
	}

	printTree(nodes[0], &Trees{Value: target}, false)
	sort.Strings(ans)
	for _, an := range ans {
		fmt.Println(an)
	}
}

//	var CardMap = map[string]int{
//		"2":  2,
//		"3":  3,
//		"4":  4,
//		"5":  5,
//		"6":  6,
//		"7":  7,
//		"8":  8,
//		"9":  9,
//		"10": 10,
//		"J":  11,
//		"Q":  12,
//		"K":  13,
//		"A":  14,
//	}
//
//	var Val2CardMap = map[int]string{
//		2:  "2",
//		3:  "3",
//		4:  "4",
//		5:  "5",
//		6:  "6",
//		7:  "7",
//		8:  "8",
//		9:  "9",
//		10: "10",
//		11: "J",
//		12: "Q",
//		13: "K",
//		14: "A",
//	}
//
//	type Card struct {
//		Value int // 牌面值
//		Count int // 牌数量
//	}
//
// 顺子
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	msg, _ := reader.ReadString('\n')
//
//	msgs := strings.Fields(msg)
//
//	carMap := make(map[string]int)
//	for _, card := range msgs {
//		carMap[card]++
//	}
//	cards := make([]Card, 0, len(carMap))
//	for key, count := range carMap {
//		cards = append(cards, Card{
//			Value: CardMap[key],
//			Count: count,
//		})
//	}
//	// 整理手牌
//	sort.Slice(cards, func(i, j int) bool {
//		return cards[i].Value > cards[j].Value // 按照值从大到小整理；
//	})
//	if cards[len(cards)-1].Value == 2 { // 最后一张手牌为2，不能组成顺子，直接移除
//		cards = cards[:len(cards)-1]
//	}
//
//	shunzi := make([]string, 0, len(cards))
//	for {
//		// 先整理牌  ：后续补充退出条件
//		sort.Slice(cards, func(i, j int) bool {
//			return cards[i].Value > cards[j].Value // 按照值从大到小整理；
//		})
//		// 双指针, 顺子长度必须连续且>=5
//		l := 0
//		noShunzi := true
//		r := 1
//		for ; r < len(cards); r++ {
//			// 保证是连续的顺子 => 这里不用判断数量，如果牌数量为0，把他从手牌中移除
//			if /*cards[r].Count > 0 &&*/ r-l == cards[l].Value-cards[r].Value {
//
//			} else {
//				if r-1-l >= 5 {
//					noShunzi = false
//					break
//				} else {
//					l = r // 重新计算顺子序列
//				}
//			}
//		}
//		if noShunzi && r-1-l == cards[l].Value-cards[r-1].Value {
//			noShunzi = false
//		}
//		if !noShunzi {
//			//	结算顺子
//			shun := fmt.Sprintf("%s", Val2CardMap[cards[r-1].Value])
//			cards[r-1].Count--
//			for i := r - 2; i >= l; i-- {
//				shun = fmt.Sprintf("%s %s", shun, Val2CardMap[cards[i].Value])
//				cards[i].Count--
//			}
//			shunzi = append(shunzi, shun)
//
//			// 整理牌
//			for i := 0; i < len(cards); {
//				if cards[i].Count <= 0 {
//					if i == len(cards)-1 {
//						cards = cards[:len(cards)-1]
//					} else {
//						cards = append(cards[:i], cards[i+1:]...)
//					}
//				} else {
//					i++
//				}
//			}
//		}
//
//		if noShunzi || len(cards) < 5 {
//			break
//		}
//	}
//
//	for i := range shunzi {
//		fmt.Println(shunzi[len(shunzi)-1-i])
//	}
//
//}

// 字符串摘要
//func main() {
//	s := ""
//	fmt.Scan(&s)
//
//	s = strings.ToLower(s) // 将所有字母转成小写
//	ns := []byte{}
//	for i := range s { // 移除非字母符号
//		if s[i] < 'a' || s[i] > 'z' {
//			continue
//		}
//		ns = append(ns, s[i])
//	}
//	s = string(ns)
//
//	ans := make([]string, 0)
//	alphaMap := make(map[byte]int) // 使用hash表记录不连续的字符，后续出现的次数
//	// 使用双指针，找到不相等的值
//	l := 0
//	for r := 0; r < len(s); r++ {
//		if s[r] != s[l] { // 有不相同的字符出现，
//			if r-l > 1 { // 说明是连续出现
//				ans = append(ans, fmt.Sprintf("%s%d", string(s[l]), r-l))
//				if _, ok := alphaMap[s[l]]; ok { // 记录后续出现的元素数量
//					alphaMap[s[l]] = r - l
//				}
//				l = r
//			} else { // 不连续出现，写入hash表中
//				alphaMap[s[l]] = 0
//				l = r
//			}
//		}
//	}
//	if l < len(s) {
//		ans = append(ans, fmt.Sprintf("%s%d", string(s[l]), len(s)-l))
//		if _, ok := alphaMap[s[l]]; ok { // 记录后续出现的元素数量
//			alphaMap[s[l]] = len(s) - l
//		}
//	}
//	for key, val := range alphaMap {
//		ans = append(ans, fmt.Sprintf("%s%d", string(key), val))
//	}
//
//	sort.Slice(ans, func(i, j int) bool {
//		if ans[i][1:] == ans[j][1:] {
//			return ans[i] < ans[j]
//		}
//		return ans[i][1:] > ans[j][1:]
//	})
//	for i := range ans {
//		fmt.Print(ans[i])
//	}
//	fmt.Printf("\n")
//}
