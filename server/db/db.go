package db

import (
	"os"
	"fmt"
	"github.com/bitly/go-simplejson"
	"bufio"
	"encoding/json"
	"errors"
)

var Load = func(name string) (*simplejson.Json, error) {

	file, err := os.Open("server/db/" + name + ".json")

	if err != nil {
		fmt.Println("Error open File", err)
		return nil, err
	}

	j, err := simplejson.NewFromReader(bufio.NewReader(file))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}


	return j, nil
}

func GetSensorIndex(sensorID int) (interface{}, error) {
	j, err := Load("sensors")

	if err != nil {
		return nil, err
	}

	arr, _ := j.Array()

	for sIndex := range arr {
		temp, err := json.Marshal(arr[sIndex])

		if err != nil {
			return nil, err
		}

		sensorSought, err := simplejson.NewJson(temp)

		if err != nil {
			return nil, err
		}

		sensorSoughtID, _ := sensorSought.Get("id").Int()

		if sensorID == sensorSoughtID {
			return sIndex, nil
		}
	}

	return nil, errors.New("Not found Index Sensor")
}

func GetSensor(sensorID int) (*simplejson.Json, error) {
	j, err := Load("sensors")

	if err != nil {
		return nil, err
	}

	arr, _ := j.Array()

	for _, sensor := range arr {
		tmp, err := json.Marshal(sensor)

		if err != nil {
			return nil, err
		}

		jSensor, _ := simplejson.NewJson(tmp)
		jSensorID, _ := jSensor.Get("id").Int()

		if jSensorID == sensorID {
			return jSensor, nil
		}
	}

	return nil, errors.New("Not found Sensor")

}

func GetUser(id int) (*simplejson.Json, error) {
	j, err := Load("users")

	if err != nil {
		return nil, err
	}

	users, _ := j.Array()

	for _, user := range users {
		tempUser, _ := json.Marshal(user)

		jUser, _ := simplejson.NewJson(tempUser)

		if id == jUser.Get("id").MustInt() {
			return jUser, nil
		}
	}

	return nil, errors.New("User not Found")
}