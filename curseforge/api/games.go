// ----------------
// Games API @see https://docs.curseforge.com/?go#curseforge-core-api-games
// ----------------

package api

import (
	"github.com/lyssar/msdcli/curseforge/schemas"
	"github.com/lyssar/msdcli/curseforge/utils"
	"net/url"
	"strconv"
)

const (
	UriGame              schemas.ApiUri = "/v1/games/{gameId}"
	UriGames             schemas.ApiUri = "/v1/games"
	UriGamesVersions     schemas.ApiUri = "/v1/games/{gameId}/versions"
	UriGamesVersionTypes schemas.ApiUri = "/v1/games/{gameId}/version-types"
)

// GetGames @see https://docs.curseforge.com/?go#get-games
func (api CurseforgeApi) GetGames(index *int, pageSize *int) (response schemas.GetGamesResponse, err error) {
	q := url.Values{}

	if index != nil {
		q.Add("index", strconv.Itoa(*index))
	}

	if pageSize != nil {
		q.Add("pageSize", strconv.Itoa(*pageSize))
	}

	client := api.newCurseforgeClientForRoute(string(UriGames))
	client.Query(&q)
	client.Request()
	err = client.GetContent(&response)

	return
}

// GetGame @see https://docs.curseforge.com/?go#get-game
func (api CurseforgeApi) GetGame(gameId int) (game schemas.GetGameResponse, err error) {
	uri := utils.ReplaceNamed(string(UriGame), map[string]string{"gameId": strconv.Itoa(gameId)})

	client := api.newCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&game)

	return
}

// GetVersions @see https://docs.curseforge.com/?go#get-versions
func (api CurseforgeApi) GetVersions(gameId int) (response schemas.GetVersionsResponse, err error) {
	uri := utils.ReplaceNamed(string(UriGamesVersions), map[string]string{"gameId": strconv.Itoa(gameId)})
	client := api.newCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)
	return
}

// GetVersionTypes @see https://docs.curseforge.com/?go#get-version-types
func (api CurseforgeApi) GetVersionTypes(gameId int) (response schemas.GetVersionTypesResponse, err error) {
	uri := utils.ReplaceNamed(string(UriGamesVersionTypes), map[string]string{"gameId": strconv.Itoa(gameId)})
	client := api.newCurseforgeClientForRoute(uri)
	client.Request()
	err = client.GetContent(&response)
	return
}