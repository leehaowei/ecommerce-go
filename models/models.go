package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName      *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=30"`
	LastName       *string            `bson:"last_name" json:"last_name" validate:"required,min=2,max=30"`
	Password       *string            `bson:"password" json:"password" validate:"required,min=6"`
	Email          *string            `bson:"email" json:"email" validate:"required"`
	Phone          *string            `bson:"phone" json:"phone" validate:"required"`
	Token          *string            `bson:"token" json:"token"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
	UserID         string             `bson:"user_id" json:"user_id"`
	UserCart       []ProductUser      `bson:"usercart" json:"usercart"`
	AddressDetails []Address          `bson:"address" json:"address"`
	OrderStatus    []Order            `bson:"orders" json:"orders"`
}

type Product struct {
	ProductID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductName *string            `bson:"product_name" json:"product_name"`
	Price       *uint64            `bson:"price" json:"price"`
	Rating      *uint8             `bson:"rating" json:"rating"`
	Image       *string            `bson:"image" json:"image"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductName *string            `bson:"product_name" json:"product_name"`
	Price       int                `bson:"price" json:"price"`
	Rating      *uint8             `bson:"rating" json:"rating"`
	Image       *string            `bson:"image" json:"image"`
}

type Address struct {
	AddressID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	House     *string            `bson:"house_name" json:"house_name"`
	Street    *string            `bson:"street_name" json:"street_name"`
	City      *string            `bson:"city_name" json:"city_name"`
	Pincode   *string            `bson:"pin_code" json:"pin_code"`
}

type Order struct {
	OrderID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OrderCart     []ProductUser      `bson:"order_list" json:"order_list"`
	OrderedAt     time.Time          `bson:"ordered_at" json:"ordered_at"`
	Price         int                `bson:"price" json:"price"`
	Discount      *int               `bson:"discount" json:"discount"`
	PaymentMethod Payment            `bson:"payment_method" json:"payment_method"`
}

type Payment struct {
	Digital bool `bson:"" json:""`
	COD     bool `bson:"" json:""`
}
