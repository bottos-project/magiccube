/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Service Layer
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
 */
package keystore

import (
    "time"	
    "os"
    "os/user"
    "crypto/ecdsa"
    "crypto/elliptic"
    "io"
    "fmt"
    "log"
    "strings"
    "encoding/hex"
   // "github.com/code/bottos/service/node/common"
   // "github.com/code/bottos/service/node/common/math"
   // "github.com/code/bottos/service/node/crypto"
   // "github.com/code/bottos/service/node/crypto/randentropy"
    "github.com/bottos-project/magiccube/service/node/config"
    "github.com/bottos-project/magiccube/service/node/api"
    "github.com/peterh/liner"
    "gopkg.in/urfave/cli.v1"  
    "runtime"
     crand "crypto/rand"
    "github.com/bottos-project/magiccube/service/node/keystore/crypto-go/crypto/aes"
	//log "github.com/jeanphorn/log4go"
        //"github.com/micro/go-micro"
        //proto "github.com/code/bottos/service/asset/proto"
        //"golang.org/x/net/context"
        //"github.com/mikemintang/go-curl"
        //"github.com/bitly/go-simplejson"
        //"time"
        //storage "github.com/code/bottos/service/storage/proto"
        //"github.com/micro/go-micro/client"
        //"bytes"
        "io/ioutil"
        //"fmt"
        //"strconv"
        //"golang.org/x/net/html/atom"
        //"encoding/json"
        //"github.com/code/bottos/config"
        //"gopkg.in/mgo.v2/bson"
        //"github.com/code/bottos/service/bean"
        //"github.com/code/bottos/tools/db/mongodb"
        //"errors"
        //cbb "github.com/code/bottos/service/asset/cbb"
	"path/filepath"
)

var (
UserPwd = ""
)

var (
	KeyStoreScheme = "keystore"
	
	myNodeAccountPubKey = ""
	myNodeAccountPriKey = ""
	myNodeAccount       = "" 
	myNodeUUID          = ""
)

type NodeConfig struct {
	// Name sets the instance name of the node. It must not contain the / character and is
	// used in the devp2p node identifier. The instance name of geth is "geth". If no
	// value is specified, the basename of the current executable is used.
	Name string `toml:"-"`

	// UserIdent, if set, is used as an additional component in the devp2p node identifier.
	UserIdent string `toml:",omitempty"`

	// Version should be set to the version number of the program. It is used
	// in the devp2p node identifier.
	Version string `toml:"-"`

	// DataDir is the file system folder the node should use for any data storage
	// requirements. The configured data directory will not be directly shared with
	// registered services, instead those can use utility methods to create/access
	// databases or flat files. This enables ephemeral nodes which can fully reside
	// in memory.
	DataDir string

	// Configuration of peer-to-peer networking.
	//P2P p2p.Config

	// KeyStoreDir is the file system folder that contains private keys. The directory can
	// be specified as a relative path, in which case it is resolved relative to the
	// current directory.
	//
	// If KeyStoreDir is empty, the default location is the "keystore" subdirectory of
	// DataDir. If DataDir is unspecified and KeyStoreDir is empty, an ephemeral directory
	// is created by New and destroyed when the node is stopped.
	KeyStoreDir string `toml:",omitempty"`

	// UseLightweightKDF lowers the memory and CPU requirements of the key store
	// scrypt KDF at the expense of security.
	UseLightweightKDF bool `toml:",omitempty"`

	// NoUSB disables hardware wallet monitoring and connectivity.
	NoUSB bool `toml:",omitempty"`

	// IPCPath is the requested location to place the IPC endpoint. If the path is
	// a simple file name, it is placed inside the data directory (or on the root
	// pipe path on Windows), whereas if it's a resolvable path name (absolute or
	// relative), then that specific path is enforced. An empty path disables IPC.
	IPCPath string `toml:",omitempty"`

	// HTTPHost is the host interface on which to start the HTTP RPC server. If this
	// field is empty, no HTTP API endpoint will be started.
	HTTPHost string `toml:",omitempty"`

	// HTTPPort is the TCP port number on which to start the HTTP RPC server. The
	// default zero value is/ valid and will pick a port number randomly (useful
	// for ephemeral nodes).
	HTTPPort int `toml:",omitempty"`

	// HTTPCors is the Cross-Origin Resource Sharing header to send to requesting
	// clients. Please be aware that CORS is a browser enforced security, it's fully
	// useless for custom HTTP clients.
	HTTPCors []string `toml:",omitempty"`

	// HTTPVirtualHosts is the list of virtual hostnames which are allowed on incoming requests.
	// This is by default {'localhost'}. Using this prevents attacks like
	// DNS rebinding, which bypasses SOP by simply masquerading as being within the same
	// origin. These attacks do not utilize CORS, since they are not cross-domain.
	// By explicitly checking the Host-header, the server will not allow requests
	// made against the server with a malicious host domain.
	// Requests using ip address directly are not affected
	HTTPVirtualHosts []string `toml:",omitempty"`

	// HTTPModules is a list of API modules to expose via the HTTP RPC interface.
	// If the module list is empty, all RPC API endpoints designated public will be
	// exposed.
	HTTPModules []string `toml:",omitempty"`

	// WSHost is the host interface on which to start the websocket RPC server. If
	// this field is empty, no websocket API endpoint will be started.
	WSHost string `toml:",omitempty"`

	// WSPort is the TCP port number on which to start the websocket RPC server. The
	// default zero value is/ valid and will pick a port number randomly (useful for
	// ephemeral nodes).
	WSPort int `toml:",omitempty"`

	// WSOrigins is the list of domain to accept websocket requests from. Please be
	// aware that the server can only act upon the HTTP request the client sends and
	// cannot verify the validity of the request header.
	WSOrigins []string `toml:",omitempty"`

	// WSModules is a list of API modules to expose via the websocket RPC interface.
	// If the module list is empty, all RPC API endpoints designated public will be
	// exposed.
	WSModules []string `toml:",omitempty"`

	// WSExposeAll exposes all API modules via the WebSocket RPC interface rather
	// than just the public ones.
	//
	// *WARNING* Only set this if the node is running in a trusted network, exposing
	// private APIs to untrusted users is a major security risk.
	WSExposeAll bool `toml:",omitempty"`

	// Logger is a custom logger to use with the p2p.Server.
	Logger log.Logger `toml:",omitempty"`
}

