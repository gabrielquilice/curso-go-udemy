package matematica

//Arquivos de teste devem terminar com a expressão "_teste"
import "testing"

const erroPadrao = "Valor esperado %v, mas o valor encontrado foi %v."

func TestMedia(t *testing.T) { //Funções de teste devem começar com Test 
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado{
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
