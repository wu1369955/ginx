package schemas

import "time"

// 客户相关结构体

// CustomerCreateRequest 创建客户请求
type CustomerCreateRequest struct {
	ID          string  `json:"id" binding:"required"`
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
	CreatedBy   string  `json:"createdBy" binding:"required"`
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
	UpdatedBy   string  `json:"updatedBy" binding:"required"`
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
	ID        string `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Position  string `json:"position"`
	Phone     string `json:"phone" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
	CreatedBy string `json:"created_by" binding:"required"`
}

// CustomerContactResponse 客户响应
type CustomerContactResponse struct {
	ID            string `json:"id"`
	AccountNo     string `json:"accountNo"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Industry      string `json:"industry"`
	Region        string `json:"region"`
	Address       string `json:"address"`
	ContactPerson string `json:"contactPerson"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Status        string `json:"status"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	CreatedBy     string `json:"createdBy"`
	UpdatedBy     string `json:"updatedBy"`
}

// ContactResponse 联系人响应
type ContactResponse struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	Name      string    `json:"name"`
	Position  string    `json:"position"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Address   string    `json:"address"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 销售线索相关结构体

// LeadCreateRequest 创建销售线索请求
type LeadCreateRequest struct {
	ID             string  `json:"id" binding:"required"`
	Code           string  `json:"code" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	Company        string  `json:"company"`
	ContactPerson  string  `json:"contactPerson" binding:"required"`
	Phone          string  `json:"phone" binding:"required"`
	Email          string  `json:"email" binding:"required,email"`
	Source         string  `json:"source" binding:"required"`
	Status         string  `json:"status"`
	Score          int     `json:"score"`
	EstimatedValue float64 `json:"estimated_value"`
	Description    string  `json:"description"`
	Remarks        string  `json:"remarks"`
	CreatedBy      string  `json:"createdBy" binding:"required"`
}

// LeadUpdateRequest 更新销售线索请求
type LeadUpdateRequest struct {
	Name           string  `json:"name"`
	Company        string  `json:"company"`
	ContactPerson  string  `json:"contactPerson"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email" binding:"omitempty,email"`
	Source         string  `json:"source"`
	Status         string  `json:"status"`
	Score          int     `json:"score"`
	EstimatedValue float64 `json:"estimated_value"`
	Description    string  `json:"description"`
	Remarks        string  `json:"remarks"`
	UpdatedBy      string  `json:"updatedBy" binding:"required"`
}

// LeadResponse 销售线索响应
type LeadResponse struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	Company       string `json:"company"`
	ContactPerson string `json:"contactPerson"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Source        string `json:"source"`
	Status        string `json:"status"`
	Score         int    `json:"score"`
	Description   string `json:"description"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

// 销售机会相关结构体

// OpportunityCreateRequest 创建销售机会请求
type OpportunityCreateRequest struct {
	ID              string  `json:"id" binding:"required"`
	OpportunityNo   string  `json:"opportunity_no" binding:"required"`
	AccountID       string  `json:"account_id" binding:"required"`
	Name            string  `json:"name" binding:"required"`
	Description     string  `json:"description"`
	Stage           string  `json:"stage" binding:"required"`
	Probability     int     `json:"probability"`
	EstimatedAmount float64 `json:"estimated_amount" binding:"required"`
	CloseDate       string  `json:"close_date" binding:"required"`
	CreatedBy       string  `json:"createdBy" binding:"required"`
}

