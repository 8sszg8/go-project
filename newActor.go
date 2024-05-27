package poor_project

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"time"
)

type customer struct {
	Name string
}
type actorA struct{}

func (a *actorA) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *customer:
		fmt.Printf("Hello, %v\n", msg.Name)
	}
}

func NewActor() {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(
		func() actor.Actor {
			return &actorA{}
		})
	pid := system.Root.Spawn(props)
	fmt.Println(pid)
	system.Root.Send(pid, &customer{Name: "LiuPei"})
	time.Sleep(1)
}
