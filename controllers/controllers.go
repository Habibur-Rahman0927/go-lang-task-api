package controllers

import (
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
}

type UpdateTaskInput struct {
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
}

// GET /tasks
// Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{AssingedTo: input.AssingedTo, Task: input.Task}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	var task models.Task
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// PATCH /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&task).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

type OperationBulletinOperation struct {
	ID                  uint   `json:"id"`
	OperationBulletinID uint   `json:"operation_bulletin_id"`
	Serial              int    `json:"serial"`
	OperationID         uint   `json:"operation_id"`
	OperationName       string `json:"operation_name"`
	MachineCategoryName string `json:"machine_category_name"`
	SMV                 string `json:"smv"`
	Attachments         string `json:"attachements"`
	OpAOPReq            string `json:"op_aop_req"`
	OpAOPAlot           uint   `json:"op_aop_alot"`
	TargetHourMan       uint   `json:"target_hour_man"`
	TotalTargetHour     uint   `json:"total_target_hour"`
	CreatedBy           uint   `json:"created_by"`
	UpdatedBy           *uint  `json:"updated_by"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}

type OperationBulletin struct {
	ID                          uint   `json:"id"`
	OrderProductTypeID          uint   `json:"order_product_type_id"`
	Status                      string `json:"status"`
	Name                        string `json:"name"`
	WorkingTime                 string `json:"working_time"`
	TargetEfficiency            string `json:"target_efficiency"`
	TheoreticalOutputHour       string `json:"theoritical_output_hour"`
	TheoreticalOutputDay        string `json:"theoritical_output_day"`
	TotalSMV                    string `json:"total_smv"`
	NoOfOperations              uint   `json:"no_of_operations"`
	ForecastOutputHour          string `json:"forecast_output_hour"`
	ForecastOutputDay           string `json:"forecast_output_day"`
	MachineSMV                  string `json:"machine_smv"`
	NoOfAOP                     uint   `json:"no_of_aop"`
	TargetOutputHour            uint   `json:"target_output_hour"`
	TargetOutputDay             uint   `json:"target_output_day"`
	ManualSMV                   string `json:"manual_smv"`
	TotalManpower               uint   `json:"total_manpower"`
	BalanceEfficiency           string `json:"balance_efficiency"`
	PitchTime                   string `json:"pitch_time"`
	StdDeviation                string `json:"std_deviation"`
	Primary                     uint
	CreatedBy                   uint   `json:"created_by"`
	UpdatedBy                   uint   `json:"updated_by"`
	CreatedAt                   string `json:"created_at"`
	UpdatedAt                   string `json:"updated_at"`
	OrderProductType            OrderProductType
	OperationBulletinOperations []OperationBulletinOperation `json:"operation_bulletin_operations"`
}

type OrderProductTypeCosting struct {
	ID                  uint   `json:"id"`
	OrderProductTypeID  uint   `json:"order_product_type_id"`
	SMV                 string `json:"smv"`
	CPM                 string `json:"cpm"`
	Efficiency          string `json:"efficiency"`
	CM                  string `json:"cm"`
	TotalFabricCost     string `json:"total_fabric_cost"`
	TrimAccessoriesCost string `json:"trim_accessories_cost"`
	PrintCost           string `json:"print_cost"`
	WashCost            string `json:"wash_cost"`
	EmbroideryCost      string `json:"embroidery_cost"`
	AllOverPrintCost    string `json:"all_over_print_cost"`
	CommercialCost      string `json:"commercial_cost"`
	Margin              string `json:"margin"`
	FOB                 string `json:"fob"`
	CreatedBy           uint   `json:"created_by"`
	UpdatedBy           uint   `json:"updated_by"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}

type OrderProductType struct {
	ID                      uint                    `json:"id"`
	OrderID                 uint                    `json:"order_id"`
	ProductTypeName         string                  `json:"product_type_name"`
	ImageURL                string                  `json:"image_url"`
	TechPackURL             string                  `json:"tech_pack_url"`
	SizeChartID             uint                    `json:"size_chart_id"`
	CreatedBy               uint                    `json:"created_by"`
	UpdatedBy               uint                    `json:"updated_by"`
	CreatedAt               string                  `json:"created_at"`
	UpdatedAt               string                  `json:"updated_at"`
	OrderProductTypeCosting OrderProductTypeCosting `json:"order_product_type_costing"`
}

