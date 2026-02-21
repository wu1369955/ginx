# 财务模块API文档

## 1. 文档概述

### 1.1 文档目的
本文档定义了ERP系统财务模块的API接口规范，包括会计科目、凭证、供应商、客户、固定资产、报表、预算和成本管理等功能的API接口设计。旨在为前端开发和后端开发提供明确的接口规范，确保系统集成的顺利进行。

### 1.2 术语定义
| 术语 | 解释 |
|------|------|
| GL | 总账（General Ledger），企业的核心财务记录 |
| AP | 应付账款（Accounts Payable），企业欠供应商的款项 |
| AR | 应收账款（Accounts Receivable），客户欠企业的款项 |
| FA | 固定资产（Fixed Assets），企业长期拥有的资产 |
| JC |  journal Entry |
| COA | 会计科目表（Chart of Accounts） |
| VAT | 增值税（Value Added Tax） |
| TB | 试算平衡表（Trial Balance） |
| BS | 资产负债表（Balance Sheet） |
| P&L | 利润表（Profit and Loss） |

## 2. 通用规范

### 2.1 接口风格
- **URL格式**：`/api/v1/finance/{resource}`
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

## 3. 会计科目管理API

### 3.1 获取会计科目列表
- **接口路径**：`/api/v1/finance/accounts`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | code | string | 否 | 科目编码 |
  | name | string | 否 | 科目名称 |
  | type | string | 否 | 科目类型（asset, liability, equity, revenue, expense） |
  | status | string | 否 | 状态（active, inactive） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "acc-001",
        "code": "1001",
        "name": "库存现金",
        "type": "asset",
        "level": 1,
        "parentId": "",
        "parentName": "",
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

### 3.2 获取会计科目树形结构
- **接口路径**：`/api/v1/finance/accounts/tree`
- **请求方法**：GET
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": "acc-001",
      "code": "1000",
      "name": "资产",
      "type": "asset",
      "level": 1,
      "parentId": "",
      "children": [
        {
          "id": "acc-002",
          "code": "1001",
          "name": "库存现金",
          "type": "asset",
          "level": 2,
          "parentId": "acc-001",
          "children": []
        }
      ]
    }
  ]
}
```

### 3.3 获取会计科目详情
- **接口路径**：`/api/v1/finance/accounts/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 科目ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "acc-001",
    "code": "1001",
    "name": "库存现金",
    "type": "asset",
    "level": 1,
    "parentId": "",
    "parentName": "",
    "description": "企业的库存现金",
    "status": "active",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 3.4 创建会计科目
- **接口路径**：`/api/v1/finance/accounts`
- **请求方法**：POST
- **请求体**：
```json
{
  "code": "1001",
  "name": "库存现金",
  "type": "asset",
  "parentId": "",
  "description": "企业的库存现金"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "acc-001"
  }
}
```

