# 人力资源模块API文档

## 1. 文档概述

### 1.1 文档目的
本文档定义了ERP系统人力资源模块的API接口规范，包括员工管理、部门管理、考勤管理、薪资管理、培训管理和招聘管理等功能的API接口设计。旨在为前端开发和后端开发提供明确的接口规范，确保系统集成的顺利进行。

### 1.2 术语定义
| 术语 | 解释 |
|------|------|
| HR | 人力资源（Human Resources），企业中负责员工管理的部门 |
| EMP | 员工（Employee），企业的在职人员 |
| DEPT | 部门（Department），企业内部的组织单位 |
| ATT | 考勤（Attendance），员工的出勤记录 |
| SAL | 薪资（Salary），员工的工资待遇 |
| TRAIN | 培训（Training），企业为员工提供的学习活动 |
| REC | 招聘（Recruitment），企业招募新员工的过程 |
| POS | 职位（Position），企业中的工作岗位 |
| PERF | 绩效（Performance），员工的工作表现评估 |

## 2. 通用规范

### 2.1 接口风格
- **URL格式**：`/api/v1/hr/{resource}`
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

## 3. 员工管理API

### 3.1 获取员工列表
- **接口路径**：`/api/v1/hr/employees`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | name | string | 否 | 员工姓名 |
  | departmentId | string | 否 | 部门ID |
  | positionId | string | 否 | 职位ID |
  | status | string | 否 | 状态（active, resigned, on_leave） |
  | entryDateStart | string | 否 | 入职开始日期，格式：YYYY-MM-DD |
  | entryDateEnd | string | 否 | 入职结束日期，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "emp-001",
        "employeeId": "EMP001",
        "name": "张三",
        "gender": "male",
        "birthDate": "1990-01-01",
        "phone": "13800138001",
        "email": "zhangsan@example.com",
        "departmentId": "dept-001",
        "departmentName": "技术部",
        "positionId": "pos-001",
        "positionName": "软件工程师",
        "entryDate": "2020-01-01",
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

### 3.2 获取员工详情
- **接口路径**：`/api/v1/hr/employees/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 员工ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "emp-001",
    "employeeId": "EMP001",
    "name": "张三",
    "gender": "male",
    "birthDate": "1990-01-01",
    "phone": "13800138001",
    "email": "zhangsan@example.com",
    "departmentId": "dept-001",
    "departmentName": "技术部",
    "positionId": "pos-001",
    "positionName": "软件工程师",
    "entryDate": "2020-01-01",
    "probationEndDate": "2020-03-31",
    "contractStartDate": "2020-01-01",
    "contractEndDate": "2023-12-31",
    "status": "active",
    "address": "北京市海淀区",
    "education": "本科",
    "major": "计算机科学与技术",
    "school": "北京大学",
    "idCard": "110101199001011234",
    "emergencyContact": "李四",
    "emergencyPhone": "13900139001",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 3.3 创建员工
- **接口路径**：`/api/v1/hr/employees`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "张三",
  "gender": "male",
  "birthDate": "1990-01-01",
  "phone": "13800138001",
  "email": "zhangsan@example.com",
  "departmentId": "dept-001",
  "positionId": "pos-001",
  "entryDate": "2020-01-01",
  "probationEndDate": "2020-03-31",
  "contractStartDate": "2020-01-01",
  "contractEndDate": "2023-12-31",
  "address": "北京市海淀区",
  "education": "本科",
  "major": "计算机科学与技术",
  "school": "北京大学",
  "idCard": "110101199001011234",
  "emergencyContact": "李四",
  "emergencyPhone": "13900139001"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "emp-001",
    "employeeId": "EMP001"
  }
}
```

