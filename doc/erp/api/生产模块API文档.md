# 生产模块API文档

## 1. 文档概述

### 1.1 文档目的
本文档定义了ERP系统生产模块的API接口规范，包括主生产计划、物料需求计划、生产订单管理、生产执行管理和物料清单管理等功能的API接口设计。旨在为前端开发和后端开发提供明确的接口规范，确保系统集成的顺利进行。

### 1.2 术语定义
| 术语 | 解释 |
|------|------|
| MPS | 主生产计划（Master Production Schedule），企业生产计划的核心部分 |
| MRP | 物料需求计划（Material Requirements Planning），根据生产计划计算物料需求的系统 |
| BOM | 物料清单（Bill of Materials），描述产品组成的清单 |
| WIP | 在制品（Work in Process），正在生产过程中的产品 |
| MO | 生产订单（Manufacturing Order），企业下达的生产任务 |
| Routing | 工艺路线，描述产品生产过程的工序和工作中心 |
| Work Center | 工作中心，生产过程中的加工单元 |

## 2. 通用规范

### 2.1 接口风格
- **URL格式**：`/api/v1/production/{resource}`
- **请求方法**：GET、POST、PUT、DELETE
- **数据格式**：JSON
- **认证方式**：JWT Token

### 2.2 响应格式

#### 2.2.1 成功响应
```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

#### 2.2.2 错误响应
```json
{
  "code": 400,
  "message": "error message",
  "data": null
}
```

### 2.3 分页响应
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

## 3. 主生产计划（MPS）API

### 3.1 获取主生产计划列表
- **接口路径**：`/api/v1/production/mps`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | productId | string | 否 | 产品ID |
  | startDate | string | 否 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 否 | 结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "mps-001",
        "productId": "prod-001",
        "productName": "产品A",
        "quantity": 100,
        "startDate": "2026-03-01",
        "endDate": "2026-03-05",
        "status": "approved",
        "createdBy": "user1",
        "createdAt": "2026-02-19T10:00:00Z",
        "updatedBy": "user1",
        "updatedAt": "2026-02-19T10:30:00Z"
      }
    ],
    "total": 10,
    "page": 1,
    "pageSize": 10
  }
}
```

### 3.2 获取主生产计划详情
- **接口路径**：`/api/v1/production/mps/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 主生产计划ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "mps-001",
    "productId": "prod-001",
    "productName": "产品A",
    "quantity": 100,
    "startDate": "2026-03-01",
    "endDate": "2026-03-05",
    "status": "approved",
    "description": "3月份生产计划",
    "createdBy": "user1",
    "createdAt": "2026-02-19T10:00:00Z",
    "updatedBy": "user1",
    "updatedAt": "2026-02-19T10:30:00Z"
  }
}
```

### 3.3 创建主生产计划
- **接口路径**：`/api/v1/production/mps`
- **请求方法**：POST
- **请求体**：
```json
{
  "productId": "prod-001",
  "quantity": 100,
  "startDate": "2026-03-01",
  "endDate": "2026-03-05",
  "description": "3月份生产计划"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "mps-001"
  }
}
```

