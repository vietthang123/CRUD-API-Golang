package main

import (
	"net/http"

	"github.com/spf13/cobra"

	"developer-orientenergy-golang/internal/app/router"
)

var rootCmd = &cobra.Command{
	Use:   "schoolAPI",
	Short: "Start app",
	Run: func(cmd *cobra.Command, args []string) {
		defineRouter()
	},
}

func defineRouter() {
	r := router.NewRouter()
	println("connect port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		print(err)
	}
}
