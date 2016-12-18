package controller

import (
    "app/model"
    "fmt"
    "golang.org/x/net/context"
    "net/http"
    "encoding/json"
)

func PersonList(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
   switch req.Method {
    case "GET":
        people, _ := model.PersonList(ctx);
        js, err := json.Marshal(people)
        if err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprint(rw, string(js))
   }
}

func Person(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
   switch req.Method {
    case "POST":
        res, status := model.PersonAdd(ctx);
        if (status != nil) {
            fmt.Fprint(rw, "false")
        }
        
        js, err := json.Marshal(res)
        if err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprint(rw, string(js))
    case "DELETE":
        res, status := model.PersonDelete(ctx);
        if (status != nil) {
            fmt.Fprint(rw, "false")
        }
        fmt.Fprint(rw, string(res))
    case "GET":
        res, status := model.PersonGet(ctx);
        if (status != nil) {
            fmt.Fprint(rw, "false")
            return
        }
        
        js, err := json.Marshal(res)
        if err != nil {
            http.Error(rw, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprint(rw, string(js))
    case "PUT":
        _, status := model.PersonEdit(ctx);
        if (status != nil) {
            fmt.Fprint(rw, "false")
            return
        }

        fmt.Fprint(rw, "true")
        
        
   }
   
   
}