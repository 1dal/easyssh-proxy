package easyssh

import (
	"os"
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKeyFile(t *testing.T) {
	// missing file
	_, err := getKeyFile("abc")
	assert.Error(t, err)
	assert.Equal(t, "open abc: no such file or directory", err.Error())

	// wrong format
	_, err = getKeyFile("../tests/.ssh/id_rsa.pub")
	assert.Error(t, err)
	assert.Equal(t, "ssh: no key found", err.Error())

	_, err = getKeyFile("../tests/.ssh/id_rsa")
	assert.NoError(t, err)
}

func TestRunCommand(t *testing.T) {
	ssh := &MakeConfig{
		Server:  "localhost",
		User:    "drone-scp",
		Port:    "22",
		KeyPath: "../tests/.ssh/id_rsa",
	}

	output, err := ssh.Run("whoami")
	assert.Equal(t, "drone-scp\n", output)
	assert.NoError(t, err)
}

func TestSCPCommand(t *testing.T) {
	ssh := &MakeConfig{
		Server:  "localhost",
		User:    "drone-scp",
		Port:    "22",
		KeyPath: "../tests/.ssh/id_rsa",
	}

	err := ssh.Scp("../tests/a.txt")
	assert.NoError(t, err)

	u, err := user.Lookup("drone-scp")
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}

	// check file exist
	if _, err := os.Stat(u.HomeDir + "/a.txt"); os.IsNotExist(err) {
		t.Fatalf("SCP-error: %v", err)
	}
}

func TestSCPCommandWithKey(t *testing.T) {
	ssh := &MakeConfig{
		Server: "localhost",
		User:   "drone-scp",
		Port:   "22",
		Key: `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA4e2D/qPN08pzTac+a8ZmlP1ziJOXk45CynMPtva0rtK/RB26
VbfAF0hIJji7ltvnYnqCU9oFfvEM33cTn7T96+od8ib/Vz25YU8ZbstqtIskPuwC
bv3K0mAHgsviJyRD7yM+QKTbBQEgbGuW6gtbMKhiYfiIB4Dyj7AdS/fk3v26wDgz
7SHI5OBqu9bv1KhxQYdFEnU3PAtAqeccgzNpbH3eYLyGzuUxEIJlhpZ/uU2G9ppj
/cSrONVPiI8Ahi4RrlZjmP5l57/sq1ClGulyLpFcMw68kP5FikyqHpHJHRBNgU57
1y0Ph33SjBbs0haCIAcmreWEhGe+/OXnJe6VUQIDAQABAoIBAH97emORIm9DaVSD
7mD6DqA7c5m5Tmpgd6eszU08YC/Vkz9oVuBPUwDQNIX8tT0m0KVs42VVPIyoj874
bgZMJoucC1G8V5Bur9AMxhkShx9g9A7dNXJTmsKilRpk2TOk7wBdLp9jZoKoZBdJ
jlp6FfaazQjjKD6zsCsMATwAoRCBpBNsmT6QDN0n0bIgY0tE6YGQaDdka0dAv68G
R0VZrcJ9voT6+f+rgJLoojn2DAu6iXaM99Gv8FK91YCymbQlXXgrk6CyS0IHexN7
V7a3k767KnRbrkqd3o6JyNun/CrUjQwHs1IQH34tvkWScbseRaFehcAm6mLT93RP
muauvMECgYEA9AXGtfDMse0FhvDPZx4mx8x+vcfsLvDHcDLkf/lbyPpu97C27b/z
ia07bu5TAXesUZrWZtKA5KeRE5doQSdTOv1N28BEr8ZwzDJwfn0DPUYUOxsN2iIy
MheO5A45Ko7bjKJVkZ61Mb1UxtqCTF9mqu9R3PBdJGthWOd+HUvF460CgYEA7QRf
Z8+vpGA+eSuu29e0xgRKnRzed5zXYpcI4aERc3JzBgO4Z0er9G8l66OWVGdMfpe6
CBajC5ToIiT8zqoYxXwqJgN+glir4gJe3mm8J703QfArZiQrdk0NTi5bY7+vLLG/
knTrtpdsKih6r3kjhuPPaAsIwmMxIydFvATKjLUCgYEAh/y4EihRSk5WKC8GxeZt
oiZ58vT4z+fqnMIfyJmD5up48JuQNcokw/LADj/ODiFM7GUnWkGxBrvDA3H67WQm
49bJjs8E+BfUQFdTjYnJRlpJZ+7Zt1gbNQMf5ENw5CCchTDqEq6pN0DVf8PBnSIF
KvkXW9KvdV5J76uCAn15mDkCgYA1y8dHzbjlCz9Cy2pt1aDfTPwOew33gi7U3skS
RTerx29aDyAcuQTLfyrROBkX4TZYiWGdEl5Bc7PYhCKpWawzrsH2TNa7CRtCOh2E
R+V/84+GNNf04ALJYCXD9/ugQVKmR1XfDRCvKeFQFE38Y/dvV2etCswbKt5tRy2p
xkCe/QKBgQCkLqafD4S20YHf6WTp3jp/4H/qEy2X2a8gdVVBi1uKkGDXr0n+AoVU
ib4KbP5ovZlrjL++akMQ7V2fHzuQIFWnCkDA5c2ZAqzlM+ZN+HRG7gWur7Bt4XH1
7XC9wlRna4b3Ln8ew3q1ZcBjXwD4ppbTlmwAfQIaZTGJUgQbdsO9YA==
-----END RSA PRIVATE KEY-----
`,
	}

	err := ssh.Scp("../tests/a.txt")
	assert.NoError(t, err)

	u, err := user.Lookup("drone-scp")
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}

	// check file exist
	if _, err := os.Stat(u.HomeDir + "/a.txt"); os.IsNotExist(err) {
		t.Fatalf("SCP-error: %v", err)
	}
}

func TestSCPCommandWithPassword(t *testing.T) {
	ssh := &MakeConfig{
		Server:   "localhost",
		User:     "drone-scp",
		Port:     "22",
		Password: "1234",
	}

	err := ssh.Scp("../tests/b.txt")
	assert.NoError(t, err)

	u, err := user.Lookup("drone-scp")
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}

	// check file exist
	if _, err := os.Stat(u.HomeDir + "/b.txt"); os.IsNotExist(err) {
		t.Fatalf("SCP-error: %v", err)
	}
}