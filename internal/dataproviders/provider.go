package dataproviders

import (
	"github.com/google/wire"
)

// Container is the container for the dataproviders domain
var DataProvidersSuperSet = wire.NewSet(
	NewTodoRepository,
)
