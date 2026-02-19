package schemas

// 账户相关结构体

// AccountCreateRequest 创建账户请求
type AccountCreateRequest struct {
	Code           string  `json:"code" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	Type           string  `json:"type" binding:"required"`
	Currency       string  `json:"currency" binding:"required"`
	InitialBalance float64 `json:"initial_balance"`
	Description    string  `json:"description"`
	Active         bool    `json:"active"`
}

// AccountUpdateRequest 更新账户请求
type AccountUpdateRequest struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Currency       string  `json:"currency"`
	InitialBalance float64 `json:"initial_balance"`
	Description    string  `json:"description"`
	Active         bool    `json:"active"`
}

// AccountResponse 账户响应
type AccountResponse struct {
	ID             uint    `json:"id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Currency       string  `json:"currency"`
	InitialBalance float64 `json:"initial_balance"`
	CurrentBalance float64 `json:"current_balance"`
	Description    string  `json:"description"`
	Active         bool    `json:"active"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// 凭证相关结构体

// VoucherCreateRequest 创建凭证请求
type VoucherCreateRequest struct {
	Code        string               `json:"code" binding:"required"`
	Date        string               `json:"date" binding:"required"`
	Description string               `json:"description" binding:"required"`
	Items       []VoucherItemRequest `json:"items" binding:"required,dive"`
	Reference   string               `json:"reference"`
	Remarks     string               `json:"remarks"`
}

// VoucherItemRequest 凭证项目请求
type VoucherItemRequest struct {
	AccountID   uint    `json:"account_id" binding:"required"`
	Debit       float64 `json:"debit"`
	Credit      float64 `json:"credit"`
	Description string  `json:"description"`
}

// VoucherUpdateRequest 更新凭证请求
type VoucherUpdateRequest struct {
	Date        string               `json:"date"`
	Description string               `json:"description"`
	Items       []VoucherItemRequest `json:"items" binding:"omitempty,dive"`
	Reference   string               `json:"reference"`
	Remarks     string               `json:"remarks"`
}

// VoucherResponse 凭证响应
type VoucherResponse struct {
	ID          uint                  `json:"id"`
	Code        string                `json:"code"`
	Date        string                `json:"date"`
	Description string                `json:"description"`
	Status      string                `json:"status"`
	TotalDebit  float64               `json:"total_debit"`
	TotalCredit float64               `json:"total_credit"`
	Reference   string                `json:"reference"`
	Items       []VoucherItemResponse `json:"items"`
	Remarks     string                `json:"remarks"`
	CreatedAt   string                `json:"created_at"`
	UpdatedAt   string                `json:"updated_at"`
}

// VoucherItemResponse 凭证项目响应
type VoucherItemResponse struct {
	ID          uint    `json:"id"`
	VoucherID   uint    `json:"voucher_id"`
	AccountID   uint    `json:"account_id"`
	AccountCode string  `json:"account_code"`
	AccountName string  `json:"account_name"`
	Debit       float64 `json:"debit"`
	Credit      float64 `json:"credit"`
	Description string  `json:"description"`
}

// 付款相关结构体

// PaymentCreateRequest 创建付款请求
type PaymentCreateRequest struct {
	Code          string  `json:"code" binding:"required"`
	Date          string  `json:"date" binding:"required"`
	SupplierID    uint    `json:"supplier_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	Currency      string  `json:"currency" binding:"required"`
	ExchangeRate  float64 `json:"exchange_rate"`
	PaymentMethod string  `json:"payment_method" binding:"required"`
	BankAccountID uint    `json:"bank_account_id"`
	Reference     string  `json:"reference"`
	Description   string  `json:"description" binding:"required"`
	Remarks       string  `json:"remarks"`
}

