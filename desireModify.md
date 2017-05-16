# 设计修改文档

## 功能需求修改

1. 一级缓存获取IP列表方式从通过资产修改为通过CNAME

## 需求分析
这个修改对于后期ip进行purge操作没有影响，所以再尽量不动后面操作的基础上进行修改

## 具体设计修改

* 增加一个CNAMEIPMap 用于存储CNAME对应的IP列表(../g/g.co)
* 增加一个UpdateCNAMEIPMap函数用来更新CNAMEIPMap(../service/portal.go)
* 增加一个定时任务来更新CNAMEIPMap(../cron/cron.go)
* 修改ServicePurgeData：增加一个CNAME字段用来一级缓存获取IP列表(../g/g.go)
* 修改buildServicePurgeData函数来适配新的结构
* 修改buildIPPurgeData函数来适配上一个函数和原来结果结构

### CNAMEIPMap结构
* key : cname `string`
* value: ips  `[]string`
### UpdateCNAMEIPMap函数流程思路

```flow
st=>start:Start
e=>end:End
op1=>operation:Get IPtoSNMap and Get SNtoIPsMap
cnamecircle=>condition:CNAME list circle finish?
op2=>operation:Get CNAME IPs
ipscircle=>condirion:ipscircle list finished?
op3=>operation:get the IP list of sn(priv+pub)
op4=>operation:append this CNAMEipList by the first ip
op5=>operation:set the CNAME map with the CNAMEipList
st->op1->cnamecircle
cnamecircle(yes)->e
cnamecircle(no)->op2->ipscircle
ipscircle(yes)->op5->cnamecircle
ipscircle(no)->op3->op4->ipscircle
```
### 根据URL确定Purge操作方法

* 首先判断cname中
