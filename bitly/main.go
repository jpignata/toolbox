package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jpignata/toolbox/pkg/ssm"
)

const (
	key = "bitly_token"
	url = "https://api-ssl.bitly.com/v4/bitlinks"
)

// Bitlink represents a shortened URL. (https://dev.bitly.com/v4_documentation.html)
type Bitlink struct {
	URL string `json:"long_url"`
}

type response struct {
	Link string `json:"link"`
}

func main() {
	token, err := ssm.GetSecureString(key)

	if err != nil {
		fmt.Printf("Couldn't read SecretString (%s): %s\n", key, err)
		os.Exit(1)
	}

	location, err := create(token, Bitlink{os.Args[1]})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(location)
}

func create(token string, bitlink Bitlink) (string, error) {
	var r response

	body, err := json.Marshal(bitlink)

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errBody, _ := ioutil.ReadAll(resp.Body)

		return "", errors.New(string(errBody))
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(b, &r)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return r.Link, nil
}
