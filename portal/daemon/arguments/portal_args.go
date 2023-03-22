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
	flag.BoolVar(&args.ServerOnly, "server-only", false, "")
	flag.StringVar(&args.TlsCaFilePath, "tls-ca", "", "")
	flag.StringVar(&args.TlsServerKeyFilePath, "tls-server-key", "", "")
	flag.StringVar(&args.TlsServerCertFilePath, "tls-server-cert", "", "")
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
