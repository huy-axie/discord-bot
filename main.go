package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	Token = os.Getenv("TOKEN")
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Check message & send
	if strings.HasPrefix(m.Content, "flarectl") {

		s.ChannelMessageSend(m.ChannelID, "executing... ```"+m.Content+"```")
		stdout, _ := execShell(m.Content)
		// slpit & send
		if len(stdout) > 1950 {
			res := splitMessage(stdout, 1950)
			for _, msg := range res {
				s.ChannelMessageSend(m.ChannelID, "```"+msg+"```")
				time.Sleep(2 * time.Second)
			}
		} else {
			// Just send normaly
			s.ChannelMessageSend(m.ChannelID, "```"+stdout+"```")
		}
	}

	// If the message is "Ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

}

// exec bash & catch output
func execShell(command string) (string, error) {

	cmd := exec.Command("bash", "-c", command)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		println(err.Error())
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())

	if len(outStr) > 0 {
		println(len(outStr))
		return outStr, nil
	} else {
		println(len(errStr))
		return errStr, nil
	}

}

// split message by discord limit 2000 character
func splitMessage(msg string, chunkSize int) []string {
	if len(msg) == 0 {
		return nil
	}
	if chunkSize >= len(msg) {
		return []string{msg}
	}
	var chunks []string = make([]string, 0, (len(msg)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range msg {
		if currentLen == chunkSize {
			chunks = append(chunks, msg[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, msg[currentStart:])
	return chunks
}
