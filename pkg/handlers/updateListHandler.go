package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/craniacshencil/got_to_do/internal/database"
	"github.com/craniacshencil/got_to_do/pkg/myJwt"
	"github.com/craniacshencil/got_to_do/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (ApiConfig *ApiCfg) UpdateListHandler(w http.ResponseWriter, r *http.Request) {
	// Authorize with JWT

	// REMOVE THESE LINES
	_, err := myJwt.AuthorizeUser(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusUnauthorized, err)
		return
	}

	// Extract user, list_id
	listIdString := chi.URLParam(r, "list_id")
	listId := uuid.MustParse(listIdString)
	dateString := chi.URLParam(r, "date")
	date, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		log.Println("ERR: couldnt' parse the date: ", err)
		utils.WriteJSON(w, http.StatusNotFound, err)
		return
	}

	// Creating a map out of updated tasks
	var updatedTasks map[uuid.UUID]Task
	err = utils.ParseJSON(r, &updatedTasks)
	if err != nil {
		log.Println("ERR: couldnt' parse incoming updated tasks: ", err)
		utils.WriteJSON(w, http.StatusNotFound, err)
		return
	}

	// Fetching original tasks from DB
	oldTasksObj, err := ApiConfig.DB.GetTasks(r.Context(), listId)
	if err != nil {
		log.Println("ERR: couldn't get tasks from DB: ", err)
		utils.WriteJSON(w, http.StatusNotFound, err)
		return
	}

	// Converting to map
	oldTasks := make(map[uuid.UUID]Task)
	for _, task := range oldTasksObj {
		oldTasks[task.TaskID] = Task{
			TaskName:   task.TaskName,
			StartTime:  task.StartTime,
			EndTime:    task.EndTime,
			Completion: task.Completion,
		}
	}

	// TODO: Create struct for validating timings
	counter := 1
	tasksValidator := make(map[int]Task)
	for _, task := range updatedTasks {
		tasksValidator[counter] = task
		counter++
	}

	// Validate timings
	err = validateTimings(tasksValidator, date)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusNotFound, err)
		return
	}

	// Update all tasks
	for id, task := range updatedTasks {
		err = ApiConfig.DB.UpdateTask(r.Context(), database.UpdateTaskParams{
			TaskName:   task.TaskName,
			StartTime:  task.StartTime,
			EndTime:    task.EndTime,
			Completion: task.Completion,
			TaskID:     id,
		})
	}

	utils.WriteJSON(w, http.StatusCreated, "updated the todos")
	return
}
