package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/phommata/appointment-scheduling/config"
	"github.com/phommata/appointment-scheduling/models"
	"github.com/phommata/appointment-scheduling/requests"
	"github.com/phommata/appointment-scheduling/responses"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

const Error = "error"

type ApppointmentRepo struct {
	Db *gorm.DB
}

func New(conf *config.Conf) *ApppointmentRepo {
	db := models.InitDb(conf)

	return &ApppointmentRepo{Db: db}
}

// GET /appointments
// Get all appointments
func (repository *ApppointmentRepo) GetAppointments(c *gin.Context) {
	var appointments	requests.Appointments
	var appointmentResp []responses.Appointment
	var count int64

	if err := c.Bind(&appointments); err != nil {
		errMsg := "cannot bind appointments"
		log.Println(errMsg, err)

		c.JSON(http.StatusBadRequest,
			gin.H{Error: err.Error()},
		)
		return
	}

	log.Println("appointments")
	log.Println(appointments)

	repository.Db.Table("appointments").Select("count(*)").Count(&count)

	if appointments.Limit != 0 {
		repository.Db.Debug().Limit(appointments.Limit).Offset(appointments.Offset).Table("appointments a").
			Select("*").
			Where("a.trainer_id = " + strconv.Itoa(appointments.TrainerID)).
			Find(&appointmentResp).Order("a.created_at")
	}

	repository.Db.Debug().Table("appointments a").
		Select("*").
		Where("a.trainer_id = " + strconv.Itoa(appointments.TrainerID)).
		Find(&appointmentResp).Order("a.created_at")

	c.JSON(http.StatusOK,
		gin.H{
			"data": appointmentResp,
			"limit": appointments.Limit,
			"offset": appointments.Offset,
			"total": count,
		},
	)
}

