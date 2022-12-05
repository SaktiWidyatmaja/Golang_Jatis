package product

import (
	"eCommerce/internal/usecase/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(useCase product.ProductUseCase) gin.HandlerFunc {
	// Jangan lupa cari tahu tentang return func!!
	// Anonymous function atau fungsi tanpa nama
	return func(c *gin.Context) {
		page := 1
		limit := 10

		// Panggil fungsi GetProducts yang di usecase
		result, err := useCase.GetProducts(c.Copy(), page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetProduct(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		// Panggil fungsi GetProduct yang di usecase
		result, err := useCase.GetProduct(c.Copy(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func AddToCart(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Panggil fungsi AddToCart yang di usecase
		err := useCase.AddToCart(c.Copy())
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, "Product added to cart")
	}
}

func RemoveFromCart(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Panggil fungsi GetProduct yang di usecase
		err := useCase.RemoveFromCart(c.Copy())
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, "Product removed from cart")
	}
}

func GetCartProducts(useCase product.ProductUseCase) gin.HandlerFunc {
	// Jangan lupa cari tahu tentang return func!!
	// Anonymous function atau fungsi tanpa nama
	return func(c *gin.Context) {
		page := 1
		limit := 10

		// Panggil fungsi GetCartProducts yang di usecase
		result, err := useCase.GetCartProducts(c.Copy(), page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func Buy(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Panggil fungsi Buy yang di usecase
		err := useCase.Buy(c.Copy())
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, "Successfully placed the order")
	}
}
