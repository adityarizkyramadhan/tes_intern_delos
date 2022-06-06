package domain

type CounterEndpoint struct {
	UniqueUserAgent int
	Count           int
	EndPoint        string
}

type TrackerUseCase interface {
	SearchUniqueUserAgent(string) (int, error)
	SearchEndpointCalled(string) (int, error)
	SaveTracker(uint, string) error
}
