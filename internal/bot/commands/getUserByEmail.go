package commands

import (
	"bytes"
	"encoding/json"

	"github.com/Sush1sui/datasets_adder/internal/repository"
	"github.com/bwmarrin/discordgo"
)

func GetUserByEmail(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Member == nil || i.GuildID == "" { return }

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Getting user by email...",
		},
	})

	email := i.ApplicationCommandData().GetOption("email").StringValue()
	if email == "" {
		mess := "Email is required."
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}

	user, err := repository.UserAccountService.GetUserByEmail(email)
	if err != nil {
		mess := "Error getting user by email: " + err.Error()
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}

	if user == nil {
		mess := "User not found."
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}

	// Marshal to pretty JSON
	jsonBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		mess := "Error marshaling user account: " + err.Error()
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}

	// Create a buffer for the file
	file := &discordgo.File{
		Name:   "user_account.json",
		Reader: bytes.NewReader(jsonBytes),
	}

	// Send the file as an attachment
	_, err = s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Content: "Here is the user account as JSON:",
		Files:   []*discordgo.File{file},
	})
	if err != nil {
		mess := "Error sending file: " + err.Error()
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}

	mess := "User account JSON file uploaded."
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &mess,
	})
}