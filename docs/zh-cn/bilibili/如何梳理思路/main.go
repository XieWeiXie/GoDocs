package main

import (
	"encoding/json"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"strconv"
)
func main() {
	u := UserResource{}
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		Doc("get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").
			DataType("string")).
		Writes(UserResource{}))
	container := restful.DefaultContainer.Add(ws)
	go func() {
		log.Fatal(http.ListenAndServe(":6060", container))
	}()




	container2 := restful.NewContainer()
	ws2 := new(restful.WebService)
	ws2.
		Path("/sss").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)
	ws2.Route(ws2.GET("/{user-id}").To(u.findUser).
		Doc("get a user").
		Param(ws2.PathParameter("user-id", "identifier of the user").
			DataType("string")).
		Writes(UserResource{}))
	container2.Add(ws2)

	server := http.Server{
		Addr:":6061",
		Handler: container2,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	select {

	}

}
type UserResource struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func NewUser()[]UserResource {
	return []UserResource{
		{
			ID: 1,
			Name: "GO",
			Age: 10,
		},
		{
			ID:2,
			Name: "Python",
			Age: 28,
		},
	}
}

func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	intID, _ := strconv.Atoi(id)
	for _, i := range NewUser() {
		if intID == i.ID{
			json.NewEncoder(response).Encode(i)
		}
	}

}