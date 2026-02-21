# 库存模块API文档

## 1. 文档概述

### 1.1 文档目的
本文档定义了ERP系统库存模块的API接口规范，包括仓库管理、物料管理、库存交易、库存盘点、库存报表等功能的API接口设计。旨在为前端开发和后端开发提供明确的接口规范，确保系统集成的顺利进行。

### 1.2 术语定义
| 术语 | 解释 |
|------|------|
| Warehouse | 仓库，存储物料的物理位置 |
| Location | 库位，仓库内的具体存储位置 |
| Item | 物料，库存管理的对象 |
| SKU | 库存保有单位（Stock Keeping Unit），物料的唯一标识符 |
| Inventory | 库存，物料的存储数量 |
| Transaction | 交易，引起库存变化的操作 |
| Movement | 移动，物料在仓库间的转移 |
| Count | 盘点，对库存数量的实物核查 |
| Cycle Count | 循环盘点，定期对部分库存进行的盘点 |
| Stock Take | 全面盘点，对所有库存进行的盘点 |

## 2. 通用规范

### 2.1 接口风格
- **URL格式**：`/api/v1/inventory/{resource}`
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

## 3. 仓库管理API

### 3.1 获取仓库列表
- **接口路径**：`/api/v1/inventory/warehouses`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | warehouseNo | string | 否 | 仓库编号 |
  | name | string | 否 | 仓库名称 |
  | status | string | 否 | 状态（active, inactive） |
  | type | string | 否 | 仓库类型（raw_material, finished_goods, semi_finished, tools） |
  | region | string | 否 | 地区 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "warehouse-001",
        "warehouseNo": "WH001",
        "name": "北京主仓库",
        "type": "finished_goods",
        "address": "北京市朝阳区",
        "region": "华北",
        "contact": "张三",
        "phone": "13800138001",
        "status": "active",
        "createdBy": "admin",
        "createdAt": "2020-01-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2020-01-01T08:00:00Z"
      }
    ],
    "total": 50,
    "page": 1,
    "pageSize": 10
  }
}
```

### 3.2 获取仓库详情
- **接口路径**：`/api/v1/inventory/warehouses/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 仓库ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "warehouse-001",
    "warehouseNo": "WH001",
    "name": "北京主仓库",
    "type": "finished_goods",
    "address": "北京市朝阳区",
    "region": "华北",
    "contact": "张三",
    "phone": "13800138001",
    "description": "主要存储成品的仓库",
    "capacity": 10000,
    "usedCapacity": 5000,
    "status": "active",
    "locations": [
      {
        "id": "location-001",
        "code": "A-01-01",
        "name": "A区-01排-01层",
        "capacity": 1000,
        "usedCapacity": 500
      }
    ],
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 3.3 创建仓库
- **接口路径**：`/api/v1/inventory/warehouses`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "北京主仓库",
  "type": "finished_goods",
  "address": "北京市朝阳区",
  "region": "华北",
  "contact": "张三",
  "phone": "13800138001",
  "description": "主要存储成品的仓库",
  "capacity": 10000
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "warehouse-001",
    "warehouseNo": "WH001"
  }
}
```

### 3.4 更新仓库
- **接口路径**：`/api/v1/inventory/warehouses/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 仓库ID |
- **请求体**：
```json
{
  "name": "北京主仓库（更新）",
  "contact": "李四",
  "phone": "13900139001",
  "capacity": 15000
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

### 3.5 删除仓库
- **接口路径**：`/api/v1/inventory/warehouses/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 仓库ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 3.6 库位管理
- **接口路径**：`/api/v1/inventory/warehouses/{id}/locations`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 仓库ID |
- **请求体**：
```json
{
  "code": "A-01-01",
  "name": "A区-01排-01层",
  "capacity": 1000,
  "description": "存储成品的区域"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "location-001",
    "code": "A-01-01"
  }
}
```

## 4. 物料管理API

