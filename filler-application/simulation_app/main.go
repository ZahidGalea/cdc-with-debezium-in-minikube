package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func randate() string {
	min := time.Date(2022, 4, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2022, 5, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02")
}

func ranNumber(max int, min int) int {
	delta := max - min
	number := rand.Intn(delta) + min
	return number
}

func ranTamanio() string {
	lista := []string{"grande", "mediano", "pequeño", "Imposibol"}
	number := ranNumber(3, 1)
	return lista[number]
}

func main() {
	message := Hello("People")
	fmt.Println(message)
	for ok := true; ok; {
		postBody, _ := json.Marshal(map[string]string{
			"fecha_envio":     randate(),
			"costo_envio":     strconv.Itoa(ranNumber(150, 100)),
			"direccion":       randomdata.Address(),
			"comuna":          "US",
			"nombre_apellido": fmt.Sprintf("%v %v", randomdata.FirstName(randomdata.Male), randomdata.FirstName(randomdata.Female)),
			"numero_telefono": randomdata.PhoneNumber(),
			"rut":             "123412312",
			"region":          strconv.Itoa(ranNumber(4, 1)),
			"peso":            strconv.Itoa(ranNumber(20, 1)),
			"tamanio":         ranTamanio(),
		})

		responseBody := bytes.NewBuffer(postBody)
		fmt.Println(responseBody)

		_, err := http.Post("http://filler-app-svc:8080/api/v1/registrar-venta", "application/json", responseBody)
		if err != nil {
			log.Fatalln(err)
		}

		time.Sleep(3 * time.Second)
	}

}
