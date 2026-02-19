package schemas

// 生产订单相关结构体

// ProductionOrderCreateRequest 创建生产订单请求
 type ProductionOrderCreateRequest struct {
	Code           string                          `json:"code" binding:"required"`
	OrderDate      string                          `json:"order_date" binding:"required"`
	ProductID      uint                            `json:"product_id" binding:"required"`
	Quantity       float64                         `json:"quantity" binding:"required"`
	ScheduledStartDate string                      `json:"scheduled_start_date" binding:"required"`
	ScheduledEndDate   string                      `json:"scheduled_end_date" binding:"required"`
	Priority       string                          `json:"priority"`
	Source         string                          `json:"source"`
	SourceID       uint                            `json:"source_id"`
	Items          []ProductionOrderItemRequest    `json:"items" binding:"required,dive"`
	Remarks        string                          `json:"remarks"`
}

// ProductionOrderItemRequest 生产订单项目请求
 type ProductionOrderItemRequest struct {
	MaterialID     uint    `json:"material_id" binding:"required"`
	Quantity       float64 `json:"quantity" binding:"required"`
	Unit           string  `json:"unit" binding:"required"`
}

// ProductionOrderUpdateRequest 更新生产订单请求
 type ProductionOrderUpdateRequest struct {
	OrderDate      string                          `json:"order_date"`
	ProductID      uint                            `json:"product_id"`
	Quantity       float64                         `json:"quantity"`
	ScheduledStartDate string                      `json:"scheduled_start_date"`
	ScheduledEndDate   string                      `json:"scheduled_end_date"`
	Priority       string                          `json:"priority"`
	Source         string                          `json:"source"`
	SourceID       uint                            `json:"source_id"`
	Items          []ProductionOrderItemRequest    `json:"items" binding:"omitempty,dive"`
	Remarks        string                          `json:"remarks"`
}

// ProductionOrderResponse 生产订单响应
 type ProductionOrderResponse struct {
	ID             uint                            `json:"id"`
	Code           string                          `json:"code"`
	OrderDate      string                          `json:"order_date"`
	ProductID      uint                            `json:"product_id"`
	ProductCode    string                          `json:"product_code"`
	ProductName    string                          `json:"product_name"`
	Quantity       float64                         `json:"quantity"`
	ProducedQuantity float64                       `json:"produced_quantity"`
	ScheduledStartDate string                      `json:"scheduled_start_date"`
	ScheduledEndDate   string                      `json:"scheduled_end_date"`
	ActualStartDate    string                      `json:"actual_start_date"`
	ActualEndDate      string                      `json:"actual_end_date"`
	Priority       string                          `json:"priority"`
	Status         string                          `json:"status"`
	Source         string                          `json:"source"`
	SourceID       uint                            `json:"source_id"`
	Items          []ProductionOrderItemResponse   `json:"items"`
	Remarks        string                          `json:"remarks"`
	CreatedAt      string                          `json:"created_at"`
	UpdatedAt      string                          `json:"updated_at"`
}

// ProductionOrderItemResponse 生产订单项目响应
 type ProductionOrderItemResponse struct {
	ID             uint    `json:"id"`
	ProductionOrderID uint    `json:"production_order_id"`
	MaterialID     uint    `json:"material_id"`
	MaterialCode   string  `json:"material_code"`
	MaterialName   string  `json:"material_name"`
	Quantity       float64 `json:"quantity"`
	IssuedQuantity float64 `json:"issued_quantity"`
	Unit           string  `json:"unit"`
}

// 生产工单相关结构体