type keyStore interface {
	// Loads and decrypts the key from disk.
	GetKey(addr aes.UUID, filename string, auth string) (*aes.Key, string, error)
	// Writes and encrypts the key.
	StoreKey(username string, filename string, k *aes.Key, auth string) error
	// Joins filename with the key directory unless it is already absolute.
	JoinPath(filename string) string
}

var DefaultConfig = NodeConfig {
	DataDir:         DefaultDataDir(),
    KeyStoreDir:     DefaultKeystoreDir(),
}

type keyStorePassphrase struct {
	keysDirPath string
	scryptN     int
	scryptP     int
}

type URL struct {
	Scheme string // Protocol scheme to identify a capable account backend
	Path   string // Path for the backend to identify a unique entity
}

// Account represents an Ethereum account located at a specific location defined
// by the optional URL field.
type Account struct {
	UUID aes.UUID `json:"address"` // Ethereum account address derived from the key
	URL     URL            `json:"url"`     // Optional resource locator within a backend
}

var PasswordFileFlag = cli.StringFlag{
		Name:  "password_filepath",
		Usage: "Password path file to use for non-interactive password input",
		Value: "",
	}

func MakePasswordList(ctx *cli.Context) []string {
    
    if len(ctx.String("password")) >0 {
        return strings.Split(ctx.String("password"), " ")
    } else if len(ctx.String("password_filepath")) <=0 {
        return nil
    }

	path := ctx.GlobalString(PasswordFileFlag.Name)
	if path == "" {
		return nil
	}
	text, err := ioutil.ReadFile(path)
	if err != nil {
		aes.Fatalf("Failed to read password file: %v", err)
	}
	lines := strings.Split(string(text), "\n")
	// Sanitise DOS line endings.
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}
	return lines
}

func keyFileName(keyAddr aes.UUID) string {
	ts := time.Now().UTC()
	return fmt.Sprintf("UTC--%s--%s", toISO8601(ts), hex.EncodeToString(keyAddr[:]))
}

func (ks keyStorePassphrase) GetKey(addr aes.UUID, filename, auth string) (*aes.Key, string, error) {
	// Load the key from the keystore and decrypt its contents
	keyjson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, "", err
	}
	key, Account, err := aes.DecryptKey(keyjson, auth)
	if err != nil {
		return nil, "", err
	}
	// Make sure we're really operating on the requested key (no swap attacks)
	if key.UUID != addr {
		return nil, "", fmt.Errorf("key content mismatch: have account %x, want %x", key.UUID, addr)
	}
	return key, Account, nil
}

