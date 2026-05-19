package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
)

// RemoverHandler é responsável por remover um aluno do banco de dados
func RemoverHandler(response http.ResponseWriter, request *http.Request) {
	// Verifica se o método da requisição é POST
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém a matrícula enviada pelo formulário
	matricula := request.FormValue("matricula")

	// Remove o aluno do banco de dados
	err := utils.RemoverAluno(matricula)
	if err != nil {
		http.Error(response, "Erro ao remover o aluno", http.StatusInternalServerError)
		return
	}

	// Redireciona para a página inicial após o sucesso
	http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
