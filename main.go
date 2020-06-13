package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// main ...
func main() {
	var (
		err             error
		list            []Tax
		bs              []byte
		baseNum         float64
		totalIncome     float64 //累计税前收入
		totalInsurance  float64 //累计社保
		totalDeduction  float64 //累计专项扣除
		totalBaseNum    float64 //累计扣除基数
		taxIncome       float64 //计税收入
		taxRate         float64 //税率
		deductionNumber float64 //速算扣除数
		lastTotalTax    float64 //上月累计应纳税额
		totalTax        float64 //累计应纳数额
		tax             float64 //当月纳税额
		finalIncome     float64 //当月税后收入
	)
	baseNum = 5000 //社保基数
	if bs, err = ioutil.ReadFile("./tax.json"); err != nil {
		return
	}
	if err = json.Unmarshal(bs, &list); err != nil {
		return
	}

	for k, v := range list {
		totalIncome += v.Income
		totalInsurance += v.Insurance
		totalDeduction += v.Deduction
		totalBaseNum += baseNum
		taxIncome = totalIncome - totalInsurance - totalDeduction - totalBaseNum
		taxRate, deductionNumber = getTaxRate(taxIncome)
		totalTax = taxIncome*taxRate - deductionNumber
		tax = totalTax - lastTotalTax
		finalIncome = v.Income - v.Insurance - tax
		lastTotalTax = totalTax
		fmt.Printf("%d月份：累计税前收入%.4f，累计社保%.4f，累计专项扣除%.4f，累计扣除基数%.4f，计税收入%.4f，税率%.4f，速算扣除数%.4f，累计应纳税额%.4f——当月纳税额%.4f，当月税后收入%.4f\n", k+1, totalIncome, totalInsurance, totalDeduction, totalBaseNum, taxIncome, taxRate, deductionNumber, totalTax, tax, finalIncome)
	}
}

type Tax struct {
	Income    float64 `json:"income"`
	Insurance float64 `json:"insurance"`
	Deduction float64 `json:"deduction"`
}

func getTaxRate(taxIncome float64) (rate float64, num float64) {
	if taxIncome < 36000 {
		rate = 0.03
		num = 0
	} else if taxIncome < 144000 {
		rate = 0.1
		num = 2520
	} else if taxIncome < 300000 {
		rate = 0.2
		num = 16920
	} else if taxIncome < 420000 {
		rate = 0.25
		num = 31920
	} else if taxIncome < 660000 {
		rate = 0.3
		num = 52920
	} else if taxIncome < 960000 {
		rate = 0.35
		num = 85920
	} else {
		rate = 0.45
		num = 181920
	}
	return
}