### 3.5 更新会计科目
- **接口路径**：`/api/v1/finance/accounts/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 科目ID |
- **请求体**：
```json
{
  "name": "库存现金",
  "description": "企业的库存现金（更新）",
  "status": "active"
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

### 3.6 删除会计科目
- **接口路径**：`/api/v1/finance/accounts/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 科目ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 4. 凭证管理API

### 4.1 获取凭证列表
- **接口路径**：`/api/v1/finance/journals`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | journalNo | string | 否 | 凭证编号 |
  | dateStart | string | 否 | 开始日期，格式：YYYY-MM-DD |
  | dateEnd | string | 否 | 结束日期，格式：YYYY-MM-DD |
  | status | string | 否 | 状态（draft, posted, cancelled） |
  | type | string | 否 | 类型（general, cash, bank） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "journal-001",
        "journalNo": "J2023060001",
        "date": "2023-06-01",
        "type": "general",
        "summary": "购买办公用品",
        "totalDebit": 1000,
        "totalCredit": 1000,
        "status": "posted",
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

### 4.2 获取凭证详情
- **接口路径**：`/api/v1/finance/journals/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 凭证ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "journal-001",
    "journalNo": "J2023060001",
    "date": "2023-06-01",
    "type": "general",
    "summary": "购买办公用品",
    "totalDebit": 1000,
    "totalCredit": 1000,
    "status": "posted",
    "entries": [
      {
        "id": "entry-001",
        "accountId": "acc-001",
        "accountCode": "1001",
        "accountName": "库存现金",
        "debit": 0,
        "credit": 1000
      },
      {
        "id": "entry-002",
        "accountId": "acc-002",
        "accountCode": "6602",
        "accountName": "管理费用",
        "debit": 1000,
        "credit": 0
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-01T09:00:00Z"
  }
}
```

### 4.3 创建凭证
- **接口路径**：`/api/v1/finance/journals`
- **请求方法**：POST
- **请求体**：
```json
{
  "date": "2023-06-01",
  "type": "general",
  "summary": "购买办公用品",
  "entries": [
    {
      "accountId": "acc-001",
      "debit": 0,
      "credit": 1000
    },
    {
      "accountId": "acc-002",
      "debit": 1000,
      "credit": 0
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
    "id": "journal-001",
    "journalNo": "J2023060001"
  }
}
```

### 4.4 更新凭证
- **接口路径**：`/api/v1/finance/journals/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 凭证ID |
- **请求体**：
```json
{
  "summary": "购买办公用品（更新）",
  "entries": [
    {
      "accountId": "acc-001",
      "debit": 0,
      "credit": 1200
    },
    {
      "accountId": "acc-002",
      "debit": 1200,
      "credit": 0
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

### 4.5 删除凭证
- **接口路径**：`/api/v1/finance/journals/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 凭证ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 4.6 过账凭证
- **接口路径**：`/api/v1/finance/journals/{id}/post`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 凭证ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 4.7 取消过账
- **接口路径**：`/api/v1/finance/journals/{id}/cancel`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 凭证ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 5. 供应商管理API

### 5.1 获取供应商列表
- **接口路径**：`/api/v1/finance/vendors`
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

### 5.2 获取供应商详情
- **接口路径**：`/api/v1/finance/vendors/{id}`
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
    "creditLimit": 100000,
    "status": "active",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 5.3 创建供应商
- **接口路径**：`/api/v1/finance/vendors`
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
  "creditLimit": 100000
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

### 5.4 更新供应商
- **接口路径**：`/api/v1/finance/vendors/{id}`
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
  "creditLimit": 150000
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

### 5.5 删除供应商
- **接口路径**：`/api/v1/finance/vendors/{id}`
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

### 5.6 获取供应商余额
- **接口路径**：`/api/v1/finance/vendors/{id}/balance`
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
    "totalAmount": 50000,
    "paidAmount": 30000,
    "balanceAmount": 20000
  }
}
```

## 6. 客户管理API

### 6.1 获取客户列表
- **接口路径**：`/api/v1/finance/customers`
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
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "cust-001",
        "customerNo": "CUS001",
        "name": "上海客户有限公司",
        "contact": "王五",
        "phone": "13700137001",
        "email": "contact@example.com",
        "address": "上海市浦东新区",
        "taxNo": "310101199001011234",
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

### 6.2 获取客户详情
- **接口路径**：`/api/v1/finance/customers/{id}`
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
    "id": "cust-001",
    "customerNo": "CUS001",
    "name": "上海客户有限公司",
    "contact": "王五",
    "phone": "13700137001",
    "email": "contact@example.com",
    "address": "上海市浦东新区",
    "taxNo": "310101199001011234",
    "bankName": "建设银行",
    "bankAccount": "6227001234567890123",
    "creditLimit": 200000,
    "status": "active",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 6.3 创建客户
- **接口路径**：`/api/v1/finance/customers`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "上海客户有限公司",
  "contact": "王五",
  "phone": "13700137001",
  "email": "contact@example.com",
  "address": "上海市浦东新区",
  "taxNo": "310101199001011234",
  "bankName": "建设银行",
  "bankAccount": "6227001234567890123",
  "creditLimit": 200000
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "cust-001",
    "customerNo": "CUS001"
  }
}
```

