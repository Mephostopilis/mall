package mapping

import (
	"github.com/blevesearch/bleve/analysis/analyzer/standard"
	"github.com/blevesearch/bleve/mapping"
)

// NewBlockDocumentMapping
func NewBlockDocumentMapping() *mapping.DocumentMapping {
	nameFieldMapping := mapping.NewTextFieldMapping()
	nameFieldMapping.Name = "name"
	nameFieldMapping.Analyzer = standard.Name

	numMapping := mapping.NewNumericFieldMapping()

	beerMapping := mapping.NewDocumentStaticMapping()
	beerMapping.AddFieldMappingsAt("height", numMapping)
	beerMapping.AddFieldMappingsAt("hash", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("coinbase", nameFieldMapping)
	return beerMapping
}

// NewNormalTxDocumentMapping
func NewNormalTxDocumentMapping() *mapping.DocumentMapping {
	nameFieldMapping := mapping.NewTextFieldMapping()
	nameFieldMapping.Name = "name"
	nameFieldMapping.Analyzer = standard.Name

	// beerMapping := mapping.NewDocumentMapping()
	beerMapping := mapping.NewDocumentStaticMapping()
	beerMapping.AddFieldMappingsAt("hash", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("from_address", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("to_address", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("memo", nameFieldMapping)
	return beerMapping
}

func NewContractDeployTxDocumentMapping() *mapping.DocumentMapping {
	nameFieldMapping := mapping.NewTextFieldMapping()
	nameFieldMapping.Name = "name"
	nameFieldMapping.Analyzer = standard.Name

	beerMapping := mapping.NewDocumentStaticMapping()
	beerMapping.AddFieldMappingsAt("hash", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("from_address", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("to_address", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("memo", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("source", nameFieldMapping)
	return beerMapping
}

// NewContractInvokeTxDocumentMapping
func NewContractInvokeTxDocumentMapping() *mapping.DocumentMapping {
	nameFieldMapping := mapping.NewTextFieldMapping()
	nameFieldMapping.Name = "name"
	nameFieldMapping.Analyzer = standard.Name

	beerMapping := mapping.NewDocumentStaticMapping()
	beerMapping.AddFieldMappingsAt("hash", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("from_address", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("to_address", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("memo", nameFieldMapping)
	beerMapping.AddFieldMappingsAt("function", nameFieldMapping)
	return beerMapping
}

// NewNormalTxDocumentMapping
func NewAccountDocumentMapping() *mapping.DocumentMapping {
	nameFieldMapping := mapping.NewTextFieldMapping()
	nameFieldMapping.Name = "name"
	nameFieldMapping.Analyzer = standard.Name

	beerMapping := mapping.NewDocumentStaticMapping()
	beerMapping.AddFieldMappingsAt("address", nameFieldMapping)
	return beerMapping
}

// NewNormalTxDocumentMapping
func NewContractDocumentMapping() *mapping.DocumentMapping {
	nameFieldMapping := mapping.NewTextFieldMapping()
	nameFieldMapping.Name = "name"
	nameFieldMapping.Analyzer = standard.Name

	beerMapping := mapping.NewDocumentStaticMapping()
	beerMapping.AddFieldMappingsAt("address", nameFieldMapping)
	return beerMapping
}
