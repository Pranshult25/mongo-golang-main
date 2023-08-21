package router

import (
	"context"
	// "encoding/json"
	"fmt"
	// "io"
	// "io/ioutil"
	"log"

	// "go/token"
	// "log"

	// "os/user"
	// "sort"

	// "net/http"
	"regexp"
	"time"

	// "github.com/gin-gonic/gin"
	// "github.com/go-delve/delve/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pranshult25/queriesportalbackend/common"
	"github.com/pranshult25/queriesportalbackend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	// "honnef.co/go/tools/config"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/gin-gonic/contrib/sessions"
	// "golang.org/x/oauth2"
)

var tokenString, _ = makeToken()
var secret = "secret123"


func Router(app *fiber.App){
	router := app
    
    
    // router.HandleFunc("/", controllers.Home).Methods("GET")
    // router.HandleFunc("/register", controllers.Register).Methods("POST")
    // router.HandleFunc("/user", controllers.User).Methods("GET")
    // router.HandleFunc("/login", controllers.Login).Methods("POST")
    // router.HandleFunc("/logout", controllers.Logout).Methods("POST")
    // router.HandleFunc("/comments", controllers.FindComments).Methods("GET")
    // router.HandleFunc("/comments/root/{rootId}", controllers.FindCommentsByRootId).Methods("GET")
    // router.HandleFunc("/comments/{id}", controllers.FindCommentsById).Methods("GET")
    // router.HandleFunc("/comments", controllers.PostComments).Methods("POST")

    router.Get("/", home)
    router.Post("/register", insertOneUser)
    router.Get("/user", findusers)
    router.Post("/login", login)
    router.Post("/logout", logout)
    router.Get("/comments", getComments)
    router.Get("/comments/root/:rootId", getCommentByRootId)
    router.Get("/comments/:id", getCommentsById)
    router.Post("/comments", postComments)
    router.Post("/votes", postVotes)
    router.Get("/vote/:commentId/:direction", voteCommentDirection)
}
// type test_user struct {
//     id int
//     username string
//     role string
// }
func makeToken() (string, error){
    // var user models.User

    // token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)
	// claims["sub"] = 1
	// claims["name"] = "token"
	// claims["admin"] = true
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// tokenString, err := token.SignedString([]byte(secret))
    // if err != nil{
    //     log.Fatal(err.Error())
    // }
    
    // fmt.Println(tokenString)
    // return tokenString

    // "email":      "test123@email.com",
    //     "exp":        time.Now().Add(time.Hour * 240).Unix(),
    //     //"exp":        time.Now().Add(-time.Hour * 8).Unix(),
    //     "role":       "testrole",
    //     "name":       "testname",
    //     "ip":         "8.8.8.8",
    //     "user_agent": "testagent",
    //     "id":         "120",


    token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{})

    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString([]byte(secret))

    return tokenString, err
}

func checkEmail(email string) bool{
    Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

//ok
func getUserFromToken(tokenString string) (models.User, error){
    coll := common.GetDBCollection("Users")
	// claims := jwt.MapClaims{}
    // claims.VerifyExpiresAt(time.Now().Add(time.Hour*24).Unix(), true)
	// _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(secret), nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var user models.User

	// userID := claims["id"].(string)
	// filter := bson.M{"_id": userID}

    // err = coll.FindOne(context.Background(), filter).Decode(user)

	// if err != nil{
	// 	log.Fatal(err)
	// }

	// return &user, nil
    
    // fmt.Println("hello2")
    claims := jwt.MapClaims{}
    // fmt.Println("hello1")

    token, err := jwt.ParseWithClaims(tokenString, claims,func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
        if err != nil{
            log.Fatal(err.Error())
        }
        // fmt.Println("hello2")
    
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
        userID := claims["id"].(string)
        var user models.User
        Id, _ := primitive.ObjectIDFromHex(userID)
        coll.FindOne(context.Background(), bson.M{"_id": Id}).Decode(&user)
        // fmt.Println("all sorted")
        return user, nil
    } else {
        return models.User{}, err
    }
    
}

//Ok
func home(c *fiber.Ctx) error{
    return c.Status(200).JSON(fiber.Map{
        "status": "Connected",
    })
    
}

//OK
func insertOneUser(c *fiber.Ctx) error{
    fmt.Println("insertoneuser")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid body",
        })
	}
    if user.Password != ""{
    bytepassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(bytepassword)
    }

    if user.Email == "" || user.Username == ""{
        return c.Status(400).JSON(fiber.Map{
            "error": "invalid details",
        })
    }

    if !checkEmail(user.Email) {
        return c.Status(500).JSON(fiber.Map{
            "error": "Enter a valid email-id",
        })
    }

	// create the book
	coll := common.GetDBCollection("Users")
	result, err := coll.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create book",
			"message": err.Error(),
		})
	}

	// return the book
    
    var cookie fiber.Cookie
    cookie.Name = "token"
    cookie.Value = tokenString
    cookie.HTTPOnly = true
    cookie.Secure = true
    
    c.Cookie(&cookie)
    

    return c.Status(201).JSON(fiber.Map{
        "result": result,
    })
}

