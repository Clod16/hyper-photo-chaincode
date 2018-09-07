package main

import(
	"bytes"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/rs/xid"
)


var logger = shim.NewLogger("HyperPhoto-log")



func (t *HyperPhoto) Init(stub shim.ChaincodeStubinterface) pb.Response {
	logger.Info("Initializing Hyper Photo")
	logger.SetLevel(shim.LogDebug)
	_, args := stub.GetFunctionAndParameters()
	//var err error

	// Upgrade Mode 1: leave ledger state as it was
	if len(args) == 0 {
		//logger.Info("Args correctly!!!")
		return shim.Success(nil)
	}

	return shim.Success(nil)
}


func (t *HyperPhoto) Invoke(stub shim.ChaincodeStubinterface) pb.Response {
	var err error


	function, args := stub.GetFunctionAndParameters()
	if function == "initPhoto"
	else if function == "putPhoto" {
		return t.putPhoto(stub, args)
	} else if function == "getPhoto" {
		return t.getPhoto(stub, args)
	}
	return shim.Error("Invalid invoke function name")
}

func (t *HyperPhoto) putPhoto(stub shim.ChaincodeStubinterface) pb.Response{
	logger.Debug("      putPhoto      ")

	var photoKEY string





}




