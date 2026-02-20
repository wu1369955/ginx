package services

import (
	"errors"
	"time"

	"github.com/wu136995/ginx/internal/api/schemas"
	"github.com/wu136995/ginx/internal/database"
	"github.com/wu136995/ginx/internal/models"
	"gorm.io/gorm"
)

// CRMService CRM服务接口
type CRMService interface {
	// 客户管理
	GetCustomerList(req map[string]interface{}) ([]schemas.CustomerContactResponse, error)
	GetCustomerDetail(id string) (*schemas.CustomerContactResponse, error)
	CreateCustomer(req schemas.CustomerCreateRequest) (*schemas.CustomerContactResponse, error)
	UpdateCustomer(id string, req schemas.CustomerUpdateRequest) (*schemas.CustomerContactResponse, error)
	DeleteCustomer(id string) error
	GetCustomerContacts(id string) ([]schemas.CustomerContactResponse, error)
	AddCustomerContact(id string, req schemas.CustomerContactCreateRequest) (*schemas.CustomerContactResponse, error)
	GetCustomerHistory(id string) ([]schemas.CustomerContactResponse, error)

	// 销售线索管理
	GetLeadList(req map[string]interface{}) ([]schemas.LeadResponse, error)
	GetLeadDetail(id string) (*schemas.LeadResponse, error)
	CreateLead(req schemas.LeadCreateRequest) (*schemas.LeadResponse, error)
	UpdateLead(id string, req schemas.LeadUpdateRequest) (*schemas.LeadResponse, error)
	DeleteLead(id string) error
	QualifyLead(id string) error
	DisqualifyLead(id string) error

	// 销售机会管理
	GetOpportunityList(req map[string]interface{}) ([]schemas.OpportunityResponse, error)
	GetOpportunityDetail(id string) (*schemas.OpportunityResponse, error)
	CreateOpportunity(req schemas.OpportunityCreateRequest) (*schemas.OpportunityResponse, error)
	UpdateOpportunity(id string, req schemas.OpportunityUpdateRequest) (*schemas.OpportunityResponse, error)
	DeleteOpportunity(id string) error
	CloseOpportunity(id string, req map[string]interface{}) error
	GetOpportunityActivities(id string) ([]schemas.ActivityResponse, error)

	// 营销活动管理
	GetCampaignList(req map[string]interface{}) ([]schemas.CampaignResponse, error)
	GetCampaignDetail(id string) (*schemas.CampaignResponse, error)
	CreateCampaign(req schemas.CampaignCreateRequest) (*schemas.CampaignResponse, error)
	UpdateCampaign(id string, req schemas.CampaignUpdateRequest) (*schemas.CampaignResponse, error)
	DeleteCampaign(id string) error
	LaunchCampaign(id string) error
	EndCampaign(id string) error
	GetCampaignResults(id string) (*schemas.CampaignResponse, error)

	// 客户活动管理
	GetActivityList(req map[string]interface{}) ([]schemas.ActivityResponse, error)
	GetActivityDetail(id string) (*schemas.ActivityResponse, error)
	CreateActivity(req schemas.ActivityCreateRequest) (*schemas.ActivityResponse, error)
	UpdateActivity(id string, req schemas.ActivityUpdateRequest) (*schemas.ActivityResponse, error)
	DeleteActivity(id string) error
	CompleteActivity(id string) error
	CancelActivity(id string) error

	// 服务请求管理
	GetServiceRequestList(req map[string]interface{}) ([]schemas.ServiceRequestResponse, error)
	GetServiceRequestDetail(id string) (*schemas.ServiceRequestResponse, error)
	CreateServiceRequest(req schemas.ServiceRequestCreateRequest) (*schemas.ServiceRequestResponse, error)
	UpdateServiceRequest(id string, req schemas.ServiceRequestUpdateRequest) (*schemas.ServiceRequestResponse, error)
	DeleteServiceRequest(id string) error
	AssignServiceRequest(id string, req map[string]interface{}) error
	ResolveServiceRequest(id string) error
	CloseServiceRequest(id string) error

	// 服务工单管理
	GetServiceTicketList(req map[string]interface{}) ([]schemas.ServiceTicketResponse, error)
	GetServiceTicketDetail(id string) (*schemas.ServiceTicketResponse, error)
	CreateServiceTicket(req schemas.ServiceTicketCreateRequest) (*schemas.ServiceTicketResponse, error)
	UpdateServiceTicket(id string, req schemas.ServiceTicketUpdateRequest) (*schemas.ServiceTicketResponse, error)
	DeleteServiceTicket(id string) error
	AssignServiceTicket(id string, req map[string]interface{}) error
	ResolveServiceTicket(id string) error
	CloseServiceTicket(id string) error

	// CRM报表管理
	GetSalesPipelineReport(req schemas.SalesPipelineReportRequest) (*schemas.SalesPipelineReportResponse, error)
	GetCustomerAnalysisReport(req schemas.CustomerAnalysisReportRequest) (map[string]interface{}, error)
	GetActivitySummaryReport(req map[string]interface{}) (map[string]interface{}, error)
	GetCampaignEffectivenessReport(req map[string]interface{}) (map[string]interface{}, error)
	GetServicePerformanceReport(req map[string]interface{}) (map[string]interface{}, error)
	ExportCRMReport(req map[string]interface{}) ([]byte, error)
}

// crmService CRM服务实现
type crmService struct {
	db *gorm.DB
}

// NewCRMService 创建CRM服务实例
func NewCRMService() CRMService {
	return &crmService{
		db: database.GetDB(),
	}
}

