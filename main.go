package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	discordToken := os.Getenv("DISCORD_TOKEN")
	gpt4oAPIKey := os.Getenv("GPT4O_API_KEY")

	if discordToken == "" || gpt4oAPIKey == "" {
		fmt.Println("DISCORD_TOKEN and GPT4O_API_KEY must be set")
		return
	}

	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !isMentioned(m) {
		return
	}

	response, err := generateResponse(m.Content)
	if err != nil {
		fmt.Println("error generating response,", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, response)
}

func isMentioned(m *discordgo.MessageCreate) bool {
	for _, user := range m.Mentions {
		if user.ID == m.Author.ID {
			return true
		}
	}
	return false
}

func generateResponse(message string) (string, error) {
	client := openai.NewClient(option.WithAPIKey(os.Getenv("GPT4O_API_KEY")))
	ctx := context.Background()

	chatCompletion, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		return "", err
	}

	return chatCompletion.Choices[0].Message.Content, nil
}