// OpportunityUpdateRequest 更新销售机会请求
type OpportunityUpdateRequest struct {
	AccountID       string  `json:"account_id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Stage           string  `json:"stage"`
	Probability     int     `json:"probability"`
	EstimatedAmount float64 `json:"estimated_amount"`
	CloseDate       string  `json:"close_date"`
	UpdatedBy       string  `json:"updatedBy" binding:"required"`
}

// OpportunityResponse 销售机会响应
type OpportunityResponse struct {
	ID              string   `json:"id"`
	OpportunityNo   string   `json:"opportunity_no"`
	AccountID       string   `json:"account_id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Stage           string   `json:"stage"`
	Probability     int      `json:"probability"`
	EstimatedAmount float64  `json:"estimated_amount"`
	ActualAmount    *float64 `json:"actual_amount,omitempty"`
	CloseDate       string   `json:"close_date"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	CreatedBy       string   `json:"created_by"`
	UpdatedBy       string   `json:"updated_by"`
}

// 营销活动相关结构体

// CampaignCreateRequest 创建营销活动请求
type CampaignCreateRequest struct {
	ID             string    `json:"id" binding:"required"`
	CampaignNo     string    `json:"campaign_no" binding:"required"`
	Name           string    `json:"name" binding:"required"`
	Type           string    `json:"type" binding:"required"`
	Description    string    `json:"description"`
	StartDate      time.Time `json:"start_date" binding:"required"`
	EndDate        time.Time `json:"end_date" binding:"required"`
	Budget         float64   `json:"budget"`
	TargetAudience string    `json:"target_audience"`
	CreatedBy      string    `json:"created_by" binding:"required"`
}

// CampaignUpdateRequest 更新营销活动请求
type CampaignUpdateRequest struct {
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	Budget         float64   `json:"budget"`
	ActualCost     float64   `json:"actual_cost"`
	TargetAudience string    `json:"target_audience"`
	UpdatedBy      string    `json:"updated_by" binding:"required"`
}

// CampaignResponse 营销活动响应
type CampaignResponse struct {
	ID             string    `json:"id"`
	CampaignNo     string    `json:"campaign_no"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	Budget         float64   `json:"budget"`
	ActualCost     float64   `json:"actual_cost"`
	TargetAudience string    `json:"target_audience"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// 客户活动相关结构体

// ActivityCreateRequest 创建客户活动请求
type ActivityCreateRequest struct {
	ID          string    `json:"id" binding:"required"`
	AccountID   string    `json:"account_id" binding:"required"`
	ContactID   string    `json:"contact_id"`
	Type        string    `json:"type" binding:"required"` // call, meeting, email, task
	Subject     string    `json:"subject" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Priority    string    `json:"priority"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	CreatedBy   string    `json:"created_by" binding:"required"`
}

// ActivityUpdateRequest 更新客户活动请求
type ActivityUpdateRequest struct {
	ContactID   string    `json:"contact_id"`
	Type        string    `json:"type"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Outcome     string    `json:"outcome"`
	UpdatedBy   string    `json:"updated_by" binding:"required"`
}

// ActivityResponse 客户活动响应
type ActivityResponse struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"account_id"`
	ContactID   string    `json:"contact_id"`
	Type        string    `json:"type"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Outcome     string    `json:"outcome"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// 服务请求相关结构体

// ServiceRequestCreateRequest 创建服务请求
type ServiceRequestCreateRequest struct {
	ID          string `json:"id" binding:"required"`
	CaseNo      string `json:"case_no" binding:"required"`
	AccountID   string `json:"account_id" binding:"required"`
	ContactID   string `json:"contact_id"`
	Subject     string `json:"subject" binding:"required"`
	Description string `json:"description" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Priority    string `json:"priority" binding:"required"`
	CreatedBy   string `json:"created_by" binding:"required"`
}

// ServiceRequestUpdateRequest 更新服务请求
type ServiceRequestUpdateRequest struct {
	ContactID   string `json:"contact_id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	Resolution  string `json:"resolution"`
	UpdatedBy   string `json:"updated_by" binding:"required"`
}

// ServiceRequestResponse 服务请求响应
type ServiceRequestResponse struct {
	ID             string     `json:"id"`
	CaseNo         string     `json:"case_no"`
	AccountID      string     `json:"account_id"`
	ContactID      string     `json:"contact_id"`
	Subject        string     `json:"subject"`
	Description    string     `json:"description"`
	Type           string     `json:"type"`
	Priority       string     `json:"priority"`
	Status         string     `json:"status"`
	Resolution     string     `json:"resolution"`
	ResponseTime   *time.Time `json:"response_time"`
	ResolutionTime *time.Time `json:"resolution_time"`
	Satisfaction   *int       `json:"satisfaction"`
	CreatedBy      string     `json:"created_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedBy      string     `json:"updated_by"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// 服务工单相关结构体

// ServiceTicketCreateRequest 创建服务工单
type ServiceTicketCreateRequest struct {
	ID          string `json:"id" binding:"required"`
	TicketNo    string `json:"ticket_no" binding:"required"`
	AccountID   string `json:"account_id" binding:"required"`
	ContactID   string `json:"contact_id"`
	Subject     string `json:"subject" binding:"required"`
	Description string `json:"description" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Priority    string `json:"priority" binding:"required"`
	CreatedBy   string `json:"created_by" binding:"required"`
}

// ServiceTicketUpdateRequest 更新服务工单
type ServiceTicketUpdateRequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	Resolution  string `json:"resolution"`
	UpdatedBy   string `json:"updated_by" binding:"required"`
}

// ServiceTicketResponse 服务工单响应
type ServiceTicketResponse struct {
	ID             string     `json:"id"`
	TicketNo       string     `json:"ticket_no"`
	AccountID      string     `json:"account_id"`
	ContactID      string     `json:"contact_id"`
	Subject        string     `json:"subject"`
	Description    string     `json:"description"`
	Type           string     `json:"type"`
	Priority       string     `json:"priority"`
	Status         string     `json:"status"`
	Resolution     string     `json:"resolution"`
	ResponseTime   *time.Time `json:"response_time"`
	ResolutionTime *time.Time `json:"resolution_time"`
	Satisfaction   *int       `json:"satisfaction"`
	CreatedBy      string     `json:"created_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedBy      string     `json:"updated_by"`
	UpdatedAt      time.Time  `json:"updated_at"`
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
