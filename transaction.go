package main

import (
)

type Transaction struct {
  ID []byte
  Vin []TXInput
  Vout []TXOutput
}

func NewCoinbaseTx(to, data string) *Transaction {
  if data == "" {
    logf("Reward to '%s'", to)
  }

  txin := TXInput{[]byte{}, -1 ,data}
  txout := TXOutput{50, to}
  tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}

  return &tx
}
