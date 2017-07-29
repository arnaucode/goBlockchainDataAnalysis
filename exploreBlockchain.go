package main

import (
	"fmt"
	"strconv"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcrpcclient"
)

func explore(client *btcrpcclient.Client, blockHash string) {
	var realBlocks int
	var nOrigin NodeModel
	nOrigin.Id = "origin"
	nOrigin.Label = "origin"
	nOrigin.Title = "origin"
	nOrigin.Group = "origin"
	nOrigin.Value = 1
	nOrigin.Shape = "dot"
	saveNode(nodeCollection, nOrigin)

	for blockHash != "" {
		//generate hash from string
		bh, err := chainhash.NewHashFromStr(blockHash)
		check(err)
		block, err := client.GetBlockVerbose(bh)
		check(err)

		if block.Height > 160 {
			for k, txHash := range block.Tx {
				if k > 0 {
					realBlocks++
					fmt.Print("Block Height: ")
					fmt.Print(block.Height)
					fmt.Print(", num of Tx: ")
					fmt.Print(k)
					fmt.Print("/")
					fmt.Println(len(block.Tx) - 1)

					th, err := chainhash.NewHashFromStr(txHash)
					check(err)
					tx, err := client.GetRawTransactionVerbose(th)
					check(err)

					var originAddresses []string
					var outputAddresses []string
					for _, Vo := range tx.Vout {
						//if Vo.Value > 0 {
						for _, outputAddr := range Vo.ScriptPubKey.Addresses {
							outputAddresses = append(outputAddresses, outputAddr)
							var n2 NodeModel
							n2.Id = outputAddr
							n2.Label = outputAddr
							n2.Title = outputAddr
							n2.Group = strconv.FormatInt(block.Height, 10)
							n2.Value = 1
							n2.Shape = "dot"
							saveNode(nodeCollection, n2)
						}
						//}
					}
					for _, Vi := range tx.Vin {
						th, err := chainhash.NewHashFromStr(Vi.Txid)
						check(err)
						txVi, err := client.GetRawTransactionVerbose(th)
						check(err)
						if len(txVi.Vout[Vi.Vout].ScriptPubKey.Addresses) > 0 {
							for _, originAddr := range txVi.Vout[Vi.Vout].ScriptPubKey.Addresses {
								originAddresses = append(originAddresses, originAddr)
								var n1 NodeModel
								n1.Id = originAddr
								n1.Label = originAddr
								n1.Title = originAddr
								n1.Group = string(block.Height)
								n1.Value = 1
								n1.Shape = "dot"
								saveNode(nodeCollection, n1)

								for _, outputAddr := range outputAddresses {
									var e EdgeModel
									e.From = originAddr
									e.To = outputAddr
									e.Label = txVi.Vout[Vi.Vout].Value
									e.Txid = tx.Txid
									e.Arrows = "to"
									e.BlockHeight = block.Height
									saveEdge(edgeCollection, e)

									//date analysis
									//dateAnalysis(e, tx.Time)
									//hour analysis
									hourAnalysis(e, tx.Time)
								}
							}
						} else {
							originAddresses = append(originAddresses, "origin")
						}

					}
					fmt.Print("originAddresses: ")
					fmt.Println(len(originAddresses))
					fmt.Print("outputAddresses: ")
					fmt.Println(len(outputAddresses))
				}
			}
		}
		//set the next block
		blockHash = block.NextHash
	}

	fmt.Print("realBlocks (blocks with Fee and Amount values): ")
	fmt.Println(realBlocks)
	fmt.Println("reached the end of blockchain")
}