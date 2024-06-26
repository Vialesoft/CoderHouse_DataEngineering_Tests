package main;

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
);

var peliculas = Peliculas{
	Pelicula{"Soy", 2013, "Pepe"},
	Pelicula{"BatGonza", 2974, "Gonza"},
	Pelicula{"La Onda Vital", 2014, "Agu Ago"},
};

var collection = getSession().DB("Curso_GO").C("Peliculas"); //C = "Collection"

//Conexión a la BD MongoDB
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost");

	if (err != nil) {
		panic(err);
	}

	return session;
}

func Index (response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "HOLA SOY SKYNET VAN A MORIR TODOS");
}

func ListaPeliculas (response http.ResponseWriter, request *http.Request) {
	var resultados []Pelicula;
	err := collection.Find(nil).Sort("-_id").All(&resultados);

	if (err != nil) {
		log.Fatal(err);
	} else {
		fmt.Println("Resultados: %s", resultados)
	}

	respuestaPeliculas(200, response, resultados);
}

func MostrarPelicula (response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request);
	movie_id := params["id"];

	if (!bson.IsObjectIdHex(movie_id)) { //Si no es hexadecimal
		response.WriteHeader(404);
		return;
	}

	object_ID := bson.ObjectIdHex(movie_id);
	resultados := Pelicula{};
	err := collection.FindId(object_ID).One(&resultados);

	if (err != nil) {
		response.WriteHeader(404);
		return;
	}

	respuestaPelicula(200, response, resultados);
}

func AgregarPelicula (response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body);

	var dataPelicula Pelicula;
	err := decoder.Decode(&dataPelicula);

	if (err != nil) {
		panic(err);
	}

	defer request.Body.Close();

	log.Println(dataPelicula);

	err = collection.Insert(dataPelicula);
	
	if (err != nil) {
		response.WriteHeader(500);
		return;
	}
	
	respuestaPelicula(200, response, dataPelicula);
}

func ActualizarPelicula (response http.ResponseWriter, request *http.Request) {
	var dataPelicula Pelicula;
	params := mux.Vars(request);
	movie_id := params["id"];

	if (!bson.IsObjectIdHex(movie_id)) { //Si no es hexadecimal
		response.WriteHeader(404);
		return;
	}

	object_ID := bson.ObjectIdHex(movie_id);

	decoder := json.NewDecoder(request.Body);
	err := decoder.Decode(&dataPelicula);

	if (err != nil) {
		response.WriteHeader(404);
		return;
	}

	defer request.Body.Close();

	log.Println(dataPelicula);

	document := bson.M{"_id": object_ID};
	change := bson.M{"$set": dataPelicula};

	err = collection.Update(document, change);

	if (err != nil) {
		response.WriteHeader(404);
		return;
	}


}

func EliminarPelicula (response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request);
	movie_id := params["id"];

	if (!bson.IsObjectIdHex(movie_id)) { //Si no es hexadecimal
		response.WriteHeader(404);
		return;
	}

	object_ID := bson.ObjectIdHex(movie_id);
	err := collection.RemoveId(object_ID);

	if (err != nil) {
		response.WriteHeader(404);
		return;
	}

	//res := Mensaje {"Success", "La película ha sido eliminada"};
	res := new (Mensaje);
	res.setStatus("Success");
	res.setMensaje("La película ha sido eliminada");

	response.Header().Set("Content-Type", "application/json");
	response.WriteHeader(200);

	json.NewEncoder(response).Encode(res);
}

func respuestaPelicula(codStatus int, writer http.ResponseWriter, resultados Pelicula) {
	writer.Header().Set("Content-Type", "application/json");
	writer.WriteHeader(codStatus);

	json.NewEncoder(writer).Encode(resultados);
}

func respuestaPeliculas(codStatus int, writer http.ResponseWriter, resultados []Pelicula) {
	writer.Header().Set("Content-Type", "application/json");
	writer.WriteHeader(codStatus);

	json.NewEncoder(writer).Encode(resultados);
}

// func mostrarError(err) {
// 	if (err != nil) {
// 		panic(err);
// 	}
// }

