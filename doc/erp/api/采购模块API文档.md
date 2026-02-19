# 采购模块API文档

## 1. 文档概述

### 1.1 文档目的
本文档定义了ERP系统采购模块的API接口规范，包括供应商管理、采购计划、采购订单、采购收货、采购发票、采购退货和采购报表等功能的API接口设计。旨在为前端开发和后端开发提供明确的接口规范，确保系统集成的顺利进行。

### 1.2 术语定义
| 术语 | 解释 |
|------|------|
| PO | 采购订单（Purchase Order），企业向供应商发出的采购请求 |
| PR | 采购申请（Purchase Requisition），内部部门提出的采购需求 |
| GRN | 收货单（Goods Receipt Note），记录货物接收的单据 |
| AP | 应付账款（Accounts Payable），企业欠供应商的款项 |
| RFQ | 询价单（Request for Quotation），向供应商请求报价的单据 |
| Vendor | 供应商，提供货物或服务的外部单位 |
| Item | 物料，采购的商品或服务 |
| Lead Time | 前置时间，从下单到收货的时间间隔 |

## 2. 通用规范

### 2.1 接口风格
- **URL格式**：`/api/v1/po/{resource}`
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

## 3. 供应商管理API

### 3.1 获取供应商列表
- **接口路径**：`/api/v1/po/vendors`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | vendorNo | string | 否 | 供应商编号 |
  | name | string | 否 | 供应商名称 |
  | contact | string | 否 | 联系人 |
  | phone | string | 否 | 电话 |
  | status | string | 否 | 状态（active, inactive） |
  | category | string | 否 | 供应商类别 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "vendor-001",
        "vendorNo": "VEN001",
        "name": "北京供应商有限公司",
        "contact": "张三",
        "phone": "13800138001",
        "email": "contact@example.com",
        "address": "北京市海淀区",
        "taxNo": "110101199001011234",
        "category": "电子元件",
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

### 3.2 获取供应商详情
- **接口路径**：`/api/v1/po/vendors/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 供应商ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "vendor-001",
    "vendorNo": "VEN001",
    "name": "北京供应商有限公司",
    "contact": "张三",
    "phone": "13800138001",
    "email": "contact@example.com",
    "address": "北京市海淀区",
    "taxNo": "110101199001011234",
    "bankName": "工商银行",
    "bankAccount": "6222021234567890123",
    "category": "电子元件",
    "creditLimit": 100000,
    "leadTime": 7,
    "remarks": "优质供应商",
    "status": "active",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 3.3 创建供应商
- **接口路径**：`/api/v1/po/vendors`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "北京供应商有限公司",
  "contact": "张三",
  "phone": "13800138001",
  "email": "contact@example.com",
  "address": "北京市海淀区",
  "taxNo": "110101199001011234",
  "bankName": "工商银行",
  "bankAccount": "6222021234567890123",
  "category": "电子元件",
  "creditLimit": 100000,
  "leadTime": 7,
  "remarks": "优质供应商"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "vendor-001",
    "vendorNo": "VEN001"
  }
}
```

