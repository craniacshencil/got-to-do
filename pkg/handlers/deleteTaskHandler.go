package handlers

import (
	"log"
	"net/http"

	"github.com/craniacshencil/got_to_do/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (ApiConfig *ApiCfg) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Extracting info out of URLs
	taskIdString := chi.URLParam(r, "task_id")
	taskId := uuid.MustParse(taskIdString)

	// Calling the sql query
	rowsAffected, err := ApiConfig.DB.DeleteTask(r.Context(), taskId)
	if err != nil {
		log.Println("ERR: while deleting task")
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected == 0 {
		incorrectTaskIdMessage := "ERR: Invalid Task ID. Task does not exist"
		log.Println(incorrectTaskIdMessage)
		utils.WriteJSON(w, http.StatusNotFound, incorrectTaskIdMessage)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Deleted task successfully!")
}
