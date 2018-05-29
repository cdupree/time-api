package main

import (
   "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
)

type Metadata struct {
    Version string `json:"version"`
    Query string   `json:"query"`
}                          

type HMS struct {
    Hour int                `json:"hour"`
    Minute int              `json:"minute"`
    Second int              `json:"second"`
}

type Date struct {
    Month string           `json:"month"`
    Day int                `json:"day"`
    Year int               `json:"year"`
}

type DatetimeReply struct {
    Metadata                   `json:"metadata"`
    Datetime string            `json:"datetime"`
    HMS                        `json:"time"`
    Date                       `json:"date"`
}

type HealthCheck struct {
    Msg string              `json:"msg"`
}

func GetDatetime(w http.ResponseWriter, r *http.Request) {
    //params := mux.Vars(r)

    t := time.Now().Round(0)

    reply_json := DatetimeReply{

        Metadata: Metadata{
            Version: "v2",
            Query: r.URL.String(),
        },

        Datetime: t.String(),

        Date: Date {
            Month: t.Month().String(),
            Day: t.Day(),
            Year: t.Year(),
        },

        HMS: HMS {
            Hour: t.Hour(),
            Minute: t.Minute(),
            Second: t.Second(),
        },
    }

    err := json.NewEncoder(w).Encode(reply_json)

    if err != nil {
        panic(err)
    }
}

func Health(w http.ResponseWriter, r *http.Request) {
    reply_json := HealthCheck {
        Msg: "Don't worry, be happy",
    }

    err := json.NewEncoder(w).Encode(reply_json)

    if err != nil {
        panic(err)
    }
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/datetime", GetDatetime).Methods("GET")
    router.HandleFunc("/healthz", Health).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}
