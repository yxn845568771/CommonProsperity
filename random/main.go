package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// type Option func(o *options)

// func Employees(e map[string]float64) Option {
// 	return func(o *options) {
// 		o.employees = e
// 	}
// }
//
// func Models(m []string) Option {
// 	return func(o *options) {
// 		o.models = m
// 	}
// }

// type options struct {
// 	employees map[string]float64
// 	models    []string
// }
//
// type design struct {
// 	opts options
// }
//
// type avg interface {
// 	GetMotherLen() int
// 	GetMotherData(int) interface{}
// 	GetSonLen() int
// 	GetSonDataAndPower(int) (interface{}, float64)
// }
//
// func
//
// func init() {
// 	design.employees = map[string]int{
// 		"晓南": 1,
// 		"晓霖": 1,
// 		"晓炜": 1,
// 		"文建": 1,
// 		"均宝": 1,
// 		"敬朋": 1,
// 		"先杰": 1,
// 	}
// 	design.models = []string{
// 		"建造者", "抽象工厂",
// 		"适配器", "桥接", "组合", "装饰", "代理", "外观", "享元",
// 		"迭代器", "中介者", "观察者", "策略", "模板方法", "命令",
// 	}
// }

func main() {
	employees := []string{"晓南", "晓霖", "晓炜", "文建", "均宝", "敬朋", "先杰"}
	models := []string{
		"建造者", "抽象工厂",
		"适配器", "桥接", "组合", "装饰", "代理", "外观", "享元",
		"迭代器", "中介者", "观察者", "策略", "模板方法", "命令",
	}
	for _, employee := range employees {
		fmt.Print(employee + "负责：")
		for i := 0; i < 2; i++ {
			l := len(models)
			data, err := rand.Int(rand.Reader, big.NewInt(int64(l)))
			if err != nil {
				return
			}
			j := data.Int64()
			fmt.Print(models[j] + " ")
			models = append(models[:j], models[j+1:]...)
		}
		fmt.Println()
	}
}
