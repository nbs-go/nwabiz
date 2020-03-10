package nwabiz

import (
	"os"
	"strconv"
	"testing"
	"time"
)

var provider *WhatsAppBiz

func TestMain(m *testing.M) {
	// Prepare opt
	opt := InitOpt{
		BaseUrl:  os.Getenv("NWABIZ_TEST_BASE_URL"),
		Username: os.Getenv("NWABIZ_TEST_USERNAME"),
		Password: os.Getenv("NWABIZ_TEST_PASSWORD"),
	}

	// Get insecure ssl opt
	insecureSSLStr := os.Getenv("NWABIZ_TEST_INSECURE_SSL")
	insecureSSL, err := strconv.ParseBool(insecureSSLStr)
	if err != nil {
		insecureSSL = false
	}
	opt.InsecureSSL = insecureSSL

	// Set timeout
	timeoutStr := os.Getenv("NWABIZ_TEST_CLIENT_TIMEOUT")
	timeout, err := strconv.ParseInt(timeoutStr, 10, 64)
	if err != nil {
		timeout = 10000
	}
	opt.Timeout = timeout

	// Init provider
	provider = NewWhatsAppBiz(opt)

	// Run test
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestLogin(t *testing.T) {
	err := provider.Login()
	if err != nil {
		t.Errorf("unable to login. Error: (%s)", err)
		return
	}

	if !provider.IsAccessValid(time.Now()) {
		t.Error("Failed: Invalid Access")
		return
	}

	t.Logf("Pass")
}

func TestCheckContactInvalid(t *testing.T) {
	// Get test case
	input := os.Getenv("NWABIZ_TEST_CASE_CONTACT_INVALID")

	// Test invalid case
	contact, err := provider.CheckContact(input)
	if err != nil {
		t.Errorf("unable to check invalid contact. Error: (%s)", err)
		return
	}

	if contact.Status != InvalidStatus {
		t.Errorf("Failed: Invalid status expected. Got %s", contact.Status)
		return
	}

	t.Log("Pass")
}

func TestCheckContactValid(t *testing.T) {
	// Get test case
	input := os.Getenv("NWABIZ_TEST_CASE_CONTACT_VALID")

	// Test invalid case
	contact, err := provider.CheckContact(input)
	if err != nil {
		t.Errorf("unable to check valid contact. Error: (%s)", err)
		return
	}

	if contact.Status != ValidStatus {
		t.Errorf("Failed: Valid status expected. Got %s", contact.Status)
		return
	}

	t.Log("Pass")
}