### 6.4 更新客户
- **接口路径**：`/api/v1/finance/customers/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 客户ID |
- **请求体**：
```json
{
  "name": "上海客户有限公司（更新）",
  "contact": "赵六",
  "phone": "13600136001",
  "creditLimit": 250000
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

### 6.5 删除客户
- **接口路径**：`/api/v1/finance/customers/{id}`
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

### 6.6 获取客户余额
- **接口路径**：`/api/v1/finance/customers/{id}/balance`
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
    "totalAmount": 80000,
    "paidAmount": 50000,
    "balanceAmount": 30000
  }
}
```

## 7. 固定资产管理API

### 7.1 获取固定资产列表
- **接口路径**：`/api/v1/finance/assets`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | assetNo | string | 否 | 资产编号 |
  | name | string | 否 | 资产名称 |
  | category | string | 否 | 资产类别 |
  | status | string | 否 | 状态（active, disposed） |
  | acquisitionDateStart | string | 否 | 购置开始日期，格式：YYYY-MM-DD |
  | acquisitionDateEnd | string | 否 | 购置结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "asset-001",
        "assetNo": "AST001",
        "name": "笔记本电脑",
        "category": "电子设备",
        "acquisitionDate": "2023-01-01",
        "cost": 8000,
        "salvageValue": 800,
        "usefulLife": 36,
        "accumulatedDepreciation": 2000,
        "netBookValue": 6000,
        "status": "active",
        "createdBy": "admin",
        "createdAt": "2023-01-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-01-01T08:00:00Z"
      }
    ],
    "total": 50,
    "page": 1,
    "pageSize": 10
  }
}
```

### 7.2 获取固定资产详情
- **接口路径**：`/api/v1/finance/assets/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 资产ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "asset-001",
    "assetNo": "AST001",
    "name": "笔记本电脑",
    "category": "电子设备",
    "description": "办公用笔记本电脑",
    "acquisitionDate": "2023-01-01",
    "cost": 8000,
    "salvageValue": 800,
    "usefulLife": 36,
    "depreciationMethod": "straight_line",
    "accumulatedDepreciation": 2000,
    "netBookValue": 6000,
    "location": "技术部",
    "responsiblePerson": "张三",
    "status": "active",
    "createdBy": "admin",
    "createdAt": "2023-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-01-01T08:00:00Z"
  }
}
```

### 7.3 创建固定资产
- **接口路径**：`/api/v1/finance/assets`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "笔记本电脑",
  "category": "电子设备",
  "description": "办公用笔记本电脑",
  "acquisitionDate": "2023-01-01",
  "cost": 8000,
  "salvageValue": 800,
  "usefulLife": 36,
  "depreciationMethod": "straight_line",
  "location": "技术部",
  "responsiblePerson": "张三"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "asset-001",
    "assetNo": "AST001"
  }
}
```

### 7.4 更新固定资产
- **接口路径**：`/api/v1/finance/assets/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 资产ID |
- **请求体**：
```json
{
  "name": "笔记本电脑（更新）",
  "description": "办公用笔记本电脑（更新）",
  "location": "市场部",
  "responsiblePerson": "李四"
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

### 7.5 删除固定资产
- **接口路径**：`/api/v1/finance/assets/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 资产ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 7.6 计提折旧
- **接口路径**：`/api/v1/finance/assets/depreciate`
- **请求方法**：POST
- **请求体**：
```json
{
  "period": "2023-06",
  "assetIds": ["asset-001", "asset-002"]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "depreciatedCount": 2,
    "totalDepreciationAmount": 1000
  }
}
```

