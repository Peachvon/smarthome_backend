package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Peachvon/smarthome/data_model"
	"github.com/gin-gonic/gin"
)

func AddItemToMobile(c *gin.Context) {
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


func SelectAirItem(c *gin.Context) {

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

	sql := "SELECT * FROM air"
	data, err := db.Query(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return
	}
	defer data.Close()
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
	c.JSON(http.StatusOK, deviceAir)

}
func SelectDoorItem(c *gin.Context) {

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

	sql := "SELECT * FROM door"
	data, err := db.Query(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return
	}
	defer data.Close()
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
	c.JSON(http.StatusOK, deviceDoor)

}

func AddAirItem(c *gin.Context) {
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
	var air data_model.DeviceAir
	//fmt.Println(c.ShouldBindJSON(&book))
	if err := c.ShouldBindJSON(&air); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	if air.Id == "" || air.Passwoed == "" || air.Model == "" || air.Topic == "" || air.Ip == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	sql := "INSERT INTO `air` (id, password, model,topic,ip) VALUES (?,?,?,?,?)"
	data, err := db.Query(sql, air.Id, air.Passwoed, air.Model, air.Topic, air.Ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return
	}
	defer data.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success",
	})

}

func AddDoorItem(c *gin.Context) {
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
	var door data_model.DeviceDoor
	//fmt.Println(c.ShouldBindJSON(&book))
	if err := c.ShouldBindJSON(&door); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	if door.Id == "" || door.Passwoed == "" || door.Model == "" || door.Topic == "" || door.Ip == "" || door.Camera == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ข้อมูลไม่ครบ",
		})
		return
	}

	sql := "INSERT INTO `door` (id, password, model,topic,ip,camera) VALUES (?,?,?,?,?,?)"
	data, err := db.Query(sql, door.Id, door.Passwoed, door.Model, door.Topic, door.Ip, door.Camera)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return
	}
	defer data.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success",
	})

}

func DeleteAirItem(c *gin.Context) {
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
	airId := c.Query("id")

	if airId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ข้อมูลไม่ครบ",
		})
		return
	}

	sql := "DELETE FROM air WHERE ID = ?"
	data, err := db.Query(sql, airId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return
	}
	defer data.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "ลบ ID: " + airId + " แล้ว",
	})

}

func DeleteDoorItem(c *gin.Context) {
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
	doorId := c.Query("id")

	if doorId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ข้อมูลไม่ครบ",
		})
		return
	}

	sql := "DELETE FROM door WHERE ID = ?"
	data, err := db.Query(sql, doorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return
	}
	defer data.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "ลบ ID: " + doorId + " แล้ว",
	})

}