// ProductionTicketCreateRequest 创建生产工单请求
 type ProductionTicketCreateRequest struct {
	Code           string                          `json:"code" binding:"required"`
	ProductionOrderID uint                          `json:"production_order_id" binding:"required"`
	RoutingID      uint                            `json:"routing_id" binding:"required"`
	WorkCenterID   uint                            `json:"work_center_id" binding:"required"`
	OperationID    uint                            `json:"operation_id" binding:"required"`
	Quantity       float64                         `json:"quantity" binding:"required"`
	ScheduledStartDate string                      `json:"scheduled_start_date" binding:"required"`
	ScheduledEndDate   string                      `json:"scheduled_end_date" binding:"required"`
	Status         string                          `json:"status"`
	Remarks        string                          `json:"remarks"`
}

// ProductionTicketUpdateRequest 更新生产工单请求
 type ProductionTicketUpdateRequest struct {
	RoutingID      uint                            `json:"routing_id"`
	WorkCenterID   uint                            `json:"work_center_id"`
	OperationID    uint                            `json:"operation_id"`
	Quantity       float64                         `json:"quantity"`
	ScheduledStartDate string                      `json:"scheduled_start_date"`
	ScheduledEndDate   string                      `json:"scheduled_end_date"`
	Status         string                          `json:"status"`
	Remarks        string                          `json:"remarks"`
}

// ProductionTicketResponse 生产工单响应
 type ProductionTicketResponse struct {
	ID             uint                            `json:"id"`
	Code           string                          `json:"code"`
	ProductionOrderID uint                          `json:"production_order_id"`
	ProductionOrderCode string                      `json:"production_order_code"`
	RoutingID      uint                            `json:"routing_id"`
	RoutingName    string                          `json:"routing_name"`
	WorkCenterID   uint                            `json:"work_center_id"`
	WorkCenterName string                          `json:"work_center_name"`
	OperationID    uint                            `json:"operation_id"`
	OperationName  string                          `json:"operation_name"`
	Quantity       float64                         `json:"quantity"`
	CompletedQuantity float64                      `json:"completed_quantity"`
	ScheduledStartDate string                      `json:"scheduled_start_date"`
	ScheduledEndDate   string                      `json:"scheduled_end_date"`
	ActualStartDate    string                      `json:"actual_start_date"`
	ActualEndDate      string                      `json:"actual_end_date"`
	Status         string                          `json:"status"`
	Remarks        string                          `json:"remarks"`
	CreatedAt      string                          `json:"created_at"`
	UpdatedAt      string                          `json:"updated_at"`
}

// 工艺路线相关结构体

// RoutingCreateRequest 创建工艺路线请求
 type RoutingCreateRequest struct {
	Code           string                          `json:"code" binding:"required"`
	Name           string                          `json:"name" binding:"required"`
	ProductID      uint                            `json:"product_id" binding:"required"`
	Description    string                          `json:"description"`
	Operations     []RoutingOperationRequest       `json:"operations" binding:"required,dive"`
}

// RoutingOperationRequest 工艺路线工序请求
 type RoutingOperationRequest struct {
	Sequence       int     `json:"sequence" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	WorkCenterID   uint    `json:"work_center_id" binding:"required"`
	StandardTime   float64 `json:"standard_time" binding:"required"`
	SetupTime      float64 `json:"setup_time"`
	RunTime        float64 `json:"run_time"`
	Description    string  `json:"description"`
}

// RoutingUpdateRequest 更新工艺路线请求
 type RoutingUpdateRequest struct {
	Code           string                          `json:"code"`
	Name           string                          `json:"name"`
	ProductID      uint                            `json:"product_id"`
	Description    string                          `json:"description"`
	Operations     []RoutingOperationRequest       `json:"operations" binding:"omitempty,dive"`
}

// RoutingResponse 工艺路线响应
 type RoutingResponse struct {
	ID             uint                            `json:"id"`
	Code           string                          `json:"code"`
	Name           string                          `json:"name"`
	ProductID      uint                            `json:"product_id"`
	ProductCode    string                          `json:"product_code"`
	ProductName    string                          `json:"product_name"`
	Description    string                          `json:"description"`
	Operations     []RoutingOperationResponse      `json:"operations"`
	CreatedAt      string                          `json:"created_at"`
	UpdatedAt      string                          `json:"updated_at"`
}

// RoutingOperationResponse 工艺路线工序响应
 type RoutingOperationResponse struct {
	ID             uint    `json:"id"`
	RoutingID      uint    `json:"routing_id"`
	Sequence       int     `json:"sequence"`
	Name           string  `json:"name"`
	WorkCenterID   uint    `json:"work_center_id"`
	WorkCenterName string  `json:"work_center_name"`
	StandardTime   float64 `json:"standard_time"`
	SetupTime      float64 `json:"setup_time"`
	RunTime        float64 `json:"run_time"`
	Description    string  `json:"description"`
}

// 工作中心相关结构体

// WorkCenterCreateRequest 创建工作中心请求
 type WorkCenterCreateRequest struct {
	Code           string  `json:"code" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	Type           string  `json:"type" binding:"required"`
	Capacity       float64 `json:"capacity" binding:"required"`
	Efficiency     float64 `json:"efficiency"`
	Location       string  `json:"location"`
	Description    string  `json:"description"`
	Active         bool    `json:"active"`
}

