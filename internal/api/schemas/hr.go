package schemas

// 员工相关结构体

// EmployeeCreateRequest 创建员工请求
 type EmployeeCreateRequest struct {
	EmployeeID     string `json:"employee_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Gender         string `json:"gender" binding:"required"`
	BirthDate      string `json:"birth_date" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Address        string `json:"address"`
	DepartmentID   uint   `json:"department_id" binding:"required"`
	PositionID     uint   `json:"position_id" binding:"required"`
	JoinDate       string `json:"join_date" binding:"required"`
	ProbationEndDate string `json:"probation_end_date"`
	ContractType   string `json:"contract_type"`
	ContractEndDate string `json:"contract_end_date"`
	BasicSalary    float64 `json:"basic_salary" binding:"required"`
	Status         string `json:"status"`
	Remarks        string `json:"remarks"`
}

// EmployeeUpdateRequest 更新员工请求
 type EmployeeUpdateRequest struct {
	Name           string `json:"name"`
	Gender         string `json:"gender"`
	BirthDate      string `json:"birth_date"`
	Phone          string `json:"phone"`
	Email          string `json:"email" binding:"omitempty,email"`
	Address        string `json:"address"`
	DepartmentID   uint   `json:"department_id"`
	PositionID     uint   `json:"position_id"`
	JoinDate       string `json:"join_date"`
	ProbationEndDate string `json:"probation_end_date"`
	ContractType   string `json:"contract_type"`
	ContractEndDate string `json:"contract_end_date"`
	BasicSalary    float64 `json:"basic_salary"`
	Status         string `json:"status"`
	Remarks        string `json:"remarks"`
}

