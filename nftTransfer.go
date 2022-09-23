/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"
	
	// We need contract api to make new chaincodes
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nathanmlh/nftTransfer"

)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