### 3.4 更新员工
- **接口路径**：`/api/v1/hr/employees/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 员工ID |
- **请求体**：
```json
{
  "name": "张三",
  "phone": "13800138002",
  "email": "zhangsan@example.com",
  "departmentId": "dept-001",
  "positionId": "pos-002",
  "address": "北京市朝阳区"
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

### 3.5 删除员工
- **接口路径**：`/api/v1/hr/employees/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 员工ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 3.6 员工离职
- **接口路径**：`/api/v1/hr/employees/{id}/resign`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 员工ID |
- **请求体**：
```json
{
  "resignDate": "2023-06-30",
  "reason": "个人发展",
  "handoverPersonId": "emp-002"
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

### 3.7 员工复职
- **接口路径**：`/api/v1/hr/employees/{id}/reinstate`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 员工ID |
- **请求体**：
```json
{
  "reinstateDate": "2023-09-01",
  "departmentId": "dept-001",
  "positionId": "pos-001"
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

### 3.8 导入员工
- **接口路径**：`/api/v1/hr/employees/import`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | file | file | 是 | 员工数据Excel文件 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "successCount": 10,
    "failCount": 0
  }
}
```

### 3.9 导出员工
- **接口路径**：`/api/v1/hr/employees/export`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | departmentId | string | 否 | 部门ID |
  | status | string | 否 | 状态 |
- **响应格式**：Excel文件

## 4. 部门管理API

### 4.1 获取部门列表
- **接口路径**：`/api/v1/hr/departments`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | name | string | 否 | 部门名称 |
  | parentId | string | 否 | 父部门ID |
  | status | string | 否 | 状态（active, inactive） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "dept-001",
        "code": "DEPT001",
        "name": "技术部",
        "parentId": "",
        "parentName": "",
        "managerId": "emp-001",
        "managerName": "张三",
        "description": "负责公司产品开发",
        "status": "active",
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

### 4.2 获取部门树形结构
- **接口路径**：`/api/v1/hr/departments/tree`
- **请求方法**：GET
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": "dept-001",
      "code": "DEPT001",
      "name": "技术部",
      "parentId": "",
      "children": [
        {
          "id": "dept-002",
          "code": "DEPT002",
          "name": "前端开发组",
          "parentId": "dept-001",
          "children": []
        }
      ]
    }
  ]
}
```

### 4.3 获取部门详情
- **接口路径**：`/api/v1/hr/departments/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 部门ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "dept-001",
    "code": "DEPT001",
    "name": "技术部",
    "parentId": "",
    "parentName": "",
    "managerId": "emp-001",
    "managerName": "张三",
    "description": "负责公司产品开发",
    "status": "active",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 4.4 创建部门
- **接口路径**：`/api/v1/hr/departments`
- **请求方法**：POST
- **请求体**：
```json
{
  "code": "DEPT001",
  "name": "技术部",
  "parentId": "",
  "managerId": "emp-001",
  "description": "负责公司产品开发"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "dept-001"
  }
}
```

