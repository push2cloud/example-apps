package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
  "crypto/rand"
  "encoding/hex"
	"io/ioutil"
	"strings"
)

var appId = "go"

func main() {
	log.SetOutput(os.Stdout)
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "4000"
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/start", startHandler)
	log.Printf(fmt.Sprintf("Listening at %s", port))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func generateKey(n int) (string, error) {
  bytes := make([]byte, n)
  if _, err := rand.Read(bytes); err != nil {
    return "", err
  }
  return hex.EncodeToString(bytes), nil
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello from go")
}

func setValue(key string, val string) {
  client := &http.Client{}
	data := url.Values{}
  data.Set("data", val)
	request, err := http.NewRequest("PUT", os.Getenv("API_URL") + "/" + key + "?" + data.Encode(), nil)
  response, err := client.Do(request)

	defer response.Body.Close()

  if err != nil {
		log.Fatal(err)
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode >= 300 {
			log.Printf("response StatusCode:", response.StatusCode)
			log.Printf("response Headers:", response.Header)
			log.Printf("response Body:", string(body))
		}
	}
}

func getValue(key string) (string) {
  client := &http.Client{}
	request, err := http.NewRequest("GET", os.Getenv("API_URL") + "/" + key, nil)
  response, err := client.Do(request)

	defer response.Body.Close()

  if err != nil {
		log.Fatal(err)
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		log.Printf("response Body of api get:", string(body))
		if response.StatusCode >= 300 {
			log.Printf("response StatusCode:", response.StatusCode)
			log.Printf("response Headers:", response.Header)
		} else {
			return string(body)
		}
	}
	return ""
}

func deleteValue(key string) {
  client := &http.Client{}
	request, err := http.NewRequest("DELETE", os.Getenv("API_URL") + "/" + key, nil)
  response, err := client.Do(request)

	defer response.Body.Close()

  if err != nil {
		log.Fatal(err)
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode >= 300 {
			log.Printf("response StatusCode:", response.StatusCode)
			log.Printf("response Headers:", response.Header)
			log.Printf("response Body:", string(body))
		}
	}
}

func forward(handle string, w http.ResponseWriter) {
  client := &http.Client{}
	request, err := http.NewRequest("GET", os.Getenv("NEXT_APP_URL") + "/start?handle=" + handle, nil)
  response, err := client.Do(request)

	defer response.Body.Close()

  if err != nil {
		log.Fatal(err)
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode >= 300 {
			log.Printf("response StatusCode:", response.StatusCode)
			log.Printf("response Headers:", response.Header)
			log.Printf("response Body:", string(body))
		}
		fmt.Fprintln(w, string(body))
	}
}

func startHandler(w http.ResponseWriter, req *http.Request) {
  log.Printf(fmt.Sprintf("called start route..."))

  handle := req.URL.Query().Get("handle")
  if handle == "" {
    log.Printf(fmt.Sprintf("no handle... so let's start..."))

    handle, _ := generateKey(20);

    setValue(handle, appId)

		forward(handle, w)

		return
  }

	var val = getValue(handle)

	if strings.Contains(val, appId) {
		log.Printf(fmt.Sprintf("round trip completed"))

		deleteValue(handle)

		var newVal =  val + ";" + appId;

		fmt.Fprintln(w, "Round trip completed!   >>>  " + strings.Replace(newVal, ";", " => ", -1))

		return
	}

	var newVal2 =  val + ";" + appId;

	log.Printf(fmt.Sprintf(handle + " save new value " + newVal2))

	setValue(handle, newVal2)

	log.Printf(fmt.Sprintf(handle + " not already seen so forward..."))

	forward(handle, w)
}
