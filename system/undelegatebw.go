package system

import (
	uos "github.com/lialvin/uos-go"
)

// NewUndelegateBW returns a `undelegatebw` action that lives on the
// `uosio.system` contract.
func NewUndelegateBW(from, receiver uos.AccountName, unstakeCPU, unstakeNet uos.Asset) *uos.Action {
	return &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("undelegatebw"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(UndelegateBW{
			From:       from,
			Receiver:   receiver,
			UnstakeNet: unstakeNet,
			UnstakeCPU: unstakeCPU,
		}),
	}
}

// UndelegateBW represents the `uosio.system::undelegatebw` action.
type UndelegateBW struct {
	From       uos.AccountName `json:"from"`
	Receiver   uos.AccountName `json:"receiver"`
	UnstakeNet uos.Asset       `json:"unstake_net_quantity"`
	UnstakeCPU uos.Asset       `json:"unstake_cpu_quantity"`
}
