package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "grpc-go/src/proto"
)

func main() {

	const address = "127.0.0.1:10081"

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	// User Request
	userClient := pb.NewUserProviderClient(conn)
	reply, err := userClient.GetByUserId(context.Background(), &pb.UserIdRequest{Id: 1})
	if err != nil {
		log.Println(err)
	}

	// Dictionary Request
	dictClient := pb.NewDictionaryServiceClient(conn)
	dict, err := dictClient.SelectorDictionary(context.Background(), &pb.DictionaryRequest{Parent: "root", Language: "zh-CN"})
	if err != nil {
		log.Println(err)
	}

	// Dictionary Page Request
	r := &pb.PageRequest{
		PageNumber: 1,
		PageSize:   2,
		Sort:       "id,desc",
	}

	dictPage, err := dictClient.SelectorDictionaryPage(context.Background(), &pb.DictionaryPageRequest{Parent: "root", Language: "zh-CN", Page: r})
	if err != nil {
		log.Println(err)
	}
	log.Println("Java Return Data : ", reply.Id, reply.Name, reply.Email)
	log.Println("Java Return Dictionary Data : ", dict.GetDictionary())
	log.Println("Java Return Dictionary Page Data : ", dictPage.Content, dictPage.First, dictPage.Last, dictPage.TotalElements, dictPage.TotalPages)

}
