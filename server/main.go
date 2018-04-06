package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/bithinalangot/LinkList/list"
)

const (
	port = ":50051"
)

// Represent a node
type Node struct {
	data int32
	next *Node
	prev *Node
}

// Represent a linked list
type List struct {
	head *Node
	tail *Node
}

//inserting data into linked list
func (L *List) InsertNode(ctx context.Context, in *pb.NodeRequest) (*pb.NodeResponse, error) {
	newNode := &Node{
		data: in.Data,
		next: nil,
		prev: nil,
	}

	if L.head == nil && L.tail == nil {
		L.head = newNode
		L.tail = newNode
	} else {
		L.tail.next = newNode
		temp := L.tail
		L.tail = newNode
		newNode.prev = temp
	}
	return &pb.NodeResponse{Success: true}, nil
}

//Printing the linked list
func (L *List) Printing(nodes *pb.LinkRequest, stream pb.List_PrintingServer) error {
	for temp := L.head; temp != nil; temp = temp.next {
		node := &pb.Nodes{
			Node: temp.data,
		}
		if err := stream.Send(node); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterListServer(s, &List{})
	s.Serve(lis)
}
