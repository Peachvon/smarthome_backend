package data_model

type DeviceAir struct {
	Id       string
	Passwoed string
	Model    string
	Ip       string
	Topic    string
}
type DeviceDoor struct {
	Id       string
	Passwoed string
	Model    string
	Ip       string
	Topic    string
	Camera   string
}

func Add(num1, num2 int) int {
	return num1 + num2
}
