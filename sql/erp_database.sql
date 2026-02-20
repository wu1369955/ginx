-- Active: 1771539750951@@127.0.0.1@3306@ginx
-- ERP系统数据库建表语句
-- 创建时间：2026-02-19

-- 设置字符集
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- 1. CRM模块

-- 1.1 客户表（crm_accounts）
CREATE TABLE IF NOT EXISTS `crm_accounts` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '客户ID',
  `account_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '客户编号',
  `name` VARCHAR(100) NOT NULL COMMENT '客户名称',
  `type` VARCHAR(20) NOT NULL COMMENT '客户类型（corporate, individual, government）',
  `industry` VARCHAR(50) COMMENT '行业',
  `region` VARCHAR(50) COMMENT '地区',
  `address` VARCHAR(255) COMMENT '地址',
  `contact_person` VARCHAR(50) COMMENT '联系人姓名',
  `phone` VARCHAR(20) COMMENT '电话',
  `email` VARCHAR(100) COMMENT '邮箱',
  `website` VARCHAR(255) COMMENT '网站',
  `tax_no` VARCHAR(30) COMMENT '税务登记号',
  `credit_limit` DECIMAL(18,2) DEFAULT 0 COMMENT '信用额度',
  `credit_used` DECIMAL(18,2) DEFAULT 0 COMMENT '已使用信用额度',
  `loyalty_level` VARCHAR(20) DEFAULT 'bronze' COMMENT '忠诚度等级',
  `remarks` TEXT COMMENT '备注',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='客户表';

-- 1.2 联系人表（crm_contacts）
CREATE TABLE IF NOT EXISTS `crm_contacts` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '联系人ID',
  `account_id` VARCHAR(36) NOT NULL COMMENT '客户ID',
  `name` VARCHAR(50) NOT NULL COMMENT '联系人姓名',
  `position` VARCHAR(50) COMMENT '职位',
  `phone` VARCHAR(20) COMMENT '电话',
  `email` VARCHAR(100) COMMENT '邮箱',
  `mobile` VARCHAR(20) COMMENT '手机',
  `fax` VARCHAR(20) COMMENT '传真',
  `address` VARCHAR(255) COMMENT '地址',
  `birthday` DATE COMMENT '生日',
  `preferences` TEXT COMMENT '偏好',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`account_id`) REFERENCES `crm_accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='联系人表';

-- 1.3 客户活动表（crm_activities）
CREATE TABLE IF NOT EXISTS `crm_activities` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '活动ID',
  `account_id` VARCHAR(36) NOT NULL COMMENT '客户ID',
  `contact_id` VARCHAR(36) COMMENT '联系人ID',
  `type` VARCHAR(20) NOT NULL COMMENT '活动类型（call, meeting, email, note）',
  `subject` VARCHAR(200) NOT NULL COMMENT '活动主题',
  `description` TEXT COMMENT '活动描述',
  `location` VARCHAR(255) COMMENT '活动地点',
  `status` VARCHAR(20) DEFAULT 'planned' COMMENT '状态（planned, completed, cancelled）',
  `priority` VARCHAR(20) DEFAULT 'medium' COMMENT '优先级（low, medium, high, urgent）',
  `start_time` DATETIME NOT NULL COMMENT '开始时间',
  `end_time` DATETIME NOT NULL COMMENT '结束时间',
  `outcome` TEXT COMMENT '活动结果',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`account_id`) REFERENCES `crm_accounts` (`id`),
  FOREIGN KEY (`contact_id`) REFERENCES `crm_contacts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='客户活动表';

-- 1.4 客户反馈表（crm_cases）
CREATE TABLE IF NOT EXISTS `crm_cases` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '反馈ID',
  `case_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '反馈编号',
  `account_id` VARCHAR(36) NOT NULL COMMENT '客户ID',
  `contact_id` VARCHAR(36) COMMENT '联系人ID',
  `subject` VARCHAR(200) NOT NULL COMMENT '主题',
  `description` TEXT NOT NULL COMMENT '描述',
  `type` VARCHAR(20) NOT NULL COMMENT '类型（question, complaint, suggestion, issue）',
  `priority` VARCHAR(20) DEFAULT 'medium' COMMENT '优先级（low, medium, high, urgent）',
  `status` VARCHAR(20) DEFAULT 'new' COMMENT '状态（new, in_progress, resolved, closed）',
  `resolution` TEXT COMMENT '解决方案',
  `response_time` DATETIME COMMENT '响应时间',
  `resolution_time` DATETIME COMMENT '解决时间',
  `satisfaction` TINYINT(1) COMMENT '满意度（1-5）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`account_id`) REFERENCES `crm_accounts` (`id`),
  FOREIGN KEY (`contact_id`) REFERENCES `crm_contacts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='客户反馈表';

-- 1.5 客户反馈评论表（crm_case_comments）
CREATE TABLE IF NOT EXISTS `crm_case_comments` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '评论ID',
  `case_id` VARCHAR(36) NOT NULL COMMENT '反馈ID',
  `content` TEXT NOT NULL COMMENT '评论内容',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  FOREIGN KEY (`case_id`) REFERENCES `crm_cases` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='客户反馈评论表';

-- 1.6 销售线索表（crm_leads）
CREATE TABLE IF NOT EXISTS `crm_leads` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '线索ID',
  `name` VARCHAR(100) NOT NULL COMMENT '线索名称',
  `company` VARCHAR(100) COMMENT '公司名称',
  `contact_person` VARCHAR(50) COMMENT '联系人姓名',
  `phone` VARCHAR(20) COMMENT '电话',
  `email` VARCHAR(100) COMMENT '邮箱',
  `source` VARCHAR(50) COMMENT '来源（website, referral, event, cold_call）',
  `status` VARCHAR(20) DEFAULT 'new' COMMENT '状态（new, qualified, unqualified, converted）',
  `score` TINYINT(3) DEFAULT 0 COMMENT '线索评分',
  `description` TEXT COMMENT '描述',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售线索表';

-- 1.7 销售机会表（crm_opportunities）
CREATE TABLE IF NOT EXISTS `crm_opportunities` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '机会ID',
  `opportunity_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '机会编号',
  `account_id` VARCHAR(36) NOT NULL COMMENT '客户ID',
  `name` VARCHAR(200) NOT NULL COMMENT '机会名称',
  `description` TEXT COMMENT '描述',
  `stage` VARCHAR(50) DEFAULT 'prospecting' COMMENT '阶段（prospecting, qualification, proposal, negotiation, closed_won, closed_lost）',
  `probability` TINYINT(3) DEFAULT 0 COMMENT '成功率（0-100）',
  `estimated_amount` DECIMAL(18,2) DEFAULT 0 COMMENT '预估金额',
  `actual_amount` DECIMAL(18,2) COMMENT '实际金额',
  `close_date` DATE NOT NULL COMMENT '预计关闭日期',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`account_id`) REFERENCES `crm_accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售机会表';

-- 1.8 营销活动表（crm_campaigns）
CREATE TABLE IF NOT EXISTS `crm_campaigns` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '活动ID',
  `campaign_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '活动编号',
  `name` VARCHAR(200) NOT NULL COMMENT '活动名称',
  `type` VARCHAR(50) NOT NULL COMMENT '活动类型（email, social, event, advertisement）',
  `description` TEXT COMMENT '描述',
  `status` VARCHAR(20) DEFAULT 'planning' COMMENT '状态（planning, active, completed, cancelled）',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `budget` DECIMAL(18,2) DEFAULT 0 COMMENT '预算',
  `actual_cost` DECIMAL(18,2) DEFAULT 0 COMMENT '实际成本',
  `target_audience` VARCHAR(255) COMMENT '目标受众',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='营销活动表';

-- 2. 销售模块

-- 2.1 客户表（sales_customers）
CREATE TABLE IF NOT EXISTS `sales_customers` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '客户ID',
  `customer_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '客户编号',
  `name` VARCHAR(100) NOT NULL COMMENT '客户名称',
  `contact_person` VARCHAR(50) COMMENT '联系人',
  `phone` VARCHAR(20) COMMENT '电话',
  `email` VARCHAR(100) COMMENT '邮箱',
  `address` VARCHAR(255) COMMENT '地址',
  `tax_no` VARCHAR(30) COMMENT '税号',
  `bank_name` VARCHAR(100) COMMENT '银行名称',
  `bank_account` VARCHAR(50) COMMENT '银行账号',
  `category` VARCHAR(50) COMMENT '客户类别',
  `region` VARCHAR(50) COMMENT '地区',
  `credit_limit` DECIMAL(18,2) DEFAULT 0 COMMENT '信用额度',
  `credit_days` INT DEFAULT 0 COMMENT '信用天数',
  `remarks` TEXT COMMENT '备注',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='客户表';

-- 2.2 产品表（sales_products）
CREATE TABLE IF NOT EXISTS `sales_products` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '产品ID',
  `product_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '产品编号',
  `name` VARCHAR(100) NOT NULL COMMENT '产品名称',
  `description` TEXT COMMENT '产品描述',
  `category_id` VARCHAR(36) COMMENT '产品类别ID',
  `unit` VARCHAR(10) NOT NULL COMMENT '单位',
  `price` DECIMAL(18,2) NOT NULL COMMENT '销售价格',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品表';

