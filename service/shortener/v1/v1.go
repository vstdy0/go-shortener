package shortener

import (
	"encoding/json"
	"fmt"
	"github.com/vstdy0/go-project/config"
	"github.com/vstdy0/go-project/service/shortener"
	"github.com/vstdy0/go-project/storage"
	infile "github.com/vstdy0/go-project/storage/infile"
	inmemory "github.com/vstdy0/go-project/storage/inmemory"
	"os"
	"strconv"
	"sync"
)

var _ shortener.URLService = (*Service)(nil)

type Service struct {
	id         int
	urlStorage storage.URLStorage
	mu         sync.Mutex
}

type Option func(*Service) error

func (s *Service) AddURL(url string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id, err := s.urlStorage.Set(strconv.Itoa(s.id+1), url)
	if err != nil {
		return "", err
	}
	s.id++
	return id, nil
}

func (s *Service) GetURL(id string) string {
	return s.urlStorage.Get(id)
}

func WithInMemoryStorage() Option {
	return func(srv *Service) error {
		var inMemory inmemory.InMemory
		inMemory.URLS = make(map[string]inmemory.URLModel)
		srv.urlStorage = &inMemory
		return nil
	}
}

func WithInFileStorage(cfg config.Config) Option {
	return func(srv *Service) error {
		var inFile infile.InFile
		file, err := os.OpenFile(cfg.FileStoragePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		inFile.Encoder = json.NewEncoder(file)
		if dec := json.NewDecoder(file); dec.More() {
			for dec.More() {
				err := dec.Decode(&inFile.URLS)
				if err != nil {
					return err
				}
				srv.id++
			}
		} else {
			inFile.URLS = make(map[string]infile.URLModel)
		}
		srv.urlStorage = &inFile
		return nil
	}
}

func NewService(opts ...Option) (*Service, error) {
	svc := &Service{}
	for _, opt := range opts {
		if err := opt(svc); err != nil {
			return nil, fmt.Errorf("initialising dependencies: %w", err)
		}
	}

	if svc.urlStorage == nil {
		return nil, fmt.Errorf("urlStorage: nil")
	}

	return svc, nil
}
