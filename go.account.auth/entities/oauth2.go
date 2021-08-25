package entities

type Authorization {
	ClientID string `json: "client_id"`
	RedirectURI string `json: "redirect_uri"`
	ResponseType string `json: "response_type"`
	scope string `json: "scope"`
	state string `json: "state"`
}