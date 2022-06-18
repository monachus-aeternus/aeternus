package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	log "github.com/withmandala/go-log"
)

func main() {
	f, err := os.Create("aeternus.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	logger := log.New(f).WithColor().WithTimestamp()
	logger.Infof(`
    ___     ______  ______    ______    ____     _   __   __  __   _____
   /   |   / ____/ /_  __/   / ____/   / __ \   / | / /  / / / /  / ___/
  / /| |  / __/     / /     / __/     / /_/ /  /  |/ /  / / / /   \__ \ 
 / ___ | / /___    / /     / /___    / _, _/  / /|  /  / /_/ /   ___/ / 
/_/  |_|/_____/   /_/     /_____/   /_/ |_|  /_/ |_/   \____/   /____/  
                                                                        
`)
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logger.Info("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		logger.Info("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})
	logger.Info(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))
}
