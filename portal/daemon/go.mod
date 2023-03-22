module github.com/kurtosis-tech/kurtosis-cloud/portal/daemon

go 1.18

replace github.com/kurtosis-tech/kurtosis-cloud/portal/api/golang => ../api/golang

require (
	github.com/google/uuid v1.3.0
	github.com/jpillora/chisel v1.8.1
	github.com/kurtosis-tech/kurtosis-cloud/portal/api/golang v0.0.0-00010101000000-000000000000
	github.com/kurtosis-tech/kurtosis/contexts-config-store v0.0.0-20230321133325-f4034e562ece
	github.com/kurtosis-tech/minimal-grpc-server/golang v0.0.0-20230317105020-7ca453c242bd
	github.com/kurtosis-tech/stacktrace v0.0.0-20211028211901-1c67a77b5409
	github.com/sirupsen/logrus v1.9.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.38.0
)

require (
	github.com/adrg/xdg v0.4.0 // indirect
	github.com/andrew-d/go-termutil v0.0.0-20150726205930-009166a695a2 // indirect
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/jpillora/ansi v1.0.2 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/jpillora/requestlog v1.0.0 // indirect
	github.com/jpillora/sizestr v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.29.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