### 4.5 更新部门
- **接口路径**：`/api/v1/hr/departments/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 部门ID |
- **请求体**：
```json
{
  "name": "技术部",
  "managerId": "emp-002",
  "description": "负责公司产品开发和技术研究"
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

### 4.6 删除部门
- **接口路径**：`/api/v1/hr/departments/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 部门ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 5. 职位管理API

### 5.1 获取职位列表
- **接口路径**：`/api/v1/hr/positions`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | name | string | 否 | 职位名称 |
  | departmentId | string | 否 | 部门ID |
  | level | string | 否 | 职级 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "pos-001",
        "code": "POS001",
        "name": "软件工程师",
        "departmentId": "dept-001",
        "departmentName": "技术部",
        "level": "P5",
        "description": "负责软件系统的开发",
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

### 5.2 获取职位详情
- **接口路径**：`/api/v1/hr/positions/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 职位ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "pos-001",
    "code": "POS001",
    "name": "软件工程师",
    "departmentId": "dept-001",
    "departmentName": "技术部",
    "level": "P5",
    "description": "负责软件系统的开发",
    "responsibilities": "1. 参与系统需求分析和设计\n2. 负责代码实现和单元测试\n3. 参与系统部署和维护",
    "requirements": "1. 本科及以上学历，计算机相关专业\n2. 3年以上相关工作经验\n3. 熟悉Java或Python编程",
    "createdBy": "admin",
    "createdAt": "2020-01-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2020-01-01T08:00:00Z"
  }
}
```

### 5.3 创建职位
- **接口路径**：`/api/v1/hr/positions`
- **请求方法**：POST
- **请求体**：
```json
{
  "code": "POS001",
  "name": "软件工程师",
  "departmentId": "dept-001",
  "level": "P5",
  "description": "负责软件系统的开发",
  "responsibilities": "1. 参与系统需求分析和设计\n2. 负责代码实现和单元测试\n3. 参与系统部署和维护",
  "requirements": "1. 本科及以上学历，计算机相关专业\n2. 3年以上相关工作经验\n3. 熟悉Java或Python编程"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "pos-001"
  }
}
```

### 5.4 更新职位
- **接口路径**：`/api/v1/hr/positions/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 职位ID |
- **请求体**：
```json
{
  "name": "高级软件工程师",
  "level": "P6",
  "description": "负责软件系统的设计和开发"
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

### 5.5 删除职位
- **接口路径**：`/api/v1/hr/positions/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 职位ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

## 6. 考勤管理API

### 6.1 获取考勤记录列表
- **接口路径**：`/api/v1/hr/attendances`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | employeeId | string | 否 | 员工ID |
  | departmentId | string | 否 | 部门ID |
  | dateStart | string | 否 | 开始日期，格式：YYYY-MM-DD |
  | dateEnd | string | 否 | 结束日期，格式：YYYY-MM-DD |
  | type | string | 否 | 类型（normal, late, early_leave, absent, overtime） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "att-001",
        "employeeId": "emp-001",
        "employeeName": "张三",
        "departmentId": "dept-001",
        "departmentName": "技术部",
        "date": "2023-06-01",
        "checkInTime": "08:30:00",
        "checkOutTime": "17:30:00",
        "type": "normal",
        "status": "confirmed",
        "createdBy": "system",
        "createdAt": "2023-06-01T08:30:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-01T17:30:00Z"
      }
    ],
    "total": 30,
    "page": 1,
    "pageSize": 10
  }
}
```

### 6.2 获取考勤详情
- **接口路径**：`/api/v1/hr/attendances/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 考勤记录ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "att-001",
    "employeeId": "emp-001",
    "employeeName": "张三",
    "departmentId": "dept-001",
    "departmentName": "技术部",
    "date": "2023-06-01",
    "checkInTime": "08:30:00",
    "checkOutTime": "17:30:00",
    "type": "normal",
    "status": "confirmed",
    "remark": "正常出勤",
    "createdBy": "system",
    "createdAt": "2023-06-01T08:30:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-01T17:30:00Z"
  }
}
```

### 6.3 打卡
- **接口路径**：`/api/v1/hr/attendances/checkin`
- **请求方法**：POST
- **请求体**：
```json
{
  "employeeId": "emp-001",
  "type": "checkin", // checkin 或 checkout
  "latitude": 39.9042,
  "longitude": 116.4074,
  "address": "北京市海淀区中关村"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "att-001",
    "time": "08:30:00"
  }
}
```

### 6.4 手动录入考勤
- **接口路径**：`/api/v1/hr/attendances`
- **请求方法**：POST
- **请求体**：
```json
{
  "employeeId": "emp-001",
  "date": "2023-06-01",
  "checkInTime": "08:30:00",
  "checkOutTime": "17:30:00",
  "type": "normal",
  "remark": "正常出勤"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "att-001"
  }
}
```

### 6.5 更新考勤
- **接口路径**：`/api/v1/hr/attendances/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 考勤记录ID |
- **请求体**：
```json
{
  "checkInTime": "08:45:00",
  "type": "late",
  "remark": "迟到15分钟"
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

### 6.6 删除考勤
- **接口路径**：`/api/v1/hr/attendances/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 考勤记录ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 6.7 获取考勤统计
- **接口路径**：`/api/v1/hr/attendances/stats`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | employeeId | string | 否 | 员工ID |
  | departmentId | string | 否 | 部门ID |
  | month | string | 是 | 月份，格式：YYYY-MM |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "normalDays": 20,
    "lateDays": 2,
    "earlyLeaveDays": 1,
    "absentDays": 0,
    "overtimeHours": 10,
    "leaveDays": 2
  }
}
```

### 6.8 导入考勤
- **接口路径**：`/api/v1/hr/attendances/import`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | file | file | 是 | 考勤数据Excel文件 |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "successCount": 100,
    "failCount": 0
  }
}
```

