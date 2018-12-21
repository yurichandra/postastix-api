package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dewadg/postastix-api/db"
)

type userResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func mapUsers(users []*db.User, f func(item *db.User) interface{}) []interface{} {
	output := make([]interface{}, 0)

	for _, user := range users {
		output = append(output, f(user))
	}

	return output
}

// GetAllUsers retrieves users and displays it as JSON.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	payload := mapUsers(userService.Get(), func(item *db.User) interface{} {
		return userResponse{
			ID:        item.ID,
			Name:      item.Name,
			FullName:  item.FullName,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: item.UpdatedAt.Format("2006-01-01 15:04:05"),
		}
	})

	fmt.Println(payload)

	res, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(res)
}
