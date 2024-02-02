package main

import (
	"context"
	gw "github.com/d1mpi/grpc-tron/api"
	"github.com/d1mpi/grpc-tron/core"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func InitGrpcClient() gw.WalletClient {
	grpcEndpoint := viper.GetString("server.host") + ":" + viper.GetString("server.port")

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(grpcEndpoint, opts...)

	if err != nil {
		panic(err)
	}

	return gw.NewWalletClient(conn)
}

func (s *WalletServer) GetContract(ctx context.Context, in *gw.BytesMessage) (*core.SmartContract, error) {
	contract, err := s.client.GetContract(ctx, in)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func (s *WalletServer) TriggerConstantContract(ctx context.Context, in *core.TriggerSmartContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.TriggerConstantContract(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) GetAccount(ctx context.Context, in *core.Account) (*core.Account, error) {
	account, err := s.client.GetAccount(ctx, in)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *WalletServer) GetAccountById(ctx context.Context, in *core.Account) (*core.Account, error) {
	account, err := s.client.GetAccountById(ctx, in)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *WalletServer) CreateTransaction(ctx context.Context, in *core.TransferContract) (*core.Transaction, error) {
	transaction, err := s.client.CreateTransaction(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) CreateTransaction2(ctx context.Context, in *core.TransferContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.CreateTransaction2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) BroadcastTransaction(ctx context.Context, in *core.Transaction) (*gw.Return, error) {
	ret, err := s.client.BroadcastTransaction(ctx, in)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *WalletServer) UpdateAccount(ctx context.Context, in *core.AccountUpdateContract) (*core.Transaction, error) {
	transaction, err := s.client.UpdateAccount(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) SetAccountId(ctx context.Context, in *core.SetAccountIdContract) (*core.Transaction, error) {
	transaction, err := s.client.SetAccountId(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) UpdateAccount2(ctx context.Context, in *core.AccountUpdateContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.UpdateAccount2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) VoteWitnessAccount(ctx context.Context, in *core.VoteWitnessContract) (*core.Transaction, error) {
	transaction, err := s.client.VoteWitnessAccount(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) UpdateSetting(ctx context.Context, in *core.UpdateSettingContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.UpdateSetting(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) UpdateEnergyLimit(ctx context.Context, in *core.UpdateEnergyLimitContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.UpdateEnergyLimit(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) VoteWitnessAccount2(ctx context.Context, in *core.VoteWitnessContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.VoteWitnessAccount2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) CreateAssetIssue(ctx context.Context, in *core.AssetIssueContract) (*core.Transaction, error) {
	transaction, err := s.client.CreateAssetIssue(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) CreateAssetIssue2(ctx context.Context, in *core.AssetIssueContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.CreateAssetIssue2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) UpdateWitness(ctx context.Context, in *core.WitnessUpdateContract) (*core.Transaction, error) {
	transaction, err := s.client.UpdateWitness(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) UpdateWitness2(ctx context.Context, in *core.WitnessUpdateContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.UpdateWitness2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) CreateAccount(ctx context.Context, in *core.AccountCreateContract) (*core.Transaction, error) {
	transaction, err := s.client.CreateAccount(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) CreateAccount2(ctx context.Context, in *core.AccountCreateContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.CreateAccount2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) CreateWitness(ctx context.Context, in *core.WitnessCreateContract) (*core.Transaction, error) {
	transaction, err := s.client.CreateWitness(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) CreateWitness2(ctx context.Context, in *core.WitnessCreateContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.CreateWitness2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) TransferAsset(ctx context.Context, in *core.TransferAssetContract) (*core.Transaction, error) {
	transaction, err := s.client.TransferAsset(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) TransferAsset2(ctx context.Context, in *core.TransferAssetContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.TransferAsset2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ParticipateAssetIssue(ctx context.Context, in *core.ParticipateAssetIssueContract) (*core.Transaction, error) {
	transaction, err := s.client.ParticipateAssetIssue(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) ParticipateAssetIssue2(ctx context.Context, in *core.ParticipateAssetIssueContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.ParticipateAssetIssue2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) FreezeBalance(ctx context.Context, in *core.FreezeBalanceContract) (*core.Transaction, error) {
	transaction, err := s.client.FreezeBalance(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) FreezeBalance2(ctx context.Context, in *core.FreezeBalanceContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.FreezeBalance2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) UnfreezeBalance(ctx context.Context, in *core.UnfreezeBalanceContract) (*core.Transaction, error) {
	transaction, err := s.client.UnfreezeBalance(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) UnfreezeBalance2(ctx context.Context, in *core.UnfreezeBalanceContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.UnfreezeBalance2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) UnfreezeAsset(ctx context.Context, in *core.UnfreezeAssetContract) (*core.Transaction, error) {
	transaction, err := s.client.UnfreezeAsset(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) UnfreezeAsset2(ctx context.Context, in *core.UnfreezeAssetContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.UnfreezeAsset2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) WithdrawBalance(ctx context.Context, in *core.WithdrawBalanceContract) (*core.Transaction, error) {
	transaction, err := s.client.WithdrawBalance(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) WithdrawBalance2(ctx context.Context, in *core.WithdrawBalanceContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.WithdrawBalance2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) UpdateAsset(ctx context.Context, in *core.UpdateAssetContract) (*core.Transaction, error) {
	transaction, err := s.client.UpdateAsset(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) UpdateAsset2(ctx context.Context, in *core.UpdateAssetContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.UpdateAsset2(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ProposalCreate(ctx context.Context, in *core.ProposalCreateContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.ProposalCreate(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ProposalApprove(ctx context.Context, in *core.ProposalApproveContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.ProposalApprove(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ProposalDelete(ctx context.Context, in *core.ProposalDeleteContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.ProposalDelete(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) BuyStorage(ctx context.Context, in *core.BuyStorageContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.BuyStorage(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) BuyStorageBytes(ctx context.Context, in *core.BuyStorageBytesContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.BuyStorageBytes(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) SellStorage(ctx context.Context, in *core.SellStorageContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.SellStorage(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ExchangeCreate(ctx context.Context, in *core.ExchangeCreateContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.ExchangeCreate(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ExchangeInject(ctx context.Context, in *core.ExchangeInjectContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.ExchangeInject(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ExchangeWithdraw(ctx context.Context, in *core.ExchangeWithdrawContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.ExchangeWithdraw(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ExchangeTransaction(ctx context.Context, in *core.ExchangeTransactionContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.ExchangeTransaction(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ListNodes(ctx context.Context, in *gw.EmptyMessage) (*gw.NodeList, error) {
	nodes, err := s.client.ListNodes(ctx, in)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func (s *WalletServer) GetAssetIssueByAccount(ctx context.Context, in *core.Account) (*gw.AssetIssueList, error) {
	assets, err := s.client.GetAssetIssueByAccount(ctx, in)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (s *WalletServer) GetAccountResource(ctx context.Context, in *core.Account) (*gw.AccountResourceMessage, error) {
	account, err := s.client.GetAccountResource(ctx, in)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *WalletServer) GetAssetIssueByName(ctx context.Context, in *gw.BytesMessage) (*core.AssetIssueContract, error) {
	asset, err := s.client.GetAssetIssueByName(ctx, in)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (s *WalletServer) GetAssetIssueListByName(ctx context.Context, in *gw.BytesMessage) (*gw.AssetIssueList, error) {
	asset, err := s.client.GetAssetIssueListByName(ctx, in)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (s *WalletServer) GetAssetIssueById(ctx context.Context, in *gw.BytesMessage) (*core.AssetIssueContract, error) {
	asset, err := s.client.GetAssetIssueById(ctx, in)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (s *WalletServer) GetNowBlock(ctx context.Context, in *gw.EmptyMessage) (*core.Block, error) {
	block, err := s.client.GetNowBlock(ctx, in)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (s *WalletServer) GetNowBlock2(ctx context.Context, in *gw.EmptyMessage) (*gw.BlockExtention, error) {
	block, err := s.client.GetNowBlock2(ctx, in)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (s *WalletServer) GetBlockByNum(ctx context.Context, in *gw.NumberMessage) (*core.Block, error) {
	block, err := s.client.GetBlockByNum(ctx, in)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (s *WalletServer) GetBlockByNum2(ctx context.Context, in *gw.NumberMessage) (*gw.BlockExtention, error) {
	blockExt, err := s.client.GetBlockByNum2(ctx, in)
	if err != nil {
		return nil, err
	}

	return blockExt, nil
}

func (s *WalletServer) GetTransactionCountByBlockNum(ctx context.Context, in *gw.NumberMessage) (*gw.NumberMessage, error) {
	numberMes, err := s.client.GetTransactionCountByBlockNum(ctx, in)
	if err != nil {
		return nil, err
	}

	return numberMes, nil
}

func (s *WalletServer) GetBlockById(ctx context.Context, in *gw.BytesMessage) (*core.Block, error) {
	block, err := s.client.GetBlockById(ctx, in)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (s *WalletServer) GetBlockByLimitNext(ctx context.Context, in *gw.BlockLimit) (*gw.BlockList, error) {
	block, err := s.client.GetBlockByLimitNext(ctx, in)
	if err != nil {
		log.Fatal("Error NewWalletClient connection")
		return nil, err
	}

	return block, nil
}

func (s *WalletServer) GetBlockByLimitNext2(ctx context.Context, in *gw.BlockLimit) (*gw.BlockListExtention, error) {
	blocksExt, err := s.client.GetBlockByLimitNext2(ctx, in)
	if err != nil {
		return nil, err
	}

	return blocksExt, nil
}

func (s *WalletServer) GetBlockByLatestNum(ctx context.Context, in *gw.NumberMessage) (*gw.BlockList, error) {
	blocks, err := s.client.GetBlockByLatestNum(ctx, in)
	if err != nil {
		return nil, err
	}

	return blocks, nil
}

func (s *WalletServer) GetBlockByLatestNum2(ctx context.Context, in *gw.NumberMessage) (*gw.BlockListExtention, error) {
	blocksExt, err := s.client.GetBlockByLatestNum2(ctx, in)
	if err != nil {
		return nil, err
	}

	return blocksExt, nil
}

func (s *WalletServer) GetTransactionById(ctx context.Context, in *gw.BytesMessage) (*core.Transaction, error) {
	transaction, err := s.client.GetTransactionById(ctx, in)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *WalletServer) DeployContract(ctx context.Context, in *core.CreateSmartContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.DeployContract(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) TriggerContract(ctx context.Context, in *core.TriggerSmartContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.TriggerContract(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ClearContractABI(ctx context.Context, in *core.ClearABIContract) (*gw.TransactionExtention, error) {
	transactionExt, err := s.client.ClearContractABI(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactionExt, nil
}

func (s *WalletServer) ListWitnesses(ctx context.Context, in *gw.EmptyMessage) (*gw.WitnessList, error) {
	witness, err := s.client.ListWitnesses(ctx, in)
	if err != nil {
		return nil, err
	}

	return witness, nil
}

func (s *WalletServer) EstimateEnergy(ctx context.Context, in *core.TriggerSmartContract) (*gw.EstimateEnergyMessage, error) {
	message, err := s.client.EstimateEnergy(ctx, in)
	if err != nil {
		return nil, err
	}

	return message, nil
}
