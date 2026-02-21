# 客户关系管理模块API文档

## 1. 文档概述

### 1.1 文档目的
本文档定义了ERP系统中客户关系管理（CRM）模块的API接口规范，包括客户管理、联系人管理、客户沟通、客户反馈、客户分析等功能的API接口设计。旨在为前端开发和后端开发提供明确的接口规范，确保系统集成的顺利进行。

### 1.2 术语定义
| 术语 | 解释 |
|------|------|
| CRM | 客户关系管理（Customer Relationship Management），是企业用来管理与客户之间关系的系统 |
| Account | 客户账户，企业的客户实体 |
| Contact | 联系人，客户中的具体联系人 |
| Lead | 销售线索，潜在的客户机会 |
| Opportunity | 销售机会，有明确购买意向的潜在交易 |
| Activity | 客户活动，与客户相关的沟通或互动 |
| Case | 客户案例/反馈，客户提出的问题或需求 |
| Campaign | 营销活动，针对客户的推广活动 |
| Loyalty | 客户忠诚度，客户对企业的忠诚程度 |

## 2. 通用规范

### 2.1 接口风格
- **URL格式**：`/api/v1/crm/{resource}`
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
- **接口路径**：`/api/v1/crm/accounts`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | accountNo | string | 否 | 客户编号 |
  | name | string | 否 | 客户名称 |
  | status | string | 否 | 状态（active, inactive） |
  | industry | string | 否 | 行业 |
  | region | string | 否 | 地区 |
  | type | string | 否 | 客户类型（corporate, individual, government） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "account-001",
        "accountNo": "ACC001",
        "name": "北京科技有限公司",
        "type": "corporate",
        "industry": "科技",
        "region": "华北",
        "address": "北京市朝阳区",
        "contactPerson": "张三",
        "phone": "13800138001",
        "email": "contact@example.com",
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

### 3.2 获取客户详情
- **接口路径**：`/api/v1/crm/accounts/{id}`
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
    "id": "account-001",
    "accountNo": "ACC001",
    "name": "北京科技有限公司",
    "type": "corporate",
    "industry": "科技",
    "region": "华北",
    "address": "北京市朝阳区",
    "contactPerson": "张三",
    "phone": "13800138001",
    "email": "contact@example.com",
    "website": "www.example.com",
    "taxNo": "110101123456789",
    "creditLimit": 1000000,
    "creditUsed": 500000,
    "loyaltyLevel": "gold",
    "remarks": "重要客户",
    "status": "active",
    "contacts": [
      {
        "id": "contact-001",
        "name": "李四",
        "position": "采购经理",
        "phone": "13900139001",
        "email": "lisi@example.com"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 3.3 创建客户
- **接口路径**：`/api/v1/crm/accounts`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "北京科技有限公司",
  "type": "corporate",
  "industry": "科技",
  "region": "华北",
  "address": "北京市朝阳区",
  "contactPerson": "张三",
  "phone": "13800138001",
  "email": "contact@example.com",
  "website": "www.example.com",
  "taxNo": "110101123456789",
  "creditLimit": 1000000,
  "remarks": "重要客户"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "account-001",
    "accountNo": "ACC001"
  }
}
```

### 3.4 更新客户
- **接口路径**：`/api/v1/crm/accounts/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 客户ID |
- **请求体**：
```json
{
  "name": "北京科技有限公司（更新）",
  "contactPerson": "王五",
  "phone": "13700137001",
  "creditLimit": 1500000
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
- **接口路径**：`/api/v1/crm/accounts/{id}`
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

### 3.6 客户状态变更
- **接口路径**：`/api/v1/crm/accounts/{id}/status`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 客户ID |
- **请求体**：
```json
{
  "status": "inactive",
  "reason": "客户暂停合作"
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

### 3.7 客户信用评估
- **接口路径**：`/api/v1/crm/accounts/{id}/credit-assessment`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 客户ID |
- **请求体**：
```json
{
  "assessmentDate": "2023-06-01",
  "assessor": "admin",
  "creditScore": 85,
  "creditLimit": 2000000,
  "remarks": "信用良好，提高信用额度"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "assessmentId": "assessment-001"
  }
}
```

## 4. 联系人管理API

### 4.1 获取联系人列表
- **接口路径**：`/api/v1/crm/contacts`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | accountId | string | 否 | 客户ID |
  | name | string | 否 | 联系人姓名 |
  | position | string | 否 | 职位 |
  | phone | string | 否 | 电话 |
  | email | string | 否 | 邮箱 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "contact-001",
        "accountId": "account-001",
        "accountName": "北京科技有限公司",
        "name": "李四",
        "position": "采购经理",
        "phone": "13900139001",
        "email": "lisi@example.com",
        "mobile": "13900139001",
        "createdBy": "admin",
        "createdAt": "2020-01-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2020-01-01T08:00:00Z"
      }
    ],
    "total": 20,
    "page": 1,
    "pageSize": 10
  }
}
```

### 4.2 获取联系人详情
- **接口路径**：`/api/v1/crm/contacts/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 联系人ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "contact-001",
    "accountId": "account-001",
    "accountName": "北京科技有限公司",
    "name": "李四",
    "position": "采购经理",
    "phone": "13900139001",
    "email": "lisi@example.com",
    "mobile": "13900139001",
    "fax": "010-12345678",
    "address": "北京市朝阳区",
    "birthday": "1980-01-01",
    "preferences": "喜欢高尔夫",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 4.3 创建联系人
