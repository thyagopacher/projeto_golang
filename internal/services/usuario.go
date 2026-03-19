package services

import (
	"errors"

	"projeto_go/internal/models"
	"projeto_go/internal/repositories"
)

type UsuarioService struct {
	repo *repositories.UsuarioRepository
}

func NewUsuarioService(repo *repositories.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo: repo}
}

/**
* GET /usuarios
*/
func (s *UsuarioService) GetUsuarios() []models.Usuario {
	return s.repo.GetAll()
}

/**
* GET /usuarios/:id
*/
func (s *UsuarioService) GetUsuarioByID(id int) (*models.Usuario, error) {
	usuario := s.repo.GetByID(id)

	if usuario == nil {
		return nil, errors.New("usuário não encontrado")
	}

	return usuario, nil
}

/**
* POST /usuarios
*/
func (s *UsuarioService) CreateUsuario(input models.Usuario) (models.Usuario, error) {

	if input.Nome == "" {
		return models.Usuario{}, errors.New("nome é obrigatório")
	}

	if input.Email == "" {
		return models.Usuario{}, errors.New("email é obrigatório")
	}

	return s.repo.Create(input), nil
}

/**
* DELETE /usuarios/:id
*/
func (s *UsuarioService) DeleteUsuario(id int) error {

	ok := s.repo.Delete(id)

	if !ok {
		return errors.New("usuário não encontrado")
	}

	return nil
}