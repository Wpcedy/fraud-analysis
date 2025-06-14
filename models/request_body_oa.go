package models

type RequestBodyOA struct {
    Model    string    `json:"model"`
    Messages []MessageOA `json:"messages"`
}
