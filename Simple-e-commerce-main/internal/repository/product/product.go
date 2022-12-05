package product

import (
	"eCommerce/internal/entity"
	"errors"
	"fmt"
	"strconv"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Repo interface {
	GetAll(ctx *gin.Context, page int, limit int) ([]entity.Product, error)
	GetByID(ctx *gin.Context, id int64) (entity.Product, error)
	AddToCart(ctx *gin.Context) (err error)
	RemoveFromCart(ctx *gin.Context) (err error)
	GetCartAll(ctx *gin.Context, page int, limit int) ([]entity.Cart, error)
	Buy(ctx *gin.Context) error
}

type repo struct {
	db *pgx.Conn
}

func NewRepo(db *pgx.Conn) repo {
	return repo{db: db}
}

func (r repo) GetAll(ctx *gin.Context, page int, limit int) ([]entity.Product, error) {
	// Fungsi GetAll yang dipanggil dari usecase
	// Untuk mengambil semua data dari tabel products
	result := make([]entity.Product, 0)

	const query = "SELECT * FROM products"
	err := pgxscan.Select(ctx, r.db, &result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repo) GetByID(ctx *gin.Context, id int64) (result entity.Product, err error) {
	// Fungsi GetByID yang dipanggil dari usecase
	// Untuk mengambil data dengan id tertentu

	const query = "SELECT * FROM products WHERE id=$1;"
	err = pgxscan.Get(ctx, r.db, &result, query, id)
	return result, err
}

func (r repo) AddToCart(ctx *gin.Context) (err error) {
	// Fungsi menambahkan id makanan ke keranjang
	// Quantity barang tidak berubah

	var query = "SELECT * FROM products WHERE id=$1;"
	var result = entity.Product{}

	id, err := strconv.Atoi(ctx.Query("id"))

	if err != nil {
		return err
	}

	err = pgxscan.Get(ctx, r.db, &result, query, id)

	if err != nil {
		return err
	}

	BuyQuantity, err := strconv.Atoi(ctx.Query("buy_quantity"))

	if err != nil {
		return err
	}

	if BuyQuantity > result.Quantity {
		err := errors.New("Buy quantity exceeds product quantity")

		return err
	}

	fmt.Println(result.Name)
	query = "INSERT INTO cart VALUES ($1,$2,$3,$4);"
	fmt.Println(query)

	_, err = r.db.Exec(ctx, query, result.ID, result.Name, result.Price, BuyQuantity)

	return err
}

func (r repo) RemoveFromCart(ctx *gin.Context) (err error) {
	// Fungsi menambahkan id makanan ke keranjang
	// Quantity barang tidak berubah

	var query = "SELECT * FROM products WHERE id=$1;"
	var result = entity.Product{}

	id, err := strconv.Atoi(ctx.Query("id"))

	if err != nil {
		return err
	}
	err = pgxscan.Get(ctx, r.db, &result, query, id)

	if err != nil {
		return err
	}

	query = "DELETE FROM CART WHERE id=$1;"
	_, err = r.db.Exec(ctx, query, id)

	return err
}

func (r repo) GetCartAll(ctx *gin.Context, page int, limit int) ([]entity.Cart, error) {
	// Fungsi GetAll yang dipanggil dari usecase
	// Untuk mengambil semua data dari tabel products
	result := make([]entity.Cart, 0)

	const query = "SELECT * FROM cart"
	err := pgxscan.Select(ctx, r.db, &result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repo) Buy(ctx *gin.Context) error {
	// Fungsi menambahkan id makanan ke keranjang
	// Quantity barang tidak berubah

	var query = "SELECT id,buy_quantity FROM CART;"

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return err
	}

	var cond string = "("

	var id int
	var buy_quantity int

	for rows.Next() {
		err = rows.Scan(&id, &buy_quantity)

		if err != nil {
			return err
		}

		cond += strconv.Itoa(id)
	}

	cond += ");"

	fmt.Println(cond)

	if cond != "();" {
		query = "UPDATE products SET QUANTITY=QUANTITY-$1 WHERE id IN " + cond
		_, err = r.db.Exec(ctx, query, buy_quantity)

		if err != nil {
			return err
		}

		query = "DELETE FROM PRODUCTS WHERE QUANTITY = 0; DELETE FROM CART;"
		_, err = r.db.Exec(ctx, query)

		if err != nil {
			return err
		}

	} else {
		err := errors.New("there is no any products in cart")
		return err
	}

	return err
}