// StoreKey generates a key, encrypts with 'auth' and stores in the given directory
func StoreKey(dir, auth string, scryptN, scryptP int) (aes.UUID, error) {
	/*key*/_, a, err := storeNewKey(&keyStorePassphrase{dir, scryptN, scryptP}, crand.Reader, auth)
    //fmt.Println("==> storeKey-> dir is: ", dir, ", key is:", key)

    return a.UUID, err
}

func (ks keyStorePassphrase) StoreKey(username string, filename string, key *aes.Key, auth string) error {
	keyjson, err := aes.EncryptKey(username, key, auth, ks.scryptN, ks.scryptP)
	if err != nil {
		return err
	}
	return aes.WriteKeyFile(filename, keyjson)
}

func (ks keyStorePassphrase) JoinPath(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	} else {
		return filepath.Join(ks.keysDirPath, filename)
	}
}

// FromECDSA exports a private key into a binary dump.
func FromECDSA(priv *ecdsa.PrivateKey) []byte {
    if priv == nil {
        return nil
    }
    return aes.PaddedBigBytes(priv.D, priv.Params().BitSize/8)
}

func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
    if pub == nil || pub.X == nil || pub.Y == nil {
            return nil
    }
    return elliptic.Marshal(aes.S256(), pub.X, pub.Y)
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}

func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()
	if home != "" {
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "bot")
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "bot\\workspace")
		} else {
			return filepath.Join(home, "bot/workspace")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}

func DefaultKeystoreDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()
	if home != "" {
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "bot")
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "bot")
		} else {
			return filepath.Join(home, "bot")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}

func defaultNodeConfig() NodeConfig {
	cfg := DefaultConfig //Dir setting
	cfg.Name = "user UUID"
	return cfg
}

func (c *NodeConfig) AccountConfig() (int, int, string) {
	scryptN := aes.StandardScryptN
	scryptP := aes.StandardScryptP
        
    keydir := c.KeyStoreDir
	return scryptN, scryptP, keydir
}

func zeroKey(k *ecdsa.PrivateKey) {
	b := k.D.Bits()
	for i := range b {
		b[i] = 0
	}
}

func storeNewKey(ks keyStore, rand io.Reader, auth string) (*aes.Key, Account, error) {
	
    key, err := aes.NewKey(rand)

	if err != nil {
		return nil, Account{}, err
	}

	a := Account{UUID: key.UUID, URL: URL{Scheme: KeyStoreScheme, Path: ks.JoinPath(keyFileName(key.UUID))}}
    
    nodeinfos := api.ReadFile(config.CONFIG_FILE)
    filename := "bto.keystore"
    filepath := ks.JoinPath(filename)

    if err := ks.StoreKey(nodeinfos.Node[0].UserName, filepath/*a.URL.Path*/, key, auth); err != nil {
		zeroKey(key.PrivateKey)
		return nil, a, err
	}
    fmt.Println("\n====================== Keystore Generated ===========================\n")
    fmt.Println("==>PATH: ", filepath, "\n==>KEY ID:",key.Id,"\n==>KEY UUID:", key.UUID,"\n==>PUB KEY:", hex.EncodeToString(FromECDSAPub(&aes.Pubkeytmp)), /*"KEY PRI KEY:", hex.EncodeToString(FromECDSA(key.PrivateKey)), "]*/"\n==>KEYSTORE FILE: ", filepath, "\n")
	
	myNodeAccountPubKey = hex.EncodeToString(FromECDSAPub(&aes.Pubkeytmp))
	myNodeAccountPriKey = hex.EncodeToString(FromECDSA(key.PrivateKey))
	myNodeAccount       = nodeinfos.Node[0].UserName
    myNodeUUID          = key.Id.String()
    
    fmt.Println("=====================================================================\n")

	return key, a, err
}

func GetPubKey() string { return myNodeAccountPubKey }
func GetPriKey() string { return myNodeAccountPriKey }
func GetAccount() string { return myNodeAccount }
func GetUUID() string { return myNodeUUID }

var Stdin = newTerminalPrompter()