// PaymentUpdateRequest 更新付款请求
type PaymentUpdateRequest struct {
	Date          string  `json:"date"`
	SupplierID    uint    `json:"supplier_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	ExchangeRate  float64 `json:"exchange_rate"`
	PaymentMethod string  `json:"payment_method"`
	BankAccountID uint    `json:"bank_account_id"`
	Reference     string  `json:"reference"`
	Description   string  `json:"description"`
	Remarks       string  `json:"remarks"`
}

// PaymentResponse 付款响应
type PaymentResponse struct {
	ID              uint    `json:"id"`
	Code            string  `json:"code"`
	Date            string  `json:"date"`
	SupplierID      uint    `json:"supplier_id"`
	SupplierName    string  `json:"supplier_name"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	ExchangeRate    float64 `json:"exchange_rate"`
	LocalAmount     float64 `json:"local_amount"`
	PaymentMethod   string  `json:"payment_method"`
	BankAccountID   uint    `json:"bank_account_id"`
	BankAccountName string  `json:"bank_account_name"`
	Status          string  `json:"status"`
	Reference       string  `json:"reference"`
	Description     string  `json:"description"`
	Remarks         string  `json:"remarks"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

// 收款相关结构体

// ReceiptCreateRequest 创建收款请求
//  type ReceiptCreateRequest struct {
// 	Code           string  `json:"code" binding:"required"`
// 	Date           string  `json:"date" binding:"required"`
// 	CustomerID     uint    `json:"customer_id" binding:"required"`
// 	Amount         float64 `json:"amount" binding:"required"`
// 	Currency       string  `json:"currency" binding:"required"`
// 	ExchangeRate   float64 `json:"exchange_rate"`
// 	PaymentMethod  string  `json:"payment_method" binding:"required"`
// 	BankAccountID  uint    `json:"bank_account_id"`
// 	Reference      string  `json:"reference"`
// 	Description    string  `json:"description" binding:"required"`
// 	Remarks        string  `json:"remarks"`
// }

// ReceiptUpdateRequest 更新收款请求
//  type ReceiptUpdateRequest struct {
// 	Date           string  `json:"date"`
// 	CustomerID     uint    `json:"customer_id"`
// 	Amount         float64 `json:"amount"`
// 	Currency       string  `json:"currency"`
// 	ExchangeRate   float64 `json:"exchange_rate"`
// 	PaymentMethod  string  `json:"payment_method"`
// 	BankAccountID  uint    `json:"bank_account_id"`
// 	Reference      string  `json:"reference"`
// 	Description    string  `json:"description"`
// 	Remarks        string  `json:"remarks"`
// }

// ReceiptResponse 收款响应
//  type ReceiptResponse struct {
// 	ID             uint    `json:"id"`
// 	Code           string  `json:"code"`
// 	Date           string  `json:"date"`
// 	CustomerID     uint    `json:"customer_id"`
// 	CustomerName   string  `json:"customer_name"`
// 	Amount         float64 `json:"amount"`
// 	Currency       string  `json:"currency"`
// 	ExchangeRate   float64 `json:"exchange_rate"`
// 	LocalAmount    float64 `json:"local_amount"`
// 	PaymentMethod  string  `json:"payment_method"`
// 	BankAccountID  uint    `json:"bank_account_id"`
// 	BankAccountName string  `json:"bank_account_name"`
// 	Status         string  `json:"status"`
// 	Reference      string  `json:"reference"`
// 	Description    string  `json:"description"`
// 	Remarks        string  `json:"remarks"`
// 	CreatedAt      string  `json:"created_at"`
// 	UpdatedAt      string  `json:"updated_at"`
// }

// 固定资产相关结构体

// AssetCreateRequest 创建固定资产请求
type AssetCreateRequest struct {
	Code               string  `json:"code" binding:"required"`
	Name               string  `json:"name" binding:"required"`
	Category           string  `json:"category" binding:"required"`
	AcquisitionDate    string  `json:"acquisition_date" binding:"required"`
	Cost               float64 `json:"cost" binding:"required"`
	SalvageValue       float64 `json:"salvage_value"`
	UsefulLife         int     `json:"useful_life" binding:"required"`
	DepreciationMethod string  `json:"depreciation_method" binding:"required"`
	Location           string  `json:"location"`
	DepartmentID       uint    `json:"department_id"`
	Description        string  `json:"description"`
}

// AssetUpdateRequest 更新固定资产请求
type AssetUpdateRequest struct {
	Code               string  `json:"code"`
	Name               string  `json:"name"`
	Category           string  `json:"category"`
	AcquisitionDate    string  `json:"acquisition_date"`
	Cost               float64 `json:"cost"`
	SalvageValue       float64 `json:"salvage_value"`
	UsefulLife         int     `json:"useful_life"`
	DepreciationMethod string  `json:"depreciation_method"`
	Location           string  `json:"location"`
	DepartmentID       uint    `json:"department_id"`
	Description        string  `json:"description"`
}

// AssetResponse 固定资产响应
type AssetResponse struct {
	ID                      uint    `json:"id"`
	Code                    string  `json:"code"`
	Name                    string  `json:"name"`
	Category                string  `json:"category"`
	AcquisitionDate         string  `json:"acquisition_date"`
	Cost                    float64 `json:"cost"`
	SalvageValue            float64 `json:"salvage_value"`
	UsefulLife              int     `json:"useful_life"`
	DepreciationMethod      string  `json:"depreciation_method"`
	AccumulatedDepreciation float64 `json:"accumulated_depreciation"`
	NetBookValue            float64 `json:"net_book_value"`
	Location                string  `json:"location"`
	DepartmentID            uint    `json:"department_id"`
	DepartmentName          string  `json:"department_name"`
	Status                  string  `json:"status"`
	Description             string  `json:"description"`
	CreatedAt               string  `json:"created_at"`
	UpdatedAt               string  `json:"updated_at"`
}

// AssetTransferRequest 固定资产转移请求
type AssetTransferRequest struct {
	AssetID          uint   `json:"asset_id" binding:"required"`
	TransferDate     string `json:"transfer_date" binding:"required"`
	FromDepartmentID uint   `json:"from_department_id"`
	ToDepartmentID   uint   `json:"to_department_id" binding:"required"`
	FromLocation     string `json:"from_location"`
	ToLocation       string `json:"to_location" binding:"required"`
	Reason           string `json:"reason" binding:"required"`
}

// AssetDisposeRequest 固定资产处置请求
type AssetDisposeRequest struct {
	AssetID       uint    `json:"asset_id" binding:"required"`
	DisposeDate   string  `json:"dispose_date" binding:"required"`
	DisposeMethod string  `json:"dispose_method" binding:"required"`
	Proceeds      float64 `json:"proceeds"`
	Reason        string  `json:"reason" binding:"required"`
}

// DepreciationRequest 折旧计算请求
type DepreciationRequest struct {
	Period   string `json:"period" binding:"required"`
	AssetIDs []uint `json:"asset_ids"`
	Category string `json:"category"`
}

// 预算相关结构体

// BudgetCreateRequest 创建预算请求
type BudgetCreateRequest struct {
	Code         string              `json:"code" binding:"required"`
	Name         string              `json:"name" binding:"required"`
	Year         int                 `json:"year" binding:"required"`
	DepartmentID uint                `json:"department_id"`
	TotalAmount  float64             `json:"total_amount" binding:"required"`
	Items        []BudgetItemRequest `json:"items" binding:"required,dive"`
	Description  string              `json:"description"`
}

// BudgetItemRequest 预算项目请求
type BudgetItemRequest struct {
	AccountID   uint    `json:"account_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Description string  `json:"description"`
}

