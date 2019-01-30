package testutil

import (
	"time"

	expect "github.com/Netflix/go-expect"
	"github.com/hinshun/vt10x"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/pkg/cliutil"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
)

func Stdio(c *expect.Console) terminal.Stdio {
	return terminal.Stdio{c.Tty(), c.Tty(), c.Tty()}
}

func ExpectInteractive(userinput func(*Console), testcli func()) {
	c, state, err := vt10x.NewVT10XConsole()
	Expect(err).NotTo(HaveOccurred())
	defer c.Close()
	cliutil.UseStdio(Stdio(c))
	// Dump the terminal's screen.
	defer func() { GinkgoWriter.Write([]byte(expect.StripTrailingEmptyLines(state.String()))) }()

	donec := make(chan struct{})
	go func() {
		defer GinkgoRecover()
		defer close(donec)

		userinput(&Console{console: c})
	}()

	//	time.Sleep(time.Hour)
	go func() {
		defer GinkgoRecover()

		testcli()

		// Close the slave end of the pty, and read the remaining bytes from the master end.
		c.Tty().Close()
		<-donec
	}()

	select {
	case <-time.After(10 * time.Second):
		c.Tty().Close()
		Fail("test timed out")
	case <-donec:
	}
}

type Console struct {
	console *expect.Console
}

func (c *Console) ExpectString(s string) string {
	ret, err := c.console.ExpectString(s)
	Expect(err).NotTo(HaveOccurred())
	return ret
}

func (c *Console) PressDown() {
	// These codes are covered here: https://en.wikipedia.org/wiki/ANSI_escape_code
	// see "Escape sequences" and "CSI sequences"
	// 27 = Escape
	// Alternatively, you can use the values written here: gopkg.in/AlecAivazis/survey.v1/terminal/sequences.go
	// But I used the CSI as I seems to be more standard

	_, err := c.console.Write([]byte{27,'[','B'})
	Expect(err).NotTo(HaveOccurred())
}

func (c *Console) Esc() {
	// I grabbed this value from here: gopkg.in/AlecAivazis/survey.v1/terminal/sequences.go
	// Originally I tried to use escape codes (https://en.wikipedia.org/wiki/ANSI_escape_code)
	// but it didnt work
	_, err := c.console.Write([]byte{27})
	Expect(err).NotTo(HaveOccurred())
}

func (c *Console) SendLine(s string) int {
	ret, err := c.console.SendLine(s)
	Expect(err).NotTo(HaveOccurred())
	return ret
}

func (c *Console) ExpectEOF() string {
	ret, err := c.console.ExpectEOF()
	Expect(err).NotTo(HaveOccurred())
	return ret
}
