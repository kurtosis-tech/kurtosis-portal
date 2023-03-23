package arguments

import (
	"flag"
	"github.com/kurtosis-tech/stacktrace"
	"os"
)

type PortalArgs struct {
	ServerOnly bool

	TlsCaFilePath string

	TlsServerKeyFilePath  string
	TlsServerCertFilePath string
}

func Parse() (*PortalArgs, error) {
	args := new(PortalArgs)
	flag.BoolVar(&args.ServerOnly, "server-only", false, "Optional: runs in server-only mode, no client will be started. Defaults to false.")
	flag.StringVar(&args.TlsCaFilePath, "tls-ca", "", "Optional: Path to the Certificate Authority file to be used by the server for TLS")
	flag.StringVar(&args.TlsServerKeyFilePath, "tls-server-key", "", "Optional: Path to the certificate key file to be used by the server for TLS")
	flag.StringVar(&args.TlsServerCertFilePath, "tls-server-cert", "", "Optional: Path to the Certificate file to be used by the server for TLS")
	flag.Parse()

	if err := validate(args); err != nil {
		return nil, stacktrace.Propagate(err, "Invalid arguments")
	}
	return args, nil
}

func validate(args *PortalArgs) error {
	if args.TlsServerCertFilePath != "" && args.TlsServerKeyFilePath == "" {
		return stacktrace.NewError("Providing a certificate file without its associated key is invalid. Please provide both or none.")
	} else if args.TlsServerCertFilePath == "" && args.TlsServerKeyFilePath != "" {
		return stacktrace.NewError("Providing a key file without its associated certificate is invalid. Please provide both or none.")
	}
	if args.TlsServerCertFilePath != "" {
		if err := validateFileExists(args.TlsServerCertFilePath); err != nil {
			return stacktrace.Propagate(err, "Cannot validate certificate file '%s' exists", args.TlsServerCertFilePath)
		}
	}
	if args.TlsServerKeyFilePath != "" {
		if err := validateFileExists(args.TlsServerKeyFilePath); err != nil {
			return stacktrace.Propagate(err, "Cannot validate key file '%s' exists", args.TlsServerKeyFilePath)
		}
	}
	if args.TlsCaFilePath != "" {
		if err := validateFileExists(args.TlsCaFilePath); err != nil {
			return stacktrace.Propagate(err, "Cannot Certificate Authority file '%s' exists", args.TlsCaFilePath)
		}
	}
	return nil
}

func validateFileExists(filePath string) error {
	if _, err := os.Stat(filePath); err != nil {
		return stacktrace.Propagate(err, "File '%s' does not exist", filePath)
	}
	return nil
}
