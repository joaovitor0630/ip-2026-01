package utils

import (
	"log"
)

// InserirAluno insere um novo aluno no banco de dados
func InserirAluno(nome, email, curso, matricula string) error {
	query := `INSERT INTO alunos (nome, email, curso, matricula) VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(query, nome, email, curso, matricula)
	if err != nil {
		log.Printf("Erro ao inserir aluno no banco de dados: %v", err)
		return err
	}
	log.Println("Aluno inserido com sucesso!")
	return nil
}
