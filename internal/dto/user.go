package dto

type UpdateAdminAccessRequest struct {
	TargetUserID int    `json:"target_user_id" binding:"required"`
	NewRole      string `json:"new_role" binding:"required"` // "admin", "superadmin"
}
