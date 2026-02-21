# 销售模块API文档

## 1. 文档概述

### 1.1 文档目的
本文档定义了ERP系统销售模块的API接口规范，包括客户管理、销售报价、销售订单、销售发货、销售退货、销售发票和销售报表等功能的API接口设计。旨在为前端开发和后端开发提供明确的接口规范，确保系统集成的顺利进行。

### 1.2 术语定义
| 术语 | 解释 |
|------|------|
| SO | 销售订单（Sales Order），企业向客户发出的销售确认单 |
| Quotation | 销售报价单，向客户提供的价格估算 |
| Delivery | 销售发货单，记录货物发出的单据 |
| AR | 应收账款（Accounts Receivable），客户欠企业的款项 |
| Customer | 客户，购买企业产品或服务的外部单位 |
| Lead | 销售线索，潜在的客户信息 |
| Opportunity | 销售机会，有购买意向的潜在交易 |
| Sales Pipeline | 销售漏斗，销售机会的进展阶段 |

## 2. 通用规范

### 2.1 接口风格
- **URL格式**：`/api/v1/sales/{resource}`
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

## 3. 客户管理API

### 3.1 获取客户列表
- **接口路径**：`/api/v1/sales/customers`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | customerNo | string | 否 | 客户编号 |
  | name | string | 否 | 客户名称 |
  | contact | string | 否 | 联系人 |
  | phone | string | 否 | 电话 |
  | status | string | 否 | 状态（active, inactive） |
  | category | string | 否 | 客户类别 |
  | region | string | 否 | 地区 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "customer-001",
        "customerNo": "CUS001",
        "name": "北京客户有限公司",
        "contact": "李四",
        "phone": "13900139001",
        "email": "contact@customer.com",
        "address": "北京市朝阳区",
        "taxNo": "110105199001011234",
        "category": "企业客户",
        "region": "华北",
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

### 3.2 获取客户详情
- **接口路径**：`/api/v1/sales/customers/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 客户ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "customer-001",
    "customerNo": "CUS001",
    "name": "北京客户有限公司",
    "contact": "李四",
    "phone": "13900139001",
    "email": "contact@customer.com",
    "address": "北京市朝阳区",
    "taxNo": "110105199001011234",
    "bankName": "建设银行",
    "bankAccount": "6227001234567890123",
    "category": "企业客户",
    "region": "华北",
    "creditLimit": 200000,
    "creditDays": 30,
    "remarks": "优质客户",
    "status": "active",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 3.3 创建客户
- **接口路径**：`/api/v1/sales/customers`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "北京客户有限公司",
  "contact": "李四",
  "phone": "13900139001",
  "email": "contact@customer.com",
  "address": "北京市朝阳区",
  "taxNo": "110105199001011234",
  "bankName": "建设银行",
  "bankAccount": "6227001234567890123",
  "category": "企业客户",
  "region": "华北",
  "creditLimit": 200000,
  "creditDays": 30,
  "remarks": "优质客户"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "customer-001",
    "customerNo": "CUS001"
  }
}
```

