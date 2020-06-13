# tax
##个人所得税计算工具

使用方法：

1.在tax.json文件中配置好各项值income（税前收入）insurance（各项社会保险费）deduction（专项附加扣除）

2.执行go run main.go

例子：
假如有张三每月收入为20000，社会保险费3000，专项扣除1500，已经发放工资至第5月份，json内容应如下。

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