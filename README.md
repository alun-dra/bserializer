# bserializer

`bserializer` is a Go package that facilitates serialization and deserialization of structures into maps (`map[string]interface{}`) with support for field filtering. Inspired by Django Rest Framework serializers, it is designed to be lightweight and easy to use.

## **Features**

- Serialization of structures into maps (`map[string]interface{}`).
- Deserialization of maps into structures.
- Optional filtering of fields on serialization.
- Customizable validations to ensure data integrity.
- Easy integration with REST APIs and HTML views.

---

## **Installation**

To install `bserializer`, use `go get`:

```bash
go get github.com/alun-dra/bserializer
```

## **Basic Use**

1. Import the package

Make sure to import bserializerinto your project:
```bash
import "github.com/alun-dra/bserializer/serializer"
```
2. Basic Serialization
```bash
package main

import (
	"fmt"
	"github.com/alun-dra/bserializer/serializer"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{
		ID:    1,
		Name:  "Alice Doe",
		Email: "alice.doe@example.com",
	}

	s := serializer.BaseSerializer{
		Fields: nil, 
	}

	serializedData, err := s.Serialize(user)
	if err != nil {
		fmt.Println("Error durante la serialización:", err)
		return
	}

	fmt.Println("Serialized Data:", serializedData)
}
```
Exit:
```bash
Serialized Data: map[id:1 name:Alice Doe email:alice.doe@example.com]
```

3. Serialization with Field Filtering
```bash
package main

import (
	"fmt"
	"github.com/alun-dra/bserializer/serializer"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{
		ID:    1,
		Name:  "Alice Doe",
		Email: "alice.doe@example.com",
	}

	s := serializer.BaseSerializer{
		Fields: []string{"id", "name"}, 
	}

	serializedData, err := s.Serialize(user)
	if err != nil {
		fmt.Println("Error durante la serialización:", err)
		return
	}

	fmt.Println("Serialized Data:", serializedData)
}
```
Exit:
```bash
Serialized Data: map[id:1 name:Alice Doe]
```

4. Deserialization
```bash
package main

import (
	"fmt"
	"github.com/alun-dra/bserializer/serializer"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	data := map[string]interface{}{
		"id":    1,
		"name":  "Alice Doe",
		"email": "alice.doe@example.com",
	}

	// Instancia del serializador
	s := serializer.BaseSerializer{}

	var user User

	err := s.Deserialize(data, &user)
	if err != nil {
		fmt.Println("Error durante la deserialización:", err)
		return
	}

	fmt.Println("Deserialized User:", user)
}
```
Exit:
```bash
Deserialized User: {1 Alice Doe alice.doe@example.com}
```

5. Use in REST APIs
Here is an example of how to use it bserializerin a REST API endpoint:
```bash
package main

import (
	"encoding/json"
	"net/http"
	"github.com/alun-dra/bserializer/serializer"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:    1,
		Name:  "Alice Doe",
		Email: "alice.doe@example.com",
	}

	s := serializer.BaseSerializer{
		Fields: []string{"id", "name"},
	}

	serializedData, _ := s.Serialize(user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serializedData)
}

func main() {
	http.HandleFunc("/api/user", userHandler)
	http.ListenAndServe(":8080", nil)
}
```

Access to "http://localhost:8080/api/user" (enpoit example):
```bash
{"id":1,"name":"Alice Doe"}
```

6. Use in HTML Views
You can pass serialized data directly to HTML templates in Go:
```bash
package main

import (
	"html/template"
	"net/http"
	"github.com/alun-dra/bserializer/serializer"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:    1,
		Name:  "Alice Doe",
		Email: "alice.doe@example.com",
	}

	s := serializer.BaseSerializer{
		Fields: []string{"id", "name"},
	}

	serializedData, _ := s.Serialize(user)

	tmpl := template.Must(template.New("view").Parse(`
		<!DOCTYPE html>
		<html>
		<head><title>User Info</title></head>
		<body>
			<h1>User Info</h1>
			<p>ID: {{.id}}</p>
			<p>Name: {{.name}}</p>
		</body>
		</html>
	`))

	tmpl.Execute(w, serializedData)
}

func main() {
	http.HandleFunc("/", userHandler)
	http.ListenAndServe(":8080", nil)
}
```

## **validation**

1. Validations
bserializer supports custom field validations to ensure data complies with specific rules. You can define validation rules for each field using validation functions.
```bash
package main

import (
	"fmt"
	"github.com/alun-dra/bserializer/serializer"
)

func main() {
	user := map[string]interface{}{
		"id":    1,
		"name":  "", // Empty field to test validation
		"email": "alice.doe@example.com",
	}

	// Create the serializer instance with validations
	s := serializer.BaseSerializer{
		Fields: []string{"id", "name", "email"},
		Validations: map[string]func(interface{}) error{
			"name":  serializer.NotEmpty, // Validate that "name" is not empty
			"email": serializer.NotEmpty, // Validate that "email" is not empty
		},
	}

	// Validate data
	if err := s.Validate(user); err != nil {
		fmt.Println("Validation Error:", err)
		return
	}

	// Serialize data
	serializedData, err := s.Serialize(user)
	if err != nil {
		fmt.Println("Serialization Error:", err)
		return
	}

	fmt.Println("Serialized Data:", serializedData)
}
```
Output: If the "name" field is empty:
```bash
Validation Error: validation failed for field 'name': value cannot be empty
```


## **Contributions**

Contributions are welcome. If you find an issue or have a suggestion, please open an issueor submit an pull requeston GitHub.

## **License**
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.