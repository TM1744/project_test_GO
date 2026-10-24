package model

import (
	"projeto-teste/enum"
	"gorm.io/gorm"
)

type IModel[Model any] interface {
	GetEntityType() enum.Entity // Retorna o tipo de entidade para deixar registrado nos registros de auditoria
	GetUniqueFields() string // Retorna os campos únicos de uma entidade, é utilizado para especificar os erros de unique constraint do BD
	ToString() string // Transcreve toda a entidade em String para ficar no registro de auditoria
	ValidateCreate() error // Realiza a validação de campos da entidade que dependem um do outro na criação
	ValidateUpdate(Model) error // A mesma coisa de validateCreate, com a diferença que é executada na atualização da entidade
	MergeUpdate(Model) // Realiza a junção da entidade antiga com a nova, serve para manter campos que não devem ser alterados
	OnCreate(*gorm.DB) error // Possibilita a realização de uma operação extra no banco de dados quando a entidade é criada
	OnUpdate(*gorm.DB, Model) error // Possibilita a realização de uma operação extra no banco de dados quando a entidade é atualizada
}
