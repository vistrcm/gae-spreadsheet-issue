// Sample app showing issue with GAE -> google spreadsheets
package main

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// let's check app engine instance scopes
	scopes, _ := metadata.Get("instance/service-accounts/default/scopes")
	log.Infof(ctx, "[DEBUG] metadata scopes: %s.\n", scopes)

	client, _ := google.DefaultClient(ctx, "https://www.googleapis.com/auth/spreadsheets.readonly")
	srv, err := sheets.New(client)

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	readRange := "Class Data!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		msg := fmt.Sprintf("Unable to retrieve data from sheet: %v\n", err)
		log.Errorf(ctx, msg)
		panic(msg)
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
