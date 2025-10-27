package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/thomaslievre/my-simple-bank/db/mock"
	db "github.com/thomaslievre/my-simple-bank/db/sqlc"
	"github.com/thomaslievre/my-simple-bank/internal/token"
	"github.com/thomaslievre/my-simple-bank/util"
)

func TestCreateTransferAPI(t *testing.T) {
	amount := int64(10)

	user1, _ := randomUser(t)
	account1 := randomAccount(user1.Username)

	user2, _ := randomUser(t)
	account2 := randomAccount(user2.Username)

	account1.Currency = util.USD
	account2.Currency = util.USD

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"from_account_id": account1.ID,
				"to_account_id":   account2.ID,
				"amount":          amount,
				"currency":        util.USD,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq((account1.ID))).
					Times(1).
					Return(account1, nil)

				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq((account2.ID))).
					Times(1).
					Return(account2, nil)

				arg := db.TransferTxParams{
					FromAccountID: account1.ID,
					ToAccountID:   account2.ID,
					Amount:        amount,
				}

				store.EXPECT().
					TransferTx(gomock.Any(), gomock.Eq(arg)).
					Times(1)

			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)

			tc.buildStubs(store)

			//  start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/transfers/create"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

// func TestSendTransfer(t *testing.T) {
// 	// Adresse du contrat déployé sur Hardhat local
// 	contractAddr := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
// 	ethClient, err := blockchain.NewEthereumClient("http://localhost:8545")
// 	if err != nil {
// 		t.Fatalf("failed to connect to eth client: %v", err)
// 	}

// 	instance, err := blockchain.NewTransfer(contractAddr, ethClient)
// 	if err != nil {
// 		t.Fatalf("failed to bind contract: %v", err)
// 	}

// 	// Adresse de destination fictive
// 	to := common.HexToAddress("0x0000000000000000000000000000000000000001")
// 	amount := big.NewInt(10000000000000000) // 0.01 ETH

// 	// Appel de la fonction Send (adapter selon ton binding)
// 	// tx, err := instance.Send(auth, to, amount)
// 	// if err != nil {
// 	//     t.Fatalf("failed to send transfer: %v", err)
// 	// }

// 	// Vérifie que la transaction a été envoyée (selon ton binding)
// }
