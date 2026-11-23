package main

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */

type ProductOfNumbers struct {
	Product []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		Product: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.Product = []int{1}
	} else {
		this.Product = append(this.Product, this.Product[len(this.Product)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	prodLen := len(this.Product)
	if k >= prodLen {
		return 0
	}
	return this.Product[prodLen-1] / this.Product[prodLen-1-k]
}