-- 2.3 产品类别表（sales_product_categories）
CREATE TABLE IF NOT EXISTS `sales_product_categories` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '类别ID',
  `code` VARCHAR(20) UNIQUE NOT NULL COMMENT '类别编码',
  `name` VARCHAR(100) NOT NULL COMMENT '类别名称',
  `parent_id` VARCHAR(36) COMMENT '父类别ID',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`parent_id`) REFERENCES `sales_product_categories` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品类别表';

-- 2.4 销售报价单表（sales_quotes）
CREATE TABLE IF NOT EXISTS `sales_quotes` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '报价单ID',
  `quote_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '报价单编号',
  `customer_id` VARCHAR(36) NOT NULL COMMENT '客户ID',
  `quote_date` DATE NOT NULL COMMENT '报价日期',
  `valid_until` DATE NOT NULL COMMENT '有效期至',
  `total_amount` DECIMAL(18,2) NOT NULL COMMENT '总金额',
  `status` VARCHAR(20) DEFAULT 'draft' COMMENT '状态（draft, submitted, approved, expired, cancelled）',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`customer_id`) REFERENCES `sales_customers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售报价单表';

-- 2.5 销售报价单明细表（sales_quote_items）
CREATE TABLE IF NOT EXISTS `sales_quote_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `quote_id` VARCHAR(36) NOT NULL COMMENT '报价单ID',
  `product_id` VARCHAR(36) NOT NULL COMMENT '产品ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `unit_price` DECIMAL(18,2) NOT NULL COMMENT '单价',
  `discount` DECIMAL(18,2) DEFAULT 0 COMMENT '折扣',
  `amount` DECIMAL(18,2) NOT NULL COMMENT '金额',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`quote_id`) REFERENCES `sales_quotes` (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `sales_products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售报价单明细表';

-- 2.6 销售订单表（sales_orders）
CREATE TABLE IF NOT EXISTS `sales_orders` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '订单ID',
  `order_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '订单编号',
  `customer_id` VARCHAR(36) NOT NULL COMMENT '客户ID',
  `order_date` DATE NOT NULL COMMENT '订单日期',
  `delivery_date` DATE COMMENT '交货日期',
  `total_amount` DECIMAL(18,2) NOT NULL COMMENT '总金额',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, confirmed, shipped, invoiced, cancelled）',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`customer_id`) REFERENCES `sales_customers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售订单表';

-- 2.7 销售订单明细表（sales_order_items）
CREATE TABLE IF NOT EXISTS `sales_order_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `order_id` VARCHAR(36) NOT NULL COMMENT '订单ID',
  `product_id` VARCHAR(36) NOT NULL COMMENT '产品ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `unit_price` DECIMAL(18,2) NOT NULL COMMENT '单价',
  `discount` DECIMAL(18,2) DEFAULT 0 COMMENT '折扣',
  `amount` DECIMAL(18,2) NOT NULL COMMENT '金额',
  `shipped_quantity` DECIMAL(18,4) DEFAULT 0 COMMENT '已发货数量',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`order_id`) REFERENCES `sales_orders` (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `sales_products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售订单明细表';

-- 2.8 销售发货单表（sales_deliveries）
CREATE TABLE IF NOT EXISTS `sales_deliveries` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '发货单ID',
  `delivery_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '发货单编号',
  `order_id` VARCHAR(36) NOT NULL COMMENT '订单ID',
  `customer_id` VARCHAR(36) NOT NULL COMMENT '客户ID',
  `delivery_date` DATE NOT NULL COMMENT '发货日期',
  `total_quantity` DECIMAL(18,4) NOT NULL COMMENT '总数量',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, shipped, cancelled）',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`order_id`) REFERENCES `sales_orders` (`id`),
  FOREIGN KEY (`customer_id`) REFERENCES `sales_customers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售发货单表';

-- 2.9 销售发货单明细表（sales_delivery_items）
CREATE TABLE IF NOT EXISTS `sales_delivery_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `delivery_id` VARCHAR(36) NOT NULL COMMENT '发货单ID',
  `order_item_id` VARCHAR(36) NOT NULL COMMENT '订单明细ID',
  `product_id` VARCHAR(36) NOT NULL COMMENT '产品ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '发货数量',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`delivery_id`) REFERENCES `sales_deliveries` (`id`),
  FOREIGN KEY (`order_item_id`) REFERENCES `sales_order_items` (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `sales_products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售发货单明细表';

-- 2.10 销售发票表（sales_invoices）
CREATE TABLE IF NOT EXISTS `sales_invoices` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '发票ID',
  `invoice_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '发票编号',
  `order_id` VARCHAR(36) NOT NULL COMMENT '订单ID',
  `customer_id` VARCHAR(36) NOT NULL COMMENT '客户ID',
  `invoice_date` DATE NOT NULL COMMENT '开票日期',
  `due_date` DATE NOT NULL COMMENT '到期日期',
  `total_amount` DECIMAL(18,2) NOT NULL COMMENT '总金额',
  `paid_amount` DECIMAL(18,2) DEFAULT 0 COMMENT '已付金额',
  `status` VARCHAR(20) DEFAULT 'unpaid' COMMENT '状态（unpaid, partially_paid, paid, cancelled）',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`order_id`) REFERENCES `sales_orders` (`id`),
  FOREIGN KEY (`customer_id`) REFERENCES `sales_customers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售发票表';

-- 2.11 销售发票明细表（sales_invoice_items）
CREATE TABLE IF NOT EXISTS `sales_invoice_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `invoice_id` VARCHAR(36) NOT NULL COMMENT '发票ID',
  `product_id` VARCHAR(36) NOT NULL COMMENT '产品ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `unit_price` DECIMAL(18,2) NOT NULL COMMENT '单价',
  `discount` DECIMAL(18,2) DEFAULT 0 COMMENT '折扣',
  `amount` DECIMAL(18,2) NOT NULL COMMENT '金额',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`invoice_id`) REFERENCES `sales_invoices` (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `sales_products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售发票明细表';

-- 2.12 销售退货单表（sales_returns）
CREATE TABLE IF NOT EXISTS `sales_returns` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '退货单ID',
  `return_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '退货单编号',
  `order_id` VARCHAR(36) NOT NULL COMMENT '订单ID',
  `customer_id` VARCHAR(36) NOT NULL COMMENT '客户ID',
  `return_date` DATE NOT NULL COMMENT '退货日期',
  `total_amount` DECIMAL(18,2) NOT NULL COMMENT '总金额',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, approved, processed, cancelled）',
  `reason` TEXT COMMENT '退货原因',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`order_id`) REFERENCES `sales_orders` (`id`),
  FOREIGN KEY (`customer_id`) REFERENCES `sales_customers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售退货单表';

-- 2.13 销售退货单明细表（sales_return_items）
CREATE TABLE IF NOT EXISTS `sales_return_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `return_id` VARCHAR(36) NOT NULL COMMENT '退货单ID',
  `product_id` VARCHAR(36) NOT NULL COMMENT '产品ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '退货数量',
  `unit_price` DECIMAL(18,2) NOT NULL COMMENT '单价',
  `amount` DECIMAL(18,2) NOT NULL COMMENT '金额',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`return_id`) REFERENCES `sales_returns` (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `sales_products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='销售退货单明细表';

