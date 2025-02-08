package user

import (
	"awesomeProject/helper"
	"awesomeProject/responsepb"
	"awesomeProject/sse"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	service    *Service
	sseService *sse.Service
}

func NewHandler(service *Service, sseService *sse.Service) *Handler {
	return &Handler{service: service, sseService: sseService}
}

func (h *Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	users := h.service.GetUsers(r.Context())

	usersPb := ModelToProtoList(users)

	// Bikin response
	response := &responsepb.Response{
		Metadata: &responsepb.Metadata{
			UserCount: int32(len(usersPb)),
		},
		Data: usersPb,
	}

	// Serialize ke Protobuf binary
	protoData, err := proto.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	debug := false
	if debug {
		// Serialize ke JSON (dari Protobuf)
		jsonData, err := protojson.Marshal(response)
		if err != nil {
			fmt.Println("Error serializing to JSON:", err)
			return
		}
		w.Write(jsonData)
	} else {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(protoData)
	}
}

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	username := "johndoe"
	password := "secret"
	_, err := h.service.CreateUser(r.Context(), username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	h.sseService.Broadcast("modify", "AddUser")
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.service.DeleteUserById(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	h.sseService.Broadcast("modify", "DeleteUser")
}

func (h *Handler) TestAddUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	fmt.Println("TestAddUser")

	body, err := io.ReadAll(r.Body)
	helper.PanicIfError(err)

	user := &responsepb.User{}
	err = proto.Unmarshal(body, user)
	helper.PanicIfError(err)

	fmt.Println("TestAddUser", user)
}
