package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	_ "net/http/pprof"
)
func serveApp() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello")

	})
	fmt.Printf("启动APP服务\n")
	http.ListenAndServe(":8081", mux)
	return errors.New("app_test")
}

func serveDebug() error {
	fmt.Printf("启动debug服务\n")
	http.ListenAndServe(":8080", http.DefaultServeMux)
	return errors.New("debug_test")
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(serveDebug)
	g.Go(serveApp)
	err := g.Wait()
	fmt.Println(err)
	fmt.Println(ctx.Err())
}