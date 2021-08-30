package seeder

import (
	"github.com/go-goyave/goyave-blog-example/database/model"

	"github.com/bxcodec/faker/v3"
	"goyave.dev/goyave/v4/database"
)

const (
	// UserCount the number of users generated by the User seeder
	UserCount = 10
)

// User seeder for users. Generate and save users in the database.
func User() {
	database.NewFactory(model.UserGenerator).Save(UserCount)

	// As user generator makes unique emails,
	// forget generated unique emails.
	// See https://github.com/bxcodec/faker/blob/master/SingleFakeData.md#unique-values
	faker.ResetUnique()
}