### 3.4 更新供应商
- **接口路径**：`/api/v1/po/vendors/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 供应商ID |
- **请求体**：
```json
{
  "name": "北京供应商有限公司（更新）",
  "contact": "李四",
  "phone": "13900139001",
  "creditLimit": 150000,
  "leadTime": 5
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

### 3.5 删除供应商
- **接口路径**：`/api/v1/po/vendors/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 供应商ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 3.6 供应商评估
- **接口路径**：`/api/v1/po/vendors/{id}/evaluate`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 供应商ID |
- **请求体**：
```json
{
  "period": "2023-06",
  "qualityScore": 95,
  "deliveryScore": 90,
  "serviceScore": 85,
  "priceScore": 80,
  "remarks": "表现良好"
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

## 4. 采购计划管理API

### 4.1 获取采购计划列表
- **接口路径**：`/api/v1/po/plans`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | planNo | string | 否 | 计划编号 |
  | name | string | 否 | 计划名称 |
  | period | string | 否 | 计划期间，格式：YYYY-MM |
  | status | string | 否 | 状态（draft, approved, executed, cancelled） |
  | departmentId | string | 否 | 部门ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "plan-001",
        "planNo": "PLN2023060001",
        "name": "2023年6月采购计划",
        "period": "2023-06",
        "departmentId": "dept-001",
        "departmentName": "技术部",
        "totalAmount": 50000,
        "status": "approved",
        "createdBy": "admin",
        "createdAt": "2023-05-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-05-10T08:00:00Z"
      }
    ],
    "total": 50,
    "page": 1,
    "pageSize": 10
  }
}
```

### 4.2 获取采购计划详情
- **接口路径**：`/api/v1/po/plans/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 计划ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "plan-001",
    "planNo": "PLN2023060001",
    "name": "2023年6月采购计划",
    "period": "2023-06",
    "departmentId": "dept-001",
    "departmentName": "技术部",
    "description": "技术部6月所需物料采购计划",
    "totalAmount": 50000,
    "status": "approved",
    "items": [
      {
        "id": "item-001",
        "materialId": "mat-001",
        "materialCode": "MAT001",
        "materialName": "内存条",
        "specification": "16GB DDR4",
        "quantity": 10,
        "unitPrice": 3000,
        "amount": 30000,
        "vendorId": "vendor-001",
        "vendorName": "北京供应商有限公司",
        "deliveryDate": "2023-06-15"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-05-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-05-10T08:00:00Z"
  }
}
```

### 4.3 创建采购计划
- **接口路径**：`/api/v1/po/plans`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "2023年6月采购计划",
  "period": "2023-06",
  "departmentId": "dept-001",
  "description": "技术部6月所需物料采购计划",
  "items": [
    {
      "materialId": "mat-001",
      "quantity": 10,
      "unitPrice": 3000,
      "vendorId": "vendor-001",
      "deliveryDate": "2023-06-15"
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
    "id": "plan-001",
    "planNo": "PLN2023060001"
  }
}
```

