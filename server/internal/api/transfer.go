package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	db "github.com/thomaslievre/my-simple-bank/db/sqlc"
	"github.com/thomaslievre/my-simple-bank/internal/blockchain"
	"github.com/thomaslievre/my-simple-bank/internal/token"
)

type trasnferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) sendTransfer() {
	// Clé privée fictive (ne jamais utiliser en prod !)
	privateKeyHex := "4f3edf983ac636a65a842ce7c78d9aa706d3b113b37e9f1e6e7a6a5a5a5a5a5a"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Erreur clé privée: %v", err)
	}

	// Création de l'authentificateur de transaction
	chainID := big.NewInt(31337) // Hardhat local
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Erreur auth: %v", err)
	}
	auth.Context = context.Background()
	auth.Value = big.NewInt(10000000000000000) // 0.01 ETH en wei
	auth.GasLimit = 300000

	// Adresse du contrat fictive (remplace par la vraie après déploiement)
	contractAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")

	// Instanciation du binding du contrat
	instance, err := blockchain.NewTransfer(contractAddress, server.ethClient)
	if err != nil {
		log.Fatalf("Erreur instance contrat: %v", err)
	}

	// Adresse de destination fictive --> pour le test adress Account #0
	to := common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")

	// Appel de la fonction send du smart contract
	tx, err := instance.Send(auth, to)
	if err != nil {
		log.Fatalf("Erreur envoi transaction: %v", err)
	}
	fmt.Printf("Transaction envoyée: %s\n", tx.Hash().Hex())

	// Attente de la confirmation (minage)
	receipt, err := bind.WaitMined(context.Background(), server.ethClient, tx)
	if err != nil {
		log.Fatalf("Erreur attente confirmation: %v", err)
	}
	if receipt.Status != 1 {
		log.Fatalf("Transaction échouée")
	}
	fmt.Printf("Transaction confirmée dans le bloc: %d\n", receipt.BlockNumber.Uint64())
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req trasnferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)

	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if _, valid = server.validAccount(ctx, req.ToAccountID, req.Currency); !valid {
		return
	}

	// Blockchain part
	server.sendTransfer()
	//     contractAddr := common.HexToAddress("0x...") // Adresse du contrat déployé
	//     instance, err := transfer.NewTransfer(contractAddr, server.ethClient)
	//     if err != nil {
	//         ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//         return
	//     }

	// privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY_HEX")
	// if err != nil {
	//     // handle error
	// }

	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(CHAIN_ID))
	// if err != nil {
	//     // handle error
	// }

	// auth.Context = context.Background()
	// auth.Value = big.NewInt(0) // montant en wei à envoyer avec la transaction (0 si pas d'ETH à transférer)
	// auth.GasLimit = 300000     // limite de gas

	//     auth := /* bind.TransactOpts avec ta clé privée */
	//     to := common.HexToAddress("0xDESTINATION") // Adresse Ethereum du destinataire
	//     amount := big.NewInt(req.Amount)           // Attention: conversion en wei si besoin

	//     tx, err := instance.Send(auth, to, amount)
	//     if err != nil {
	//         ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//         return
	//     }

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	txResult, err := server.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, txResult)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	acc, err := server.store.GetAccount(ctx, accountID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return acc, false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return acc, false
	}

	if acc.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", acc.ID, acc.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return acc, false
	}

	return acc, true
}
