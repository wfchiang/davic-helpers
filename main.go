package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
	"html/template"
	"github.com/gorilla/mux"
	"io/ioutil"
	"github.com/wfchiang/davic"
)

const (
	KEY_STORE_HTTP_REQUEST  = "http-reqt"
	KEY_STORE_HTTP_RESPONSE = "http-resp" 
	TYPE_BOOLEAN = "(boolean)"
	TYPE_NUMBER = "(number)"
	TYPE_STRING = "(string)"
	TYPE_OBJECT = "(object)"
	TYPE_EXPR = "(expression)"
	VALUE_INDEX_NUMBER = "Index " + TYPE_NUMBER
	VALUE_KEY_STRING = "Key " + TYPE_STRING 
	VALUE_KEY_EXPR = "Key " + TYPE_EXPR
	VALUE_VALUE_EXPR = "Value " + TYPE_EXPR
	VALUE_LHS_EXPR = "LHS Value " + TYPE_EXPR
	VALUE_RHS_EXPR = "RHS Value " + TYPE_EXPR
	VALUE_ARRAY_EXPR = "Array " + TYPE_EXPR
	VALUE_OBJ_EXPR = "Object " + TYPE_EXPR
)

type OptData struct {
	Name string 
	Type string 
	Symbol string 
	OpdNames []string
}

type OptList struct {
	SymbolOptMark string
	KeyHttpMethod string 
	KeyHttpUrl string 
	KeyHttpHeaders string 
	KeyHttpBody string 
	Operations []OptData 
}

// ====
// Globals 
// ====
var OptListData = OptList {
	SymbolOptMark: davic.SYMBOL_OPT_MARK, 
	KeyHttpMethod: davic.KEY_HTTP_METHOD, 
	KeyHttpUrl: davic.KEY_HTTP_URL, 
	KeyHttpHeaders: davic.KEY_HTTP_HEADERS, 
	KeyHttpBody: davic.KEY_HTTP_BODY, 
	Operations: []OptData {
		OptData {
			Name: "Http Call", 
			Symbol: davic.OPT_HTTP_CALL, 
			OpdNames: []string {"Http Request " + TYPE_OBJECT}}, 
		OptData {
			Name: "Lambda", 
			Symbol: davic.OPT_LAMBDA, 
			OpdNames: []string {"Function " + TYPE_EXPR}},
		OptData {
			Name: "Function Call",  
			Symbol: davic.OPT_FUNC_CALL, 
			OpdNames: []string {"Lambda " + TYPE_EXPR, VALUE_VALUE_EXPR}},  
		OptData {
			Name: "Stack Read", 
			Symbol: davic.OPT_STACK_READ, 
			OpdNames: []string {}}, 
		OptData {
			Name: "Store Read", 
			Symbol: davic.OPT_STORE_READ, 
			OpdNames: []string {VALUE_KEY_STRING}}, 
		OptData {
			Name: "Store Write", 
			Symbol: davic.OPT_STORE_WRITE, 
			OpdNames: []string {VALUE_KEY_STRING, VALUE_VALUE_EXPR}},
		OptData {
			Name: "Relation Equal", 
			Symbol: davic.OPT_RELATION_EQ, 
			OpdNames: []string {VALUE_LHS_EXPR, VALUE_RHS_EXPR}}, 
		OptData {
			Name: "Arithmetic Add", 
			Symbol: davic.OPT_ARITHMETIC_ADD, 
			OpdNames: []string {VALUE_LHS_EXPR, VALUE_RHS_EXPR}}, 
		OptData {
			Name: "Arithmetic Subtract", 
			Symbol: davic.OPT_ARITHMETIC_SUB, 
			OpdNames: []string {VALUE_LHS_EXPR, VALUE_RHS_EXPR}}, 
		OptData {
			Name: "Arithmetic Multiply", 
			Symbol: davic.OPT_ARITHMETIC_MUL, 
			OpdNames: []string {VALUE_LHS_EXPR, VALUE_RHS_EXPR}}, 
		OptData {
			Name: "Arithmetic Division", 
			Symbol: davic.OPT_ARITHMETIC_DIV, 
			OpdNames: []string {VALUE_LHS_EXPR, VALUE_RHS_EXPR}}, 
		OptData {
			Name: "String concat", 
			Symbol: davic.OPT_STRING_CONCAT, 
			OpdNames: []string {VALUE_LHS_EXPR, VALUE_RHS_EXPR}}, 
		OptData {
			Name: "Array Get/Read", 
			Symbol: davic.OPT_ARRAY_GET, 
			OpdNames: []string {VALUE_ARRAY_EXPR, VALUE_INDEX_NUMBER}}, 
		OptData {
			Name: "Object Read", 
			Symbol: davic.OPT_OBJ_READ,
			OpdNames: []string {VALUE_OBJ_EXPR, VALUE_KEY_EXPR}},  
		OptData {
			Name: "Object Update", 
			Symbol: davic.OPT_OBJ_UPDATE, 
			OpdNames: []string {VALUE_OBJ_EXPR, VALUE_KEY_EXPR, VALUE_VALUE_EXPR}} }}

