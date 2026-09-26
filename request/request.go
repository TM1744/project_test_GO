package request

import (

)


type IRequest[M any] interface {
	Validate() error
	Instance()
	ToModel() M
}

