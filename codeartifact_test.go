package main_test

import (
	"context"
	. "settings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
)

type SpyCodeArtifactClient struct {
	Domain      string
	DomainOwner string
}

func (c *SpyCodeArtifactClient) GetAuthorizationToken(ctx context.Context, params *codeartifact.GetAuthorizationTokenInput, optFns ...func(*codeartifact.Options)) (*codeartifact.GetAuthorizationTokenOutput, error) {
	c.Domain = *params.Domain
	c.DomainOwner = *params.DomainOwner

	expirationDate, _ := time.Parse("02 January 2006, 15:04", "29 July 2023, 23:00")
	return &codeartifact.GetAuthorizationTokenOutput{
		AuthorizationToken: aws.String("some-valid-token"),
		Expiration:         &expirationDate,
	}, nil
}

func TestGetAuthorizationToken(t *testing.T) {
	client := &SpyCodeArtifactClient{}

	tokenOut, err := GetAuthorizationToken(client)
	if err != nil {
		t.Fatalf("should not have error getting token: %s", err)
	}

	wantToken := "some-valid-token"
	if wantToken != tokenOut.Token {
		t.Errorf("got token %q, want %q", tokenOut.Token, wantToken)
	}

	wantExpiration := "29 July 2023, 20:00"
	if wantExpiration != tokenOut.Expiration {
		t.Errorf("got expiration %q, want %q", tokenOut.Expiration, wantExpiration)
	}

	if client.Domain != "zup" {
		t.Errorf("got %q as domain, should be %q", client.Domain, "zup")
	}
	if client.DomainOwner != "546045978864" {
		t.Errorf("got %q as domain owner, should be %q", client.DomainOwner, "546045978864")
	}
}
