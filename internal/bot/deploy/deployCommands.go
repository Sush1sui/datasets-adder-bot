package deploy

import (
	"fmt"
	"log"

	"github.com/Sush1sui/datasets_adder/internal/config"
	"github.com/bwmarrin/discordgo"
)

var slashCommands = []*discordgo.ApplicationCommand{
	
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	
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