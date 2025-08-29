package commands

import (
	"github.com/Sush1sui/datasets_adder/internal/repository"
	"github.com/bwmarrin/discordgo"
)

func DeleteAccountByEmail(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Member == nil || i.GuildID == "" {
		return
	}

	// defer reply
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Deleting account...",
		},
	})

	// get string option named as "email"
	email := i.ApplicationCommandData().GetOption("email").StringValue()

	if email == "" {
		mess := "Email is required."
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}

	count, err := repository.UserAccountService.DeleteUserAccountByEmail(email)
	if err != nil {
		mess := "Error deleting account: " + err.Error()
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}
	if count == 0 {
		mess := "No account found with email " + email + "."
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}

	mess := "Account with email " + email + " has been deleted."
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &mess,
	})
}