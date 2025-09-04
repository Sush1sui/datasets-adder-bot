package commands

import (
	"fmt"

	"github.com/Sush1sui/datasets_adder/internal/repository"
	"github.com/bwmarrin/discordgo"
)

func DeleteAllUsers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Member == nil || i.GuildID == "" {
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Deleting accounts...",
		},
	})

	count, err := repository.UserAccountService.DeleteAllUserAccounts()
	if err != nil {
		mess := "Error deleting accounts: " + err.Error()
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &mess,
		})
		return
	}
	mess := "Deleted " + fmt.Sprint(count) + " accounts."
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &mess,
	})
}