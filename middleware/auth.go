package middleware

import (
	"context"
	"errors"
	"os"
	"procurement-be/utils"
	"strings"
)

func Authentication(ctx context.Context, operations string) error {
	authHeader := ctx.Value("Authorization").(string)
	if authHeader == "" {
		return errors.New("header authorization cannot be empty")
	}
	token := strings.Split(authHeader, " ")
	if len(token) == 2 {
		authToken := token[1]
		authorized, err := utils.IsAuthorized(authToken, os.Getenv("ACCESS_TOKEN_SECRET"))
		if err != nil || !authorized {
			return errors.New(err.Error())
		}
	} else {
		return errors.New("invalid yoken")
	}

	// only check auth
	if operations == "" {
		return nil
	}
	return nil
}
