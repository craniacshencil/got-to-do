package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/craniacshencil/got_to_do/internal/database"
	"github.com/craniacshencil/got_to_do/pkg/myJwt"
	"github.com/craniacshencil/got_to_do/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Task struct {
	TaskName   string    `json:"task_name"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Completion bool      `json:"completion"`
}

func (ApiConfig *ApiCfg) CreateListHandler(w http.ResponseWriter, r *http.Request) {
	// NOTE: Even though there the path contains both supposedly contains userID and date
	// I forgot about them and didn't use them from the path

	// Authorize user and get userID
	token, err := myJwt.AuthorizeUser(r)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	claimsMap := token.Claims.(jwt.MapClaims)
	userIDString := claimsMap["sub"].(string)
	userID := uuid.MustParse(userIDString)

	// Parse the json
	var Todo map[int]Task
	err = utils.ParseJSON(r, &Todo)
	if err != nil {
		log.Println("ERR: While parsing Todo JSON", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get date (Not checking the date of every task as coming from the frontend all the tasks would have the same date)
	// NOTE: Probably didn't need to do this
	// 1. Could have retrieved it from path
	// 2. Postgres would've converted time.Time to appropriate date format automatically(it did, with time, below in the tasks relation)
	dateTimeString := Todo[1].StartTime.Format(time.DateOnly)
	date, err := time.Parse(time.DateOnly, dateTimeString)
	if err != nil {
		log.Println("ERR: While extracting date out of time:", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Validate all todos
	err = validateTimings(Todo, date)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Adding the list to DB
	todoEntry, err := ApiConfig.DB.CreateList(r.Context(), database.CreateListParams{
		UserID: userID,
		ListID: uuid.New(),
		Date:   date,
	})
	if err != nil {
		log.Println("ERR: While storing list:", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Adding each task to DB
	for taskNo, task := range Todo {
		taskDB, err := ApiConfig.DB.CreateTask(r.Context(), database.CreateTaskParams{
			TaskID:     uuid.New(),
			ListID:     todoEntry.ListID,
			TaskName:   task.TaskName,
			StartTime:  task.StartTime, // voila, didn't have to convert time.Time to a specific format
			EndTime:    task.EndTime,
			Completion: task.Completion,
		})
		if err != nil {
			log.Printf("ERR: Couldn't save task %d to DB: %v", taskNo, err)
			utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		log.Println(taskDB)
	}

	// Send appropriate response
	utils.WriteJSON(
		w,
		http.StatusCreated,
		"Todo list and the tasks have been added to the database",
	)
}

// Checks for the following:
//   - STEP 1: Make sure date is valid
//   - STEP 2: Making sure that all the tasks have endtime > starttime
//   - STEP 3: Validate the timings make sure there is no clash Eg: One task starts and ends at: 13:00 to 14:00, if other task starts at 13:30 then error should be reported
func validateTimings(taskTimings map[int]Task, date time.Time) error {
	// STEP 1
	if date.Before(time.Now()) {
		return fmt.Errorf("date is invalid(before today)")
	}

	var startTimeArr, endTimeArr []time.Time
	var i, j int
	for idx, task := range taskTimings {
		// Had to use this because directly using task.StartTime.String() does not yield RFC3339 format
		startTimeString := task.StartTime.Format(time.TimeOnly)
		start, err := time.Parse(time.TimeOnly, startTimeString)
		if err != nil {
			return fmt.Errorf("couldn't parse startime %v: %v", task.StartTime, err)
		}

		endTimeString := task.EndTime.Format(time.TimeOnly)
		end, err := time.Parse(time.TimeOnly, endTimeString)
		if err != nil {
			return fmt.Errorf("couldn't parse startime %v: %v", task.EndTime, err)
		}

		// STEP 2
		if start.After(end) {
			return fmt.Errorf(
				"endtime: %v hours is before starttime: %v hours for task %d",
				start,
				end,
				idx,
			)
		}

		startTimeArr = append(startTimeArr, start)
		endTimeArr = append(endTimeArr, end)
	}

	// STEP 3
	for i = 0; i < len(startTimeArr); i++ {
		for j = 0; j < len(startTimeArr); j++ {
			if i == j {
				continue
			}
			// Condition says that if startTime of a task lies in time assigned for other task or startime of 2 tasks is same
			if startTimeArr[i].Before(endTimeArr[j]) &&
				(startTimeArr[i].After(startTimeArr[j]) || startTimeArr[i] == startTimeArr[j]) {
				return fmt.Errorf(
					"starting time of task %d : %v hours\nend time of task %d : %v hours\nstarting time of task %d: %v hours\nThis is clashing",
					i,
					startTimeArr[i],
					i,
					endTimeArr[i],
					j,
					startTimeArr[j],
				)
			}
		}
	}
	return nil
}
