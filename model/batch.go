package model

import "project/dingdangke-dataoke/entity"

func batchResult(size int, data []entity.DTKResult2) [][]entity.DTKResult2 {
	batch := [][]entity.DTKResult2{}
	for i, n := 0, len(data); i < n; i += size {
		end := i + size
		if end > n {
			batch = append(batch, data[i:])
		} else {
			batch = append(batch, data[i:end])
		}
	}
	return batch
}