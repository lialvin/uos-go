package system

import (
	uos "github.com/lialvin/uos-go"
)

// NewUnregProducer returns a `unregprod` action that lives on the
// `uosio.system` contract.
func NewUnregProducer(producer uos.AccountName) *uos.Action {
	return &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("unregprod"),
		Authorization: []uos.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(UnregProducer{
			Producer: producer,
		}),
	}
}

// UnregProducer represents the `uosio.system::unregprod` action
type UnregProducer struct {
	Producer uos.AccountName `json:"producer"`
}
