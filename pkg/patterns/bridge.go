package patterns

import (
	"fmt"
)

// Implementer defines the interface for implementation classes
type Color interface {
	Paint() string
}
// Red is a concrete implementer of the color interface
type Red struct{}
func (r *Red) Paint() string {
	return "red"
}
// Bule is another concrete implementer of the color interface
type Blue struct{}
func (b *Blue) Paint() string {
	return "blue"
}

// Shape is the abstraction that holds a reference to the Color implementer
type Shape struct {
	Color Color
}
func(s *Shape) Draw() string {
	return fmt.Sprintf("Shape drawn with %s color",s.Color.Paint())
}

// Circle is a refined abstraction of Shape 
type Circle struct {
	Shape
}
func (c *Circle) Draw() string {
	return fmt.Sprintf("Circle with %s color",c.Color.Paint())
}

// Square is refinde abstraction of Shape
type Square struct {
	Shape
}
func (s *Square) Draw() string {
	return fmt.Sprintf("Square with %s color",s.Color.Paint())
}

// Other example of bridge pattern

// Converter is the implementer interface for converting video formats
type Converter interface {
	Convert(videoFile string) string
}

// MP4Converter is a concrete implementer of the Converter interface
type AVIConverter struct {}
func (c *AVIConverter) Convert(videoFile string) string {
	return fmt.Sprintf("Converting %s to AVI format",videoFile)
}
// MP4Converter is another concrete implementer of the Converter interface
type MP4Converter struct {}
func (c *MP4Converter) Convert(videoFile string) string {
	return fmt.Sprintf("Converting %s to MP4 format",videoFile)
}
// MKVConverter is another concrete implementer of the Converter interface
type MKVConverter struct{}
func (c *MKVConverter) Convert(videoFile string) string {
	return fmt.Sprintf("Converting %s to KMV format",videoFile)
}

// Video is the abstraction that holds a referrance to a Converter
type Video struct {
	FileName string
	Converter Converter
}

func (v *Video) Convert() string {
	return v.Converter.Convert(v.FileName)
}

// MOVVideo is a refined abstraction of video
type MOVVideo struct{
	Video
}
// WMVVideo is refined abstraction of Video
type WMVVideo struct{
	Video
}


func CircleDrawer(){
	red := &Red{}
	circle := &Circle{Shape:Shape{Color:red}}
	fmt.Println(circle.Draw())
	// Create a blue square
	blue := &Blue{}
	square := &Square{Shape: Shape{Color: blue}}
	fmt.Println(square.Draw())
}

func VideoConverter(){
			// Conver MOV video to MP4 format
			mp4Converter := new(MP4Converter)
			movVideo := &MOVVideo{
				Video:Video{
					FileName:"example.mov",
					Converter:mp4Converter,
				},
			}
			fmt.Println(movVideo.Convert()) 
	
			// Convert WMV video to AVI format
			AVIConverter:=new(AVIConverter)
			wmvVideo := &WMVVideo{
				Video:Video{
					FileName:"example.wmv",
					Converter:AVIConverter,
				},
			}
			fmt.Println(wmvVideo.Convert())
}