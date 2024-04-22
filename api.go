package main

import(
  "encoding/json"
  "net/http"
"fmt"
"log"
  "github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
  w.WriteHeader(status)
  w.Header().Set("Content-Type", "application/json")

  return json.NewEncoder(w).Encode(v)
}

type APIServer struct {
  listenAddr string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
  Error string
}



func MakeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if err := f(w, r); err != nil {

      WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})

    }
  }}

func NewAPIServer(listenAddr string) *APIServer {
  return &APIServer{
    listenAddr: listenAddr,
  }
}
func (s *APIServer) Run() {
  router:= mux.NewRouter()
  log.Println("api server running on port: ", s.listenAddr)
  router.HandleFunc("/account", MakeHTTPHandleFunc(s.handleAccount))
  http.ListenAndServe(s.listenAddr, router)
}


func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
  if r.Method == "GET" {
    return s.handleGetAccount(w, r)
  }
  if r.Method == "POST" {
    return s.handleCreateAccount(w, r)
  }
  if r.Method == "DELETE" {
    return s.handleDeleteAccount(w, r)
  }
  return fmt.Errorf("method not allowed: %s", r.Method)
}
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
  account := NewAccount("Anthony", "GG")
  return WriteJSON(w, http.StatusOK, account)
}
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
  return nil
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
  return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
  return nil
}