-- 2.14 客户信用评估表（sales_customer_credits）
CREATE TABLE IF NOT EXISTS `sales_customer_credits` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '评估ID',
  `customer_id` VARCHAR(36) NOT NULL COMMENT '客户ID',
  `evaluation_date` DATE NOT NULL COMMENT '评估日期',
  `credit_score` INT NOT NULL COMMENT '信用评分',
  `risk_level` VARCHAR(20) NOT NULL COMMENT '风险等级（low, medium, high）',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`customer_id`) REFERENCES `sales_customers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='客户信用评估表';

-- 3. 库存模块

-- 3.1 仓库表（inventory_warehouses）
CREATE TABLE IF NOT EXISTS `inventory_warehouses` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '仓库ID',
  `code` VARCHAR(20) UNIQUE NOT NULL COMMENT '仓库编码',
  `name` VARCHAR(100) NOT NULL COMMENT '仓库名称',
  `type` VARCHAR(20) NOT NULL COMMENT '仓库类型（raw_material, finished_goods, semi_finished, tools）',
  `address` VARCHAR(255) COMMENT '仓库地址',
  `region` VARCHAR(50) COMMENT '地区',
  `contact` VARCHAR(50) COMMENT '联系人',
  `phone` VARCHAR(20) COMMENT '联系电话',
  `description` TEXT COMMENT '仓库描述',
  `capacity` INT DEFAULT 0 COMMENT '容量',
  `used_capacity` INT DEFAULT 0 COMMENT '已使用容量',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='仓库表';

-- 3.2 库位表（inventory_locations）
CREATE TABLE IF NOT EXISTS `inventory_locations` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '库位ID',
  `warehouse_id` VARCHAR(36) NOT NULL COMMENT '仓库ID',
  `code` VARCHAR(20) NOT NULL COMMENT '库位编码',
  `name` VARCHAR(100) NOT NULL COMMENT '库位名称',
  `type` VARCHAR(20) COMMENT '库位类型',
  `capacity` DECIMAL(18,4) COMMENT '容量',
  `used_capacity` DECIMAL(18,4) DEFAULT 0 COMMENT '已使用容量',
  `status` VARCHAR(20) DEFAULT 'available' COMMENT '状态（available, occupied, blocked）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`warehouse_id`) REFERENCES `inventory_warehouses` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库位表';

-- 3.3 物料表（inventory_items）
CREATE TABLE IF NOT EXISTS `inventory_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '物料ID',
  `item_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '物料编号',
  `name` VARCHAR(100) NOT NULL COMMENT '物料名称',
  `description` TEXT COMMENT '物料描述',
  `category_id` VARCHAR(36) COMMENT '物料类别ID',
  `unit` VARCHAR(10) NOT NULL COMMENT '单位',
  `type` VARCHAR(20) NOT NULL COMMENT '物料类型（raw_material, finished_goods, semi_finished, tool）',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive, obsolete）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物料表';

-- 3.4 物料类别表（inventory_item_categories）
CREATE TABLE IF NOT EXISTS `inventory_item_categories` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '类别ID',
  `code` VARCHAR(20) UNIQUE NOT NULL COMMENT '类别编码',
  `name` VARCHAR(100) NOT NULL COMMENT '类别名称',
  `parent_id` VARCHAR(36) COMMENT '父类别ID',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`parent_id`) REFERENCES `inventory_item_categories` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物料类别表';

-- 3.5 库存表（inventory_on_hand）
CREATE TABLE IF NOT EXISTS `inventory_on_hand` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '库存ID',
  `item_id` VARCHAR(36) NOT NULL COMMENT '物料ID',
  `warehouse_id` VARCHAR(36) NOT NULL COMMENT '仓库ID',
  `location_id` VARCHAR(36) COMMENT '库位ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `unit_cost` DECIMAL(18,2) NOT NULL COMMENT '单位成本',
  `total_cost` DECIMAL(18,2) NOT NULL COMMENT '总成本',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`item_id`) REFERENCES `inventory_items` (`id`),
  FOREIGN KEY (`warehouse_id`) REFERENCES `inventory_warehouses` (`id`),
  FOREIGN KEY (`location_id`) REFERENCES `inventory_locations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存表';

-- 3.6 库存交易表（inventory_transactions）
CREATE TABLE IF NOT EXISTS `inventory_transactions` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '交易ID',
  `transaction_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '交易编号',
  `item_id` VARCHAR(36) NOT NULL COMMENT '物料ID',
  `warehouse_id` VARCHAR(36) NOT NULL COMMENT '仓库ID',
  `location_id` VARCHAR(36) COMMENT '库位ID',
  `type` VARCHAR(20) NOT NULL COMMENT '交易类型（purchase, sales, transfer, adjustment, production）',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `unit_cost` DECIMAL(18,2) NOT NULL COMMENT '单位成本',
  `total_cost` DECIMAL(18,2) NOT NULL COMMENT '总成本',
  `reference_type` VARCHAR(50) COMMENT '参考类型（purchase_order, sales_order, transfer_order, etc.）',
  `reference_id` VARCHAR(36) COMMENT '参考ID',
  `transaction_date` DATETIME NOT NULL COMMENT '交易日期',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  FOREIGN KEY (`item_id`) REFERENCES `inventory_items` (`id`),
  FOREIGN KEY (`warehouse_id`) REFERENCES `inventory_warehouses` (`id`),
  FOREIGN KEY (`location_id`) REFERENCES `inventory_locations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存交易表';

-- 3.7 库存盘点表（inventory_counts）
CREATE TABLE IF NOT EXISTS `inventory_counts` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '盘点ID',
  `count_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '盘点编号',
  `warehouse_id` VARCHAR(36) NOT NULL COMMENT '仓库ID',
  `count_date` DATE NOT NULL COMMENT '盘点日期',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, in_progress, completed, cancelled）',
  `total_items` INT NOT NULL COMMENT '总物料数',
  `total_variance` DECIMAL(18,2) DEFAULT 0 COMMENT '总差异金额',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`warehouse_id`) REFERENCES `inventory_warehouses` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存盘点表';

-- 3.8 库存盘点明细表（inventory_count_items）
CREATE TABLE IF NOT EXISTS `inventory_count_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `count_id` VARCHAR(36) NOT NULL COMMENT '盘点ID',
  `item_id` VARCHAR(36) NOT NULL COMMENT '物料ID',
  `location_id` VARCHAR(36) COMMENT '库位ID',
  `system_quantity` DECIMAL(18,4) NOT NULL COMMENT '系统数量',
  `actual_quantity` DECIMAL(18,4) NOT NULL COMMENT '实际数量',
  `variance_quantity` DECIMAL(18,4) NOT NULL COMMENT '差异数量',
  `unit_cost` DECIMAL(18,2) NOT NULL COMMENT '单位成本',
  `variance_amount` DECIMAL(18,2) NOT NULL COMMENT '差异金额',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`count_id`) REFERENCES `inventory_counts` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `inventory_items` (`id`),
  FOREIGN KEY (`location_id`) REFERENCES `inventory_locations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存盘点明细表';

-- 4. 采购模块

-- 4.1 供应商表（purchase_vendors）
CREATE TABLE IF NOT EXISTS `purchase_vendors` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '供应商ID',
  `vendor_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '供应商编号',
  `name` VARCHAR(100) NOT NULL COMMENT '供应商名称',
  `contact_person` VARCHAR(50) COMMENT '联系人',
  `phone` VARCHAR(20) COMMENT '电话',
  `email` VARCHAR(100) COMMENT '邮箱',
  `address` VARCHAR(255) COMMENT '地址',
  `tax_no` VARCHAR(30) COMMENT '税号',
  `bank_name` VARCHAR(100) COMMENT '银行名称',
  `bank_account` VARCHAR(50) COMMENT '银行账号',
  `category` VARCHAR(50) COMMENT '供应商类别',
  `credit_limit` DECIMAL(18,2) DEFAULT 0 COMMENT '信用额度',
  `lead_time` INT DEFAULT 0 COMMENT '前置时间（天）',
  `payment_terms` VARCHAR(50) COMMENT '付款条件',
  `remarks` TEXT COMMENT '备注',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='供应商表';

-- 4.2 采购计划体表（purchase_plans）
CREATE TABLE IF NOT EXISTS `purchase_plans` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '计划ID',
  `plan_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '计划编号',
  `name` VARCHAR(100) NOT NULL COMMENT '计划名称',
  `description` TEXT COMMENT '计划描述',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `total_amount` DECIMAL(18,2) NOT NULL COMMENT '总金额',
  `status` VARCHAR(20) DEFAULT 'draft' COMMENT '状态（draft, approved, executed, cancelled）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购计划体表';

-- 4.3 采购计划表明细表（purchase_plan_items）
CREATE TABLE IF NOT EXISTS `purchase_plan_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `plan_id` VARCHAR(36) NOT NULL COMMENT '计划ID',
  `item_id` VARCHAR(36) NOT NULL COMMENT '物料ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `estimated_price` DECIMAL(18,2) NOT NULL COMMENT '预估价格',
  `estimated_amount` DECIMAL(18,2) NOT NULL COMMENT '预估金额',
  `need_date` DATE NOT NULL COMMENT '需求日期',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`plan_id`) REFERENCES `purchase_plans` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `inventory_items` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购计划表明细表';

-- 4.4 采购订单表（purchase_orders）
CREATE TABLE IF NOT EXISTS `purchase_orders` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '订单ID',
  `order_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '订单编号',
  `vendor_id` VARCHAR(36) NOT NULL COMMENT '供应商ID',
  `order_date` DATE NOT NULL COMMENT '订单日期',
  `delivery_date` DATE COMMENT '交货日期',
  `total_amount` DECIMAL(18,2) NOT NULL COMMENT '总金额',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, approved, delivered, invoiced, cancelled）',
  `payment_terms` VARCHAR(50) COMMENT '付款条件',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`vendor_id`) REFERENCES `purchase_vendors` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购订单表';

-- 4.5 采购订单明细表（purchase_order_items）
CREATE TABLE IF NOT EXISTS `purchase_order_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `order_id` VARCHAR(36) NOT NULL COMMENT '订单ID',
  `item_id` VARCHAR(36) NOT NULL COMMENT '物料ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `unit_price` DECIMAL(18,2) NOT NULL COMMENT '单价',
  `discount` DECIMAL(18,2) DEFAULT 0 COMMENT '折扣',
  `amount` DECIMAL(18,2) NOT NULL COMMENT '金额',
  `received_quantity` DECIMAL(18,4) DEFAULT 0 COMMENT '已收货数量',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`order_id`) REFERENCES `purchase_orders` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `inventory_items` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购订单明细表';