// 客户管理方法
func (s *crmService) GetCustomerList(req map[string]interface{}) ([]schemas.CustomerContactResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户数据
	var accounts []models.CRMAccount
	result := s.db.Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	customers := make([]schemas.CustomerContactResponse, len(accounts))
	for i, account := range accounts {
		customers[i] = schemas.CustomerContactResponse{
			ID:            account.ID,
			Name:          account.Name,
			Type:          account.Type,
			Industry:      account.Industry,
			Region:        account.Region,
			Address:       account.Address,
			ContactPerson: account.ContactPerson,
			Phone:         account.Phone,
			Email:         account.Email,
			Status:        account.Status,
		}
	}

	return customers, nil
}

func (s *crmService) GetCustomerDetail(id string) (*schemas.CustomerContactResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户详情
	var account models.CRMAccount
	result := s.db.First(&account, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	customer := &schemas.CustomerContactResponse{
		ID:            account.ID,
		Name:          account.Name,
		Type:          account.Type,
		Industry:      account.Industry,
		Region:        account.Region,
		Address:       account.Address,
		ContactPerson: account.ContactPerson,
		Phone:         account.Phone,
		Email:         account.Email,
		Status:        account.Status,
	}

	return customer, nil
}

func (s *crmService) CreateCustomer(req schemas.CustomerCreateRequest) (*schemas.CustomerContactResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建客户模型
	account := models.CRMAccount{
		ID:            req.ID,
		AccountNo:     req.AccountNo,
		Name:          req.Name,
		Type:          req.Type,
		Industry:      req.Industry,
		Region:        req.Region,
		Address:       req.Address,
		ContactPerson: req.ContactPerson,
		Phone:         req.Phone,
		Email:         req.Email,
		Status:        "active",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CreatedBy:     req.CreatedBy,
		UpdatedBy:     req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&account)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	customer := &schemas.CustomerContactResponse{
		ID:            account.ID,
		Name:          account.Name,
		Type:          account.Type,
		Industry:      account.Industry,
		Region:        account.Region,
		Address:       account.Address,
		ContactPerson: account.ContactPerson,
		Phone:         account.Phone,
		Email:         account.Email,
		Status:        account.Status,
	}

	return customer, nil
}

func (s *crmService) UpdateCustomer(id string, req schemas.CustomerUpdateRequest) (*schemas.CustomerContactResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户
	var account models.CRMAccount
	result := s.db.First(&account, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	account.Name = req.Name
	account.Type = req.Type
	account.Industry = req.Industry
	account.Region = req.Region
	account.Address = req.Address
	account.ContactPerson = req.ContactPerson
	account.Phone = req.Phone
	account.Email = req.Email
	account.Status = req.Status
	account.UpdatedAt = time.Now()
	account.UpdatedBy = req.UpdatedBy

	// 保存到数据库
	result = s.db.Save(&account)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	customer := &schemas.CustomerContactResponse{
		ID:            account.ID,
		Name:          account.Name,
		Type:          account.Type,
		Industry:      account.Industry,
		Region:        account.Region,
		Address:       account.Address,
		ContactPerson: account.ContactPerson,
		Phone:         account.Phone,
		Email:         account.Email,
		Status:        account.Status,
	}

	return customer, nil
}

func (s *crmService) DeleteCustomer(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除客户
	result := s.db.Delete(&models.CRMAccount{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) GetCustomerContacts(id string) ([]schemas.CustomerContactResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户联系人
	var contacts []models.CRMContact
	result := s.db.Where("account_id = ?", id).Find(&contacts)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	contactList := make([]schemas.CustomerContactResponse, len(contacts))
	for i, contact := range contacts {
		contactList[i] = schemas.CustomerContactResponse{
			ID:            contact.ID,
			Name:          contact.Name,
			ContactPerson: contact.Name,
			Phone:         contact.Phone,
			Email:         contact.Email,
			Position:      contact.Position,
		}
	}

	return contactList, nil
}

func (s *crmService) AddCustomerContact(id string, req schemas.CustomerContactCreateRequest) (*schemas.CustomerContactResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建联系人模型
	contact := models.CRMContact{
		ID:        req.ID,
		AccountID: id,
		Name:      req.Name,
		Position:  req.Position,
		Phone:     req.Phone,
		Email:     req.Email,
		Mobile:    req.Mobile,
		Address:   req.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&contact)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	contactResponse := &schemas.CustomerContactResponse{
		ID:            contact.ID,
		Name:          contact.Name,
		ContactPerson: contact.Name,
		Phone:         contact.Phone,
		Email:         contact.Email,
		Position:      contact.Position,
	}

	return contactResponse, nil
}

func (s *crmService) GetCustomerHistory(id string) ([]schemas.CustomerContactResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户活动历史
	var activities []models.CRMActivity
	result := s.db.Where("account_id = ?", id).Order("created_at DESC").Find(&activities)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	historyList := make([]schemas.CustomerContactResponse, len(activities))
	for i, activity := range activities {
		historyList[i] = schemas.CustomerContactResponse{
			ID:            activity.ID,
			Name:          activity.Subject,
			ContactPerson: activity.Type,
			Status:        activity.Status,
		}
	}

	return historyList, nil
}

// 销售线索管理方法
func (s *crmService) GetLeadList(req map[string]interface{}) ([]schemas.LeadResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取销售线索数据
	var leads []models.CRMLead
	result := s.db.Find(&leads)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	leadList := make([]schemas.LeadResponse, len(leads))
	for i, lead := range leads {
		leadList[i] = schemas.LeadResponse{
			ID:            lead.ID,
			Name:          lead.Name,
			Company:       lead.Company,
			ContactPerson: lead.ContactPerson,
			Phone:         lead.Phone,
			Email:         lead.Email,
			Source:        lead.Source,
			Status:        lead.Status,
			Score:         lead.Score,
			Description:   lead.Description,
		}
	}

	return leadList, nil
}

func (s *crmService) GetLeadDetail(id string) (*schemas.LeadResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取销售线索详情
	var lead models.CRMLead
	result := s.db.First(&lead, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	leadDetail := &schemas.LeadResponse{
		ID:            lead.ID,
		Name:          lead.Name,
		Company:       lead.Company,
		ContactPerson: lead.ContactPerson,
		Phone:         lead.Phone,
		Email:         lead.Email,
		Source:        lead.Source,
		Status:        lead.Status,
		Score:         lead.Score,
		Description:   lead.Description,
	}

	return leadDetail, nil
}

func (s *crmService) CreateLead(req schemas.LeadCreateRequest) (*schemas.LeadResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建销售线索模型
	lead := models.CRMLead{
		ID:            req.ID,
		Name:          req.Name,
		Company:       req.Company,
		ContactPerson: req.ContactPerson,
		Phone:         req.Phone,
		Email:         req.Email,
		Source:        req.Source,
		Status:        "new",
		Score:         req.Score,
		Description:   req.Description,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CreatedBy:     req.CreatedBy,
		UpdatedBy:     req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&lead)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	leadResponse := &schemas.LeadResponse{
		ID:            lead.ID,
		Name:          lead.Name,
		Company:       lead.Company,
		ContactPerson: lead.ContactPerson,
		Phone:         lead.Phone,
		Email:         lead.Email,
		Source:        lead.Source,
		Status:        lead.Status,
		Score:         lead.Score,
		Description:   lead.Description,
	}

	return leadResponse, nil
}

func (s *crmService) UpdateLead(id string, req schemas.LeadUpdateRequest) (*schemas.LeadResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取销售线索
	var lead models.CRMLead
	result := s.db.First(&lead, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	lead.Name = req.Name
	lead.Company = req.Company
	lead.ContactPerson = req.ContactPerson
	lead.Phone = req.Phone
	lead.Email = req.Email
	lead.Source = req.Source
	lead.Status = req.Status
	lead.Score = req.Score
	lead.Description = req.Description
	lead.UpdatedAt = time.Now()
	lead.UpdatedBy = req.UpdatedBy

	// 保存到数据库
	result = s.db.Save(&lead)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	leadResponse := &schemas.LeadResponse{
		ID:            lead.ID,
		Name:          lead.Name,
		Company:       lead.Company,
		ContactPerson: lead.ContactPerson,
		Phone:         lead.Phone,
		Email:         lead.Email,
		Source:        lead.Source,
		Status:        lead.Status,
		Score:         lead.Score,
		Description:   lead.Description,
	}

	return leadResponse, nil
}

func (s *crmService) DeleteLead(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除销售线索
	result := s.db.Delete(&models.CRMLead{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) QualifyLead(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取销售线索
	var lead models.CRMLead
	result := s.db.First(&lead, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为qualified
	lead.Status = "qualified"
	lead.UpdatedAt = time.Now()
	lead.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&lead)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) DisqualifyLead(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取销售线索
	var lead models.CRMLead
	result := s.db.First(&lead, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为disqualified
	lead.Status = "disqualified"
	lead.UpdatedAt = time.Now()
	lead.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&lead)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 销售机会管理方法
func (s *crmService) GetOpportunityList(req map[string]interface{}) ([]schemas.OpportunityResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取销售机会数据
	var opportunities []models.CRMOpportunity
	result := s.db.Find(&opportunities)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	opportunityList := make([]schemas.OpportunityResponse, len(opportunities))
	for i, opportunity := range opportunities {
		opportunityList[i] = schemas.OpportunityResponse{
			ID:              opportunity.ID,
			OpportunityNo:   opportunity.OpportunityNo,
			AccountID:       opportunity.AccountID,
			Name:            opportunity.Name,
			Description:     opportunity.Description,
			Stage:           opportunity.Stage,
			Probability:     opportunity.Probability,
			EstimatedAmount: opportunity.EstimatedAmount,
			ActualAmount:    opportunity.ActualAmount,
			CloseDate:       opportunity.CloseDate,
		}
	}

	return opportunityList, nil
}

func (s *crmService) GetOpportunityDetail(id string) (*schemas.OpportunityResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取销售机会详情
	var opportunity models.CRMOpportunity
	result := s.db.First(&opportunity, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	opportunityDetail := &schemas.OpportunityResponse{
		ID:              opportunity.ID,
		OpportunityNo:   opportunity.OpportunityNo,
		AccountID:       opportunity.AccountID,
		Name:            opportunity.Name,
		Description:     opportunity.Description,
		Stage:           opportunity.Stage,
		Probability:     opportunity.Probability,
		EstimatedAmount: opportunity.EstimatedAmount,
		ActualAmount:    opportunity.ActualAmount,
		CloseDate:       opportunity.CloseDate,
	}

	return opportunityDetail, nil
}

func (s *crmService) CreateOpportunity(req schemas.OpportunityCreateRequest) (*schemas.OpportunityResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建销售机会模型
	opportunity := models.CRMOpportunity{
		ID:              req.ID,
		OpportunityNo:   req.OpportunityNo,
		AccountID:       req.AccountID,
		Name:            req.Name,
		Description:     req.Description,
		Stage:           "prospecting",
		Probability:     req.Probability,
		EstimatedAmount: req.EstimatedAmount,
		CloseDate:       req.CloseDate,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		CreatedBy:       req.CreatedBy,
		UpdatedBy:       req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&opportunity)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	opportunityResponse := &schemas.OpportunityResponse{
		ID:              opportunity.ID,
		OpportunityNo:   opportunity.OpportunityNo,
		AccountID:       opportunity.AccountID,
		Name:            opportunity.Name,
		Description:     opportunity.Description,
		Stage:           opportunity.Stage,
		Probability:     opportunity.Probability,
		EstimatedAmount: opportunity.EstimatedAmount,
		ActualAmount:    opportunity.ActualAmount,
		CloseDate:       opportunity.CloseDate,
	}

	return opportunityResponse, nil
}

func (s *crmService) UpdateOpportunity(id string, req schemas.OpportunityUpdateRequest) (*schemas.OpportunityResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取销售机会
	var opportunity models.CRMOpportunity
	result := s.db.First(&opportunity, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	opportunity.Name = req.Name
	opportunity.Description = req.Description
	opportunity.Stage = req.Stage
	opportunity.Probability = req.Probability
	opportunity.EstimatedAmount = req.EstimatedAmount
	opportunity.ActualAmount = req.ActualAmount
	opportunity.CloseDate = req.CloseDate
	opportunity.UpdatedAt = time.Now()
	opportunity.UpdatedBy = req.UpdatedBy

	// 保存到数据库
	result = s.db.Save(&opportunity)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	opportunityResponse := &schemas.OpportunityResponse{
		ID:              opportunity.ID,
		OpportunityNo:   opportunity.OpportunityNo,
		AccountID:       opportunity.AccountID,
		Name:            opportunity.Name,
		Description:     opportunity.Description,
		Stage:           opportunity.Stage,
		Probability:     opportunity.Probability,
		EstimatedAmount: opportunity.EstimatedAmount,
		ActualAmount:    opportunity.ActualAmount,
		CloseDate:       opportunity.CloseDate,
	}

	return opportunityResponse, nil
}

func (s *crmService) DeleteOpportunity(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除销售机会
	result := s.db.Delete(&models.CRMOpportunity{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) CloseOpportunity(id string, req map[string]interface{}) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取销售机会
	var opportunity models.CRMOpportunity
	result := s.db.First(&opportunity, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为closed
	opportunity.Stage = "closed"
	opportunity.UpdatedAt = time.Now()
	opportunity.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&opportunity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) GetOpportunityActivities(id string) ([]schemas.ActivityResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取销售机会相关的活动
	var activities []models.CRMActivity
	result := s.db.Where("account_id = ?", id).Find(&activities)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	activityList := make([]schemas.ActivityResponse, len(activities))
	for i, activity := range activities {
		activityList[i] = schemas.ActivityResponse{
			ID:          activity.ID,
			AccountID:   activity.AccountID,
			ContactID:   activity.ContactID,
			Type:        activity.Type,
			Subject:     activity.Subject,
			Description: activity.Description,
			Location:    activity.Location,
			Status:      activity.Status,
			Priority:    activity.Priority,
			StartTime:   activity.StartTime,
			EndTime:     activity.EndTime,
			Outcome:     activity.Outcome,
		}
	}

	return activityList, nil
}

// 营销活动管理方法
func (s *crmService) GetCampaignList(req map[string]interface{}) ([]schemas.CampaignResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取营销活动数据
	var campaigns []models.CRMCampaign
	result := s.db.Find(&campaigns)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	campaignList := make([]schemas.CampaignResponse, len(campaigns))
	for i, campaign := range campaigns {
		campaignList[i] = schemas.CampaignResponse{
			ID:             campaign.ID,
			CampaignNo:     campaign.CampaignNo,
			Name:           campaign.Name,
			Type:           campaign.Type,
			Description:    campaign.Description,
			Status:         campaign.Status,
			StartDate:      campaign.StartDate,
			EndDate:        campaign.EndDate,
			Budget:         campaign.Budget,
			ActualCost:     campaign.ActualCost,
			TargetAudience: campaign.TargetAudience,
		}
	}

	return campaignList, nil
}

func (s *crmService) GetCampaignDetail(id string) (*schemas.CampaignResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取营销活动详情
	var campaign models.CRMCampaign
	result := s.db.First(&campaign, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	campaignDetail := &schemas.CampaignResponse{
		ID:             campaign.ID,
		CampaignNo:     campaign.CampaignNo,
		Name:           campaign.Name,
		Type:           campaign.Type,
		Description:    campaign.Description,
		Status:         campaign.Status,
		StartDate:      campaign.StartDate,
		EndDate:        campaign.EndDate,
		Budget:         campaign.Budget,
		ActualCost:     campaign.ActualCost,
		TargetAudience: campaign.TargetAudience,
	}

	return campaignDetail, nil
}

func (s *crmService) CreateCampaign(req schemas.CampaignCreateRequest) (*schemas.CampaignResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建营销活动模型
	campaign := models.CRMCampaign{
		ID:             req.ID,
		CampaignNo:     req.CampaignNo,
		Name:           req.Name,
		Type:           req.Type,
		Description:    req.Description,
		Status:         "planning",
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		Budget:         req.Budget,
		ActualCost:     0,
		TargetAudience: req.TargetAudience,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		CreatedBy:      req.CreatedBy,
		UpdatedBy:      req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&campaign)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	campaignResponse := &schemas.CampaignResponse{
		ID:             campaign.ID,
		CampaignNo:     campaign.CampaignNo,
		Name:           campaign.Name,
		Type:           campaign.Type,
		Description:    campaign.Description,
		Status:         campaign.Status,
		StartDate:      campaign.StartDate,
		EndDate:        campaign.EndDate,
		Budget:         campaign.Budget,
		ActualCost:     campaign.ActualCost,
		TargetAudience: campaign.TargetAudience,
	}

	return campaignResponse, nil
}

func (s *crmService) UpdateCampaign(id string, req schemas.CampaignUpdateRequest) (*schemas.CampaignResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取营销活动
	var campaign models.CRMCampaign
	result := s.db.First(&campaign, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	campaign.Name = req.Name
	campaign.Type = req.Type
	campaign.Description = req.Description
	campaign.Status = req.Status
	campaign.StartDate = req.StartDate
	campaign.EndDate = req.EndDate
	campaign.Budget = req.Budget
	campaign.ActualCost = req.ActualCost
	campaign.TargetAudience = req.TargetAudience
	campaign.UpdatedAt = time.Now()
	campaign.UpdatedBy = req.UpdatedBy

	// 保存到数据库
	result = s.db.Save(&campaign)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	campaignResponse := &schemas.CampaignResponse{
		ID:             campaign.ID,
		CampaignNo:     campaign.CampaignNo,
		Name:           campaign.Name,
		Type:           campaign.Type,
		Description:    campaign.Description,
		Status:         campaign.Status,
		StartDate:      campaign.StartDate,
		EndDate:        campaign.EndDate,
		Budget:         campaign.Budget,
		ActualCost:     campaign.ActualCost,
		TargetAudience: campaign.TargetAudience,
	}

	return campaignResponse, nil
}

func (s *crmService) DeleteCampaign(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除营销活动
	result := s.db.Delete(&models.CRMCampaign{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) LaunchCampaign(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取营销活动
	var campaign models.CRMCampaign
	result := s.db.First(&campaign, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为active
	campaign.Status = "active"
	campaign.UpdatedAt = time.Now()
	campaign.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&campaign)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) EndCampaign(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取营销活动
	var campaign models.CRMCampaign
	result := s.db.First(&campaign, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为completed
	campaign.Status = "completed"
	campaign.UpdatedAt = time.Now()
	campaign.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&campaign)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) GetCampaignResults(id string) (*schemas.CampaignResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取营销活动结果
	var campaign models.CRMCampaign
	result := s.db.First(&campaign, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	campaignResult := &schemas.CampaignResponse{
		ID:             campaign.ID,
		CampaignNo:     campaign.CampaignNo,
		Name:           campaign.Name,
		Type:           campaign.Type,
		Description:    campaign.Description,
		Status:         campaign.Status,
		StartDate:      campaign.StartDate,
		EndDate:        campaign.EndDate,
		Budget:         campaign.Budget,
		ActualCost:     campaign.ActualCost,
		TargetAudience: campaign.TargetAudience,
	}

	return campaignResult, nil
}

// 客户活动管理方法
func (s *crmService) GetActivityList(req map[string]interface{}) ([]schemas.ActivityResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户活动数据
	var activities []models.CRMActivity
	result := s.db.Find(&activities)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	activityList := make([]schemas.ActivityResponse, len(activities))
	for i, activity := range activities {
		activityList[i] = schemas.ActivityResponse{
			ID:          activity.ID,
			AccountID:   activity.AccountID,
			ContactID:   activity.ContactID,
			Type:        activity.Type,
			Subject:     activity.Subject,
			Description: activity.Description,
			Location:    activity.Location,
			Status:      activity.Status,
			Priority:    activity.Priority,
			StartTime:   activity.StartTime,
			EndTime:     activity.EndTime,
			Outcome:     activity.Outcome,
		}
	}

	return activityList, nil
}

func (s *crmService) GetActivityDetail(id string) (*schemas.ActivityResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户活动详情
	var activity models.CRMActivity
	result := s.db.First(&activity, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	activityDetail := &schemas.ActivityResponse{
		ID:          activity.ID,
		AccountID:   activity.AccountID,
		ContactID:   activity.ContactID,
		Type:        activity.Type,
		Subject:     activity.Subject,
		Description: activity.Description,
		Location:    activity.Location,
		Status:      activity.Status,
		Priority:    activity.Priority,
		StartTime:   activity.StartTime,
		EndTime:     activity.EndTime,
		Outcome:     activity.Outcome,
	}

	return activityDetail, nil
}

func (s *crmService) CreateActivity(req schemas.ActivityCreateRequest) (*schemas.ActivityResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建客户活动模型
	activity := models.CRMActivity{
		ID:          req.ID,
		AccountID:   req.AccountID,
		ContactID:   req.ContactID,
		Type:        req.Type,
		Subject:     req.Subject,
		Description: req.Description,
		Location:    req.Location,
		Status:      "planned",
		Priority:    req.Priority,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&activity)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	activityResponse := &schemas.ActivityResponse{
		ID:          activity.ID,
		AccountID:   activity.AccountID,
		ContactID:   activity.ContactID,
		Type:        activity.Type,
		Subject:     activity.Subject,
		Description: activity.Description,
		Location:    activity.Location,
		Status:      activity.Status,
		Priority:    activity.Priority,
		StartTime:   activity.StartTime,
		EndTime:     activity.EndTime,
		Outcome:     activity.Outcome,
	}

	return activityResponse, nil
}

func (s *crmService) UpdateActivity(id string, req schemas.ActivityUpdateRequest) (*schemas.ActivityResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户活动
	var activity models.CRMActivity
	result := s.db.First(&activity, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	activity.Type = req.Type
	activity.Subject = req.Subject
	activity.Description = req.Description
	activity.Location = req.Location
	activity.Status = req.Status
	activity.Priority = req.Priority
	activity.StartTime = req.StartTime
	activity.EndTime = req.EndTime
	activity.Outcome = req.Outcome
	activity.UpdatedAt = time.Now()
	activity.UpdatedBy = req.UpdatedBy

	// 保存到数据库
	result = s.db.Save(&activity)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	activityResponse := &schemas.ActivityResponse{
		ID:          activity.ID,
		AccountID:   activity.AccountID,
		ContactID:   activity.ContactID,
		Type:        activity.Type,
		Subject:     activity.Subject,
		Description: activity.Description,
		Location:    activity.Location,
		Status:      activity.Status,
		Priority:    activity.Priority,
		StartTime:   activity.StartTime,
		EndTime:     activity.EndTime,
		Outcome:     activity.Outcome,
	}

	return activityResponse, nil
}

func (s *crmService) DeleteActivity(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除客户活动
	result := s.db.Delete(&models.CRMActivity{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) CompleteActivity(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取客户活动
	var activity models.CRMActivity
	result := s.db.First(&activity, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为completed
	activity.Status = "completed"
	activity.UpdatedAt = time.Now()
	activity.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&activity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) CancelActivity(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取客户活动
	var activity models.CRMActivity
	result := s.db.First(&activity, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为cancelled
	activity.Status = "cancelled"
	activity.UpdatedAt = time.Now()
	activity.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&activity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 服务请求管理方法
func (s *crmService) GetServiceRequestList(req map[string]interface{}) ([]schemas.ServiceRequestResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取服务请求数据
	var cases []models.CRMCase
	result := s.db.Find(&cases)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	requestList := make([]schemas.ServiceRequestResponse, len(cases))
	for i, caseItem := range cases {
		requestList[i] = schemas.ServiceRequestResponse{
			ID:             caseItem.ID,
			CaseNo:         caseItem.CaseNo,
			AccountID:      caseItem.AccountID,
			ContactID:      caseItem.ContactID,
			Subject:        caseItem.Subject,
			Description:    caseItem.Description,
			Type:           caseItem.Type,
			Priority:       caseItem.Priority,
			Status:         caseItem.Status,
			Resolution:     caseItem.Resolution,
			ResponseTime:   caseItem.ResponseTime,
			ResolutionTime: caseItem.ResolutionTime,
			Satisfaction:   caseItem.Satisfaction,
		}
	}

	return requestList, nil
}

func (s *crmService) GetServiceRequestDetail(id string) (*schemas.ServiceRequestResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取服务请求详情
	var caseItem models.CRMCase
	result := s.db.First(&caseItem, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	requestDetail := &schemas.ServiceRequestResponse{
		ID:             caseItem.ID,
		CaseNo:         caseItem.CaseNo,
		AccountID:      caseItem.AccountID,
		ContactID:      caseItem.ContactID,
		Subject:        caseItem.Subject,
		Description:    caseItem.Description,
		Type:           caseItem.Type,
		Priority:       caseItem.Priority,
		Status:         caseItem.Status,
		Resolution:     caseItem.Resolution,
		ResponseTime:   caseItem.ResponseTime,
		ResolutionTime: caseItem.ResolutionTime,
		Satisfaction:   caseItem.Satisfaction,
	}

	return requestDetail, nil
}

func (s *crmService) CreateServiceRequest(req schemas.ServiceRequestCreateRequest) (*schemas.ServiceRequestResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建服务请求模型
	caseItem := models.CRMCase{
		ID:          req.ID,
		CaseNo:      req.CaseNo,
		AccountID:   req.AccountID,
		ContactID:   req.ContactID,
		Subject:     req.Subject,
		Description: req.Description,
		Type:        req.Type,
		Priority:    req.Priority,
		Status:      "new",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&caseItem)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	requestResponse := &schemas.ServiceRequestResponse{
		ID:             caseItem.ID,
		CaseNo:         caseItem.CaseNo,
		AccountID:      caseItem.AccountID,
		ContactID:      caseItem.ContactID,
		Subject:        caseItem.Subject,
		Description:    caseItem.Description,
		Type:           caseItem.Type,
		Priority:       caseItem.Priority,
		Status:         caseItem.Status,
		Resolution:     caseItem.Resolution,
		ResponseTime:   caseItem.ResponseTime,
		ResolutionTime: caseItem.ResolutionTime,
		Satisfaction:   caseItem.Satisfaction,
	}

	return requestResponse, nil
}

func (s *crmService) UpdateServiceRequest(id string, req schemas.ServiceRequestUpdateRequest) (*schemas.ServiceRequestResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取服务请求
	var caseItem models.CRMCase
	result := s.db.First(&caseItem, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	caseItem.Subject = req.Subject
	caseItem.Description = req.Description
	caseItem.Type = req.Type
	caseItem.Priority = req.Priority
	caseItem.Status = req.Status
	caseItem.Resolution = req.Resolution
	caseItem.UpdatedAt = time.Now()
	caseItem.UpdatedBy = req.UpdatedBy

	// 保存到数据库
	result = s.db.Save(&caseItem)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	requestResponse := &schemas.ServiceRequestResponse{
		ID:             caseItem.ID,
		CaseNo:         caseItem.CaseNo,
		AccountID:      caseItem.AccountID,
		ContactID:      caseItem.ContactID,
		Subject:        caseItem.Subject,
		Description:    caseItem.Description,
		Type:           caseItem.Type,
		Priority:       caseItem.Priority,
		Status:         caseItem.Status,
		Resolution:     caseItem.Resolution,
		ResponseTime:   caseItem.ResponseTime,
		ResolutionTime: caseItem.ResolutionTime,
		Satisfaction:   caseItem.Satisfaction,
	}

	return requestResponse, nil
}

func (s *crmService) DeleteServiceRequest(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除服务请求
	result := s.db.Delete(&models.CRMCase{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) AssignServiceRequest(id string, req map[string]interface{}) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取服务请求
	var caseItem models.CRMCase
	result := s.db.First(&caseItem, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为assigned
	caseItem.Status = "assigned"
	caseItem.UpdatedAt = time.Now()
	caseItem.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&caseItem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) ResolveServiceRequest(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取服务请求
	var caseItem models.CRMCase
	result := s.db.First(&caseItem, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为resolved
	caseItem.Status = "resolved"
	caseItem.ResolutionTime = &time.Time{}
	*caseItem.ResolutionTime = time.Now()
	caseItem.UpdatedAt = time.Now()
	caseItem.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&caseItem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) CloseServiceRequest(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取服务请求
	var caseItem models.CRMCase
	result := s.db.First(&caseItem, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为closed
	caseItem.Status = "closed"
	caseItem.UpdatedAt = time.Now()
	caseItem.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&caseItem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 服务工单管理方法
func (s *crmService) GetServiceTicketList(req map[string]interface{}) ([]schemas.ServiceTicketResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取服务工单数据
	var cases []models.CRMCase
	result := s.db.Find(&cases)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	ticketList := make([]schemas.ServiceTicketResponse, len(cases))
	for i, caseItem := range cases {
		ticketList[i] = schemas.ServiceTicketResponse{
			ID:             caseItem.ID,
			TicketNo:       caseItem.CaseNo,
			AccountID:      caseItem.AccountID,
			ContactID:      caseItem.ContactID,
			Subject:        caseItem.Subject,
			Description:    caseItem.Description,
			Type:           caseItem.Type,
			Priority:       caseItem.Priority,
			Status:         caseItem.Status,
			Resolution:     caseItem.Resolution,
			ResponseTime:   caseItem.ResponseTime,
			ResolutionTime: caseItem.ResolutionTime,
			Satisfaction:   caseItem.Satisfaction,
		}
	}

	return ticketList, nil
}

func (s *crmService) GetServiceTicketDetail(id string) (*schemas.ServiceTicketResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取服务工单详情
	var caseItem models.CRMCase
	result := s.db.First(&caseItem, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	ticketDetail := &schemas.ServiceTicketResponse{
		ID:             caseItem.ID,
		TicketNo:       caseItem.CaseNo,
		AccountID:      caseItem.AccountID,
		ContactID:      caseItem.ContactID,
		Subject:        caseItem.Subject,
		Description:    caseItem.Description,
		Type:           caseItem.Type,
		Priority:       caseItem.Priority,
		Status:         caseItem.Status,
		Resolution:     caseItem.Resolution,
		ResponseTime:   caseItem.ResponseTime,
		ResolutionTime: caseItem.ResolutionTime,
		Satisfaction:   caseItem.Satisfaction,
	}

	return ticketDetail, nil
}

func (s *crmService) CreateServiceTicket(req schemas.ServiceTicketCreateRequest) (*schemas.ServiceTicketResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 创建服务工单模型
	caseItem := models.CRMCase{
		ID:          req.ID,
		CaseNo:      req.TicketNo,
		AccountID:   req.AccountID,
		ContactID:   req.ContactID,
		Subject:     req.Subject,
		Description: req.Description,
		Type:        req.Type,
		Priority:    req.Priority,
		Status:      "new",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.CreatedBy,
	}

	// 保存到数据库
	result := s.db.Create(&caseItem)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	ticketResponse := &schemas.ServiceTicketResponse{
		ID:             caseItem.ID,
		TicketNo:       caseItem.CaseNo,
		AccountID:      caseItem.AccountID,
		ContactID:      caseItem.ContactID,
		Subject:        caseItem.Subject,
		Description:    caseItem.Description,
		Type:           caseItem.Type,
		Priority:       caseItem.Priority,
		Status:         caseItem.Status,
		Resolution:     caseItem.Resolution,
		ResponseTime:   caseItem.ResponseTime,
		ResolutionTime: caseItem.ResolutionTime,
		Satisfaction:   caseItem.Satisfaction,
	}

	return ticketResponse, nil
}

func (s *crmService) UpdateServiceTicket(id string, req schemas.ServiceTicketUpdateRequest) (*schemas.ServiceTicketResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取服务工单
	var caseItem models.CRMCase
	result := s.db.First(&caseItem, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	// 更新字段
	caseItem.Subject = req.Subject
	caseItem.Description = req.Description
	caseItem.Type = req.Type
	caseItem.Priority = req.Priority
	caseItem.Status = req.Status
	caseItem.Resolution = req.Resolution
	caseItem.UpdatedAt = time.Now()
	caseItem.UpdatedBy = req.UpdatedBy

	// 保存到数据库
	result = s.db.Save(&caseItem)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将模型转换为响应格式
	ticketResponse := &schemas.ServiceTicketResponse{
		ID:             caseItem.ID,
		TicketNo:       caseItem.CaseNo,
		AccountID:      caseItem.AccountID,
		ContactID:      caseItem.ContactID,
		Subject:        caseItem.Subject,
		Description:    caseItem.Description,
		Type:           caseItem.Type,
		Priority:       caseItem.Priority,
		Status:         caseItem.Status,
		Resolution:     caseItem.Resolution,
		ResponseTime:   caseItem.ResponseTime,
		ResolutionTime: caseItem.ResolutionTime,
		Satisfaction:   caseItem.Satisfaction,
	}

	return ticketResponse, nil
}

func (s *crmService) DeleteServiceTicket(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库删除服务工单
	result := s.db.Delete(&models.CRMCase{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) AssignServiceTicket(id string, req map[string]interface{}) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取服务工单
	var caseItem models.CRMCase
	result := s.db.First(&caseItem, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为assigned
	caseItem.Status = "assigned"
	caseItem.UpdatedAt = time.Now()
	caseItem.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&caseItem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) ResolveServiceTicket(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取服务工单
	var caseItem models.CRMCase
	result := s.db.First(&caseItem, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为resolved
	caseItem.Status = "resolved"
	if caseItem.ResolutionTime == nil {
		caseItem.ResolutionTime = &time.Time{}
	}
	*caseItem.ResolutionTime = time.Now()
	caseItem.UpdatedAt = time.Now()
	caseItem.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&caseItem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *crmService) CloseServiceTicket(id string) error {
	// 检查数据库连接
	if s.db == nil {
		return errors.New("database connection is nil")
	}

	// 从数据库读取服务工单
	var caseItem models.CRMCase
	result := s.db.First(&caseItem, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	// 更新状态为closed
	caseItem.Status = "closed"
	caseItem.UpdatedAt = time.Now()
	caseItem.UpdatedBy = "system"

	// 保存到数据库
	result = s.db.Save(&caseItem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// CRM报表管理方法
func (s *crmService) GetSalesPipelineReport(req schemas.SalesPipelineReportRequest) (*schemas.SalesPipelineReportResponse, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取销售机会数据
	var opportunities []models.CRMOpportunity
	result := s.db.Find(&opportunities)
	if result.Error != nil {
		return nil, result.Error
	}

	// 构建销售漏斗报表数据
	report := &schemas.SalesPipelineReportResponse{
		TotalOpportunities: len(opportunities),
		TotalValue:         0,
		Stages:             make([]schemas.PipelineStage, 0),
	}

	// 计算总价值和按阶段分组
	stageMap := make(map[string][]models.CRMOpportunity)
	for _, opportunity := range opportunities {
		report.TotalValue += opportunity.EstimatedAmount
		stageMap[opportunity.Stage] = append(stageMap[opportunity.Stage], opportunity)
	}

	// 构建阶段数据
	for stage, stageOpportunities := range stageMap {
		stageValue := 0.0
		for _, opportunity := range stageOpportunities {
			stageValue += opportunity.EstimatedAmount
		}

		report.Stages = append(report.Stages, schemas.PipelineStage{
			Stage: stage,
			Count: len(stageOpportunities),
			Value: stageValue,
		})
	}

	return report, nil
}

func (s *crmService) GetCustomerAnalysisReport(req schemas.CustomerAnalysisReportRequest) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取客户数据
	var accounts []models.CRMAccount
	result := s.db.Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}

	// 构建客户分析报表数据
	report := make(map[string]interface{})
	report["total_customers"] = len(accounts)

	// 按行业分组
	industryMap := make(map[string]int)
	for _, account := range accounts {
		industryMap[account.Industry]++
	}
	report["industry_breakdown"] = industryMap

	// 按状态分组
	statusMap := make(map[string]int)
	for _, account := range accounts {
		statusMap[account.Status]++
	}
	report["status_breakdown"] = statusMap

	return report, nil
}

func (s *crmService) GetActivitySummaryReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取活动数据
	var activities []models.CRMActivity
	result := s.db.Find(&activities)
	if result.Error != nil {
		return nil, result.Error
	}

	// 构建活动汇总报表数据
	report := make(map[string]interface{})
	report["total_activities"] = len(activities)

	// 按类型分组
	typeMap := make(map[string]int)
	for _, activity := range activities {
		typeMap[activity.Type]++
	}
	report["activity_types"] = typeMap

	// 按状态分组
	statusMap := make(map[string]int)
	for _, activity := range activities {
		statusMap[activity.Status]++
	}
	report["status_breakdown"] = statusMap

	return report, nil
}

func (s *crmService) GetCampaignEffectivenessReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取营销活动数据
	var campaigns []models.CRMCampaign
	result := s.db.Find(&campaigns)
	if result.Error != nil {
		return nil, result.Error
	}

	// 构建营销活动效果报表数据
	report := make(map[string]interface{})
	report["total_campaigns"] = len(campaigns)

	// 计算总预算和实际成本
	totalBudget := 0.0
	totalActualCost := 0.0
	for _, campaign := range campaigns {
		totalBudget += campaign.Budget
		totalActualCost += campaign.ActualCost
	}
	report["total_budget"] = totalBudget
	report["total_actual_cost"] = totalActualCost
	report["cost_efficiency"] = totalBudget / (totalActualCost + 1) // 避免除以零

	// 按状态分组
	statusMap := make(map[string]int)
	for _, campaign := range campaigns {
		statusMap[campaign.Status]++
	}
	report["status_breakdown"] = statusMap

	return report, nil
}

func (s *crmService) GetServicePerformanceReport(req map[string]interface{}) (map[string]interface{}, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 从数据库读取服务请求数据
	var cases []models.CRMCase
	result := s.db.Find(&cases)
	if result.Error != nil {
		return nil, result.Error
	}

	// 构建服务绩效报表数据
	report := make(map[string]interface{})
	report["total_service_requests"] = len(cases)

	// 按状态分组
	statusMap := make(map[string]int)
	for _, caseItem := range cases {
		statusMap[caseItem.Status]++
	}
	report["status_breakdown"] = statusMap

	// 按优先级分组
	priorityMap := make(map[string]int)
	for _, caseItem := range cases {
		priorityMap[caseItem.Priority]++
	}
	report["priority_breakdown"] = priorityMap

	return report, nil
}

func (s *crmService) ExportCRMReport(req map[string]interface{}) ([]byte, error) {
	// 检查数据库连接
	if s.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// 这里简单返回一个空字节数组，实际实现中应该根据请求生成相应的报表文件
	return []byte{}, nil
}
