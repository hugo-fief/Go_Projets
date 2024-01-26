package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// Server représente le serveur du chat en temps réel.
type Server struct {
	IP       string
	PORT     string
	Listener net.Listener
	clients  map[net.Conn]string // @key : client and @value : username
}

// New crée une nouvelle instance de Server.
func New(IP string, PORT string) (*Server, error) {
	server := &Server{
		IP:      IP,
		PORT:    PORT,
		clients: make(map[net.Conn]string),
	}

	err := server.startConnection()
	if err != nil {
		return nil, err
	}

	return server, nil
}

// startConnection démarre le listener du serveur.
func (server *Server) startConnection() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.IP, server.PORT))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
		return err
	}
	server.Listener = listener
	log.Println("Server is running ...")
	return nil
}

// check vérifie une erreur et ferme le serveur en cas d'erreur.
func (server *Server) check(err error) {
	if err != nil {
		if server.Listener != nil {
			server.Listener.Close()
		}
		log.Println("Server is shutdown")
		log.Println(err)
		os.Exit(2)
	}
}

// addLog ajoute une ligne dans le fichier journal.
func (server *Server) addLog(line string) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		log.Println("ERROR :", err)
	}

	line = datetimeLine(line)
	_, err = file.WriteString(line)
	if err != nil {
		log.Println("ERROR :", err)
	}
	log.Print(line)
}

// datetimeLine ajoute la date et l'heure actuelles à une ligne.
func datetimeLine(text string) string {
	datetimeNow := time.Now().Format("02/01/2006 15:04:05") // Format MDD-MM-YYYY hh:mm:ss
	return fmt.Sprintf("[%s] %s", datetimeNow, text)
}

// addClient ajoute un client au serveur.
func (server *Server) addClient(client net.Conn) {
	if !server.checkUsernameClient(client) {
		return // exit the goroutine because the client interrupts the username input
	}

	message := fmt.Sprintf("[INFO] %s join the server", strings.TrimSpace(server.clients[client]))

	if len(server.clients) > 0 {
		client.Write([]byte("List of usernames in the server:\n"))
		for c := range server.clients {
			client.Write([]byte(fmt.Sprintf("-> %s\n", server.clients[c])))
		}
	}

	client.Write([]byte("You can start the discussion with guests ...\n\n"))
	server.sendToAll(client, message, true)
	server.addLog(fmt.Sprintf("%s connected from %s\n", server.clients[client], client.RemoteAddr()))
	server.receive(client)
}

// checkUsernameClient vérifie si le nom d'utilisateur du client existe déjà.
func (server *Server) checkUsernameClient(client net.Conn) bool {
	var (
		username string
		err      error
	)

	for {
		username, err = server.catchClientUsername(client)
		if err != nil {
			return false
		}

		if server.isUsernameExists(username, client) {
			client.Write([]byte("badUsername"))
		} else {
			client.Write([]byte("goodUsername"))
			break
		}
	}

	server.clients[client] = username // add client to the clients Map()
	return true
}

// isUsernameExists vérifie si le nom d'utilisateur existe déjà dans le Map.
func (server *Server) isUsernameExists(username string, client net.Conn) bool {
	findUsername := false
	for c := range server.clients {
		if server.clients[c] == username && client != c {
			log.Printf("The username %s already exists!", username)
			findUsername = true
			break
		}
	}
	return findUsername
}

// catchClientUsername récupère le nom d'utilisateur envoyé par le client.
func (server *Server) catchClientUsername(client net.Conn) (string, error) {
	var err error

	usernameBuffer := make([]byte, 4096)
	length, err := client.Read(usernameBuffer)

	username := string(usernameBuffer[:length])
	if err != nil {
		server.addLog(fmt.Sprintf("Client from %s interrupted the username input\n", client.RemoteAddr()))
		username = "error"
		err = errors.New("client interrupted input")
	}

	return strings.TrimSuffix(username, "\n"), err
}

// sendToAll envoie un message à tous les clients.
func (server *Server) sendToAll(sender net.Conn, message string, ignoreItself bool) {
	for client := range server.clients {
		if ignoreItself {
			if client != sender {
				client.Write([]byte(fmt.Sprintln(message)))
			}
		} else {
			client.Write([]byte(fmt.Sprintln(message)))
		}
	}
}

// receive reçoit les messages du client.
func (server *Server) receive(client net.Conn) {
	buf := bufio.NewReader(client)
	for {
		message, err := buf.ReadString('\n')
		if err != nil {
			server.removeClient(client)
			break
		}

		message = strings.TrimRight(message, "\r\n")

		clientName, ok := server.clients[client]
		if !ok {
			log.Println("Client not found:", client)
			continue
		}

		formattedMessage := fmt.Sprintf("Message de %s : %s", strings.TrimSpace(clientName), strings.TrimSpace(message))
		server.sendToAll(client, formattedMessage, true)
	}
}

// removeClient supprime un client du Map.
func (server *Server) removeClient(client net.Conn) {
	message := fmt.Sprintf("[INFO] %s is now disconnected\n", server.clients[client])
	server.sendToAll(client, message, true)
	server.addLog(fmt.Sprintf("%s is disconnected [total clients %d]\n", server.clients[client], len(server.clients)-1))
	delete(server.clients, client)
}

// Run démarre la connexion du serveur et attend les clients.
func (server *Server) Run() {
	for {
		client, err := server.Listener.Accept()
		server.check(err)
		go server.addClient(client)
	}
}
