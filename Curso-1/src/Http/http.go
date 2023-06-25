package http

import (
	"github.com/gin-gonic/gin"
)

type Aluno struct {
	name     string `json:"name"`
	age      int    `json:"age"`
	lastname string `json:"lastname"`
}

var aluno = []Aluno{
	{
		name:     "Diogo",
		age:      17,
		lastname: "Almeida",
	},
	{
		name:     "Mariana",
		age:      17,
		lastname: "Pereira",
	},
}

var constructAluno *Aluno

func main() {

	router := gin.Default()

	router.GET("/aluno", getAluno)

	router.Run("localhost:3030")

	// res, error := http.Get("https://google.com")
	// // if res.Body != nil {
	// // 	defer fmt.Println(res.Cookies()) // defer delay of execution
	// // }
	// // fmt.Println(res.StatusCode, error)

	// // constructAluno = new(Aluno)
	// // constructAluno.name = "Diogo"
	// // constructAluno.age = 17
	// // constructAluno.lastname = "Almeida"

	// // fmt.Println(constructAluno)
}

func getAluno(c *gin.Context) {
	c.indentedJSON(200, aluno)
}
