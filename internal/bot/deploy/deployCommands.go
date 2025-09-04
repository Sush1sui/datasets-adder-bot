package deploy

import (
	"fmt"
	"log"

	"github.com/Sush1sui/datasets_adder/internal/bot/commands"
	"github.com/Sush1sui/datasets_adder/internal/config"
	"github.com/bwmarrin/discordgo"
)

var slashCommands = []*discordgo.ApplicationCommand{
	{
		Name: "delete-acc-by-email",
		Description: "Delete user account by email",
		Type: discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: func() *int64 { p := int64(discordgo.PermissionAdministrator); return &p }(),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "email",
				Description: "Email of the account to delete",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	{
		Name: "get-all-users",
		Description: "Get all user accounts",
		Type: discordgo.ChatApplicationCommand,
	},
	{
		Name: "get-user-by-email",
		Description: "Get user account by email",
		Type: discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "email",
				Description: "Email of the account to get",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	{
		Name: "delete-all-users",
		Description: "Delete all user accounts",
		Type: discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: func() *int64 { p := int64(discordgo.PermissionAdministrator); return &p }(),
	},
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"delete-acc-by-email": commands.DeleteAccountByEmail,
	"get-all-users": commands.GetAllUsers,
	"get-user-by-email": commands.GetUserByEmail,
	"delete-all-users": commands.DeleteAllUsers,
}

func DeployCommands(s *discordgo.Session) {
	globalCmds, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		for _, cmd := range globalCmds {
			err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
			if err != nil {
				log.Printf("Failed to delete command %s: %v", cmd.Name, err)
			} else {
				log.Printf("Deleted command: %s", cmd.Name)
			}
		}
	}

	if len(slashCommands) == 0 { return }

	guilds := s.State.Guilds
	for _, guild := range guilds {
		if guild.ID == config.Global.GuildID {
			_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, guild.ID, slashCommands)
			if err != nil {
				log.Printf("Failed to deploy commands to guild %s: %v", guild.Name, err)
			} else {
				log.Printf("Successfully deployed commands to guild: %s", guild.Name)
			}
			break
		}
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand { return }

		if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		} else {
			fmt.Println("Unknown command:", i.ApplicationCommandData().Name)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Unknown command. Please try again.",
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
		}
	})
}