### 3.4 更新主生产计划
- **接口路径**：`/api/v1/production/mps/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 主生产计划ID |
- **请求体**：
```json
{
  "quantity": 120,
  "startDate": "2026-03-01",
  "endDate": "2026-03-06",
  "description": "3月份生产计划（更新）"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 3.5 删除主生产计划
- **接口路径**：`/api/v1/production/mps/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 主生产计划ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 3.6 审核主生产计划
- **接口路径**：`/api/v1/production/mps/{id}/approve`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 主生产计划ID |
- **请求体**：
```json
{
  "approved": true,
  "comment": "审核通过"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 3.7 生成主生产计划
- **接口路径**：`/api/v1/production/mps/generate`
- **请求方法**：POST
- **请求体**：
```json
{
  "startDate": "2026-03-01",
  "endDate": "2026-03-31",
  "productIds": ["prod-001", "prod-002"]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "mpsIds": ["mps-001", "mps-002"]
  }
}
```

## 4. 物料需求计划（MRP）API

### 4.1 获取物料需求计划列表
- **接口路径**：`/api/v1/production/mrp`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | itemId | string | 否 | 物料ID |
  | startDate | string | 否 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 否 | 结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "mrp-001",
        "itemId": "item-001",
        "itemName": "原材料A",
        "requirementQuantity": 200,
        "availableQuantity": 50,
        "netRequirement": 150,
        "startDate": "2026-03-01",
        "endDate": "2026-03-05",
        "status": "planned",
        "createdAt": "2026-02-19T11:00:00Z"
      }
    ],
    "total": 15,
    "page": 1,
    "pageSize": 10
  }
}
```

### 4.2 获取物料需求计划详情
- **接口路径**：`/api/v1/production/mrp/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 物料需求计划ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "mrp-001",
    "itemId": "item-001",
    "itemName": "原材料A",
    "requirementQuantity": 200,
    "availableQuantity": 50,
    "netRequirement": 150,
    "startDate": "2026-03-01",
    "endDate": "2026-03-05",
    "status": "planned",
    "source": "mps-001",
    "createdAt": "2026-02-19T11:00:00Z"
  }
}
```

### 4.3 生成物料需求计划
- **接口路径**：`/api/v1/production/mrp/generate`
- **请求方法**：POST
- **请求体**：
```json
{
  "startDate": "2026-03-01",
  "endDate": "2026-03-31",
  "mpsIds": ["mps-001", "mps-002"]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "mrpCount": 15
  }
}
```

### 4.4 转换为采购计划
- **接口路径**：`/api/v1/production/mrp/convert-to-purchase`
- **请求方法**：POST
- **请求体**：
```json
{
  "mrpIds": ["mrp-001", "mrp-002"]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "purchasePlanIds": ["pp-001", "pp-002"]
  }
}
```

## 5. 物料清单（BOM）API

### 5.1 获取物料清单列表
- **接口路径**：`/api/v1/production/bom`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | productId | string | 否 | 产品ID |
  | version | string | 否 | 版本号 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "bom-001",
        "productId": "prod-001",
        "productName": "产品A",
        "version": "V1.0",
        "status": "active",
        "createdBy": "user1",
        "createdAt": "2026-02-19T12:00:00Z",
        "updatedBy": "user1",
        "updatedAt": "2026-02-19T12:30:00Z"
      }
    ],
    "total": 5,
    "page": 1,
    "pageSize": 10
  }
}
```

### 5.2 获取物料清单详情
- **接口路径**：`/api/v1/production/bom/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 物料清单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "bom-001",
    "productId": "prod-001",
    "productName": "产品A",
    "version": "V1.0",
    "status": "active",
    "description": "产品A的物料清单",
    "items": [
      {
        "itemId": "item-001",
        "itemName": "原材料A",
        "quantity": 2,
        "unit": "个",
        "usage": "100%"
      },
      {
        "itemId": "item-002",
        "itemName": "原材料B",
        "quantity": 1,
        "unit": "个",
        "usage": "100%"
      }
    ],
    "createdBy": "user1",
    "createdAt": "2026-02-19T12:00:00Z",
    "updatedBy": "user1",
    "updatedAt": "2026-02-19T12:30:00Z"
  }
}
```

### 5.3 创建物料清单
- **接口路径**：`/api/v1/production/bom`
- **请求方法**：POST
- **请求体**：
```json
{
  "productId": "prod-001",
  "version": "V1.0",
  "description": "产品A的物料清单",
  "items": [
    {
      "itemId": "item-001",
      "quantity": 2,
      "unit": "个",
      "usage": "100%"
    },
    {
      "itemId": "item-002",
      "quantity": 1,
      "unit": "个",
      "usage": "100%"
    }
  ]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "bom-001"
  }
}
```