// GET /available-appointments
// Get all available appointments
func (repository *ApppointmentRepo) ListAvailableAppointments(c *gin.Context) {
	var availAppts 		requests.AvailableAppointments
	var appointments 	[]responses.Appointment
	var errMsg 			string
	var startsAt 		string
	var endsAt	 		string
	var dateTImesAvail  []string
	var timeTmp  		time.Time
	dateTimeIntervals := make(map[string]bool)

	if err := c.Bind(&availAppts); err != nil {
		errMsg := "cannot bind availAppts"
		log.Println(errMsg, err)

		c.JSON(http.StatusBadRequest,
			gin.H{Error: err.Error()},
		)
		return
	}

	log.Println("availAppts")
	log.Println(availAppts)

	startsAt = availAppts.StartsAt + " 08:00:00"
	endsAt = availAppts.EndsAt + " 16:30:00"

	/* build intervals start - end date
	query for intervals not matching*/
	startDateT, err := time.Parse("2006-01-02 15:04:05", startsAt)

	if err != nil {
		errMsg = "start_date cannot be parsed"
		log.Println(errMsg, err)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	endDateT, err := time.Parse("2006-01-02 15:04:05", endsAt)

	if err != nil {
		errMsg = "end_date cannot be parsed"
		log.Println(errMsg, err)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	for date := startDateT; !date.After(endDateT) ; date = date.Add(time.Minute * 30) {
		log.Println("date")
		log.Println(date)

		if (date.Weekday() != time.Saturday || date.Weekday() != time.Sunday) && (date.Hour() >= 8 && date.Hour() <= 16) {
			dateTimeIntervals[date.Format("2006-01-02T15:04:05Z")] = true
		}

		fmt.Println(date.Format("2006-01-02 15:04:05"))
	}

	repository.Db.Debug().Table("appointments a").
		Select("*").
		Where("a.trainer_id = " + strconv.Itoa(availAppts.TrainerID)).
		Where("a.starts_at >= '" + startsAt + "'").
		Where("a.ends_at <= '" + endsAt + "'").
		Find(&appointments).Order("a.created_at")

	log.Println("appointments")
	log.Println(appointments)

	/*
	filter results for date time intervals
	 */
	log.Println("dateTimeIntervals")
	log.Println(dateTimeIntervals)

	for _, val := range appointments {

		log.Println("val.StartsAt")
		log.Printf("%T %v %s\n", val.StartsAt, val.StartsAt, val.StartsAt)

		if _, ok := dateTimeIntervals[val.StartsAt]; ok {
			log.Println("delete val.StartsAt")
			delete(dateTimeIntervals, val.StartsAt)
		}
	}

	for key, _ := range dateTimeIntervals {
		timeTmp, err = time.Parse("2006-01-02T15:04:05Z", key)
		timeStr := timeTmp.Format("2006-01-02 15:04:05")

		dateTImesAvail = append(dateTImesAvail, timeStr)
	}

	sort.Strings(dateTImesAvail)

	log.Println("dateTimeIntervals")
	log.Println(dateTimeIntervals)

	c.JSON(http.StatusOK,
		gin.H{
			"data": dateTImesAvail,
		},
	)
}

// POST /appointment
// Create appointment
func (repository *ApppointmentRepo) CreateAppointment(c *gin.Context) {
	var startsAtT 			time.Time
	var endsAtT 			time.Time
	var appointmentResp 	responses.Appointment
	var appointmentEmpty 	responses.Appointment
	var appointmentReq		requests.Appointment
	var trainerId			string
	var err					error

	if err = c.Bind(&appointmentReq); err != nil {
		errMsg := "cannot bind appointmentReq"
		log.Println(errMsg, err)

		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	trainerId = strconv.Itoa(appointmentReq.TrainerID)
	userId := strconv.Itoa(appointmentReq.UserID)
	startsAtT, err = time.Parse("2006-01-02 15:04", appointmentReq.StartsAt)

	if err != nil {
		errMsg := "starts_at cannot be parsed"
		log.Println(errMsg, err)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	endsAtT, err = time.Parse("2006-01-02 15:04", appointmentReq.EndsAt)

	if err != nil {
		errMsg := "ends_at cannot be parsed"
		log.Println(errMsg, err)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	// start date M - F
	if startsAtT.Weekday() == time.Saturday || startsAtT.Weekday() == time.Saturday {
		errMsg := "starts_at cannot be weekend"
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	// start time 8 - 16:30
	log.Println("startsAtT.Hour()")
	log.Println(startsAtT.Hour())
	log.Println("startsAtT.Minute()")
	log.Println(startsAtT.Minute())

	if startsAtT.Hour() < 8 || startsAtT.Hour() >= 16 || (startsAtT.Hour() >= 16 && startsAtT.Minute() > 30) {
		errMsg := "starts_at must be 8 - 16:30"
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	// start time is 00 or 30
	log.Println("startsAtT.Minute()")
	log.Println(startsAtT.Minute())

	if startsAtT.Minute() != 0 && startsAtT.Minute() != 30 {
		errMsg := "starts_at must be at 00 or 30 minutes"
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}


	// end time 8:30 - 17:00
	if endsAtT.Hour() < 8 || endsAtT.Hour() > 17 {
		errMsg := "ends_at must be 8:30 - 17:00"
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	// end time is 00 or 30
	if endsAtT.Minute() != 0 && endsAtT.Minute() != 30 {
		errMsg := "ends_at must be at 00 or 30 minutes"
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	// end date is 30 interval
	if startsAtT.Add(30 * time.Minute) != endsAtT {
		errMsg := "starts_at - ends_at must be 30 minutes"
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	// start date < end date
	if startsAtT.After(endsAtT) {
		errMsg := "starts_at cannot be after ends_at"
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	// start date avail
	repository.Db.Debug().Table("appointments a").
		Select("*").
		Where("a.trainer_id = " + trainerId).
		Where("a.user_id = " + userId).
		Where("a.starts_at = '" + appointmentReq.StartsAt + "'").
		Find(&appointmentResp).Order("a.created_at")

	log.Println("appointmentResp")
	log.Println(appointmentResp)

	if appointmentResp != appointmentEmpty {
		errMsg := "appointment is not available"
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest,
			gin.H{Error: errMsg},
		)
		return
	}

	repository.Db.Debug().Table("appointments").
		Create(&models.Appointment{
				TrainerID:	appointmentReq.TrainerID,
				UserID:		appointmentReq.UserID,
				StartsAt:	startsAtT,
				EndsAt:		endsAtT,
			},
		)

	c.JSON(http.StatusCreated,
		gin.H{},
	)
}