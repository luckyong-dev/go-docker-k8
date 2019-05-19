package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"

	"github.com/julienschmidt/httprouter"
)

type config struct {
	redisAddr      string
	httpServerAddr string
}

type handler struct {
	redis *redis.Pool
}

func loadConfig() *config {
	return &config{
		redisAddr:      os.Getenv("REDIS_ADDRESS"),
		httpServerAddr: os.Getenv("HTTP_ADDRESS"),
	}
}

func newRedis(addr string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}

func newHandler(redis *redis.Pool) *handler {
	return &handler{
		redis: redis,
	}
}

func (h *handler) getHello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello World"))
}

func (h *handler) incr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	conn := h.redis.Get()
	incr, err := redis.Int64(conn.Do("INCR", "data"))
	conn.Close()
	if err != nil {
		fmt.Fprintf(w, "err: %s\n", err)
		return
	}
	fmt.Fprintf(w, "incr: %d", incr)
}

func main() {
	var (
		cfg     = loadConfig()
		redis   = newRedis(cfg.redisAddr)
		handler = newHandler(redis)
	)

	r := httprouter.New()
	r.GET("/hello", handler.getHello)
	r.GET("/incr", handler.incr)

	server := &http.Server{Addr: cfg.httpServerAddr, Handler: r}
	server.ListenAndServe()
}
