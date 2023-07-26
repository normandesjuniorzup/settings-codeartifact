package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
)

func main() {
	ssoLogin := exec.Command("aws", "sso", "login", "--profile", "CodeArtifact")
	var out strings.Builder
	ssoLogin.Stdout = &out
	if err := Login(ssoLogin, &out); err != nil {
		log.Fatalf("error login into AWS: %s", err)
	}

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile("CodeArtifact"),
	)
	if err != nil {
		log.Fatalf("error loading config: %s", err)
	}
	client := codeartifact.NewFromConfig(cfg)
	tokenOut, err := GetAuthorizationToken(client)
	if err != nil {
		log.Fatalf("error getting authorization token: %s", err)
	}
	fmt.Printf("The token will expire in BRT Time: %s\n", tokenOut.Expiration)

	settingsRenderer, err := NewSettingsRenderer()
	if err != nil {
		log.Fatalf("error preparing settings.xml: %s", err)
	}

	settingsFile, err := os.Create("settings.xml")
	if err != nil {
		log.Fatalf("error creating settings.xml file: %s", err)
	}
	defer settingsFile.Close()

	if err := settingsRenderer.Render(settingsFile, tokenOut.Token); err != nil {
		log.Fatalf("error rendering settings.xml: %s", err)
	}

	fmt.Println("\n\ndone")
}
