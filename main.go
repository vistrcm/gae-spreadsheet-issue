// Sample app showing issue with GAE -> google spreadsheets
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func main() {
	http.HandleFunc("/", indexHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s\n", port)
	}

	// let's check app engine instance scopes
	scopes, _ := metadata.Get("instance/service-accounts/default/scopes")
	log.Printf("[DEBUG] metadata scopes: %s.\n", scopes)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	// [END setting_port]
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, _ := google.DefaultClient(ctx, "https://www.googleapis.com/auth/spreadsheets.readonly")
	srv, err := sheets.New(client)

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	readRange := "Class Data!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v\n", err)
	}

	if len(resp.Values) == 0 {
		fmt.Fprintf(w, "No data found.\n")
	} else {
		fmt.Fprintf(w, "Name, Major:\n")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Fprintf(w, "%s, %s\n", row[0], row[4])
		}
	}

}
