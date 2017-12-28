package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dkumor/acmewrapper"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloServer)

	sslHost := os.Getenv("SSL_HOST")

	if sslHost != "" {
		w, err := acmewrapper.New(acmewrapper.Config{
			Domains: []string{sslHost},
			Address: ":443",

			TLSCertFile: "cert.pem",
			TLSKeyFile:  "key.pem",

			RegistrationFile: "user.reg",
			PrivateKeyFile:   "user.pem",

			TOSCallback: acmewrapper.TOSAgree,
		})

		if err != nil {
			log.Fatal("acmewrapper: ", err)
		}

		tlsconfig := w.TLSConfig()

		listener, err := tls.Listen("tcp", ":443", tlsconfig)
		if err != nil {
			log.Fatal("Listener: ", err)
		}
		// To enable http2, we need http.Server to have reference to tlsconfig
		// https://github.com/golang/go/issues/14374
		server := &http.Server{
			Addr:      ":443",
			Handler:   mux,
			TLSConfig: tlsconfig,
		}
		server.Serve(listener)
	} else {
		log.Fatal(http.ListenAndServe(":8080", mux))
	}

}