- **接口路径**：`/api/v1/crm/contacts`
- **请求方法**：POST
- **请求体**：
```json
{
  "accountId": "account-001",
  "name": "李四",
  "position": "采购经理",
  "phone": "13900139001",
  "email": "lisi@example.com",
  "mobile": "13900139001",
  "address": "北京市朝阳区"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "contact-001"
  }
}
```

### 4.4 更新联系人
- **接口路径**：`/api/v1/crm/contacts/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 联系人ID |
- **请求体**：
```json
{
  "name": "李四（更新）",
  "position": "采购总监",
  "phone": "13900139002",
  "email": "lisi-updated@example.com"
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

### 4.5 删除联系人
- **接口路径**：`/api/v1/crm/contacts/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 联系人ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 5. 客户活动API

### 5.1 获取活动列表
- **接口路径**：`/api/v1/crm/activities`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | accountId | string | 否 | 客户ID |
  | type | string | 否 | 活动类型（call, meeting, email, note） |
  | status | string | 否 | 状态（planned, completed, cancelled） |
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
        "id": "activity-001",
        "accountId": "account-001",
        "accountName": "北京科技有限公司",
        "contactId": "contact-001",
        "contactName": "李四",
        "type": "meeting",
        "subject": "产品演示会议",
        "description": "介绍新产品功能",
        "status": "completed",
        "startTime": "2023-06-01T10:00:00Z",
        "endTime": "2023-06-01T11:00:00Z",
        "createdBy": "admin",
        "createdAt": "2023-05-30T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-01T11:00:00Z"
      }
    ],
    "total": 30,
    "page": 1,
    "pageSize": 10
  }
}
```

### 5.2 获取活动详情
- **接口路径**：`/api/v1/crm/activities/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 活动ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "activity-001",
    "accountId": "account-001",
    "accountName": "北京科技有限公司",
    "contactId": "contact-001",
    "contactName": "李四",
    "type": "meeting",
    "subject": "产品演示会议",
    "description": "介绍新产品功能",
    "location": "公司会议室A",
    "status": "completed",
    "priority": "high",
    "startTime": "2023-06-01T10:00:00Z",
    "endTime": "2023-06-01T11:00:00Z",
    "outcome": "客户对产品很感兴趣，需要提供详细方案",
    "attachments": [
      {
        "id": "file-001",
        "name": "产品手册.pdf",
        "url": "/api/v1/files/file-001"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-05-30T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-01T11:00:00Z"
  }
}
```

### 5.3 创建活动
- **接口路径**：`/api/v1/crm/activities`
- **请求方法**：POST
- **请求体**：
```json
{
  "accountId": "account-001",
  "contactId": "contact-001",
  "type": "meeting",
  "subject": "产品演示会议",
  "description": "介绍新产品功能",
  "location": "公司会议室A",
  "priority": "high",
  "startTime": "2023-06-01T10:00:00Z",
  "endTime": "2023-06-01T11:00:00Z"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "activity-001"
  }
}
```

### 5.4 更新活动
- **接口路径**：`/api/v1/crm/activities/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 活动ID |
- **请求体**：
```json
{
  "subject": "产品演示会议（更新）",
  "location": "公司会议室B",
  "status": "completed",
  "outcome": "客户对产品很感兴趣，需要提供详细方案"
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

### 5.5 删除活动
- **接口路径**：`/api/v1/crm/activities/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 活动ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 6. 客户反馈API

### 6.1 获取反馈列表
- **接口路径**：`/api/v1/crm/cases`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | accountId | string | 否 | 客户ID |
  | caseNo | string | 否 | 反馈编号 |
  | subject | string | 否 | 主题 |
  | status | string | 否 | 状态（new, in_progress, resolved, closed） |
  | priority | string | 否 | 优先级（low, medium, high, urgent） |
  | type | string | 否 | 类型（question, complaint, suggestion, issue） |
  | createdDateStart | string | 否 | 创建开始日期，格式：YYYY-MM-DD |
  | createdDateEnd | string | 否 | 创建结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "case-001",
        "caseNo": "CAS001",
        "accountId": "account-001",
        "accountName": "北京科技有限公司",
        "contactId": "contact-001",
        "contactName": "李四",
        "subject": "系统登录问题",
        "type": "issue",
        "priority": "high",
        "status": "resolved",
        "createdBy": "admin",
        "createdAt": "2023-06-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-02T09:00:00Z"
      }
    ],
    "total": 25,
    "page": 1,
    "pageSize": 10
  }
}
```

### 6.2 获取反馈详情
- **接口路径**：`/api/v1/crm/cases/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 反馈ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "case-001",
    "caseNo": "CAS001",
    "accountId": "account-001",
    "accountName": "北京科技有限公司",
    "contactId": "contact-001",
    "contactName": "李四",
    "subject": "系统登录问题",
    "description": "用户无法登录系统，提示密码错误",
    "type": "issue",
    "priority": "high",
    "status": "resolved",
    "resolution": "重置用户密码，问题已解决",
    "responseTime": "2023-06-01T08:30:00Z",
    "resolutionTime": "2023-06-01T09:00:00Z",
    "attachments": [
      {
        "id": "file-002",
        "name": "错误截图.png",
        "url": "/api/v1/files/file-002"
      }
    ],
    "comments": [
      {
        "id": "comment-001",
        "content": "已收到反馈，正在处理",
        "createdBy": "admin",
        "createdAt": "2023-06-01T08:15:00Z"
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-06-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-02T09:00:00Z"
  }
}
```

### 6.3 创建反馈
- **接口路径**：`/api/v1/crm/cases`
- **请求方法**：POST
- **请求体**：
```json
{
  "accountId": "account-001",
  "contactId": "contact-001",
  "subject": "系统登录问题",
  "description": "用户无法登录系统，提示密码错误",
  "type": "issue",
  "priority": "high"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "case-001",
    "caseNo": "CAS001"
  }
}
```

### 6.4 更新反馈
- **接口路径**：`/api/v1/crm/cases/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 反馈ID |
- **请求体**：
```json
{
  "status": "resolved",
  "resolution": "重置用户密码，问题已解决"
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

### 6.5 添加反馈评论
- **接口路径**：`/api/v1/crm/cases/{id}/comments`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 反馈ID |
- **请求体**：
```json
{
  "content": "已收到反馈，正在处理"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "commentId": "comment-001"
  }
}
```

### 6.6 关闭反馈
- **接口路径**：`/api/v1/crm/cases/{id}/close`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 反馈ID |
- **请求体**：
```json
{
  "resolution": "问题已完全解决",
  "satisfaction": 5
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

## 7. 销售线索与机会API

### 7.1 获取销售线索列表
- **接口路径**：`/api/v1/crm/leads`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | name | string | 否 | 线索名称 |
  | status | string | 否 | 状态（new, qualified, unqualified, converted） |
  | source | string | 否 | 来源（website, referral, event, cold_call） |
  | createdDateStart | string | 否 | 创建开始日期，格式：YYYY-MM-DD |
  | createdDateEnd | string | 否 | 创建结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "lead-001",
        "name": "潜在客户A",
        "company": "科技公司",
        "contactPerson": "张三",
        "phone": "13800138001",
        "email": "lead@example.com",
        "source": "website",
        "status": "qualified",
        "score": 85,
        "createdBy": "admin",
        "createdAt": "2023-06-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-02T09:00:00Z"
      }
    ],
    "total": 40,
    "page": 1,
    "pageSize": 10
  }
}
```

### 7.2 创建销售线索
- **接口路径**：`/api/v1/crm/leads`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "潜在客户A",
  "company": "科技公司",
  "contactPerson": "张三",
  "phone": "13800138001",
  "email": "lead@example.com",
  "source": "website",
  "description": "通过网站表单提交的潜在客户"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "lead-001"
  }
}
```

### 7.3 转换销售线索为客户
- **接口路径**：`/api/v1/crm/leads/{id}/convert`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 线索ID |
- **请求体**：
```json
{
  "createAccount": true,
  "createContact": true,
  "opportunityName": "产品采购机会",
  "estimatedAmount": 1000000,
  "closeDate": "2023-07-31"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "accountId": "account-002",
    "contactId": "contact-002",
    "opportunityId": "opp-001"
  }
}
```

### 7.4 获取销售机会列表
- **接口路径**：`/api/v1/crm/opportunities`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | accountId | string | 否 | 客户ID |
  | opportunityNo | string | 否 | 机会编号 |
  | name | string | 否 | 机会名称 |
  | stage | string | 否 | 阶段（prospecting, qualification, proposal, negotiation, closed_won, closed_lost） |
  | probability | int | 否 | 成功率（0-100） |
  | estimatedAmountMin | number | 否 | 预估金额最小值 |
  | estimatedAmountMax | number | 否 | 预估金额最大值 |
  | closeDateStart | string | 否 | 预计关闭开始日期，格式：YYYY-MM-DD |
  | closeDateEnd | string | 否 | 预计关闭结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "opp-001",
        "opportunityNo": "OPP001",
        "accountId": "account-001",
        "accountName": "北京科技有限公司",
        "name": "产品采购机会",
        "description": "采购一批新产品",
        "stage": "proposal",
        "probability": 60,
        "estimatedAmount": 1000000,
        "actualAmount": null,
        "closeDate": "2023-07-31",
        "createdBy": "admin",
        "createdAt": "2023-06-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-15T09:00:00Z"
      }
    ],
    "total": 15,
    "page": 1,
    "pageSize": 10
  }
}
```

