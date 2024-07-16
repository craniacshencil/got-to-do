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

func (ApiConfig *ApiCfg) AddToListHandler(w http.ResponseWriter, r *http.Request) {
	// Authorize user
	_, err := myJwt.AuthorizeUser(r)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get list_id from URL
	listIDString := chi.URLParam(r, "list_id")
	listID := uuid.MustParse(listIDString)
	dateString := chi.URLParam(r, "date")
	date, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		log.Println("ERR:While retreiving tasks from DB", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Parse new tasks from JSON to a map
	newTasks := make(map[int]Task)
	utils.ParseJSON(r, &newTasks)
	newTasksCount := len(newTasks)

	// Get Previous tasks and add them to the map
	existingTasks, err := ApiConfig.DB.GetTasks(r.Context(), listID)
	if err != nil {
		log.Println("ERR:While retreiving tasks from DB", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	var taskStruct Task
	for _, task := range existingTasks {
		taskStruct = Task{
			TaskName:   task.TaskName,
			StartTime:  task.StartTime,
			EndTime:    task.EndTime,
			Completion: task.Completion,
		}
		// because i started indexing from 1
		newTasks[len(newTasks)+1] = taskStruct
	}

	// Validate my tasks
	log.Println(newTasks)
	err = validateTimings(newTasks, date)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(newTasksCount)
	// Add new tasks to the database
	for i := 1; i <= newTasksCount; i++ {
		ApiConfig.DB.CreateTask(r.Context(), database.CreateTaskParams{
			TaskID:     uuid.New(),
			ListID:     listID,
			TaskName:   newTasks[i].TaskName,
			StartTime:  newTasks[i].StartTime,
			EndTime:    newTasks[i].EndTime,
			Completion: newTasks[i].Completion,
		})
	}

	utils.WriteJSON(w, http.StatusAccepted, "Added tasks successfully")
}
