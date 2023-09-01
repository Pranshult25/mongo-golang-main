package router

import (
	"context"
	"fmt"
	"regexp"
	"unicode"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pranshult25/queriesportalbackend/common"
	"github.com/pranshult25/queriesportalbackend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/mgo.v2/bson"
)

// var tokenString, _ = makeToken()
// var secret = "secret123"

func Router(app *fiber.App) {
	router := app

	router.Get("/", home)
	router.Post("/register", insertOneUser)
	router.Get("/user", findusers)
	router.Post("/login", Login)
	router.Post("/logout", logout)
	router.Get("/comments", getComments)
	router.Get("commentsbycategory/:category", getCommentsbyCategory)
	router.Get("/comments/root/:rootId", getCommentByRootId)
	router.Get("/comments/:id", getCommentsById)
	router.Post("/comments", postComments)
	router.Post("/votes", postVotes)
	router.Get("/vote/:commentId/:direction", voteCommentDirection)
	router.Get("/comments/user/:username", getUserComments)

}

func checkEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func home(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "Connected",
	})

}

// OK
func insertOneUser(c *fiber.Ctx) error {
	fmt.Println("insertoneuser")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}
	if user.Password != "" {
		bytepassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(bytepassword)
	}

	if user.Email == "" || user.Username == "" {
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
	coll := common.GetDBCollection("users")
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
	cookie.Value = user.Username
	cookie.HTTPOnly = true
	cookie.Secure = true

	c.Cookie(&cookie)

	return c.Status(201).JSON(fiber.Map{
		"result": result,
	})
}

// Ok
func findusers(c *fiber.Ctx) error {
	fmt.Println("findUsers")
	coll := common.GetDBCollection("users")

	// find all bookss
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

// OK
func Login(c *fiber.Ctx) error {
	fmt.Println("login")
	coll := common.GetDBCollection("users")

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	ExistedUser := models.User{}

	err := coll.FindOne(c.Context(), bson.M{"username": user.Username}).Decode(&ExistedUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Can't find the user with this username.",
		})
	}

	passOk := bcrypt.CompareHashAndPassword([]byte(ExistedUser.Password), []byte(user.Password))
	if passOk != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "Invalid Username or Password",
		})
	}
	var cookie fiber.Cookie
	cookie.Name = "token"
	cookie.Value = ExistedUser.Username
	cookie.HTTPOnly = true
	cookie.Secure = true

	c.Cookie(&cookie)

	return c.Status(200).JSON(fiber.Map{
		"status": "Logged-In successfully",
		"data":   ExistedUser,
	})
}

//	func logout(c *fiber.Ctx) error {
//		fmt.Println("logout")
//		var cookie *fiber.Cookie
//		cookie.Value = ""
//		c.Cookie(cookie)
//		return c.Status(200).JSON(fiber.Map{
//			"status": "Successfully logged-out",
//		})
//	} Fixed the Logout issue
func logout(c *fiber.Ctx) error {
	fmt.Println("logout")
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	c.Cookie(cookie)
	return c.Status(200).JSON(fiber.Map{
		"status": "Successfully logged-out",
	})
}

