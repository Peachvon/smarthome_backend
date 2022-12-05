package main

import (
	"fmt"

	"net/http"

	"github.com/Peachvon/smarthome/api"
	"github.com/Peachvon/smarthome/ffun"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func main() {

	r := gin.Default()

	fmt.Println(ffun.Asd())
	fmt.Println(ffun.Add2(10, 101))
	r.Use(cors.Default())
	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	r.GET("/api/add_item_tomobile", api.AddItemToMobile)
	r.GET("/api/select_air_item", api.SelectAirItem)
	r.GET("/api/select_door_item", api.SelectDoorItem)
	r.POST("/api/add_air_item", api.AddAirItem)
	r.POST("/api/add_door_item", api.AddDoorItem)
	r.DELETE("/api/delete_door_item", api.DeleteDoorItem)
	r.DELETE("/api/delete_air_item", api.DeleteAirItem)

	r.POST("/books", func(c *gin.Context) {

		var book Book
		//fmt.Println(c.ShouldBindJSON(&book))
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Printf("asd %+v \n", book)
		books = append(books, book)

		c.JSON(http.StatusCreated, books)
	})

	r.Run(":3001")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// fmt.Println("When's Saturday?")
	// today := time.Now().Day()
	// fmt.Println(today)

	// mux := http.NewServeMux()
	// //mux.Handle("/kfc", apiHandler{})

	// mux.HandleFunc("/door", ServeHTTP)
	// mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	if req.URL.Path != "/" {

	// 		http.NotFound(w, req)

	// 		return
	// 	}

	// 	fmt.Fprintln(w, "peachvon")
	// })

	// http.ListenAndServe(":3001", mux)

}
