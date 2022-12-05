package delivery

import (
	"context"
	"eCommerce/internal/delivery/product"
	product3 "eCommerce/internal/repository/product"
	product2 "eCommerce/internal/usecase/product"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Routes(r *gin.Engine) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Urutran program

	repo := product3.NewRepo(conn)       // Dari router ambil db dengan repo
	usecase := product2.NewUseCase(repo) // Olah repo dengan usecase

	// Routing diisi dengan fungsi GetProducts tanpa parameter
	r.GET("/products", product.GetProducts(usecase)) // Panggil fungsi GetProducts yang di delivery

	// Routing diisi dengan fungsi GetProduct dengan parameter id
	r.GET("/products/:id", product.GetProduct(usecase)) // Panggil fungsi GetProduct yang di delivery

	// Routing diisi dengan fungsi GetCartProduct tanpa parameter
	r.GET("/cart", product.GetCartProducts(usecase)) // Panggil fungsi GetCartProduct yang di delivery

	// Routing diisi dengan fungsi AddToCart dengan parameter id dan buy_quantity
	r.GET("/add_cart", product.AddToCart(usecase)) // Panggil fungsi GetCartProduct yang di delivery

	// Routing diisi dengan fungsi RemoveFromCart tanpa parameter
	r.GET("/remove_cart", product.RemoveFromCart(usecase)) // Panggil fungsi GetCartProduct yang di delivery

	// Routing diisi dengan fungsi Buy tanpa parameter
	r.GET("/buy", product.Buy(usecase)) // Panggil fungsi GetCartProduct yang di delivery

}
