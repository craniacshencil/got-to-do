package handlers

import (
	"log"
	"net/http"

	"github.com/craniacshencil/got_to_do/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (ApiConfig *ApiCfg) DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	// Extract list_id from URL
	listIdString := chi.URLParam(r, "list_id")
	listId := uuid.MustParse(listIdString)

	// Delete the list using list_id
	rowsAffected, err := ApiConfig.DB.DeleteList(r.Context(), listId)
	if err != nil {
		log.Println("ERR: While deleting the list:", err.Error())
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// When the list_id is not valid, there's no such list in the database
	if rowsAffected == 0 {
		invalidListIdMessage := "Invalid list_id. This list does not exist in database"
		log.Println(invalidListIdMessage)
		utils.WriteJSON(w, http.StatusNotFound, invalidListIdMessage)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "List was deleted successfully")
	return
}
