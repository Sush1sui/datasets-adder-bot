package events

import (
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func OnUploadFood(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Member == nil || m.GuildID == "" { return }
	if m.ChannelID != "1350438584548589660" { return }
	if !strings.HasPrefix(m.Content, "!food ") { return }

	// Parse food name
    args := strings.Fields(m.Content)
    if len(args) < 2 {
        s.ChannelMessageSend(m.ChannelID, "Please provide a food name. Usage: `!food <food_name>`")
        return
    }
    foodName := args[1]

	// Find channel with the food name
    channels, err := s.GuildChannels(m.GuildID)
    if err != nil {
        s.ChannelMessageSend(m.ChannelID, "Error fetching channels.")
        return
    }

	var targetChannel *discordgo.Channel
    for _, ch := range channels {
        if ch.Type == discordgo.ChannelTypeGuildText && ch.Name == foodName {
            targetChannel = ch
            break
        }
    }

	if targetChannel == nil {
        s.ChannelMessageSend(m.ChannelID, "No channel found with the name: "+foodName)
        return
    }

	// Check for attachments
    if len(m.Attachments) == 0 {
        s.ChannelMessageSend(m.ChannelID, "Please attach a file to upload.")
        return
    }

	// Download and re-upload each attachment to the target channel
    for _, att := range m.Attachments {
        resp, err := http.Get(att.URL)
        if err != nil {
            s.ChannelMessageSend(m.ChannelID, "Failed to download: "+att.Filename)
            continue
        }
        defer resp.Body.Close()

        _, err = s.ChannelMessageSendComplex(targetChannel.ID, &discordgo.MessageSend{
            Content: "Uploaded by " + m.Author.Mention(),
            Files: []*discordgo.File{
                {
                    Name:        att.Filename,
                    ContentType: att.ContentType,
                    Reader:      resp.Body,
                },
            },
        })
        if err != nil {
            s.ChannelMessageSend(m.ChannelID, "Failed to upload: "+att.Filename)
        }
    }
    s.ChannelMessageSend(m.ChannelID, "Uploaded to #"+foodName)
}