package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractBlockReference(t *testing.T) {
	tests := []struct {
		name     string
		request  *JsonRpcRequest
		expected string
		expInt   int64
		expErr   bool
	}{
		{
			name:     "nil request",
			request:  nil,
			expected: "",
			expInt:   0,
			expErr:   true,
		},
		{
			name: "eth_getBlockByNumber",
			request: &JsonRpcRequest{
				Method: "eth_getBlockByNumber",
				Params: []interface{}{"0xc5043f", false},
			},
			expected: "12911679",
			expInt:   12911679,
			expErr:   false,
		},
		{
			name: "invalid hex in eth_getBlockByNumber",
			request: &JsonRpcRequest{
				Method: "eth_getBlockByNumber",
				Params: []interface{}{"invalidHex"},
			},
			expected: "",
			expInt:   0,
			expErr:   false,
		},
		{
			name: "eth_getUncleByBlockNumberAndIndex",
			request: &JsonRpcRequest{
				Method: "eth_getUncleByBlockNumberAndIndex",
				Params: []interface{}{"0x1b4"},
			},
			expected: "436",
			expInt:   436,
			expErr:   false,
		},
		{
			name: "eth_getTransactionByBlockNumberAndIndex",
			request: &JsonRpcRequest{
				Method: "eth_getTransactionByBlockNumberAndIndex",
				Params: []interface{}{"0xc5043f", "0x0"},
			},
			expected: "12911679",
			expInt:   12911679,
			expErr:   false,
		},
		{
			name: "eth_getUncleCountByBlockNumber",
			request: &JsonRpcRequest{
				Method: "eth_getUncleCountByBlockNumber",
				Params: []interface{}{"0xc5043f"},
			},
			expected: "12911679",
			expInt:   12911679,
			expErr:   false,
		},
		{
			name: "eth_getBlockTransactionCountByNumber",
			request: &JsonRpcRequest{
				Method: "eth_getBlockTransactionCountByNumber",
				Params: []interface{}{"0xc5043f"},
			},
			expected: "12911679",
			expInt:   12911679,
			expErr:   false,
		},
		{
			name: "eth_getBlockReceipts",
			request: &JsonRpcRequest{
				Method: "eth_getBlockReceipts",
				Params: []interface{}{"0xc5043f"},
			},
			expected: "12911679",
			expInt:   12911679,
			expErr:   false,
		},
		{
			name: "eth_getLogs",
			request: &JsonRpcRequest{
				Method: "eth_getLogs",
				Params: []interface{}{
					map[string]interface{}{
						"fromBlock": "0x1b4",
						"toBlock":   "0x1b5",
					},
				},
			},
			expected: "0x1b4-0x1b5",
			expInt:   437,
			expErr:   false,
		},
		{
			name: "missing parameters in eth_getLogs",
			request: &JsonRpcRequest{
				Method: "eth_getLogs",
				Params: []interface{}{},
			},
			expected: "",
			expInt:   0,
			expErr:   false,
		},
		{
			name: "eth_getBalance",
			request: &JsonRpcRequest{
				Method: "eth_getBalance",
				Params: []interface{}{"0xAddress", "0x1b4"},
			},
			expected: "436",
			expInt:   436,
			expErr:   false,
		},
		{
			name: "missing parameters in eth_getBalance",
			request: &JsonRpcRequest{
				Method: "eth_getBalance",
				Params: []interface{}{"0xAddress"},
			},
			expected: "",
			expInt:   0,
			expErr:   true,
		},
		{
			name: "eth_getCode",
			request: &JsonRpcRequest{
				Method: "eth_getCode",
				Params: []interface{}{"0xAddress", "0x1b4"},
			},
			expected: "436",
			expInt:   436,
			expErr:   false,
		},
		{
			name: "eth_getTransactionCount",
			request: &JsonRpcRequest{
				Method: "eth_getTransactionCount",
				Params: []interface{}{"0xAddress", "0x1b4"},
			},
			expected: "436",
			expInt:   436,
			expErr:   false,
		},
		{
			name: "eth_call",
			request: &JsonRpcRequest{
				Method: "eth_call",
				Params: []interface{}{
					map[string]interface{}{
						"from": nil,
						"to":   "0x6b175474e89094c44da98b954eedeac495271d0f",
						"data": "0x70a082310000000000000000000000006E0d01A76C3Cf4288372a29124A26D4353EE51BE",
					},
					"0x1b4",
					map[string]interface{}{
						"0x1111111111111111111111111111111111111111": map[string]interface{}{
							"balance": "0xFFFFFFFFFFFFFFFFFFFF",
						},
					},
				},
			},
			expected: "436",
			expInt:   436,
			expErr:   false,
		},
		{
			name: "eth_feeHistory",
			request: &JsonRpcRequest{
				Method: "eth_feeHistory",
				Params: []interface{}{"0x8D97689C9818892B700e27F316cc3E41e17fBeb9", "0x1b4"},
			},
			expected: "436",
			expInt:   436,
			expErr:   false,
		},
		{
			name: "eth_getAccount",
			request: &JsonRpcRequest{
				Method: "eth_getAccount",
				Params: []interface{}{4, "0x1b4", []interface{}{25, 75}},
			},
			expected: "436",
			expInt:   436,
			expErr:   false,
		},
		{
			name: "eth_getBlockByHash",
			request: &JsonRpcRequest{
				Method: "eth_getBlockByHash",
				Params: []interface{}{"0x3f07a9c83155594c000642e7d60e8a8a00038d03e9849171a05ed0e2d47acbb3", false},
			},
			expected: "0x3f07a9c83155594c000642e7d60e8a8a00038d03e9849171a05ed0e2d47acbb3",
			expInt:   0,
			expErr:   false,
		},
		{
			name: "eth_getTransactionByBlockHashAndIndex",
			request: &JsonRpcRequest{
				Method: "eth_getTransactionByBlockHashAndIndex",
				Params: []interface{}{"0x829df9bb801fc0494abf2f443423a49ffa32964554db71b098d332d87b70a48b", "0x0"},
			},
			expected: "0x829df9bb801fc0494abf2f443423a49ffa32964554db71b098d332d87b70a48b",
			expInt:   0,
			expErr:   false,
		},
		{
			name: "eth_getBlockTransactionCountByHash",
			request: &JsonRpcRequest{
				Method: "eth_getBlockTransactionCountByHash",
				Params: []interface{}{"0x829df9bb801fc0494abf2f443423a49ffa32964554db71b098d332d87b70a48b"},
			},
			expected: "0x829df9bb801fc0494abf2f443423a49ffa32964554db71b098d332d87b70a48b",
			expInt:   0,
			expErr:   false,
		},
		{
			name: "eth_getUncleCountByBlockHash",
			request: &JsonRpcRequest{
				Method: "eth_getUncleCountByBlockHash",
				Params: []interface{}{"0x829df9bb801fc0494abf2f443423a49ffa32964554db71b098d332d87b70a48b"},
			},
			expected: "0x829df9bb801fc0494abf2f443423a49ffa32964554db71b098d332d87b70a48b",
			expInt:   0,
			expErr:   false,
		},
		{
			name: "eth_getProof",
			request: &JsonRpcRequest{
				Method: "eth_getProof",
				Params: []interface{}{
					"0x7F0d15C7FAae65896648C8273B6d7E43f58Fa842",
					[]interface{}{"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"},
					"0x1b4",
				},
			},
			expected: "436",
			expInt:   436,
			expErr:   false,
		},
		{
			name: "eth_getStorageAt",
			request: &JsonRpcRequest{
				Method: "eth_getStorageAt",
				Params: []interface{}{
					"0xE592427A0AEce92De3Edee1F18E0157C05861564",
					"0x0",
					"0x1b4",
				},
			},
			expected: "436",
			expInt:   436,
			expErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, resultUint, err := ExtractEvmBlockReferenceFromRequest(tt.request)
			if tt.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.expInt, resultUint)
		})
	}
}