type BundleTag struct {
	ID                           uint   `json:"id"`
	JobCardID                    uint   `json:"job_card_id"`
	RFID                         string `json:"rfid"`
	OrderPurchaseID              uint   `json:"order_purchase_id"`
	OrderPurchaseCountryID       uint   `json:"order_purchase_country_id"`
	SizeLabel                    string `json:"size_label"`
	SerialStart                  uint   `json:"serial_start"`
	SerialEnd                    uint   `json:"serial_end"`
	Quantity                     uint
	ChallanType                  string `json:"challan_type"`
	OperationBulletinID          uint   `json:"operation_bulletin_id"`
	OperationBulletinOperationID uint   `json:"operation_bulletin_operation_id"`
	Status                       string
	CreatedBy                    uint                    `json:"created_by"`
	UpdatedBy                    uint                    `json:"updated_by"`
	CreatedAt                    string                  `json:"created_at"`
	UpdatedAt                    string                  `json:"updated_at"`
	SupervisorID                 uint                    `json:"supervisor_id"`
	TagType                      string                  `json:"tag_type"`
	ParentTagID                  *uint                   `json:"parent_tag_id"`
	BundleTagStatusChange        []BundleTagStatusChange `json:"bundle_tag_status_changes"`
	OperationBulletin            OperationBulletin       `json:"operation_bulletin"`
}

type BundleTagStatusChange struct {
	ID                           uint                        `json:"id"`
	BundleTagID                  uint                        `json:"bundle_tag_id"`
	OperationBulletinOperationID uint                        `json:"operation_bulletin_operation_id"`
	SewingMachineID              string                      `json:"sewing_machine_id"`
	Status                       string                      `json:"status"`
	Quantity                     uint                        `json:"quantity"`
	CreatedAt                    string                      `json:"created_at"`
	UpdatedAt                    string                      `json:"updated_at"`
	OperationBulletinOperation   OperationBulletinOperation  `gorm:"foreignkey:OperationBulletinOperationID"`
	BundleTagStatusChangeUsers   []BundleTagStatusChangeUser `gorm:"foreignkey:BundleTagStatusChangeID"`
}

type BundleTagStatusChangeUser struct {
	ID                      uint   `json:"id"`
	BundleTagID             uint   `json:"bundle_tag_id"`
	BundleTagStatusChangeID uint   `json:"bundle_tag_status_change_id"`
	UserID                  uint   `json:"user_id"`
	CreatedAt               string `json:"created_at"`
	UpdatedAt               string `json:"updated_at"`
	User                    User   `json:"user"`
}

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FilterAndRetrieveBundleTags(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var filterParams struct {
		Status []string `json:"status"`
		Date   string   `json:"date"`
	}

	if err := c.ShouldBindJSON(&filterParams); err != nil {
		c.JSON(400, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	var bundleTags []BundleTag

	query := db.Preload("OperationBulletin.OrderProductType.OrderProductTypeCosting").
		Preload("OperationBulletin.OperationBulletinOperations").
		// Preload("BundleTagStatusChange", "bundle_tag_status_changes.status IN (?)", filterParams.Status).
		Preload("BundleTagStatusChange", func(db *gorm.DB) *gorm.DB {
			return db.Preload("OperationBulletinOperation").
				Preload("BundleTagStatusChangeUsers.User")
		}, "bundle_tag_status_changes.status IN (?)", filterParams.Status).
		Where("bundle_tags.status IN (?)", filterParams.Status)

	if filterParams.Date != "" {
		query = query.Where("DATE(bundle_tags.updated_at) = ?", filterParams.Date)
	}

	if err := query.Find(&bundleTags).Error; err != nil {
		c.JSON(500, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":    "Filtered and retrieved data successfully",
		"bundleTags": bundleTags,
	})
}
