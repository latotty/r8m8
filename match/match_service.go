package match

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/rating"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	Add(transaction transaction.Transaction, players []player.Player, league league.League, reporterPlayer player.Player) error
}

type matchService struct {
	matchRepository Repository
	ratingService   rating.Service
	playerService   player.Service
	leagueService   league.Service
}

func (m *matchService) Add(transaction transaction.Transaction, players []player.Player, league league.League, reporterPlayer player.Player) error {
	if isPlayerCountUneven(players) {
		return &errors.UnevenMatchPlayersError{}
	}

	if isReporterPlayerNotInLeague(reporterPlayer, players) {
		return &errors.ReporterPlayerNotInLeagueError{}
	}

	repoLeague, err := m.leagueService.GetOrAddLeague(transaction, league)
	if err != nil {
		return err
	}

	leagueID := repoLeague.ID
	repoPlayers, err := m.playerService.GetOrAddPlayers(transaction, players, leagueID)
	if err != nil {
		return err
	}

	reporterRepoPlayer := getReporterRepoPlayer(reporterPlayer, repoPlayers)
	reporterRepoPlayerID := reporterRepoPlayer.ID
	matchID, err := m.matchRepository.Create(transaction, leagueID, reporterRepoPlayerID)
	if err != nil {
		return err
	}

	repoPlayerIDs := mapToIDs(repoPlayers)
	err = m.ratingService.UpdateRatings(transaction, repoPlayerIDs, matchID)
	return err
}

func isPlayerCountUneven(players []player.Player) bool {
	return (len(players) % 2) != 0
}

func isReporterPlayerNotInLeague(reporterPlayer player.Player, players []player.Player) bool {
	missingFromLeague := true
	for i := range players {
		if players[i].UniqueName == reporterPlayer.UniqueName {
			missingFromLeague = false
		}
	}
	return missingFromLeague
}

func getReporterRepoPlayer(reporterPlayer player.Player, repoPlayers []player.RepoPlayer) player.RepoPlayer {
	var reporterRepoPlayer player.RepoPlayer

	for i := range repoPlayers {
		if repoPlayers[i].UniqueName == reporterPlayer.UniqueName {
			reporterRepoPlayer = repoPlayers[i]
		}
	}

	return reporterRepoPlayer
}

func mapToIDs(repoPlayers []player.RepoPlayer) []int64 {
	IDs := make([]int64, len(repoPlayers))
	for i := range repoPlayers {
		IDs[i] = repoPlayers[i].ID
	}
	return IDs
}

// NewService creates a service
func NewService(matchRepository Repository, ratingService rating.Service, playerService player.Service, leagueService league.Service) Service {
	return &matchService{
		matchRepository: matchRepository,
		ratingService:   ratingService,
		playerService:   playerService,
		leagueService:   leagueService,
	}
}
