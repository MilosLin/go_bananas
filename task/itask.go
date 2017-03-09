package task

type Itask interface {
	//run task
	Run(argu *string) error
}
