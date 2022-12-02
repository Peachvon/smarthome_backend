package main

import (
	"database/sql"
	"fmt"

	"net/http"

	"github.com/Peachvon/smarthome/data_model"
	"github.com/Peachvon/smarthome/ffun"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func addItem(c *gin.Context) {
	//firstname := c.DefaultQuery("firstname", "Guest")
	model := c.Query("model")
	id := c.Query("id")
	password := c.Query("password") // shortcut for c.Request.URL.Query().Get("lastname")

	db, err := sql.Open("mysql", "Peach:Pe@ch123@tcp(35.240.190.171)/smarthome")

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return

	} else {
		fmt.Println("Database Con")

	}
	defer db.Close()
	var version string = ""
	if model == "1" {
		version = "SELECT * FROM `air` WHERE `id` =\"" + id + "\"AND `password` =\"" + password + "\""
	} else if model == "2" {
		version = "SELECT * FROM `door` WHERE `id` =\"" + id + "\"AND `password` =\"" + password + "\""
	}

	fmt.Println(version)
	data, err := db.Query(version)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return
	}
	defer data.Close()
	if model == "2" {
		var deviceDoor []data_model.DeviceDoor
		for data.Next() {
			var device data_model.DeviceDoor
			err := data.Scan(&device.Id, &device.Passwoed, &device.Model, &device.Topic, &device.Ip, &device.Camera)

			if err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "error",
				})
				return

			}
			deviceDoor = append(deviceDoor, device)
		}
		if len(deviceDoor) == 1 {
			fmt.Println(deviceDoor[0].Id)
			c.JSON(http.StatusOK, gin.H{
				"status":   "success",
				"id":       deviceDoor[0].Id,
				"password": deviceDoor[0].Passwoed,
				"model":    deviceDoor[0].Model,
				"topic":    deviceDoor[0].Topic,
				"ip":       deviceDoor[0].Ip,
				"camera":   deviceDoor[0].Camera,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
			})
			return

		}
	} else if model == "1" {
		var deviceAir []data_model.DeviceAir
		for data.Next() {
			var device data_model.DeviceAir
			err := data.Scan(&device.Id, &device.Passwoed, &device.Model, &device.Topic, &device.Ip)

			if err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "error",
				})
				return

			}
			deviceAir = append(deviceAir, device)
		}
		if len(deviceAir) == 1 {
			fmt.Println(deviceAir[0].Id)
			c.JSON(http.StatusOK, gin.H{
				"status":   "success",
				"id":       deviceAir[0].Id,
				"password": deviceAir[0].Passwoed,
				"model":    deviceAir[0].Model,
				"topic":    deviceAir[0].Topic,
				"ip":       deviceAir[0].Ip,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
			})
			return

		}
	}

}

func main() {

	r := gin.Default()

	fmt.Println(ffun.Asd())
	fmt.Println(ffun.Add2(10, 101))
	r.GET("/add_item", addItem)
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
