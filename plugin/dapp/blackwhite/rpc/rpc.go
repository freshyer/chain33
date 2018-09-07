package rpc

import (
	"context"
	"math/rand"

	"gitlab.33.cn/chain33/chain33/client"
	"gitlab.33.cn/chain33/chain33/common/address"
	bw "gitlab.33.cn/chain33/chain33/plugin/dapp/blackwhite/types"
	"gitlab.33.cn/chain33/chain33/queue"
	"gitlab.33.cn/chain33/chain33/types"
)

type channelClient struct {
	client.QueueProtocolAPI
}

func (c *channelClient) Init(q queue.Client) {
	c.QueueProtocolAPI, _ = client.New(q, nil)
}

func (c *channelClient) Create(ctx context.Context, head *bw.BlackwhiteCreate) (*types.UnsignTx, error) {
	val := &bw.BlackwhiteAction{
		Ty:    bw.BlackwhiteActionCreate,
		Value: &bw.BlackwhiteAction_Create{head},
	}
	tx := &types.Transaction{
		Execer:  bw.ExecerBlackwhite,
		Payload: types.Encode(val),
		Fee:     0,
		Nonce:   rand.New(rand.NewSource(types.Now().UnixNano())).Int63(),
		To:      address.ExecAddress(string(bw.ExecerBlackwhite)),
	}
	tx.SetRealFee(types.MinFee)
	data := types.Encode(tx)
	return &types.UnsignTx{Data: data}, nil
}

func (c *channelClient) Show(ctx context.Context, head *bw.BlackwhiteShow) (*types.UnsignTx, error) {
	val := &bw.BlackwhiteAction{
		Ty:    bw.BlackwhiteActionShow,
		Value: &bw.BlackwhiteAction_Show{head},
	}
	tx := &types.Transaction{
		Execer:  bw.ExecerBlackwhite,
		Payload: types.Encode(val),
		Fee:     0,
		Nonce:   rand.New(rand.NewSource(types.Now().UnixNano())).Int63(),
		To:      address.ExecAddress(string(bw.ExecerBlackwhite)),
	}
	tx.SetRealFee(types.MinFee)
	data := types.Encode(tx)
	return &types.UnsignTx{Data: data}, nil
}

func (c *channelClient) Play(ctx context.Context, head *bw.BlackwhitePlay) (*types.UnsignTx, error) {
	val := &bw.BlackwhiteAction{
		Ty:    bw.BlackwhiteActionPlay,
		Value: &bw.BlackwhiteAction_Play{head},
	}
	tx := &types.Transaction{
		Execer:  bw.ExecerBlackwhite,
		Payload: types.Encode(val),
		Fee:     0,
		Nonce:   rand.New(rand.NewSource(types.Now().UnixNano())).Int63(),
		To:      address.ExecAddress(string(bw.ExecerBlackwhite)),
	}
	tx.SetRealFee(types.MinFee)
	data := types.Encode(tx)
	return &types.UnsignTx{Data: data}, nil
}

func (c *channelClient) TimeoutDone(ctx context.Context, head *bw.BlackwhiteTimeoutDone) (*types.UnsignTx, error) {
	val := &bw.BlackwhiteAction{
		Ty:    bw.BlackwhiteActionTimeoutDone,
		Value: &bw.BlackwhiteAction_TimeoutDone{head},
	}
	tx := &types.Transaction{
		Execer:  bw.ExecerBlackwhite,
		Payload: types.Encode(val),
		Fee:     0,
		Nonce:   rand.New(rand.NewSource(types.Now().UnixNano())).Int63(),
		To:      address.ExecAddress(string(bw.ExecerBlackwhite)),
	}
	tx.SetRealFee(types.MinFee)
	data := types.Encode(tx)
	return &types.UnsignTx{Data: data}, nil
}