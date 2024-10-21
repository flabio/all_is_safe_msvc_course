package helpers

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/safe_msvc_course/insfractruture/utils"
	"github.com/safe_msvc_course/usecase/dto"
)

var fieldCache sync.Map // Caché para almacenar los índices de los campos
func MapToStructCourse(courseDto *dto.CourseDTO, dataMap map[string]interface{}) {
	course := dto.CourseDTO{
		Id:     0,
		Name:   dataMap[utils.NAME].(string),
		Active: dataMap[utils.ACTIVE].(bool),
	}
	*courseDto = course
}

func MapToStructTopic(topicDto *dto.TopicDTO, dataMap map[string]interface{}) {
	topic := dto.TopicDTO{
		Id:        0,
		Title:     dataMap[utils.TITLE].(string),
		TimeHours: dataMap[utils.TIME_HOURS].(string),
		CourseId:  uint(dataMap[utils.COURSE_ID].(float64)),
		Active:    dataMap[utils.ACTIVE].(bool),
	}
	*topicDto = topic
}
func MapToStructLanguage(lenguageDto *dto.LanguageDTO, dataMap map[string]interface{}) {
	lenguage := dto.LanguageDTO{
		Id:     0,
		Name:   dataMap[utils.NAME].(string),
		Active: dataMap[utils.ACTIVE].(bool),
	}
	*lenguageDto = lenguage
}

// MapToStruct es una función genérica que mapea un map[string]interface{} a cualquier estructura proporcionada
func MapToStruct(dataMap map[string]interface{}, structPointer interface{}) error {
	// Verificamos que structPointer sea un puntero a una estructura
	value := reflect.ValueOf(structPointer)
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		return errors.New("structPointer debe ser un puntero a una estructura")
	}

	structValue := value.Elem()
	structType := structValue.Type()

	for key, mapValue := range dataMap {
		// Intentamos obtener el índice del campo desde el caché
		fieldIdx, found := getFieldIndex(structType, key)
		if !found {
			// Si no encontramos el campo, continuamos con la siguiente clave del map
			continue
		}

		fieldValue := structValue.Field(fieldIdx)
		if !fieldValue.CanSet() {
			// Si el campo no se puede asignar (no exportado), continuamos
			continue
		}

		mapValueReflect := reflect.ValueOf(mapValue)

		// Verificamos si el valor es asignable al tipo del campo
		if mapValueReflect.Type().AssignableTo(fieldValue.Type()) {
			fieldValue.Set(mapValueReflect)
		} else {
			// Intentamos convertir el valor si es posible
			convertedValue, err := tryConvert(mapValueReflect, fieldValue.Type())
			if err == nil {
				fieldValue.Set(convertedValue)
			} else {
				return fmt.Errorf("error asignando el campo '%s': %v", key, err)
			}
		}
	}

	return nil
}

// getFieldIndex intenta obtener el índice del campo desde el caché, o lo calcula y almacena si no está en el caché
func getFieldIndex(structType reflect.Type, fieldName string) (int, bool) {
	cacheKey := structType.String() + "." + fieldName

	if index, ok := fieldCache.Load(cacheKey); ok {
		return index.(int), true
	}
	/*
		field, found := structType.FieldByName(fieldName)
		if found {
			fieldCache.Store(cacheKey, field.Index[0])
			return field.Index[0], true
		}*/
	// Buscar el campo correspondiente a la etiqueta json
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		} else {
			tag = strings.Split(tag, ",")[0] // Obtenemos el nombre antes de cualquier opción (como omitir)
		}

		if tag == fieldName {
			fieldCache.Store(cacheKey, i)
			return i, true
		}
	}
	return -1, false
}

// tryConvert intenta convertir un valor a un tipo objetivo compatible
func tryConvert(value reflect.Value, targetType reflect.Type) (reflect.Value, error) {
	if value.Type().ConvertibleTo(targetType) {
		return value.Convert(targetType), nil
	}
	return reflect.Value{}, fmt.Errorf("no se puede convertir el valor '%v' al tipo '%v'", value, targetType)
}

/*
// Optimizada función genérica para mapear datos de un mapa a una estructura
func MapToStructOptimizada(structPointer interface{}, dataMap map[string]interface{}) error {
	// Verificamos que structPointer sea un puntero a una estructura
	value := reflect.ValueOf(structPointer)
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		return errors.New("structPointer debe ser un puntero a una estructura")
	}

	// Obtener el valor y tipo de la estructura
	structValue := value.Elem()
	structType := structValue.Type()
	log.Println(structType)
	// Iteramos a través del map y asignamos los valores a los campos correspondientes en la estructura
	for key, mapValue := range dataMap {
		// Buscamos si la estructura tiene un campo con el nombre correspondiente a la clave del map
		// Nota: Podríamos usar etiquetas "json" en vez de nombres de campo si lo deseas
		field := structValue.FieldByName(key)
		if !field.IsValid() {
			// Si no hay campo con este nombre en la estructura, lo ignoramos
			continue
		}

		// Comprobamos si el campo es asignable
		if !field.CanSet() {
			continue
		}

		// Obtenemos el valor del mapa y lo convertimos a reflect.Value
		fieldValue := reflect.ValueOf(mapValue)

		// Verificamos si el tipo es asignable
		if fieldValue.Type().AssignableTo(field.Type()) {
			field.Set(fieldValue)
		} else {
			// Intentamos hacer una conversión si los tipos son compatibles (por ejemplo, de int a float64)
			convertedValue, err := tryConvert(fieldValue, field.Type())
			if err == nil {
				field.Set(convertedValue)
			}
		}
	}
	return nil
}

// Intentar convertir el valor si es posible (por ejemplo, de int a float64)
func tryConvert(value reflect.Value, targetType reflect.Type) (reflect.Value, error) {
	if value.Type().ConvertibleTo(targetType) {
		return value.Convert(targetType), nil
	}
	return reflect.Value{}, errors.New("no se puede convertir el valor")
}
*/
