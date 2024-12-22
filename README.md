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


## **Advanced Validations**

bserializer supports advanced validations, including multiple validations per field and predefined validation functions. This makes it easy to enforce complex rules on your data.

# Built-in Validation Functions
1. NotEmpty
Ensures the field is not empty
```bash
serializer.NotEmpty
```
2. Positive
Ensures the field is a positive number.
```bash
serializer.Positive
```

3. ValidPassword
Ensures the password meets the following criteria:

- At least 8 characters long.
- Contains at least one uppercase letter.
- Contains at least one lowercase letter.
- Contains at least one number.
- Contains at least one special character (e.g., !@#$%^&*()).

```bash
serializer.ValidPassword
```

Example: Multiple Validations with Password
This example demonstrates using multiple validations, including the password validation:

```bash
package main

import (
	"fmt"
	"github.com/alun-dra/bserializer/serializer"
)

func main() {
	user := map[string]interface{}{
		"name":     "Alice Doe",
		"email":    "alice.doe@example.com",
		"password": "Short1!", // Invalid password for testing
	}

	// Create the serializer instance with validations
	s := serializer.BaseSerializer{
		Validations: map[string][]func(interface{}) error{
			"name":     {serializer.NotEmpty},
			"email":    {serializer.ValidEmail},
			"password": {serializer.ValidPassword}, // Add password validation
		},
	}

	// Validate the data
	if err := s.Validate(user); err != nil {
		fmt.Println("Validation Error:", err)
		return
	}

	fmt.Println("Validation passed!")
}
```
# Output

1. If the password field is invalid:
```bash
Validation Error: validation failed for field 'password': password must be at least 8 characters long
```

2. If all validations pass:
```bash
Validation passed!
```

# How It Works
1. Multiple Validations:
You can define multiple validation functions for a single field by using a slice of validation functions:
```bash
Validations: map[string][]func(interface{}) error{
    "name":     {serializer.NotEmpty},
    "email":    {serializer.ValidEmail},
    "password": {serializer.ValidPassword},
}
```








































































## **Field Transformations**

1. Transforming Fields

Here’s how to define and use transformations in `BaseSerializer`:

```bash
package main

import (
	"fmt"
	"strings"

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

	// Create a serializer instance with transformations
	s := serializer.BaseSerializer{
		Fields: []string{"id", "name", "email"}, // Optional: Filter only the fields you want
		Transformations: map[string]func(interface{}) interface{}{
			"name": func(value interface{}) interface{} {
				return strings.ToUpper(value.(string)) // Convert name to uppercase
			},
			"email": func(value interface{}) interface{} {
				return strings.ReplaceAll(value.(string), "@example.com", "@mydomain.com") // Change email domain
			},
		},
	}

	// Serialize the data
	serializedData, err := s.Serialize(user)
	if err != nil {
		fmt.Println("Serialization Error:", err)
		return
	}

	fmt.Println("Serialized Data with Transformations:", serializedData)
}

```
Output
If executed with the provided user struct, the output will be:
```bash
Serialized Data with Transformations: map[id:1 name:ALICE DOE email:alice.doe@mydomain.com]
```

How It Works
1. Transformations Field:
You define the Transformations field as a map where:

- Key: The field name in the struct you want to transform.
- Value: A function (func(interface{}) interface{}) that takes the field's value and returns the transformed value.

2. Order of Operations:
- Transformations are applied before field filtering, meaning the modified values will appear in the final serialized map if the field is included.

3. Flexible Application:
- You can mix transformations with validations and field filtering to fully customize your serialization process.




Additional Use Cases
Anonymizing Data:
Hide sensitive information like emails or phone numbers.


```bash
"email": func(value interface{}) interface{} {
    return "hidden@example.com" // Replace with anonymized value
}

```
Formatting Dates:

Convert timestamps into a human-readable format.
```bash
"created_at": func(value interface{}) interface{} {
    return time.Now().Format("2006-01-02") // Format as YYYY-MM-DD
}
```
Conditional Modifications:
Adjust values based on business logic.
```bash
"status": func(value interface{}) interface{} {
    if value.(string) == "inactive" {
        return "archived"
    }
    return value
}
```


## **Contributions**

Contributions are welcome. If you find an issue or have a suggestion, please open an issueor submit an pull requeston GitHub.

## **License**
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.