// EmployeeResponse 员工响应
 type EmployeeResponse struct {
	ID             uint    `json:"id"`
	EmployeeID     string  `json:"employee_id"`
	Name           string  `json:"name"`
	Gender         string  `json:"gender"`
	BirthDate      string  `json:"birth_date"`
	Age            int     `json:"age"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email"`
	Address        string  `json:"address"`
	DepartmentID   uint    `json:"department_id"`
	DepartmentName string  `json:"department_name"`
	PositionID     uint    `json:"position_id"`
	PositionName   string  `json:"position_name"`
	JoinDate       string  `json:"join_date"`
	Tenure         float64 `json:"tenure"`
	ProbationEndDate string  `json:"probation_end_date"`
	ContractType   string  `json:"contract_type"`
	ContractEndDate string  `json:"contract_end_date"`
	BasicSalary    float64 `json:"basic_salary"`
	Status         string  `json:"status"`
	Remarks        string  `json:"remarks"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// 部门相关结构体

// DepartmentCreateRequest 创建部门请求
 type DepartmentCreateRequest struct {
	Code           string `json:"code" binding:"required"`
	Name           string `json:"name" binding:"required"`
	ParentID       uint   `json:"parent_id"`
	ManagerID      uint   `json:"manager_id"`
	Description    string `json:"description"`
	Active         bool   `json:"active"`
}

// DepartmentUpdateRequest 更新部门请求
 type DepartmentUpdateRequest struct {
	Code           string `json:"code"`
	Name           string `json:"name"`
	ParentID       uint   `json:"parent_id"`
	ManagerID      uint   `json:"manager_id"`
	Description    string `json:"description"`
	Active         bool   `json:"active"`
}

// DepartmentResponse 部门响应
 type DepartmentResponse struct {
	ID             uint    `json:"id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	ParentID       uint    `json:"parent_id"`
	ParentName     string  `json:"parent_name"`
	ManagerID      uint    `json:"manager_id"`
	ManagerName    string  `json:"manager_name"`
	EmployeeCount  int     `json:"employee_count"`
	Description    string  `json:"description"`
	Active         bool    `json:"active"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// 职位相关结构体

// PositionCreateRequest 创建职位请求
 type PositionCreateRequest struct {
	Code           string `json:"code" binding:"required"`
	Name           string `json:"name" binding:"required"`
	DepartmentID   uint   `json:"department_id" binding:"required"`
	Level          string `json:"level"`
	Description    string `json:"description"`
	Active         bool   `json:"active"`
}

// PositionUpdateRequest 更新职位请求
 type PositionUpdateRequest struct {
	Code           string `json:"code"`
	Name           string `json:"name"`
	DepartmentID   uint   `json:"department_id"`
	Level          string `json:"level"`
	Description    string `json:"description"`
	Active         bool   `json:"active"`
}

// PositionResponse 职位响应
 type PositionResponse struct {
	ID             uint    `json:"id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	DepartmentID   uint    `json:"department_id"`
	DepartmentName string  `json:"department_name"`
	Level          string  `json:"level"`
	EmployeeCount  int     `json:"employee_count"`
	Description    string  `json:"description"`
	Active         bool    `json:"active"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// 考勤相关结构体

// AttendanceCreateRequest 创建考勤请求
 type AttendanceCreateRequest struct {
	EmployeeID     uint   `json:"employee_id" binding:"required"`
	Date           string `json:"date" binding:"required"`
	CheckInTime    string `json:"check_in_time" binding:"required"`
	CheckOutTime   string `json:"check_out_time"`
	Status         string `json:"status" binding:"required"` // present, absent, late, leave
	HoursWorked    float64 `json:"hours_worked"`
	OvertimeHours  float64 `json:"overtime_hours"`
	Remarks        string  `json:"remarks"`
}

// AttendanceUpdateRequest 更新考勤请求
 type AttendanceUpdateRequest struct {
	CheckInTime    string `json:"check_in_time"`
	CheckOutTime   string `json:"check_out_time"`
	Status         string `json:"status"`
	HoursWorked    float64 `json:"hours_worked"`
	OvertimeHours  float64 `json:"overtime_hours"`
	Remarks        string  `json:"remarks"`
}

// AttendanceResponse 考勤响应
 type AttendanceResponse struct {
	ID             uint    `json:"id"`
	EmployeeID     uint    `json:"employee_id"`
	EmployeeName   string  `json:"employee_name"`
	Date           string  `json:"date"`
	CheckInTime    string  `json:"check_in_time"`
	CheckOutTime   string  `json:"check_out_time"`
	Status         string  `json:"status"`
	HoursWorked    float64 `json:"hours_worked"`
	OvertimeHours  float64 `json:"overtime_hours"`
	Remarks        string  `json:"remarks"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// AttendanceReportRequest 考勤报表请求
 type AttendanceReportRequest struct {
	StartDate      string `json:"start_date" binding:"required"`
	EndDate        string `json:"end_date" binding:"required"`
	EmployeeID     uint   `json:"employee_id"`
	DepartmentID   uint   `json:"department_id"`
	Status         string `json:"status"`
}

// AttendanceReportResponse 考勤报表响应
 type AttendanceReportResponse struct {
	Period         string                    `json:"period"`
	TotalDays      int                       `json:"total_days"`
	Employees      []EmployeeAttendanceReport `json:"employees"`
}

// EmployeeAttendanceReport 员工考勤报表
 type EmployeeAttendanceReport struct {
	EmployeeID     uint    `json:"employee_id"`
	EmployeeName   string  `json:"employee_name"`
	DepartmentName string  `json:"department_name"`
	PresentDays    int     `json:"present_days"`
	AbsentDays     int     `json:"absent_days"`
	LateDays       int     `json:"late_days"`
	LeaveDays      int     `json:"leave_days"`
	TotalHours     float64 `json:"total_hours"`
	OvertimeHours  float64 `json:"overtime_hours"`
	AttendanceRate float64 `json:"attendance_rate"`
}

// 薪资相关结构体

// SalaryCreateRequest 创建薪资请求
 type SalaryCreateRequest struct {
	EmployeeID     uint    `json:"employee_id" binding:"required"`
	PayPeriod      string  `json:"pay_period" binding:"required"`
	BasicSalary    float64 `json:"basic_salary" binding:"required"`
	Allowances     float64 `json:"allowances"`
	Deductions     float64 `json:"deductions"`
	OvertimePay    float64 `json:"overtime_pay"`
	Bonus          float64 `json:"bonus"`
	Tax            float64 `json:"tax"`
	NetSalary      float64 `json:"net_salary" binding:"required"`
	PaymentDate    string  `json:"payment_date"`
	Status         string  `json:"status"`
	Remarks        string  `json:"remarks"`
}

// SalaryUpdateRequest 更新薪资请求
 type SalaryUpdateRequest struct {
	PayPeriod      string  `json:"pay_period"`
	BasicSalary    float64 `json:"basic_salary"`
	Allowances     float64 `json:"allowances"`
	Deductions     float64 `json:"deductions"`
	OvertimePay    float64 `json:"overtime_pay"`
	Bonus          float64 `json:"bonus"`
	Tax            float64 `json:"tax"`
	NetSalary      float64 `json:"net_salary"`
	PaymentDate    string  `json:"payment_date"`
	Status         string  `json:"status"`
	Remarks        string  `json:"remarks"`
}

// SalaryResponse 薪资响应
 type SalaryResponse struct {
	ID             uint    `json:"id"`
	EmployeeID     uint    `json:"employee_id"`
	EmployeeName   string  `json:"employee_name"`
	DepartmentName string  `json:"department_name"`
	PayPeriod      string  `json:"pay_period"`
	BasicSalary    float64 `json:"basic_salary"`
	Allowances     float64 `json:"allowances"`
	Deductions     float64 `json:"deductions"`
	OvertimePay    float64 `json:"overtime_pay"`
	Bonus          float64 `json:"bonus"`
	Tax            float64 `json:"tax"`
	NetSalary      float64 `json:"net_salary"`
	PaymentDate    string  `json:"payment_date"`
	Status         string  `json:"status"`
	Remarks        string  `json:"remarks"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// SalaryCalculateRequest 薪资计算请求
 type SalaryCalculateRequest struct {
	PayPeriod      string `json:"pay_period" binding:"required"`
	DepartmentID   uint   `json:"department_id"`
	CalculateAll   bool   `json:"calculate_all"`
}

// 培训相关结构体

// TrainingCreateRequest 创建培训请求
 type TrainingCreateRequest struct {
	Code           string `json:"code" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Type           string `json:"type" binding:"required"`
	StartDate      string `json:"start_date" binding:"required"`
	EndDate        string `json:"end_date" binding:"required"`
	Location       string `json:"location" binding:"required"`
	Trainer        string `json:"trainer"`
	Description    string `json:"description"`
	Status         string `json:"status"`
}

// TrainingUpdateRequest 更新培训请求
 type TrainingUpdateRequest struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Location       string `json:"location"`
	Trainer        string `json:"trainer"`
	Description    string `json:"description"`
	Status         string `json:"status"`
}

