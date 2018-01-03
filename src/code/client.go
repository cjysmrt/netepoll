package main

import (
	_"time"
	"os"
	"os/signal"
	"syscall"
	"boost/cellnet"
	_"boost/cellnet/benchmark"
	_ "boost/cellnet/codec/pb"
	"boost/cellnet/proto/pb/gamedef"
	"boost/cellnet/socket"
	_"boost/cellnet/util"
	"boost/golog"
)

var log *golog.Logger = golog.New("test")

const benchmarkAddress = "127.0.0.1:7201"

const clientCount = 50

const benchmarkSeconds = 10


func main() {
	golog.SetLevelByString("cellnet", "error")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	queue := cellnet.NewEventQueue()

	evd := socket.NewConnector(queue).Start(benchmarkAddress)

	cellnet.RegisterMessage(evd, "gamedef.TestEchoACK", func(ev *cellnet.Event) {

		log.Infof("client recive Message happen")
		ev.Send(&gamedef.TestEchoACK{})
	})

	cellnet.RegisterMessage(evd, "coredef.SessionConnected", func(ev *cellnet.Event) {

		log.Infof("client Connected happen")
		ev.Send(&gamedef.TestEchoACK{})

	})

	queue.StartLoop()
	
	<-sc
}