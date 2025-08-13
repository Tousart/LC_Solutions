package main

import (
	"fmt"
	"math/rand/v2"
)

type RandomizedSet struct {
	valInd map[int]int
	set    []int
}

func Constructor() RandomizedSet {
	valInd := make(map[int]int)
	set := []int{-1}
	return RandomizedSet{
		valInd: valInd,
		set:    set,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	ind := this.valInd[val]
	if ind == 0 {
		this.valInd[val] = len(this.set)
		this.set = append(this.set, val)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	ind := this.valInd[val]
	if ind != 0 {
		n := len(this.set)

		if n > 2 {
			last := this.set[n-1]
			this.set[ind] = last
			this.valInd[this.set[ind]] = ind
		}

		this.set = this.set[:n-1]
		delete(this.valInd, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	ind := rand.IntN(len(this.set)-1) + 1
	return this.set[ind]
}

func main() {
	rs := Constructor()
	fmt.Println(rs.Insert(0))
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Remove(0))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.GetRandom())
}
