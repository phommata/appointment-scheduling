package requests

type Appointments struct {
	TrainerID   int  	`form:"trainer_id" binding:"required"`
	Limit   	int  	`form:"limit"`
	Offset   	int  	`form:"offset"`
}