//Ok
func findusers(c *fiber.Ctx) error{
    fmt.Println("findUsers")
    coll := common.GetDBCollection("Users")

	// find all books
	users := make([]models.User, 0)
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		user := models.User{}
		err := cursor.Decode(&user)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		users = append(users, user)
	}

	return c.Status(200).JSON(fiber.Map{"data": users})
}

//OK
func login(c *fiber.Ctx) error{
    fmt.Println("login")
    coll := common.GetDBCollection("Users")
    
    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid body",
        })
    }

    existedUser := models.User{}

    err := coll.FindOne(c.Context(), bson.M{"username": user.Username}).Decode(&existedUser)
    if err != nil{
        return c.Status(500).JSON(fiber.Map{
            "status": "Can't find the user with this username.",
        })
    } 
    
    passOk := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(user.Password))
    if passOk != nil{
        return c.Status(404).JSON(fiber.Map{
            "status": "Invalid Username or Password",
        })
    } 

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": existedUser.Id,
    })
    tokenString_login, err := token.SignedString([]byte(secret))
    if err != nil{
        return c.Status(400).JSON(fiber.Map{
            "status": "Can't convert to a tokenString",
        })
    }

    var cookie fiber.Cookie
    cookie.Name = "token"
    cookie.Value = tokenString_login
    cookie.HTTPOnly = true
    cookie.Secure = true
    
    c.Cookie(&cookie)

    return c.Status(200).JSON(fiber.Map{
        "status": "Logged-In successfully",
        "data": existedUser,
    }) 
}


func logout(c *fiber.Ctx) error{
    fmt.Println("logout")
    return c.Status(200).JSON(fiber.Map{
        "status": "Successfully logged-out",
    })
}

func getComments(c *fiber.Ctx) error{
    fmt.Println("getcomments")
    coll := common.GetDBCollection("Comments")

//     search := string(c.Request().URI().QueryString())
//     var filters map[string]interface{}
// 	if search != "" {
// 		filters = map[string]interface{}{
// 			"body": regexp.MustCompile(".*" + search + ".*"),
// 		}
// 	} else {
// 		filters = map[string]interface{}{
// 			"rootId": nil,
// 		}
// }

//     comments, err := coll.Find(c.Context(), filters)
//     if err != nil{
//         return c.Status(500).JSON(fiber.Map{
//             "Status": "No comments were found.",
//         })
//     }

//     return c.Status(200).JSON(fiber.Map{
//         "data" : comments,
//     })

    search := c.Query("search")
    var filters bson.M

    if search != "" {
        filters = bson.M{
            "body": bson.M{
                "$regex": ".*" + search + ".*",
            },
        }
    } else {
        filters = bson.M{"rootId": nil}
    }

    cur, err := coll.Find(context.Background(), filters)
    if err != nil {
        fmt.Println("error finding comments\n", err)
        return c.Status(500).SendString(err.Error())
    }
    defer cur.Close(context.Background())

    var comments []models.Comment
    for cur.Next(context.Background()) {
        var comment models.Comment
        err := cur.Decode(&comment)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
        }
        comments = append(comments, comment)
    }
    
    if comments == nil{
        return c.Status(500).JSON("Null")
    } else{
    return c.Status(200).JSON(comments)
    }
}

// func escapeRegExp(s string) string {
// 	return regexp.QuoteMeta(s)
// }
// func x(w http.ResponseWriter, r *http.Request) {
//     search := r.URL.Query().Get("search")
// }

//Ok
func postComments(c *fiber.Ctx) error{
    fmt.Println("postcomments")
    tokenString := c.Cookies("token")
    if tokenString == ""{
        return c.Status(fiber.StatusUnauthorized).SendString("Unauthorised")
    }
    // fmt.Println("hello")
    userInfo, err := getUserFromToken(tokenString)
    // fmt.Println("ok")
    if err != nil{
        c.Status(500).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    // fmt.Println("ok2")
    coll := common.GetDBCollection("Comments")
    // fmt.Println("ok2")
    var comment models.Comment
    if err := c.BodyParser(&comment); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid body",
        })
	}
    
    // fmt.Println("ok3")
    comment.PostedAt = time.Now()
    // fmt.Println("lol")
    comment.Author = userInfo.Username
    // fmt.Println("ok4")
    _, err = coll.InsertOne(c.Context(), &comment)


    if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create comment",
			"message": err.Error(),
		})
	} 
    fmt.Println("all sorted postComments")
    return c.Status(200).JSON(fiber.Map{
        "data": comment,
    })
}

