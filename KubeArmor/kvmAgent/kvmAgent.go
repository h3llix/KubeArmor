// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeArmor

package kvmagent

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"os"
	"time"

	kg "github.com/kubearmor/KubeArmor/KubeArmor/log"
	tp "github.com/kubearmor/KubeArmor/KubeArmor/types"

	pb "github.com/kubearmor/KubeArmor/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// KVMAgent Structure
type KVMAgent struct {
	Identity         string
	gRPCServer       string
	gRPCConnection   *grpc.ClientConn
	gRPCClient       pb.KVMClient
	UpdateHostPolicy func(tp.K8sKubeArmorHostPolicyEvent)
}

func getgRPCAddress() (string, error) {
	serverAddr := net.JoinHostPort(os.Getenv("CLUSTER_IP"), os.Getenv("CLUSTER_PORT"))
	if serverAddr == ":" {
		return "", errors.New("host and port value is empty")
	}
	return serverAddr, nil
}

// NewKVMAgent Function
func NewKVMAgent(eventCb tp.KubeArmorHostPolicyEventCallback) *KVMAgent {
	kvm := &KVMAgent{}

	// Get identity
	kvm.Identity = os.Getenv("WORKLOAD_IDENTITY")

	// Get the address of gRPC server
	gRPCServer, err := getgRPCAddress()
	if err != nil {
		kg.Warn(err.Error())
		return nil
	}

	kvm.gRPCServer = gRPCServer

	// Connect to gRPC server
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	gRPCConnection, err := grpc.DialContext(ctx, kvm.gRPCServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		kg.Warn(err.Error())
		return nil
	}

	kvm.gRPCConnection = gRPCConnection
	kvm.gRPCClient = pb.NewKVMClient(gRPCConnection)

	// Register identity
	response, err := kvm.gRPCClient.RegisterAgentIdentity(context.Background(), &pb.AgentIdentity{Identity: kvm.Identity})
	if err != nil || response.Status != 0 {
		kg.Warn("failed to register KVM agent identity")
		return nil
	}

	// Link ParseAndUpdateHostSecurityPolicy()
	kvm.UpdateHostPolicy = eventCb

	return kvm
}

// DestroyKVMAgent Function
func (kvm *KVMAgent) DestroyKVMAgent() error {
	if err := kvm.gRPCConnection.Close(); err != nil {
		return err
	}
	return nil
}

// ConnectToKVMService Function
func (kvm *KVMAgent) ConnectToKVMService() {
	for {
		md := metadata.New(map[string]string{"identity": kvm.Identity})
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		stream, err := kvm.gRPCClient.SendPolicy(ctx)
		if err != nil {
			kg.Warn("Failed to connect stream")

			// close the connection
			if err = kvm.gRPCConnection.Close(); err != nil {
				kg.Warn("Failed to close the current connection")
			}

			// connect to gRPC server again
			ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
			gRPCConnection, err := grpc.DialContext(ctx, kvm.gRPCServer, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				kg.Warn("gRPC server is not accessable")
				return
			}

			// update Connection and Client
			kvm.gRPCConnection = gRPCConnection
			kvm.gRPCClient = pb.NewKVMClient(gRPCConnection)

			continue
		}

		for {
			status := int32(0)
			policyEvent := tp.K8sKubeArmorHostPolicyEvent{}

			policy, err := stream.Recv()
			if err == io.EOF {
				continue
			} else if err != nil {
				// close the connection
				if err = kvm.gRPCConnection.Close(); err != nil {
					kg.Warn("Failed to close the current connection")
				}

				// connect to gRPC server again
				ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
				gRPCConnection, err := grpc.DialContext(ctx, kvm.gRPCServer, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					kg.Warn("gRPC server is not accessable")
					return
				}

				// update Connection and Client
				kvm.gRPCConnection = gRPCConnection
				kvm.gRPCClient = pb.NewKVMClient(gRPCConnection)

				break
			}

			// get a policy
			err = json.Unmarshal(policy.PolicyData, &policyEvent)
			if err == nil {
				// update the policy
				kvm.UpdateHostPolicy(policyEvent)
			} else {
				kg.Warn("Failed to load a policy")
				status = 1
			}

			// return the status
			if err = stream.Send(&pb.Status{Status: status}); err != nil {
				kg.Warn(err.Error())
			}
		}
	}
}
