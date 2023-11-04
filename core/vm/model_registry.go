package vm

import (
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/ethclient"
)

type ModelRegistry struct {
    contract *bind.BoundContract 
}

func NewModelRegistry(contractAddr common.Address, client *ethclient.Client) (*ModelRegistry, error) {
    contract, err := bindModelRegistry(contractAddr, client)
    if err != nil {
        return nil, err
    }
    return &ModelRegistry{contract: contract}, nil
}

func (amr *ModelRegistry) RegisterModel(opts *bind.TransactOpts, modelId *big.Int, modelHash [32]byte) error {
    _, err := amr.contract.Transact(opts, "registerModel", modelId, modelHash)
    return err
}

func (amr *ModelRegistry) UpdateModel(opts *bind.TransactOpts, modelId *big.Int, newModelHash [32]byte) error {
    _, err := amr.contract.Transact(opts, "updateModel", modelId, newModelHash)
    return err
}

func (amr *ModelRegistry) GetModelHash(opts *bind.CallOpts, modelId *big.Int) ([32]byte, error) {
    var modelHash [32]byte
    _, err := amr.contract.Call(opts, &modelHash, "getModelHash", modelId)
    return modelHash, err
}
