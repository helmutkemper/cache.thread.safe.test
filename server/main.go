package main

import (
	cache "github.com/helmutkemper/cache.thread.safe/server"
	"log"
	"os"
	"sync"
)

func main() {
	var err error
	var wg sync.WaitGroup

	cacheServer := &cache.Server{}
	err = cacheServer.Init()
	if err != nil {
		log.Printf("error: %v", err.Error())
	}

	log.Printf("starting server main(): %v", os.Getenv("DEBUG_NAME"))

	wg.Add(1)
	wg.Wait()
}

//func dial() {
//  time.Sleep(1 * time.Second)
//  conn, err := grpc.Dial("0.0.0.0:3000", grpc.WithInsecure())
//  if err != nil {
//    panic(err)
//  }
//  defer conn.Close()
//
//  c := cache.NewSyncCacheDaraServerClient(conn)
//  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//  defer cancel()
//
//  r, err := c.TakeMyIp(ctx, &cache.TakeMyIpRequest{IP: "127.0.0.1"})
//  if err != nil {
//    panic(err)
//  }
//
//  log.Printf("Ok: %v", r.Ok)
//}

func _main() {
	var err error
	//go func() {
	//  time.Sleep(3*time.Second)
	//  for i := 0; i != 170; i += 1 {
	//    go dial()
	//  }
	//}()

	//port := cache.KServerPort
	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()

	cacheServer := &cache.Server{}
	err = cacheServer.Init()
	if err != nil {
		log.Printf("error: %v", err.Error())
	}

	log.Printf("starting server: %v", os.Getenv("DEBUG_NAME"))

	//cache.RegisterSyncCacheDaraServerServer(s, cacheServer)
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
