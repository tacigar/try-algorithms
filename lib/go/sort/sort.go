package sort

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func IsSorted(data Interface) bool {
	for i := 0; i < data.Len()-1; i++ {
		if data.Less(i+1, i) {
			return false
		}
	}
	return true
}

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