### 4.4 更新采购计划
- **接口路径**：`/api/v1/po/plans/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 计划ID |
- **请求体**：
```json
{
  "name": "2023年6月采购计划（更新）",
  "description": "技术部6月所需物料采购计划（更新）",
  "items": [
    {
      "materialId": "mat-001",
      "quantity": 15,
      "unitPrice": 2900,
      "vendorId": "vendor-001",
      "deliveryDate": "2023-06-15"
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

### 4.5 删除采购计划
- **接口路径**：`/api/v1/po/plans/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 计划ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 4.6 审批采购计划
- **接口路径**：`/api/v1/po/plans/{id}/approve`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 计划ID |
- **请求体**：
```json
{
  "approved": true,
  "remarks": "同意采购计划"
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

## 5. 采购订单管理API

### 5.1 获取采购订单列表
- **接口路径**：`/api/v1/po/orders`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | orderNo | string | 否 | 订单编号 |
  | vendorId | string | 否 | 供应商ID |
  | vendorName | string | 否 | 供应商名称 |
  | status | string | 否 | 状态（draft, approved, sent, partially_received, fully_received, cancelled） |
  | createDateStart | string | 否 | 创建开始日期，格式：YYYY-MM-DD |
  | createDateEnd | string | 否 | 创建结束日期，格式：YYYY-MM-DD |
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
        "orderNo": "PO2023060001",
        "vendorId": "vendor-001",
        "vendorName": "北京供应商有限公司",
        "vendorContact": "张三",
        "vendorPhone": "13800138001",
        "orderDate": "2023-06-01",
        "deliveryDate": "2023-06-15",
        "totalAmount": 30000,
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

### 5.2 获取采购订单详情
- **接口路径**：`/api/v1/po/orders/{id}`
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
    "orderNo": "PO2023060001",
    "vendorId": "vendor-001",
    "vendorName": "北京供应商有限公司",
    "vendorContact": "张三",
    "vendorPhone": "13800138001",
    "vendorAddress": "北京市海淀区",
    "orderDate": "2023-06-01",
    "deliveryDate": "2023-06-15",
    "paymentTerms": "30天内付款",
    "shippingTerms": "运费由供应商承担",
    "remarks": "急需物料，请尽快发货",
    "totalAmount": 30000,
    "status": "approved",
    "items": [
      {
        "id": "item-001",
        "materialId": "mat-001",
        "materialCode": "MAT001",
        "materialName": "内存条",
        "specification": "16GB DDR4",
        "unit": "条",
        "quantity": 10,
        "unitPrice": 3000,
        "amount": 30000,
        "receivedQuantity": 0,
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

### 5.3 创建采购订单
- **接口路径**：`/api/v1/po/orders`
- **请求方法**：POST
- **请求体**：
```json
{
  "vendorId": "vendor-001",
  "orderDate": "2023-06-01",
  "deliveryDate": "2023-06-15",
  "paymentTerms": "30天内付款",
  "shippingTerms": "运费由供应商承担",
  "remarks": "急需物料，请尽快发货",
  "items": [
    {
      "materialId": "mat-001",
      "quantity": 10,
      "unitPrice": 3000,
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
    "orderNo": "PO2023060001"
  }
}
```

### 5.4 更新采购订单
- **接口路径**：`/api/v1/po/orders/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 订单ID |
- **请求体**：
```json
{
  "deliveryDate": "2023-06-10",
  "remarks": "急需物料，请尽快发货（更新）",
  "items": [
    {
      "materialId": "mat-001",
      "quantity": 15,
      "unitPrice": 2900,
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

### 5.5 删除采购订单
- **接口路径**：`/api/v1/po/orders/{id}`
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

### 5.6 审批采购订单
- **接口路径**：`/api/v1/po/orders/{id}/approve`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 订单ID |
- **请求体**：
```json
{
  "approved": true,
  "remarks": "同意采购订单"
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

### 5.7 发送采购订单
- **接口路径**：`/api/v1/po/orders/{id}/send`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 订单ID |
- **请求体**：
```json
{
  "sendMethod": "email",
  "recipient": "contact@example.com"
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

### 5.8 取消采购订单
- **接口路径**：`/api/v1/po/orders/{id}/cancel`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 订单ID |
- **请求体**：
```json
{
  "reason": "项目变更，取消采购"
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

## 6. 采购收货管理API

### 6.1 获取采购收货列表
- **接口路径**：`/api/v1/po/receipts`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | receiptNo | string | 否 | 收货单编号 |
  | orderId | string | 否 | 采购订单ID |
  | orderNo | string | 否 | 采购订单编号 |
  | vendorId | string | 否 | 供应商ID |
  | vendorName | string | 否 | 供应商名称 |
  | receiptDateStart | string | 否 | 收货开始日期，格式：YYYY-MM-DD |
  | receiptDateEnd | string | 否 | 收货结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "receipt-001",
        "receiptNo": "GRN2023060001",
        "orderId": "order-001",
        "orderNo": "PO2023060001",
        "vendorId": "vendor-001",
        "vendorName": "北京供应商有限公司",
        "receiptDate": "2023-06-15",
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

### 6.2 获取采购收货详情
- **接口路径**：`/api/v1/po/receipts/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 收货单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "receipt-001",
    "receiptNo": "GRN2023060001",
    "orderId": "order-001",
    "orderNo": "PO2023060001",
    "vendorId": "vendor-001",
    "vendorName": "北京供应商有限公司",
    "receiptDate": "2023-06-15",
    "remarks": "收货正常",
    "status": "completed",
    "items": [
      {
        "id": "item-001",
        "orderItemId": "order-item-001",
        "materialId": "mat-001",
        "materialCode": "MAT001",
        "materialName": "内存条",
        "specification": "16GB DDR4",
        "unit": "条",
        "orderedQuantity": 10,
        "receivedQuantity": 10,
        "qualifiedQuantity": 10,
        "defectiveQuantity": 0,
        "remark": "质量良好"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-15T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-15T09:00:00Z"
  }
}
```

### 6.3 创建采购收货
- **接口路径**：`/api/v1/po/receipts`
- **请求方法**：POST
- **请求体**：
```json
{
  "orderId": "order-001",
  "receiptDate": "2023-06-15",
  "remarks": "收货正常",
  "items": [
    {
      "orderItemId": "order-item-001",
      "receivedQuantity": 10,
      "qualifiedQuantity": 10,
      "defectiveQuantity": 0,
      "remark": "质量良好"
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
    "id": "receipt-001",
    "receiptNo": "GRN2023060001"
  }
}
```

### 6.4 更新采购收货
- **接口路径**：`/api/v1/po/receipts/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 收货单ID |
- **请求体**：
```json
{
  "remarks": "收货正常（更新）",
  "items": [
    {
      "orderItemId": "order-item-001",
      "receivedQuantity": 10,
      "qualifiedQuantity": 9,
      "defectiveQuantity": 1,
      "remark": "有1个质量问题"
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

### 6.5 删除采购收货
- **接口路径**：`/api/v1/po/receipts/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 收货单ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 7. 采购发票管理API

### 7.1 获取采购发票列表
- **接口路径**：`/api/v1/po/invoices`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | invoiceNo | string | 否 | 发票编号 |
  | vendorId | string | 否 | 供应商ID |
  | vendorName | string | 否 | 供应商名称 |
  | orderId | string | 否 | 采购订单ID |
  | orderNo | string | 否 | 采购订单编号 |
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
        "vendorId": "vendor-001",
        "vendorName": "北京供应商有限公司",
        "orderId": "order-001",
        "orderNo": "PO2023060001",
        "invoiceDate": "2023-06-15",
        "dueDate": "2023-07-15",
        "totalAmount": 30000,
        "paidAmount": 0,
        "balanceAmount": 30000,
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

### 7.2 获取采购发票详情
- **接口路径**：`/api/v1/po/invoices/{id}`
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
    "vendorId": "vendor-001",
    "vendorName": "北京供应商有限公司",
    "vendorTaxNo": "110101199001011234",
    "orderId": "order-001",
    "orderNo": "PO2023060001",
    "invoiceDate": "2023-06-15",
    "dueDate": "2023-07-15",
    "totalAmount": 30000,
    "taxAmount": 3451.33,
    "netAmount": 26548.67,
    "paidAmount": 0,
    "balanceAmount": 30000,
    "status": "unpaid",
    "items": [
      {
        "id": "item-001",
        "materialId": "mat-001",
        "materialCode": "MAT001",
        "materialName": "内存条",
        "specification": "16GB DDR4",
        "quantity": 10,
        "unitPrice": 2654.87,
        "taxRate": 13,
        "taxAmount": 345.13,
        "totalAmount": 3000
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-15T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-15T08:00:00Z"
  }
}
```

### 7.3 创建采购发票
- **接口路径**：`/api/v1/po/invoices`
- **请求方法**：POST
- **请求体**：
```json
{
  "vendorId": "vendor-001",
  "orderId": "order-001",
  "invoiceDate": "2023-06-15",
  "dueDate": "2023-07-15",
  "items": [
    {
      "materialId": "mat-001",
      "quantity": 10,
      "unitPrice": 2654.87,
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

### 7.4 更新采购发票
- **接口路径**：`/api/v1/po/invoices/{id}`
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
      "materialId": "mat-001",
      "quantity": 10,
      "unitPrice": 2654.87,
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

### 7.5 删除采购发票
- **接口路径**：`/api/v1/po/invoices/{id}`
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

### 7.6 标记发票已支付
- **接口路径**：`/api/v1/po/invoices/{id}/pay`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 发票ID |
- **请求体**：
```json
{
  "amount": 30000,
  "paymentDate": "2023-07-10",
  "paymentMethod": "bank_transfer",
  "remarks": "全额支付"
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

## 8. 采购退货管理API

### 8.1 获取采购退货列表
- **接口路径**：`/api/v1/po/returns`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | returnNo | string | 否 | 退货单编号 |
  | vendorId | string | 否 | 供应商ID |
  | vendorName | string | 否 | 供应商名称 |
  | orderId | string | 否 | 采购订单ID |
  | orderNo | string | 否 | 采购订单编号 |
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
        "vendorId": "vendor-001",
        "vendorName": "北京供应商有限公司",
        "orderId": "order-001",
        "orderNo": "PO2023060001",
        "returnDate": "2023-06-20",
        "totalAmount": 3000,
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

### 8.2 获取采购退货详情
- **接口路径**：`/api/v1/po/returns/{id}`
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
    "vendorId": "vendor-001",
    "vendorName": "北京供应商有限公司",
    "orderId": "order-001",
    "orderNo": "PO2023060001",
    "returnDate": "2023-06-20",
    "reason": "质量问题",
    "remarks": "退货处理",
    "totalAmount": 3000,
    "status": "completed",
    "items": [
      {
        "id": "item-001",
        "materialId": "mat-001",
        "materialCode": "MAT001",
        "materialName": "内存条",
        "specification": "16GB DDR4",
        "unit": "条",
        "quantity": 1,
        "unitPrice": 3000,
        "amount": 3000,
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

### 8.3 创建采购退货
- **接口路径**：`/api/v1/po/returns`
- **请求方法**：POST
- **请求体**：
```json
{
  "vendorId": "vendor-001",
  "orderId": "order-001",
  "returnDate": "2023-06-20",
  "reason": "质量问题",
  "remarks": "退货处理",
  "items": [
    {
      "materialId": "mat-001",
      "quantity": 1,
      "unitPrice": 3000,
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

### 8.4 更新采购退货
- **接口路径**：`/api/v1/po/returns/{id}`
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
      "materialId": "mat-001",
      "quantity": 1,
      "unitPrice": 3000,
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

### 8.5 删除采购退货
- **接口路径**：`/api/v1/po/returns/{id}`
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

## 9. 采购报表管理API

### 9.1 获取采购订单执行报表
- **接口路径**：`/api/v1/po/reports/order-execution`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
  | vendorId | string | 否 | 供应商ID |
  | departmentId | string | 否 | 部门ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "period": "2023-06-01 至 2023-06-30",
    "totalOrders": 10,
    "totalAmount": 500000,
    "orders": [
      {
        "orderNo": "PO2023060001",
        "vendorName": "北京供应商有限公司",
        "orderDate": "2023-06-01",
        "amount": 50000,
        "status": "fully_received",
        "receivedAmount": 50000,
        "receivedRate": 100
      }
    ]
  }
}
```

### 9.2 获取供应商采购分析报表
- **接口路径**：`/api/v1/po/reports/vendor-analysis`
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
    "vendors": [
      {
        "vendorId": "vendor-001",
        "vendorName": "北京供应商有限公司",
        "totalOrders": 5,
        "totalAmount": 250000,
        "deliveryRate": 95,
        "qualityRate": 98
      }
    ]
  }
}
```

### 9.3 获取物料采购趋势报表
- **接口路径**：`/api/v1/po/reports/material-trend`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
  | materialId | string | 否 | 物料ID |
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
        "amount": 100000
      },
      {
        "period": "2023-02",
        "quantity": 120,
        "amount": 120000
      }
    ]
  }
}
```

### 9.4 导出采购报表
- **接口路径**：`/api/v1/po/reports/export`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | type | string | 是 | 报表类型（order_execution, vendor_analysis, material_trend） |
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
  | vendorId | string | 否 | 供应商ID |
  | departmentId | string | 否 | 部门ID |
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
- 《ERP系统采购管理模块设计与实现》
- 《采购管理实务》
- 《API设计最佳实践》

### 11.2 接口测试
- 使用Postman或Swagger进行接口测试
- 测试用例应覆盖正常流程和异常流程
- 测试数据应符合业务规则和数据约束