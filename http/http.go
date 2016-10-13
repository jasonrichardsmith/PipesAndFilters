package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"log"
)

var (
	Exec string
	Port string
)

func init() {
	flag.StringVar(&Exec, "exec-service", os.Getenv("EXEC_SERVICE"), "Executable")
	flag.StringVar(&Port, "port", os.Getenv("PORT"), "Port to use")
}

func main() {
	flag.Parse()
	err := http.ListenAndServe(fmt.Sprintf(":%v", Port), http.HandlerFunc(ExecHandler))
	log.Print(err)
}

func ExecHandler(w http.ResponseWriter, r *http.Request) {
	c := exec.Command("bash", "-c", Exec)
	in, err := c.StdinPipe()
	if err != nil {
		log.Print(err)
	}
	out, err := c.StdoutPipe()
	if err != nil {
		log.Print(err)
	}
	if err := c.Start(); err != nil {
		log.Print(err)
	}
	io.Copy(in, r.Body)
	io.Copy(w, out)
	if err := c.Wait(); err != nil {
		log.Print(err)
	}
	c.Run()
}
