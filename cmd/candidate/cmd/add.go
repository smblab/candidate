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
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/smblab/candidate/pkg/transactions/pb"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <message> <amount>",
	Short: "Add transaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Usage()
			os.Exit(1)
		}

		conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		client := pb.NewTransactionServiceClient(conn)
		ctx := context.Background()

		uuid, err := uuid.NewUUID()
		if err != nil {
			panic(err)
		}

		amount, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			panic(err)
		}

		t, err := client.CreateTransaction(ctx, &pb.CreateTransactionRequest{
			RequestId: uuid.String(),
			Transaction: &pb.Transaction{
				Name:    fmt.Sprintf("accounts/1/transactions/%s", uuid.String()),
				Parent:  "accounts/1",
				Message: args[0],
				Amount: &pb.Money{
					Amount:       amount,
					CurrencyCode: "NOK",
				},
				TransactionTime: ptypes.TimestampNow(),
			},
		})
		if err != nil {
			panic(err)
		}

		log.Printf("Created transaction: %s\n", t.String())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