### 6.9 导出考勤
- **接口路径**：`/api/v1/hr/attendances/export`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | departmentId | string | 否 | 部门ID |
  | dateStart | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | dateEnd | string | 是 | 结束日期，格式：YYYY-MM-DD |
- **响应格式**：Excel文件

## 7. 请假管理API

### 7.1 获取请假记录列表
- **接口路径**：`/api/v1/hr/leaves`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | employeeId | string | 否 | 员工ID |
  | departmentId | string | 否 | 部门ID |
  | type | string | 否 | 类型（annual, sick, personal, maternity, paternity, other） |
  | status | string | 否 | 状态（pending, approved, rejected） |
  | startDateStart | string | 否 | 开始日期范围开始，格式：YYYY-MM-DD |
  | startDateEnd | string | 否 | 开始日期范围结束，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "leave-001",
        "leaveNo": "LEAVE2023060001",
        "employeeId": "emp-001",
        "employeeName": "张三",
        "departmentId": "dept-001",
        "departmentName": "技术部",
        "type": "annual",
        "startDate": "2023-06-01",
        "endDate": "2023-06-02",
        "days": 2,
        "reason": "个人休息",
        "status": "approved",
        "approverId": "emp-003",
        "approverName": "王五",
        "createdBy": "emp-001",
        "createdAt": "2023-05-20T10:00:00Z",
        "updatedBy": "emp-003",
        "updatedAt": "2023-05-21T10:00:00Z"
      }
    ],
    "total": 20,
    "page": 1,
    "pageSize": 10
  }
}
```

### 7.2 获取请假详情
- **接口路径**：`/api/v1/hr/leaves/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 请假记录ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "leave-001",
    "leaveNo": "LEAVE2023060001",
    "employeeId": "emp-001",
    "employeeName": "张三",
    "departmentId": "dept-001",
    "departmentName": "技术部",
    "type": "annual",
    "startDate": "2023-06-01",
    "endDate": "2023-06-02",
    "days": 2,
    "reason": "个人休息",
    "status": "approved",
    "approverId": "emp-003",
    "approverName": "王五",
    "approvalRemark": "同意",
    "createdBy": "emp-001",
    "createdAt": "2023-05-20T10:00:00Z",
    "updatedBy": "emp-003",
    "updatedAt": "2023-05-21T10:00:00Z"
  }
}
```

### 7.3 申请请假
- **接口路径**：`/api/v1/hr/leaves`
- **请求方法**：POST
- **请求体**：
```json
{
  "type": "annual",
  "startDate": "2023-06-01",
  "endDate": "2023-06-02",
  "days": 2,
  "reason": "个人休息",
  "approverId": "emp-003"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "leave-001",
    "leaveNo": "LEAVE2023060001"
  }
}
```

### 7.4 更新请假
- **接口路径**：`/api/v1/hr/leaves/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 请假记录ID |
- **请求体**：
```json
{
  "endDate": "2023-06-03",
  "days": 3,
  "reason": "个人休息和处理私事"
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

### 7.5 删除请假
- **接口路径**：`/api/v1/hr/leaves/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 请假记录ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 7.6 审批请假
- **接口路径**：`/api/v1/hr/leaves/{id}/approve`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 请假记录ID |
- **请求体**：
```json
{
  "approved": true,
  "remark": "同意"
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

### 7.7 获取请假余额
- **接口路径**：`/api/v1/hr/leaves/balance/{employeeId}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | employeeId | string | 是 | 员工ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "annual": {
      "total": 15,
      "used": 2,
      "remaining": 13
    },
    "sick": {
      "total": 10,
      "used": 0,
      "remaining": 10
    }
  }
}
```

## 8. 加班管理API

### 8.1 获取加班记录列表
- **接口路径**：`/api/v1/hr/overtimes`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | employeeId | string | 否 | 员工ID |
  | departmentId | string | 否 | 部门ID |
  | dateStart | string | 否 | 开始日期，格式：YYYY-MM-DD |
  | dateEnd | string | 否 | 结束日期，格式：YYYY-MM-DD |
  | status | string | 否 | 状态（pending, approved, rejected） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "ot-001",
        "overtimeNo": "OT2023060001",
        "employeeId": "emp-001",
        "employeeName": "张三",
        "departmentId": "dept-001",
        "departmentName": "技术部",
        "date": "2023-06-01",
        "startTime": "18:00:00",
        "endTime": "21:00:00",
        "hours": 3,
        "reason": "项目赶工",
        "type": "workday", // workday, weekend, holiday
        "status": "approved",
        "approverId": "emp-003",
        "approverName": "王五",
        "createdBy": "emp-001",
        "createdAt": "2023-06-01T21:00:00Z",
        "updatedBy": "emp-003",
        "updatedAt": "2023-06-02T10:00:00Z"
      }
    ],
    "total": 15,
    "page": 1,
    "pageSize": 10
  }
}
```

