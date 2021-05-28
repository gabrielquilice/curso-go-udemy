package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //operação de bloqueio, atribuindo valor no canal
	fmt.Println("... isso só aparece depois que foi lido ...") //só vai ser executado quando for lido o dado no canal
}

func main(){
	c := make(chan int)

	go rotina(c)

	fmt.Println(<-c) //operação de bloqueio (só recebe o dado quando ele estiver pronto)
	fmt.Println("Agora foi lido!")
	fmt.Println(<-c)//deadlock (não há mais nada no canal, pois as goroutines já acabaram, então dará erro)

	 fmt.Println("Fim") //por conta do deadlock, essa linha nunca será executada, pois o programa vai ser encerrado
}
