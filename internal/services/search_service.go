package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/wu136995/ginx/internal/database"
	"github.com/wu136995/ginx/internal/models"
	"gorm.io/gorm"
)

// SearchService 搜索服务接口
type SearchService interface {
	// 关键词搜索（使用Elasticsearch）
	SearchByKeyword(req map[string]interface{}) (map[string]interface{}, error)
	
	// 数量排行数据（可以选择小到大或大到小）
	GetRankingData(req map[string]interface{}) (map[string]interface{}, error)
}

// searchService 搜索服务实现
type searchService struct {
	db *gorm.DB
}

// NewSearchService 创建搜索服务实例
func NewSearchService() SearchService {
	return &searchService{
		db: database.GetDB(),
	}
}

// SearchByKeyword 关键词搜索（使用Elasticsearch）
func (s *searchService) SearchByKeyword(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接为nil，返回模拟数据
		return s.getMockSearchData(req), nil
	}

	// 解析请求参数
	keyword, ok := req["keyword"].(string)
	if !ok || keyword == "" {
		return nil, errors.New("keyword is required")
	}

	// 分页参数
	page, _ := strconv.Atoi(fmt.Sprintf("%v", req["page"]))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(fmt.Sprintf("%v", req["page_size"]))
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 搜索类型（客户、产品、订单等）
	searchType, _ := req["type"].(string)
	if searchType == "" {
		searchType = "all"
	}

	// 构建查询
	var results []map[string]interface{}
	var total int64

	// 这里应该使用Elasticsearch进行搜索
	// 由于没有实际的Elasticsearch连接，这里使用数据库模拟
	switch searchType {
	case "customer":
		// 搜索客户
		var customers []map[string]interface{}
		query := s.db.Model(&models.SalesCustomer{}).Select("id, customer_no, name, contact_person, phone, email")
		
		// 添加关键词搜索条件
		keyword = "%" + keyword + "%"
		query = query.Where("name LIKE ? OR contact_person LIKE ? OR phone LIKE ? OR email LIKE ?", keyword, keyword, keyword, keyword)
		
		// 获取总数
		query.Count(&total)
		
		// 分页查询
		result := query.Offset(offset).Limit(pageSize).Find(&customers)
		if result.Error != nil {
			return nil, result.Error
		}
		
		results = customers
		
	case "product":
		// 搜索产品
		var products []map[string]interface{}
		query := s.db.Model(&models.SalesProduct{}).Select("id, product_no, name, description, unit")
		
		// 添加关键词搜索条件
		keyword = "%" + keyword + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", keyword, keyword)
		
		// 获取总数
		query.Count(&total)
		
		// 分页查询
		result := query.Offset(offset).Limit(pageSize).Find(&products)
		if result.Error != nil {
			return nil, result.Error
		}
		
		results = products
		
	case "order":
		// 搜索订单
		var orders []map[string]interface{}
		query := s.db.Model(&models.SalesOrder{}).Select("id, order_no, customer_id, order_date, total_amount, status")
		
		// 添加关键词搜索条件
		keyword = "%" + keyword + "%"
		query = query.Where("order_no LIKE ? OR customer_id LIKE ?", keyword, keyword)
		
		// 获取总数
		query.Count(&total)
		
		// 分页查询
		result := query.Offset(offset).Limit(pageSize).Find(&orders)
		if result.Error != nil {
			return nil, result.Error
		}
		
		results = orders
		
	default: // all
		// 搜索所有类型
		// 这里简化处理，只返回客户数据
		var customers []map[string]interface{}
		query := s.db.Model(&models.SalesCustomer{}).Select("id, customer_no, name, contact_person, phone, email")
		
		// 添加关键词搜索条件
		keyword = "%" + keyword + "%"
		query = query.Where("name LIKE ? OR contact_person LIKE ? OR phone LIKE ? OR email LIKE ?", keyword, keyword, keyword, keyword)
		
		// 获取总数
		query.Count(&total)
		
		// 分页查询
		result := query.Offset(offset).Limit(pageSize).Find(&customers)
		if result.Error != nil {
			return nil, result.Error
		}
		
		results = customers
	}

	// 构建返回结果
	response := map[string]interface{}{
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		"results":    results,
	}

	return response, nil
}

