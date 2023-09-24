package transport

import (
	"net/http"
	"workflow-service/transport/util"
)

type Subscription struct {
	PubsubName string            `json:"pubsubname"`
	Topic      string            `json:"topic"`
	Metadata   map[string]string `json:"metadata,omitempty"`
	Route      Route             `json:"routes"`
}

type Route struct {
	Rules   []Rule `json:"rules,omitempty"`
	Default string `json:"default,omitempty"`
}

type Rule struct {
	Match string `json:"match"`
	Path  string `json:"path"`
}

func ConfigureSubscribeHandler(w http.ResponseWriter, _ *http.Request) {
	pubsubName := "mrf-pub-sub"
	v := []Subscription{
		{
			PubsubName: pubsubName,
			Topic:      "user-delete",
			Route: Route{
				Rules: []Rule{
					{
						Match: `event.type == "user-delete"`,
						Path:  "/users/delete",
					},
				},
				Default: "/users/delete",
			},
		},
	}

	util.WriteResponse(w, http.StatusOK, v)
}
