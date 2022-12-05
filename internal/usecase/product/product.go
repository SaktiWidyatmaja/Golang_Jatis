package product

import (
	"eCommerce/internal/entity"
	"eCommerce/internal/repository/product"

	"github.com/gin-gonic/gin"
)

type ProductUseCase interface {
	GetProducts(ctx *gin.Context, page int, limit int) ([]entity.Product, error)
	GetProduct(ctx *gin.Context, id int64) (entity.Product, error)
	AddToCart(ctx *gin.Context) (err error)
	RemoveFromCart(ctx *gin.Context) (err error)
	GetCartProducts(ctx *gin.Context, page int, limit int) ([]entity.Cart, error)
	Buy(ctx *gin.Context) error
}

type usecase struct {
	repo product.Repo
}

func NewUseCase(repo product.Repo) usecase {
	return usecase{repo: repo}
}

// Fungsi GetProducts yang dipanggil di delivery
func (u usecase) GetProducts(ctx *gin.Context, page int, limit int) ([]entity.Product, error) {
	// Memanggil GetAll dari repo
	return u.repo.GetAll(ctx, page, limit)
}

func (u usecase) GetProduct(ctx *gin.Context, id int64) (entity.Product, error) {
	return u.repo.GetByID(ctx, id)
}

func (u usecase) AddToCart(ctx *gin.Context) (err error) {
	return u.repo.AddToCart(ctx)
}
func (u usecase) RemoveFromCart(ctx *gin.Context) (err error) {
	return u.repo.RemoveFromCart(ctx)
}
func (u usecase) GetCartProducts(ctx *gin.Context, page int, limit int) ([]entity.Cart, error) {
	return u.repo.GetCartAll(ctx, page, limit)
}
func (u usecase) Buy(ctx *gin.Context) error {
	return u.repo.Buy(ctx)
}
