package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	razorpay "github.com/razorpay/razorpay-go"
)

type PageVariables struct {
	OrderId string
}

func main() {
	http.HandleFunc("/", App)
	log.Fatal(http.ListenAndServe(":8089", nil))
}

func App(w http.ResponseWriter, r *http.Request) {

	//Create order_id from the server
	client := razorpay.NewClient("rzp_test_kEtg65WKqGTpKd", "gPURxG4gzTmeNJKqqz61YCHV")

	data := map[string]interface{}{
		"amount":   50000,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)

	if err != nil {
		fmt.Println("Problem getting the repository information", err)
		os.Exit(1)
	}

	value := body["id"]
	str := value.(string)

	HomePageVars := PageVariables{ //store the order_id in a struct
		OrderId: str,
	}

	t, err := template.ParseFiles("app.html")
	if err != nil {
		log.Print("template parsing error", err)
	}

	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template excecuting error", err)
	}
}
