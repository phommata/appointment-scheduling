package requests

type AvailableAppointments struct {
	TrainerID   int  	`form:"trainer_id" binding:"required"`
	StartsAt 	string  `form:"starts_at" binding:"required"`
	EndsAt  	string 	`form:"ends_at" binding:"required"`
}
