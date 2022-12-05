package main

type User struct {
	Name string
	Age  int
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "apa kabar!")
// }

// func main() {
// 	r := echo.New()

// 	r.GET("/", func(ctx echo.Context) error {
// 		data := "Hello from /index"
// 		return ctx.String(http.StatusOK, data)
// 	})

// 	r.GET("/page1", func(ctx echo.Context) error {
// 		name := ctx.QueryParam("name")
// 		data := fmt.Sprintf("Hello %s", name)
// 		return ctx.String(http.StatusOK, data)
// 	})

// 	r.POST("/page4", func(ctx echo.Context) error {
// 		name := ctx.FormValue("name")
// 		message := ctx.FormValue("message")
// 		data := fmt.Sprintf(
// 			"Hello %s, I have message for you: %s",
// 			name,
// 			strings.Replace(message, "/", "", 1),
// 		)
// 		return ctx.String(http.StatusOK, data)
// 	})

// 	r.Start(":8080")
// }
