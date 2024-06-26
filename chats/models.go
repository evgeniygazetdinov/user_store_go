package chats

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	created_at       time.Time          `json:"created_at",bson:"created_at"`
	user_ids         string             `json:"user_ids",bson:"user_ids"`
	updated_at       time.Time          `json:"updated_at",bson:"updated_at"`
	is_system        bool               `json:"is_system",bson:"is_system"`
	is_visited_users string             `json:"is_visited_users",bson:"is_visited_users"`
	is_pure          bool               `json:"is_pure",bson:"is_pure"`
	is_active        bool               `json:"is_active",bson:"is_active"`
}

type ListCarQuery struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type ListCarQueryV1 struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func (q *ListCarQuery) Filter(email string) bson.M {
	andFilters := []bson.M{
		bson.M{
			"email": email,
		},
	}

	if q.Make != "" {
		orFilters := []bson.M{
			// Exact match
			bson.M{
				"make": q.Make,
			},
			// Similar match
			bson.M{
				"make": bson.M{
					"$regex": primitive.Regex{
						Pattern: "^" + q.Make + "*",
						Options: "i",
					},
				},
			},
		}

		andFilters = append(andFilters, bson.M{"$or": orFilters})
	}

	if q.Model != "" {
		orFilters := []bson.M{
			// Exact match
			bson.M{
				"model": q.Model,
			},
			// Similar match
			bson.M{
				"model": bson.M{
					"$regex": primitive.Regex{
						Pattern: "^" + q.Model + "*",
						Options: "i",
					},
				},
			},
		}

		andFilters = append(andFilters, bson.M{"$or": orFilters})
	}

	if q.Year != 0 {
		andFilters = append(andFilters, bson.M{"year": q.Year})
	}

	if len(andFilters) == 0 {
		// Handle empty and, since there must be one item.
		return bson.M{}
	}
	return bson.M{"$and": andFilters}
}

type CreateCar struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type CreateCarV1 struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func (c *CreateCar) Valid() error {
	if c.Make == "" {
		return errors.New("Error: Make is missing")
	}
	if c.Model == "" {
		return errors.New("Error: Model is missing")
	}
	if c.Year == 0 {
		return errors.New("Error: Year is missing")
	}
	return nil
}

type UpdateCar struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Status string `json:"status"`
}

func (u *UpdateCar) Valid() error {
	return nil
}

func (u *UpdateCar) Update() bson.M {
	update := bson.M{}
	if u.Make != "" {
		update["make"] = u.Make
	}
	if u.Model != "" {
		update["model"] = u.Model
	}
	if u.Year != 0 {
		update["year"] = u.Year
	}
	if u.Status != "" {
		update["status"] = u.Status
	}
	if len(update) == 0 {
		return nil
	}
	return bson.M{"$set": update}
}

type UpdateCarV1 struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Status string `json:"status"`
}
