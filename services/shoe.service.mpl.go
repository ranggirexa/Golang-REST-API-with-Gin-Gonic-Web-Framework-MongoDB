package services

import (
	"context"
	"errors"

	"example.com/sarang-apis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShoeServiceImpl struct {
	shoecollection *mongo.Collection
	ctx            context.Context
}

func NewShoeService(shoecollection *mongo.Collection, ctx context.Context) ShoeService {
	return &ShoeServiceImpl{
		shoecollection: shoecollection,
		ctx:            ctx,
	}
}

func (u *ShoeServiceImpl) CreateShoe(shoe *models.Shoe, user_id *string) error {
	// _, err := u.shoecollection.InsertOne(u.ctx, shoe, bson.D{{Key: "user_id", Value: shoe.User_id}})
	_, err := u.shoecollection.InsertOne(u.ctx, bson.D{{Key: "user_id", Value: user_id}, {Key: "brand", Value: shoe.Brand},
		{Key: "size", Value: shoe.Size},
	})
	return err
}

func (u *ShoeServiceImpl) GetShoe(brand *string) (*models.Shoe, error) {
	var shoe *models.Shoe
	query := bson.D{bson.E{Key: "brand", Value: brand}}
	err := u.shoecollection.FindOne(u.ctx, query).Decode(&shoe)
	return shoe, err
}

func (u *ShoeServiceImpl) GetAll() ([]*models.Shoe, error) {
	var shoes []*models.Shoe
	cursor, err := u.shoecollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var shoe models.Shoe
		err := cursor.Decode(&shoe)
		if err != nil {
			return nil, err
		}
		shoes = append(shoes, &shoe)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(shoes) == 0 {
		return nil, errors.New("documents not found")
	}
	return shoes, nil
}

func (u *ShoeServiceImpl) UpdateShoe(shoe *models.Shoe) error {
	filter := bson.D{primitive.E{Key: "brand", Value: shoe.Brand}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "brand", Value: shoe.Brand}, primitive.E{Key: "size", Value: shoe.Size}}}}
	result, _ := u.shoecollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *ShoeServiceImpl) DeleteShoe(brand *string) error {
	filter := bson.D{primitive.E{Key: "brand ", Value: brand}}
	result, _ := u.shoecollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
