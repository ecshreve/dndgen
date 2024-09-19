package utils

type IndxCache struct {
	IndxToId map[string]int
	IdToIndx map[int]string
}

func NewIndxCache() *IndxCache {
	return &IndxCache{
		IndxToId: make(map[string]int),
		IdToIndx: make(map[int]string),
	}
}