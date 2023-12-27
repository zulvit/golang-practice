package main

import "fmt"

type FilterStrategy interface {
	Apply(image string) string
}

type BlackAndWhiteFilter struct{}

func (f *BlackAndWhiteFilter) Apply(image string) string {
	return fmt.Sprintf("Применён чёрно-белый фильтр к %s", image)
}

type HighContrastFilter struct{}

func (f *HighContrastFilter) Apply(image string) string {
	return fmt.Sprintf("Применён фильтр высокой контрастности к %s", image)
}

type ImageProcessor struct {
	filter FilterStrategy
}

func (p *ImageProcessor) SetFilter(filter FilterStrategy) {
	p.filter = filter
}

func (p *ImageProcessor) Process(image string) string {
	return p.filter.Apply(image)
}

func main() {
	processor := ImageProcessor{}

	bwFilter := &BlackAndWhiteFilter{}
	processor.SetFilter(bwFilter)
	fmt.Println(processor.Process("image.jpg"))

	hcFilter := &HighContrastFilter{}
	processor.SetFilter(hcFilter)
	fmt.Println(processor.Process("image.jpg"))
}
