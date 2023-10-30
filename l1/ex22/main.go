// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

var (
	a = int64(300000000 * 1000000)
	b = int64(200000000 * 1000000)
)

type BigInt big.Int

func NewBigInt(numb int64) *BigInt {
	return (*BigInt)(big.NewInt(numb))
}

func (i *BigInt) Add(i2 *BigInt) *BigInt {
	c := new(big.Int).Add((*big.Int)(i), (*big.Int)(i2))
	return (*BigInt)(c)
}

func (i *BigInt) Sub(i2 *BigInt) *BigInt {
	c := new(big.Int).Sub((*big.Int)(i), (*big.Int)(i2))
	return (*BigInt)(c)
}

func (i *BigInt) Mul(i2 *BigInt) *BigInt {
	c := new(big.Int).Mul((*big.Int)(i), (*big.Int)(i2))
	return (*BigInt)(c)
}

func (i *BigInt) Div(i2 *BigInt) *BigInt {
	c := new(big.Int).Div((*big.Int)(i), (*big.Int)(i2))
	return (*BigInt)(c)
}

func (i *BigInt) String() string {
	return (*big.Int)(i).String()
}

func main() {
	n1 := NewBigInt(a)
	n2 := NewBigInt(b)
	fmt.Println("n1 = ", n1)
	fmt.Println("n2 = ", n2)

	add := n1.Add(n2)
	sub := n1.Sub(n2)
	mul := n1.Mul(n2)
	div := n1.Div(n2)
	fmt.Println("add = ", add)
	fmt.Println("sub = ", sub)
	fmt.Println("mul = ", mul)
	fmt.Println("div = ", div)

}
