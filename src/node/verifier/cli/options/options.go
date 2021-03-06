package options

import "github.com/c-bata/go-prompt"

var Initial = map[string]string{
	"eth":    "ETH functions",
	"plasma": "ETH functions",
	"quit":   "Quit from app",
}
var Eth = map[string]string{
	"transfer":     "example transfer 4 wei:\neth transfer 4 0x4ED6d26c6885247fA22746AB2c5328076597a5DF",
	"balance":      "example get balance of account 0x4ED6d26c6885247fA22746AB2c5328076597a5DF:\neth balance 0x4ED6d26c6885247fA22746AB2c5328076597a5DF",
	"ownerBalance": "example get balance of user(config) eth account:\neth ownerBalance",
}

var Plasma = map[string]string{
	"deposit":  "example deposit 1 wei:\nplasma deposit 1",
	"balance":  "example get balance of plasma:\nplasma balance",
	"transfer": "example transfer:\nplasma tr block txN out value address",
	"utxo":     "example get list UTXOs from plasma:\nplasma utxo",
	"exit":     "example exit from plasma:\nplasma exit",
}

var InitialOptions = []prompt.Suggest{
	{Text: "eth", Description: Initial["eth"]},
	{Text: "plasma", Description: Initial["plasma"]},
	{Text: "quit", Description: Initial["quit"]},
}

var EthOptions = []prompt.Suggest{
	{Text: "transfer", Description: Eth["transfer"]},
	{Text: "balance", Description: Eth["balance"]},
	{Text: "ownerBalance", Description: Eth["ownerBalance"]},
}

var PlasmaOptions = []prompt.Suggest{
	{Text: "deposit", Description: Plasma["deposit"]},
	{Text: "balance", Description: Plasma["balance"]},
	{Text: "transfer", Description: Plasma["transfer"]},
	{Text: "utxo", Description: Plasma["utxo"]},
	{Text: "exit", Description: Plasma["exit"]},
	// {Text: "withdraw", Description: "ex: Withdraw eth from plasma"},
}
