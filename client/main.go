package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/songgao/water"
)

func main() {
	ip := "10.10.10.1"

	iface, err := createTun(ip)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", "${YOUR_PUBLIC_VPN_SERVER_ADDR}:8990")
	if err != nil {
		panic(err)
	}

	go listenTcpConnection(conn, iface)
	go listenTunDevice(conn, iface)

	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, os.Interrupt, syscall.SIGTERM)
	<-termSignal
	fmt.Println("terminating")
}

func createTun(ip string) (*water.Interface, error) {
	iface, err := water.New(water.Config{DeviceType: water.TUN})
	if err != nil {
		return nil, err
	}

	Run(fmt.Sprintf("sudo ip addr add %s/24 dev %s", ip, iface.Name()))
	Run(fmt.Sprintf("sudo ip link set dev %s up", iface.Name()))
	return iface, nil
}

func listenTunDevice(conn net.Conn, iface *water.Interface) {
	for {
		m := make([]byte, 1500)
		n, err := iface.Read(m)
		if err != nil {
			log.Printf("error occurred while reading interface: %v", err)
			continue
		}
		message := m[:n]
		conn.Write(message)
	}
}

func listenTcpConnection(conn net.Conn, iface *water.Interface) {
	for {
		m := make([]byte, 1500)
		n, err := conn.Read(m)
		if err != nil {
			log.Printf("error occurred while reading connection: %v", err)
			continue
		}
		message := m[:n]
		iface.Write(message)
	}
}

func Run(command string) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String(), fmt.Errorf("command execution failed: %w", err)
	}
	return stdout.String(), nil
}