func getComments(c *fiber.Ctx) error {
	coll := common.GetDBCollection("comments")

	search := c.Query("search")
	var filters bson.M

	if search != "" {
		caser := cases.Title(language.English)
		filters = bson.M{
			"body": bson.M{
				"$regex": ".*" + caser.String(search) + ".*",
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
	if comments == nil {
		return c.Status(500).JSON("Null")
	} else {
		return c.Status(200).JSON(comments)
	}

}

func addSpaceBeforeCapital(s string) string {
    if len(s) <= 1 {
        return s
    }

    var result []rune
    result = append(result, rune(s[0])) // Initialize the result with the first letter

    for i := 1; i < len(s); i++ {
        if unicode.IsUpper(rune(s[i])) {
            result = append(result, ' ') // Add a space before the capital letter
        }
        result = append(result, rune(s[i])) // Add the current character
    }
    return string(result)
}



func getCommentsbyCategory(c *fiber.Ctx) error {
	category := c.Params("category")
    refinedCategory := addSpaceBeforeCapital(category)
	fmt.Println(refinedCategory)
	coll := common.GetDBCollection(refinedCategory)
  
	search := c.Query("search")
	var filters bson.M

	if search != "" {
		caser := cases.Title(language.English)
		filters = bson.M{
			"body": bson.M{
				"$regex": ".*" + caser.String(search) + ".*",
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
	if comments == nil {
		return c.Status(500).JSON("Null")
	} else {
		return c.Status(200).JSON(comments)
	}

}

// func escapeRegExp(s string) string {
// 	return regexp.QuoteMeta(s)
// }
// func x(w http.ResponseWriter, r *http.Request) {
//     search := r.URL.Query().Get("search")
// }

// Ok
func postComments(c *fiber.Ctx) error {
	fmt.Println("postcomments")
	username := c.Cookies("token")
	if username == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorised")
	}
	
	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	comment.PostedAt = time.Now()
	comment.Author = username
	coll := common.GetDBCollection("comments")
	category_coll := common.GetDBCollection(comment.Category)

	_, err := category_coll.InsertOne(c.Context(), &comment)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create comment",
			"message": err.Error(),
		})
	}

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

func getCommentByRootId(c *fiber.Ctx) error {
	// fmt.Println("getcommentbyrootid")
	coll := common.GetDBCollection("comments")

	rootId := c.Params("rootId")
	if rootId == "" {
		c.Status(404).JSON(fiber.Map{
			"error": "Please provide an valid root Id",
		})
	}
	realrootId, _ := primitive.ObjectIDFromHex(rootId)
	cur, err := coll.Find(context.Background(), bson.M{"rootId": realrootId}, options.Find().SetSort(bson.M{"postedAt": -1}))
	if err != nil {
		c.Status(404).JSON(fiber.Map{
			"error": "Cannot find an comment with this rootId",
		})
	}
	defer cur.Close(c.Context())

	var comments []models.Comment
	for cur.Next(c.Context()) {
		var comment models.Comment
		if err := cur.Decode(&comment); err != nil {
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

func getCommentsById(c *fiber.Ctx) error {
	// fmt.Println("getcommentbyid")
	coll := common.GetDBCollection("comments")

	Id := c.Params("id")

	if Id == "" {
		c.Status(404).JSON(fiber.Map{
			"error": "Please provide an valid root Id",
		})
	}
	realId, _ := primitive.ObjectIDFromHex(Id)
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	var comments models.Comment

	err := coll.FindOne(c.Context(), bson.M{"_id": realId}).Decode(&comments)
	if err != nil {
		c.Status(404).JSON(fiber.Map{
			"error": "Cannot find an comment with this Id",
		})
	}

	return c.Status(200).JSON(comments)
}

type request struct {
	CommentsIds []string `json:"commentsIds"`
}

func postVotes(c *fiber.Ctx) error {
	// fmt.Println("postVotes")
	coll := common.GetDBCollection("votes")
	// req := c.Body()
	var req request

	if err := c.BodyParser(&req); err != nil {
		fmt.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	fmt.Println(req.CommentsIds)

	username := c.Cookies("token")

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

		if vote.Author == username {
			userVotes[vote.CommentId.Hex()] = vote.Direction
		}
	}
	return c.JSON(fiber.Map{"commentsTotals": commentsTotals, "userVotes": userVotes})
}

// func getVotesFromDatabase(req []byte) ([]models.Vote, error){
//     coll := common.GetDBCollection("votes")

//     filter := bson.M{"commentId": bson.M{"$in": req}}

//     cur, err := coll.Find(context.Background(), filter)
//     if err != nil {
//         return nil, err
//     }
//     defer cur.Close(context.Background())

//     var votes []models.Vote

//     for cur.Next(context.Background()) {
//         var vote models.Vote
//         if err := cur.Decode(&vote); err != nil {
//             return nil, err
//         }
//         votes = append(votes, vote)
//     }

//     if err := cur.Err(); err != nil {
//         return nil, err
//     }

//     return votes, nil
// }

func voteCommentDirection(c *fiber.Ctx) error {
	coll := common.GetDBCollection("votes")
	username := c.Cookies("token")
	// userInfo, err := getUserFromToken(token)
	// if err != nil {
	//     return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	// }

	commentID := c.Params("commentId")
	direction := c.Params("direction")

	// Remove existing votes by the user for the given comment
	_, err := coll.DeleteMany(context.Background(), bson.M{"commentId": commentID, "author": username})
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
		Author:    username,
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

func getUserComments(c *fiber.Ctx) error {
	username := c.Cookies("token")
	if username == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	coll := common.GetDBCollection("comments")

	cur, err := coll.Find(c.Context(), bson.M{"author": username})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer cur.Close(c.Context())

	var userComments []models.Comment

	// Iterate through the query results and decode into Comment objects
	for cur.Next(c.Context()) {
		var comment models.Comment
		if err := cur.Decode(&comment); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		userComments = append(userComments, comment)
	}
	return c.Status(fiber.StatusOK).JSON(userComments)
}
