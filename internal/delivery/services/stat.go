package services

import (
	"URLProject/internal/repository"
	"URLProject/pkg/event"
	"log"
)

type StatServiceDeps struct {
	EventBus *event.EventBus
	*repository.StatRepository
}

type StatService struct {
	EventBus *event.EventBus
	*repository.StatRepository
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (ss *StatService) AddClick() {
	for msg := range ss.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Printf("Bad EventLinkVisited data: %s\n", msg.Data)
				continue
			}
			ss.StatRepository.AddClick(id)
		}
	}
}
