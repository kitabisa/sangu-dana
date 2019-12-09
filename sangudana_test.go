package dana

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/suite"
)

type DanaSanguTestSuite struct {
	suite.Suite
	client Client
}

type credentials struct {
	BaseUrl      string
	Version      string
	ClientId     string
	ClientSecret string
}

func TestDanaSanguTestSuite(t *testing.T) {
	suite.Run(t, new(DanaSanguTestSuite))
}

func (d *DanaSanguTestSuite) SetupSuite() {
	theToml, err := ioutil.ReadFile("credential_test.toml")
	if err != nil {
		d.T().Log(err)
		d.T().FailNow()
	}

	var cred credentials
	if _, err := toml.Decode(string(theToml), &cred); err != nil {
		d.T().Log(err)
		d.T().FailNow()
	}

	d.client = NewClient()
	d.client.BaseUrl = cred.BaseUrl
	d.client.Version = cred.Version
	d.client.ClientId = cred.ClientId
	d.client.ClientSecret = cred.ClientSecret
	d.client.LogLevel = 3
	d.client.SignatureEnabled = true

	privateKey, err := ioutil.ReadFile("my_private.pem")
	if err != nil {
		d.T().Log(err)
		d.T().FailNow()
	}
	d.client.PrivateKey = privateKey

	publicKey, err := ioutil.ReadFile("dana_public.pem")
	if err != nil {
		d.T().Log(err)
		d.T().FailNow()
	}
	d.client.PublicKey = publicKey
}

func (d *DanaSanguTestSuite) TestApplyTokenSuccess() {
	coreGateway := CoreGateway{
		Client: d.client,
	}

	reqBody := &RequestApplyAccessToken{
		GrantType:    "AUTHORIZATION_CODE",
		AuthCode:     "2P8pZedLvCW70sBKc7kWP5SZxlULgWDoeGCO7300",
		RefreshToken: "",
	}

	resp, err := coreGateway.ApplyAccessToken(reqBody)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("response: %v\n", resp.Response.Body)
	}
}
