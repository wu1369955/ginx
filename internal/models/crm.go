package models

import (
	"time"

	"gorm.io/gorm"
)

// CRMAccount 客户表模型
type CRMAccount struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	AccountNo     string         `json:"account_no" gorm:"unique;not null;type:varchar(20)"`
	Name          string         `json:"name" gorm:"not null;type:varchar(100)"`
	Type          string         `json:"type" gorm:"not null;type:varchar(20)"`
	Industry      string         `json:"industry" gorm:"type:varchar(50)"`
	Region        string         `json:"region" gorm:"type:varchar(50)"`
	Address       string         `json:"address" gorm:"type:varchar(255)"`
	ContactPerson string         `json:"contact_person" gorm:"type:varchar(50)"`
	Phone         string         `json:"phone" gorm:"type:varchar(20)"`
	Email         string         `json:"email" gorm:"type:varchar(100)"`
	Website       string         `json:"website" gorm:"type:varchar(255)"`
	TaxNo         string         `json:"tax_no" gorm:"type:varchar(30)"`
	CreditLimit   float64        `json:"credit_limit" gorm:"type:decimal(18,2);default:0"`
	CreditUsed    float64        `json:"credit_used" gorm:"type:decimal(18,2);default:0"`
	LoyaltyLevel  string         `json:"loyalty_level" gorm:"type:varchar(20);default:'bronze'"`
	Remarks       string         `json:"remarks" gorm:"type:text"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'active'"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Contacts   []CRMContact   `json:"contacts,omitempty" gorm:"foreignKey:AccountID"`
	Activities []CRMActivity  `json:"activities,omitempty" gorm:"foreignKey:AccountID"`
	Cases      []CRMCase      `json:"cases,omitempty" gorm:"foreignKey:AccountID"`
	Opportunities []CRMOpportunity `json:"opportunities,omitempty" gorm:"foreignKey:AccountID"`
}

// TableName 指定表名
func (CRMAccount) TableName() string {
	return "crm_accounts"
}

// CRMContact 联系人表模型
type CRMContact struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	AccountID string         `json:"account_id" gorm:"not null;type:varchar(36)"`
	Name      string         `json:"name" gorm:"not null;type:varchar(50)"`
	Position  string         `json:"position" gorm:"type:varchar(50)"`
	Phone     string         `json:"phone" gorm:"type:varchar(20)"`
	Email     string         `json:"email" gorm:"type:varchar(100)"`
	Mobile    string         `json:"mobile" gorm:"type:varchar(20)"`
	Fax       string         `json:"fax" gorm:"type:varchar(20)"`
	Address   string         `json:"address" gorm:"type:varchar(255)"`
	Birthday  *time.Time     `json:"birthday" gorm:"type:date"`
	Preferences string       `json:"preferences" gorm:"type:text"`
	CreatedBy string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Account    CRMAccount    `json:"account,omitempty" gorm:"foreignKey:AccountID"`
	Activities []CRMActivity `json:"activities,omitempty" gorm:"foreignKey:ContactID"`
	Cases      []CRMCase     `json:"cases,omitempty" gorm:"foreignKey:ContactID"`
}

// TableName 指定表名
func (CRMContact) TableName() string {
	return "crm_contacts"
}

// CRMActivity 客户活动表模型
type CRMActivity struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	AccountID   string         `json:"account_id" gorm:"not null;type:varchar(36)"`
	ContactID   string         `json:"contact_id" gorm:"type:varchar(36)"`
	Type        string         `json:"type" gorm:"not null;type:varchar(20)"`
	Subject     string         `json:"subject" gorm:"not null;type:varchar(200)"`
	Description string         `json:"description" gorm:"type:text"`
	Location    string         `json:"location" gorm:"type:varchar(255)"`
	Status      string         `json:"status" gorm:"type:varchar(20);default:'planned'"`
	Priority    string         `json:"priority" gorm:"type:varchar(20);default:'medium'"`
	StartTime   time.Time      `json:"start_time" gorm:"not null"`
	EndTime     time.Time      `json:"end_time" gorm:"not null"`
	Outcome     string         `json:"outcome" gorm:"type:text"`
	CreatedBy   string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy   string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Account CRMAccount  `json:"account,omitempty" gorm:"foreignKey:AccountID"`
	Contact *CRMContact `json:"contact,omitempty" gorm:"foreignKey:ContactID"`
}

// TableName 指定表名
func (CRMActivity) TableName() string {
	return "crm_activities"
}

// CRMCase 客户反馈表模型
type CRMCase struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CaseNo        string         `json:"case_no" gorm:"unique;not null;type:varchar(20)"`
	AccountID     string         `json:"account_id" gorm:"not null;type:varchar(36)"`
	ContactID     string         `json:"contact_id" gorm:"type:varchar(36)"`
	Subject       string         `json:"subject" gorm:"not null;type:varchar(200)"`
	Description   string         `json:"description" gorm:"not null;type:text"`
	Type          string         `json:"type" gorm:"not null;type:varchar(20)"`
	Priority      string         `json:"priority" gorm:"type:varchar(20);default:'medium'"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'new'"`
	Resolution    string         `json:"resolution" gorm:"type:text"`
	ResponseTime  *time.Time     `json:"response_time" gorm:"type:datetime"`
	ResolutionTime *time.Time    `json:"resolution_time" gorm:"type:datetime"`
	Satisfaction  *int           `json:"satisfaction" gorm:"type:tinyint(1)"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Account  CRMAccount      `json:"account,omitempty" gorm:"foreignKey:AccountID"`
	Contact  *CRMContact     `json:"contact,omitempty" gorm:"foreignKey:ContactID"`
	Comments []CRMCaseComment `json:"comments,omitempty" gorm:"foreignKey:CaseID"`
}

// TableName 指定表名
func (CRMCase) TableName() string {
	return "crm_cases"
}

// CRMCaseComment 客户反馈评论表模型
type CRMCaseComment struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CaseID    string         `json:"case_id" gorm:"not null;type:varchar(36)"`
	Content   string         `json:"content" gorm:"not null;type:text"`
	CreatedBy string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Case CRMCase `json:"case,omitempty" gorm:"foreignKey:CaseID"`
}

// TableName 指定表名
func (CRMCaseComment) TableName() string {
	return "crm_case_comments"
}

// CRMLead 销售线索表模型
type CRMLead struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name          string         `json:"name" gorm:"not null;type:varchar(100)"`
	Company       string         `json:"company" gorm:"type:varchar(100)"`
	ContactPerson string         `json:"contact_person" gorm:"type:varchar(50)"`
	Phone         string         `json:"phone" gorm:"type:varchar(20)"`
	Email         string         `json:"email" gorm:"type:varchar(100)"`
	Source        string         `json:"source" gorm:"type:varchar(50)"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'new'"`
	Score         int            `json:"score" gorm:"type:tinyint(3);default:0"`
	Description   string         `json:"description" gorm:"type:text"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName 指定表名
func (CRMLead) TableName() string {
	return "crm_leads"
}

// CRMOpportunity 销售机会表模型
type CRMOpportunity struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	OpportunityNo string         `json:"opportunity_no" gorm:"unique;not null;type:varchar(20)"`
	AccountID     string         `json:"account_id" gorm:"not null;type:varchar(36)"`
	Name          string         `json:"name" gorm:"not null;type:varchar(200)"`
	Description   string         `json:"description" gorm:"type:text"`
	Stage         string         `json:"stage" gorm:"type:varchar(50);default:'prospecting'"`
	Probability   int            `json:"probability" gorm:"type:tinyint(3);default:0"`
	EstimatedAmount float64      `json:"estimated_amount" gorm:"type:decimal(18,2);default:0"`
	ActualAmount  *float64       `json:"actual_amount" gorm:"type:decimal(18,2)"`
	CloseDate     time.Time      `json:"close_date" gorm:"not null;type:date"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// 关联
	Account CRMAccount `json:"account,omitempty" gorm:"foreignKey:AccountID"`
}

// TableName 指定表名
func (CRMOpportunity) TableName() string {
	return "crm_opportunities"
}

// CRMCampaign 营销活动表模型
type CRMCampaign struct {
	ID            string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CampaignNo    string         `json:"campaign_no" gorm:"unique;not null;type:varchar(20)"`
	Name          string         `json:"name" gorm:"not null;type:varchar(200)"`
	Type          string         `json:"type" gorm:"not null;type:varchar(50)"`
	Description   string         `json:"description" gorm:"type:text"`
	Status        string         `json:"status" gorm:"type:varchar(20);default:'planning'"`
	StartDate     time.Time      `json:"start_date" gorm:"not null;type:date"`
	EndDate       time.Time      `json:"end_date" gorm:"not null;type:date"`
	Budget        float64        `json:"budget" gorm:"type:decimal(18,2);default:0"`
	ActualCost    float64        `json:"actual_cost" gorm:"type:decimal(18,2);default:0"`
	TargetAudience string        `json:"target_audience" gorm:"type:varchar(255)"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(36)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(36)"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName 指定表名
func (CRMCampaign) TableName() string {
	return "crm_campaigns"
}