// ==== 
// Recovery function 
// ====
func recoverFromPanic (http_resp http.ResponseWriter, id_service string) {
	if r := recover(); r != nil {
		err_message := fmt.Sprintf("%v", r)
		log.Println(fmt.Sprintf("[%s] %s", id_service, err_message))
		fmt.Fprintf(http_resp, err_message)
	}
}

// ====
// Handlers 
// ====
func homepageHandler (http_resp http.ResponseWriter, http_reqt *http.Request) {
	log.Println("Davic-helpers is Hit!")
	fmt.Fprintf(http_resp, "Davic Helpers are Here!")	
}

func optDataHandler (http_resp http.ResponseWriter, http_reqt *http.Request) {
	defer recoverFromPanic(http_resp, "opt-data")

	resp_body, err := json.Marshal(OptListData) 
	if err != nil {
		panic(fmt.Sprintf("Response marshalling failed: %v", err))
	} 
	
	fmt.Fprintf(http_resp, string(resp_body))
}

func davicHelperHandler (http_resp http.ResponseWriter, http_reqt *http.Request) {
	defer recoverFromPanic(http_resp, "davic-helper")

	log.Println("Davic Helper is Hit!")
	
	template_fname := "davic-helpers.html"
	tmpl, err := template.New(template_fname).Delims("<<", ">>").ParseFiles(template_fname)
	if (err != nil) {
		panic(fmt.Sprintf("Template load failed: %v", err))
	}

	tmpl.Execute(http_resp, nil)
	
	log.Println("Davic Helper responded")
}

func runDavicHandler (http_resp http.ResponseWriter, http_reqt *http.Request) {
	defer recoverFromPanic(http_resp, "run-davic")

	log.Println("Run-davic is Hit!")

	template_fname := "run-davic.html"
	tmpl, err := template.New(template_fname).Delims("<<", ">>").ParseFiles(template_fname)
	if (err != nil) {
		panic(fmt.Sprintf("Template load failed: %v", err))
	}

	tmpl.Execute(http_resp, nil)

	log.Println("Run-davic responded")
}

func davicHandler (http_resp http.ResponseWriter, http_reqt *http.Request) {
	defer recoverFromPanic(http_resp, "davic")

	log.Println("Davic is Hit!")

	// Read the request body 
	bytes_reqt_body, err := ioutil.ReadAll(http_reqt.Body)
	if err != nil {
		panic("Failed to read the request body")
	}

	reqt_body := string(bytes_reqt_body)
	log.Println("Davic/Go is Hit! Body: " + reqt_body)

	// Convert the string type request body to object
	reqt_obj := davic.CreateObjFromBytes(bytes_reqt_body)
	
	_, ok := reqt_obj["data"]
	if (!ok) {
		panic("Data field is missed in the request object")
	}
	opt_obj, ok := reqt_obj["opt"]
	if (!ok) {
		panic("Opt field is missed in the request object")
	}

	// Setup the Davic environment 
	env := davic.CreateNewEnvironment()
	env.Store = davic.CastInterfaceToObj(reqt_obj)

	// Execute the operation 
	rel_obj := davic.EvalExpr(env, opt_obj)

	// Marshal the response 
	resp_body, err := json.Marshal(rel_obj) 
	if err != nil {
		panic(fmt.Sprintf("Response marshalling failed: %v", err))
	} 
	
	fmt.Fprintf(http_resp, string(resp_body))
}

// ====
// Main
// ====
func main () {
	log.Println("Init File Server...")
	file_server := http.FileServer(http.Dir("./static/"))

	log.Println("Starting Davic-helpers...")
	mux_router := mux.NewRouter()

	mux_router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", file_server))
	mux_router.HandleFunc("/opt-data", optDataHandler).Methods("GET")
	mux_router.HandleFunc("/davic-helpers", davicHelperHandler).Methods("GET")
	// mux_router.HandleFunc("/run-davic", runDavicHandler).Methods("GET")
	// mux_router.HandleFunc("/davic", davicHandler).Methods("POST")
	mux_router.HandleFunc("/", homepageHandler)

	http.Handle("/", mux_router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Use defailt port %s", port)
	} 
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}