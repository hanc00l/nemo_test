package main

import (
	"fmt"
	"github.com/hanc00l/nemo_test/pkg/HFish/core/protocol/elasticsearch"
	"github.com/hanc00l/nemo_test/pkg/HFish/core/protocol/ftp"
	"github.com/hanc00l/nemo_test/pkg/HFish/core/protocol/httpp"
	"github.com/hanc00l/nemo_test/pkg/HFish/core/protocol/redis"
	"github.com/hanc00l/nemo_test/pkg/HFish/core/protocol/ssh"
	"github.com/hanc00l/nemo_test/pkg/HFish/core/protocol/telnet"
	"github.com/hanc00l/nemo_test/pkg/HFish/core/protocol/vnc"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	banner := `nemo test server
version: 1.0.0
forked from: HFish project

starting server...`
	fmt.Println(banner)

	startTestServer()
	time.Sleep(time.Second * 3)
	setupCloseHandler()
}

func startTestServer() {
	ftpServer := "0.0.0.0:21"
	go ftp.Start(ftpServer)
	fmt.Println("start ftp server:", ftpServer)
	sshServer := "0.0.0.0:22"
	go ssh.Start(sshServer)
	fmt.Println("start ssh server:", sshServer)
	telnetServer := "0.0.0.0:23"
	go telnet.Start(telnetServer)
	fmt.Println("start telnet server:", telnetServer)
	httpServer := "0.0.0.0:80"
	go httpp.Start(httpServer, httpp.ServerHttp)
	fmt.Println("start http server:", httpServer)
	httpsServer := "0.0.0.0:443"
	go httpp.Start(httpsServer, httpp.ServerHttps)
	fmt.Println("start https server:", httpsServer)
	vncServer := "0.0.0.0:5900"
	go vnc.Start(vncServer)
	fmt.Println("start vnc server:", vncServer)
	redisServer := "0.0.0.0:6379"
	go redis.Start(redisServer)
	fmt.Println("start redis server:", redisServer)
	elasticsearchServer := "0.0.0.0:9200"
	go elasticsearch.Start(elasticsearchServer)
	fmt.Println("start elasticsearch server:", elasticsearchServer)
	actuatorServer := "0.0.0.0:18080"
	go httpp.Start(actuatorServer, httpp.ServerActuator)
	fmt.Println("start actuator server:", actuatorServer)
}

func setupCloseHandler() {
	quitSignal := make(chan os.Signal, 1)
	fmt.Println("Press Ctrl+C to Terminal...")
	signal.Notify(quitSignal, syscall.SIGTERM, syscall.SIGINT)
	<-quitSignal
	fmt.Println("Ctrl+C pressed in Terminal...")
	os.Exit(0)
}
