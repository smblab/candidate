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
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/smblab/candidate/pkg/transactions"
	"github.com/smblab/candidate/pkg/transactions/pb"
)

// transactionserviceCmd represents the transactionservice command
var transactionserviceCmd = &cobra.Command{
	Use:   "transactionservice",
	Short: "Launch transaction service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		repository, err := transactions.NewRepository("db.sqlite")
		if err != nil {
			panic(err)
		}

		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			panic(err)
		}

		service := transactions.NewService(repository)
		server := grpc.NewServer()
		pb.RegisterTransactionServiceServer(server, service)

		log.Println("Transaction service listening on :50052")
		if err := server.Serve(lis); err != nil {
			panic(err)
		}
	},
}

func init() {
	serveCmd.AddCommand(transactionserviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transactionserviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transactionserviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
