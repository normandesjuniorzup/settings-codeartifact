package main

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
)

type CodeArtifactClient interface {
	GetAuthorizationToken(ctx context.Context, params *codeartifact.GetAuthorizationTokenInput, optFns ...func(*codeartifact.Options)) (*codeartifact.GetAuthorizationTokenOutput, error)
}

type TokenOutput struct {
	Token      string
	Expiration string
}

func GetAuthorizationToken(client CodeArtifactClient) (*TokenOutput, error) {
	ctx := context.TODO()

	input := &codeartifact.GetAuthorizationTokenInput{
		Domain:      aws.String("claro-codeartifact"),
		DomainOwner: aws.String("319569500149"),
	}

	tokenOut, err := client.GetAuthorizationToken(ctx, input)
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	sp := tokenOut.Expiration.In(loc)

	return &TokenOutput{
		Token:      *tokenOut.AuthorizationToken,
		Expiration: sp.Format("02 January 2006, 15:04"),
	}, nil
}
