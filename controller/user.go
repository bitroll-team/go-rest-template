package controller

import (
	"com/example/model"
	"context"
	"time"
)

func (ctrl *Controller) Register(req model.ReqRegister) error {

	// write

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := `
		INSERT INTO users (username, password_hash)
		VALUES ($1, $2)
	`

	_, err := ctrl.DB.ExecContext(
		ctx,
		query,
		req.Username,
		req.Password,
	)
	if err != nil {
		return err
	}

	return nil
}
