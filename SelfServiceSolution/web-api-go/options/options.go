package options

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"os"

	"github.com/gorilla/securecookie"
)

type Options struct {
	KratosPublicURL     *url.URL      // KratosAdminURL is the URL where ORY Kratos's Admin API is located at.
	KratosAdminURL      *url.URL      //KratosPublicURL is the URL where ORY Kratos's Public API is located at.
	KratosBrowserURL    *url.URL      // KratosBrowserURL is the URL where ORY Kratos's self service browser endpoints are located at.
	BaseURL             *url.URL      // BaseURL is the base url of this app. If served e.g. behind a proxy or via GitHub pages. Must be absolute!
	Host                string        // Host that the app is listening on. Used together with Port
	Port                int           // Port that this app is listening on. Used together with Host
	ShutdownWait        time.Duration // Duration to wait when asked to shutdown gracefully
	TLSCertPath         string        // TLSCertPath is an optional Path to certificate file. Should be set up together with TLSKeyPath to enable HTTPS.
	TLSKeyPath          string        // TLSCertPath is an optional path to key file. Should be set up together with TLSCertPath to enable HTTPS.
	CookieStoreKeyPairs [][]byte      // Pairs of authentication and encryption keys for Cookies
}

func NewOptions() *Options {
	return &Options{
		KratosAdminURL:   &url.URL{},
		KratosPublicURL:  &url.URL{},
		KratosBrowserURL: &url.URL{},
		BaseURL:          &url.URL{},
	}
}

// SetFromCommandLine will parse the command line, and populate the Options.
// The special case is when the 'gen-cookie-store-key-pair' is detected, will genrate the keys and exit
// Will also exit if key-pairs passed in are invalid
// TODO get these to be read from the command line https://github.com/davidoram/kratos-selfservice-ui-go/blob/main/options/options.go
func (o *Options) SetFronEnvVariables() *Options {
	o.KratosPublicURL, _ = url.Parse(os.Getenv("KRATOS_PUBLIC_URL"))
	o.KratosAdminURL, _ = url.Parse(os.Getenv("KRATOS_ADMIN_URL"))
	o.KratosBrowserURL, _ = url.Parse(os.Getenv("KRATOS_BROWSER_URL"))
	o.BaseURL, _ = url.Parse(os.Getenv("BASE_URL"))
	o.Host = os.Getenv("WEB_API_GO_HOST")
	o.Port = parseInt(os.Getenv("WEB_API_GO_PORT"))

	// flag.DurationVar(&o.ShutdownWait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	// flag.StringVar(&o.TLSCertPath, "tls-cert-path", "", "Optional path to the certificate file. Use in conjunction with tls-key-path to enable https.")

	// flag.StringVar(&o.TLSKeyPath, "tls-key-path", "", "Optional path to the key file. Use in conjunction with tls-cert-path to enable https.")

	// var allCookieStoreKeyPairs string
	// flag.StringVar(&allCookieStoreKeyPairs, "cookie-store-key-pairs", os.Getenv("COOKIE_STORE_KEY_PAIRS"), "Pairs of authentication and encryption keys, enclose then in quotes. See the gen-cookie-store-key-pair flag to generate")

	// genCookieStoreKeys := false
	// flag.BoolVar(&genCookieStoreKeys, "gen-cookie-store-key-pair", false, "Pass this flag to generate a pairs of authentication and encryption keys and exit")

	// flag.Parse()

	cookieKeyPairs := os.Getenv("COOKIE_KEY_PAIR")
	if cookieKeyPairs == "" {
		fmt.Println("No cookie key pair found. Generating cookies key pair.")
		authKey := securecookie.GenerateRandomKey(32)
		encrKey := securecookie.GenerateRandomKey(32)
		authKeyString := base64.StdEncoding.EncodeToString(authKey)
		encrKeyString := base64.StdEncoding.EncodeToString(encrKey)
		fmt.Printf("The following key pair was generated:\n")
		fmt.Printf("-> %s %s\n", authKeyString, encrKeyString)
		// os.Exit(0)
		cookieKeyPairs = authKeyString + " " + encrKeyString
	}

	// o.KratosAdminURL = KratosAdminURL.URL
	// o.KratosPublicURL = KratosPublicURL.URL
	// o.KratosBrowserURL = KratosBrowserURL.URL
	// o.BaseURL = BaseURL.URL
	// o.CookieStoreKeyPairs = make([][]byte, 0)
	pairs := strings.Split(cookieKeyPairs, " ")
	for _, s := range pairs {
		decoded, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			log.Fatalf("Error decoding 'cookie-store-key-pairs' value: '%s' , did you use 'gen-cookie-store-key-pair' to generate them? Error: %v", s, err)
		}
		o.CookieStoreKeyPairs = append(o.CookieStoreKeyPairs, []byte(decoded))
	}
	return o
}

