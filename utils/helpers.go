package utils

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// CopyObject ... copy struct to struct
func CopyObject(from interface{}, to interface{}) error {
	if from == nil {
		to = nil
		return nil
	}

	return copier.Copy(to, from)
}

// HashAndSalt... to encrypt password
func HashAndSalt(password []byte) string {

	// Use GenerateFromPassword to hash & salt password.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// ComparePasswords... to compare password given and encrypted password
func ComparePasswords(hashedPassword string, password []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// ConvStringToInt... to convert string to int data type
func ConvStringToInt(data string) int {
	a, b := strconv.Atoi(data)
	if b != nil {
		return 0
	}
	return a
}

// GetCollection... to get db collection base on collection name given
func GetCollection(name string, db *mongo.Client) *mongo.Collection {
	return db.Database(os.Getenv("DATABASE_NAME")).Collection(name)
}

// RandomChar... generate random string base on length, camelCase and prefix given
func RandomChar(length int, camelCase bool, prefix string) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,/?!@#$%&"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 0)
	if prefix != "" {
		prefix += "-"
		b = append(b, []byte(prefix)...)
	}
	for i := 0; i < length; i++ {
		tmp := charset[seededRand.Intn(len(charset))]
		if camelCase {
			if i%2 == 0 {
				lower := strings.ToLower(string(tmp))
				b = append(b, []byte(lower)...)
			} else {
				upper := strings.ToUpper(string(tmp))
				b = append(b, []byte(upper)...)
			}
		} else {
			b = append(b, tmp)
		}
	}
	return string(b)
}
