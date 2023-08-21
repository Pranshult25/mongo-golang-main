package common

import (
	"context"
	"errors"
	"os"

	// "encoding/json"
	// "errors"
	// "log"
	// "os"
	// "regexp"
	// "time"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "golang.org/x/crypto/bcrypt"

	// "context"
	// "encoding/json"
	// "net/http"
	// "regexp"
	// "time"

	// "encoding/json"
	"fmt"
	// "log"

	// "mongoapi/model"
	// "net/http"
	// "github.com/gorilla/mux"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"

	// "github.com/gin-gonic/gin"
	// "github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt"
	// "github.com/gorilla/mux"

	// "github.com/gorilla/mux"
	// "github.com/pranshult25/queriesportalbackend/common"
	// "github.com/pranshult25/queriesportalbackend/common"
	// "github.com/pranshult25/queriesportalbackend/models"
	// "go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "golang.org/x/crypto/bcrypt"
	// "go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

var db *mongo.Database

// const connectionString = "mongodb://localhost:27017"
// const dbName = "QueriesPortal"
const col = "Users"
// // const colName2 = "Comments"
// const secret = "secret123"

// var tokenString = makeToken()

// var collection_users *mongo.Collection
// var collection_comments *mongo.Collection

func GetDBCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func InitDB() error {
	uri := os.Getenv("MONGODB_URI")
	fmt.Println(uri)
	if uri == "" {
		return errors.New("you must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("err found")
		return err
	}

	db = client.Database("queriesportal")

	return nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.Background())
}



// func Register(w http.ResponseWriter, r *http.Request){

// 	var user models.User
// 	// user = {
// 	// 	Email: r.FormValue("email"),
// 	// 	Username: r.FormValue("username"),
// 	// 	Password: bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost),
// 	// }

// 	// user.Email = r.FormValue("email")
// 	// user.Username = r.FormValue("username")
// 	// user.Password, _ = bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)

// 	json.NewDecoder(r.Body).Decode(&user)

// 	user.Password, _ = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)


// 	insertOneUser(user)
	
// 	http.SetCookie(w, &http.Cookie{
// 		Name: "token",
// 		Value: tokenString,
// 		HttpOnly: true,
// 		Secure: true,
// 	})

// 	json.NewEncoder(w).Encode(user)

// 	w.WriteHeader(http.StatusCreated)
// }

// func makeToken() (tokenString string){
// 	var user models.User
// 	claims := jwt.MapClaims{
// 		"user": &user,
// 		"exp": time.Now().Add(time.Hour*1).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err  := token.SignedString([]byte(secret))

// 	if err != nil{
// 		log.Fatal(err)
// 	}
    
// 	fmt.Println("JWT", tokenString)

// 	return tokenString
// }

// func getUserFromToken(tokenString string) (*models.User, error){
// 	claims := jwt.MapClaims{}
// 	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(secret), nil
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var user models.User

// 	userID := claims["id"].(string)
// 	filter := bson.M{"_id": userID}

//     err = collection_users.FindOne(context.Background(), filter).Decode(user)

// 	if err != nil{
// 		log.Fatal(err)
// 	}

// 	return &user, nil
// }

// // ///////////// HELPERS
// func insertOneUser(user models.User){
// 	inserted, err := collection_users.InsertOne(context.Background(), user)

// 	if err != nil{
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Inserted user id is: ", inserted.InsertedID)
// }

// func insertOneComment(comment *models.Comment){
// 	inserted, err := collection_comments.InsertOne(context.Background(), comment)

// 	if err != nil{
// 		fmt.Println("The comment can't be inserted")
// 	}

// 	fmt.Println("The inserted id is: ", inserted.InsertedID)

// }

// func Home(w http.ResponseWriter, r *http.Request){
// 	json.NewEncoder(w).Encode(http.StatusOK)
// }


// func User(w http.ResponseWriter, r *http.Request){
// 	tokenString, err := r.Cookie("token")
    
// 	if err != nil{
// 		log.Fatal(err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	user, err := getUserFromToken(tokenString.Value)

// 	if err != nil{
// 		log.Fatal(err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	type response struct{
// 		Username string  `json:"username"`
// 	}
	
// 	var result response
// 	result.Username = user.Username
    
//     json.NewEncoder(w).Encode(result)
// }

// func Login(w http.ResponseWriter, r *http.Request){
// 	var user models.User
	
// 	// username := r.FormValue("username")
// 	// password := r.FormValue("password")
// 	var response models.User
// 	json.NewDecoder(r.Body).Decode(&response)

	

// 	err := collection_users.FindOne(context.Background(), bson.M{"username": response.Username}).Decode(&user)

// 	if err != nil{
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte("Invalid username or password"))
// 		fmt.Println("hello")
// 		return
// 	}

