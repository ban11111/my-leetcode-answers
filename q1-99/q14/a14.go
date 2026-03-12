package q14

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	for index := 0; ; index++ {
		var letter *byte
		for _, str := range strs {
			if len(str) < index+1 {
				return str[:index]
			}
			if letter == nil {
				let := str[index]
				letter = &let
			}
			if str[index] != *letter {
				return str[:index]
			}
		}
	}
}

// todo 多核优化
//func longestCommonPrefixMultiCore(strs []string) string {
//	if len(strs) <= 0 {
//		return ""
//	}
//	numCPU := runtime.NumCPU()
//	var indexCh = make(chan int, numCPU)
//	for i := 0; i < numCPU; i++ {
//		go compare(indexCh)
//	}
//	go func() {
//		for index := 0; ; index++ {
//			indexCh <- index
//		}
//	}()
//
//	var result []byte
//	for index := 0; ; index++ {
//		var letter *byte
//		for _, str := range strs {
//			if len(str) < index+1 {
//				return string(result)
//			}
//			if letter == nil {
//				let := str[index]
//				letter = &let
//			}
//			if str[index] != *letter {
//				return string(result)
//			}
//		}
//		result = append(result, *letter)
//	}
//}
//
//var failedAt int = -1
//
//func compare(indexCh, failedIndexCh chan int, strs []string) {
//	for {
//		select {
//		case index := <-indexCh:
//			var letter *byte
//			for _, str := range strs {
//				if failedAt >= 0 && index >= failedAt {
//					return false
//				}
//				if letter != nil && str[index] != *letter {
//					return false
//				}
//			}
//			return true
//			//case failedAt := <- failedIndexCh:
//		}
//	}
//
//}
