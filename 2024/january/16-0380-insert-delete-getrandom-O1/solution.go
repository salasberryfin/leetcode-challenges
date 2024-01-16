package main

import "math/rand"

type RandomizedSet struct {
	Vals map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{Vals: make(map[int]bool)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Vals[val]; !ok {
		this.Vals[val] = true
		return true
	}

	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.Vals[val]; ok {
		delete(this.Vals, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	randomized := rand.Intn(len(this.Vals))
	var counter int
	for k := range this.Vals {
		if counter == randomized {
			return k
		}
		counter++
	}

	return -1
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
