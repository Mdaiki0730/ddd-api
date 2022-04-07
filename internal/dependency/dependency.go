package dependency

import (
	"api/internal/application/service"
	"api/internal/domain/repoif"
	"api/internal/infrastructure/repository"
)

// person dependency injection
func NewPersonApplicationService(r repoif.Person) service.PersonApplicationServiceIF {
	return service.NewPersonApplicationService(r)
}

func NewPersonRepository() repoif.Person {
	return repository.NewPerson()
}
