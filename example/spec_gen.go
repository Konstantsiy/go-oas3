// This file is generated by github.com/mikekonan/go-oas3. DO NOT EDIT.

package example

import "net/http"

var spec = []byte("{\"components\":{\"schemas\":{\"CountryAlpha2\":{\"$ref\":\"https://raw.githubusercontent.com/mikekonan/go-types/main/swagger.yaml#/components/schemas/CountryAlpha2\"},\"CreateTransactionRequest\":{\"properties\":{\"Amount\":{\"exclusiveMinimum\":true,\"minimum\":0.009,\"type\":\"number\"},\"AmountCents\":{\"maximum\":100,\"type\":\"integer\"},\"CallbackURL\":{\"$ref\":\"#/components/schemas/URL\"},\"Country\":{\"$ref\":\"#/components/schemas/CountryAlpha2\"},\"Currency\":{\"$ref\":\"#/components/schemas/CurrencyCode\"},\"Description\":{\"maxLength\":100,\"minLength\":8,\"type\":\"string\"},\"Details\":{\"type\":\"string\",\"x-go-pointer\":true},\"Email\":{\"$ref\":\"#/components/schemas/Email\"},\"RegexParam\":{\"type\":\"string\",\"x-go-regex\":\"^[.?\\\\d]+$\"},\"Title\":{\"maxLength\":50,\"minLength\":8,\"type\":\"string\"},\"TransactionID\":{\"format\":\"uuid\",\"type\":\"string\"}},\"required\":[\"Description\"]},\"CurrencyCode\":{\"$ref\":\"https://raw.githubusercontent.com/mikekonan/go-types/main/swagger.yaml#/components/schemas/CurrencyCode\"},\"Email\":{\"$ref\":\"https://raw.githubusercontent.com/mikekonan/go-types/main/swagger.yaml#/components/schemas/Email\"},\"GenericResponse\":{\"properties\":{\"result\":{\"enum\":[\"success\",\"failed\"],\"example\":\"success\",\"type\":\"string\"}}},\"RawPayload\":{\"format\":\"binary\",\"type\":\"string\"},\"URL\":{\"$ref\":\"https://raw.githubusercontent.com/mikekonan/go-types/main/swagger.yaml#/components/schemas/URL\"},\"WithEnum\":{\"enum\":[\"one\",\"two\"],\"type\":\"string\"}}},\"info\":{\"description\":\"go-oas3 example\",\"title\":\"go-oas3 example\",\"version\":\"1.0.0\"},\"openapi\":\"3.0.0\",\"paths\":{\"/callbacks/{callbackType}\":{\"post\":{\"description\":\"callbacks\",\"parameters\":[{\"description\":\"callback type\",\"in\":\"path\",\"name\":\"callbackType\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"requestBody\":{\"content\":{\"application/octet-stream\":{\"schema\":{\"$ref\":\"#/components/schemas/RawPayload\"}}},\"description\":\"Callback data\",\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/octet-stream\":{\"schema\":{\"$ref\":\"#/components/schemas/RawPayload\"}}},\"description\":\"OK\"}},\"summary\":\"Callback\",\"tags\":[\"callbacks\"]}},\"/transaction\":{\"post\":{\"parameters\":[{\"in\":\"header\",\"name\":\"x-signature\",\"schema\":{\"maxLength\":5,\"type\":\"string\"}}],\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/CreateTransactionRequest\"}}},\"description\":\"address to observe\"},\"responses\":{\"201\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/GenericResponse\"}}},\"description\":\"transaction created\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/GenericResponse\"}}},\"description\":\"bad request\"},\"500\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/GenericResponse\"}}},\"description\":\"unhandled error\"}},\"tags\":[\"transactions\"]}},\"/transactions/{uuid}\":{\"delete\":{\"parameters\":[{\"in\":\"header\",\"name\":\"x-signature\",\"schema\":{\"maxLength\":5,\"type\":\"string\"}},{\"description\":\"uuid\",\"in\":\"path\",\"name\":\"uuid\",\"required\":true,\"schema\":{\"type\":\"string\"},\"x-go-type\":\"github.com/satori/go.uuid.UUID\",\"x-go-type-string-parse\":\"github.com/satori/go.uuid.FromString\"},{\"in\":\"path\",\"name\":\"regexParam\",\"required\":true,\"schema\":{\"minLength\":5,\"type\":\"string\",\"x-go-regex\":\"^[.?\\\\d]+$\"}}],\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/GenericResponse\"}}},\"description\":\"transaction deleted\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/GenericResponse\"}}},\"description\":\"bad request\"}},\"tags\":[\"transactions\"]}}},\"servers\":[{\"url\":\"https://example.com\"}]}")

func Spec(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write(spec)
}
