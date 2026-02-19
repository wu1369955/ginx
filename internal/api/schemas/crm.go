package schemas

// 客户相关结构体

// CustomerCreateRequest 创建客户请求
type CustomerCreateRequest struct {
	Code        string  `json:"code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Type        string  `json:"type" binding:"required"` // individual, company
	ContactName string  `json:"contact_name" binding:"required"`
	Phone       string  `json:"phone" binding:"required"`
	Email       string  `json:"email" binding:"required,email"`
	Address     string  `json:"address"`
	TaxNumber   string  `json:"tax_number"`
	BankName    string  `json:"bank_name"`
	BankAccount string  `json:"bank_account"`
	Category    string  `json:"category"`
	Region      string  `json:"region"`
	CreditLimit float64 `json:"credit_limit"`
	CreditDays  int     `json:"credit_days"`
	Status      string  `json:"status"`
	Remarks     string  `json:"remarks"`
}

// CustomerUpdateRequest 更新客户请求
type CustomerUpdateRequest struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	ContactName string  `json:"contact_name"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email" binding:"omitempty,email"`
	Address     string  `json:"address"`
	TaxNumber   string  `json:"tax_number"`
	BankName    string  `json:"bank_name"`
	BankAccount string  `json:"bank_account"`
	Category    string  `json:"category"`
	Region      string  `json:"region"`
	CreditLimit float64 `json:"credit_limit"`
	CreditDays  int     `json:"credit_days"`
	Status      string  `json:"status"`
	Remarks     string  `json:"remarks"`
}

// CustomerResponse 客户响应
// type CustomerResponse struct {
// 	ID          uint    `json:"id"`
// 	Code        string  `json:"code"`
// 	Name        string  `json:"name"`
// 	Type        string  `json:"type"`
// 	ContactName string  `json:"contact_name"`
// 	Phone       string  `json:"phone"`
// 	Email       string  `json:"email"`
// 	Address     string  `json:"address"`
// 	TaxNumber   string  `json:"tax_number"`
// 	BankName    string  `json:"bank_name"`
// 	BankAccount string  `json:"bank_account"`
// 	Category    string  `json:"category"`
// 	Region      string  `json:"region"`
// 	CreditLimit float64 `json:"credit_limit"`
// 	CreditDays  int     `json:"credit_days"`
// 	Balance     float64 `json:"balance"`
// 	Status      string  `json:"status"`
// 	Remarks     string  `json:"remarks"`
// 	CreatedAt   string  `json:"created_at"`
// 	UpdatedAt   string  `json:"updated_at"`
// }

// CustomerContactCreateRequest 创建客户联系人请求
type CustomerContactCreateRequest struct {
	CustomerID uint   `json:"customer_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Position   string `json:"position"`
	Phone      string `json:"phone" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	IsPrimary  bool   `json:"is_primary"`
	Remarks    string `json:"remarks"`
}