### 7.5 创建销售机会
- **接口路径**：`/api/v1/crm/opportunities`
- **请求方法**：POST
- **请求体**：
```json
{
  "accountId": "account-001",
  "name": "产品采购机会",
  "description": "采购一批新产品",
  "stage": "prospecting",
  "probability": 30,
  "estimatedAmount": 1000000,
  "closeDate": "2023-07-31"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "opp-001",
    "opportunityNo": "OPP001"
  }
}
```

## 8. 客户分析API

### 8.1 获取客户概览
- **接口路径**：`/api/v1/crm/analytics/overview`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | period | string | 否 | 时间周期（week, month, quarter, year） |
  | startDate | string | 否 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 否 | 结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "totalCustomers": 100,
    "newCustomers": 10,
    "activeCustomers": 80,
    "totalOpportunities": 50,
    "openOpportunities": 30,
    "wonOpportunities": 15,
    "lostOpportunities": 5,
    "totalRevenue": 5000000,
    "averageDealSize": 333333,
    "customerAcquisitionCost": 5000,
    "customerLifetimeValue": 50000
  }
}
```

### 8.2 获取客户趋势
- **接口路径**：`/api/v1/crm/analytics/trends`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | metric | string | 是 | 指标类型（customers, opportunities, revenue, activities） |
  | period | string | 否 | 时间周期（week, month, quarter, year） |
  | startDate | string | 否 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 否 | 结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "metric": "customers",
    "period": "month",
    "data": [
      {
        "date": "2023-01",
        "value": 80
      },
      {
        "date": "2023-02",
        "value": 85
      },
      {
        "date": "2023-03",
        "value": 90
      },
      {
        "date": "2023-04",
        "value": 95
      },
      {
        "date": "2023-05",
        "value": 98
      },
      {
        "date": "2023-06",
        "value": 100
      }
    ]
  }
}
```

