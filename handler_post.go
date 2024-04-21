package main

import (
	"fmt"
	"net/http"

	"github.com/mkappworks/rss-aggregator-gobackend/internal/database"
)

func (apiConfig *apiConfig) handlerGetPostsForUsers(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiConfig.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching posts %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
