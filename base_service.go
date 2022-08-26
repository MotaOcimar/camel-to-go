package cameltogo

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

type Status int

const (
	NEW Status = iota
	BUILT
	INITIALIZING
	INITIALIZED
	STARTING
	STARTED
	SUSPENDING
	SUSPENDED
	STOPPING
	STOPPED
	SHUTTING_DOWN
	SHUTDOWN
	FAILED
)

type IBaseService interface {
	DoBuild() error
	DoFail(error)

	GetStatus() Status
	SetStatus(Status)
	Lock()
	Unlock()
}

type BaseService struct {
	status Status
	mutex  sync.Mutex
}

func (baseService *BaseService) GetStatus() Status {
	return baseService.status
}

func (baseService *BaseService) SetStatus(status Status) {
	baseService.status = status
}

func (baseService *BaseService) Lock() {
	baseService.mutex.Lock()
}

func (baseService *BaseService) Unlock() {
	baseService.mutex.Unlock()
}

func Build(baseService IBaseService) {
	baseService.Lock()
	defer baseService.Unlock()

	if baseService.GetStatus() == NEW {
		log.Tracef("Building service: %+v", baseService)
		err := baseService.DoBuild()
		if err != nil {
			baseService.DoFail(err)
		}
		baseService.SetStatus(BUILT)
		log.Tracef("Built service: %+v", baseService)
	}
}