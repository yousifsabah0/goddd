package mongo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yousifsabah0/goddd/aggregates"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	mongo    *mongo.Database
	customer *mongo.Collection
}

type mongoCustomer struct {
	Id   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c aggregates.Customer) mongoCustomer {
	return mongoCustomer{
		Id:   c.GetId(),
		Name: c.GetName(),
	}
}

func (m mongoCustomer) ToAggregate() aggregates.Customer {
	c := aggregates.Customer{}

	c.SetId(m.Id)
	c.SetName(m.Name)

	return c
}

func New(ctx context.Context, uri string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database("goddd")
	customers := db.Collection("customers")

	return &MongoRepository{
		mongo:    db,
		customer: customers,
	}, nil
}

func (mr *MongoRepository) Get(id uuid.UUID) (aggregates.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.customer.FindOne(ctx, bson.M{"id": id})
	var c mongoCustomer
	err := result.Decode(c)
	if err != nil {
		return aggregates.Customer{}, err
	}

	return c.ToAggregate(), nil
}

func (mr *MongoRepository) Add(c aggregates.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)
	if _, err := mr.customer.InsertOne(ctx, internal); err != nil {
		return err
	}

	return nil
}

func (mr *MongoRepository) Update(c aggregates.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)
	if _, err := mr.customer.UpdateOne(ctx, bson.M{"id": c.GetId()}, internal); err != nil {
		return err
	}

	return nil
}
