package main

import (
    "app/model"
    "app/controller"
    "log"
    "golang.org/x/net/context"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type ContextHandler interface {
    ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
    ctx = context.WithValue(ctx, "req", req)
    h(ctx, rw, req)
}

type ContextAdapter struct {
    ctx     context.Context
    handler ContextHandler
}

func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    ca.handler.ServeHTTPContext(ca.ctx, rw, req)
}

type Configuration struct {
    Database databaseConfig
    Server serverConfig
}

type databaseConfig struct {
    Driver string `json:"driver"`
    Host string `json:"host"`
    Port int `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
    Database string `json:"database"`
}

type serverConfig struct {
    Address string `json:"address"`
}

func LoadConfig(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return config
}

func main() {
    config := LoadConfig("./config/config.json")

    db, err := model.NewDB(config.Database.Driver, config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Database)
    if err != nil {
        log.Panic(err)
    }
    
    ctx := context.Background()
    ctx = context.WithValue(ctx, "db", db)
    
    model.PersonCreate(ctx);

    http.Handle("/personList", &ContextAdapter{ctx, ContextHandlerFunc(controller.PersonList)})
    http.Handle("/person", &ContextAdapter{ctx, ContextHandlerFunc(controller.Person)})
    
    http.ListenAndServe(config.Server.Address, nil)
}