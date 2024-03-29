# DDD布局

## 原文

基于DDD落地Go微服务开发总结：https://tkstorm.com/ln/ddd-summary-2024

## DDD 理解

DDD 是一种复杂业务（战略设计、战术实施）软件架构方法，**通过分治思想将大型复杂业务完成业务域划分、领域建模确定，指导微服务落地，以便更好地管理业务复杂性并提高代码可维护性**

## DDD 具体步骤

DDD 具体步骤包含: **DDD-战略设计**、**DDD-战术实施**两部份，战略设计聚焦在业务分治，通过梳理业务，领域划分，完成领域建模，战略实施聚焦在领域概念代码映射，完成微服务的编码落地

### DDD - 战略设计（梳理业务，领域划分，完成领域建模）

1. **业务梳理（事件风暴、领域故事分析）**：搞清楚业务流、用例，深入理解业务需求流程和规则
2. **领域划分（业务域、子域划分）**：基于业务流、用例划分业务域、子域，区分哪些是核心域、支持域
3. **领域建模（领域对象梳理）**：明确领域内核心要素,提炼业务域内的**实体、值对象、聚合、聚合根**
4. **服务提炼（划定限定文边界）**：基于业务域能力内聚原则，根据不同子领域之间的关系和依赖性，明确领域边界，提炼出领域下的服务；基于聚合根定义服务接口（RPC），确保单一职责，
   **清晰且有意义，如果存在模糊，重复步骤 2、3**

![DDD-战略设计](/doc/DDD-战略设计.png)

### DDD - 战术实施（将领域概念映射成具体的服务代码）

> 战术实施即服务开发，旨在将领域核心映射成服务代码
>
> Go 微服务 DDD 代码布局参考: https://github.com/lupguo/go-ddd-layout

服务代码 DDD 目录分层推荐（接口层、应用层、领域层、基础设施层），每层职责和边界清晰，可以具体团队规模、业务复杂实际情况，控制服务的内聚程度，避免一开始服务拆分粒度过细导致服务数量增长过快，增大服务维护和治理成本。

- **接口层**：协议转换（DTO）、参数校验、错误处理
- **应用层**：业务流程编排（不含实际处理逻辑），针对业务缓存、分布式锁、事件通知、任务调度非业务流操作也放应用层（不含实现内容）
- **领域层**：领域服务核心能力，提供业务逻辑实现（比如赤字数据生成读取、报表下载、GMV 等)，核心要素包括领域服务 `service`
  、领域实体 `entity`、仓储接口 `repository`
- **基础设施层**：实现仓储接口，RPC 具体实现、DB 存储、缓存、消息中间件等，都会去实现领域层定义好的仓储接口

<img src="/doc/DDD-战术实施.png" width="500px" />

## 目录说明

```
├── Makefile
├── README.md
├── app
│     ├── application                 -- 应用层（业务编排）
│     │     ├── image_search_app.go
│     │     └── image_upload_app.go
│     ├── domain                      -- 领域层（核心业务逻辑处理）
│     │     ├── entity                -- 领域层-实体
│     │     ├── repository            -- 领域层-仓储接口
│     │     └── service               -- 领域层-领域服务
│     ├── infras                      -- 基础设施层
│     │     ├── dbs
│     │     ├── esearch
│     │     ├── httpclient
│     │     ├── mqs
│     │     ├── openaix
│     │     └── rds
│     └── interfaces                  -- 接口层（RPC接口、WebHandler处理函数实现)
│         ├── iamge_search_intf.go
│         └── image_upload_intf.go
│         
├── configs
├── go.mod
├── go.sum
├── internel
├── main.go  -- 服务注册、启动服务
└── test
    └── testdata
```