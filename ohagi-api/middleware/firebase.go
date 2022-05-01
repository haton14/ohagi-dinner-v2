package middleware

import (
	"context"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func FirebaseAuth() echo.MiddlewareFunc {
	return firebaseAuth
}

func firebaseAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := option.WithCredentialsFile("./serviceAccountKey.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return err
		}
		client, err := app.Auth(context.Background())
		if err != nil {
			return err
		}
		auth := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(auth, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return err
		}
		c.Set("firebase_user_id", token.UID)
		return next(c)
	}
}
