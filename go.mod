module github.com/theovidal/craftlink

go 1.16

require (
	github.com/bwmarrin/discordgo v0.23.2
	github.com/fatih/color v1.10.0
	github.com/gorilla/websocket v1.4.2
	github.com/theovidal/onyxcord v0.1.0
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/theovidal/onyxcord => ../onyxcord

replace github.com/bwmarrin/discordgo => github.com/FedorLap2006/discordgo v0.22.1-0.20210526221316-e7fb87fa3c1b
