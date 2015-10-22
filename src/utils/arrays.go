package utils

func Contains(list []uint64, e uint64) bool {
    for _, a := range list {
        if a == e {
            return true
        }
    }
    return false
}