### 8.3 获取客户细分
- **接口路径**：`/api/v1/crm/analytics/segmentation`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | segmentBy | string | 是 | 细分维度（industry, region, size, loyalty） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "segmentBy": "industry",
    "segments": [
      {
        "name": "科技",
        "count": 30,
        "revenue": 2000000,
        "percentage": 30
      },
      {
        "name": "制造",
        "count": 25,
        "revenue": 1500000,
        "percentage": 25
      },
      {
        "name": "零售",
        "count": 20,
        "revenue": 1000000,
        "percentage": 20
      },
      {
        "name": "其他",
        "count": 25,
        "revenue": 500000,
        "percentage": 25
      }
    ]
  }
}
```

### 8.4 获取客户忠诚度分析
- **接口路径**：`/api/v1/crm/analytics/loyalty`
- **请求方法**：GET
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "loyaltyLevels": [
      {
        "level": "platinum",
        "count": 5,
        "revenue": 2000000,
        "retentionRate": 95
      },
      {
        "level": "gold",
        "count": 15,
        "revenue": 1500000,
        "retentionRate": 90
      },
      {
        "level": "silver",
        "count": 30,
        "revenue": 1000000,
        "retentionRate": 80
      },
      {
        "level": "bronze",
        "count": 50,
        "revenue": 500000,
        "retentionRate": 60
      }
    ],
    "recommendation": "针对青铜级客户，建议实施客户关怀计划，提高客户忠诚度"
  }
}
```

