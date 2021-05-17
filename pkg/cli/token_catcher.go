package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
)

func StartServer(config Config) {
	serverAddress := fmt.Sprintf("localhost:%v", config.EmbeddedServerConfig.Port)

	http.HandleFunc("/sso-callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code != "" {
			request, err := BuildTokenExchangeRequest(config, code)
			if err == nil {
				var resp *http.Response
				var body []byte
				resp, err = http.DefaultClient.Do(request)
				if err == nil {
					body, err = ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
					defer resp.Body.Close()
					if resp.StatusCode == 200 {
						content, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
						switch content {
						case "application/json":
							var f interface{}
							json.Unmarshal(body, &f)
							m := f.(map[string]interface{})
							fmt.Println("Here's your Token:")
							fmt.Println(m["access_token"].(string))
						default:
							fmt.Println("Printing out raw body:")
							fmt.Println(body)
						}
					} else {
						err = fmt.Errorf("invalid Status code (%v)", resp.StatusCode)
					}
					fmt.Fprintf(w, "You can close this page. The token is in your CLI logs...")
					CloseApp.Done()
				}
			}
		}
	})

	go func() {
		log.Print("Booting up the server")
		if err := http.ListenAndServe(serverAddress , nil); err != nil {
			log.Fatalf("Unable to start server: %v\n", err)
			CloseApp.Done()
		}
	}()
}