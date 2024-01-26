package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

const maxLengthUsername int = 20

// Client représente le client du chat en temps réel.
type Client struct {
	IP          string
	PORT        string
	Username    string
	Conn        net.Conn
	IsConnected bool
}

// New crée une nouvelle instance de Client.
func New(IP string, PORT string) (*Client, error) {
	client := &Client{
		IP:          IP,
		PORT:        PORT,
		IsConnected: false,
	}

	err := client.connect()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// usernameHandle gère la saisie du nom d'utilisateur par le client.
func (client *Client) usernameHandle() {
	for {
		username := client.getUsernameInput()
		lengthUsername := len(strings.TrimSuffix(username, "\n"))

		if lengthUsername > maxLengthUsername || lengthUsername == 0 {
			fmt.Println("[ERROR] Your username must not be empty or exceed", maxLengthUsername, "characters")
			continue
		}

		client.Conn.Write([]byte(username))
		serverResponse, err := client.read()
		client.check(err)

		if client.isUsernameGood(serverResponse) {
			client.Username = strings.TrimSuffix(username, "\n")
			fmt.Print("[SUCCESS] You are successfully connected!\n")
			break
		}
	}
}

// getUsernameInput récupère la saisie du nom d'utilisateur.
func (client *Client) getUsernameInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username : ")
	username, err := reader.ReadString('\n')

	client.check(err)
	return username
}

// isUsernameGood vérifie si le nom d'utilisateur est accepté par le serveur.
func (client *Client) isUsernameGood(response string) bool {
	if response == "badUsername" {
		fmt.Println("[ERROR] Your username already exists in the server, please enter another username")
		return false
	} else {
		fmt.Println("[SUCCESS] Your username is accepted by the server")
		return true
	}
}

// connect établit la connexion avec le serveur.
func (client *Client) connect() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", client.IP, client.PORT))
	if err != nil {
		return err
	}
	fmt.Println("Connecting to", conn.RemoteAddr(), "SERVER ...")
	client.IsConnected = true
	client.Conn = conn
	return nil
}

// check vérifie une erreur et ferme la connexion du client en cas d'erreur.
func (client *Client) check(err error) {
	if err != nil {
		client.IsConnected = false
		if client.Conn != nil {
			client.Conn.Close()
		}
		fmt.Println(err)
		fmt.Println("You are now disconnected !")
		os.Exit(2)
	}
}

// send obtient l'entrée de l'utilisateur et l'envoie au serveur.
func (client *Client) send() {
	defer wg.Done()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Votre message : ")

		message, err := reader.ReadString('\n')
		if !client.IsConnected {
			break
		}
		client.check(err)

		client.Conn.Write([]byte(message))
	}
}

// read lit les messages envoyés par le serveur.
func (client *Client) read() (string, error) {
	messageBuffer := make([]byte, 4096)
	length, err := client.Conn.Read(messageBuffer)
	if err != nil {
		fmt.Println("[INFO] Server is down, click Enter to close the session")
		client.IsConnected = false
	}
	message := string(messageBuffer[:length])
	return message, err
}

// receive lit tous les messages envoyés par le serveur.
func (client *Client) receive() {
	defer wg.Done()
	for {
		message, err := client.read()
		if !client.IsConnected {
			break
		}
		if err != nil {
			fmt.Println("[INFO] Server is down, click Enter to close the session")
			client.IsConnected = false
			break
		}
		fmt.Print(message)
	}
}

// Run démarre la connexion du client et commence la conversation avec les autres clients du serveur.
func (client *Client) Run() {
	client.connect()
	client.usernameHandle()
	wg.Add(2)
	go client.send()
	go client.receive()
	wg.Wait()
	fmt.Println("You are now disconnected !")
}