### 5.4 更新物料清单
- **接口路径**：`/api/v1/production/bom/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 物料清单ID |
- **请求体**：
```json
{
  "description": "产品A的物料清单（更新）",
  "items": [
    {
      "itemId": "item-001",
      "quantity": 2,
      "unit": "个",
      "usage": "100%"
    },
    {
      "itemId": "item-002",
      "quantity": 1,
      "unit": "个",
      "usage": "100%"
    },
    {
      "itemId": "item-003",
      "quantity": 0.5,
      "unit": "千克",
      "usage": "95%"
    }
  ]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 5.5 删除物料清单
- **接口路径**：`/api/v1/production/bom/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 物料清单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 5.6 审核物料清单
- **接口路径**：`/api/v1/production/bom/{id}/approve`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 物料清单ID |
- **请求体**：
```json
{
  "approved": true,
  "comment": "审核通过"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 5.7 复制物料清单
- **接口路径**：`/api/v1/production/bom/{id}/copy`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 物料清单ID |
- **请求体**：
```json
{
  "productId": "prod-002",
  "version": "V1.0",
  "description": "基于产品A的物料清单"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "bom-002"
  }
}
```

## 6. 生产订单API

### 6.1 获取生产订单列表
- **接口路径**：`/api/v1/production/orders`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | orderNo | string | 否 | 订单编号 |
  | productId | string | 否 | 产品ID |
  | status | string | 否 | 状态（planned, released, in_progress, completed, cancelled） |
  | startDate | string | 否 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 否 | 结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "mo-001",
        "orderNo": "MO2026030001",
        "productId": "prod-001",
        "productName": "产品A",
        "quantity": 100,
        "status": "released",
        "scheduledStartDate": "2026-03-01",
        "scheduledEndDate": "2026-03-05",
        "actualStartDate": "2026-03-01",
        "actualEndDate": null,
        "createdBy": "user1",
        "createdAt": "2026-02-19T13:00:00Z",
        "updatedBy": "user1",
        "updatedAt": "2026-02-19T13:30:00Z"
      }
    ],
    "total": 8,
    "page": 1,
    "pageSize": 10
  }
}
```

### 6.2 获取生产订单详情
- **接口路径**：`/api/v1/production/orders/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产订单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "mo-001",
    "orderNo": "MO2026030001",
    "productId": "prod-001",
    "productName": "产品A",
    "quantity": 100,
    "status": "released",
    "scheduledStartDate": "2026-03-01",
    "scheduledEndDate": "2026-03-05",
    "actualStartDate": "2026-03-01",
    "actualEndDate": null,
    "bomId": "bom-001",
    "routingId": "routing-001",
    "description": "3月份生产订单",
    "createdBy": "user1",
    "createdAt": "2026-02-19T13:00:00Z",
    "updatedBy": "user1",
    "updatedAt": "2026-02-19T13:30:00Z"
  }
}
```

### 6.3 创建生产订单
- **接口路径**：`/api/v1/production/orders`
- **请求方法**：POST
- **请求体**：
```json
{
  "productId": "prod-001",
  "quantity": 100,
  "scheduledStartDate": "2026-03-01",
  "scheduledEndDate": "2026-03-05",
  "bomId": "bom-001",
  "routingId": "routing-001",
  "description": "3月份生产订单"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "mo-001",
    "orderNo": "MO2026030001"
  }
}
```

### 6.4 更新生产订单
- **接口路径**：`/api/v1/production/orders/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产订单ID |
- **请求体**：
```json
{
  "quantity": 120,
  "scheduledEndDate": "2026-03-06",
  "description": "3月份生产订单（更新）"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 6.5 删除生产订单
- **接口路径**：`/api/v1/production/orders/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产订单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 6.6 释放生产订单
- **接口路径**：`/api/v1/production/orders/{id}/release`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产订单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 6.7 开始生产订单
- **接口路径**：`/api/v1/production/orders/{id}/start`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产订单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 6.8 完成生产订单
- **接口路径**：`/api/v1/production/orders/{id}/complete`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产订单ID |
- **请求体**：
```json
{
  "actualQuantity": 98,
  "scrapQuantity": 2,
  "comment": "生产完成"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 6.9 取消生产订单
- **接口路径**：`/api/v1/production/orders/{id}/cancel`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产订单ID |
- **请求体**：
```json
{
  "reason": "原材料短缺"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 6.10 基于MPS生成生产订单
- **接口路径**：`/api/v1/production/orders/generate-from-mps`
- **请求方法**：POST
- **请求体**：
```json
{
  "mpsIds": ["mps-001", "mps-002"]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "orderIds": ["mo-001", "mo-002"]
  }
}
```

## 7. 生产执行API

### 7.1 获取生产工单列表
- **接口路径**：`/api/v1/production/work-orders`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | orderId | string | 否 | 生产订单ID |
  | operationId | string | 否 | 工序ID |
  | status | string | 否 | 状态（pending, in_progress, completed, cancelled） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "wo-001",
        "orderId": "mo-001",
        "orderNo": "MO2026030001",
        "operationId": "op-001",
        "operationName": "加工",
        "workCenterId": "wc-001",
        "workCenterName": "加工中心1",
        "quantity": 100,
        "status": "in_progress",
        "scheduledStartTime": "2026-03-01T08:00:00Z",
        "scheduledEndTime": "2026-03-01T16:00:00Z",
        "actualStartTime": "2026-03-01T08:00:00Z",
        "actualEndTime": null,
        "createdAt": "2026-02-19T14:00:00Z"
      }
    ],
    "total": 5,
    "page": 1,
    "pageSize": 10
  }
}
```

### 7.2 获取生产工单详情
- **接口路径**：`/api/v1/production/work-orders/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产工单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "wo-001",
    "orderId": "mo-001",
    "orderNo": "MO2026030001",
    "operationId": "op-001",
    "operationName": "加工",
    "workCenterId": "wc-001",
    "workCenterName": "加工中心1",
    "quantity": 100,
    "status": "in_progress",
    "scheduledStartTime": "2026-03-01T08:00:00Z",
    "scheduledEndTime": "2026-03-01T16:00:00Z",
    "actualStartTime": "2026-03-01T08:00:00Z",
    "actualEndTime": null,
    "createdAt": "2026-02-19T14:00:00Z"
  }
}
```

### 7.3 开始生产工单
- **接口路径**：`/api/v1/production/work-orders/{id}/start`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产工单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 7.4 完成生产工单
- **接口路径**：`/api/v1/production/work-orders/{id}/complete`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产工单ID |
- **请求体**：
```json
{
  "actualQuantity": 98,
  "scrapQuantity": 2,
  "comment": "加工完成"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 7.5 报工
- **接口路径**：`/api/v1/production/work-orders/{id}/report`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产工单ID |
- **请求体**：
```json
{
  "employeeId": "emp-001",
  "quantity": 50,
  "startTime": "2026-03-01T08:00:00Z",
  "endTime": "2026-03-01T12:00:00Z",
  "comment": "上午完成50个"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "reportId": "report-001"
  }
}
```

### 7.6 获取报工记录
- **接口路径**：`/api/v1/production/work-orders/{id}/reports`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 生产工单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": "report-001",
      "employeeId": "emp-001",
      "employeeName": "张三",
      "quantity": 50,
      "startTime": "2026-03-01T08:00:00Z",
      "endTime": "2026-03-01T12:00:00Z",
      "comment": "上午完成50个",
      "createdAt": "2026-03-01T12:00:00Z"
    }
  ]
}
```

## 8. 工艺路线API

### 8.1 获取工艺路线列表
- **接口路径**：`/api/v1/production/routings`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | productId | string | 否 | 产品ID |
  | version | string | 否 | 版本号 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "routing-001",
        "productId": "prod-001",
        "productName": "产品A",
        "version": "V1.0",
        "status": "active",
        "createdBy": "user1",
        "createdAt": "2026-02-19T15:00:00Z",
        "updatedBy": "user1",
        "updatedAt": "2026-02-19T15:30:00Z"
      }
    ],
    "total": 3,
    "page": 1,
    "pageSize": 10
  }
}
```

