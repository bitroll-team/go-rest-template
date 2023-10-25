package controller

import (
	"com/example/model"
	"context"
	"time"
)

func (ctrl *Controller) Login(req model.ReqLogin) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := "SELECT id FROM users WHERE username = $1 AND password_hash = $2"

	row := ctrl.DB.QueryRowContext(
		ctx,
		query,
		req.Username,
		req.Password,
	)

	var userId string
	if err := row.Scan(&userId); err != nil {
		return "", err
	}

	return userId, nil
}
