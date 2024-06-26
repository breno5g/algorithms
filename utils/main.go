package utils

var (
	Array []int
)

func Init() {
	Array = ArrayNotSequential()
}

func GetArray() []int {
	return Array
}
