package main

import (
    "net/http"
    //"fmt"
    "encoding/json"
    "os"
    "time"
)

type HandOn struct {
    Time time.Time `json:"time"`
    Hostname string `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.NotFound(w,r)
        return
    }

    resp := HandOn {
        Time: time.Now(),
        Hostname: os.Getenv("HOSTNAME"),
    }

    jsonResp, err := json.Marshal(&resp)

    if(err != nil){
        w.Write([]byte("Error"))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    //resp := fmt.Sprintf("La hora es %v y hostname es %v", time.Now(), os.Getenv("HOSTNAME"))
    //w.Write([]byte(resp))
    w.Write(jsonResp)
}

func main() {
    http.HandleFunc("/", ServeHTTP)
    http.ListenAndServe(":9090", nil)
}