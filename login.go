package main

import (
	"errors"
	"fmt"
	"strings"
)

const SUCCESS_LOGIN_MESSAGE = "Successfully logged into Start URL: https://zup.awsapps.com/start"

type Command interface {
	Run() error
}

func Login(cmd Command, out *strings.Builder) error {
	if err := cmd.Run(); err != nil {
		return err
	}

	outMessage := out.String()
	if !strings.Contains(outMessage, SUCCESS_LOGIN_MESSAGE) {
		errMsg := fmt.Sprintf("Success login should have the correct message %q, but found %q", SUCCESS_LOGIN_MESSAGE, outMessage)
		return errors.New(errMsg)
	}

	return nil
}
