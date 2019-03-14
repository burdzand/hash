package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/burdzand/hash/server/cfg"
	"github.com/burdzand/hash/server/discount"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/burdzand/hash/server/connections/database"
	"github.com/burdzand/hash/server/models"
)

var maxDiscountAllow float32 = 10

func main() {
	srv := grpc.NewServer()
	var server Server
	discount.RegisterDiscountServiceServer(srv, server)
	ip := *cfg.GetString("api/bind", "0.0.0.0:3000")
	l, err := net.Listen("tcp", ip)
	if err != nil {
		log.Fatalf("could not listen to %v %v", ip, err)
	}
	log.Fatal(srv.Serve(l))
}

type Server struct{}

func (Server) Get(ctx context.Context, request *discount.DiscountResquest) (*discount.Discount, error) {
	DB := database.GetConn()
	// Get Product
	product := models.Product{}
	if DB.First(&product, request.ProductID).RecordNotFound() {
		return &discount.Discount{}, fmt.Errorf("Product Not Found")
	}

	// Get User
	user := models.User{}
	if DB.First(&user, request.UserID).RecordNotFound() {
		return &discount.Discount{}, fmt.Errorf("User Not Found")
	}
	return getDiscount(&user, &product), nil
}

func getDiscount(user *models.User, product *models.Product) *discount.Discount {
	log.Printf("Discount Calculator user id: %v product id: %v ", user.ID, product.ID)
	var pct float32

	if user.CheckIsBirthDay() {
		pct += 5
	}

	if checkBlackFriday(time.Now()) {
		pct += 10
	}

	if pct > maxDiscountAllow {
		pct = maxDiscountAllow
	}

	price := float32(product.Price_in_cents) / 100
	priceDiscounted := price - (price * (pct / 100))

	return &discount.Discount{
		Pct:          pct,
		ValueInCents: int32(priceDiscounted * 100),
	}
}

func checkBlackFriday(date time.Time) bool {
	return date.Day() == 25 && date.Month() == 11
}
