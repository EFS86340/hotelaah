package hotelaah

import ()

type AahServer struct {
	UnimplementedHaalServer

	redisdb Redisor

	cachedCities []*City
}

func NewAahServer() (*AahServer, error) {
		return &AahServer{
cachedCities: &{

							}
		}

}

func (s *AahServer) Init() error {

}

func (s *AahServer) PullCities(preq *PullRequest) (*PullResponse, error) {

}

func (s *AahServer) PushProvCities(pcr *ProvCitiesRequest) (*ProvCitiesResponse, error) {

}

func (s *AahServer) DelProvCities(pcr *ProvCitiesRequest) (*ProvCitiesResponse, error) {

}
