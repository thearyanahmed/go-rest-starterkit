package user


import (
	"github.com/thearyanahmed/kloudlabllc/utility"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name,omitempty" bson:"name,omitempty"`
	Email     string        `json:"email,omitempty" bson:"email,omitempty"`
	Password  string        `json:"-" bson:"password,omitempty"`
	Salt      string        `json:"-" bson:"salt,omitempty"`
	Role      string        `json:"role,omitempty" bson:"role,omitempty"`
	IsActive  bool          `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAT int64         `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAT int64         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DeletedAt int64         `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	IsActive bool   `json:"isActive,omitempty" bson:"isActive,omitempty"`
}

// ComparePassword , used to compared
// hashed password with input text password
// return error if any otherwise nil
func (u *User) ComparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

// Initialize , will set the hashed password, createdAt and updatedAt
// date in milliseconds
func (u *User) Initialize() error {
	salt := uuid.New().String()
	passwordBytes := []byte(u.Password + salt)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash[:])
	u.Salt = salt
	u.CreatedAT = utility.CurrentTimeInMilli()
	u.UpdatedAT = utility.CurrentTimeInMilli()
	u.IsActive = true
	u.Role = utility.UserRole
	return nil
}
