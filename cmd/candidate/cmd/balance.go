// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
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
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/smblab/candidate/pkg/balance/pb"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance <account> <from> <to>",
	Short: "Get balance for account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
			os.Exit(1)
		}

		conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		client := pb.NewBalanceServiceClient(conn)
		ctx := context.Background()

		res, err := client.GetBalance(ctx, &pb.GetBalanceRequest{})
		if err != nil {
			panic(err)
		}

		log.Printf("Balance: %s\n", res.String())
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)
}