### 7.7 资产处置
- **接口路径**：`/api/v1/finance/assets/{id}/dispose`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 资产ID |
- **请求体**：
```json
{
  "disposalDate": "2023-06-30",
  "disposalMethod": "sale",
  "disposalAmount": 5000,
  "remarks": "资产出售"
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

## 8. 报表管理API

### 8.1 获取资产负债表
- **接口路径**：`/api/v1/finance/reports/balance-sheet`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | date | string | 是 | 报表日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "date": "2023-06-30",
    "assets": {
      "currentAssets": 1000000,
      "nonCurrentAssets": 2000000,
      "totalAssets": 3000000
    },
    "liabilities": {
      "currentLiabilities": 800000,
      "nonCurrentLiabilities": 1200000,
      "totalLiabilities": 2000000
    },
    "equity": {
      "paidInCapital": 500000,
      "retainedEarnings": 500000,
      "totalEquity": 1000000
    }
  }
}
```

### 8.2 获取利润表
- **接口路径**：`/api/v1/finance/reports/profit-loss`
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
    "period": "2023-06-01 至 2023-06-30",
    "revenues": {
      "operatingRevenue": 500000,
      "otherRevenue": 50000,
      "totalRevenue": 550000
    },
    "expenses": {
      "costOfSales": 300000,
      "operatingExpenses": 100000,
      "administrativeExpenses": 50000,
      "financialExpenses": 20000,
      "totalExpenses": 470000
    },
    "profit": {
      "grossProfit": 200000,
      "operatingProfit": 50000,
      "netProfit": 80000
    }
  }
}
```

### 8.3 获取现金流量表
- **接口路径**：`/api/v1/finance/reports/cash-flow`
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
    "period": "2023-06-01 至 2023-06-30",
    "operatingActivities": {
      "cashInflows": 600000,
      "cashOutflows": 400000,
      "netCashFlow": 200000
    },
    "investingActivities": {
      "cashInflows": 50000,
      "cashOutflows": 150000,
      "netCashFlow": -100000
    },
    "financingActivities": {
      "cashInflows": 100000,
      "cashOutflows": 50000,
      "netCashFlow": 50000
    },
    "netIncreaseInCash": 150000,
    "beginningCashBalance": 200000,
    "endingCashBalance": 350000
  }
}
```

### 8.4 获取试算平衡表
- **接口路径**：`/api/v1/finance/reports/trial-balance`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | date | string | 是 | 报表日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "date": "2023-06-30",
    "accounts": [
      {
        "code": "1001",
        "name": "库存现金",
        "debit": 50000,
        "credit": 0
      },
      {
        "code": "1002",
        "name": "银行存款",
        "debit": 150000,
        "credit": 0
      }
    ],
    "totalDebit": 5000000,
    "totalCredit": 5000000
  }
}
```

### 8.5 导出报表
- **接口路径**：`/api/v1/finance/reports/export`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | type | string | 是 | 报表类型（balance_sheet, profit_loss, cash_flow, trial_balance） |
  | startDate | string | 否 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 否 | 结束日期，格式：YYYY-MM-DD |
  | date | string | 否 | 报表日期，格式：YYYY-MM-DD |
- **响应格式**：Excel文件

## 9. 预算管理API

### 9.1 获取预算列表
- **接口路径**：`/api/v1/finance/budgets`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | budgetNo | string | 否 | 预算编号 |
  | name | string | 否 | 预算名称 |
  | year | string | 否 | 预算年度 |
  | status | string | 否 | 状态（draft, approved, executed） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "budget-001",
        "budgetNo": "BUD20230001",
        "name": "2023年度预算",
        "year": "2023",
        "totalAmount": 5000000,
        "status": "approved",
        "createdBy": "admin",
        "createdAt": "2022-12-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2022-12-15T08:00:00Z"
      }
    ],
    "total": 10,
    "page": 1,
    "pageSize": 10
  }
}
```

### 9.2 获取预算详情
- **接口路径**：`/api/v1/finance/budgets/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 预算ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "budget-001",
    "budgetNo": "BUD20230001",
    "name": "2023年度预算",
    "year": "2023",
    "description": "2023年度公司整体预算",
    "totalAmount": 5000000,
    "status": "approved",
    "details": [
      {
        "department": "技术部",
        "amount": 2000000
      },
      {
        "department": "市场部",
        "amount": 1500000
      }
    ],
    "createdBy": "admin",
    "createdAt": "2022-12-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2022-12-15T08:00:00Z"
  }
}
```

