package responses

type Appointment struct {
	ID          uint	`json:"id" gorm:"primary_key"`
	TrainerID   int 	`json:"trainer_id"`
	UserID   	int 	`json:"user_id"`
	StartsAt 	string 	`json:"starts_at"`
	EndsAt 		string 	`json:"ends_at"`
}