### 8.2 获取加班详情
- **接口路径**：`/api/v1/hr/overtimes/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 加班记录ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "ot-001",
    "overtimeNo": "OT2023060001",
    "employeeId": "emp-001",
    "employeeName": "张三",
    "departmentId": "dept-001",
    "departmentName": "技术部",
    "date": "2023-06-01",
    "startTime": "18:00:00",
    "endTime": "21:00:00",
    "hours": 3,
    "reason": "项目赶工",
    "type": "workday",
    "status": "approved",
    "approverId": "emp-003",
    "approverName": "王五",
    "approvalRemark": "同意",
    "createdBy": "emp-001",
    "createdAt": "2023-06-01T21:00:00Z",
    "updatedBy": "emp-003",
    "updatedAt": "2023-06-02T10:00:00Z"
  }
}
```

### 8.3 申请加班
- **接口路径**：`/api/v1/hr/overtimes`
- **请求方法**：POST
- **请求体**：
```json
{
  "date": "2023-06-01",
  "startTime": "18:00:00",
  "endTime": "21:00:00",
  "hours": 3,
  "reason": "项目赶工",
  "type": "workday",
  "approverId": "emp-003"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "ot-001",
    "overtimeNo": "OT2023060001"
  }
}
```

### 8.4 更新加班
- **接口路径**：`/api/v1/hr/overtimes/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 加班记录ID |
- **请求体**：
```json
{
  "endTime": "22:00:00",
  "hours": 4,
  "reason": "项目赶工，需要完成更多工作"
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

### 8.5 删除加班
- **接口路径**：`/api/v1/hr/overtimes/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 加班记录ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 8.6 审批加班
- **接口路径**：`/api/v1/hr/overtimes/{id}/approve`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 加班记录ID |
- **请求体**：
```json
{
  "approved": true,
  "remark": "同意"
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

### 8.7 获取加班统计
- **接口路径**：`/api/v1/hr/overtimes/stats`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | employeeId | string | 否 | 员工ID |
  | departmentId | string | 否 | 部门ID |
  | month | string | 是 | 月份，格式：YYYY-MM |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "workdayHours": 10,
    "weekendHours": 8,
    "holidayHours": 4,
    "totalHours": 22
  }
}
```

## 9. 薪资管理API

### 9.1 获取薪资列表
- **接口路径**：`/api/v1/hr/salaries`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | employeeId | string | 否 | 员工ID |
  | departmentId | string | 否 | 部门ID |
  | month | string | 否 | 月份，格式：YYYY-MM |
  | status | string | 否 | 状态（draft, confirmed, paid） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "sal-001",
        "salaryNo": "SAL2023060001",
        "employeeId": "emp-001",
        "employeeName": "张三",
        "departmentId": "dept-001",
        "departmentName": "技术部",
        "month": "2023-06",
        "basicSalary": 10000,
        "allowance": 2000,
        "overtimePay": 1000,
        "bonus": 3000,
        "deduction": 1500,
        "tax": 500,
        "netSalary": 14000,
        "status": "paid",
        "createdBy": "admin",
        "createdAt": "2023-06-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-15T10:00:00Z"
      }
    ],
    "total": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

### 9.2 获取薪资详情
- **接口路径**：`/api/v1/hr/salaries/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 薪资记录ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "sal-001",
    "salaryNo": "SAL2023060001",
    "employeeId": "emp-001",
    "employeeName": "张三",
    "departmentId": "dept-001",
    "departmentName": "技术部",
    "month": "2023-06",
    "basicSalary": 10000,
    "allowance": 2000,
    "overtimePay": 1000,
    "bonus": 3000,
    "deduction": 1500,
    "tax": 500,
    "netSalary": 14000,
    "details": {
      "allowances": [
        { "name": "交通补贴", "amount": 1000 },
        { "name": "餐补", "amount": 1000 }
      ],
      "deductions": [
        { "name": "社保", "amount": 1000 },
        { "name": "公积金", "amount": 500 }
      ],
      "overtimes": [
        { "date": "2023-06-01", "hours": 3, "amount": 500 },
        { "date": "2023-06-02", "hours": 2, "amount": 500 }
      ]
    },
    "status": "paid",
    "paidDate": "2023-06-15",
    "createdBy": "admin",
    "createdAt": "2023-06-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-15T10:00:00Z"
  }
}
```

### 9.3 创建薪资
- **接口路径**：`/api/v1/hr/salaries`
- **请求方法**：POST
- **请求体**：
```json
{
  "employeeId": "emp-001",
  "month": "2023-06",
  "basicSalary": 10000,
  "allowance": 2000,
  "overtimePay": 1000,
  "bonus": 3000,
  "deduction": 1500,
  "tax": 500,
  "netSalary": 14000,
  "details": {
    "allowances": [
      { "name": "交通补贴", "amount": 1000 },
      { "name": "餐补", "amount": 1000 }
    ],
    "deductions": [
      { "name": "社保", "amount": 1000 },
      { "name": "公积金", "amount": 500 }
    ]
  }
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "sal-001",
    "salaryNo": "SAL2023060001"
  }
}
```

### 9.4 更新薪资
- **接口路径**：`/api/v1/hr/salaries/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 薪资记录ID |
- **请求体**：
```json
{
  "bonus": 4000,
  "netSalary": 15000,
  "details": {
    "allowances": [
      { "name": "交通补贴", "amount": 1000 },
      { "name": "餐补", "amount": 1000 }
    ],
    "deductions": [
      { "name": "社保", "amount": 1000 },
      { "name": "公积金", "amount": 500 }
    ]
  }
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

### 9.5 删除薪资
- **接口路径**：`/api/v1/hr/salaries/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 薪资记录ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 9.6 确认薪资
- **接口路径**：`/api/v1/hr/salaries/{id}/confirm`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 薪资记录ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 9.7 标记薪资为已支付
- **接口路径**：`/api/v1/hr/salaries/{id}/pay`
- **请求方法**：POST
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 薪资记录ID |
- **请求体**：
```json
{
  "paidDate": "2023-06-15"
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

### 9.8 批量生成薪资
- **接口路径**：`/api/v1/hr/salaries/generate`
- **请求方法**：POST
- **请求体**：
```json
{
  "month": "2023-06",
  "departmentIds": ["dept-001", "dept-002"]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "generatedCount": 50
  }
}
```

### 9.9 导出薪资单
- **接口路径**：`/api/v1/hr/salaries/export`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | month | string | 是 | 月份，格式：YYYY-MM |
  | departmentId | string | 否 | 部门ID |
- **响应格式**：Excel文件

## 10. 培训管理API

### 10.1 获取培训列表
- **接口路径**：`/api/v1/hr/trainings`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | name | string | 否 | 培训名称 |
  | type | string | 否 | 培训类型（internal, external） |
  | status | string | 否 | 状态（planning, ongoing, completed, cancelled） |
  | startDateStart | string | 否 | 开始日期范围开始，格式：YYYY-MM-DD |
  | startDateEnd | string | 否 | 开始日期范围结束，格式：YYYY-MM-DD |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "train-001",
        "trainingNo": "TRAIN2023060001",
        "name": "新员工入职培训",
        "type": "internal",
        "startDate": "2023-06-01",
        "endDate": "2023-06-02",
        "location": "公司会议室A",
        "trainer": "李四",
        "status": "completed",
        "createdBy": "admin",
        "createdAt": "2023-05-01T08:00:00Z",
        "updatedBy": "admin",
        "updatedAt": "2023-06-02T17:00:00Z"
      }
    ],
    "total": 10,
    "page": 1,
    "pageSize": 10
  }
}
```

