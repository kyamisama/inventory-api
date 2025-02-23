package dto

type CreateItemDto struct {
	Name        string `json:"name" binding:"required" err_msg:"Name is required"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity" binding:"required,min=0,max=9999"`
	CreatedBy   string `json:"created_by" binding:"required"`
}

type UpdateItemDto struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required" err_msg:"Name is required"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity" binding:"required,min=0,max=9999"`
	UpdatedBy   string `json:"updated_by" binding:"required"`
}
