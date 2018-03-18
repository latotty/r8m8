package league

import (
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/slack"
)

type getLeaderboardInputAdapterSlack struct {
	slackService       slack.Service
	leagueSlackService SlackService
}

func (g *getLeaderboardInputAdapterSlack) Handle(data interface{}) (model.GetLeaderboardInput, error) {
	var input model.GetLeaderboardInput
	values := data.(string)

	requestValues, err := g.slackService.ParseRequestValues(values)
	if err != nil {
		return input, err
	}

	teamID := requestValues.TeamID
	teamDomain := requestValues.TeamDomain
	channelID := requestValues.ChannelID
	channelName := requestValues.ChannelName
	league := g.leagueSlackService.ToLeague(teamID, teamDomain, channelID, channelName)

	input = model.GetLeaderboardInput{
		League: league,
	}
	return input, nil
}

// NewGetLeaderboardInputAdapterSlack factory
func NewGetLeaderboardInputAdapterSlack(slackService slack.Service, leagueSlackService SlackService) GetLeaderboardInputAdapter {
	return &getLeaderboardInputAdapterSlack{
		slackService:       slackService,
		leagueSlackService: leagueSlackService,
	}
}
