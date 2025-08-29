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
		mess := "Error getting user accounts: " + err.Error()
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}

	// Send the user accounts as a message
	var content = "# User Accounts:\n\n"
	for _, user := range userAccounts {
		if user.Email != nil {
			content += *user.Email
		}
		content += "\n"
	}

	_, _ = s.ChannelMessageSend(i.ChannelID, content)

	mess := "User accounts successfully fetched"
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &mess,
	})
}