package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	user := "admin"
	pass := "secret123"
	token := auth(user, pass)
	mainSubNetId := getSubnetId(token, "10.0.0.0", "8")
	fmt.Println(requestNewSubnet(token, mainSubNetId, "21"))
	fmt.Println(createNewSubnet(token, mainSubNetId, "21"))
}

func auth(user, pass string) string {
	type IpamApiResp struct {
		Code    int  `json:"code"`
		Success bool `json:"success"`
		Data    struct {
			Token   string `json:"token"`
			Expires string `json:"expires"`
		} `json:"data"`
		Time float64 `json:"time"`
	}

	var authResp IpamApiResp

	c := http.Client{
		Transport: &http.Transport{},
	}
	req, err := http.NewRequest("POST", "http://127.0.0.1/api/apiclient/user", nil)

	if err != nil {
		log.Fatalf("error creating new request for authentication %v\n", err)
		return ""
	}

	req.SetBasicAuth(user, pass)
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("error doing a request for authentication %v\n", err)
		return ""
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		log.Fatalf("error decoding resp json %v\n", err)
	}

	return fmt.Sprintf(authResp.Data.Token)
}

func getSubnetId(token, subnet, mask string) int {
	type IpamSubnets struct {
		Code    int  `json:"code"`
		Success bool `json:"success"`
		Data    []struct {
			ID          int    `json:"id"`
			Subnet      string `json:"subnet"`
			Mask        string `json:"mask"`
			SectionID   int    `json:"sectionId"`
			Description string `json:"description"`
			IsFull      int    `json:"isFull"`
			IsPool      int    `json:"isPool"`
			Tag         int    `json:"tag"`
		} `json:"data"`
		Time float64 `json:"time"`
	}

	var subnetsResp IpamSubnets
	c := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost/api/apiclient/subnets", nil)

	if err != nil {
		log.Fatalf("error creating new request for subnets list %v\n", err)
		return 0
	}
	req.Header.Add("token", token)

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("error doing a request for subnets list %v\n", err)
		return 0
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&subnetsResp)
	if err != nil {
		log.Fatalf("error decoding subnets list resp json %v\n", err)
	}

	for _, v := range subnetsResp.Data {
		if v.Subnet == subnet && v.Mask == mask {
			return v.ID
		}
	}

	return 0
}

func requestNewSubnet(token string, mainSubNetId int, mask string) string {
	type reqNewSubnet struct {
		Code    int    `json:"code"`
		Success bool   `json:"success"`
		Data    string `json:"data"`
	}
	var newSubnet reqNewSubnet
	c := &http.Client{}
	reqUrl := fmt.Sprintf("http://localhost/api/apiclient/subnets/%d/first_subnet/%s", mainSubNetId, mask)
	req, err := http.NewRequest("GET", reqUrl, nil)

	if err != nil {
		log.Fatalf("error creating new request for subnets list %v\n", err)
		return ""
	}
	req.Header.Add("token", token)

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("error doing a request for subnets list %v\n", err)
		return ""
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&newSubnet)
	if err != nil {
		log.Fatalf("error decoding subnets list resp json %v\n", err)
	}
	return fmt.Sprint(newSubnet)
}

func createNewSubnet(token string, mainSubNetId int, mask string) string {
	type createNewSubnet struct {
		Code    int     `json:"code"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
		ID      string  `json:"id"`
		Data    string  `json:"data"`
		Time    float64 `json:"time"`
	}

	var createSubnet createNewSubnet
	c := &http.Client{}
	reqUrl := fmt.Sprintf("http://localhost/api/apiclient/subnets/%d/first_subnet/%s", mainSubNetId, mask)
	req, err := http.NewRequest("POST", reqUrl, nil)

	if err != nil {
		log.Fatalf("error creating new creation request for subnets list %v\n", err)
		return ""
	}
	req.Header.Add("token", token)

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("error doing new creation request for subnets list %v\n", err)
		return ""
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&createSubnet)
	if err != nil {
		log.Fatalf("error decoding subnets list resp json %v\n", err)
	}

	return fmt.Sprint(createSubnet)
}
