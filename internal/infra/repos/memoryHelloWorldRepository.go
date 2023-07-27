package repos

import (
	"fmt"
	"go-webapi-template/internal/domain/repos"
	"time"
)

type HelloWorldRepository struct {
}

func NewHelloWorldRepository() repos.IHelloWorldRepository {
	return &HelloWorldRepository{}
}

func (h *HelloWorldRepository) HelloWorld() string {
	return fmt.Sprintf("Hello World from Memory! time:%s", time.Now().UTC().Format(time.RFC3339))
}