// UserPrompter defines the methods needed by the console to promt the user for
// various types of inputs.
type UserPrompter interface {
	// PromptInput displays the given prompt to the user and requests some textual
	// data to be entered, returning the input of the user.
	PromptInput(prompt string) (string, error)

	// PromptPassword displays the given prompt to the user and requests some textual
	// data to be entered, but one which must not be echoed out into the terminal.
	// The method returns the input provided by the user.
	PromptPassword(prompt string) (string, error)

	// PromptConfirm displays the given prompt to the user and requests a boolean
	// choice to be made, returning that choice.
	PromptConfirm(prompt string) (bool, error)

	// SetHistory sets the the input scrollback history that the prompter will allow
	// the user to scroll back to.
	SetHistory(history []string)

	// AppendHistory appends an entry to the scrollback history. It should be called
	// if and only if the prompt to append was a valid command.
	AppendHistory(command string)

	// ClearHistory clears the entire history
	ClearHistory()

	// SetWordCompleter sets the completion function that the prompter will call to
	// fetch completion candidates when the user presses tab.
	SetWordCompleter(completer WordCompleter)
}

// WordCompleter takes the currently edited line with the cursor position and
// returns the completion candidates for the partial word to be completed. If
// the line is "Hello, wo!!!" and the cursor is before the first '!', ("Hello,
// wo!!!", 9) is passed to the completer which may returns ("Hello, ", {"world",
// "Word"}, "!!!") to have "Hello, world!!!".
type WordCompleter func(line string, pos int) (string, []string, string)

// terminalPrompter is a UserPrompter backed by the liner package. It supports
// prompting the user for various input, among others for non-echoing password
// input.
type terminalPrompter struct {
	*liner.State
	warned     bool
	supported  bool
	normalMode liner.ModeApplier
	rawMode    liner.ModeApplier
}

// newTerminalPrompter creates a liner based user input prompter working off the
// standard input and output streams.
func newTerminalPrompter() *terminalPrompter {
	p := new(terminalPrompter)
	// Get the original mode before calling NewLiner.
	// This is usually regular "cooked" mode where characters echo.
	normalMode, _ := liner.TerminalMode()
	// Turn on liner. It switches to raw mode.
	p.State = liner.NewLiner()
	rawMode, err := liner.TerminalMode()
	if err != nil || !liner.TerminalSupported() {
		p.supported = false
	} else {
		p.supported = true
		p.normalMode = normalMode
		p.rawMode = rawMode
		// Switch back to normal mode while we're not prompting.
		normalMode.ApplyMode()
	}
	p.SetCtrlCAborts(true)
	p.SetTabCompletionStyle(liner.TabPrints)
	p.SetMultiLineMode(true)
	return p
}

// PromptInput displays the given prompt to the user and requests some textual
// data to be entered, returning the input of the user.
func (p *terminalPrompter) PromptInput(prompt string) (string, error) {
	if p.supported {
		p.rawMode.ApplyMode()
		defer p.normalMode.ApplyMode()
	} else {
		// liner tries to be smart about printing the prompt
		// and doesn't print anything if input is redirected.
		// Un-smart it by printing the prompt always.
		fmt.Print(prompt)
		prompt = ""
		defer fmt.Println()
	}
	return p.State.Prompt(prompt)
}

// PromptPassword displays the given prompt to the user and requests some textual
// data to be entered, but one which must not be echoed out into the terminal.
// The method returns the input provided by the user.
func (p *terminalPrompter) PromptPassword(prompt string) (passwd string, err error) {
	if p.supported {
		p.rawMode.ApplyMode()
		defer p.normalMode.ApplyMode()
		return p.State.PasswordPrompt(prompt)
	}
	if !p.warned {
		fmt.Println("!! Unsupported terminal, password will be echoed.")
		p.warned = true
	}
	// Just as in Prompt, handle printing the prompt here instead of relying on liner.
	fmt.Print(prompt)
	passwd, err = p.State.Prompt("")
	fmt.Println()
	return passwd, err
}

// PromptConfirm displays the given prompt to the user and requests a boolean
// choice to be made, returning that choice.
func (p *terminalPrompter) PromptConfirm(prompt string) (bool, error) {
	input, err := p.Prompt(prompt + " [y/N] ")
	if len(input) > 0 && strings.ToUpper(input[:1]) == "Y" {
		return true, nil
	}
	return false, err
}

// SetHistory sets the the input scrollback history that the prompter will allow
// the user to scroll back to.
func (p *terminalPrompter) SetHistory(history []string) {
	p.State.ReadHistory(strings.NewReader(strings.Join(history, "\n")))
}

// AppendHistory appends an entry to the scrollback history.
func (p *terminalPrompter) AppendHistory(command string) {
	p.State.AppendHistory(command)
}

// ClearHistory clears the entire history
func (p *terminalPrompter) ClearHistory() {
	p.State.ClearHistory()
}

