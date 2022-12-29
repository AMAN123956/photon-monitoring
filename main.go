package main

import (
   "bytes"
   "fmt"
   "io"
	"io/ioutil"
   //  "html"
   "os"
   "log"
   "net/http"
	// "github.com/gorilla/mux"
   "github.com/rs/cors"
   // "strings"
   // "encoding/json"
   // "reflect"
) 

type Payload struct {
	Timestamp                   float64 `json:"timestamp"`
	Type                        string  `json:"type"`
	MediaType                   string  `json:"mediaType"`
	Jitter                      float64 `json:"jitter"`
	PacketsLost                 int     `json:"packetsLost"`
	PacketsReceived             int     `json:"packetsReceived"`
	BytesReceived               int     `json:"bytesReceived"`
	LastPacketReceivedTimestamp int64   `json:"lastPacketReceivedTimestamp"`
	JitterBufferDelay           float64 `json:"jitterBufferDelay"`
	FramesReceived              int     `json:"framesReceived"`
	FrameWidth                  int     `json:"frameWidth"`
	FrameHeight                 int     `json:"frameHeight"`
	FramesPerSecond             int     `json:"framesPerSecond"`
	KeyFramesDecoded            int     `json:"keyFramesDecoded"`
	FramesDropped               int     `json:"framesDropped"`
	State                       string  `json:"state"`
	MessagesSent                int     `json:"messagesSent"`
	BytesSent                   int     `json:"bytesSent"`
	MessagesReceived            int     `json:"messagesReceived"`
}

//main function
func main() {
   // Init the mux router
   mux := http.NewServeMux()
  
   // route to handle request
   mux.HandleFunc("/make", func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      requestBody, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
      sb := string(requestBody)
      fmt.Println("request Body",sb)
      resp, err := http.Post("http://localhost:9091/metrics/job/some_job","application/json",bytes.NewBuffer(requestBody))
      if err != nil {
         log.Fatalln(err)
      }

      // close response body
      defer resp.Body.Close()

      // we read the response body
      responseBody, err := ioutil.ReadAll(resp.Body)

      // handle error
      if err != nil {
         log.Fatalln(err)
	      os.Exit(1)
      }

      // log.Printf(sb)
      fmt.Println(string(responseBody))
      fmt.Println("Status Code: ",resp.StatusCode)
  })


   handler := cors.Default().Handler(mux)
   log.Println("Listening on localhost:8080")
   log.Fatal(http.ListenAndServe(":8080", handler))
}