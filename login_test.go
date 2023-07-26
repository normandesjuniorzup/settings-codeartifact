package main_test

import (
	"fmt"
	"io"
	. "settings"
	"strings"
	"testing"
)

type FakeLoginCommand struct {
	Stdout io.Writer
}

func (f *FakeLoginCommand) Run() error {
	if _, err := fmt.Fprint(f.Stdout, "Some other text that may happen here\n"); err != nil {
		return err
	}
	if _, err := fmt.Fprint(f.Stdout, SUCCESS_LOGIN_MESSAGE); err != nil {
		return err
	}
	return nil
}

func TestLogin(t *testing.T) {
	cmd := &FakeLoginCommand{}
	var out strings.Builder
	cmd.Stdout = &out

	err := Login(cmd, &out)

	if err != nil {
		t.Fatalf("should not have error login, but got err: %s", err)
	}
}