### 4.1 获取物料列表
- **接口路径**：`/api/v1/inventory/items`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | itemNo | string | 否 | 物料编号 |
  | name | string | 否 | 物料名称 |
  | sku | string | 否 | SKU编码 |
  | category | string | 否 | 物料类别 |
  | status | string | 否 | 状态（active, inactive） |
  | unit | string | 否 | 单位 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "item-001",
        "itemNo": "MAT001",
        "name": "内存条",
        "sku": "RAM-16GB-DDR4",
        "category": "电子元件",
        "specification": "16GB DDR4",
        "unit": "条",
        "status": "active",
        "createdBy": "admin",
        "createdAt": "2020-01-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2020-01-01T08:00:00Z"
      }
    ],
    "total": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

### 4.2 获取物料详情
- **接口路径**：`/api/v1/inventory/items/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 物料ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "item-001",
    "itemNo": "MAT001",
    "name": "内存条",
    "sku": "RAM-16GB-DDR4",
    "category": "电子元件",
    "specification": "16GB DDR4",
    "unit": "条",
    "description": "电脑内存条",
    "weight": 0.1,
    "volume": 0.001,
    "minStock": 10,
    "maxStock": 100,
    "reorderPoint": 20,
    "leadTime": 7,
    "status": "active",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 4.3 创建物料
- **接口路径**：`/api/v1/inventory/items`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "内存条",
  "sku": "RAM-16GB-DDR4",
  "category": "电子元件",
  "specification": "16GB DDR4",
  "unit": "条",
  "description": "电脑内存条",
  "weight": 0.1,
  "volume": 0.001,
  "minStock": 10,
  "maxStock": 100,
  "reorderPoint": 20,
  "leadTime": 7
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "item-001",
    "itemNo": "MAT001"
  }
}
```

### 4.4 更新物料
- **接口路径**：`/api/v1/inventory/items/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 物料ID |
- **请求体**：
```json
{
  "name": "内存条（更新）",
  "specification": "16GB DDR4 3200MHz",
  "minStock": 15,
  "maxStock": 120,
  "reorderPoint": 25
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

### 4.5 删除物料
- **接口路径**：`/api/v1/inventory/items/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 物料ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 4.6 获取物料库存
- **接口路径**：`/api/v1/inventory/items/{id}/stock`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 物料ID |
  | warehouseId | string | 否 | 仓库ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "totalStock": 50,
    "warehouses": [
      {
        "warehouseId": "warehouse-001",
        "warehouseName": "北京主仓库",
        "quantity": 30,
        "locations": [
          {
            "locationId": "location-001",
            "locationCode": "A-01-01",
            "quantity": 30
          }
        ]
      }
    ]
  }
}
```

## 5. 库存交易API

### 5.1 获取库存交易列表
- **接口路径**：`/api/v1/inventory/transactions`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | transactionNo | string | 否 | 交易编号 |
  | type | string | 否 | 交易类型（purchase, sales, transfer, adjustment, production） |
  | itemId | string | 否 | 物料ID |
  | warehouseId | string | 否 | 仓库ID |
  | transactionDateStart | string | 否 | 交易开始日期，格式：YYYY-MM-DD |
  | transactionDateEnd | string | 否 | 交易结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "transaction-001",
        "transactionNo": "TRX2023060001",
        "type": "purchase",
        "referenceNo": "PO2023060001",
        "warehouseId": "warehouse-001",
        "warehouseName": "北京主仓库",
        "transactionDate": "2023-06-01",
        "totalQuantity": 10,
        "createdBy": "admin",
        "createdAt": "2023-06-01T08:00:00Z"
      }
    ],
    "total": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

### 5.2 获取库存交易详情
- **接口路径**：`/api/v1/inventory/transactions/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 交易ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "transaction-001",
    "transactionNo": "TRX2023060001",
    "type": "purchase",
    "referenceNo": "PO2023060001",
    "warehouseId": "warehouse-001",
    "warehouseName": "北京主仓库",
    "transactionDate": "2023-06-01",
    "remarks": "采购入库",
    "items": [
      {
        "id": "item-001",
        "itemId": "mat-001",
        "itemCode": "MAT001",
        "itemName": "内存条",
        "quantity": 10,
        "unit": "条",
        "unitCost": 3000,
        "totalCost": 30000,
        "locationId": "location-001",
        "locationCode": "A-01-01"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-01T08:00:00Z"
  }
}
```

