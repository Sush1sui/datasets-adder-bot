package commands

import (
	"bytes"
	"encoding/json"

	"github.com/Sush1sui/datasets_adder/internal/repository"
	"github.com/bwmarrin/discordgo"
)

func GetAllUsers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Member == nil || i.GuildID == "" {
		return
	}

		// defer reply
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Getting all users...",
		},
	})

	// Get all user accounts
	userAccounts, err := repository.UserAccountService.GetAllUserAccounts()
	if err != nil {
		mess := "Error getting user accounts: " + err.Error()
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}

	// Marshal to pretty JSON
    jsonBytes, err := json.MarshalIndent(userAccounts, "", "  ")
    if err != nil {
        mess := "Error marshaling user accounts: " + err.Error()
        s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
            Content: &mess,
        })
        return
    }

    // Create a buffer for the file
    file := &discordgo.File{
        Name:   "user_accounts.json",
        Reader: bytes.NewReader(jsonBytes),
    }

    // Send the file as an attachment
    _, err = s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
        Content: "Here are all user accounts as JSON:",
        Files:   []*discordgo.File{file},
    })
    if err != nil {
        mess := "Error sending file: " + err.Error()
        s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
            Content: &mess,
        })
        return
    }

    mess := "User accounts JSON file uploaded."
    s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
        Content: &mess,
    })
}