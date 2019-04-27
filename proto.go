package main

import (
	"encoding/json"
	"fmt"
)

type GetPlacementExtraResponse struct {
	Data *GetPlacementExtraResponse_Data `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

type GetPlacementExtraResponse_Data struct {
	Items *PlacementExtraItems `protobuf:"bytes,1,opt,name=items" json:"items,omitempty"`
}

type PlacementExtraItems struct {
	Extras []*PlacementExtra `protobuf:"bytes,1,rep,name=extras" json:"extras,omitempty"`
	Valid  bool              `protobuf:"varint,2,opt,name=Valid" json:"Valid,omitempty"`
}

type PlacementExtra struct {
	Name string `json:"name"`
}

func main() {
	extra := &PlacementExtra{Name: "baka"}
	extras := []*PlacementExtra{extra}
	items := &PlacementExtraItems{extras, true}
	responseData := &GetPlacementExtraResponse_Data{items}
	response := &GetPlacementExtraResponse{responseData}

	bytes, err := json.Marshal(response)
	fmt.Println("struct: ", response)
	fmt.Println("bytes: ", string(bytes))
	fmt.Println("err: ", err)
}