## 9. 营销活动API

### 9.1 获取营销活动列表
- **接口路径**：`/api/v1/crm/campaigns`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | campaignNo | string | 否 | 活动编号 |
  | name | string | 否 | 活动名称 |
  | status | string | 否 | 状态（planning, active, completed, cancelled） |
  | type | string | 否 | 活动类型（email, social, event, advertisement） |
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
        "id": "campaign-001",
        "campaignNo": "CMP001",
        "name": "新产品发布会",
        "type": "event",
        "status": "active",
        "startDate": "2023-06-01",
        "endDate": "2023-06-30",
        "budget": 100000,
        "actualCost": 80000,
        "createdBy": "admin",
        "createdAt": "2023-05-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-01T09:00:00Z"
      }
    ],
    "total": 10,
    "page": 1,
    "pageSize": 10
  }
}
```

### 9.2 创建营销活动
- **接口路径**：`/api/v1/crm/campaigns`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "新产品发布会",
  "type": "event",
  "description": "发布新产品的线上线下活动",
  "startDate": "2023-06-01",
  "endDate": "2023-06-30",
  "budget": 100000,
  "targetAudience": "existing_customers, prospects"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "campaign-001",
    "campaignNo": "CMP001"
  }
}
```

### 9.3 获取营销活动效果
- **接口路径**：`/api/v1/crm/campaigns/{id}/analytics`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 活动ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "campaignId": "campaign-001",
    "campaignName": "新产品发布会",
    "metrics": {
      "impressions": 10000,
      "clicks": 1000,
      "conversions": 100,
      "conversionRate": 10,
      "costPerClick": 80,
      "returnOnInvestment": 200,
      "revenueGenerated": 200000
    },
    "leadsGenerated": 50,
    "opportunitiesCreated": 20,
    "salesClosed": 10
  }
}
```

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
| 4001 | 客户不存在 |
| 4002 | 联系人不存在 |
| 4003 | 活动不存在 |
| 4004 | 反馈不存在 |
| 4005 | 线索不存在 |
| 4006 | 机会不存在 |
| 4007 | 活动类型错误 |
| 4008 | 反馈状态错误 |
| 4009 | 线索状态错误 |
| 4010 | 机会阶段错误 |

## 11. 附录

### 11.1 参考文档
- 《CRM系统设计与实现》
- 《客户关系管理实务》
- 《API设计最佳实践》

### 11.2 接口测试
- 使用Postman或Swagger进行接口测试
- 测试用例应覆盖正常流程和异常流程
- 测试数据应符合业务规则和数据约束

### 11.3 接口安全
- 所有接口都应进行身份认证和授权
- 敏感数据应进行加密传输
- 接口应进行速率限制，防止恶意请求
- 应记录接口访问日志，便于审计和问题排查