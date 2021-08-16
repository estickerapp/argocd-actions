package main

import (
	"log"
	"os"

	"github.com/omegion/argocd-actions/internal/argocd"
	ctrl "github.com/omegion/argocd-actions/internal/controller"
)

func main() {
	PlainText := os.Getenv("INPUT_PLAIN")
	IsPlainText := false

	if PlainText == "true" {
		IsPlainText = true
	}

	options := argocd.APIOptions{
		Address: os.Getenv("INPUT_ADDRESS"),
		Token:   os.Getenv("INPUT_TOKEN"),
		PlainText: IsPlainText,
	}

	api := argocd.NewAPI(options)
	controller := ctrl.NewController(api)

	err := controller.Sync(os.Getenv("INPUT_APPNAME"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
