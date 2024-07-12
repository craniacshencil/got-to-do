package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/craniacshencil/got_to_do/pkg/myJwt"
	"github.com/craniacshencil/got_to_do/utils"
)

type Task struct {
	TaskName  string    `json:"task_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (ApiConfig *ApiCfg) CreateListHandler(w http.ResponseWriter, r *http.Request) {
	// Retreive jwt token from cookies
	token, err := r.Cookie("jwt")
	if err != nil {
		log.Println("ERR: Couldn't find cookie", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Validate the cookie
	_, err = myJwt.ValidateToken(token.Value)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Parse the json
	var Todo map[int]Task
	err = utils.ParseJSON(r, &Todo)
	if err != nil {
		log.Println("ERR: While parsing Todo JSON", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Validate the timings make sure there is no clash
	// Eg: One task starts and ends at: 13:00 to 14:00, if other task starts at 13:30 then error should be reported
	err = validateTimings(Todo)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func validateTimings(todo map[int]Task) error {
	// TODO: Primary validation
	// Making sure that all the tasks have endtime > starttime
	// Making sure all the tasks are of proper type(i think using a struct basically makes sure that this is taken care of)
	var start, end time.Time
	var i, j, startVal, endVal int
	var startTimeArr, endTimeArr []int
	for _, task := range todo {
		start = task.StartTime
		end = task.EndTime

		// getting time in form: 13:00 to 14:00 -> 1300, 1400
		startVal = start.Hour()*1000 + start.Minute()
		endVal = end.Hour()*1000 + end.Minute()

		startTimeArr = append(startTimeArr, startVal)
		endTimeArr = append(endTimeArr, endVal)
	}
	for i = 0; i < len(startTimeArr); i++ {
		for j = 0; j < len(startTimeArr); j++ {
			if i == j {
				continue
			}
			// Condition says that if startTime of a task lies in time assigned for other task
			if startTimeArr[i] < endTimeArr[j] && startTimeArr[i] >= startTimeArr[j] {
				return fmt.Errorf(
					"starting time of task %d: %d\nend time of task %d: %d\nstarting time of task %d: %d\nThis is clashing",
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