### 10.2 获取培训详情
- **接口路径**：`/api/v1/hr/trainings/{id}`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 培训ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "train-001",
    "trainingNo": "TRAIN2023060001",
    "name": "新员工入职培训",
    "type": "internal",
    "startDate": "2023-06-01",
    "endDate": "2023-06-02",
    "location": "公司会议室A",
    "trainer": "李四",
    "description": "新员工入职必备知识培训",
    "content": "1. 公司简介\n2. 企业文化\n3. 规章制度\n4. 安全培训",
    "status": "completed",
    "attendees": [
      {
        "employeeId": "emp-001",
        "employeeName": "张三",
        "departmentName": "技术部",
        "attendance": "present",
        "score": 90
      }
    ],
    "createdBy": "admin",
    "createdAt": "2023-05-01T08:00:00Z",
    "updatedBy": "admin",
    "updatedAt": "2023-06-02T17:00:00Z"
  }
}
```

### 10.3 创建培训
- **接口路径**：`/api/v1/hr/trainings`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "新员工入职培训",
  "type": "internal",
  "startDate": "2023-06-01",
  "endDate": "2023-06-02",
  "location": "公司会议室A",
  "trainer": "李四",
  "description": "新员工入职必备知识培训",
  "content": "1. 公司简介\n2. 企业文化\n3. 规章制度\n4. 安全培训",
  "attendeeIds": ["emp-001", "emp-002"]
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "train-001",
    "trainingNo": "TRAIN2023060001"
  }
}
```