// 	passOk := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(response.Password))
// 	if passOk != nil{
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte("Invalid username or password"))
// 		fmt.Println("hello2")
// 		return
// 	} else{

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id": user.Id,
// 	})
// 	tokenString, err := token.SignedString([]byte(secret))
// 	if err != nil{
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Internal server error"))
// 		return
// 	}

// 	http.SetCookie(w, &http.Cookie{
// 		Name:"token",
// 		Value: tokenString,
// 		HttpOnly: true,
// 		Secure: true,
// 	})
	

// 	json.NewEncoder(w).Encode("You are successfully logedIn.")
// 	json.NewEncoder(w).Encode(user)
	

// 	w.WriteHeader(http.StatusOK)
// }
// }


// func Logout(w http.ResponseWriter, r *http.Request){
// 	http.SetCookie(w, &http.Cookie{
// 		Name: "token",
// 		Value: "",
// 		HttpOnly: true,
// 		Secure: true,
// 	})
// 	w.WriteHeader(http.StatusOK)
// }

// func FindComments(w http.ResponseWriter, r *http.Request){
// 	result := r.URL.Query().Get("search")
// 	var filters map[string]interface{}

// 	if result != "" {
// 		filters = map[string]interface{}{
// 			"body": regexp.MustCompile(".*" + result + ".*"),
// 		}
// 	} else {
// 		filters = map[string]interface{}{
// 			"rootId": nil,
// 		}
// 	}

// 	var comment models.Comment

// 	err := collection_comments.FindOne(context.TODO(), filters).Decode(&comment)																																			
// 	if err != nil {
//         w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Internal server error"))
// 		fmt.Println("Unable to find the comment.")
// 		return
// 	}
// 	json.NewEncoder(w).Encode(comment)
// }

// func FindCommentsByRootId(w http.ResponseWriter, r *http.Request){
// 	params := mux.Vars(r)

// 	rootId := params["rootId"]

// 	var comment models.Comment

// 	err := collection_comments.FindOne(context.TODO(), bson.M{"rootId": rootId}).Decode(&comment)
    
//     if err != nil{
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Println("The following comment with the rootId doesn't exist.")
// 	}
   
//     json.NewEncoder(w).Encode(comment)
// }

// func FindCommentsById(w http.ResponseWriter, r *http.Request){
// 	params := mux.Vars(r)
    
// 	Id := params["id"]

// 	var comment models.Comment

// 	err := collection_comments.FindOne(context.TODO(), bson.M{"parentId":Id}).Decode(&comment)

// 	if err != nil{
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Println("The following comment with the Id doesn't exist.")
// 	}

// 	json.NewEncoder(w).Encode(comment)
// }

// func PostComments(w http.ResponseWriter, r *http.Request){
// 	tokenString, err := r.Cookie("token")
// 	if err != nil || tokenString == nil{
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	userInfo, err := getUserFromToken(tokenString.Value)
//     if err != nil{
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	var comment models.Comment
//     err = json.NewDecoder(r.Body).Decode(&comment)
//     if err != nil{
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	comment.Author = userInfo.Username
// 	comment.PostedAt = time.Now()

// 	insertOneComment(&comment)

// 	json.NewEncoder(w).Encode(comment)
// }
