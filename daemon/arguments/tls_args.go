package arguments

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/kurtosis-tech/stacktrace"
	"google.golang.org/grpc/credentials"
	"os"
)

func BuildTlsCredentials(ca []byte, cert []byte, key []byte) (credentials.TransportCredentials, error) {
	certPool, keyPair, err := buildTlsRootCaAndKeyPair(ca, cert, key)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Error building root CAs and key pair")
	}
	// nolint: exhaustruct
	tlsCredentialsConfig := &tls.Config{
		Certificates: []tls.Certificate{
			*keyPair,
		},
		RootCAs: certPool,
	}
	return credentials.NewTLS(tlsCredentialsConfig), nil
}

func BuildTlsRootCaAndKeyPairFromFiles(caFilePath string, certFilePath string, keyFilePath string) (*x509.CertPool, *tls.Certificate, error) {
	if certFilePath == "" || keyFilePath == "" {
		return nil, nil, stacktrace.NewError("Cannot build TLS credentials with empty certificate of key file path")
	}
	cert, err := os.ReadFile(certFilePath)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, "Unable to read content of certificate file at '%s'", certFilePath)
	}
	key, err := os.ReadFile(keyFilePath)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, "Unable to read content of key file at '%s'", keyFilePath)
	}

	var ca []byte
	if caFilePath != "" {
		ca, err = os.ReadFile(caFilePath)
		if err != nil {
			return nil, nil, stacktrace.Propagate(err, "Unable to read content of CA file at '%s'", caFilePath)
		}
	}

	certPool, keyPair, err := buildTlsRootCaAndKeyPair(ca, cert, key)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, "Error building Root CA and key pair from CA file '%s', cert "+
			"file '%s' and key file '%s'", caFilePath, certFilePath, keyFilePath)
	}
	return certPool, keyPair, nil
}

func buildTlsRootCaAndKeyPair(ca []byte, cert []byte, key []byte) (*x509.CertPool, *tls.Certificate, error) {
	var certPool *x509.CertPool
	if ca != nil {
		certPool = x509.NewCertPool()
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			return nil, nil, stacktrace.NewError("Unable to add the content of the CA to the certificate pool")
		}
	}

	serverCert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return nil, nil, stacktrace.Propagate(err, "Unable to build a key pair from the certificate and key "+
			"files provided")
	}
	return certPool, &serverCert, nil
}
