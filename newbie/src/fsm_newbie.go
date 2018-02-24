package main

import (
	"github.com/looplab/fsm"
	"log"
	"encoding/json"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}
type DoorInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name:"sureOpen",Src:[]string{"closed","open"},Dst:"open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
			{Name: "ready",Src:[]string{"open"},Dst:"ready"},
		},
		fsm.Callbacks{
			"before_open": func(e *fsm.Event) { d.enterState(e,"bOpen") },
			"before_close": func(e *fsm.Event) { d.enterState(e,"bClose") },
		},
	)

	return d
}

func (d *Door) enterState(e *fsm.Event,name string) {
	log.Printf("[%s]The door to %s is %s\n",name, d.To, e.Dst)
}

func main() {
	door := NewDoor("heaven")
	log.Printf("Current Status:%s",door.FSM.Current())


	err := door.FSM.Event("sureOpen")
	if err != nil {
		log.Printf("Error %v",err)
	}
	log.Printf("Current Status:%s",door.FSM.Current())

	err = door.FSM.Event("open")
	if err != nil {
		log.Printf("Error %v",err)
	}
	log.Printf("Current Status:%s",door.FSM.Current())
	err=door.FSM.Event("ready")
	if err!=nil{
		log.Fatalf("to ready status failed.%v",err)
	}
	log.Printf("Current Status:%s",door.FSM.Current())

	doorInfo:=&DoorInfo{
		Name:"SmartDoor",
		Age:2,
	}
	jsons,err:=json.Marshal(doorInfo)
	if err!=nil{
		log.Fatalf("Door Info To Json Failed,%v",err)
	}
	log.Printf("DoorInfo:%s",string(jsons))
}