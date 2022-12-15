package main

import (
	"bufio"
	"context"
	"fmt"
	proto "github.com/CSC354/mamar/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var mp = map[string]string{}

const MASTER = 8000

type Mamar struct {
	proto.UnimplementedMamarServer
}

func (*Mamar) GetPort(ctx context.Context, s *proto.Service) (*proto.Port, error) {
	if mp[s.Name] == "db" {
		p := proto.Port{Address: fmt.Sprintf("server=qaida;user id=sa;password=rBwiY3JgqmG26q@;port=1433;database=%s;", s.Name)}
		return &p, nil
	}
	port := proto.Port{Address: fmt.Sprintf("%s:%s", s.Name, mp[s.Name])}
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
		if err != nil && columns[1] != "db" {
			return fmt.Errorf("line %d: %s is not a valid port numbering", c, columns[1])
		} else if i == MASTER {
			return fmt.Errorf("line %d: can not map to master port", c)
		}
		if v, ok := mp[columns[0]]; ok {
			return fmt.Errorf("line %d Port %s occured more than once", c, v)
		}

		// if columns[1] == "db" {
		// 	mp[columns[0]] = fmt.Sprintf("server=qaida;user id=sa;password=rBwiY3JgqmG26q@;port=1433;database=%s;", columns[0])
		// } else {
		// 	mp[columns[0]] = columns[1]
		// }

	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil

}
