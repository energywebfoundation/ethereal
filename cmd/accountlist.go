// Copyright © 2017-2019 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens/v2"
	string2eth "github.com/wealdtech/go-string2eth"
)

// accountListCmd represents the account list command
var accountListCmd = &cobra.Command{
	Use:   "list",
	Short: "List visible accounts",
	Long: `List accounts that are visible to Ethereal.  For example:

    ethereal account list

In quiet mode this will return 0 if any accounts are found, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		wallets, err := cli.ObtainWallets(chainID)
		foundAccounts := false
		if err == nil {
			for _, wallet := range wallets {
				for _, account := range wallet.Accounts() {
					foundAccounts = true
					if !quiet {
						if !verbose {
							fmt.Println(account.Address.Hex())
						} else {
							fmt.Printf("Location:\t%s\n", account.URL)
							fmt.Printf("Address:\t%s\n", account.Address.Hex())
							if !offline {
								name, err := ens.ReverseResolve(client, account.Address)
								if err == nil {
									fmt.Printf("Name:\t\t%s\n", name)
								}
								ctx, cancel := localContext()
								defer cancel()
								balance, err := client.BalanceAt(ctx, account.Address, nil)
								if err == nil {
									fmt.Printf("Balance:\t%s\n", string2eth.WeiToString(balance, true))
								}
								nonce, err := client.PendingNonceAt(ctx, account.Address)
								if err == nil {
									fmt.Printf("Next nonce:\t%v\n", nonce)
								}
							}
							fmt.Println("")
						}
					}
				}
			}
		}

		if quiet {
			if foundAccounts {
				os.Exit(_exit_success)
			} else {
				os.Exit(_exit_failure)
			}
		}
	},
}

func init() {
	accountCmd.AddCommand(accountListCmd)
}