### 3.4 更新客户
- **接口路径**：`/api/v1/sales/customers/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 客户ID |
- **请求体**：
```json
{
  "name": "北京客户有限公司（更新）",
  "contact": "王五",
  "phone": "13800138001",
  "creditLimit": 250000,
  "creditDays": 45
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

### 3.5 删除客户
- **接口路径**：`/api/v1/sales/customers/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 客户ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 3.6 客户信用评估
- **接口路径**：`/api/v1/sales/customers/{id}/credit-evaluate`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 客户ID |
- **请求体**：
```json
{
  "evaluationDate": "2023-06-01",
  "creditScore": 90,
  "riskLevel": "low",
  "remarks": "信用良好"
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

## 4. 销售报价管理API

### 4.1 获取销售报价列表
- **接口路径**：`/api/v1/sales/quotations`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | quotationNo | string | 否 | 报价单编号 |
  | customerId | string | 否 | 客户ID |
  | customerName | string | 否 | 客户名称 |
  | status | string | 否 | 状态（draft, submitted, approved, rejected, expired） |
  | createDateStart | string | 否 | 创建开始日期，格式：YYYY-MM-DD |
  | createDateEnd | string | 否 | 创建结束日期，格式：YYYY-MM-DD |
  | expiryDateStart | string | 否 | 有效期开始日期，格式：YYYY-MM-DD |
  | expiryDateEnd | string | 否 | 有效期结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "quotation-001",
        "quotationNo": "QT2023060001",
        "customerId": "customer-001",
        "customerName": "北京客户有限公司",
        "contact": "李四",
        "phone": "13900139001",
        "quotationDate": "2023-06-01",
        "expiryDate": "2023-06-30",
        "totalAmount": 50000,
        "status": "approved",
        "createdBy": "admin",
        "createdAt": "2023-06-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-02T08:00:00Z"
      }
    ],
    "total": 50,
    "page": 1,
    "pageSize": 10
  }
}
```

### 4.2 获取销售报价详情
- **接口路径**：`/api/v1/sales/quotations/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 报价单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "quotation-001",
    "quotationNo": "QT2023060001",
    "customerId": "customer-001",
    "customerName": "北京客户有限公司",
    "contact": "李四",
    "phone": "13900139001",
    "address": "北京市朝阳区",
    "quotationDate": "2023-06-01",
    "expiryDate": "2023-06-30",
    "paymentTerms": "30天内付款",
    "shippingTerms": "运费由我方承担",
    "remarks": "批量采购可享受折扣",
    "totalAmount": 50000,
    "status": "approved",
    "items": [
      {
        "id": "item-001",
        "productId": "prod-001",
        "productCode": "PROD001",
        "productName": "笔记本电脑",
        "specification": "i7 16GB 512GB",
        "unit": "台",
        "quantity": 10,
        "unitPrice": 5000,
        "amount": 50000,
        "remark": "正品保证"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-02T08:00:00Z"
  }
}
```

### 4.3 创建销售报价
- **接口路径**：`/api/v1/sales/quotations`
- **请求方法**：POST
- **请求体**：
```json
{
  "customerId": "customer-001",
  "quotationDate": "2023-06-01",
  "expiryDate": "2023-06-30",
  "paymentTerms": "30天内付款",
  "shippingTerms": "运费由我方承担",
  "remarks": "批量采购可享受折扣",
  "items": [
    {
      "productId": "prod-001",
      "quantity": 10,
      "unitPrice": 5000,
      "remark": "正品保证"
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
    "id": "quotation-001",
    "quotationNo": "QT2023060001"
  }
}
```

### 4.4 更新销售报价
- **接口路径**：`/api/v1/sales/quotations/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 报价单ID |
- **请求体**：
```json
{
  "expiryDate": "2023-07-15",
  "remarks": "批量采购可享受折扣（更新）",
  "items": [
    {
      "productId": "prod-001",
      "quantity": 15,
      "unitPrice": 4800,
      "remark": "正品保证"
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

### 4.5 删除销售报价
- **接口路径**：`/api/v1/sales/quotations/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 报价单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 4.6 审批销售报价
- **接口路径**：`/api/v1/sales/quotations/{id}/approve`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 报价单ID |
- **请求体**：
```json
{
  "approved": true,
  "remarks": "同意报价单"
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

### 4.7 从报价单生成销售订单
- **接口路径**：`/api/v1/sales/quotations/{id}/generate-order`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 报价单ID |
- **请求体**：
```json
{
  "orderDate": "2023-06-10",
  "deliveryDate": "2023-06-20"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "orderId": "order-001",
    "orderNo": "SO2023060001"
  }
}
```

## 5. 销售订单管理API

### 5.1 获取销售订单列表
- **接口路径**：`/api/v1/sales/orders`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | orderNo | string | 否 | 订单编号 |
  | customerId | string | 否 | 客户ID |
  | customerName | string | 否 | 客户名称 |
  | status | string | 否 | 状态（draft, approved, partially_delivered, fully_delivered, cancelled） |
  | createDateStart | string | 否 | 创建开始日期，格式：YYYY-MM-DD |
  | createDateEnd | string | 否 | 创建结束日期，格式：YYYY-MM-DD |
  | deliveryDateStart | string | 否 | 交货开始日期，格式：YYYY-MM-DD |
  | deliveryDateEnd | string | 否 | 交货结束日期，格式：YYYY-MM-DD |
  | totalAmountStart | number | 否 | 总金额起始 |
  | totalAmountEnd | number | 否 | 总金额结束 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "order-001",
        "orderNo": "SO2023060001",
        "customerId": "customer-001",
        "customerName": "北京客户有限公司",
        "contact": "李四",
        "phone": "13900139001",
        "orderDate": "2023-06-01",
        "deliveryDate": "2023-06-15",
        "totalAmount": 50000,
        "status": "approved",
        "createdBy": "admin",
        "createdAt": "2023-06-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-01T09:00:00Z"
      }
    ],
    "total": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

### 5.2 获取销售订单详情
- **接口路径**：`/api/v1/sales/orders/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 订单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "order-001",
    "orderNo": "SO2023060001",
    "customerId": "customer-001",
    "customerName": "北京客户有限公司",
    "contact": "李四",
    "phone": "13900139001",
    "address": "北京市朝阳区",
    "orderDate": "2023-06-01",
    "deliveryDate": "2023-06-15",
    "paymentTerms": "30天内付款",
    "shippingTerms": "运费由我方承担",
    "remarks": "加急订单，请尽快发货",
    "totalAmount": 50000,
    "status": "approved",
    "items": [
      {
        "id": "item-001",
        "productId": "prod-001",
        "productCode": "PROD001",
        "productName": "笔记本电脑",
        "specification": "i7 16GB 512GB",
        "unit": "台",
        "quantity": 10,
        "unitPrice": 5000,
        "amount": 50000,
        "deliveredQuantity": 0,
        "remark": "正品保证"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-01T09:00:00Z"
  }
}
```

### 5.3 创建销售订单
- **接口路径**：`/api/v1/sales/orders`
- **请求方法**：POST
- **请求体**：
```json
{
  "customerId": "customer-001",
  "orderDate": "2023-06-01",
  "deliveryDate": "2023-06-15",
  "paymentTerms": "30天内付款",
  "shippingTerms": "运费由我方承担",
  "remarks": "加急订单，请尽快发货",
  "items": [
    {
      "productId": "prod-001",
      "quantity": 10,
      "unitPrice": 5000,
      "remark": "正品保证"
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
    "id": "order-001",
    "orderNo": "SO2023060001"
  }
}
```

### 5.4 更新销售订单
- **接口路径**：`/api/v1/sales/orders/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 订单ID |
- **请求体**：
```json
{
  "deliveryDate": "2023-06-10",
  "remarks": "加急订单，请尽快发货（更新）",
  "items": [
    {
      "productId": "prod-001",
      "quantity": 15,
      "unitPrice": 4800,
      "remark": "正品保证"
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

### 5.5 删除销售订单
- **接口路径**：`/api/v1/sales/orders/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 订单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 5.6 审批销售订单
- **接口路径**：`/api/v1/sales/orders/{id}/approve`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 订单ID |
- **请求体**：
```json
{
  "approved": true,
  "remarks": "同意销售订单"
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

### 5.7 取消销售订单
- **接口路径**：`/api/v1/sales/orders/{id}/cancel`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 订单ID |
- **请求体**：
```json
{
  "reason": "客户需求变更，取消订单"
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

### 5.8 从销售订单生成发货单
- **接口路径**：`/api/v1/sales/orders/{id}/generate-delivery`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 订单ID |
- **请求体**：
```json
{
  "deliveryDate": "2023-06-15",
  "warehouseId": "warehouse-001",
  "items": [
    {
      "orderItemId": "item-001",
      "quantity": 10
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
    "deliveryId": "delivery-001",
    "deliveryNo": "DL2023060001"
  }
}
```

## 6. 销售发货管理API

### 6.1 获取销售发货列表
- **接口路径**：`/api/v1/sales/deliveries`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | deliveryNo | string | 否 | 发货单编号 |
  | orderId | string | 否 | 销售订单ID |
  | orderNo | string | 否 | 销售订单编号 |
  | customerId | string | 否 | 客户ID |
  | customerName | string | 否 | 客户名称 |
  | deliveryDateStart | string | 否 | 发货开始日期，格式：YYYY-MM-DD |
  | deliveryDateEnd | string | 否 | 发货结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "delivery-001",
        "deliveryNo": "DL2023060001",
        "orderId": "order-001",
        "orderNo": "SO2023060001",
        "customerId": "customer-001",
        "customerName": "北京客户有限公司",
        "deliveryDate": "2023-06-15",
        "totalQuantity": 10,
        "status": "completed",
        "createdBy": "admin",
        "createdAt": "2023-06-15T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-15T09:00:00Z"
      }
    ],
    "total": 50,
    "page": 1,
    "pageSize": 10
  }
}
```

### 6.2 获取销售发货详情
- **接口路径**：`/api/v1/sales/deliveries/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 发货单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "delivery-001",
    "deliveryNo": "DL2023060001",
    "orderId": "order-001",
    "orderNo": "SO2023060001",
    "customerId": "customer-001",
    "customerName": "北京客户有限公司",
    "contact": "李四",
    "phone": "13900139001",
    "address": "北京市朝阳区",
    "deliveryDate": "2023-06-15",
    "warehouseId": "warehouse-001",
    "warehouseName": "北京仓库",
    "carrier": "顺丰快递",
    "trackingNo": "SF1234567890",
    "remarks": "加急发货",
    "status": "completed",
    "items": [
      {
        "id": "item-001",
        "orderItemId": "order-item-001",
        "productId": "prod-001",
        "productCode": "PROD001",
        "productName": "笔记本电脑",
        "specification": "i7 16GB 512GB",
        "unit": "台",
        "orderedQuantity": 10,
        "deliveredQuantity": 10,
        "remark": "正品保证"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-15T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-15T09:00:00Z"
  }
}
```

### 6.3 创建销售发货
- **接口路径**：`/api/v1/sales/deliveries`
- **请求方法**：POST
- **请求体**：
```json
{
  "orderId": "order-001",
  "deliveryDate": "2023-06-15",
  "warehouseId": "warehouse-001",
  "carrier": "顺丰快递",
  "trackingNo": "SF1234567890",
  "remarks": "加急发货",
  "items": [
    {
      "orderItemId": "order-item-001",
      "deliveredQuantity": 10,
      "remark": "正品保证"
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
    "id": "delivery-001",
    "deliveryNo": "DL2023060001"
  }
}
```

### 6.4 更新销售发货
- **接口路径**：`/api/v1/sales/deliveries/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 发货单ID |
- **请求体**：
```json
{
  "carrier": "京东物流",
  "trackingNo": "JD9876543210",
  "remarks": "加急发货（更新）",
  "items": [
    {
      "orderItemId": "order-item-001",
      "deliveredQuantity": 10,
      "remark": "正品保证"
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

### 6.5 删除销售发货
- **接口路径**：`/api/v1/sales/deliveries/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 发货单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 7. 销售发票管理API

### 7.1 获取销售发票列表
- **接口路径**：`/api/v1/sales/invoices`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | invoiceNo | string | 否 | 发票编号 |
  | customerId | string | 否 | 客户ID |
  | customerName | string | 否 | 客户名称 |
  | orderId | string | 否 | 销售订单ID |
  | orderNo | string | 否 | 销售订单编号 |
  | invoiceDateStart | string | 否 | 发票日期开始，格式：YYYY-MM-DD |
  | invoiceDateEnd | string | 否 | 发票日期结束，格式：YYYY-MM-DD |
  | status | string | 否 | 状态（unpaid, partially_paid, fully_paid） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "invoice-001",
        "invoiceNo": "INV2023060001",
        "customerId": "customer-001",
        "customerName": "北京客户有限公司",
        "orderId": "order-001",
        "orderNo": "SO2023060001",
        "invoiceDate": "2023-06-15",
        "dueDate": "2023-07-15",
        "totalAmount": 50000,
        "paidAmount": 0,
        "balanceAmount": 50000,
        "status": "unpaid",
        "createdBy": "admin",
        "createdAt": "2023-06-15T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-15T08:00:00Z"
      }
    ],
    "total": 50,
    "page": 1,
    "pageSize": 10
  }
}
```

### 7.2 获取销售发票详情
- **接口路径**：`/api/v1/sales/invoices/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 发票ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "invoice-001",
    "invoiceNo": "INV2023060001",
    "customerId": "customer-001",
    "customerName": "北京客户有限公司",
    "customerTaxNo": "110105199001011234",
    "orderId": "order-001",
    "orderNo": "SO2023060001",
    "invoiceDate": "2023-06-15",
    "dueDate": "2023-07-15",
    "totalAmount": 50000,
    "taxAmount": 5752.21,
    "netAmount": 44247.79,
    "paidAmount": 0,
    "balanceAmount": 50000,
    "status": "unpaid",
    "items": [
      {
        "id": "item-001",
        "productId": "prod-001",
        "productCode": "PROD001",
        "productName": "笔记本电脑",
        "specification": "i7 16GB 512GB",
        "quantity": 10,
        "unitPrice": 4424.78,
        "taxRate": 13,
        "taxAmount": 575.22,
        "totalAmount": 5000
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-15T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-15T08:00:00Z"
  }
}
```

### 7.3 创建销售发票
- **接口路径**：`/api/v1/sales/invoices`
- **请求方法**：POST
- **请求体**：
```json
{
  "customerId": "customer-001",
  "orderId": "order-001",
  "invoiceDate": "2023-06-15",
  "dueDate": "2023-07-15",
  "items": [
    {
      "productId": "prod-001",
      "quantity": 10,
      "unitPrice": 4424.78,
      "taxRate": 13
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
    "id": "invoice-001",
    "invoiceNo": "INV2023060001"
  }
}
```

### 7.4 更新销售发票
- **接口路径**：`/api/v1/sales/invoices/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 发票ID |
- **请求体**：
```json
{
  "dueDate": "2023-07-20",
  "items": [
    {
      "productId": "prod-001",
      "quantity": 10,
      "unitPrice": 4424.78,
      "taxRate": 13
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

### 7.5 删除销售发票
- **接口路径**：`/api/v1/sales/invoices/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 发票ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 7.6 标记发票已收款
- **接口路径**：`/api/v1/sales/invoices/{id}/receive`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 发票ID |
- **请求体**：
```json
{
  "amount": 50000,
  "receiveDate": "2023-07-10",
  "paymentMethod": "bank_transfer",
  "remarks": "全额收款"
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

## 8. 销售退货管理API

### 8.1 获取销售退货列表
- **接口路径**：`/api/v1/sales/returns`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | returnNo | string | 否 | 退货单编号 |
  | customerId | string | 否 | 客户ID |
  | customerName | string | 否 | 客户名称 |
  | orderId | string | 否 | 销售订单ID |
  | orderNo | string | 否 | 销售订单编号 |
  | returnDateStart | string | 否 | 退货开始日期，格式：YYYY-MM-DD |
  | returnDateEnd | string | 否 | 退货结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "return-001",
        "returnNo": "RTN2023060001",
        "customerId": "customer-001",
        "customerName": "北京客户有限公司",
        "orderId": "order-001",
        "orderNo": "SO2023060001",
        "returnDate": "2023-06-20",
        "totalAmount": 5000,
        "status": "completed",
        "createdBy": "admin",
        "createdAt": "2023-06-20T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-20T09:00:00Z"
      }
    ],
    "total": 20,
    "page": 1,
    "pageSize": 10
  }
}
```

### 8.2 获取销售退货详情
- **接口路径**：`/api/v1/sales/returns/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 退货单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "return-001",
    "returnNo": "RTN2023060001",
    "customerId": "customer-001",
    "customerName": "北京客户有限公司",
    "contact": "李四",
    "phone": "13900139001",
    "orderId": "order-001",
    "orderNo": "SO2023060001",
    "returnDate": "2023-06-20",
    "reason": "质量问题",
    "remarks": "退货处理",
    "totalAmount": 5000,
    "status": "completed",
    "items": [
      {
        "id": "item-001",
        "orderItemId": "order-item-001",
        "productId": "prod-001",
        "productCode": "PROD001",
        "productName": "笔记本电脑",
        "specification": "i7 16GB 512GB",
        "unit": "台",
        "quantity": 1,
        "unitPrice": 5000,
        "amount": 5000,
        "reason": "质量问题"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-20T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-20T09:00:00Z"
  }
}
```

### 8.3 创建销售退货
- **接口路径**：`/api/v1/sales/returns`
- **请求方法**：POST
- **请求体**：
```json
{
  "customerId": "customer-001",
  "orderId": "order-001",
  "returnDate": "2023-06-20",
  "reason": "质量问题",
  "remarks": "退货处理",
  "items": [
    {
      "orderItemId": "order-item-001",
      "quantity": 1,
      "unitPrice": 5000,
      "reason": "质量问题"
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
    "id": "return-001",
    "returnNo": "RTN2023060001"
  }
}
```

### 8.4 更新销售退货
- **接口路径**：`/api/v1/sales/returns/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 退货单ID |
- **请求体**：
```json
{
  "remarks": "退货处理（更新）",
  "items": [
    {
      "orderItemId": "order-item-001",
      "quantity": 1,
      "unitPrice": 5000,
      "reason": "质量问题"
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

### 8.5 删除销售退货
- **接口路径**：`/api/v1/sales/returns/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 退货单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 9. 销售报表管理API

### 9.1 获取销售订单执行报表
- **接口路径**：`/api/v1/sales/reports/order-execution`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
  | customerId | string | 否 | 客户ID |
  | productId | string | 否 | 产品ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "period": "2023-06-01 至 2023-06-30",
    "totalOrders": 10,
    "totalAmount": 1000000,
    "orders": [
      {
        "orderNo": "SO2023060001",
        "customerName": "北京客户有限公司",
        "orderDate": "2023-06-01",
        "amount": 100000,
        "status": "fully_delivered",
        "deliveredAmount": 100000,
        "deliveredRate": 100
      }
    ]
  }
}
```

### 9.2 获取客户销售分析报表
- **接口路径**：`/api/v1/sales/reports/customer-analysis`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "period": "2023-01-01 至 2023-06-30",
    "customers": [
      {
        "customerId": "customer-001",
        "customerName": "北京客户有限公司",
        "totalOrders": 5,
        "totalAmount": 500000,
        "avgOrderAmount": 100000,
        "paymentRate": 95
      }
    ]
  }
}
```

### 9.3 获取产品销售趋势报表
- **接口路径**：`/api/v1/sales/reports/product-trend`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
  | productId | string | 否 | 产品ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "period": "2023-01-01 至 2023-06-30",
    "trends": [
      {
        "period": "2023-01",
        "quantity": 100,
        "amount": 500000
      },
      {
        "period": "2023-02",
        "quantity": 120,
        "amount": 600000
      }
    ]
  }
}
```

### 9.4 导出销售报表
- **接口路径**：`/api/v1/sales/reports/export`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | type | string | 是 | 报表类型（order_execution, customer_analysis, product_trend） |
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
  | customerId | string | 否 | 客户ID |
  | productId | string | 否 | 产品ID |
- **响应格式**：Excel文件

## 10. 错误码定义

| 错误码 | 描述 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |
| 501 | 接口未实现 |

## 11. 附录

### 11.1 参考文档
- 《ERP系统销售管理模块设计与实现》
- 《销售管理实务》
- 《API设计最佳实践》

### 11.2 接口测试
- 使用Postman或Swagger进行接口测试
- 测试用例应覆盖正常流程和异常流程
- 测试数据应符合业务规则和数据约束