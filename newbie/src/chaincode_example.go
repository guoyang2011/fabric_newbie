package main

import (
	"github.com/hyperledger.back/fabric/core/chaincode/shim"
	pb "github.com/hyperledger.back/fabric/protos/peer"
	"fmt"
	"strconv"
	"golang.org/x/crypto/sha3"
	"crypto/sha256"
	"crypto/sha512"
	"log"
	"time"
)

type SimpleChainCode struct {
}
func (t *SimpleChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response{
	fmt.Printf("ex02 Init")
	_,args:=stub.GetFunctionAndParameters()
	var A,B string
	var AVal,BVal int
	var err error
	if len(args)!=4{
		return shim.Error("Incorrect number of arguments, Expecting 4")
	}
	A=args[0]
	AVal,err=strconv.Atoi(args[1])
	if err!=nil{
		return shim.Error("Expecting integer value for asset holding")
	}
	B=args[2]
	BVal,err=strconv.Atoi(args[3])
	if err!=nil{
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d , Bval = %d \n",AVal,BVal)
	err =stub.PutState(A,[]byte(strconv.Itoa(AVal)))

	if err !=nil{
		return shim.Error(err.Error())
	}
	err = stub.PutState(B,[]byte(strconv.Itoa(BVal)))
	if err !=nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
func (t *SimpleChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	fmt.Printf("ex02 Invoke")
	function,args:=stub.GetFunctionAndParameters()
	if function == "invoke" {
		return t.invoke(stub,args)
	}else if function == "delete"{
		return t.delete(stub,args)
	}else if function == "query"{
		return t.query(stub,args)
	}
	return shim.ERROR("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

func (t *SimpleChainCode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	return nil
}
func (t *SimpleChainCode) delete(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	return nil
}
func (t *SimpleChainCode) query(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	return nil
}
func main() {
	ch := make(chan bool)
	//err:=shim.Start(new(SimpleChainCode))
	//if err!=nil{
	//	fmt.Printf("Error starting Simple chaincode:%s",err)
	//}
	checksum:=sha256.Sum256([]byte("Hello world...."))
	log.Printf("checksum:%s",checksum)
	go func(){
		for {
			select {
			case <- ch:
				log.Printf("receive channel message ...")
			}
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second)
	close(ch)

	time.Sleep(time.Second*1000)
}
