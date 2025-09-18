package validate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const maxFileSize = 10 * 1024 * 1024 // 10MB

// Estrutura genérica para validar arquivos
type FileUploadDTO struct {
	Buffer       []byte `validate:"required"`
	MimeType     string `validate:"required,oneof=image/png image/jpeg image/jpg"`
	Size         int    `validate:"lte=10485760"` // 10MB
	OriginalName string `validate:"required"`
}

// Função genérica para validar e parsear o corpo JSON
func ParseBody[T any](ctx *fiber.Ctx) (*T, error) {
	// Lê o corpo da requisição
	body := ctx.Body()

	// Cria uma instância do tipo genérico
	var parsedBody T

	// Decodifica o JSON no struct do tipo T
	if err := json.Unmarshal(body, &parsedBody); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	// Valida o struct
	validate := validator.New()
	if err := validate.Struct(&parsedBody); err != nil {
		return nil, fmt.Errorf("erro na validação do corpo: %w", err)
	}

	return &parsedBody, nil
}

// Função genérica para validar e mapear parâmetros
func ParseParams[T any](ctx *fiber.Ctx) (*T, error) {
	var parsedParams T
	paramStruct := reflect.ValueOf(&parsedParams).Elem()

	for i := 0; i < paramStruct.NumField(); i++ {
		field := paramStruct.Type().Field(i)
		paramName := field.Tag.Get("params")
		if paramName == "" {
			continue
		}

		paramValue := ctx.Params(paramName)
		if paramValue == "" {
			continue
		}

		fieldValue := paramStruct.Field(i)

		switch field.Type.Kind() {
		case reflect.String:
			fieldValue.SetString(paramValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intValue, err := strconv.Atoi(paramValue)
			if err != nil {
				return nil, fmt.Errorf("erro ao converter %s para inteiro: %w", paramName, err)
			}
			fieldValue.SetInt(int64(intValue))
		default:
			return nil, fmt.Errorf("tipo não suportado: %s", field.Type.Name())
		}
	}

	// Validação com go-playground/validator
	validate := validator.New()
	if err := validate.Struct(&parsedParams); err != nil {
		return nil, fmt.Errorf("erro na validação dos parâmetros: %w", err)
	}

	return &parsedParams, nil
}

// Função genérica para validar e mapear query params
func ParseQueryParams[T any](ctx *fiber.Ctx) (*T, error) {
	var parsedParams T
	paramStruct := reflect.ValueOf(&parsedParams).Elem()

	for i := 0; i < paramStruct.NumField(); i++ {
		field := paramStruct.Type().Field(i)
		paramName := field.Tag.Get("query")
		if paramName == "" {
			continue
		}

		paramValue := ctx.Query(paramName)
		if paramValue == "" {
			continue
		}

		fieldValue := paramStruct.Field(i)

		switch field.Type.Kind() {
		case reflect.String:
			fieldValue.SetString(paramValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intValue, err := strconv.Atoi(paramValue)
			if err != nil {
				return nil, fmt.Errorf("erro ao converter %s para inteiro: %w", paramName, err)
			}
			fieldValue.SetInt(int64(intValue))
		default:
			return nil, fmt.Errorf("tipo não suportado: %s", field.Type.Name())
		}
	}

	// Validação com go-playground/validator
	validate := validator.New()
	if err := validate.Struct(&parsedParams); err != nil {
		return nil, fmt.Errorf("erro na validação dos parâmetros: %w", err)
	}

	return &parsedParams, nil
}

// Função genérica para validar arquivos enviados por multipart/form-data
func ParseFile(ctx *fiber.Ctx) (*FileUploadDTO, error) {
	// Obter o arquivo enviado
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("nenhum arquivo enviado: %w", err)
	}

	// Verificar tamanho do arquivo
	if fileHeader.Size > maxFileSize {
		return nil, fmt.Errorf("o tamanho do arquivo excede o limite de 10MB")
	}

	// Abrir o arquivo para leitura
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	// Ler o arquivo em um buffer
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(file); err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo: %w", err)
	}

	// Detectar o tipo MIME com base no conteúdo
	mimeType := http.DetectContentType(buf.Bytes())

	// Construir os dados do arquivo
	fileData := FileUploadDTO{
		Buffer:       buf.Bytes(),
		MimeType:     mimeType,
		Size:         int(fileHeader.Size),
		OriginalName: fileHeader.Filename,
	}

	// Validar os dados do arquivo
	validate := validator.New()
	if err := validate.Struct(&fileData); err != nil {
		return nil, fmt.Errorf("erro na validação do arquivo: %w", err)
	}

	return &fileData, nil
}

// Função para validar e-mail invalido
func UserEmail(email string) bool {
	var dominiosEmailsTemporarios = []string{
		"tuamaeaquelaursa.com",
		"temp-mail.org",
		"mailinator.com",
		"guerrillamail.com",
		"10minutemail.com",
		"tempemail.co",
		"fakeinbox.com",
		"throwawaymail.com",
		"e4ward.com",
		"nada.email",
		"yopmail.com",
		"dispostable.com",
		"trashmail.com",
		"getnada.com",
		"tempail.com",
		"mytemp.email",
		"temp-mail.io",
		"mohmal.com",
		"fakemail.net",
		"maildrop.cc",
		"spambog.com",
		"spamgourmet.com",
		"tmail.com",
		"instantemailaddress.com",
		"10minemail.com",
		"sharklasers.com",
		"mailnesia.com",
		"easytrashmail.com",
		"spamdecoy.net",
	}

	for _, dominio := range dominiosEmailsTemporarios {
		if strings.Contains(email, dominio) {
			return true
		}
	}

	return false
}

// Função para validar CPF (BR)
func ValidateCPF(cpf string) bool {
	re := regexp.MustCompile(`\D`)
	cpf = re.ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return false
	}

	// Verifica se todos os dígitos são iguais
	firstDigit := rune(cpf[0])
	allSame := true
	for _, digit := range cpf {
		if digit != firstDigit {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	calcResto := func(factor int) int {
		sum := 0
		for i := 0; i < factor-1; i++ {
			digit, _ := strconv.Atoi(string(cpf[i]))
			sum += digit * (factor - i)
		}
		resto := (sum * 10) % 11
		if resto == 10 {
			resto = 0
		}
		return resto
	}

	resto1 := calcResto(10)
	resto2 := calcResto(11)

	dv1, _ := strconv.Atoi(string(cpf[9]))
	dv2, _ := strconv.Atoi(string(cpf[10]))
	if resto1 != dv1 || resto2 != dv2 {
		return false
	}

	return true
}
