
type MyHashSet struct {
	list []int
}


func Constructor() MyHashSet {
	list = make([]int, 0)
}


func (this *MyHashSet) Add(key int)  {
    list = append(list, key)
}


func (this *MyHashSet) Remove(key int)  {
    for i := 0; i < len(this.list); i++ {
		if this.list[i] == key {
			this.list = append(list[:i], list[i+1:]...)
			break
		}
	}
}


func (this *MyHashSet) Contains(key int) bool {
    for i := 0; i < len(this.list); i++ {
		if this.list[i] == key {
			return true
		}
	}
	return false
}