// BudgetUpdateRequest 更新预算请求
type BudgetUpdateRequest struct {
	Name         string              `json:"name"`
	DepartmentID uint                `json:"department_id"`
	TotalAmount  float64             `json:"total_amount"`
	Items        []BudgetItemRequest `json:"items" binding:"omitempty,dive"`
	Description  string              `json:"description"`
}

// BudgetResponse 预算响应
type BudgetResponse struct {
	ID              uint                 `json:"id"`
	Code            string               `json:"code"`
	Name            string               `json:"name"`
	Year            int                  `json:"year"`
	DepartmentID    uint                 `json:"department_id"`
	DepartmentName  string               `json:"department_name"`
	TotalAmount     float64              `json:"total_amount"`
	UsedAmount      float64              `json:"used_amount"`
	RemainingAmount float64              `json:"remaining_amount"`
	Status          string               `json:"status"`
	Items           []BudgetItemResponse `json:"items"`
	Description     string               `json:"description"`
	CreatedAt       string               `json:"created_at"`
	UpdatedAt       string               `json:"updated_at"`
}

// BudgetItemResponse 预算项目响应
type BudgetItemResponse struct {
	ID              uint    `json:"id"`
	BudgetID        uint    `json:"budget_id"`
	AccountID       uint    `json:"account_id"`
	AccountCode     string  `json:"account_code"`
	AccountName     string  `json:"account_name"`
	Amount          float64 `json:"amount"`
	UsedAmount      float64 `json:"used_amount"`
	RemainingAmount float64 `json:"remaining_amount"`
	Description     string  `json:"description"`
}

