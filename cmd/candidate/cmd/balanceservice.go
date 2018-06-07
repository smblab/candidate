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

	"github.com/smblab/candidate/pkg/balance"
	"github.com/smblab/candidate/pkg/balance/pb"
	tpb "github.com/smblab/candidate/pkg/transactions/pb"
)

// balanceserviceCmd represents the balanceservice command
var balanceserviceCmd = &cobra.Command{
	Use:   "balanceservice",
	Short: "Launch balance service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		client := tpb.NewTransactionServiceClient(conn)

		lis, err := net.Listen("tcp", ":50053")
		if err != nil {
			panic(err)
		}

		service := balance.NewService(client)
		server := grpc.NewServer()
		pb.RegisterBalanceServiceServer(server, service)

		log.Println("Balance service listening on :50053")
		if err := server.Serve(lis); err != nil {
			panic(err)
		}
	},
}

func init() {
	serveCmd.AddCommand(balanceserviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// balanceserviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// balanceserviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
