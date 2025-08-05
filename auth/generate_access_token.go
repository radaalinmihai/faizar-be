package auth

import (
	"be/database/entities"
	"encoding/json"
)

func GenerateAccessToken(user entities.User) map[string]interface{} {
	accessTokenPayload := &AccessToken{Username: user.Username, Email: user.Email, ID: user.ID}
	var accessTokenInterface map[string]interface{}
	inrec, _ := json.Marshal(accessTokenPayload)
	json.Unmarshal(inrec, &accessTokenInterface)
	return accessTokenInterface
}