// 财务报表相关结构体

// BalanceSheetRequest 资产负债表请求
type BalanceSheetRequest struct {
	Date           string `json:"date" binding:"required"`
	ComparisonDate string `json:"comparison_date"`
}

// BalanceSheetResponse 资产负债表响应
type BalanceSheetResponse struct {
	Date             string                  `json:"date"`
	Assets           []BalanceSheetItem      `json:"assets"`
	TotalAssets      float64                 `json:"total_assets"`
	Liabilities      []BalanceSheetItem      `json:"liabilities"`
	TotalLiabilities float64                 `json:"total_liabilities"`
	Equity           []BalanceSheetItem      `json:"equity"`
	TotalEquity      float64                 `json:"total_equity"`
	Comparison       *BalanceSheetComparison `json:"comparison,omitempty"`
}

// BalanceSheetItem 资产负债表项目
type BalanceSheetItem struct {
	AccountCode string  `json:"account_code"`
	AccountName string  `json:"account_name"`
	Amount      float64 `json:"amount"`
	Percentage  float64 `json:"percentage"`
}

// BalanceSheetComparison 资产负债表比较
type BalanceSheetComparison struct {
	Date             string  `json:"date"`
	TotalAssets      float64 `json:"total_assets"`
	TotalLiabilities float64 `json:"total_liabilities"`
	TotalEquity      float64 `json:"total_equity"`
	AssetChange      float64 `json:"asset_change"`
	LiabilityChange  float64 `json:"liability_change"`
	EquityChange     float64 `json:"equity_change"`
}

// IncomeStatementRequest 利润表请求
type IncomeStatementRequest struct {
	StartDate           string `json:"start_date" binding:"required"`
	EndDate             string `json:"end_date" binding:"required"`
	ComparisonStartDate string `json:"comparison_start_date"`
	ComparisonEndDate   string `json:"comparison_end_date"`
}

// IncomeStatementResponse 利润表响应
type IncomeStatementResponse struct {
	Period        string                     `json:"period"`
	Revenues      []IncomeStatementItem      `json:"revenues"`
	TotalRevenues float64                    `json:"total_revenues"`
	Expenses      []IncomeStatementItem      `json:"expenses"`
	TotalExpenses float64                    `json:"total_expenses"`
	NetIncome     float64                    `json:"net_income"`
	Comparison    *IncomeStatementComparison `json:"comparison,omitempty"`
}

// IncomeStatementItem 利润表项目
type IncomeStatementItem struct {
	AccountCode string  `json:"account_code"`
	AccountName string  `json:"account_name"`
	Amount      float64 `json:"amount"`
	Percentage  float64 `json:"percentage"`
}

// IncomeStatementComparison 利润表比较
type IncomeStatementComparison struct {
	Period        string  `json:"period"`
	TotalRevenues float64 `json:"total_revenues"`
	TotalExpenses float64 `json:"total_expenses"`
	NetIncome     float64 `json:"net_income"`
	RevenueChange float64 `json:"revenue_change"`
	ExpenseChange float64 `json:"expense_change"`
	IncomeChange  float64 `json:"income_change"`
}

// CashFlowStatementRequest 现金流量表请求
type CashFlowStatementRequest struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

// CashFlowStatementResponse 现金流量表响应
type CashFlowStatementResponse struct {
	Period              string         `json:"period"`
	OperatingActivities []CashFlowItem `json:"operating_activities"`
	TotalOperating      float64        `json:"total_operating"`
	InvestingActivities []CashFlowItem `json:"investing_activities"`
	TotalInvesting      float64        `json:"total_investing"`
	FinancingActivities []CashFlowItem `json:"financing_activities"`
	TotalFinancing      float64        `json:"total_financing"`
	NetChangeCash       float64        `json:"net_change_cash"`
	BeginningCash       float64        `json:"beginning_cash"`
	EndingCash          float64        `json:"ending_cash"`
}

// CashFlowItem 现金流量表项目
type CashFlowItem struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}
