package utils

import (
	"log"
)

// RemoverAluno remove um aluno do banco de dados pela matrícula
func RemoverAluno(matricula string) error {
	query := `DELETE FROM alunos WHERE matricula = $1`
	_, err := DB.Exec(query, matricula)
	if err != nil {
		log.Printf("Erro ao remover aluno do banco de dados: %v", err)
		return err
	}
	log.Println("Aluno removido com sucesso!")
	return nil
}
