package services

import (
	"go-webapi-template/internal/domain/errors"
	"go-webapi-template/internal/domain/repos"
	"os"
)

type HelloWorldService struct {
	helloWorldRepository repos.IHelloWorldRepository
}

func NewHelloWorldService(helloWorldRepository repos.IHelloWorldRepository) *HelloWorldService {
	return &HelloWorldService{
		helloWorldRepository: helloWorldRepository,
	}
}

func (h *HelloWorldService) HelloWorld() (string, error) {
	return h.helloWorldRepository.HelloWorld(), nil
}

func (h *HelloWorldService) DoSomething() error {
	_, err := os.ReadFile("notfound.txt")

	if err != nil {
		return errors.HelloWorldError.New(err).WithMessage("DoSomething error")
	}

	return err
}
