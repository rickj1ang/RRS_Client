package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL = "http://localhost:4000" // Replace with your API base URL

func Login(email, password string) (string, error) {
	loginData := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)
	resp, err := http.Post(baseURL+"/tokens/authentication", "application/json", strings.NewReader(loginData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("login failed: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	token, ok := result["authentication_token"].(string)
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}

	return token, nil
}

func Get(endpoint, token string) (any, error) {
	req, err := http.NewRequest("GET", baseURL+endpoint, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		return "", fmt.Errorf("Request error: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	records, ok := result["records"]
	if !ok {
		return "", fmt.Errorf("record not found in response")
	}

	return records, nil
}

func Post(endpoint, data, token string) (string, error) {
	req, err := http.NewRequest("POST", baseURL+endpoint, strings.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Put(endpoint, data, token string) (string, error) {
	req, err := http.NewRequest("PUT", baseURL+endpoint, strings.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Delete(endpoint, token string) (string, error) {
	req, err := http.NewRequest("DELETE", baseURL+endpoint, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
