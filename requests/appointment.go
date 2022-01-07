package requests

type Appointment struct {
	TrainerID   int  	`json:"trainer_id" binding:"required"`
	UserID  	int  	`json:"user_id" binding:"required"`
	StartsAt 	string  `json:"starts_at" binding:"required"`
	EndsAt  	string 	`json:"ends_at" binding:"required"`
}
