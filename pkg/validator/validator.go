package validator

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/lantonster/liberate/pkg/errors"
	"github.com/lantonster/liberate/pkg/errors/reason"
	"github.com/lantonster/liberate/pkg/log"
	"github.com/microcosm-cc/bluemonday"
)

type Checker interface {
	Check() (fields ValidationErrors, err error)
}

type Error struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ValidationErrors []*Error

func (v ValidationErrors) Error() string {
	buff := bytes.NewBufferString("")
	for _, fieldError := range v {
		buff.WriteString(fieldError.Error)
		buff.WriteString("\n")
	}
	return strings.TrimSpace(buff.String())
}

func Check(c context.Context, value any) (fields ValidationErrors, err error) {
	// 进行默认的结构体校验
	if err = defaultValidator.Struct(value); err != nil {
		var errs validator.ValidationErrors
		if !errors.As(err, &errs) {
			// 记录校验出错时的错误信息
			log.WithContext(c).Errorf("validate.Struct(%v) error: %v", value, err)
			return nil, fmt.Errorf("validate check exception: %w", err)
		}

		for _, field := range errs {
			fieldErr := &Error{
				Field: field.Field(),
				Error: field.Error(),
			}

			structNamespace := field.StructNamespace()
			if _, filedName, found := strings.Cut(structNamespace, "."); found {
				if originalTag := getObjectTagByFieldName(c, value, filedName); len(originalTag) > 0 {
					fieldErr.Field = originalTag
				}
			}

			fields = append(fields, fieldErr)
		}

		if len(fields) > 0 {
			return fields, errors.BadRequest(reason.RequestFormatError).WithMsg("%s", ValidationErrors(fields).Error())
		}
	}

	// 如果值实现了 Checker 接口，进行额外的校验
	if value, ok := value.(Checker); ok {
		if fields, err = value.Check(); err == nil {
			return nil, nil
		}

		return fields, errors.BadRequest(reason.RequestFormatError).WithMsg("%s", ValidationErrors(fields).Error())
	}

	return nil, nil
}

var defaultValidator = validator.New()

func init() {
	// 注册自定义的验证函数 "notblank"
	defaultValidator.RegisterValidation("notblank", notBlank)
	// 注册自定义的验证函数 "sanitizer"
	defaultValidator.RegisterValidation("sanitizer", sanitizer)
	// 注册用于获取字段标签名的函数
	defaultValidator.RegisterTagNameFunc(tagNameFunc)
}

// notBlank 检查字段是否为空
func notBlank(fl validator.FieldLevel) bool {
	// 获取要检查的字段
	field := fl.Field()

	// 根据字段的类型进行不同的检查
	switch field.Kind() {
	case reflect.String:
		// 去除字符串两端的空格
		trimSpace := strings.TrimSpace(field.String())
		// 如果字符串为空，将其设置为去除空格后的结果
		if len(trimSpace) == 0 {
			field.SetString(trimSpace)
		}
		return true
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		// 对于通道、映射、切片和数组，检查长度是否大于 0
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		// 对于指针、接口和函数，检查是否不为空
		return !field.IsNil()
	default:
		// 对于其他类型，检查是否有效且不等于零值
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

// sanitizer 清理字段中的 HTML 标签和特殊字符
func sanitizer(fl validator.FieldLevel) bool {
	// 获取要处理的字段
	field := fl.Field()

	// 根据字段的类型进行不同的操作
	switch field.Kind() {
	case reflect.String:
		// 使用 bluemonday 的 UGCPolicy 进行清理
		filter := bluemonday.UGCPolicy()
		// 清理并替换特定字符
		content := strings.Replace(filter.Sanitize(field.String()), "&amp;", "&", -1)
		// 设置清理后的结果到字段
		field.SetString(content)
		return true
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		// 对于通道、映射、切片和数组，检查长度是否大于 0
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		// 对于指针、接口和函数，检查是否不为空
		return !field.IsNil()
	default:
		// 对于其他类型，检查是否有效且不等于零值
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

// tagNameFunc 定义用于获取字段标签名的函数
func tagNameFunc(fld reflect.StructField) string {
	if jsonTag := fld.Tag.Get("json"); len(jsonTag) > 0 {
		if jsonTag == "-" {
			return ""
		}
		return jsonTag
	}
	if formTag := fld.Tag.Get("form"); len(formTag) > 0 {
		return formTag
	}
	return fld.Name
}

func getObjectTagByFieldName(c context.Context, obj any, fieldName string) (tag string) {
	defer func() {
		if err := recover(); err != nil {
			log.WithContext(c).Errorf("getObjectTagByFieldName(%v, %s) error: %v", obj, fieldName, err)
		}
	}()

	// 获取对象的类型
	objT := reflect.TypeOf(obj)
	// 获取指针指向的实际类型
	objT = objT.Elem()

	// 根据字段名查找字段，如果字段不存在，返回空字符串
	structField, exists := objT.FieldByName(fieldName)
	if !exists {
		return ""
	}

	// 尝试获取 "json" 标签的值，如果有且不为空则返回
	if tag = structField.Tag.Get("json"); len(tag) > 0 {
		return tag
	}
	// 否则返回 "form" 标签的值
	return structField.Tag.Get("form")
}
