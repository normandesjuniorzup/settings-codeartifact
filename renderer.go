package main

import (
	"embed"
	"io"
	"text/template"
)

//go:embed "templates/*"
var settingsTemplate embed.FS

type SettingsRenderer struct {
	templ *template.Template
}

func NewSettingsRenderer() (*SettingsRenderer, error) {
	templ, err := template.ParseFS(settingsTemplate, "templates/*.goxml")
	if err != nil {
		return nil, err
	}

	return &SettingsRenderer{templ: templ}, nil
}

func (r *SettingsRenderer) Render(w io.Writer, token string) error {
	if err := r.templ.Execute(w, token); err != nil {
		return err
	}

	return nil
}
