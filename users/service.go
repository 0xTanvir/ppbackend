package users

import (
	"errors"
	"time"

	"github.com/0xTanvir/pp/db"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const userCollection = "users"

// Service all logic functionality of User
type Service struct {
	DB *db.DB
}

// Create new user
func (r *Service) Create(userInfo *UserInfo) (*bson.ObjectId, error) {

	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(userCollection)

	user := User{ID: bson.NewObjectId(),
		UserInfo: *userInfo,
	}

	return &user.ID, collection.Insert(user)
}

// EmailExist check if a emaol exists
func (r *Service) EmailExist(email string) bool {

	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	var user User
	collection := session.DB("").C(userCollection)
	err := collection.Find(bson.M{"email": email}).One(&user)

	return err == nil
}

// Login a user
func (r *Service) Login(login *Login) (*Ssion, error) {

	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(userCollection)

	var user *User

	err := collection.Find(bson.M{"email": login.Email}).One(&user)
	if err != nil {
		// don't return the real error, this could be used to probe the system
		return nil, errors.New("invalid username or email")
	}

	if user.Password != login.Password {
		return nil, errors.New("invalid username or password")
	}

	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	expiry := time.Now().UTC().Add(time.Hour * 6)

	token.Claims = jwt.MapClaims{
		"name":  user.FirstName,
		"login": true,
		"user":  user.ID.Hex(),
		"exp":   expiry.Unix(),
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("secure-development-key"))
	if err != nil {
		return nil, err
	}

	return &Ssion{&user.ID, tokenString, expiry}, nil
}
