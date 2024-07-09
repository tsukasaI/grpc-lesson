package main

import (
	"fmt"
	"log"
	"os"
	"protobuf-lesson/pb"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Sizilkoi",
		Email:       "test.com",
		Occupation:  pb.Occipation_ENGINEER,
		PhoneNumber: []string{"080-1234566667", "12312312312"},
		Project:     map[string]*pb.Company_Project{"A": &pb.Company_Project{}},
		Profile:     &pb.Employee_Text{Text: "My Name is "},
		Birthday:    &pb.Date{Year: 2000, Month: 1, Day: 1},
	}

	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Cannnot serialize", err)
	}

	fileName := "test.bin"
	if err := os.WriteFile(fileName, binData, 0666); err != nil {
		log.Fatalln("Cannnot write", err)
	}

	in, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Cannot read file", err)
	}

	readEmp := new(pb.Employee)
	err = proto.Unmarshal(in, readEmp)
	if err != nil {
		log.Fatalln("Cannot deserialize", err)
	}

	fmt.Println(readEmp)

	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Cannot marshal to json", err)
	}

	fmt.Println(out)
}
