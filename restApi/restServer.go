package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

func defaultRoute(res http.ResponseWriter, req *http.Request) {
	queryStr := queryStr.URL.Query()
	queryName := queryStr.Get("name")
	if queryName == "" {
		queryName = "datCooCoo"
	}
}