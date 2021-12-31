package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"learn.oauth.client/model"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strings"
	"text/template"
	"time"
)

var config = struct {
	authURL             string
	logout              string
	afterLogoutRedirect string
	appID               string
	authCodeCallback    string
	tokenEndpoint       string
	appPassword         string
	servicesEndpoint    string
}{
	appPassword:         "W2gPUx5e3J7HCMgtLrbIf5HcmL9JnGDP",
	appID:               "billingApp",
	afterLogoutRedirect: "http://127.0.0.1:8080",
	authURL:             "http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/auth",
	logout:              "http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/logout",
	authCodeCallback:    "http://127.0.0.1:8080/authCodeRedirect",
	tokenEndpoint:       "http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/token",
	servicesEndpoint:    "http://127.0.0.1:8081/billing/v1/services",
}

var t = template.Must(template.ParseFiles("template/index.html"))
var tService = template.Must(template.ParseFiles("template/index.html", "template/services.html"))

type AppVar struct {
	AuthCode     string
	SessionState string
	AccessToken  string
	RefreshToken string
	Scope        string
	Services     []string
	State        map[string]struct{}
}

var appVar = newAppVar()

func newAppVar() AppVar {
	return AppVar{State: make(map[string]struct{})}
}
func init() {
	log.SetFlags(log.Ltime)
}
func enabledLog(handler func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		log.SetPrefix(handlerName + " ")
		log.Println("--> Home" + handlerName)
		log.Printf("request: %+v \n", r.RequestURI)
		log.Printf("response: %+v \n", w)
		handler(w, r)
		log.Println("<-- Home" + handlerName)
	}
}

func main() {
	fmt.Println("hello")
	http.HandleFunc("/", enabledLog(home))
	http.HandleFunc("/login", enabledLog(login))
	http.HandleFunc("/authCodeRedirect", enabledLog(authCodeRedirect))
	http.HandleFunc("/logout", enabledLog(logout))
	http.HandleFunc("/exchangeToken", enabledLog(exchangeToken))
	http.HandleFunc("/services", enabledLog(services))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	t.Execute(w, appVar)
}

func login(w http.ResponseWriter, r *http.Request) {
	// create a redirect url for authentication endpoint
	fmt.Println("login")

	req, err := http.NewRequest("GET", config.authURL, nil)
	if err != nil {
		log.Print(err)
		return
	}

	qs := url.Values{}
	state := uuid.New().String()
	appVar.State[state] = struct{}{}
	qs.Add("state", state)
	qs.Add("client_id", config.appID)
	qs.Add("response_type", "code")
	qs.Add("redirect_uri", config.authCodeCallback)
	req.URL.RawQuery = qs.Encode()

	//req.URL.RawQuery = "state=123abc&client_id=billingApp&response_type=code"
	http.Redirect(w, r, req.URL.String(), http.StatusFound)
}

func authCodeRedirect(w http.ResponseWriter, r *http.Request) {
	appVar.AuthCode = r.URL.Query().Get("code")
	callBackState := r.URL.Query().Get("state")
	if _, ok := appVar.State[callBackState]; !ok {
		fmt.Fprintf(w, "Error")
		return
	}
	delete(appVar.State, callBackState)
	appVar.SessionState = r.URL.Query().Get("session_state")
	r.URL.RawQuery = ""
	fmt.Printf("Request queries: %+v\n", appVar)
	//t.Execute(w, nil)
	http.Redirect(w, r, "http://127.0.0.1:8080", http.StatusFound)
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout")
	q := url.Values{}
	q.Add("redirect_uri", config.afterLogoutRedirect)
	logoutURL, err := url.Parse(config.logout)
	if err != nil {
		log.Println(err)
	}
	logoutURL.RawQuery = q.Encode()
	appVar = newAppVar()
	http.Redirect(w, r, logoutURL.String(), http.StatusFound)
}

func exchangeToken(w http.ResponseWriter, r *http.Request) {
	// Request
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", appVar.AuthCode)
	form.Add("redirect_uri", config.authCodeCallback)
	form.Add("client_id", config.appID)
	req, err := http.NewRequest("POST", config.tokenEndpoint, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Println(err)
		return
	}
	req.SetBasicAuth(config.appID, config.appPassword)

	// Client
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Println("couldn't get access token", err)
		return
	}

	// Process response
	byteBody, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}

	accessTokenResponse := &model.AccessTokenResponse{}
	json.Unmarshal(byteBody, accessTokenResponse)
	appVar.AccessToken = accessTokenResponse.AccessToken
	appVar.Scope = accessTokenResponse.Scope
	appVar.RefreshToken = accessTokenResponse.RefreshToken
	t.Execute(w, appVar)

}

func services(w http.ResponseWriter, r *http.Request) {
	// request
	req, err := http.NewRequest("GET", config.servicesEndpoint, nil)
	if err != nil {
		log.Println(err)
		tService.Execute(w, appVar)
		return
	}
	req.Header.Add("Authorization", "Bearer "+appVar.AccessToken)

	// client
	ctx, cancelFunc := context.WithTimeout(context.Background(), 50000*time.Microsecond)
	defer cancelFunc()
	c := http.Client{}
	res, err := c.Do(req.WithContext(ctx))
	if err != nil {
		log.Println(err)
		tService.Execute(w, appVar)
		return
	}
	byteBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		tService.Execute(w, appVar)
		return
	}

	log.Println("tgao4")
	// process response
	if res.StatusCode != 200 {
		log.Println("tgao102")
		log.Println(string(byteBody))
		tService.Execute(w, appVar)
		return
	}
	billingResponse := &model.Billing{}
	err = json.Unmarshal(byteBody, billingResponse)
	if err != nil {
		log.Println(err)
		tService.Execute(w, appVar)
		return
	}
	appVar.Services = billingResponse.Services

	log.Println("tgao3")
	log.Println(billingResponse)
	log.Println(appVar.Services)
	tService.Execute(w, appVar)
}