### 10.4 更新培训
- **接口路径**：`/api/v1/hr/trainings/{id}`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 培训ID |
- **请求体**：
```json
{
  "name": "新员工入职培训（更新）",
  "location": "公司会议室B",
  "content": "1. 公司简介\n2. 企业文化\n3. 规章制度\n4. 安全培训\n5. 职业发展"
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

### 10.5 删除培训
- **接口路径**：`/api/v1/hr/trainings/{id}`
- **请求方法**：DELETE
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 培训ID |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

### 10.6 导出培训记录
- **接口路径**：`/api/v1/hr/trainings/export`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | startDate | string | 是 | 开始日期，格式：YYYY-MM-DD |
  | endDate | string | 是 | 结束日期，格式：YYYY-MM-DD |
- **响应格式**：Excel文件

## 11. 招聘管理API

### 11.1 获取招聘需求列表
- **接口路径**：`/api/v1/hr/recruit/requirements`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | positionId | string | 否 | 职位ID |
  | departmentId | string | 否 | 部门ID |
  | status | string | 否 | 状态（pending, approved, cancelled, completed） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "req-001",
        "requirementNo": "REQ2023060001",
        "positionId": "pos-001",
        "positionName": "软件工程师",
        "departmentId": "dept-001",
        "departmentName": "技术部",
        "headcount": 2,
        "status": "approved",
        "createdBy": "emp-003",
        "createdAt": "2023-05-01T10:00:00Z",
        "updatedBy": "emp-004",
        "updatedAt": "2023-05-02T10:00:00Z"
      }
    ],
    "total": 10,
    "page": 1,
    "pageSize": 10
  }
}
```