-- 4.6 采购收货单表（purchase_receipts）
CREATE TABLE IF NOT EXISTS `purchase_receipts` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '收货单ID',
  `receipt_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '收货单编号',
  `order_id` VARCHAR(36) NOT NULL COMMENT '订单ID',
  `vendor_id` VARCHAR(36) NOT NULL COMMENT '供应商ID',
  `receipt_date` DATE NOT NULL COMMENT '收货日期',
  `total_quantity` DECIMAL(18,4) NOT NULL COMMENT '总数量',
  `total_amount` DECIMAL(18,2) NOT NULL COMMENT '总金额',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, received, cancelled）',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`order_id`) REFERENCES `purchase_orders` (`id`),
  FOREIGN KEY (`vendor_id`) REFERENCES `purchase_vendors` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购收货单表';

-- 4.7 采购收货单明细表（purchase_receipt_items）
CREATE TABLE IF NOT EXISTS `purchase_receipt_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `receipt_id` VARCHAR(36) NOT NULL COMMENT '收货单ID',
  `order_item_id` VARCHAR(36) NOT NULL COMMENT '订单明细ID',
  `item_id` VARCHAR(36) NOT NULL COMMENT '物料ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '收货数量',
  `unit_price` DECIMAL(18,2) NOT NULL COMMENT '单价',
  `amount` DECIMAL(18,2) NOT NULL COMMENT '金额',
  `warehouse_id` VARCHAR(36) NOT NULL COMMENT '仓库ID',
  `location_id` VARCHAR(36) COMMENT '库位ID',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`receipt_id`) REFERENCES `purchase_receipts` (`id`),
  FOREIGN KEY (`order_item_id`) REFERENCES `purchase_order_items` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `inventory_items` (`id`),
  FOREIGN KEY (`warehouse_id`) REFERENCES `inventory_warehouses` (`id`),
  FOREIGN KEY (`location_id`) REFERENCES `inventory_locations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购收货单明细表';

-- 4.8 采购发票表（purchase_invoices）
CREATE TABLE IF NOT EXISTS `purchase_invoices` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '发票ID',
  `invoice_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '发票编号',
  `order_id` VARCHAR(36) NOT NULL COMMENT '订单ID',
  `vendor_id` VARCHAR(36) NOT NULL COMMENT '供应商ID',
  `invoice_date` DATE NOT NULL COMMENT '开票日期',
  `due_date` DATE NOT NULL COMMENT '到期日期',
  `total_amount` DECIMAL(18,2) NOT NULL COMMENT '总金额',
  `paid_amount` DECIMAL(18,2) DEFAULT 0 COMMENT '已付金额',
  `status` VARCHAR(20) DEFAULT 'unpaid' COMMENT '状态（unpaid, partially_paid, paid, cancelled）',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`order_id`) REFERENCES `purchase_orders` (`id`),
  FOREIGN KEY (`vendor_id`) REFERENCES `purchase_vendors` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购发票表';

-- 4.9 采购发票明细表（purchase_invoice_items）
CREATE TABLE IF NOT EXISTS `purchase_invoice_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `invoice_id` VARCHAR(36) NOT NULL COMMENT '发票ID',
  `item_id` VARCHAR(36) NOT NULL COMMENT '物料ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `unit_price` DECIMAL(18,2) NOT NULL COMMENT '单价',
  `discount` DECIMAL(18,2) DEFAULT 0 COMMENT '折扣',
  `amount` DECIMAL(18,2) NOT NULL COMMENT '金额',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`invoice_id`) REFERENCES `purchase_invoices` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `inventory_items` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购发票明细表';

-- 4.10 采购退货单表（purchase_returns）
CREATE TABLE IF NOT EXISTS `purchase_returns` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '退货单ID',
  `return_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '退货单编号',
  `order_id` VARCHAR(36) NOT NULL COMMENT '订单ID',
  `vendor_id` VARCHAR(36) NOT NULL COMMENT '供应商ID',
  `return_date` DATE NOT NULL COMMENT '退货日期',
  `total_amount` DECIMAL(18,2) NOT NULL COMMENT '总金额',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, approved, processed, cancelled）',
  `reason` TEXT COMMENT '退货原因',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`order_id`) REFERENCES `purchase_orders` (`id`),
  FOREIGN KEY (`vendor_id`) REFERENCES `purchase_vendors` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购退货单表';

-- 4.11 采购退货单明细表（purchase_return_items）
CREATE TABLE IF NOT EXISTS `purchase_return_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `return_id` VARCHAR(36) NOT NULL COMMENT '退货单ID',
  `item_id` VARCHAR(36) NOT NULL COMMENT '物料ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '退货数量',
  `unit_price` DECIMAL(18,2) NOT NULL COMMENT '单价',
  `amount` DECIMAL(18,2) NOT NULL COMMENT '金额',
  `warehouse_id` VARCHAR(36) NOT NULL COMMENT '仓库ID',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`return_id`) REFERENCES `purchase_returns` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `inventory_items` (`id`),
  FOREIGN KEY (`warehouse_id`) REFERENCES `inventory_warehouses` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采购退货单明细表';

-- 5. 财务模块

-- 5.1 会计科目表（finance_accounts）
CREATE TABLE IF NOT EXISTS `finance_accounts` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '科目ID',
  `code` VARCHAR(20) UNIQUE NOT NULL COMMENT '科目编码',
  `name` VARCHAR(100) NOT NULL COMMENT '科目名称',
  `type` VARCHAR(20) NOT NULL COMMENT '科目类型（asset, liability, equity, revenue, expense）',
  `level` INT NOT NULL COMMENT '科目级别',
  `parent_id` VARCHAR(36) COMMENT '父科目ID',
  `category` VARCHAR(50) COMMENT '科目类别',
  `description` TEXT COMMENT '科目描述',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`parent_id`) REFERENCES `finance_accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='会计科目表';

-- 5.2 凭证表（finance_journals）
CREATE TABLE IF NOT EXISTS `finance_journals` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '凭证ID',
  `journal_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '凭证编号',
  `date` DATE NOT NULL COMMENT '凭证日期',
  `reference` VARCHAR(100) COMMENT '参考号',
  `description` TEXT COMMENT '凭证描述',
  `total_debit` DECIMAL(18,2) NOT NULL COMMENT '借方合计',
  `total_credit` DECIMAL(18,2) NOT NULL COMMENT '贷方合计',
  `status` VARCHAR(20) DEFAULT 'draft' COMMENT '状态（draft, posted, cancelled）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='凭证表';

-- 5.3 凭证明细表（finance_journal_items）
CREATE TABLE IF NOT EXISTS `finance_journal_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `journal_id` VARCHAR(36) NOT NULL COMMENT '凭证ID',
  `account_id` VARCHAR(36) NOT NULL COMMENT '科目ID',
  `description` TEXT COMMENT '明细描述',
  `debit` DECIMAL(18,2) DEFAULT 0 COMMENT '借方金额',
  `credit` DECIMAL(18,2) DEFAULT 0 COMMENT '贷方金额',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`journal_id`) REFERENCES `finance_journals` (`id`),
  FOREIGN KEY (`account_id`) REFERENCES `finance_accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='凭证明细表';

-- 5.4 固定资产表（finance_fixed_assets）
CREATE TABLE IF NOT EXISTS `finance_fixed_assets` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '资产ID',
  `asset_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '资产编号',
  `name` VARCHAR(100) NOT NULL COMMENT '资产名称',
  `description` TEXT COMMENT '资产描述',
  `category` VARCHAR(50) NOT NULL COMMENT '资产类别',
  `acquisition_date` DATE NOT NULL COMMENT '购置日期',
  `cost` DECIMAL(18,2) NOT NULL COMMENT '购置成本',
  `salvage_value` DECIMAL(18,2) DEFAULT 0 COMMENT '残值',
  `useful_life` INT NOT NULL COMMENT '使用年限（月）',
  `depreciation_method` VARCHAR(50) NOT NULL COMMENT '折旧方法',
  `accumulated_depreciation` DECIMAL(18,2) DEFAULT 0 COMMENT '累计折旧',
  `net_book_value` DECIMAL(18,2) NOT NULL COMMENT '净值',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, disposed, sold）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='固定资产表';

