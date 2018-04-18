package shared

type User struct {
	ID			int		`json:"id"`
	Cpf			string 	`json:"cpf"`
	Name     	string 	`json:"name"`
	Email		string 	`json:"email"`
	Admin  		bool 	`json:"admin"`
	SensorID 	int 	`json:"sensorID"`
}

type Request struct {
	Method string `json:"method"`
	Action string `json:"action"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type Response struct {
	Success bool                   `json:"success"`
	Error   string                 `json:"error,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}