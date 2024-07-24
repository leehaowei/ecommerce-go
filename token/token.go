package token

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/leehaowei/ecommerce-go/db"
	"github.com/leehaowei/ecommerce-go/errors"
	"github.com/leehaowei/ecommerce-go/models"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Uid        string
	jwt.MapClaims
}

var UserData *mongo.Collection = db.UserData(db.Client, "Users")
var SECRET_KEY = os.Getenv("SECRET_LOVE")

func CreateTokenFromUser(user *models.User) string {
	now := time.Now()
	expires := now.Add(time.Hour * 4).Unix()
	claims := jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expires": expires,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("failed to sign token with secret", err)
	}
	return tokenStr
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, errors.ErrUnauthorized()
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return nil, errors.ErrUnauthorized()
	}
	if !token.Valid {
		fmt.Println("invalid token:")
		return nil, errors.ErrUnauthorized()
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.ErrUnauthorized()
	}
	return claims, nil
}

// func UpdateAllTokens(signedtoken string, signedrefreshtoken string, userid string) {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	var updateobj primitive.D
// 	updateobj = append(updateobj, bson.E{Key: "token", Value: signedtoken})
// 	updateobj = append(updateobj, bson.E{Key: "refresh_token", Value: signedrefreshtoken})
// 	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 	updateobj = append(updateobj, bson.E{Key: "updatedat", Value: updated_at})
// 	upsert := true
// 	filter := bson.M{"user_id": userid}
// 	opt := options.UpdateOptions{
// 		Upsert: &upsert,
// 	}
// 	_, err := UserData.UpdateOne(ctx, filter, bson.D{
// 		{Key: "$set", Value: updateobj},
// 	},
// 		&opt)
// 	defer cancel()
// 	if err != nil {
// 		log.Panic(err)
// 		return
// 	}

// }