-- 5.5 预算表（finance_budgets）
CREATE TABLE IF NOT EXISTS `finance_budgets` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '预算ID',
  `budget_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '预算编号',
  `name` VARCHAR(100) NOT NULL COMMENT '预算名称',
  `description` TEXT COMMENT '预算描述',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `total_amount` DECIMAL(18,2) NOT NULL COMMENT '总预算金额',
  `status` VARCHAR(20) DEFAULT 'draft' COMMENT '状态（draft, approved, executed, closed）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='预算表';

-- 5.6 预算明细表（finance_budget_items）
CREATE TABLE IF NOT EXISTS `finance_budget_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `budget_id` VARCHAR(36) NOT NULL COMMENT '预算ID',
  `account_id` VARCHAR(36) NOT NULL COMMENT '科目ID',
  `amount` DECIMAL(18,2) NOT NULL COMMENT '预算金额',
  `description` TEXT COMMENT '明细描述',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`budget_id`) REFERENCES `finance_budgets` (`id`),
  FOREIGN KEY (`account_id`) REFERENCES `finance_accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='预算明细表';

-- 5.7 成本中心表（finance_cost_centers）
CREATE TABLE IF NOT EXISTS `finance_cost_centers` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '成本中心ID',
  `code` VARCHAR(20) UNIQUE NOT NULL COMMENT '成本中心编码',
  `name` VARCHAR(100) NOT NULL COMMENT '成本中心名称',
  `description` TEXT COMMENT '成本中心描述',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='成本中心表';

-- 6. 生产模块

-- 6.1 物料清单表（production_boms）
CREATE TABLE IF NOT EXISTS `production_boms` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT 'BOM ID',
  `bom_no` VARCHAR(20) UNIQUE NOT NULL COMMENT 'BOM编号',
  `product_id` VARCHAR(36) NOT NULL COMMENT '产品ID',
  `version` VARCHAR(10) NOT NULL COMMENT '版本号',
  `description` TEXT COMMENT 'BOM描述',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`product_id`) REFERENCES `sales_products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物料清单表';

-- 6.2 物料清单明细表（production_bom_items）
CREATE TABLE IF NOT EXISTS `production_bom_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `bom_id` VARCHAR(36) NOT NULL COMMENT 'BOM ID',
  `item_id` VARCHAR(36) NOT NULL COMMENT '物料ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `unit` VARCHAR(10) NOT NULL COMMENT '单位',
  `scrap_rate` DECIMAL(18,4) DEFAULT 0 COMMENT '损耗率',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`bom_id`) REFERENCES `production_boms` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `inventory_items` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物料清单明细表';

-- 6.3 主生产计划表（production_mps）
CREATE TABLE IF NOT EXISTS `production_mps` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT 'MPS ID',
  `mps_no` VARCHAR(20) UNIQUE NOT NULL COMMENT 'MPS编号',
  `name` VARCHAR(100) NOT NULL COMMENT 'MPS名称',
  `description` TEXT COMMENT 'MPS描述',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `status` VARCHAR(20) DEFAULT 'draft' COMMENT '状态（draft, approved, executed, closed）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='主生产计划表';

-- 6.4 主生产计划明细表（production_mps_items）
CREATE TABLE IF NOT EXISTS `production_mps_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `mps_id` VARCHAR(36) NOT NULL COMMENT 'MPS ID',
  `product_id` VARCHAR(36) NOT NULL COMMENT '产品ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `schedule_date` DATE NOT NULL COMMENT '计划日期',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`mps_id`) REFERENCES `production_mps` (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `sales_products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='主生产计划明细表';

-- 6.5 物料需求计划表（production_mrp）
CREATE TABLE IF NOT EXISTS `production_mrp` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT 'MRP ID',
  `mrp_no` VARCHAR(20) UNIQUE NOT NULL COMMENT 'MRP编号',
  `name` VARCHAR(100) NOT NULL COMMENT 'MRP名称',
  `description` TEXT COMMENT 'MRP描述',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `status` VARCHAR(20) DEFAULT 'draft' COMMENT '状态（draft, approved, executed, closed）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物料需求计划表';

-- 6.6 物料需求计划明细表（production_mrp_items）
CREATE TABLE IF NOT EXISTS `production_mrp_items` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '明细ID',
  `mrp_id` VARCHAR(36) NOT NULL COMMENT 'MRP ID',
  `item_id` VARCHAR(36) NOT NULL COMMENT '物料ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `schedule_date` DATE NOT NULL COMMENT '计划日期',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`mrp_id`) REFERENCES `production_mrp` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `inventory_items` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物料需求计划明细表';

-- 6.7 生产订单表（production_orders）
CREATE TABLE IF NOT EXISTS `production_orders` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '生产订单ID',
  `order_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '生产订单编号',
  `product_id` VARCHAR(36) NOT NULL COMMENT '产品ID',
  `mps_item_id` VARCHAR(36) COMMENT 'MPS明细ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `status` VARCHAR(20) DEFAULT 'planned' COMMENT '状态（planned, released, in_progress, completed, cancelled）',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`product_id`) REFERENCES `sales_products` (`id`),
  FOREIGN KEY (`mps_item_id`) REFERENCES `production_mps_items` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='生产订单表';

-- 6.8 生产工单表（production_work_orders）
CREATE TABLE IF NOT EXISTS `production_work_orders` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '生产工单ID',
  `work_order_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '生产工单编号',
  `production_order_id` VARCHAR(36) NOT NULL COMMENT '生产订单ID',
  `operation_seq` INT NOT NULL COMMENT '工序序号',
  `work_center_id` VARCHAR(36) NOT NULL COMMENT '工作中心ID',
  `quantity` DECIMAL(18,4) NOT NULL COMMENT '数量',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, in_progress, completed, cancelled）',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`production_order_id`) REFERENCES `production_orders` (`id`),
  FOREIGN KEY (`work_center_id`) REFERENCES `production_work_centers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='生产工单表';

-- 6.9 工艺路线表（production_routings）
CREATE TABLE IF NOT EXISTS `production_routings` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '工艺路线ID',
  `routing_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '工艺路线编号',
  `product_id` VARCHAR(36) NOT NULL COMMENT '产品ID',
  `version` VARCHAR(10) NOT NULL COMMENT '版本号',
  `description` TEXT COMMENT '工艺路线描述',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`product_id`) REFERENCES `sales_products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工艺路线表';

-- 6.10 工艺路线工序表（production_routing_operations）
CREATE TABLE IF NOT EXISTS `production_routing_operations` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '工序ID',
  `routing_id` VARCHAR(36) NOT NULL COMMENT '工艺路线ID',
  `operation_seq` INT NOT NULL COMMENT '工序序号',
  `work_center_id` VARCHAR(36) NOT NULL COMMENT '工作中心ID',
  `operation_name` VARCHAR(100) NOT NULL COMMENT '工序名称',
  `setup_time` DECIMAL(18,2) DEFAULT 0 COMMENT '准备时间（小时）',
  `run_time` DECIMAL(18,4) DEFAULT 0 COMMENT '加工时间（小时/件）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`routing_id`) REFERENCES `production_routings` (`id`),
  FOREIGN KEY (`work_center_id`) REFERENCES `production_work_centers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工艺路线工序表';

-- 6.11 工作中心表（production_work_centers）
CREATE TABLE IF NOT EXISTS `production_work_centers` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '工作中心ID',
  `code` VARCHAR(20) UNIQUE NOT NULL COMMENT '工作中心编码',
  `name` VARCHAR(100) NOT NULL COMMENT '工作中心名称',
  `description` TEXT COMMENT '工作中心描述',
  `location` VARCHAR(255) COMMENT '位置',
  `capacity` INT DEFAULT 0 COMMENT '产能（小时/天）',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工作中心表';

-- 7. 人力资源模块

-- 7.1 员工表（hr_employees）
CREATE TABLE IF NOT EXISTS `hr_employees` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '员工ID',
  `employee_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '员工编号',
  `name` VARCHAR(50) NOT NULL COMMENT '员工姓名',
  `gender` VARCHAR(10) COMMENT '性别',
  `birth_date` DATE COMMENT '出生日期',
  `id_card_no` VARCHAR(18) UNIQUE COMMENT '身份证号',
  `phone` VARCHAR(20) COMMENT '电话',
  `email` VARCHAR(100) COMMENT '邮箱',
  `address` VARCHAR(255) COMMENT '地址',
  `department_id` VARCHAR(36) COMMENT '部门ID',
  `position_id` VARCHAR(36) COMMENT '职位ID',
  `join_date` DATE NOT NULL COMMENT '入职日期',
  `probation_end_date` DATE COMMENT '试用期结束日期',
  `contract_start_date` DATE COMMENT '合同开始日期',
  `contract_end_date` DATE COMMENT '合同结束日期',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, resigned, on_leave）',
  `salary` DECIMAL(18,2) DEFAULT 0 COMMENT '薪资',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`department_id`) REFERENCES `hr_departments` (`id`),
  FOREIGN KEY (`position_id`) REFERENCES `hr_positions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='员工表';