func getCommentByRootId(c *fiber.Ctx) error{
    // fmt.Println("getcommentbyrootid")
    coll := common.GetDBCollection("Comments")

    rootId := c.Params("rootId")
    if rootId == ""{
        c.Status(404).JSON(fiber.Map{
            "error": "Please provide an valid root Id",
        })
    }   
    realrootId, _ := primitive.ObjectIDFromHex(rootId)
    cur, err := coll.Find(context.Background(), bson.M{"rootId": realrootId}, options.Find().SetSort(bson.M{"postedAt": -1}))
    if err != nil{
        c.Status(404).JSON(fiber.Map{
            "error": "Cannot find an comment with this rootId",
        })
    }
    defer cur.Close(c.Context())

    var comments []models.Comment
    for cur.Next(c.Context()) {
        var comment models.Comment
        if err := cur.Decode(&comment); err!=nil{
            return c.JSON(err.Error())
        }
        comments = append(comments, comment)
    }

    // sort.Slice(comments, func(i, j int) bool {
	// 	return (comments[i].PostedAt.Hour() > comments[j].PostedAt.Hour())
	// })

    return c.Status(200).JSON(fiber.Map{
        "data": comments,
    })
}

func getCommentsById(c *fiber.Ctx) error{
    // fmt.Println("getcommentbyid")
    coll := common.GetDBCollection("Comments")

    Id := c.Params("id")

    if Id == ""{
        c.Status(404).JSON(fiber.Map{
            "error": "Please provide an valid root Id",
        })
    }
    realId, _ := primitive.ObjectIDFromHex(Id)
    // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

    var comments models.Comment

    err := coll.FindOne(c.Context(), bson.M{"_id": realId}).Decode(&comments)
    if err != nil{
        c.Status(404).JSON(fiber.Map{
            "error": "Cannot find an comment with this Id",
        })
    }

    return c.Status(200).JSON(comments)
}

type request struct {
    CommentsIds []string `json:"commentsIds"`
}

func postVotes(c *fiber.Ctx) error{
    // fmt.Println("postVotes")
    coll := common.GetDBCollection("Votes")
    // req := c.Body()
    var req request

    if err := c.BodyParser(&req); err != nil {
        fmt.Println("Error parsing request body:", err)
        return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
    }  
    
    fmt.Println(req.CommentsIds)

    token := c.Cookies("token")
    userInfo, err := getUserFromToken(token)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
    }
    
	voteFilter := bson.M{"commentId": bson.M{"$in": req.CommentsIds}}
	cursor, err := coll.Find(c.Context(), voteFilter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer cursor.Close(c.Context())

		votes := []models.Vote{}
		if err := cursor.All(c.Context(), &votes); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

    commentsTotals := make(map[string]int)
	userVotes := make(map[string]int)
	for _, vote := range votes {
		if _, ok := commentsTotals[vote.CommentId.Hex()]; !ok {
			commentsTotals[vote.CommentId.Hex()] = 0
		}
		commentsTotals[vote.CommentId.Hex()] += vote.Direction

		if vote.Author == userInfo.Username {
			userVotes[vote.CommentId.Hex()] = vote.Direction
		}
	}
    return c.JSON(fiber.Map{"commentsTotals": commentsTotals, "userVotes": userVotes})
}

func getVotesFromDatabase(req []byte) ([]models.Vote, error){
    coll := common.GetDBCollection("Votes")

    filter := bson.M{"commentId": bson.M{"$in": req}}

    cur, err := coll.Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }
    defer cur.Close(context.Background())
   
    var votes []models.Vote

    for cur.Next(context.Background()) {
        var vote models.Vote
        if err := cur.Decode(&vote); err != nil {
            return nil, err
        }
        votes = append(votes, vote)
    }

    if err := cur.Err(); err != nil {
        return nil, err
    }
    
    return votes, nil
}

func voteCommentDirection(c *fiber.Ctx) error{
    coll := common.GetDBCollection("Votes")
    token := c.Cookies("token")
    userInfo, err := getUserFromToken(token)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
    }

    commentID := c.Params("commentId")
    direction := c.Params("direction")

    // Remove existing votes by the user for the given comment
    _, err = coll.DeleteMany(context.Background(), bson.M{"commentId": commentID, "author": userInfo.Username})
    if err != nil {
        fmt.Println("Error deleting votes:", err)
        return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
    }

    if direction != "up" && direction != "down" {
        return c.JSON(true)
    }

    NewCommentId, _ := primitive.ObjectIDFromHex(commentID)


    // Create a new vote for the user
    newVote := models.Vote{
        Author:    userInfo.Username,
        Direction: 1, // Default to upvote, you can change it to -1 for downvote
        CommentId: NewCommentId,
    }

    // Save the vote to the database
    _, err = coll.InsertOne(context.Background(), newVote)
    if err != nil {
        fmt.Println("Error inserting vote:", err)
        return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
    }

    return c.JSON(true)

}



// func PostComments(w http.ResponseWriter, r *http.Request){
//     tokenString, err := r.Cookie("token")
//     if err != nil || tokenString == nil{
//         w.WriteHeader(http.StatusUnauthorized)
//         return
//     }

//     userInfo, err := getUserFromToken(tokenString.Value)
//     if err != nil{
//         w.WriteHeader(http.StatusUnauthorized)
//         return
//     }

//     var comment models.Comment
//     err = json.NewDecoder(r.Body).Decode(&comment)
//     if err != nil{
//         w.WriteHeader(http.StatusBadRequest)
//         return
//     }

//     comment.Author = userInfo.Username
//     comment.PostedAt = time.Now()

//     insertOneComment(&comment)

//     json.NewEncoder(w).Encode(comment)
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