func (o *Options) Validate() error {
	if o.KratosPublicURL == nil || o.KratosPublicURL.String() == "" {
		return errors.New("'kratos-public-url' URL missing")
	}
	if o.KratosAdminURL == nil || o.KratosAdminURL.String() == "" {
		return errors.New("'kratos-admin-url' URL missing")
	}
	if o.KratosBrowserURL == nil || o.KratosBrowserURL.String() == "" {
		return errors.New("'kratos-browser-url' URL missing")
	}
	if o.BaseURL == nil || o.BaseURL.String() == "" {
		return errors.New("'base-url' URL missing")
	}
	// if o.TLSCertPath != "" && !fileExists(o.TLSCertPath) {
	// 	return fmt.Errorf("'tls-cert-path' file '%s' invalid", o.TLSCertPath)
	// }
	// if o.TLSKeyPath != "" && !fileExists(o.TLSKeyPath) {
	// 	return fmt.Errorf("'tls-key-path' file '%s' invalid", o.TLSKeyPath)
	// }
	// if (o.TLSCertPath == "" && o.TLSKeyPath != "") || (o.TLSCertPath != "" && o.TLSKeyPath == "") {
	// 	return fmt.Errorf("To enable HTTPS, provide 'tls-key-path' and 'tls-cert-path'")
	// }
	// if !(len(o.CookieStoreKeyPairs) == 1 || len(o.CookieStoreKeyPairs)%2 == 0) {
	// 	return fmt.Errorf("'cookie-store-key-pairs' has %d values, it should contain one auth key, or even pairs of auth & encryption keys separated by a space", len(o.CookieStoreKeyPairs))
	// }
	return nil
}

func (o *Options) Display() {
	fmt.Println("Options:")
	fmt.Printf("- Kratos Admin URL   -> %v\n", o.KratosAdminURL)
	fmt.Printf("-- Kratos Admin URL Scheme -> %v\n", o.KratosAdminURL.Scheme)
	fmt.Printf("-- Kratos Admin URL Host   -> %v\n", o.KratosAdminURL.Host)
	fmt.Printf("-- Kratos Admin URL Path   -> %v\n", o.KratosAdminURL.Path)
	fmt.Printf("- Kratos Public URL  -> %v\n", o.KratosPublicURL)
	fmt.Printf("-- Kratos Public URL Scheme -> %v\n", o.KratosPublicURL.Scheme)
	fmt.Printf("-- Kratos Public URL Host   -> %v\n", o.KratosPublicURL.Host)
	fmt.Printf("-- Kratos Public URL Path   -> %v\n", o.KratosPublicURL.Path)
	fmt.Printf("- Kratos Browser URL -> %v\n", o.KratosBrowserURL)
	fmt.Printf("-- Kratos Browser URL Scheme -> %v\n", o.KratosBrowserURL.Scheme)
	fmt.Printf("-- Kratos Browser URL Host   -> %v\n", o.KratosBrowserURL.Host)
	fmt.Printf("-- Kratos Browser URL Path   -> %v\n", o.KratosBrowserURL.Path)
	fmt.Printf("- BaseURL -> %v\n", o.BaseURL)
	fmt.Printf("-- Base URL Scheme -> %v\n", o.BaseURL.Scheme)
	fmt.Printf("-- Base URL Host   -> %v\n", o.BaseURL.Host)
	fmt.Printf("-- Base URL Path   -> %v\n", o.BaseURL.Path)
	fmt.Printf("- WEB-API-GO Host -> %v\n", o.Host)
	fmt.Printf("- WEB-API-GP Port -> %v\n", o.Port)
}

// Address that this application will listen on
func (o *Options) Address() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

// LoginURL returns the URL to redirect to that shows the login page
func (o *Options) LoginPageURL() string {
	url := o.BaseURL
	url.Path = "/auth/login"
	return url.String()
}

// WhoAmIURL returns the URL to POST to to get the session
func (o *Options) WhoAmIFlowURL() string {
	url := o.KratosPublicURL
	url.Path = "/sessions/whoami"
	return url.String()
}

// RegistrationURL returns the URL to redirect to that will
// start the registration flow
func (o *Options) RegistrationFlowURL() string {
	url := o.KratosBrowserURL
	url.Path = "/self-service/registration/browser"
	return url.String()
}

// LoginFlowURL returns the URL to redirect to that will
// start the login flow
func (o *Options) LoginFlowURL() string {
	url := o.KratosBrowserURL
	url.Path = "/self-service/login/browser"
	return url.String()
}

// SettingsURL returns the URL to redirect to that will
// start the settings flow
func (o *Options) SettingsFlowURL() string {
	url := o.KratosBrowserURL
	url.Path = "/self-service/settings/browser"
	return url.String()
}

// RecoveryFlowURL returns the URL to redirect to that will
// start the recovery flow
func (o *Options) RecoveryFlowURL() string {
	url := o.KratosBrowserURL
	url.Path = "/self-service/recovery/browser"
	return url.String()
}

// LogoutFlowURL returns the URL to redirect to that will
// start the logout flow
func (o *Options) LogoutFlowURL() string {
	url := o.KratosBrowserURL
	url.Path = "/self-service/browser/flows/logout"
	return url.String()
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// func fileExists(filename string) bool {
// 	info, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return !info.IsDir()
// }