-- 7.2 部门表（hr_departments）
CREATE TABLE IF NOT EXISTS `hr_departments` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '部门ID',
  `code` VARCHAR(20) UNIQUE NOT NULL COMMENT '部门编码',
  `name` VARCHAR(100) NOT NULL COMMENT '部门名称',
  `description` TEXT COMMENT '部门描述',
  `parent_id` VARCHAR(36) COMMENT '父部门ID',
  `manager_id` VARCHAR(36) COMMENT '部门经理ID',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`parent_id`) REFERENCES `hr_departments` (`id`),
  FOREIGN KEY (`manager_id`) REFERENCES `hr_employees` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门表';

-- 7.3 职位表（hr_positions）
CREATE TABLE IF NOT EXISTS `hr_positions` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '职位ID',
  `code` VARCHAR(20) UNIQUE NOT NULL COMMENT '职位编码',
  `name` VARCHAR(100) NOT NULL COMMENT '职位名称',
  `description` TEXT COMMENT '职位描述',
  `department_id` VARCHAR(36) COMMENT '部门ID',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`department_id`) REFERENCES `hr_departments` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='职位表';

-- 7.4 考勤表（hr_attendances）
CREATE TABLE IF NOT EXISTS `hr_attendances` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '考勤ID',
  `employee_id` VARCHAR(36) NOT NULL COMMENT '员工ID',
  `attendance_date` DATE NOT NULL COMMENT '考勤日期',
  `check_in_time` DATETIME COMMENT '上班时间',
  `check_out_time` DATETIME COMMENT '下班时间',
  `status` VARCHAR(20) DEFAULT 'present' COMMENT '状态（present, absent, late, leave_early, overtime）',
  `hours_worked` DECIMAL(18,2) DEFAULT 0 COMMENT '工作时长',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`employee_id`) REFERENCES `hr_employees` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='考勤表';

-- 7.5 请假表（hr_leave_applications）
CREATE TABLE IF NOT EXISTS `hr_leave_applications` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '请假ID',
  `leave_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '请假编号',
  `employee_id` VARCHAR(36) NOT NULL COMMENT '员工ID',
  `leave_type` VARCHAR(20) NOT NULL COMMENT '请假类型（annual, sick, personal, maternity）',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `days` DECIMAL(18,2) NOT NULL COMMENT '请假天数',
  `reason` TEXT NOT NULL COMMENT '请假原因',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, approved, rejected, cancelled）',
  `approved_by` VARCHAR(36) COMMENT '审批人ID',
  `approved_at` DATETIME COMMENT '审批时间',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`employee_id`) REFERENCES `hr_employees` (`id`),
  FOREIGN KEY (`approved_by`) REFERENCES `hr_employees` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='请假表';

-- 7.6 加班申请表（hr_overtime_applications）
CREATE TABLE IF NOT EXISTS `hr_overtime_applications` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '加班ID',
  `overtime_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '加班编号',
  `employee_id` VARCHAR(36) NOT NULL COMMENT '员工ID',
  `overtime_date` DATE NOT NULL COMMENT '加班日期',
  `start_time` DATETIME NOT NULL COMMENT '开始时间',
  `end_time` DATETIME NOT NULL COMMENT '结束时间',
  `hours` DECIMAL(18,2) NOT NULL COMMENT '加班时长',
  `reason` TEXT NOT NULL COMMENT '加班原因',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, approved, rejected, cancelled）',
  `approved_by` VARCHAR(36) COMMENT '审批人ID',
  `approved_at` DATETIME COMMENT '审批时间',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`employee_id`) REFERENCES `hr_employees` (`id`),
  FOREIGN KEY (`approved_by`) REFERENCES `hr_employees` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='加班申请表';

-- 7.7 薪资表（hr_salaries）
CREATE TABLE IF NOT EXISTS `hr_salaries` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '薪资ID',
  `salary_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '薪资编号',
  `employee_id` VARCHAR(36) NOT NULL COMMENT '员工ID',
  `month` VARCHAR(7) NOT NULL COMMENT '月份（YYYY-MM）',
  `basic_salary` DECIMAL(18,2) DEFAULT 0 COMMENT '基本工资',
  `allowance` DECIMAL(18,2) DEFAULT 0 COMMENT '津贴',
  `bonus` DECIMAL(18,2) DEFAULT 0 COMMENT '奖金',
  `overtime_pay` DECIMAL(18,2) DEFAULT 0 COMMENT '加班费',
  `deductions` DECIMAL(18,2) DEFAULT 0 COMMENT '扣除项',
  `net_salary` DECIMAL(18,2) DEFAULT 0 COMMENT '实发工资',
  `status` VARCHAR(20) DEFAULT 'calculated' COMMENT '状态（calculated, approved, paid）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`employee_id`) REFERENCES `hr_employees` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='薪资表';

-- 7.8 培训表（hr_trainings）
CREATE TABLE IF NOT EXISTS `hr_trainings` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '培训ID',
  `training_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '培训编号',
  `name` VARCHAR(100) NOT NULL COMMENT '培训名称',
  `description` TEXT COMMENT '培训描述',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `location` VARCHAR(255) COMMENT '培训地点',
  `trainer` VARCHAR(100) COMMENT '培训师',
  `status` VARCHAR(20) DEFAULT 'planned' COMMENT '状态（planned, in_progress, completed, cancelled）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='培训表';

-- 7.9 培训参加表（hr_training_participants）
CREATE TABLE IF NOT EXISTS `hr_training_participants` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '参加ID',
  `training_id` VARCHAR(36) NOT NULL COMMENT '培训ID',
  `employee_id` VARCHAR(36) NOT NULL COMMENT '员工ID',
  `attendance` VARCHAR(20) DEFAULT 'registered' COMMENT '考勤状态（registered, attended, absent）',
  `score` DECIMAL(18,2) COMMENT '分数',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`training_id`) REFERENCES `hr_trainings` (`id`),
  FOREIGN KEY (`employee_id`) REFERENCES `hr_employees` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='培训参加表';

-- 7.10 招聘表（hr_recruitments）
CREATE TABLE IF NOT EXISTS `hr_recruitments` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '招聘ID',
  `recruitment_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '招聘编号',
  `position_id` VARCHAR(36) NOT NULL COMMENT '职位ID',
  `department_id` VARCHAR(36) NOT NULL COMMENT '部门ID',
  `recruit_count` INT NOT NULL COMMENT '招聘人数',
  `requirement` TEXT COMMENT '职位要求',
  `salary_range` VARCHAR(100) COMMENT '薪资范围',
  `start_date` DATE NOT NULL COMMENT '开始日期',
  `end_date` DATE NOT NULL COMMENT '结束日期',
  `status` VARCHAR(20) DEFAULT 'open' COMMENT '状态（open, closed, cancelled）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`position_id`) REFERENCES `hr_positions` (`id`),
  FOREIGN KEY (`department_id`) REFERENCES `hr_departments` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='招聘表';

-- 7.11 简历表（hr_resumes）
CREATE TABLE IF NOT EXISTS `hr_resumes` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '简历ID',
  `resume_no` VARCHAR(20) UNIQUE NOT NULL COMMENT '简历编号',
  `recruitment_id` VARCHAR(36) NOT NULL COMMENT '招聘ID',
  `name` VARCHAR(50) NOT NULL COMMENT '姓名',
  `gender` VARCHAR(10) COMMENT '性别',
  `phone` VARCHAR(20) NOT NULL COMMENT '电话',
  `email` VARCHAR(100) NOT NULL COMMENT '邮箱',
  `education` VARCHAR(50) COMMENT '学历',
  `work_experience` TEXT COMMENT '工作经历',
  `skills` TEXT COMMENT '技能',
  `status` VARCHAR(20) DEFAULT 'pending' COMMENT '状态（pending, screening, interview, offer, rejected）',
  `interview_result` TEXT COMMENT '面试结果',
  `remarks` TEXT COMMENT '备注',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`recruitment_id`) REFERENCES `hr_recruitments` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='简历表';

-- 8. 系统模块

-- 8.1 用户表（system_users）
CREATE TABLE IF NOT EXISTS `system_users` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '用户ID',
  `username` VARCHAR(50) UNIQUE NOT NULL COMMENT '用户名',
  `password_hash` VARCHAR(255) NOT NULL COMMENT '密码哈希',
  `employee_id` VARCHAR(36) COMMENT '员工ID',
  `name` VARCHAR(50) NOT NULL COMMENT '姓名',
  `phone` VARCHAR(20) COMMENT '电话',
  `email` VARCHAR(100) COMMENT '邮箱',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive, locked）',
  `last_login_at` DATETIME COMMENT '最后登录时间',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  FOREIGN KEY (`employee_id`) REFERENCES `hr_employees` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 8.2 角色表（system_roles）
CREATE TABLE IF NOT EXISTS `system_roles` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '角色ID',
  `code` VARCHAR(20) UNIQUE NOT NULL COMMENT '角色编码',
  `name` VARCHAR(100) NOT NULL COMMENT '角色名称',
  `description` TEXT COMMENT '角色描述',
  `status` VARCHAR(20) DEFAULT 'active' COMMENT '状态（active, inactive）',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- 8.3 用户角色表（system_user_roles）
CREATE TABLE IF NOT EXISTS `system_user_roles` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT 'ID',
  `user_id` VARCHAR(36) NOT NULL COMMENT '用户ID',
  `role_id` VARCHAR(36) NOT NULL COMMENT '角色ID',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  FOREIGN KEY (`user_id`) REFERENCES `system_users` (`id`),
  FOREIGN KEY (`role_id`) REFERENCES `system_roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色表';

