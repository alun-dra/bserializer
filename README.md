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
go get github.com/alun-dra/bserializer@v1.6.0
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

2. Custom Validation Functions:
You can easily add custom validation functions tailored to your specific needs.

3. Comprehensive Error Handling:
The Validate method returns detailed error messages indicating which field failed validation and why.

# Use Cases

Password Validation: Ensure passwords meet security requirements:
```bash
"password": {serializer.ValidPassword}
```

Email Validation: Validate email format:
```bash
"email": {serializer.ValidEmail}
```
Combined Rules: Apply multiple rules to a single field:
```bash
"name": {serializer.NotEmpty, serializer.MaxLength(50)}
```

# **Field Transformations**

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
## **Output**
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


# **Conditional Fields**


bserializer supports conditionally including or excluding fields during serialization based on custom logic. You can define a map of conditions for specific fields, and each condition is a function that evaluates whether a field should be included.

How It Works

1. ConditionalFields Field

You define a ConditionalFields map in the BaseSerializer:
Key: The field name to evaluate.
Value: A function (func(map[string]interface{}) bool) that takes the serialized data as input and returns:
true: The field will be included.
false: The field will be excluded.

2. Order of Operations:

Conditional field evaluation happens before field filtering (Fields) to ensure only necessary fields are included in the serialized result.

Example: Conditional Inclusion of Fields
In this example, the email field is included only if the user's role is "admin".

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
	Role  string `json:"role"`
}

func main() {
	user := User{
		ID:    1,
		Name:  "Alice Doe",
		Email: "alice.doe@example.com",
		Role:  "user", // Change to "admin" to include the email
	}

	// Create the serializer instance with conditional fields
	s := serializer.BaseSerializer{
		ConditionalFields: map[string]func(map[string]interface{}) bool{
			"email": func(data map[string]interface{}) bool {
				return data["role"] == "admin" // Include "email" only if role is "admin"
			},
		},
	}

	// Serialize the data
	serializedData, err := s.Serialize(user)
	if err != nil {
		fmt.Println("Serialization Error:", err)
		return
	}

	fmt.Println("Serialized Data with Conditional Fields:", serializedData)
}

```

## **Outpu**

1. If the role is "user" (not admin):
```bash
Serialized Data with Conditional Fields: map[id:1 name:Alice Doe role:user]
```

2. If the role is "admin":
```bash
Serialized Data with Conditional Fields: map[id:1 name:Alice Doe email:alice.doe@example.com role:admin]
```

## **Use Cases**

1. Restricting Data for Certain Roles:
Include fields only for specific roles, such as admin users.

```bash
"email": func(data map[string]interface{}) bool {
    return data["role"] == "admin"
}

```

2. Dynamic Field Visibility:
Show sensitive information only if a condition is met.
```bash
"ssn": func(data map[string]interface{}) bool {
    return data["is_verified"] == true
}

```
3. Custom Business Logic:
Exclude fields based on complex logic, such as subscription status.
```bash
"subscription_details": func(data map[string]interface{}) bool {
    return data["subscription_active"] == true
}

```

Adding Conditional Fields to BaseSerializer
Define the ConditionalFields map in your BaseSerializer instance:

```bash
s := serializer.BaseSerializer{
    ConditionalFields: map[string]func(map[string]interface{}) bool{
        "email": func(data map[string]interface{}) bool {
            return data["role"] == "admin" // Only include "email" for admins
        },
    },
}

```



# **Support for Other Formats: XML and YAM**

`bserializer` allows serializing data not only to JSON but also to XML and YAML formats. This feature is useful when working with APIs or systems that require these formats.

## **Serialize to XML**

To serialize a struct to XML, use the `SerializeToXML` method:



```bash
xmlOutput, err := serializer.SerializeToXML(data)
if err != nil {
    fmt.Println("XML Serialization Error:", err)
    return
}
fmt.Println("Serialized to XML:")
fmt.Println(xmlOutput)
```

## **Serialize to YAML**
To serialize a struct to YAML, use the SerializeToYAML method:
```bash
yamlOutput, err := serializer.SerializeToYAML(data)
if err != nil {
    fmt.Println("YAML Serialization Error:", err)
    return
}
fmt.Println("Serialized to YAML:")
fmt.Println(yamlOutput)
```

## **Example Usage** 
```bash
user := User{
    ID:    1,
    Name:  "Alice Doe",
    Email: "alice.doe@example.com",
    Role:  "admin",
}

s := serializer.BaseSerializer{}

// Serialize to XML
xmlOutput, err := s.SerializeToXML(user)
if err != nil {
    fmt.Println("XML Serialization Error:", err)
    return
}
fmt.Println("Serialized to XML:")
fmt.Println(xmlOutput)

// Serialize to YAML
yamlOutput, err := s.SerializeToYAML(user)
if err != nil {
    fmt.Println("YAML Serialization Error:", err)
    return
}
fmt.Println("Serialized to YAML:")
fmt.Println(yamlOutput)

```
## **Output**

XML:
```bash
<user>
  <id>1</id>
  <name>Alice Doe</name>
  <email>alice.doe@example.com</email>
  <role>admin</role>
</user>
```
YAML:
```bash
id: 1
name: Alice Doe
email: alice.doe@example.com
role: admin
```

## **How It Works**
1. XML Serialization:

Uses Go's encoding/xml package to convert structs to XML.
Automatically applies field tags (xml:"fieldname") for struct fields.

2. YAML Serialization:

Uses the gopkg.in/yaml.v3 package to serialize structs to YAML.
Requires installing the YAML library

```bash
go get gopkg.in/yaml.v3

```

# **Contributions**

Contributions are welcome. If you find an issue or have a suggestion, please open an issueor submit an pull requeston GitHub.

# **License**
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.