package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wu136995/ginx/internal/api/handlers"
)

// SetupHRRoutes 设置人力资源模块路由
func SetupHRRoutes(router *gin.Engine, hrHandler *handlers.HRHandler) {
	// 人力资源模块API路由组
	hr := router.Group("/api/v1/hr")
	{
		// 员工管理
		employees := hr.Group("/employees")
		{
			employees.GET("", hrHandler.GetEmployeeList)
			employees.GET("/:id", hrHandler.GetEmployeeDetail)
			employees.POST("", hrHandler.CreateEmployee)
			employees.PUT("/:id", hrHandler.UpdateEmployee)
			employees.DELETE("/:id", hrHandler.DeleteEmployee)
		}

		// 部门管理
		departments := hr.Group("/departments")
		{
			departments.GET("", hrHandler.GetDepartmentList)
			departments.GET("/:id", hrHandler.GetDepartmentDetail)
			departments.POST("", hrHandler.CreateDepartment)
			departments.PUT("/:id", hrHandler.UpdateDepartment)
			departments.DELETE("/:id", hrHandler.DeleteDepartment)
		}

		// 职位管理
		positions := hr.Group("/positions")
		{
			positions.GET("", hrHandler.GetPositionList)
			positions.GET("/:id", hrHandler.GetPositionDetail)
			positions.POST("", hrHandler.CreatePosition)
			positions.PUT("/:id", hrHandler.UpdatePosition)
			positions.DELETE("/:id", hrHandler.DeletePosition)
		}

		// 考勤管理
		attendance := hr.Group("/attendance")
		{
			attendance.GET("", hrHandler.GetAttendanceList)
			attendance.GET("/:id", hrHandler.GetAttendanceDetail)
			attendance.POST("", hrHandler.CreateAttendance)
			attendance.PUT("/:id", hrHandler.UpdateAttendance)
			attendance.DELETE("/:id", hrHandler.DeleteAttendance)
		}

		// 薪资管理
		salary := hr.Group("/salary")
		{
			salary.GET("", hrHandler.GetSalaryList)
			salary.GET("/:id", hrHandler.GetSalaryDetail)
			salary.POST("", hrHandler.CreateSalary)
			salary.PUT("/:id", hrHandler.UpdateSalary)
			salary.DELETE("/:id", hrHandler.DeleteSalary)
		}

		// 培训管理
		training := hr.Group("/training")
		{
			training.GET("", hrHandler.GetTrainingList)
			training.GET("/:id", hrHandler.GetTrainingDetail)
			training.POST("", hrHandler.CreateTraining)
			training.PUT("/:id", hrHandler.UpdateTraining)
			training.DELETE("/:id", hrHandler.DeleteTraining)
		}

		// 招聘管理
		recruitment := hr.Group("/recruitment")
		{
			recruitment.GET("", hrHandler.GetRecruitmentList)
			recruitment.GET("/:id", hrHandler.GetRecruitmentDetail)
			recruitment.POST("", hrHandler.CreateRecruitment)
			recruitment.PUT("/:id", hrHandler.UpdateRecruitment)
			recruitment.DELETE("/:id", hrHandler.DeleteRecruitment)
		}

		// 绩效评估管理
		performance := hr.Group("/performance")
		{
			performance.GET("", hrHandler.GetPerformanceList)
			performance.GET("/:id", hrHandler.GetPerformanceDetail)
			performance.POST("", hrHandler.CreatePerformance)
			performance.PUT("/:id", hrHandler.UpdatePerformance)
			performance.DELETE("/:id", hrHandler.DeletePerformance)
		}
	}
}


