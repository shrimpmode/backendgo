package handlers

import "webserver/models"

func CanCreateChat(user *models.User, server *models.Server) bool {
	return server.OwnerID == user.ID
}