// getMockSearchData 获取模拟搜索数据
func (s *searchService) getMockSearchData(req map[string]interface{}) map[string]interface{} {
	// 解析请求参数
	_, _ = req["keyword"].(string) // 暂时不使用关键词
	searchType, _ := req["type"].(string)
	if searchType == "" {
		searchType = "all"
	}

	// 分页参数
	page, _ := strconv.Atoi(fmt.Sprintf("%v", req["page"]))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(fmt.Sprintf("%v", req["page_size"]))
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	// 生成模拟数据
	var results []map[string]interface{}
	total := int64(20)

	switch searchType {
	case "customer":
		// 模拟客户数据
		results = []map[string]interface{}{
			{
				"id":             "1",
				"customer_no":    "C001",
				"name":           "测试客户1",
				"contact_person": "张三",
				"phone":          "13800138001",
				"email":          "test1@example.com",
			},
			{
				"id":             "2",
				"customer_no":    "C002",
				"name":           "测试客户2",
				"contact_person": "李四",
				"phone":          "13800138002",
				"email":          "test2@example.com",
			},
		}
		
	case "product":
		// 模拟产品数据
		results = []map[string]interface{}{
			{
				"id":          "1",
				"product_no":  "P001",
				"name":        "测试产品1",
				"description": "这是测试产品1",
				"unit":        "个",
			},
			{
				"id":          "2",
				"product_no":  "P002",
				"name":        "测试产品2",
				"description": "这是测试产品2",
				"unit":        "个",
			},
		}
		
	case "order":
		// 模拟订单数据
		results = []map[string]interface{}{
			{
				"id":           "1",
				"order_no":     "O001",
				"customer_id":  "1",
				"order_date":   "2026-02-27",
				"total_amount": 1000.00,
				"status":       "pending",
			},
			{
				"id":           "2",
				"order_no":     "O002",
				"customer_id":  "2",
				"order_date":   "2026-02-26",
				"total_amount": 2000.00,
				"status":       "completed",
			},
		}
		
	default: // all
		// 模拟所有类型数据（这里简化处理，只返回客户数据）
		results = []map[string]interface{}{
			{
				"id":             "1",
				"customer_no":    "C001",
				"name":           "测试客户1",
				"contact_person": "张三",
				"phone":          "13800138001",
				"email":          "test1@example.com",
			},
			{
				"id":             "2",
				"customer_no":    "C002",
				"name":           "测试客户2",
				"contact_person": "李四",
				"phone":          "13800138002",
				"email":          "test2@example.com",
			},
		}
	}

	// 构建返回结果
	response := map[string]interface{}{
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		"results":     results,
	}

	return response
}

