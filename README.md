# Appointment Scheduling

## Duration
Around 60-90 minutes (please make sure to send across your code around the 90-minute mark at the most).

## Motivation
Clients need to be able to schedule an appointment with their trainer through an HTTP API.

## Instructions

The client should be able to pick from a list of available times. Appointments for a trainer should not overlap.
Appointments are 30 minutes long.
Appointments should be scheduled at :00, :30 minutes after the hour during business hours.
Business hours are M-F 8am-5pm Pacific Time

Your job is to create an HTTP JSON API written in Go with the following endpoints:

* Get a list of available appointment times for a trainer between two dates
  Parameters:
    trainer_id
    starts_at
    ends_at
  Returns:
    list of available appointment times
* Post an appointment (as JSON)
  Fields:
    trainer_id
    user_id
    starts_at
    ends_at
* Get a list of scheduled appointments for a trainer
  Parameters:
    trainer_id

appointments.json contains the current list of appointments in this format:

     [
        {
            "id": 1
            "trainer_id": 1
            "user_id": 2,
            "starts_at": "2019-01-25T09:00:00-08:00",
            "ends_at": "2019-01-25T09:30:00-08:00"
        }
    ]

You can store appointments in this file, a database or any back end storage you prefer.

## Getting started 
Build our custom image(i.e dockerfile provided on build key)

    docker-compose build 

Starts all our container(service) configured on yaml file

    docker-compose up

Removes all our container(service)

    docker-compose down

## Quick start

    docker-compose down; docker-compose build; docker-compose up

## Examples

### Get appointments 

Request

    curl -X GET \
      'http://localhost:8081/appointments?trainer_id=1'

Response

    {
        "data": [
            {
                "id": 1,
                "trainer_id": 1,
                "user_id": 1,
                "starts_at": "2019-01-25T09:00:00Z",
                "ends_at": "2019-01-25T09:30:00Z"
            }
        ],
        "limit": 12,
        "offset": 0,
        "total": 1
    }
    
### Get available appointments

Request

    curl -X GET \
      'http://localhost:8081/available-appointments?trainer_id=1&starts_at=2019-01-25&ends_at=2019-01-26'

Response   

    {
        "data": [
            "2019-01-25 08:00:00",
            "2019-01-25 08:30:00",
            "2019-01-25 09:30:00",
            "2019-01-25 10:00:00",
            "2019-01-25 10:30:00",
            "2019-01-25 11:00:00",
            "2019-01-25 11:30:00",
            "2019-01-25 12:00:00",
            "2019-01-25 12:30:00",
            "2019-01-25 13:00:00",
            "2019-01-25 13:30:00",
            "2019-01-25 14:00:00",
            "2019-01-25 14:30:00",
            "2019-01-25 15:00:00",
            "2019-01-25 15:30:00",
            "2019-01-25 16:00:00",
            "2019-01-25 16:30:00",
            "2019-01-26 08:00:00",
            "2019-01-26 08:30:00",
            "2019-01-26 09:00:00",
            "2019-01-26 09:30:00",
            "2019-01-26 10:00:00",
            "2019-01-26 10:30:00",
            "2019-01-26 11:00:00",
            "2019-01-26 11:30:00",
            "2019-01-26 12:00:00",
            "2019-01-26 12:30:00",
            "2019-01-26 13:00:00",
            "2019-01-26 13:30:00",
            "2019-01-26 14:00:00",
            "2019-01-26 14:30:00",
            "2019-01-26 15:00:00",
            "2019-01-26 15:30:00",
            "2019-01-26 16:00:00",
            "2019-01-26 16:30:00"
        ]
    } 
    
### Create appointment

Request

    curl -X POST \
      http://localhost:8080/appointment \
      -H 'content-type: application/json' \
      -d '{
        "trainer_id": 1,
        "user_id": 1,
        "starts_at":"2019-01-25 09:00",
        "ends_at":"2019-01-25 09:30"
    }'

Response    

    {}