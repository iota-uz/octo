package octoapi

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type EnvConfiguration struct {
	OctoShopId int32  `env:"OCTO_SHOP_ID"`
	OctoSecret string `env:"OCTO_SECRET"`
	NotifyUrl  string `env:"OCTO_NOTIFY_URL"`
}

// cleanInvisible removes all invisible and control characters from string.
func cleanInvisible(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsControl(r) && r != '\n' && r != '\t' {
			return -1
		}
		return r
	}, s)
}

func LoadEnvConfiguration() *EnvConfiguration {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalf("Error opening .env file: %v", err)
	}
	defer file.Close()

	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := cleanInvisible(strings.TrimSpace(scanner.Text()))
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := cleanInvisible(strings.TrimSpace(parts[0]))
		value := cleanInvisible(strings.TrimSpace(parts[1]))
		if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) && len(value) > 1 {
			value = value[1 : len(value)-1]
		}
		envMap[key] = value
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	config := &EnvConfiguration{}
	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		envKey := cleanInvisible(strings.TrimSpace(field.Tag.Get("env")))
		envVal, ok := envMap[envKey]
		if ok {
			fieldVal := val.Field(i)
			if !fieldVal.CanSet() {
				log.Fatalf("Cannot set field: %s", field.Name)
			}
			switch field.Type.Kind() {
			case reflect.String:
				fieldVal.SetString(envVal)
			case reflect.Int32:
				intVal, err := strconv.ParseInt(envVal, 10, 32)
				if err != nil {
					log.Fatalf("Failed to parse int32 for %s: %v", envKey, err)
				}
				fieldVal.SetInt(intVal)
			default:
				log.Fatalf("Unsupported kind: %s", field.Type.Kind())
			}
		}
	}
	return config
}
