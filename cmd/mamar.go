package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	proto "github.com/CSC354/mamar/proto"
	"google.golang.org/grpc"
)

var mp = map[string]string{}

const MASTER = 8000

type Mamar struct {
	proto.UnimplementedMamarServer
}

func (*Mamar) GetPort(ctx context.Context, s *proto.Service) (*proto.Port, error) {
	port := proto.Port{Address: mp[s.Name]}
	return &port, nil
}

func main() {
	read()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterMamarServer(grpcServer, &Mamar{})
	grpcServer.Serve(lis)
}

func read() {
	if len(os.Args) < 2 {
		log.Fatal("No file was provided as argument")
	}
	name := os.Args[1]
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	err = scan(scanner)
	if err != nil {
		log.Fatal(err)
	}
}

func scan(scanner *bufio.Scanner) error {
	var c int
	for scanner.Scan() {
		c++
		if scanner.Text()[0] == '#' {
			continue
		}
		columns := strings.Split(scanner.Text(), " ")
		nlines := len(columns)
		if nlines != 2 {
			return fmt.Errorf("line %d: Schema break", nlines)
		}
		i, err := strconv.Atoi(columns[1])
		if err != nil {
			return fmt.Errorf("line %d: %s is not a valid port numbering", c, columns[1])
		} else if i == MASTER {

			return fmt.Errorf("line %d: can not map to master port", c)
		}
		mp[columns[0]] = columns[1]
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil

}
