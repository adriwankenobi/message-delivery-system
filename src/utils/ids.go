package utils

/**
	GENERATES UNIQUE ID
*/
var currentId uint64 = 0
func NextId() uint64 {

	currentId++
	return currentId
}