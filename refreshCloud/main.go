package main

import (
	"fmt"
	"encoding/base64"
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"os"
	"bytes"
)

var usage = `Usage:
refreshCloud [user] [password]`

type VAppRecord struct {
	Href string `xml:"href,attr"`
	Name string `xml:"name,attr"`
}
type VAppRecords struct {
	VAppRecord   []VAppRecord
}

func getAuthHeader(username, org, passwd string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s@%s:%s", username, org, passwd)))
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	var user = os.Args[1]
	var password = os.Args[2]

	header := http.Header{}
	header.Add("Authorization", getAuthHeader(user, "technology", password))
	header.Add("Accept", "application/*+xml;version=1.5")

	// login
	login, error := http.NewRequest("POST", "https://tech-cloud.microstrategy.com/api/sessions", nil)
	if error != nil {
		panic(error)
	}
	login.Header = header
	res, error := http.DefaultClient.Do(login)
	if error != nil {
		panic(error)
	}
	//fmt.Print("login response: ")
	//fmt.Println(res.Header)

	// list vms
	header.Add("x-vcloud-authorization", res.Header.Get("x-vcloud-authorization"))
	listVms, error := http.NewRequest("GET", "https://tech-cloud.microstrategy.com/api/query?type=vApp&filter=(ownerName==qye)", nil)
	if error != nil {
		panic(error)
	}
	listVms.Header = header
	res, error = http.DefaultClient.Do(listVms)
	if error != nil {
		panic(error)
	}

	vapps, error := ioutil.ReadAll(res.Body)
	if error != nil {
		panic(error)
	}
	v := VAppRecords{}
	error = xml.Unmarshal(vapps, &v)
	if error != nil {
		panic(error)
	}
	fmt.Println(v)

	for _, vapp := range v.VAppRecord {
		// refresh
		fmt.Println("refresh ", vapp)
		refreshget, error := http.NewRequest("GET", vapp.Href, nil)
		if error != nil {
			panic(error)
		}
		refreshget.Header = header
		res, error := http.DefaultClient.Do(refreshget)
		if error != nil {
			panic(error)
		}
		text, error := ioutil.ReadAll(res.Body)
		if error != nil {
			panic(error)
		}
		//fmt.Println("get response: ", string(text))

		refreshpost, error := http.NewRequest("PUT", vapp.Href, bytes.NewReader(text))
		if error != nil {
			panic(error)
		}
		refreshpost.Header = header
		_, error = http.DefaultClient.Do(refreshpost)
		if error != nil {
			panic(error)
		}
//		text, error = ioutil.ReadAll(res1.Body)
//		if error != nil {
//			panic(error)
//		}
//		fmt.Println("put response: ", string(text))
	}
}
