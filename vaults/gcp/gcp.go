package gcp

import (
	"encoding/json"
	"fmt"
)

type VaultClientData struct {
	RequestID     string `json:"request_id"`
	LeaseID       string `json:"lease_id"`
	Renewable     bool   `json:"renewable"`
	LeaseDuration int    `json:"lease_duration"`
	Data          struct {
	} `json:"data"`
	WrapInfo interface{} `json:"wrap_info"`
	Warnings []string    `json:"warnings"`
	Auth     struct {
		ClientToken   string   `json:"client_token"`
		Accessor      string   `json:"accessor"`
		Policies      []string `json:"policies"`
		TokenPolicies []string `json:"token_policies"`
		Metadata      struct {
			Username string `json:"username"`
		} `json:"metadata"`
		LeaseDuration int    `json:"lease_duration"`
		Renewable     bool   `json:"renewable"`
		EntityID      string `json:"entity_id"`
		TokenType     string `json:"token_type"`
		Orphan        bool   `json:"orphan"`
	} `json:"auth"`
}

func LoginServiceUser(serviceUser string, password string, vaultID string) string {

	base_url := urls.Base("service")

	url := base_url.String() + "/" + vaultID

	body := make(map[string]string)
	body["password"] = password

	req, err := request.New("POST", url, body)
	if err != nil {
		panic(err)
	}

	request.DisableSSLTLSValidation()

	data := request.Do(req)

	var vaultClientData VaultClientData
	err = json.Unmarshal([]byte(data), &vaultClientData)
	if err != nil {
		panic(err)
	}

	return vaultClientData.GetServiceUserToken(serviceUser)

}

type ServiceUserData struct {
	RequestID     string `json:"request_id"`
	LeaseID       string `json:"lease_id"`
	Renewable     bool   `json:"renewable"`
	LeaseDuration int    `json:"lease_duration"`
	Data          struct {
		CurrentPassword string `json:"current_password"`
		Username        string `json:"username"`
	} `json:"data"`
	WrapInfo interface{} `json:"wrap_info"`
	Warnings interface{} `json:"warnings"`
	Auth     interface{} `json:"auth"`
}

func (vaultClientData VaultClientData) GetServiceUserToken(serviceUser string) string {
	base_url := urls.Base("service")

	url := base_url.String() + "/" + serviceUser

	req, err := request.New("GET", url, nil)
	if err != nil {
		panic(err)
	}

	request.AddHeaders(req, "X-Vault-Token", vaultClientData.Auth.ClientToken)

	data := request.Do(req)

	var serviceUserData ServiceUserData
	err = json.Unmarshal([]byte(data), &serviceUserData)
	if err != nil {
		panic(err)
	}

	return serviceUserData.Data.CurrentPassword
}

func LoginVault(username string, password string, environment string) string {

	base_url := urls.Base(environment)

	url := base_url.String() + "/" + username

	body := make(map[string]string)
	body["password"] = password

	req, err := request.New("POST", url, body)
	if err != nil {
		panic(err)
	}

	data := request.Do(req)

	var vaultClientData VaultClientData
	err = json.Unmarshal([]byte(data), &vaultClientData)
	if err != nil {
		fmt.Println(err.Error())
	}

	return vaultClientData.Auth.ClientToken

}

func UpdateVault() {

}

func GetRoleSecretID(username string) {
}

func Get(vaultID string, token string) {
}

func CreateSecret(topdomain string, domain string, vaultname string, environment string) {
}

func Delete() {

}

func Update() {

}
