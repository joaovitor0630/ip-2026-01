package utils

import (
	"log"
)

// BuscarAlunoPorMatricula busca um aluno no banco de dados pela matrícula
func BuscarAlunoPorMatricula(matricula string) (*Aluno, error) {
	query := `SELECT nome, email, curso, matricula FROM alunos WHERE matricula = $1`
	var aluno Aluno
	err := DB.QueryRow(query, matricula).Scan(&aluno.Nome, &aluno.Email, &aluno.Curso, &aluno.Matricula)
	if err != nil {
		log.Printf("Erro ao buscar aluno no banco de dados: %v", err)
		return nil, err
	}
	return &aluno, nil
}
