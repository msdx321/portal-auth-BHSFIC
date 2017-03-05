package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	boolToken bool
	portalURL = "http://192.168.1.254:8888"
	token     = "MjAxNDAxMDZ8c2hpbWVpd2Fuc3Vp"
	userID    string
	userPass  string
)

func main() {
	parseFlags()
	if boolToken {
		mResp := postForm()
		parseResult(mResp)
	} else {
		encodeInfoToBase64()
		mResp := postForm()
		parseResult(mResp)
	}
}

func parseFlags() {
	boolToken = *flag.Bool("lazy", false, "Use built-in token instead")
	userID = *flag.String("userid", "none", "User ID")
	userPass = *flag.String("userpass", "none", "User Password")
	flag.Parse()
}

func postForm() *http.Response {
	mForm := url.Values{
		"actionType":    {"umlogin"},
		"authorization": {token},
		"language":      {"0"},
		"userIpMac":     {" "},
	}
	resp, err := http.PostForm(portalURL, mForm)
	checkErr(err)
	defer resp.Body.Close()
	return resp
}

func parseResult(mResp *http.Response) {
	mRespBody, err := ioutil.ReadAll(mResp.Body)
	checkErr(err)
	mBodyString := string(mRespBody)
	if strings.Contains(mBodyString, "Login successful") || strings.Contains(mBodyString, "User is online") {
		fmt.Println("Login Successful")
	} else {
		fmt.Println("Something went wrong")
		fmt.Println("Resp Body:" + mBodyString)
	}
}

func encodeInfoToBase64() {
	mInfo := []byte(userID + "|" + userPass)
	token = base64.StdEncoding.EncodeToString(mInfo)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
