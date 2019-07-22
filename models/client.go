package models

import "encoding/json"
import "gopkg.in/oauth2.v3"

// Client client model
type Client struct {
	ID     string `json:"clientId"`
	Secret string `json:"clientSecret"`
	Domain string `json:"domain"`
	UserID string `json:"userId"`
}

// GetID client id
func (c Client) GetID() string {
	return c.ID
}

// GetSecret client domain
func (c Client) GetSecret() string {
	return c.Secret
}

// GetDomain client domain
func (c Client) GetDomain() string {
	return c.Domain
}

// GetUserID user id
func (c Client) GetUserID() string {
	return c.UserID
}

func Marshall(cli oauth2.ClientInfo, val *string) error {
	raw, err := json.Marshal(cli)
	if err != nil {
		return err
	}
	*val = string(raw)
	return nil
}

func Unmarshall(cli *oauth2.ClientInfo, val string) error {
	if err := json.Unmarshal([]byte(val), cli); err != nil {
		return err
	}
	return nil
}
