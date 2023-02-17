package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var s = ""

type XYZ struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type CDL struct {
	Uid        string  `json:"uid"`
	Name       string  `json:"name"`
	Path       string  `json:"path"`
	Slope      XYZ     `json:"slope"`
	Offset     XYZ     `json:"offset"`
	Power      XYZ     `json:"power"`
	Saturation float64 `json:"saturation"`
}

type CDLPost struct {
	Uid        string  `json:"uid"`
	Name       string  `json:"name"`
	Slope      XYZ     `json:"slope"`
	Offset     XYZ     `json:"offset"`
	Power      XYZ     `json:"power"`
	Saturation float64 `json:"saturation"`
}

type OldCDL struct {
	Uid        string  `json:"uid"`
	Slope      XYZ     `json:"slope"`
	Offset     XYZ     `json:"offset"`
	Power      XYZ     `json:"power"`
	Saturation float64 `json:"saturation"`
}

type CDLs struct {
	Status map[string]string `json:"status"`
	Result []CDL             `json:"result"`
}

type uid struct {
	Uid string `json:"uid"`
}

type uids struct {
	Uids []uid `json:"uids"`
}

func getCDLs(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get(fmt.Sprintf("http://%s/api/session/colour/cdls", s))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	var dat CDLs
	json.Unmarshal(body, &dat)
	//x := make(map[string][]string)
	//dat["result"]
	uids := make([]uid, len(dat.Result))
	for i := 0; i < len(dat.Result); i++ {
		uids[i] = uid{Uid: dat.Result[i].Uid}
	}
	x := make(map[string][]uid)
	x["uids"] = uids
	t, err := json.Marshal(x)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s", t)
	//w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(t)
}

func getCDL(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/api/v1/colour/cdl/get/")
	resp, err := http.Get(fmt.Sprintf("http://%s/api/session/colour/cdls", s))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var dat CDLs
	json.Unmarshal(body, &dat)
	var rp []byte
	x := make(map[string]CDL)

	for i := 0; i < len(dat.Result); i++ {
		if dat.Result[i].Uid == id {
			dat.Result[i].Path = fmt.Sprintf("objects/cdl/%s%s", dat.Result[i].Name, ".apx")
			x["cdl"] = dat.Result[i]
			rp, err = json.Marshal(x)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	log.Printf("%s", rp)

	//Convert the body to type string
	//w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(rp)
}

func setCDL(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var dat OldCDL
	json.Unmarshal(body, &dat)

	resp, err := http.Get(fmt.Sprintf("http://%s/api/session/colour/cdls", s))
	if err != nil {
		log.Fatalln(err)
	}

	bd2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var cdls CDLs
	json.Unmarshal(bd2, &cdls)
	var updatedCDL CDLPost

	for i := 0; i < len(cdls.Result); i++ {
		if cdls.Result[i].Uid == dat.Uid {
			//updatedCDL.Name = cdls.Result[i].Name
			updatedCDL.Uid = cdls.Result[i].Uid
			updatedCDL.Name = cdls.Result[i].Name
			break
		}
	}

	updatedCDL.Offset = dat.Offset
	updatedCDL.Power = dat.Power
	updatedCDL.Saturation = dat.Saturation
	updatedCDL.Slope = dat.Slope

	x := make(map[string]CDLPost)

	var rp []byte
	x["cdl"] = updatedCDL

	rp, err = json.Marshal(x)

	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s", rp)

	resp3, err := http.Post(fmt.Sprintf("http://%s/api/session/colour/cdl", s), "application/json", bytes.NewBuffer(rp))

	if err != nil {
		log.Fatalln(err)
	}

	bd3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%s", bd3)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("{}"))

}

func main() {
	s = os.Args[1]
	log.Printf("Starting server on %s", s)
	http.HandleFunc("/api/v1/colour/cdl/list", getCDLs)
	http.HandleFunc("/api/v1/colour/cdl/get/", getCDL)
	http.HandleFunc("/api/v1/colour/cdl/set", setCDL)
	//http.HandleFunc("/", getCDL)

	http.ListenAndServe(":8008", nil)
}