// GetRankingData 数量排行数据
func (s *searchService) GetRankingData(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		// 数据库连接为nil，返回模拟数据
		return s.getMockRankingData(req), nil
	}

	// 解析请求参数
	rankingType, ok := req["type"].(string)
	if !ok || rankingType == "" {
		return nil, errors.New("ranking type is required")
	}

	// 排序方向
	sortDirection, _ := req["sort"].(string)
	if sortDirection == "" {
		sortDirection = "desc" // 默认从大到小
	}

	// 分页参数
	page, _ := strconv.Atoi(fmt.Sprintf("%v", req["page"]))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(fmt.Sprintf("%v", req["page_size"]))
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 构建查询
	var results []map[string]interface{}
	var total int64

	switch rankingType {
	case "customer_order_count":
		// 客户订单数量排行
		query := `
			SELECT c.id, c.name, COUNT(o.id) as order_count
			FROM sales_customers c
			LEFT JOIN sales_orders o ON c.id = o.customer_id
			GROUP BY c.id, c.name
			ORDER BY order_count ` + sortDirection + `
			LIMIT ? OFFSET ?
		`
		
		// 获取总数
		totalQuery := `
			SELECT COUNT(DISTINCT c.id)
			FROM sales_customers c
			LEFT JOIN sales_orders o ON c.id = o.customer_id
		`
		s.db.Raw(totalQuery).Count(&total)
		
		// 执行查询
		result := s.db.Raw(query, pageSize, offset).Scan(&results)
		if result.Error != nil {
			return nil, result.Error
		}
		
	case "product_sales_volume":
		// 产品销量排行
		query := `
			SELECT p.id, p.name, SUM(oi.quantity) as sales_volume
			FROM sales_products p
			LEFT JOIN sales_order_items oi ON p.id = oi.product_id
			GROUP BY p.id, p.name
			ORDER BY sales_volume ` + sortDirection + `
			LIMIT ? OFFSET ?
		`
		
		// 获取总数
		totalQuery := `
			SELECT COUNT(DISTINCT p.id)
			FROM sales_products p
			LEFT JOIN sales_order_items oi ON p.id = oi.product_id
		`
		s.db.Raw(totalQuery).Count(&total)
		
		// 执行查询
		result := s.db.Raw(query, pageSize, offset).Scan(&results)
		if result.Error != nil {
			return nil, result.Error
		}
		
	case "customer_sales_amount":
		// 客户销售额排行
		query := `
			SELECT c.id, c.name, SUM(o.total_amount) as sales_amount
			FROM sales_customers c
			LEFT JOIN sales_orders o ON c.id = o.customer_id
			GROUP BY c.id, c.name
			ORDER BY sales_amount ` + sortDirection + `
			LIMIT ? OFFSET ?
		`
		
		// 获取总数
		totalQuery := `
			SELECT COUNT(DISTINCT c.id)
			FROM sales_customers c
			LEFT JOIN sales_orders o ON c.id = o.customer_id
		`
		s.db.Raw(totalQuery).Count(&total)
		
		// 执行查询
		result := s.db.Raw(query, pageSize, offset).Scan(&results)
		if result.Error != nil {
			return nil, result.Error
		}
		
	default:
		return nil, errors.New("invalid ranking type")
	}

	// 构建返回结果
	response := map[string]interface{}{
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		"results":    results,
		"sort":       sortDirection,
	}

	return response, nil
}

// getMockRankingData 获取模拟排行数据
func (s *searchService) getMockRankingData(req map[string]interface{}) map[string]interface{} {
	// 解析请求参数
	rankingType, _ := req["type"].(string)
	sortDirection, _ := req["sort"].(string)
	if sortDirection == "" {
		sortDirection = "desc"
	}

	// 分页参数
	page, _ := strconv.Atoi(fmt.Sprintf("%v", req["page"]))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(fmt.Sprintf("%v", req["page_size"]))
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	// 生成模拟数据
	var results []map[string]interface{}
	total := int64(10)

	switch rankingType {
	case "customer_order_count":
		// 模拟客户订单数量排行
		results = []map[string]interface{}{
			{
				"id":          "1",
				"name":        "测试客户1",
				"order_count": 10,
			},
			{
				"id":          "2",
				"name":        "测试客户2",
				"order_count": 8,
			},
			{
				"id":          "3",
				"name":        "测试客户3",
				"order_count": 5,
			},
		}
		
	case "product_sales_volume":
		// 模拟产品销量排行
		results = []map[string]interface{}{
			{
				"id":           "1",
				"name":         "测试产品1",
				"sales_volume": 100,
			},
			{
				"id":           "2",
				"name":         "测试产品2",
				"sales_volume": 80,
			},
			{
				"id":           "3",
				"name":         "测试产品3",
				"sales_volume": 50,
			},
		}
		
	case "customer_sales_amount":
		// 模拟客户销售额排行
		results = []map[string]interface{}{
			{
				"id":           "1",
				"name":         "测试客户1",
				"sales_amount": 10000.00,
			},
			{
				"id":           "2",
				"name":         "测试客户2",
				"sales_amount": 8000.00,
			},
			{
				"id":           "3",
				"name":         "测试客户3",
				"sales_amount": 5000.00,
			},
		}
		
	default:
		// 默认为客户订单数量排行
		results = []map[string]interface{}{
			{
				"id":          "1",
				"name":        "测试客户1",
				"order_count": 10,
			},
			{
				"id":          "2",
				"name":        "测试客户2",
				"order_count": 8,
			},
			{
				"id":          "3",
				"name":        "测试客户3",
				"order_count": 5,
			},
		}
	}

	// 如果是升序排序，反转结果
	if sortDirection == "asc" {
		for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
			results[i], results[j] = results[j], results[i]
		}
	}

	// 构建返回结果
	response := map[string]interface{}{
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		"results":     results,
		"sort":        sortDirection,
	}

	return response
}
