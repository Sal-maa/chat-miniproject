package main

import (
	"chat-miniproject/config"
	"chat-miniproject/contract"
	"chat-miniproject/handler"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func DatabaseMigration() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	db.AutoMigrate(contract.Users{}, contract.Chat{})
}

func main() {
	DatabaseMigration()

	app := fiber.New()

	app.Post("/users", handler.CreateNewUser)
	app.Get("/users", handler.GetAllUser)
	app.Post("/users/:id", handler.SendNewText)
	app.Get("/users/:id", handler.ReadNewText)
	app.Get("/users/:id?name", handler.GetChatRoom)

	app.Listen(":8080")
}

// func main() {

// 	config.NewMysqlDatabase()
// 	config.DB.AutoMigrate(&contract.Users{})
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == "GET" {
// 			res := contract.JSONResponse{
// 				http.StatusOK,
// 				true,
// 				"Success get the chat list",
// 				chatBox,
// 			}

// 			resJSON, err := json.Marshal(res)
// 			if err != nil {
// 				fmt.Println("An error occurred")
// 				http.Error(w, "An error occurred", http.StatusInternalServerError)
// 				return
// 			}

// 			w.Header().Add("Content-Type", "application/json")
// 			w.Write(resJSON)
// 			return
// 		} else if r.Method == "POST" {
// 			jsonDecode := json.NewDecoder(r.Body)
// 			newChat := contract.Chat{}
// 			res := contract.JSONResponse{}

// 			if err := jsonDecode.Decode(&newChat); err != nil {
// 				fmt.Println("An error occurred")
// 				http.Error(w, "An error occurred reading the input", http.StatusInternalServerError)
// 				return
// 			}

// 			res.Code = http.StatusCreated
// 			res.Success = true
// 			res.Massage = "Successfully added chat"
// 			res.Data = newChat

// 			chatBox = append(chatBox, newChat)

// 			resJSON, err := json.Marshal(res)

// 			if err != nil {
// 				fmt.Println("An error occurred")
// 				http.Error(w, "An error occurred parsing to json", http.StatusInternalServerError)
// 				return
// 			}

// 			w.Header().Add("Content-Type", "application/json")
// 			w.Write(resJSON)
// 			return
// 		}
// 	})
// 	fmt.Println("Listening on: 8080...")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