### 8.2 获取工艺路线详情
- **接口路径**：`/api/v1/production/routings/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 工艺路线ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "routing-001",
    "productId": "prod-001",
    "productName": "产品A",
    "version": "V1.0",
    "status": "active",
    "description": "产品A的工艺路线",
    "operations": [
      {
        "sequence": 1,
        "operationId": "op-001",
        "operationName": "加工",
        "workCenterId": "wc-001",
        "workCenterName": "加工中心1",
        "standardTime": 1800,
        "setupTime": 600,
        "teardownTime": 300
      },
      {
        "sequence": 2,
        "operationId": "op-002",
        "operationName": "装配",
        "workCenterId": "wc-002",
        "workCenterName": "装配中心1",
        "standardTime": 1200,
        "setupTime": 300,
        "teardownTime": 180
      }
    ],
    "createdBy": "user1",
    "createdAt": "2026-02-19T15:00:00Z",
    "updatedBy": "user1",
    "updatedAt": "2026-02-19T15:30:00Z"
  }
}
```

### 8.3 创建工艺路线
- **接口路径**：`/api/v1/production/routings`
- **请求方法**：POST
- **请求体**：
```json
{
  "productId": "prod-001",
  "version": "V1.0",
  "description": "产品A的工艺路线",
  "operations": [
    {
      "sequence": 1,
      "operationId": "op-001",
      "workCenterId": "wc-001",
      "standardTime": 1800,
      "setupTime": 600,
      "teardownTime": 300
    },
    {
      "sequence": 2,
      "operationId": "op-002",
      "workCenterId": "wc-002",
      "standardTime": 1200,
      "setupTime": 300,
      "teardownTime": 180
    }
  ]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "routing-001"
  }
}
```

