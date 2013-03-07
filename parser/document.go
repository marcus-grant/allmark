package parser

import (
	"github.com/andreaskoch/docs/indexer"
)

func ParseDocument(item *indexer.Item, lines []string) *indexer.Item {

	// meta data
	item, lines = ParseMetaData(item, lines)

	// title
	item.Title, lines = getMatchingValue(lines, TitlePattern)

	// description
	item.Description, lines = getMatchingValue(lines, DescriptionPattern)

	// content
	item.AddBlock("content", getContent(lines))

	return item
}