### 5.3 创建库存交易
- **接口路径**：`/api/v1/inventory/transactions`
- **请求方法**：POST
- **请求体**：
```json
{
  "type": "purchase",
  "referenceNo": "PO2023060001",
  "warehouseId": "warehouse-001",
  "transactionDate": "2023-06-01",
  "remarks": "采购入库",
  "items": [
    {
      "itemId": "mat-001",
      "quantity": 10,
      "unitCost": 3000,
      "locationId": "location-001"
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
    "id": "transaction-001",
    "transactionNo": "TRX2023060001"
  }
}
```

### 5.4 库存调整
- **接口路径**：`/api/v1/inventory/transactions/adjustment`
- **请求方法**：POST
- **请求体**：
```json
{
  "warehouseId": "warehouse-001",
  "transactionDate": "2023-06-01",
  "reason": "库存差异调整",
  "remarks": "盘点调整",
  "items": [
    {
      "itemId": "mat-001",
      "quantity": -2,
      "unitCost": 3000,
      "locationId": "location-001"
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
    "id": "transaction-002",
    "transactionNo": "TRX2023060002"
  }
}
```

### 5.5 仓库间转移
- **接口路径**：`/api/v1/inventory/transactions/transfer`
- **请求方法**：POST
- **请求体**：
```json
{
  "fromWarehouseId": "warehouse-001",
  "toWarehouseId": "warehouse-002",
  "transactionDate": "2023-06-01",
  "remarks": "仓库间转移",
  "items": [
    {
      "itemId": "mat-001",
      "quantity": 5,
      "fromLocationId": "location-001",
      "toLocationId": "location-002"
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
    "id": "transaction-003",
    "transactionNo": "TRX2023060003"
  }
}
```

## 6. 库存盘点API

### 6.1 获取盘点列表
- **接口路径**：`/api/v1/inventory/counts`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | countNo | string | 否 | 盘点单号 |
  | type | string | 否 | 盘点类型（cycle, stock_take） |
  | warehouseId | string | 否 | 仓库ID |
  | status | string | 否 | 状态（draft, in_progress, completed, cancelled） |
  | startDateStart | string | 否 | 开始日期开始，格式：YYYY-MM-DD |
  | startDateEnd | string | 否 | 开始日期结束，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "count-001",
        "countNo": "CNT2023060001",
        "type": "cycle",
        "warehouseId": "warehouse-001",
        "warehouseName": "北京主仓库",
        "startDate": "2023-06-01",
        "endDate": "2023-06-02",
        "status": "completed",
        "createdBy": "admin",
        "createdAt": "2023-06-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-02T09:00:00Z"
      }
    ],
    "total": 50,
    "page": 1,
    "pageSize": 10
  }
}
```

### 6.2 获取盘点详情
- **接口路径**：`/api/v1/inventory/counts/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 盘点单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "count-001",
    "countNo": "CNT2023060001",
    "type": "cycle",
    "warehouseId": "warehouse-001",
    "warehouseName": "北京主仓库",
    "startDate": "2023-06-01",
    "endDate": "2023-06-02",
    "remarks": "月度循环盘点",
    "status": "completed",
    "items": [
      {
        "id": "item-001",
        "itemId": "mat-001",
        "itemCode": "MAT001",
        "itemName": "内存条",
        "systemQuantity": 10,
        "actualQuantity": 8,
        "difference": -2,
        "unitCost": 3000,
        "differenceAmount": -6000,
        "locationId": "location-001",
        "locationCode": "A-01-01"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-02T09:00:00Z"
  }
}
```

### 6.3 创建盘点单
- **接口路径**：`/api/v1/inventory/counts`
- **请求方法**：POST
- **请求体**：
```json
{
  "type": "cycle",
  "warehouseId": "warehouse-001",
  "startDate": "2023-06-01",
  "remarks": "月度循环盘点",
  "items": [
    {
      "itemId": "mat-001",
      "locationId": "location-001"
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
    "id": "count-001",
    "countNo": "CNT2023060001"
  }
}
```

### 6.4 更新盘点结果
- **接口路径**：`/api/v1/inventory/counts/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 盘点单ID |
- **请求体**：
```json
{
  "endDate": "2023-06-02",
  "items": [
    {
      "id": "item-001",
      "actualQuantity": 8
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

### 6.5 完成盘点
- **接口路径**：`/api/v1/inventory/counts/{id}/complete`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 盘点单ID |
- **请求体**：
```json
{
  "adjustInventory": true
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "adjustmentTransactionId": "transaction-004"
  }
}
```

### 6.6 取消盘点
- **接口路径**：`/api/v1/inventory/counts/{id}/cancel`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 盘点单ID |
- **请求体**：
```json
{
  "reason": "盘点错误，取消盘点"
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

## 7. 库存报表API

### 7.1 获取库存余额报表
- **接口路径**：`/api/v1/inventory/reports/balance`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | warehouseId | string | 否 | 仓库ID |
  | itemId | string | 否 | 物料ID |
  | category | string | 否 | 物料类别 |
  | asOfDate | string | 否 | 截至日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "asOfDate": "2023-06-30",
    "totalItems": 100,
    "totalValue": 5000000,
    "items": [
      {
        "itemId": "mat-001",
        "itemCode": "MAT001",
        "itemName": "内存条",
        "warehouseId": "warehouse-001",
        "warehouseName": "北京主仓库",
        "quantity": 30,
        "unitCost": 3000,
        "totalValue": 90000
      }
    ]
  }
}
```

### 7.2 获取库存变动报表
- **接口路径**：`/api/v1/inventory/reports/movement`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
  | warehouseId | string | 否 | 仓库ID |
  | itemId | string | 否 | 物料ID |
  | type | string | 否 | 交易类型 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "period": "2023-06-01 至 2023-06-30",
    "items": [
      {
        "itemId": "mat-001",
        "itemCode": "MAT001",
        "itemName": "内存条",
        "beginningBalance": 20,
        "inQuantity": 15,
        "outQuantity": 5,
        "endingBalance": 30,
        "transactions": [
          {
            "transactionNo": "TRX2023060001",
            "type": "purchase",
            "date": "2023-06-01",
            "quantity": 10
          }
        ]
      }
    ]
  }
}
```

