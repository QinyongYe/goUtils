package main

import (
	"fmt"
	"encoding/base64"
	"net/http"
	"io/ioutil"
)

func getAuthHeader(username, org, passwd string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s@%s:%s", username, org, passwd)))
}

func main() {
	vappids := []string{
		"2ce84708-cfba-46ad-aab7-f6202dab5d96",
		"07f40418-12a7-4fd9-94f5-fb10079c3025",
		"847bc0e6-524e-47bd-8cf2-50997ea71fcd",
		"68556ef0-8518-4e58-9d4d-6e63eea45bb7"}

	vUrl := "https://tech-cloud.microstrategy.com/api/vApp/vapp-" + vappids[0] + "/leaseSettingsSection/"
	header := http.Header{}
	header.Add("Authorization", getAuthHeader("qye", "technology", "mstr.1984"))
	header.Add("Accept", "application/*+xml;version=1.5")
	req, error := http.NewRequest("GET", vUrl, nil)
	if error != nil {
		panic(error)
	}
	req.Header = header
	res, error := http.DefaultClient.Do(req)
	if error != nil {
		panic(error)
	}

	str, error := ioutil.ReadAll(res.Body)
	if error != nil {
		panic(error)
	}
	fmt.Println(string(str))
}

//r = requests.post('https://tech-cloud.microstrategy.com/api/sessions', verify=False, headers=headers)
//headers = {'x-vcloud-authorization' : r.headers['x-vcloud-authorization'], 'Accept': 'application/*+xml;version=1.5'}
//
//vapp_uuids = [ '2ce84708-cfba-46ad-aab7-f6202dab5d96','07f40418-12a7-4fd9-94f5-fb10079c3025' , '847bc0e6-524e-47bd-8cf2-50997ea71fcd' , '68556ef0-8518-4e58-9d4d-6e63eea45bb7' ]
//
//for vapp_id in vapp_uuids:
//    v_url = 'https://tech-cloud.microstrategy.com/api/vApp/vapp-%s/leaseSettingsSection/' % vapp_id
//
//    r = requests.get(v_url, headers=headers, verify=False);
//    r = requests.put(v_url, headers=headers, verify=False, data=r.text)
//
//    v_url = etree.fromstring(r.text).get('href')
//
//    for i in xrange(60):
//        r = requests.get(v_url, headers=headers, verify=False);
//        if etree.fromstring(r.text).get('status') == 'success':
//            break
//        time.sleep(1)
//
