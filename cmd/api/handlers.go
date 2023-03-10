package main

import (
	"net/http"
)

// jsonResponse is the type used for generic JSON responses
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Login is the handler used to attempt to log a user into the api
func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		UserName string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	// err := json.NewDecoder(r.Body).Decode(&creds)
	// if err != nil {
	// 	// send back error message
	// 	app.errorLog.Println("invalid json")
	// 	payload.Error = true
	// 	payload.Message = "invalid json"

	// 	out, err := json.MarshalIndent(payload, "", "\t")
	// 	if err != nil {
	// 		app.errorLog.Println(err)
	// 	}

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(out)
	// 	return
	// }

	// TODO authenticate
	app.infoLog.Println(creds.UserName, creds.Password)

	// send back a response
	payload.Error = false
	payload.Message = "Signed in"

	err = app.writeJSON(w, http.StatusOK, payload)

	// out, err := json.MarshalIndent(payload, "", "\t")

	if err != nil {
		app.errorLog.Println(err)
	}

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(out)
}
