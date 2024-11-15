package helper

const (
	defaultHttpEchoImage = "us-docker.pkg.dev/developers-369321/rav-test/kennship-http-echo:latest"
	HttpEchoName         = "http-echo"
	HttpEchoPort         = 3000
)

// Deprecated
// ported to test/kubernetes/e2e/defaults/testdata/http_echo.yaml
func NewEchoHttp(namespace string) (TestContainer, error) {
	return newTestContainer(namespace, defaultHttpEchoImage, HttpEchoName, HttpEchoPort, true, nil)
}

const (
	defaultTcpEchoImage = "soloio/tcp-echo:latest"
	TcpEchoName         = "tcp-echo"
	TcpEchoPort         = 1025
)

// Deprecated
// ported to test/kubernetes/e2e/defaults/testdata/tcp_echo.yaml
func NewEchoTcp(namespace string) (TestContainer, error) {
	return newTestContainer(namespace, defaultTcpEchoImage, TcpEchoName, TcpEchoPort, true, nil)
}
