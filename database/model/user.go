package model

import (
	"reflect"
	"time"

	"github.com/System-Glitch/goyave/v3/config"
	"github.com/System-Glitch/goyave/v3/database"
	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

func init() {
	database.RegisterModel(&User{})

	config.Register("app.bcryptCost", config.Entry{
		Value:            10,
		Type:             reflect.Int,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})
}

// User represents a user.
type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string      `gorm:"type:char(100);unique;unique_index;not null"`
	Email     string      `gorm:"type:char(100);unique;unique_index;not null" auth:"username"`
	Image     null.String `gorm:"type:char(100);default:null"` // TODO file storage
	Password  string      `gorm:"type:char(60);not null" auth:"password" model:"hide" json:",omitempty"`
}

// BeforeCreate hook executed before a User record is inserted in the database.
// Ensures the password is encrypted using bcrypt, with the cost defined by the
// config entry "app.bcryptCost".
func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.brcyptPassword(tx)
}

// BeforeUpdate hook executed before a User record is updated in the database.
// Ensures the password is encrypted using bcrypt, with the cost defined by the
// config entry "app.bcryptCost".
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("Password") {
		return u.brcyptPassword(tx)
	}

	return nil
}

func (u *User) brcyptPassword(tx *gorm.DB) error {
	var newPass string
	if v, ok := tx.Statement.Dest.(map[string]interface{}); ok {
		newPass = v["password"].(string)
	} else {
		newPass = tx.Statement.Dest.(*User).Password
	}
	b, err := bcrypt.GenerateFromPassword([]byte(newPass), config.GetInt("app.bcryptCost"))
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("password", b)
	return nil
}

// UserGenerator generator function for the User model.
// Generate users using the following:
//  database.NewFactory(model.UserGenerator).Generate(5)
func UserGenerator() interface{} {
	user := &User{}
	user.Username = faker.Name()

	faker.SetGenerateUniqueValues(true)
	user.Email = faker.Email()
	faker.SetGenerateUniqueValues(false)

	user.Username = faker.Name()
	return user
}
