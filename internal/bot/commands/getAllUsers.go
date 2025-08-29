package commands

import (
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
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Error getting user accounts: " + err.Error(),
			},
		})
		return
	}

	// Send the user accounts as a message
	var content string
	for _, user := range userAccounts {
		if user.Email != nil {
			content += *user.Email
		}
		content += "\n"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "User accounts:\n" + content,
		},
	})
}