// SetWordCompleter sets the completion function that the prompter will call to
// fetch completion candidates when the user presses tab.
func (p *terminalPrompter) SetWordCompleter(completer WordCompleter) {
	p.State.SetWordCompleter(liner.WordCompleter(completer))
}

// getPassPhrase retrieves the password associated with an account, either fetched
// from a list of preloaded passphrases, or requested interactively from the user.
func getPassPhrase(prompt string, confirmation bool, i int, passwords []string) string {
	// If a list of passwords was supplied, retrieve from them
	if len(passwords) > 0 {
		if i < len(passwords) {
			return passwords[i]
		}
		return passwords[len(passwords)-1]
	}
	// Otherwise prompt the user for the password
	if prompt != "" {
		fmt.Println(prompt)
	}
	password, err := Stdin.PromptPassword("Passphrase: ")
	if err != nil {
		aes.Fatalf("Failed to read passphrase: %v", err)
	}
	if confirmation {
		confirm, err := Stdin.PromptPassword("Repeat passphrase: ")
		if err != nil {
			aes.Fatalf("Failed to read passphrase confirmation: %v", err)
		}
		if password != confirm {
			aes.Fatalf("Passphrases do not match")
		}
	}
	return password
}

func toISO8601(t time.Time) string {
	var tz string
	name, offset := t.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tz)
}

func accountCreate(ctx *cli.Context) error {
    
    if len(ctx.String("keystore_decrypt")) > 0 {
        if len(ctx.String("password")) <= 0 {
            fmt.Println("WRONG! Please input format with --keystore_decrypt <keystore file path> --password <your password>!")
            return nil
        }
        UserPwd = ctx.String("password")
        fmt.Println("===> test keystore_decrypt: ", ctx.String("keystore_decrypt"))
        aes.KeyDecrypt(ctx.String("keystore_decrypt"), ctx.String("password"))
        return nil
    }

	NodeCfg := defaultNodeConfig()

    if len(ctx.String("data_dir")) > 0 {
        NodeCfg.DataDir = ctx.String("data_dir")
    }

    if len(ctx.String("keystore_dir")) > 0 {
        NodeCfg.KeyStoreDir = ctx.String("keystore_dir")
	}
    
    scryptN, scryptP, keydir := NodeCfg.AccountConfig()
    //fmt.Println("===> parameters is:", os.Args, ", data_dir is: ", NodeCfg.DataDir, ", keystore_dir is:", NodeCfg.KeyStoreDir)

	password := getPassPhrase("Your new account is locked with a password. Please give a password. Do not forget this password.", true, 0, MakePasswordList(ctx))
    
    UserPwd = password

	/*UUID, err := */StoreKey(keydir, password, scryptN, scryptP)

	//fmt.Printf("UUID: {%x}, err: %s\n", UUID, err)
	return nil
}

func AccountCreate_Ex( datadir string , keydir string, password string) error {
    if len(datadir) <= 0 || len(keydir) <= 0 || len(password) <= 0 {
        log.Println("keydir or UserPwd invalid: datadir:", datadir ,", keydir:", keydir, ", password:", password)
        return nil
    }

	NodeCfg := defaultNodeConfig()

    NodeCfg.DataDir = datadir

    NodeCfg.KeyStoreDir = keydir
    
    scryptN, scryptP, keydir := NodeCfg.AccountConfig()
    //log.Println("===>  data_dir is: ", NodeCfg.DataDir, ", keystore_dir is:", NodeCfg.KeyStoreDir, ", password is:", password)
    
    UserPwd = password

	/*UUID, err := */StoreKey(keydir, password, scryptN, scryptP)

	//fmt.Printf("UUID: {%x}, err: %s\n", UUID, err)
	return nil
}

func main() {

    app := cli.NewApp()

    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "keystore_dir",
            Value: "",
            Usage: "keystore file's dir",
        },
        cli.StringFlag{
            Name: "data_dir",
            Value: "",
            Usage: "data files' dir",
        },
        cli.StringFlag{
            Name: "password",
            Value: "",
            Usage: "Password by default",
        },
        cli.StringFlag{
            Name: "password_filepath",
            Value: "",
            Usage: "Password path file stored in path",
        },
        cli.StringFlag{
            Name: "keystore_decrypt",
            Value: "",
            Usage: "Decryption keystore file, need user input same password as before",
        },
    }

    app.Action = accountCreate

    app.Run(os.Args)
}
