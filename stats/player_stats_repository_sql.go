package stats

import (
	"github.com/szokodiakos/r8m8/transaction"
)

type playerStatsRepositorySQL struct{}

func (s *playerStatsRepositorySQL) GetMultipleByLeagueUniqueName(tr transaction.Transaction, uniqueName string) ([]RepoPlayerStats, error) {
	repoPlayersStats := []RepoPlayerStats{}

	query := `
		SELECT
			p.display_name,
			r.rating,
			COUNT(CASE WHEN d.has_won THEN 1 END) AS won_match_count,
			COUNT(*) AS total_match_count
		FROM
			players p,
			ratings r,
			leagues l,
			matches m,
			details d
		WHERE
			l.unique_name = $1 AND
			l.id = r.league_id AND
			r.player_id = p.id AND
			p.id = d.player_id AND
			m.league_id = l.id AND
			d.match_id = m.id
		GROUP BY
			p.display_name,
			r.rating
		ORDER BY
			r.rating DESC;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&repoPlayersStats, query, uniqueName)

	return repoPlayersStats, err
}

// NewPlayerRepositorySQL factory
func NewPlayerRepositorySQL() PlayerRepository {
	return &playerStatsRepositorySQL{}
}