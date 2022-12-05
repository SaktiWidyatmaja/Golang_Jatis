// package main

// import (
// 	"net/http"
// 	"time"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/labstack/echo/v4"
// )

// // jwtCustomClaims are custom claims extending default ones.
// // See https://github.com/golang-jwt/jwt for more examples
// type jwtCustomClaims struct {
// 	Name  string `json:"name"`
// 	Admin bool   `json:"admin"`
// 	jwt.StandardClaims
// }

// func login(c echo.Context) error {
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")

// 	// Throws unauthorized error
// 	if username != "jon" || password != "shhh!" {
// 		return echo.ErrUnauthorized
// 	}

// 	// Set custom claims
// 	claims := &jwtCustomClaims{
// 		"Jon Snow",
// 		true,
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
// 		},
// 	}

// 	// Create token with claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Generate encoded token and send it as response.
// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"token": t,
// 	})
// }

// func accessible(c echo.Context) error {
// 	return c.String(http.StatusOK, "Accessible")
// }

// func restricted(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*jwtCustomClaims)
// 	name := claims.Name
// 	return c.String(http.StatusOK, "Welcome "+name+"!")
// }

// // func main() {
// // 	e := echo.New()

// // 	// Middleware
// // 	e.Use(middleware.Logger())
// // 	e.Use(middleware.Recover())

// // 	// Login route
// // 	e.POST("/login", login)

// // 	// Unauthenticated route
// // 	e.GET("/", accessible)

// // 	// Restricted group
// // 	r := e.Group("/restricted")

// // 	// Configure middleware with the custom claims type
// // 	config := middleware.JWTConfig{
// // 		Claims:     &jwtCustomClaims{},
// // 		SigningKey: []byte("secret"),
// // 	}
// // 	r.Use(middleware.JWTWithConfig(config))
// // 	r.GET("", restricted)

// // 	e.Logger.Fatal(e.Start(":1323"))
// // }