### 7.3 获取库存预警报表
- **接口路径**：`/api/v1/inventory/reports/alert`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | warehouseId | string | 否 | 仓库ID |
  | type | string | 否 | 预警类型（low_stock, excess_stock, reorder_point） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "lowStockItems": [
      {
        "itemId": "mat-002",
        "itemCode": "MAT002",
        "itemName": "硬盘",
        "warehouseId": "warehouse-001",
        "warehouseName": "北京主仓库",
        "currentStock": 5,
        "minStock": 10,
        "reorderPoint": 15
      }
    ],
    "excessStockItems": [
      {
        "itemId": "mat-003",
        "itemCode": "MAT003",
        "itemName": "鼠标",
        "warehouseId": "warehouse-001",
        "warehouseName": "北京主仓库",
        "currentStock": 150,
        "maxStock": 100
      }
    ]
  }
}
```

### 7.4 获取ABC分析报表
- **接口路径**：`/api/v1/inventory/reports/abc`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
  | warehouseId | string | 否 | 仓库ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "period": "2023-01-01 至 2023-06-30",
    "items": [
      {
        "itemId": "mat-001",
        "itemCode": "MAT001",
        "itemName": "内存条",
        "category": "A",
        "quantity": 100,
        "value": 300000,
        "percentage": 60
      }
    ]
  }
}
```

### 7.5 导出库存报表
- **接口路径**：`/api/v1/inventory/reports/export`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | type | string | 是 | 报表类型（balance, movement, alert, abc） |
  | startDate | string | 否 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 否 | 结束日期，格式：YYYY-MM-DD |
  | warehouseId | string | 否 | 仓库ID |
  | itemId | string | 否 | 物料ID |
- **响应格式**：Excel文件

## 8. 错误码定义

| 错误码 | 描述 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |
| 501 | 接口未实现 |
| 4001 | 库存不足 |
| 4002 | 仓库不存在 |
| 4003 | 物料不存在 |
| 4004 | 库位不存在 |
| 4005 | 交易类型错误 |
| 4006 | 盘点单状态错误 |

## 9. 附录

### 9.1 参考文档
- 《ERP系统库存管理模块设计与实现》
- 《库存管理实务》
- 《API设计最佳实践》

### 9.2 接口测试
- 使用Postman或Swagger进行接口测试
- 测试用例应覆盖正常流程和异常流程
- 测试数据应符合业务规则和数据约束