// WorkCenterUpdateRequest 更新工作中心请求
 type WorkCenterUpdateRequest struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Capacity       float64 `json:"capacity"`
	Efficiency     float64 `json:"efficiency"`
	Location       string  `json:"location"`
	Description    string  `json:"description"`
	Active         bool    `json:"active"`
}

// WorkCenterResponse 工作中心响应
 type WorkCenterResponse struct {
	ID             uint    `json:"id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Capacity       float64 `json:"capacity"`
	Efficiency     float64 `json:"efficiency"`
	AvailableCapacity float64 `json:"available_capacity"`
	LoadPercentage float64 `json:"load_percentage"`
	Location       string  `json:"location"`
	Description    string  `json:"description"`
	Active         bool    `json:"active"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// WorkCenterCapacityResponse 工作中心产能响应
 type WorkCenterCapacityResponse struct {
	WorkCenterID   uint    `json:"work_center_id"`
	WorkCenterName string  `json:"work_center_name"`
	TotalCapacity  float64 `json:"total_capacity"`
	UsedCapacity   float64 `json:"used_capacity"`
	AvailableCapacity float64 `json:"available_capacity"`
	LoadPercentage float64 `json:"load_percentage"`
	Period         string  `json:"period"`
	Details        []WorkCenterCapacityDetail `json:"details"`
}

// WorkCenterCapacityDetail 工作中心产能详情
 type WorkCenterCapacityDetail struct {
	Date           string  `json:"date"`
	TotalCapacity  float64 `json:"total_capacity"`
	UsedCapacity   float64 `json:"used_capacity"`
	AvailableCapacity float64 `json:"available_capacity"`
	LoadPercentage float64 `json:"load_percentage"`
}

// 物料需求计划相关结构体

// MRPRunRequest 运行物料需求计划请求
 type MRPRunRequest struct {
	StartDate      string `json:"start_date" binding:"required"`
	EndDate        string `json:"end_date" binding:"required"`
	PlantID        uint   `json:"plant_id"`
	ProductIDs     []uint `json:"product_ids"`
	Regenerate     bool   `json:"regenerate"`
}

// MRPResultResponse 物料需求计划结果响应
 type MRPResultResponse struct {
	RunID          uint                `json:"run_id"`
	RunDate        string              `json:"run_date"`
	StartDate      string              `json:"start_date"`
	EndDate        string              `json:"end_date"`
	TotalProducts  int                 `json:"total_products"`
	TotalMaterials int                 `json:"total_materials"`
	Suggestions    []MRPSuggestionResponse `json:"suggestions"`
}

// MRPSuggestionResponse 物料需求计划建议响应
 type MRPSuggestionResponse struct {
	ID             uint    `json:"id"`
	MRPRunID       uint    `json:"mrp_run_id"`
	Type           string  `json:"type"` // purchase, production, transfer
	MaterialID     uint    `json:"material_id"`
	MaterialCode   string  `json:"material_code"`
	MaterialName   string  `json:"material_name"`
	Quantity       float64 `json:"quantity"`
	RequiredDate   string  `json:"required_date"`
	CurrentStock   float64 `json:"current_stock"`
	Shortage       float64 `json:"shortage"`
	SuggestedAction string  `json:"suggested_action"`
}

// 生产报表相关结构体

// ProductionPlanReportRequest 生产计划报表请求
 type ProductionPlanReportRequest struct {
	StartDate      string `json:"start_date" binding:"required"`
	EndDate        string `json:"end_date" binding:"required"`
	ProductID      uint   `json:"product_id"`
	WorkCenterID   uint   `json:"work_center_id"`
	Status         string `json:"status"`
}

// ProductionPlanReportResponse 生产计划报表响应
 type ProductionPlanReportResponse struct {
	Period         string                      `json:"period"`
	TotalOrders    int                         `json:"total_orders"`
	TotalQuantity  float64                     `json:"total_quantity"`
	Orders         []ProductionPlanOrderResponse `json:"orders"`
}

// ProductionPlanOrderResponse 生产计划订单响应
 type ProductionPlanOrderResponse struct {
	OrderID        uint    `json:"order_id"`
	OrderCode      string  `json:"order_code"`
	ProductName    string  `json:"product_name"`
	Quantity       float64 `json:"quantity"`
	ScheduledStart string  `json:"scheduled_start"`
	ScheduledEnd   string  `json:"scheduled_end"`
	Status         string  `json:"status"`
	Priority       string  `json:"priority"`
}

// ProductionExecutionReportRequest 生产执行报表请求
 type ProductionExecutionReportRequest struct {
	StartDate      string `json:"start_date" binding:"required"`
	EndDate        string `json:"end_date" binding:"required"`
	ProductID      uint   `json:"product_id"`
	WorkCenterID   uint   `json:"work_center_id"`
}

// ProductionExecutionReportResponse 生产执行报表响应
 type ProductionExecutionReportResponse struct {
	Period         string                        `json:"period"`
	TotalOrders    int                           `json:"total_orders"`
	TotalQuantity  float64                       `json:"total_quantity"`
	ProducedQuantity float64                     `json:"produced_quantity"`
	Efficiency     float64                       `json:"efficiency"`
	Orders         []ProductionExecutionOrderResponse `json:"orders"`
}

// ProductionExecutionOrderResponse 生产执行订单响应
 type ProductionExecutionOrderResponse struct {
	OrderID        uint    `json:"order_id"`
	OrderCode      string  `json:"order_code"`
	ProductName    string  `json:"product_name"`
	Quantity       float64 `json:"quantity"`
	ProducedQuantity float64 `json:"produced_quantity"`
	ActualStart    string  `json:"actual_start"`
	ActualEnd      string  `json:"actual_end"`
	Status         string  `json:"status"`
}

// WorkCenterLoadReportRequest 工作中心负荷报表请求
 type WorkCenterLoadReportRequest struct {
	StartDate      string `json:"start_date" binding:"required"`
	EndDate        string `json:"end_date" binding:"required"`
	WorkCenterID   uint   `json:"work_center_id"`
}

// WorkCenterLoadReportResponse 工作中心负荷报表响应
 type WorkCenterLoadReportResponse struct {
	Period         string                      `json:"period"`
	WorkCenters    []WorkCenterLoadDetail      `json:"work_centers"`
}

// WorkCenterLoadDetail 工作中心负荷详情
 type WorkCenterLoadDetail struct {
	WorkCenterID   uint    `json:"work_center_id"`
	WorkCenterName string  `json:"work_center_name"`
	TotalCapacity  float64 `json:"total_capacity"`
	UsedCapacity   float64 `json:"used_capacity"`
	LoadPercentage float64 `json:"load_percentage"`
	Status         string  `json:"status"` // underloaded, normal, overloaded
}