### 11.2 创建招聘需求
- **接口路径**：`/api/v1/hr/recruit/requirements`
- **请求方法**：POST
- **请求体**：
```json
{
  "positionId": "pos-001",
  "departmentId": "dept-001",
  "headcount": 2,
  "reason": "业务扩展需要",
  "requirement": "1. 本科及以上学历，计算机相关专业\n2. 3年以上相关工作经验\n3. 熟悉Java或Python编程",
  "expectedDate": "2023-06-30"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "req-001",
    "requirementNo": "REQ2023060001"
  }
}
```

### 11.3 获取候选人列表
- **接口路径**：`/api/v1/hr/recruit/candidates`
- **请求方法**：GET
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | page | int | 否 | 页码，默认1 |
  | pageSize | int | 否 | 每页条数，默认10 |
  | name | string | 否 | 候选人姓名 |
  | positionId | string | 否 | 应聘职位ID |
  | status | string | 否 | 状态（applied, screening, interview, offer, hired, rejected） |
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "cand-001",
        "candidateNo": "CAND2023060001",
        "name": "李四",
        "phone": "13900139001",
        "email": "lisi@example.com",
        "positionId": "pos-001",
        "positionName": "软件工程师",
        "departmentId": "dept-001",
        "departmentName": "技术部",
        "status": "interview",
        "createdAt": "2023-06-01T10:00:00Z"
      }
    ],
    "total": 20,
    "page": 1,
    "pageSize": 10
  }
}
```

### 11.4 创建候选人
- **接口路径**：`/api/v1/hr/recruit/candidates`
- **请求方法**：POST
- **请求体**：
```json
{
  "name": "李四",
  "phone": "13900139001",
  "email": "lisi@example.com",
  "positionId": "pos-001",
  "departmentId": "dept-001",
  "resume": "https://example.com/resume.pdf",
  "source": "linkedin",
  "remarks": "有5年相关工作经验"
}
```
- **响应格式**：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": "cand-001",
    "candidateNo": "CAND2023060001"
  }
}
```

### 11.5 更新候选人状态
- **接口路径**：`/api/v1/hr/recruit/candidates/{id}/status`
- **请求方法**：PUT
- **请求参数**：
  | 参数名 | 类型 | 必填 | 描述 |
  |--------|------|------|------|
  | id | string | 是 | 候选人ID |
- **请求体**：
```json
{
  "status": "offer",
  "remarks": "面试通过，发送offer"
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

## 12. 错误码定义

| 错误码 | 描述 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |
| 501 | 接口未实现 |

## 13. 附录

### 13.1 参考文档
- 《ERP系统人力资源管理模块设计与实现》
- 《人力资源管理实务》
- 《API设计最佳实践》

### 13.2 接口测试
- 使用Postman或Swagger进行接口测试
- 测试用例应覆盖正常流程和异常流程
- 测试数据应符合业务规则和数据约束