### 9.3 创建预算
- **接口路径**：`/api/v1/finance/budgets`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "2023年度预算",
  "year": "2023",
  "description": "2023年度公司整体预算",
  "details": [
    {
      "department": "技术部",
      "amount": 2000000
    },
    {
      "department": "市场部",
      "amount": 1500000
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
    "id": "budget-001",
    "budgetNo": "BUD20230001"
  }
}
```

### 9.4 更新预算
- **接口路径**：`/api/v1/finance/budgets/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 预算ID |
- **请求体**：
```json
{
  "name": "2023年度预算（更新）",
  "description": "2023年度公司整体预算（更新）",
  "details": [
    {
      "department": "技术部",
      "amount": 2200000
    },
    {
      "department": "市场部",
      "amount": 1600000
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

### 9.5 删除预算
- **接口路径**：`/api/v1/finance/budgets/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 预算ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 9.6 审批预算
- **接口路径**：`/api/v1/finance/budgets/{id}/approve`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 预算ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 9.7 获取预算执行情况
- **接口路径**：`/api/v1/finance/budgets/{id}/execution`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 预算ID |
  | period | string | 是 | 期间，格式：YYYY-MM |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "budgetId": "budget-001",
    "period": "2023-06",
    "details": [
      {
        "department": "技术部",
        "budgetAmount": 166667,
        "actualAmount": 150000,
        "usageRate": 90
      },
      {
        "department": "市场部",
        "budgetAmount": 125000,
        "actualAmount": 130000,
        "usageRate": 104
      }
    ],
    "totalBudgetAmount": 291667,
    "totalActualAmount": 280000,
    "totalUsageRate": 96
  }
}
```

## 10. 成本管理API

### 10.1 获取成本中心列表
- **接口路径**：`/api/v1/finance/cost-centers`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | code | string | 否 | 成本中心编码 |
  | name | string | 否 | 成本中心名称 |
  | status | string | 否 | 状态（active, inactive） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "cc-001",
        "code": "CC001",
        "name": "技术部",
        "description": "技术部成本中心",
        "status": "active",
        "createdBy": "admin",
        "createdAt": "2020-01-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2020-01-01T08:00:00Z"
      }
    ],
    "total": 10,
    "page": 1,
    "pageSize": 10
  }
}
```

### 10.2 获取成本中心详情
- **接口路径**：`/api/v1/finance/cost-centers/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 成本中心ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "cc-001",
    "code": "CC001",
    "name": "技术部",
    "description": "技术部成本中心",
    "manager": "张三",
    "status": "active",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 10.3 创建成本中心
- **接口路径**：`/api/v1/finance/cost-centers`
- **请求方法**：POST
- **请求体**：
```json
{
  "code": "CC001",
  "name": "技术部",
  "description": "技术部成本中心",
  "manager": "张三"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "cc-001"
  }
}
```

### 10.4 更新成本中心
- **接口路径**：`/api/v1/finance/cost-centers/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 成本中心ID |
- **请求体**：
```json
{
  "name": "技术部（更新）",
  "description": "技术部成本中心（更新）",
  "manager": "李四"
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

### 10.5 删除成本中心
- **接口路径**：`/api/v1/finance/cost-centers/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 成本中心ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 10.6 获取成本分析
- **接口路径**：`/api/v1/finance/costs/analysis`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | period | string | 是 | 期间，格式：YYYY-MM |
  | costCenterId | string | 否 | 成本中心ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "period": "2023-06",
    "costs": [
      {
        "costCenterId": "cc-001",
        "costCenterName": "技术部",
        "amount": 150000,
        "breakdown": {
          "salary": 100000,
          "equipment": 30000,
          "other": 20000
        }
      }
    ],
    "totalCost": 150000
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
- 《ERP系统财务管理模块设计与实现》
- 《财务管理实务》
- 《API设计最佳实践》

### 12.2 接口测试
- 使用Postman或Swagger进行接口测试
- 测试用例应覆盖正常流程和异常流程
- 测试数据应符合业务规则和数据约束