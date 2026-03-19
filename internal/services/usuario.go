package services

import (
	"errors"
	"fmt"
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
func (s *UsuarioService) GetUsuarios() ([]models.Usuario, error) {
	return s.repo.GetAll()
}

/**
* GET /usuarios/:id
*/
func (s *UsuarioService) GetByID(id int) (models.Usuario, error) {
	usuario, err := s.repo.GetByID(id)
	
	if err != nil {
		return models.Usuario{}, err  // propaga erro do repo (ex: conexão, query falha)
	}

	if usuario.ID == 0 {
		return models.Usuario{}, errors.New("usuário não encontrado")
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

	criado, err := s.repo.Create(input)
	if err != nil {
		// Você pode melhorar o erro com wrap (opcional, mas recomendado)
		return models.Usuario{}, fmt.Errorf("falha ao criar usuário: %w", err)
	}

	return criado, nil
}

/**
* PUT /usuarios
*/
func (s *UsuarioService) UpdateUsuario(id int, input models.Usuario) (models.Usuario, error) {

	if input.Nome == "" {
		return models.Usuario{}, errors.New("nome é obrigatório")
	}

	if input.Email == "" {
		return models.Usuario{}, errors.New("email é obrigatório")
	}

	criado, err := s.repo.Update(id, input)
	if err != nil {
		// Você pode melhorar o erro com wrap (opcional, mas recomendado)
		return models.Usuario{}, fmt.Errorf("falha ao criar usuário: %w", err)
	}

	return criado, nil
}

/**
* DELETE /usuarios/:id
*/
func (s *UsuarioService) DeleteUsuario(id int) (bool, error) {

	ok, err := s.repo.Delete(id)

	if err != nil {
		// Você pode melhorar o erro com wrap (opcional, mas recomendado)
		return false, fmt.Errorf("falha ao excluir usuário: %w", err)
	}
	if !ok {
		return false, errors.New("usuário não encontrado")
	}

	return true, nil
}