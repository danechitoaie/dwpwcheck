package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	var dwHostname string
	flag.StringVar(&dwHostname, "h", "", "Sandbox hostname")

	var dwUsername string
	flag.StringVar(&dwUsername, "u", "", "Sandbox username")

	var dwPassword string
	flag.StringVar(&dwPassword, "p", "", "Sandbox password")

	flag.Parse()

	if dwHostname == "" {
		fmt.Println("ERROR! Missing Sandbox hostname argument.")
		os.Exit(1)
	}

	if dwUsername == "" {
		fmt.Println("ERROR! Missing Sandbox username argument.")
		os.Exit(1)
	}

	if dwPassword == "" {
		fmt.Println("ERROR! Missing Sandbox password argument.")
		os.Exit(1)
	}

	reqUrl := fmt.Sprintf("https://%s/on/demandware.servlet/studiosvc/Sites", dwHostname)
	reqBody := strings.NewReader(`{"getSitesReq" : ""}`)

	req, err := http.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(dwUsername, dwPassword)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}

	if res.StatusCode == 401 {
		fmt.Println("ERROR: Invalid username or password!")
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
	res.Body.Close()

	var resJSON map[string]interface{}
	if err := json.Unmarshal(resBody, &resJSON); err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}

	resHasErrors := false
	resGetSitesResp := resJSON["getSitesResp"].(map[string]interface{})
	resSiteNames := resGetSitesResp["siteNames"].(map[string]interface{})

	for _, resSiteIdInterface := range resSiteNames["string"].([]interface{}) {
		if resSiteId := resSiteIdInterface.(string); resSiteId != "Sites-Site" {
			pwCheckUrl := fmt.Sprintf("https://%s/on/demandware.store/%s", dwHostname, resSiteId)

			pwCheckRes, err := http.Head(pwCheckUrl)
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				os.Exit(1)
			}

			if pwCheckRes.StatusCode != 401 {
				resHasErrors = true
				fmt.Println(fmt.Sprintf("-> %s (ERROR %d)", resSiteId, pwCheckRes.StatusCode))
			} else {
				fmt.Println(fmt.Sprintf("-> %s (Ok %d)", resSiteId, pwCheckRes.StatusCode))
			}
		}
	}

	if resHasErrors {
		os.Exit(1)
	}
}
