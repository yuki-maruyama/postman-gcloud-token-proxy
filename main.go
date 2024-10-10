package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

const port = "8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("gcloud", "auth", "print-identity-token")
		output, err := cmd.Output()
		if err != nil {
			http.Error(w, "Failed to get identity token", http.StatusInternalServerError)
			return
		}
		res := string(output)
		fmt.Fprintf(w, "%s", res[:len(res)-1])
	})

	fmt.Println("Server listening on port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
