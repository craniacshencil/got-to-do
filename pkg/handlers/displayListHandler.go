package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/craniacshencil/got_to_do/internal/database"
	"github.com/craniacshencil/got_to_do/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (ApiConfig *ApiCfg) DisplayListHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve parameters from the url
	userIDString := chi.URLParam(r, "user_id")
	dateString := chi.URLParam(r, "date")
	userID := uuid.MustParse(userIDString)
	date, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		log.Println("ERR: While parsing date from string to time.Time:", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get ListID
	listID, err := ApiConfig.DB.GetListID(r.Context(), database.GetListIDParams{
		Date:   date,
		UserID: userID,
	})
	if err != nil {
		log.Println("ERR: While retrieving listID:", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get all the tasks from the ListID
	taskSlice, err := ApiConfig.DB.GetTasks(r.Context(), listID)
	if err != nil {
		log.Println("ERR: While retrieving tasks using listID:", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Sorting the tasks by time
	sortTasksByTime(taskSlice)

	// Iterating through tasks to create a map
	Todos := make(map[int]Task)
	var taskStruct Task
	for idx, task := range taskSlice {
		// NOTE: Indexing here starts from 0 (was using indexing from 1 while cURLing)
		taskStruct.TaskName = task.TaskName
		taskStruct.StartTime = task.StartTime
		taskStruct.EndTime = task.EndTime
		Todos[idx] = taskStruct
	}

	// Send Response as JSON
	utils.WriteJSON(w, http.StatusCreated, Todos)
}

func sortTasksByTime(taskSlice []database.Task) {
	for i := 0; i < len(taskSlice)-1; i++ {
		for j := 0; j < len(taskSlice)-1-i; j++ {
			if taskSlice[j].StartTime.After(taskSlice[j+1].StartTime) {
				taskSlice[j], taskSlice[j+1] = taskSlice[j+1], taskSlice[j]
			}
		}
	}
}
