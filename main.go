package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)
func serveApp() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello")

	})
	s := http.Server{
		Addr: ":8081",
		Handler: mux,
	}
	fmt.Printf("启动APP服务\n")
	go func() {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-quit:
			fmt.Printf("收到信号: %v\n", sig)
			s.Shutdown(context.Background())
		}
	}()
	return s.ListenAndServe()
}


func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(serveApp)
	err := g.Wait()
	fmt.Println(err)
	fmt.Println(ctx.Err())
}