package utils

// type sortByAge []models.TrxItemsMongo

// func (s sortByAge) Len() int           { return len(s) }
// func (s sortByAge) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
// func (s sortByAge) Less(i, j int) bool { return s[i].MerchantKey != s[j].MerchantKey }

// //Slice grouping
// func GroupingBatch(list []models.TrxItemsMongo) [][]models.TrxItemsMongo {
// 	sort.Sort(sortByAge(list))
// 	returnData := make([][]models.TrxItemsMongo, 0)
// 	i := 0
// 	var j int
// 	for {
// 		if i >= len(list) {
// 			break
// 		}
// 		for j = i + 1; j < len(list) && list[i].DeviceID == list[j].DeviceID; j++ {
// 		}

// 		returnData = append(returnData, list[i:j])
// 		i = j
// 	}
// 	return returnData
// }