// TrainingResponse 培训响应
 type TrainingResponse struct {
	ID             uint    `json:"id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	StartDate      string  `json:"start_date"`
	EndDate        string  `json:"end_date"`
	Duration       int     `json:"duration"`
	Location       string  `json:"location"`
	Trainer        string  `json:"trainer"`
	ParticipantCount int    `json:"participant_count"`
	Description    string  `json:"description"`
	Status         string  `json:"status"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// EnrollEmployeeRequest 员工报名培训请求
 type EnrollEmployeeRequest struct {
	TrainingID     uint   `json:"training_id" binding:"required"`
	EmployeeIDs    []uint `json:"employee_ids" binding:"required,dive"`
}

// 招聘相关结构体

// RecruitmentCreateRequest 创建招聘请求
 type RecruitmentCreateRequest struct {
	Code           string `json:"code" binding:"required"`
	PositionID     uint   `json:"position_id" binding:"required"`
	DepartmentID   uint   `json:"department_id" binding:"required"`
	RecruitmentDate string `json:"recruitment_date" binding:"required"`
	RequiredCount  int    `json:"required_count" binding:"required"`
	Description    string `json:"description"`
	Status         string `json:"status"`
}

// RecruitmentUpdateRequest 更新招聘请求
 type RecruitmentUpdateRequest struct {
	PositionID     uint   `json:"position_id"`
	DepartmentID   uint   `json:"department_id"`
	RecruitmentDate string `json:"recruitment_date"`
	RequiredCount  int    `json:"required_count"`
	Description    string `json:"description"`
	Status         string `json:"status"`
}

// RecruitmentResponse 招聘响应
 type RecruitmentResponse struct {
	ID             uint    `json:"id"`
	Code           string  `json:"code"`
	PositionID     uint    `json:"position_id"`
	PositionName   string  `json:"position_name"`
	DepartmentID   uint    `json:"department_id"`
	DepartmentName string  `json:"department_name"`
	RecruitmentDate string  `json:"recruitment_date"`
	RequiredCount  int     `json:"required_count"`
	ApplicantCount int     `json:"applicant_count"`
	HiredCount     int     `json:"hired_count"`
	Description    string  `json:"description"`
	Status         string  `json:"status"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// ApplicantCreateRequest 创建应聘者请求
 type ApplicantCreateRequest struct {
	RecruitmentID  uint   `json:"recruitment_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Gender         string `json:"gender" binding:"required"`
	BirthDate      string `json:"birth_date"`
	Phone          string `json:"phone" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Education      string `json:"education"`
	Experience     string `json:"experience"`
	InterviewDate  string `json:"interview_date"`
	InterviewResult string `json:"interview_result"`
	Status         string `json:"status"`
	Remarks        string `json:"remarks"`
}

// ApplicantResponse 应聘者响应
 type ApplicantResponse struct {
	ID             uint    `json:"id"`
	RecruitmentID  uint    `json:"recruitment_id"`
	RecruitmentCode string  `json:"recruitment_code"`
	Name           string  `json:"name"`
	Gender         string  `json:"gender"`
	BirthDate      string  `json:"birth_date"`
	Age            int     `json:"age"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email"`
	Education      string  `json:"education"`
	Experience     string  `json:"experience"`
	InterviewDate  string  `json:"interview_date"`
	InterviewResult string  `json:"interview_result"`
	Status         string  `json:"status"`
	Remarks        string  `json:"remarks"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// 绩效评估相关结构体

// PerformanceCreateRequest 创建绩效评估请求
 type PerformanceCreateRequest struct {
	EmployeeID     uint   `json:"employee_id" binding:"required"`
	EvaluationPeriod string `json:"evaluation_period" binding:"required"`
	EvaluationDate string `json:"evaluation_date" binding:"required"`
	EvaluatorID    uint   `json:"evaluator_id" binding:"required"`
	Goals          string `json:"goals"`
	Achievements   string `json:"achievements"`
	Competencies   string `json:"competencies"`
	OverallScore   float64 `json:"overall_score"`
	Status         string `json:"status"`
	Remarks        string `json:"remarks"`
}

// PerformanceUpdateRequest 更新绩效评估请求
 type PerformanceUpdateRequest struct {
	EvaluationPeriod string `json:"evaluation_period"`
	EvaluationDate string `json:"evaluation_date"`
	EvaluatorID    uint   `json:"evaluator_id"`
	Goals          string `json:"goals"`
	Achievements   string `json:"achievements"`
	Competencies   string `json:"competencies"`
	OverallScore   float64 `json:"overall_score"`
	Status         string `json:"status"`
	Remarks        string `json:"remarks"`
}

// PerformanceResponse 绩效评估响应
 type PerformanceResponse struct {
	ID             uint    `json:"id"`
	EmployeeID     uint    `json:"employee_id"`
	EmployeeName   string  `json:"employee_name"`
	DepartmentName string  `json:"department_name"`
	EvaluationPeriod string  `json:"evaluation_period"`
	EvaluationDate string  `json:"evaluation_date"`
	EvaluatorID    uint    `json:"evaluator_id"`
	EvaluatorName  string  `json:"evaluator_name"`
	Goals          string  `json:"goals"`
	Achievements   string  `json:"achievements"`
	Competencies   string  `json:"competencies"`
	OverallScore   float64 `json:"overall_score"`
	Status         string  `json:"status"`
	Remarks        string  `json:"remarks"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// PerformanceEvaluateRequest 绩效评估请求
 type PerformanceEvaluateRequest struct {
	PerformanceID  uint    `json:"performance_id" binding:"required"`
	OverallScore   float64 `json:"overall_score" binding:"required"`
	Feedback       string  `json:"feedback" binding:"required"`
	Status         string  `json:"status" binding:"required"`
}