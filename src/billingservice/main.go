package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learn.oauth.billing/model"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strings"
)

type BillingError struct {
	Error string `json:"error"`
}
type Billing struct {
	Services []string `json:"services"`
}
type TokenIntrospect struct {
	Jti      string `json:"jti"`
	Exp      int    `json:"exp"`
	Nbf      int    `json:"nbf"`
	Iat      int    `json:"iat"`
	Aud      string `json:"aud"`
	Typ      string `json:"typ"`
	AuthTime int    `json:"auth_time"`
	Acr      string `json:"acr"`
	Active   bool   `json:"active"`
}

var config = struct {
	tokenIntroSpection string
}{
	tokenIntroSpection: "http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/token/introspect",
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
func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	fmt.Println("hello")
	http.HandleFunc("/billing/v1/services", enabledLog(services))
	http.ListenAndServe(":8081", nil)
}

func services(w http.ResponseWriter, r *http.Request) {
	token, err := getToken(r)
	if err != nil {
		log.Println(err)
		makeErrorMessage(w, err.Error())
		return
	}
	log.Println("Token : ")
	log.Println(token)

	// Validate token
	if !validateToken(token) {
		makeErrorMessage(w, "InvalidToken")
		return
	}
	claimBytes, err := getClaim(token)
	if err != nil {
		log.Println(err)
		makeErrorMessage(w, "cannot parse token claim")
		return
	}
	tokenClaim := &model.Tokenclaim{}
	err = json.Unmarshal(claimBytes, tokenClaim)
	if err != nil {
		log.Println(err)
		makeErrorMessage(w, err.Error())
		return
	}
	scopes := strings.Split(tokenClaim.Scope, " ")
	for _, v := range scopes {
		log.Println("Scope: ", v)
	}
	if !strings.Contains(tokenClaim.Scope, "getBillingService") {
		makeErrorMessage(w, "invalid token. Required scope getBillingService")
		return
	}

	s := Billing{
		Services: []string{
			"electric",
			"phone",
			"internet",
			"water",
		},
	}
	encoder := json.NewEncoder(w)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-type", "application/json")
	encoder.Encode(s)
}

func getToken(r *http.Request) (string, error) {
	// header
	token := r.Header.Get("Authorization")
	if token != "" {
		auths := strings.Split(token, " ")
		if len(auths) != 2 {
			return "", fmt.Errorf("invalid Authorization header format")
		}
		return auths[1], nil
	}
	// form body
	token = r.FormValue("access_token")
	if token != "" {
		return token, nil
	}
	// query
	token = r.URL.Query().Get("access_token")
	if token != "" {
		return token, nil
	}
	return token, fmt.Errorf("access token is not presented")
}

func validateToken(token string) bool {
	// Request
	form := url.Values{}
	form.Add("token", token)
	form.Add("token_type_hint", "requesting_party_token")
	req, err := http.NewRequest("POST", config.tokenIntroSpection, strings.NewReader(form.Encode()))
	if err != nil {
		log.Println(err)
		return false
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("tokenChecker", "fmdXLM8pbKF2U7kGMGAVxbqMSXS4Bfax")
	// client
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	if res.StatusCode != 200 {
		log.Println("status is not 200 : ", res.StatusCode)
		return false
	}
	byteBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()

	introSpect := &TokenIntrospect{}
	err = json.Unmarshal(byteBody, introSpect)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("tgao1")

	log.Println(introSpect)
	return introSpect.Active
}

func makeErrorMessage(w http.ResponseWriter, errMsg string) {
	s := &BillingError{Error: errMsg}
	encoder := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusBadRequest)
	encoder.Encode(s)
}

func getClaim(token string) ([]byte, error) {
	tokenParts := strings.Split(token, ".")
	claim, err := base64.RawURLEncoding.DecodeString(tokenParts[1])
	if err != nil {
		return []byte{}, err
	}
	return claim, nil
}
