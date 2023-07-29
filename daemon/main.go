package main

import (
	"context"
	"fmt"
	portal_constructors "github.com/kurtosis-tech/kurtosis-portal/api/golang/constructors"
	portal_api "github.com/kurtosis-tech/kurtosis-portal/api/golang/generated"
	"github.com/kurtosis-tech/kurtosis-portal/daemon/arguments"
	"github.com/kurtosis-tech/kurtosis-portal/daemon/client"
	"github.com/kurtosis-tech/kurtosis-portal/daemon/server"
	minimal_grpc_server "github.com/kurtosis-tech/minimal-grpc-server/golang/server"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	successExitCode = 0
	failureExitCode = 1

	grpcServerStopGracePeriod = 5 * time.Second

	forceColors   = true
	fullTimestamp = true

	logMethodAlongWithLogLine = true
	functionPathSeparator     = "."
	emptyFunctionName         = ""
)

func main() {
	ctx := context.Background()
	logrus.SetLevel(logrus.DebugLevel)
	// This allows the filename & function to be reported
	logrus.SetReportCaller(logMethodAlongWithLogLine)
	// NOTE: we'll want to change the ForceColors to false if we ever want structured logging
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:               forceColors,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             fullTimestamp,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			fullFunctionPath := strings.Split(f.Function, functionPathSeparator)
			functionName := fullFunctionPath[len(fullFunctionPath)-1]
			_, filename := path.Split(f.File)
			return emptyFunctionName, formatFilenameFunctionForLogs(filename, functionName)
		},
	})

	mainArgs, err := arguments.Parse()
	if err != nil {
		fmt.Fprintln(logrus.StandardLogger().Out, err)
		os.Exit(failureExitCode)
	}

	// TODO(gb): add flags to run only server or only client. On Kurtosis cloud server, we technically won't need to run
	//  the client (only server). That being said, client is very lightweight and probably not worth the effort right
	//  at this point
	errorGroup, cancellableCtx := errgroup.WithContext(ctx)
	errorGroup.Go(func() error {
		return runServer(cancellableCtx, mainArgs)
	})
	errorGroup.Go(func() error {
		return runClient(cancellableCtx, mainArgs)
	})

	if err = errorGroup.Wait(); err != nil {
		logrus.Errorf("An error occurred when running the portal")
		fmt.Fprintln(logrus.StandardLogger().Out, err)
		os.Exit(failureExitCode)
	}
	os.Exit(successExitCode)
}

func runClient(ctx context.Context, mainArgs *arguments.PortalArgs) error {
	if mainArgs.ServerOnly {
		logrus.Infof("Running portal in server-only mode. Not starting client.")
		return nil
	}

	kurtosisPortalClient := client.NewKurtosisClient()
	defer kurtosisPortalClient.Close()

	// Whatever current context should be used, switch to it
	if _, err := kurtosisPortalClient.SwitchContext(ctx, portal_constructors.NewSwitchContextArgs()); err != nil {
		return stacktrace.Propagate(err, "Unable to apply current context configuration, daemon cannot start")
	}

	kurtosisPortalDaemonRegistrationFunc := func(grpcServer *grpc.Server) {
		portal_api.RegisterKurtosisPortalClientServer(grpcServer, kurtosisPortalClient)
	}
	kurtosisPortalClientDaemon := minimal_grpc_server.NewMinimalGRPCServer(
		client.PortalClientGrpcPort,
		grpcServerStopGracePeriod,
		[]func(*grpc.Server){
			kurtosisPortalDaemonRegistrationFunc,
		},
	)

	logrus.Infof("Kurtosis Portal Client running and listening on port %d", client.PortalClientGrpcPort)
	if err := kurtosisPortalClientDaemon.RunUntilStopped(ctx.Done()); err != nil {
		return stacktrace.Propagate(err, "An error occurred running the Kurtosis Portal Daemon")
	}
	return nil
}

func runServer(ctx context.Context, mainArgs *arguments.PortalArgs) error {
	kurtosisPortalServer := server.NewKurtosisPortalServer(mainArgs.TlsCaFilePath, mainArgs.TlsServerKeyFilePath, mainArgs.TlsServerCertFilePath, mainArgs.RemoteHost)
	defer kurtosisPortalServer.Close()

	err := kurtosisPortalServer.StartTunnelServer(ctx, server.PortalServerTunnelListeningHost, server.PortalServerTunnelPort)
	if err != nil {
		return stacktrace.Propagate(err, "Unable to start server-side tunnel")
	}

	kurtosisPortalDaemonRegistrationFunc := func(grpcServer *grpc.Server) {
		portal_api.RegisterKurtosisPortalServerServer(grpcServer, kurtosisPortalServer)
	}

	var kurtosisPortalServerDaemon *minimal_grpc_server.MinimalGRPCServer
	if mainArgs.TlsServerKeyFilePath == "" && mainArgs.TlsServerCertFilePath == "" {
		kurtosisPortalServerDaemon = minimal_grpc_server.NewMinimalGRPCServer(
			server.PortalServerGrpcPort,
			grpcServerStopGracePeriod,
			[]func(*grpc.Server){
				kurtosisPortalDaemonRegistrationFunc,
			},
		)
	} else {
		caCertPool, serverCert, err := arguments.BuildTlsRootCaAndKeyPairFromFiles(mainArgs.TlsCaFilePath, mainArgs.TlsServerCertFilePath, mainArgs.TlsServerKeyFilePath)
		if err != nil {
			return stacktrace.Propagate(err, "Unable to build TLS configurations from the provided files.")
		}
		kurtosisPortalServerDaemon = minimal_grpc_server.NewMinimalHttpsGRPCServer(
			server.PortalServerGrpcPort,
			grpcServerStopGracePeriod,
			caCertPool,
			serverCert,
			[]func(*grpc.Server){
				kurtosisPortalDaemonRegistrationFunc,
			},
		)
	}

	logrus.Infof("Kurtosis Portal Server running and listening on port %d", server.PortalServerGrpcPort)
	if err := kurtosisPortalServerDaemon.RunUntilStopped(ctx.Done()); err != nil {
		return stacktrace.Propagate(err, "An error occurred running the Kurtosis Portal Daemon")
	}
	return nil
}

func formatFilenameFunctionForLogs(filename string, functionName string) string {
	var output strings.Builder
	output.WriteString("[")
	output.WriteString(filename)
	output.WriteString(":")
	output.WriteString(functionName)
	output.WriteString("]")
	return output.String()
}