// CustomerContactResponse 客户联系人响应
type CustomerContactResponse struct {
	ID           uint   `json:"id"`
	CustomerID   uint   `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	IsPrimary    bool   `json:"is_primary"`
	Remarks      string `json:"remarks"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// 销售线索相关结构体

// LeadCreateRequest 创建销售线索请求
type LeadCreateRequest struct {
	Code           string  `json:"code" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	Company        string  `json:"company"`
	Phone          string  `json:"phone" binding:"required"`
	Email          string  `json:"email" binding:"required,email"`
	Source         string  `json:"source" binding:"required"`
	Status         string  `json:"status"`
	Rating         int     `json:"rating"`
	EstimatedValue float64 `json:"estimated_value"`
	Description    string  `json:"description"`
	Remarks        string  `json:"remarks"`
}

// LeadUpdateRequest 更新销售线索请求
type LeadUpdateRequest struct {
	Name           string  `json:"name"`
	Company        string  `json:"company"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email" binding:"omitempty,email"`
	Source         string  `json:"source"`
	Status         string  `json:"status"`
	Rating         int     `json:"rating"`
	EstimatedValue float64 `json:"estimated_value"`
	Description    string  `json:"description"`
	Remarks        string  `json:"remarks"`
}

// LeadResponse 销售线索响应
type LeadResponse struct {
	ID             uint    `json:"id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Company        string  `json:"company"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email"`
	Source         string  `json:"source"`
	Status         string  `json:"status"`
	Rating         int     `json:"rating"`
	EstimatedValue float64 `json:"estimated_value"`
	Description    string  `json:"description"`
	Remarks        string  `json:"remarks"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// 销售机会相关结构体

// OpportunityCreateRequest 创建销售机会请求
type OpportunityCreateRequest struct {
	Code           string  `json:"code" binding:"required"`
	CustomerID     uint    `json:"customer_id" binding:"required"`
	LeadID         uint    `json:"lead_id"`
	Name           string  `json:"name" binding:"required"`
	Stage          string  `json:"stage" binding:"required"`
	Probability    float64 `json:"probability"`
	EstimatedValue float64 `json:"estimated_value" binding:"required"`
	CloseDate      string  `json:"close_date" binding:"required"`
	Description    string  `json:"description"`
	Remarks        string  `json:"remarks"`
}

// OpportunityUpdateRequest 更新销售机会请求
type OpportunityUpdateRequest struct {
	CustomerID     uint    `json:"customer_id"`
	LeadID         uint    `json:"lead_id"`
	Name           string  `json:"name"`
	Stage          string  `json:"stage"`
	Probability    float64 `json:"probability"`
	EstimatedValue float64 `json:"estimated_value"`
	CloseDate      string  `json:"close_date"`
	Description    string  `json:"description"`
	Remarks        string  `json:"remarks"`
}

// OpportunityResponse 销售机会响应
type OpportunityResponse struct {
	ID             uint    `json:"id"`
	Code           string  `json:"code"`
	CustomerID     uint    `json:"customer_id"`
	CustomerName   string  `json:"customer_name"`
	LeadID         uint    `json:"lead_id"`
	LeadName       string  `json:"lead_name"`
	Name           string  `json:"name"`
	Stage          string  `json:"stage"`
	Probability    float64 `json:"probability"`
	EstimatedValue float64 `json:"estimated_value"`
	ActualValue    float64 `json:"actual_value"`
	CloseDate      string  `json:"close_date"`
	Status         string  `json:"status"`
	Description    string  `json:"description"`
	Remarks        string  `json:"remarks"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// 营销活动相关结构体

// CampaignCreateRequest 创建营销活动请求
type CampaignCreateRequest struct {
	Code             string  `json:"code" binding:"required"`
	Name             string  `json:"name" binding:"required"`
	Type             string  `json:"type" binding:"required"`
	Status           string  `json:"status"`
	StartDate        string  `json:"start_date" binding:"required"`
	EndDate          string  `json:"end_date" binding:"required"`
	Budget           float64 `json:"budget"`
	ExpectedResponse float64 `json:"expected_response"`
	Description      string  `json:"description"`
	Remarks          string  `json:"remarks"`
}

// CampaignUpdateRequest 更新营销活动请求
type CampaignUpdateRequest struct {
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	Status           string  `json:"status"`
	StartDate        string  `json:"start_date"`
	EndDate          string  `json:"end_date"`
	Budget           float64 `json:"budget"`
	ExpectedResponse float64 `json:"expected_response"`
	Description      string  `json:"description"`
	Remarks          string  `json:"remarks"`
}

// CampaignResponse 营销活动响应
type CampaignResponse struct {
	ID               uint    `json:"id"`
	Code             string  `json:"code"`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	Status           string  `json:"status"`
	StartDate        string  `json:"start_date"`
	EndDate          string  `json:"end_date"`
	Budget           float64 `json:"budget"`
	ActualCost       float64 `json:"actual_cost"`
	ExpectedResponse float64 `json:"expected_response"`
	ActualResponse   float64 `json:"actual_response"`
	Description      string  `json:"description"`
	Remarks          string  `json:"remarks"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

// 客户活动相关结构体

// ActivityCreateRequest 创建客户活动请求
type ActivityCreateRequest struct {
	Code          string `json:"code" binding:"required"`
	Type          string `json:"type" binding:"required"` // call, meeting, email, task
	CustomerID    uint   `json:"customer_id"`
	LeadID        uint   `json:"lead_id"`
	OpportunityID uint   `json:"opportunity_id"`
	Subject       string `json:"subject" binding:"required"`
	Description   string `json:"description"`
	StartDate     string `json:"start_date" binding:"required"`
	EndDate       string `json:"end_date" binding:"required"`
	Status        string `json:"status"`
	Priority      string `json:"priority"`
	AssignedTo    uint   `json:"assigned_to"`
	Remarks       string `json:"remarks"`
}

// ActivityUpdateRequest 更新客户活动请求
type ActivityUpdateRequest struct {
	Type          string `json:"type"`
	CustomerID    uint   `json:"customer_id"`
	LeadID        uint   `json:"lead_id"`
	OpportunityID uint   `json:"opportunity_id"`
	Subject       string `json:"subject"`
	Description   string `json:"description"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Status        string `json:"status"`
	Priority      string `json:"priority"`
	AssignedTo    uint   `json:"assigned_to"`
	Remarks       string `json:"remarks"`
}

// ActivityResponse 客户活动响应
type ActivityResponse struct {
	ID              uint   `json:"id"`
	Code            string `json:"code"`
	Type            string `json:"type"`
	CustomerID      uint   `json:"customer_id"`
	CustomerName    string `json:"customer_name"`
	LeadID          uint   `json:"lead_id"`
	LeadName        string `json:"lead_name"`
	OpportunityID   uint   `json:"opportunity_id"`
	OpportunityName string `json:"opportunity_name"`
	Subject         string `json:"subject"`
	Description     string `json:"description"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	Status          string `json:"status"`
	Priority        string `json:"priority"`
	AssignedTo      uint   `json:"assigned_to"`
	AssignedName    string `json:"assigned_name"`
	Remarks         string `json:"remarks"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// 服务请求相关结构体

// ServiceRequestCreateRequest 创建服务请求
type ServiceRequestCreateRequest struct {
	Code        string `json:"code" binding:"required"`
	CustomerID  uint   `json:"customer_id" binding:"required"`
	ContactID   uint   `json:"contact_id"`
	Subject     string `json:"subject" binding:"required"`
	Description string `json:"description" binding:"required"`
	Priority    string `json:"priority" binding:"required"`
	Status      string `json:"status"`
	AssignedTo  uint   `json:"assigned_to"`
	DueDate     string `json:"due_date"`
	Remarks     string `json:"remarks"`
}

// ServiceRequestUpdateRequest 更新服务请求
type ServiceRequestUpdateRequest struct {
	CustomerID  uint   `json:"customer_id"`
	ContactID   uint   `json:"contact_id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	AssignedTo  uint   `json:"assigned_to"`
	DueDate     string `json:"due_date"`
	Remarks     string `json:"remarks"`
}

// ServiceRequestResponse 服务请求响应
type ServiceRequestResponse struct {
	ID           uint   `json:"id"`
	Code         string `json:"code"`
	CustomerID   uint   `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	ContactID    uint   `json:"contact_id"`
	ContactName  string `json:"contact_name"`
	Subject      string `json:"subject"`
	Description  string `json:"description"`
	Priority     string `json:"priority"`
	Status       string `json:"status"`
	AssignedTo   uint   `json:"assigned_to"`
	AssignedName string `json:"assigned_name"`
	DueDate      string `json:"due_date"`
	Resolution   string `json:"resolution"`
	Remarks      string `json:"remarks"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// 服务工单相关结构体

// ServiceTicketCreateRequest 创建服务工单
type ServiceTicketCreateRequest struct {
	Code           string  `json:"code" binding:"required"`
	RequestID      uint    `json:"request_id"`
	CustomerID     uint    `json:"customer_id" binding:"required"`
	Subject        string  `json:"subject" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	Type           string  `json:"type" binding:"required"`
	Priority       string  `json:"priority" binding:"required"`
	Status         string  `json:"status"`
	AssignedTo     uint    `json:"assigned_to"`
	DueDate        string  `json:"due_date"`
	EstimatedHours float64 `json:"estimated_hours"`
	Remarks        string  `json:"remarks"`
}

// ServiceTicketUpdateRequest 更新服务工单
type ServiceTicketUpdateRequest struct {
	RequestID      uint    `json:"request_id"`
	CustomerID     uint    `json:"customer_id"`
	Subject        string  `json:"subject"`
	Description    string  `json:"description"`
	Type           string  `json:"type"`
	Priority       string  `json:"priority"`
	Status         string  `json:"status"`
	AssignedTo     uint    `json:"assigned_to"`
	DueDate        string  `json:"due_date"`
	EstimatedHours float64 `json:"estimated_hours"`
	Remarks        string  `json:"remarks"`
}

// ServiceTicketResponse 服务工单响应
type ServiceTicketResponse struct {
	ID             uint    `json:"id"`
	Code           string  `json:"code"`
	RequestID      uint    `json:"request_id"`
	RequestCode    string  `json:"request_code"`
	CustomerID     uint    `json:"customer_id"`
	CustomerName   string  `json:"customer_name"`
	Subject        string  `json:"subject"`
	Description    string  `json:"description"`
	Type           string  `json:"type"`
	Priority       string  `json:"priority"`
	Status         string  `json:"status"`
	AssignedTo     uint    `json:"assigned_to"`
	AssignedName   string  `json:"assigned_name"`
	DueDate        string  `json:"due_date"`
	EstimatedHours float64 `json:"estimated_hours"`
	ActualHours    float64 `json:"actual_hours"`
	Resolution     string  `json:"resolution"`
	Remarks        string  `json:"remarks"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// CRM报表相关结构体

// SalesPipelineReportRequest 销售漏斗报表请求
type SalesPipelineReportRequest struct {
	StartDate     string `json:"start_date" binding:"required"`
	EndDate       string `json:"end_date" binding:"required"`
	SalespersonID uint   `json:"salesperson_id"`
	CustomerID    uint   `json:"customer_id"`
}

// SalesPipelineReportResponse 销售漏斗报表响应
type SalesPipelineReportResponse struct {
	Period             string               `json:"period"`
	TotalOpportunities int                  `json:"total_opportunities"`
	TotalValue         float64              `json:"total_value"`
	Stages             []SalesPipelineStage `json:"stages"`
}

// SalesPipelineStage 销售漏斗阶段
type SalesPipelineStage struct {
	Stage      string  `json:"stage"`
	Count      int     `json:"count"`
	Value      float64 `json:"value"`
	Percentage float64 `json:"percentage"`
}

// CustomerAnalysisReportRequest 客户分析报表请求
type CustomerAnalysisReportRequest struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	Category  string `json:"category"`
	Region    string `json:"region"`
}

// CustomerAnalysisReportResponse 客户分析报表响应
//  type CustomerAnalysisReportResponse struct {
// 	Period         string                  `json:"period"`
// 	TotalCustomers int                     `json:"total_customers"`
// 	TotalRevenue   float64                 `json:"total_revenue"`
// 	AverageRevenue float64                 `json:"average_revenue"`
// 	TopCustomers   []TopCustomer           `json:"top_customers"`
// 	CategoryStats  []CustomerCategoryStat  `json:"category_stats"`
// 	RegionStats    []CustomerRegionStat    `json:"region_stats"`
// }

// TopCustomer 顶级客户
type TopCustomer struct {
	CustomerID   uint    `json:"customer_id"`
	CustomerName string  `json:"customer_name"`
	Revenue      float64 `json:"revenue"`
	Percentage   float64 `json:"percentage"`
}

// CustomerCategoryStat 客户类别统计
type CustomerCategoryStat struct {
	Category   string  `json:"category"`
	Count      int     `json:"count"`
	Revenue    float64 `json:"revenue"`
	Percentage float64 `json:"percentage"`
}

// CustomerRegionStat 客户地区统计
type CustomerRegionStat struct {
	Region     string  `json:"region"`
	Count      int     `json:"count"`
	Revenue    float64 `json:"revenue"`
	Percentage float64 `json:"percentage"`
}