### 8.4 更新工艺路线
- **接口路径**：`/api/v1/production/routings/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 工艺路线ID |
- **请求体**：
```json
{
  "description": "产品A的工艺路线（更新）",
  "operations": [
    {
      "sequence": 1,
      "operationId": "op-001",
      "workCenterId": "wc-001",
      "standardTime": 1800,
      "setupTime": 600,
      "teardownTime": 300
    },
    {
      "sequence": 2,
      "operationId": "op-002",
      "workCenterId": "wc-002",
      "standardTime": 1000,
      "setupTime": 300,
      "teardownTime": 180
    }
  ]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 8.5 删除工艺路线
- **接口路径**：`/api/v1/production/routings/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 工艺路线ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 9. 工作中心API

### 9.1 获取工作中心列表
- **接口路径**：`/api/v1/production/work-centers`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | name | string | 否 | 工作中心名称 |
  | type | string | 否 | 工作中心类型 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "wc-001",
        "code": "WC001",
        "name": "加工中心1",
        "type": "machining",
        "capacity": 8,
        "status": "active",
        "createdBy": "user1",
        "createdAt": "2026-02-19T16:00:00Z",
        "updatedBy": "user1",
        "updatedAt": "2026-02-19T16:30:00Z"
      }
    ],
    "total": 5,
    "page": 1,
    "pageSize": 10
  }
}
```

### 9.2 获取工作中心详情
- **接口路径**：`/api/v1/production/work-centers/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 工作中心ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "wc-001",
    "code": "WC001",
    "name": "加工中心1",
    "type": "machining",
    "capacity": 8,
    "status": "active",
    "description": "数控加工中心",
    "createdBy": "user1",
    "createdAt": "2026-02-19T16:00:00Z",
    "updatedBy": "user1",
    "updatedAt": "2026-02-19T16:30:00Z"
  }
}
```

### 9.3 创建工作中心
- **接口路径**：`/api/v1/production/work-centers`
- **请求方法**：POST
- **请求体**：
```json
{
  "code": "WC001",
  "name": "加工中心1",
  "type": "machining",
  "capacity": 8,
  "description": "数控加工中心"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "wc-001"
  }
}
```

### 9.4 更新工作中心
- **接口路径**：`/api/v1/production/work-centers/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 工作中心ID |
- **请求体**：
```json
{
  "name": "加工中心1（更新）",
  "capacity": 10,
  "description": "数控加工中心（更新）"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 9.5 删除工作中心
- **接口路径**：`/api/v1/production/work-centers/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 工作中心ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 10. 系统集成API

### 10.1 获取生产报表
- **接口路径**：`/api/v1/production/reports`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | reportType | string | 是 | 报表类型（production_plan, production_execution, work_center_load） |
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
  | productId | string | 否 | 产品ID |
  | workCenterId | string | 否 | 工作中心ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "reportType": "production_plan",
    "startDate": "2026-03-01",
    "endDate": "2026-03-31",
    "data": [
      {
        "productId": "prod-001",
        "productName": "产品A",
        "plannedQuantity": 1000,
        "actualQuantity": 950,
        "completionRate": 95
      }
    ]
  }
}
```

### 10.2 获取生产数据看板
- **接口路径**：`/api/v1/production/dashboard`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | period | string | 否 | 时间段（day, week, month），默认month |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "productionOrders": {
      "total": 10,
      "planned": 2,
      "inProgress": 5,
      "completed": 3,
      "cancelled": 0
    },
    "workCenterLoad": [
      {
        "workCenterId": "wc-001",
        "workCenterName": "加工中心1",
        "loadRate": 85
      }
    ],
    "productionTrend": [
      {
        "date": "2026-03-01",
        "quantity": 50
      }
    ]
  }
}
```

## 11. 错误码定义

| 错误码 | 描述 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |
| 501 | 接口未实现 |

## 12. 附录

### 12.1 参考文档
- 《ERP系统生产管理模块设计与实现》
- 《生产计划与控制》
- 《API设计最佳实践》

### 12.2 接口测试
- 使用Postman或Swagger进行接口测试
- 测试用例应覆盖正常流程和异常流程
- 测试数据应符合业务规则和数据约束