-- 8.4 权限表（system_permissions）
CREATE TABLE IF NOT EXISTS `system_permissions` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '权限ID',
  `code` VARCHAR(50) UNIQUE NOT NULL COMMENT '权限编码',
  `name` VARCHAR(100) NOT NULL COMMENT '权限名称',
  `description` TEXT COMMENT '权限描述',
  `module` VARCHAR(50) NOT NULL COMMENT '模块',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` VARCHAR(36) COMMENT '更新人',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限表';

-- 8.5 角色权限表（system_role_permissions）
CREATE TABLE IF NOT EXISTS `system_role_permissions` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT 'ID',
  `role_id` VARCHAR(36) NOT NULL COMMENT '角色ID',
  `permission_id` VARCHAR(36) NOT NULL COMMENT '权限ID',
  `created_by` VARCHAR(36) COMMENT '创建人',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  FOREIGN KEY (`role_id`) REFERENCES `system_roles` (`id`),
  FOREIGN KEY (`permission_id`) REFERENCES `system_permissions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限表';

-- 8.6 日志表（system_logs）
CREATE TABLE IF NOT EXISTS `system_logs` (
  `id` VARCHAR(36) PRIMARY KEY COMMENT '日志ID',
  `user_id` VARCHAR(36) COMMENT '用户ID',
  `action` VARCHAR(100) NOT NULL COMMENT '操作',
  `module` VARCHAR(50) NOT NULL COMMENT '模块',
  `ip_address` VARCHAR(50) COMMENT 'IP地址',
  `user_agent` VARCHAR(255) COMMENT '用户代理',
  `request_url` VARCHAR(255) COMMENT '请求URL',
  `request_method` VARCHAR(10) COMMENT '请求方法',
  `status` VARCHAR(20) NOT NULL COMMENT '状态（success, error）',
  `message` TEXT COMMENT '消息',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  FOREIGN KEY (`user_id`) REFERENCES `system_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日志表';

-- 9. 创建索引

-- CRM模块索引
CREATE INDEX `idx_crm_accounts_account_no` ON `crm_accounts` (`account_no`);
CREATE INDEX `idx_crm_accounts_name` ON `crm_accounts` (`name`);
CREATE INDEX `idx_crm_contacts_account_id` ON `crm_contacts` (`account_id`);
CREATE INDEX `idx_crm_activities_account_id` ON `crm_activities` (`account_id`);
CREATE INDEX `idx_crm_activities_contact_id` ON `crm_activities` (`contact_id`);
CREATE INDEX `idx_crm_cases_account_id` ON `crm_cases` (`account_id`);
CREATE INDEX `idx_crm_cases_contact_id` ON `crm_cases` (`contact_id`);
CREATE INDEX `idx_crm_case_comments_case_id` ON `crm_case_comments` (`case_id`);
CREATE INDEX `idx_crm_leads_status` ON `crm_leads` (`status`);
CREATE INDEX `idx_crm_opportunities_account_id` ON `crm_opportunities` (`account_id`);
CREATE INDEX `idx_crm_opportunities_stage` ON `crm_opportunities` (`stage`);
CREATE INDEX `idx_crm_campaigns_status` ON `crm_campaigns` (`status`);

-- 销售模块索引
CREATE INDEX `idx_sales_customers_customer_no` ON `sales_customers` (`customer_no`);
CREATE INDEX `idx_sales_customers_name` ON `sales_customers` (`name`);
CREATE INDEX `idx_sales_products_product_no` ON `sales_products` (`product_no`);
CREATE INDEX `idx_sales_products_name` ON `sales_products` (`name`);
CREATE INDEX `idx_sales_product_categories_code` ON `sales_product_categories` (`code`);
CREATE INDEX `idx_sales_quotes_quote_no` ON `sales_quotes` (`quote_no`);
CREATE INDEX `idx_sales_quotes_customer_id` ON `sales_quotes` (`customer_id`);
CREATE INDEX `idx_sales_quote_items_quote_id` ON `sales_quote_items` (`quote_id`);
CREATE INDEX `idx_sales_quote_items_product_id` ON `sales_quote_items` (`product_id`);
CREATE INDEX `idx_sales_orders_order_no` ON `sales_orders` (`order_no`);
CREATE INDEX `idx_sales_orders_customer_id` ON `sales_orders` (`customer_id`);
CREATE INDEX `idx_sales_order_items_order_id` ON `sales_order_items` (`order_id`);
CREATE INDEX `idx_sales_order_items_product_id` ON `sales_order_items` (`product_id`);
CREATE INDEX `idx_sales_deliveries_delivery_no` ON `sales_deliveries` (`delivery_no`);
CREATE INDEX `idx_sales_deliveries_order_id` ON `sales_deliveries` (`order_id`);
CREATE INDEX `idx_sales_delivery_items_delivery_id` ON `sales_delivery_items` (`delivery_id`);
CREATE INDEX `idx_sales_invoices_invoice_no` ON `sales_invoices` (`invoice_no`);
CREATE INDEX `idx_sales_invoices_order_id` ON `sales_invoices` (`order_id`);
CREATE INDEX `idx_sales_invoices_customer_id` ON `sales_invoices` (`customer_id`);
CREATE INDEX `idx_sales_invoice_items_invoice_id` ON `sales_invoice_items` (`invoice_id`);
CREATE INDEX `idx_sales_returns_return_no` ON `sales_returns` (`return_no`);
CREATE INDEX `idx_sales_returns_order_id` ON `sales_returns` (`order_id`);
CREATE INDEX `idx_sales_return_items_return_id` ON `sales_return_items` (`return_id`);
CREATE INDEX `idx_sales_customer_credits_customer_id` ON `sales_customer_credits` (`customer_id`);

-- 库存模块索引
CREATE INDEX `idx_inventory_warehouses_code` ON `inventory_warehouses` (`code`);
CREATE INDEX `idx_inventory_warehouses_name` ON `inventory_warehouses` (`name`);
CREATE INDEX `idx_inventory_locations_warehouse_id` ON `inventory_locations` (`warehouse_id`);
CREATE INDEX `idx_inventory_locations_code` ON `inventory_locations` (`code`);
CREATE INDEX `idx_inventory_items_item_no` ON `inventory_items` (`item_no`);
CREATE INDEX `idx_inventory_items_name` ON `inventory_items` (`name`);
CREATE INDEX `idx_inventory_item_categories_code` ON `inventory_item_categories` (`code`);
CREATE INDEX `idx_inventory_on_hand_item_id` ON `inventory_on_hand` (`item_id`);
CREATE INDEX `idx_inventory_on_hand_warehouse_id` ON `inventory_on_hand` (`warehouse_id`);
CREATE INDEX `idx_inventory_transactions_transaction_no` ON `inventory_transactions` (`transaction_no`);
CREATE INDEX `idx_inventory_transactions_item_id` ON `inventory_transactions` (`item_id`);
CREATE INDEX `idx_inventory_transactions_warehouse_id` ON `inventory_transactions` (`warehouse_id`);
CREATE INDEX `idx_inventory_transactions_type` ON `inventory_transactions` (`type`);
CREATE INDEX `idx_inventory_transactions_transaction_date` ON `inventory_transactions` (`transaction_date`);
CREATE INDEX `idx_inventory_counts_count_no` ON `inventory_counts` (`count_no`);
CREATE INDEX `idx_inventory_counts_warehouse_id` ON `inventory_counts` (`warehouse_id`);
CREATE INDEX `idx_inventory_count_items_count_id` ON `inventory_count_items` (`count_id`);
CREATE INDEX `idx_inventory_count_items_item_id` ON `inventory_count_items` (`item_id`);

-- 采购模块索引
CREATE INDEX `idx_purchase_vendors_vendor_no` ON `purchase_vendors` (`vendor_no`);
CREATE INDEX `idx_purchase_vendors_name` ON `purchase_vendors` (`name`);
CREATE INDEX `idx_purchase_plans_plan_no` ON `purchase_plans` (`plan_no`);
CREATE INDEX `idx_purchase_plan_items_plan_id` ON `purchase_plan_items` (`plan_id`);
CREATE INDEX `idx_purchase_plan_items_item_id` ON `purchase_plan_items` (`item_id`);
CREATE INDEX `idx_purchase_orders_order_no` ON `purchase_orders` (`order_no`);
CREATE INDEX `idx_purchase_orders_vendor_id` ON `purchase_orders` (`vendor_id`);
CREATE INDEX `idx_purchase_order_items_order_id` ON `purchase_order_items` (`order_id`);
CREATE INDEX `idx_purchase_order_items_item_id` ON `purchase_order_items` (`item_id`);
CREATE INDEX `idx_purchase_receipts_receipt_no` ON `purchase_receipts` (`receipt_no`);
CREATE INDEX `idx_purchase_receipts_order_id` ON `purchase_receipts` (`order_id`);
CREATE INDEX `idx_purchase_receipt_items_receipt_id` ON `purchase_receipt_items` (`receipt_id`);
CREATE INDEX `idx_purchase_receipt_items_order_item_id` ON `purchase_receipt_items` (`order_item_id`);
CREATE INDEX `idx_purchase_invoices_invoice_no` ON `purchase_invoices` (`invoice_no`);
CREATE INDEX `idx_purchase_invoices_order_id` ON `purchase_invoices` (`order_id`);
CREATE INDEX `idx_purchase_invoices_vendor_id` ON `purchase_invoices` (`vendor_id`);
CREATE INDEX `idx_purchase_invoice_items_invoice_id` ON `purchase_invoice_items` (`invoice_id`);
CREATE INDEX `idx_purchase_returns_return_no` ON `purchase_returns` (`return_no`);
CREATE INDEX `idx_purchase_returns_order_id` ON `purchase_returns` (`order_id`);
CREATE INDEX `idx_purchase_return_items_return_id` ON `purchase_return_items` (`return_id`);

-- 财务模块索引
CREATE INDEX `idx_finance_accounts_code` ON `finance_accounts` (`code`);
CREATE INDEX `idx_finance_accounts_name` ON `finance_accounts` (`name`);
CREATE INDEX `idx_finance_accounts_type` ON `finance_accounts` (`type`);
CREATE INDEX `idx_finance_journals_journal_no` ON `finance_journals` (`journal_no`);
CREATE INDEX `idx_finance_journals_date` ON `finance_journals` (`date`);
CREATE INDEX `idx_finance_journal_items_journal_id` ON `finance_journal_items` (`journal_id`);
CREATE INDEX `idx_finance_journal_items_account_id` ON `finance_journal_items` (`account_id`);
CREATE INDEX `idx_finance_fixed_assets_asset_no` ON `finance_fixed_assets` (`asset_no`);
CREATE INDEX `idx_finance_fixed_assets_name` ON `finance_fixed_assets` (`name`);
CREATE INDEX `idx_finance_budgets_budget_no` ON `finance_budgets` (`budget_no`);
CREATE INDEX `idx_finance_budget_items_budget_id` ON `finance_budget_items` (`budget_id`);
CREATE INDEX `idx_finance_budget_items_account_id` ON `finance_budget_items` (`account_id`);
CREATE INDEX `idx_finance_cost_centers_code` ON `finance_cost_centers` (`code`);
CREATE INDEX `idx_finance_cost_centers_name` ON `finance_cost_centers` (`name`);

-- 生产模块索引
CREATE INDEX `idx_production_boms_bom_no` ON `production_boms` (`bom_no`);
CREATE INDEX `idx_production_boms_product_id` ON `production_boms` (`product_id`);
CREATE INDEX `idx_production_bom_items_bom_id` ON `production_bom_items` (`bom_id`);
CREATE INDEX `idx_production_bom_items_item_id` ON `production_bom_items` (`item_id`);
CREATE INDEX `idx_production_mps_mps_no` ON `production_mps` (`mps_no`);
CREATE INDEX `idx_production_mps_items_mps_id` ON `production_mps_items` (`mps_id`);
CREATE INDEX `idx_production_mps_items_product_id` ON `production_mps_items` (`product_id`);
CREATE INDEX `idx_production_mrp_mrp_no` ON `production_mrp` (`mrp_no`);
CREATE INDEX `idx_production_mrp_items_mrp_id` ON `production_mrp_items` (`mrp_id`);
CREATE INDEX `idx_production_mrp_items_item_id` ON `production_mrp_items` (`item_id`);
CREATE INDEX `idx_production_orders_order_no` ON `production_orders` (`order_no`);
CREATE INDEX `idx_production_orders_product_id` ON `production_orders` (`product_id`);
CREATE INDEX `idx_production_work_orders_work_order_no` ON `production_work_orders` (`work_order_no`);
CREATE INDEX `idx_production_work_orders_production_order_id` ON `production_work_orders` (`production_order_id`);
CREATE INDEX `idx_production_work_orders_work_center_id` ON `production_work_orders` (`work_center_id`);
CREATE INDEX `idx_production_routings_routing_no` ON `production_routings` (`routing_no`);
CREATE INDEX `idx_production_routings_product_id` ON `production_routings` (`product_id`);
CREATE INDEX `idx_production_routing_operations_routing_id` ON `production_routing_operations` (`routing_id`);
CREATE INDEX `idx_production_routing_operations_work_center_id` ON `production_routing_operations` (`work_center_id`);
CREATE INDEX `idx_production_work_centers_code` ON `production_work_centers` (`code`);
CREATE INDEX `idx_production_work_centers_name` ON `production_work_centers` (`name`);

-- 人力资源模块索引
CREATE INDEX `idx_hr_employees_employee_no` ON `hr_employees` (`employee_no`);
CREATE INDEX `idx_hr_employees_name` ON `hr_employees` (`name`);
CREATE INDEX `idx_hr_employees_department_id` ON `hr_employees` (`department_id`);
CREATE INDEX `idx_hr_employees_position_id` ON `hr_employees` (`position_id`);
CREATE INDEX `idx_hr_departments_code` ON `hr_departments` (`code`);
CREATE INDEX `idx_hr_departments_name` ON `hr_departments` (`name`);
CREATE INDEX `idx_hr_positions_code` ON `hr_positions` (`code`);
CREATE INDEX `idx_hr_positions_name` ON `hr_positions` (`name`);
CREATE INDEX `idx_hr_positions_department_id` ON `hr_positions` (`department_id`);
CREATE INDEX `idx_hr_attendances_employee_id` ON `hr_attendances` (`employee_id`);
CREATE INDEX `idx_hr_attendances_attendance_date` ON `hr_attendances` (`attendance_date`);
CREATE INDEX `idx_hr_leave_applications_leave_no` ON `hr_leave_applications` (`leave_no`);
CREATE INDEX `idx_hr_leave_applications_employee_id` ON `hr_leave_applications` (`employee_id`);
CREATE INDEX `idx_hr_overtime_applications_overtime_no` ON `hr_overtime_applications` (`overtime_no`);
CREATE INDEX `idx_hr_overtime_applications_employee_id` ON `hr_overtime_applications` (`employee_id`);
CREATE INDEX `idx_hr_salaries_salary_no` ON `hr_salaries` (`salary_no`);
CREATE INDEX `idx_hr_salaries_employee_id` ON `hr_salaries` (`employee_id`);
CREATE INDEX `idx_hr_salaries_month` ON `hr_salaries` (`month`);
CREATE INDEX `idx_hr_trainings_training_no` ON `hr_trainings` (`training_no`);
CREATE INDEX `idx_hr_training_participants_training_id` ON `hr_training_participants` (`training_id`);
CREATE INDEX `idx_hr_training_participants_employee_id` ON `hr_training_participants` (`employee_id`);
CREATE INDEX `idx_hr_recruitments_recruitment_no` ON `hr_recruitments` (`recruitment_no`);
CREATE INDEX `idx_hr_recruitments_position_id` ON `hr_recruitments` (`position_id`);
CREATE INDEX `idx_hr_recruitments_department_id` ON `hr_recruitments` (`department_id`);
CREATE INDEX `idx_hr_resumes_resume_no` ON `hr_resumes` (`resume_no`);
CREATE INDEX `idx_hr_resumes_recruitment_id` ON `hr_resumes` (`recruitment_id`);

-- 系统模块索引
CREATE INDEX `idx_system_users_username` ON `system_users` (`username`);
CREATE INDEX `idx_system_users_employee_id` ON `system_users` (`employee_id`);
CREATE INDEX `idx_system_roles_code` ON `system_roles` (`code`);
CREATE INDEX `idx_system_roles_name` ON `system_roles` (`name`);
CREATE INDEX `idx_system_user_roles_user_id` ON `system_user_roles` (`user_id`);
CREATE INDEX `idx_system_user_roles_role_id` ON `system_user_roles` (`role_id`);
CREATE INDEX `idx_system_permissions_code` ON `system_permissions` (`code`);
CREATE INDEX `idx_system_permissions_module` ON `system_permissions` (`module`);
CREATE INDEX `idx_system_role_permissions_role_id` ON `system_role_permissions` (`role_id`);
CREATE INDEX `idx_system_role_permissions_permission_id` ON `system_role_permissions` (`permission_id`);
CREATE INDEX `idx_system_logs_user_id` ON `system_logs` (`user_id`);
CREATE INDEX `idx_system_logs_created_at` ON `system_logs` (`created_at`);

-- 10. 设置外键检查
SET FOREIGN_KEY_CHECKS = 1;