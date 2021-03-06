# 个人所得税计算工具

使用方法：

1.在tax.json文件中配置好各项值income（税前收入）insurance（各项社会保险费和公积金）deduction（专项附加扣除）

2.执行go run main.go

例子：
假如有张三每月收入为20000，各项社会保险费和公积金共3000，专项扣除共1500，已经发放工资至第5月份，json内容应如下。

```
[
  {
    "income": 20000,
    "insurance": 3000,
    "deduction": 1500
  },
  {
    "income": 20000,
    "insurance": 3000,
    "deduction": 1500
  },
  {
    "income": 20000,
    "insurance": 3000,
    "deduction": 1500
  },
  {
    "income": 20000,
    "insurance": 3000,
    "deduction": 1500
  },
  {
    "income": 20000,
    "insurance": 3000,
    "deduction": 1500
  }
]
```


个人所得税计算方式：

第一个月工资的个税=[(第一个月税前工资-5000元-各项社会保险费和公积金-专项扣除)*适用税率-速算扣除数]

第二个月工资的个税=[(第一个月税前工资+第二个月税前工资）-（5000元 * 2）-（各项社会保险费和公积金 * 2）-（专项扣除 * 2）) * 适用税率 - 速算扣除数] - 第一个月工资个税

第三个月工资的个税=[(第一个月税前工资+第二个月税前工资+第三个月税前工资）-（5000元 * 3）-（各项社会保险费和公积金 * 3）-（专项扣除 * 3）) * 适用税率 - 速算扣除数] - (第一个月工资个税+第二个月工资个税)

以此类推...

注意：如果当年换过公司，那么新公司算税款一般是重新开始，所以到年底会需要补税

|  级数   | 累计预扣预缴应纳税所得额  |  （%）   | 速算扣除数  |
|  ----  | ----  | ----  | ----  |
| 1| 不超过36,000元的部分 | 3  | 0 |
| 2  | 超过36,000元至144,000元的部分 | 10  | 2520 |
| 3  | 超过144,000元至300,000元的部分 | 20  | 16920 |
| 4  | 超过300,000元至420,000元的部分 | 25  | 31920 |
| 5  | 超过420,000元至660,000元的部分 | 30  | 52920 |
| 6  | 超过660,000元至960,000元的部分 | 35  | 85920 |
| 7  | 超过960,000元的部分 | 45 | 181920 |
