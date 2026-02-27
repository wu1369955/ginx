package test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wu136995/ginx/internal/services"
)

func TestSearchService_SearchByKeyword(t *testing.T) {
	// 创建搜索服务实例
	service := services.NewSearchService()

	// 测试结果集合
	var testResults []map[string]interface{}

	// 测试用例1：正常的关键词搜索
	t.Run("NormalSearch", func(t *testing.T) {
		req := map[string]interface{}{
			"keyword":   "test",
			"type":      "customer",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.SearchByKeyword(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Contains(t, result, "total")
		assert.Contains(t, result, "page")
		assert.Contains(t, result, "page_size")
		assert.Contains(t, result, "total_pages")
		assert.Contains(t, result, "results")

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "NormalSearch",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例2：不指定搜索类型
	t.Run("NoType", func(t *testing.T) {
		req := map[string]interface{}{
			"keyword":   "test",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.SearchByKeyword(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "NoType",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例3：不指定页码
	t.Run("NoPage", func(t *testing.T) {
		req := map[string]interface{}{
			"keyword":   "test",
			"type":      "customer",
			"page_size": 10,
		}

		result, err := service.SearchByKeyword(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "NoPage",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例4：不指定每页大小
	t.Run("NoPageSize", func(t *testing.T) {
		req := map[string]interface{}{
			"keyword": "test",
			"type":    "customer",
			"page":    1,
		}

		result, err := service.SearchByKeyword(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "NoPageSize",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例5：产品搜索
	t.Run("ProductSearch", func(t *testing.T) {
		req := map[string]interface{}{
			"keyword":   "test",
			"type":      "product",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.SearchByKeyword(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "ProductSearch",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例6：订单搜索
	t.Run("OrderSearch", func(t *testing.T) {
		req := map[string]interface{}{
			"keyword":   "test",
			"type":      "order",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.SearchByKeyword(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "OrderSearch",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例7：所有类型搜索
	t.Run("AllSearch", func(t *testing.T) {
		req := map[string]interface{}{
			"keyword":   "test",
			"type":      "all",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.SearchByKeyword(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "AllSearch",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 将测试结果输出到文件
	testOutput := map[string]interface{}{
		"test_name":    "TestSearchService_SearchByKeyword",
		"test_results": testResults,
	}

	outputFile, err := os.Create("./search_by_keyword_test.json")
	if err != nil {
		t.Errorf("创建测试输出文件失败: %v", err)
		return
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(testOutput)
	if err != nil {
		t.Errorf("编码测试结果失败: %v", err)
		return
	}
	t.Logf("测试结果已输出到 ./search_by_keyword_test.json")
}

func TestSearchService_GetRankingData(t *testing.T) {
	// 创建搜索服务实例
	service := services.NewSearchService()

	// 测试结果集合
	var testResults []map[string]interface{}

	// 测试用例1：客户订单数量排行（降序）
	t.Run("CustomerOrderCountDesc", func(t *testing.T) {
		req := map[string]interface{}{
			"type":      "customer_order_count",
			"sort":      "desc",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.GetRankingData(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Contains(t, result, "total")
		assert.Contains(t, result, "page")
		assert.Contains(t, result, "page_size")
		assert.Contains(t, result, "total_pages")
		assert.Contains(t, result, "results")
		assert.Contains(t, result, "sort")

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "CustomerOrderCountDesc",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例2：客户订单数量排行（升序）
	t.Run("CustomerOrderCountAsc", func(t *testing.T) {
		req := map[string]interface{}{
			"type":      "customer_order_count",
			"sort":      "asc",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.GetRankingData(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "asc", result["sort"])

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "CustomerOrderCountAsc",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例3：产品销量排行
	t.Run("ProductSalesVolume", func(t *testing.T) {
		req := map[string]interface{}{
			"type":      "product_sales_volume",
			"sort":      "desc",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.GetRankingData(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "ProductSalesVolume",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例4：客户销售额排行
	t.Run("CustomerSalesAmount", func(t *testing.T) {
		req := map[string]interface{}{
			"type":      "customer_sales_amount",
			"sort":      "desc",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.GetRankingData(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "CustomerSalesAmount",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例5：不指定排序方向
	t.Run("NoSort", func(t *testing.T) {
		req := map[string]interface{}{
			"type":      "customer_order_count",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.GetRankingData(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "desc", result["sort"])

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "NoSort",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例6：不指定页码
	t.Run("NoPage", func(t *testing.T) {
		req := map[string]interface{}{
			"type":      "customer_order_count",
			"sort":      "desc",
			"page_size": 10,
		}

		result, err := service.GetRankingData(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "NoPage",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例7：不指定每页大小
	t.Run("NoPageSize", func(t *testing.T) {
		req := map[string]interface{}{
			"type": "customer_order_count",
			"sort": "desc",
			"page": 1,
		}

		result, err := service.GetRankingData(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "NoPageSize",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 测试用例8：无效的排行类型
	t.Run("InvalidType", func(t *testing.T) {
		req := map[string]interface{}{
			"type":      "invalid_type",
			"sort":      "desc",
			"page":      1,
			"page_size": 10,
		}

		result, err := service.GetRankingData(req)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// 添加测试结果到集合
		testResults = append(testResults, map[string]interface{}{
			"test_case": "InvalidType",
			"request":   req,
			"result":    result,
			"error":     err,
		})
	})

	// 将测试结果输出到文件
	testOutput := map[string]interface{}{
		"test_name":    "TestSearchService_GetRankingData",
		"test_results": testResults,
	}

	outputFile, err := os.Create("./get_ranking_data_test.json")
	if err != nil {
		t.Errorf("创建测试输出文件失败: %v", err)
		return
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(testOutput)
	if err != nil {
		t.Errorf("编码测试结果失败: %v", err)
		return
	}
	t.Logf("测试结果已输出到 ./get_ranking_data_test.json")
}
