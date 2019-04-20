// package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	var Name, TradeLocation, TradeQuality strings, ID   // 定义了名字，交易地点，交易性质, ID
	var TradeAmount  int // 定义了交易金额（交易金额大于0）
	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
    //名字参数0，交易地点参数1，交易性质参数2，ID参数3，交易金额参数4

	Name = args[0]
	TradeLocation = args[1]
    TradeQuality = args[2]
	ID = arg[3]

    TradeAmount, err = strconv.Atoi(args[4])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}

	fmt.Printf("Name = %s, TradeLocation = %s, TradeQuality = %s, ID = %s, TradeAmount = %d", 
    Name, TradeLocation, TradeQuality, ID, TradeAmount)

	// Write the state to the ledger，写入账本

	err = stub.PutState(ID, Name ,TradeLocation, TradeQuality, []byte(strconv.Itoa(TradeAmount))) //改动，把key设为ID
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "query" {
		//查询ID，返回该ID历史记录
		return t.query(stub, args)
	} else if function == "model" {
		// 用于判断客户数据是否合规
		return t.model(stub, args)
	} else if function == "entering" {
		// 录入和更新客户数据
		return t.entering(stub, args)
	} 

	return shim.Error("Invalid invoke function name. Expecting \"query\" \"model\" \"entering\"")
}

func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var Name, TradeLocation, TradeQuality strings // 定义了名字，交易地点，交易性质
	var ID, TradeAmount  int // 定义了身份ID，交易金额，（最后定义ID只能是固定长度，交易金额大于0）
	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	Name = args[0]
	TradeLocation = args[1]
    TradeQuality = args[2]
	ID = args[3]  
	TradeAmount = args[4] ////注意，该参数声明时将int转为了string

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Name, TradeLocation, TradeQuality, TradeAmountbyte, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get ID")
	}
	if Name == nil {
		return shim.Error("Name not found")
	}
	if TradeLocation == nil {
		return shim.Error("TradeLocation not found")
	}
	if TradeQuality == nil {
		return shim.Error("TradeQuality not found")
	}
	if TradeAmountbyte == nil {
		return shim.Error("TradeAmount not found")
	}
	TradeAmount, _ = strconv.Atoi(string(TradeAmountbyte))
	
	fmt.Printf("Name = %s, TradeLocation = %s, TradeQuality = %s, ID = %s, TradeAmount = %d", 
    Name, TradeLocation, TradeQuality, ID, TradeAmount)

	return shim.Success(nil)
}

	func (t *SimpleChaincode) model(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var Name, TradeLocation, TradeQuality， ID strings // 定义了名字，交易地点，交易性质，ID
	var  TradeAmount  int // 定义了交易金额（交易金额大于0）
	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	Name = args[0]
	TradeLocation = args[1]
    TradeQuality = args[2]
	ID = args[3]  
	TradeAmount = args[4] ////注意，该参数声明时将int转为了string

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Name, TradeLocation, TradeQuality, TradeAmountbyte, err := stub.GetState(ID)

	if TradeAmountbyte == nil {
		return shim.Error("TradeAmount not found")
	}
	TradeAmount, _ = strconv.Atoi(string(TradeAmountbyte))

	if TradeAmount > 5000 {
		return shim.Error("Trade is not legal")
	}
}


// 录入客户数据
func (t *SimpleChaincode) entering(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var Name, TradeLocation, TradeQuality ID strings // 定义了名字，交易地点，交易性质
	var TradeAmount  int // 定义了身份ID，交易金额，（最后定义ID只能是固定长度，交易金额大于0）
	var err error
	
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	Name = args[0]
	TradeLocation = args[1]
    TradeQuality = args[2]
	ID = args[3]  
	TradeAmount = strconv.Atoi(args[4]) //将字符串转换为int
	
	Name, TradeLocation, TradeQuality, TradeAmountbyte, err := stub.GetState(ID)
	if err != nil {
	//如果ledger上不存在客户的历史信息
		fmt.Printf("This is a new ID, which does not exist in the ledger.")
		
		fmt.Printf("Please enter customer transaction information")
		fmt.Printf("Please enter customer name:")
		fmt.Scanln(Name)
		fmt.Printf("Please enter trade location:")
		fmt.Scanln(TradeLocation)
		fmt.Printf("Please enter trade quality:")
		fmt.Scanln(TradeQuality)
		fmt.Printf("Please enter customer ID:")
		fmt.Scanln(ID)	
		fmt.Printf("Please enter trade amount:")
		fmt.Scanln(TradeAmount)
		//写入客户信息
		err= stub.PutState(ID,Name, TradeLocation, TradeQuality,[]byte(strconv.Itoa(TradeAmount)))
		if(err!=nil){
			return shim.Error(err.Error())
			fmt.Printf("Failed to write in ledger.")
		}
	}
	else{
		//调用upgrade函数
	}


	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


