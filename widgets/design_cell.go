package widgets

import (
	"goapp_fyne"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type DesignCellConfigData struct {
	// SizeFillColor      color.Color
	SizeBorderColor color.Color
	// MinSizeFillColor   color.Color
	MinSizeBorderColor color.Color
	BorderWidth        float32
}

var DesignCellConfig = DesignCellConfigData{}

func (this DesignCellConfigData) getColor(p color.Color, prefn string, defv func() color.Color) color.Color {
	if p == nil {
		s := fyne.CurrentApp().Preferences().String(prefn)
		c, ok := goapp_fyne.ParseColor(s)
		if !ok {
			c = defv()
		}
		return c
	}
	return p
}

func (this DesignCellConfigData) GetSizeFillColor() color.Color {
	return color.Alpha16{0}
}

func (this DesignCellConfigData) GetSizeBorderColor() color.Color {
	this.SizeBorderColor = this.getColor(this.SizeBorderColor, "DesignCellSizeBorderColor", func() color.Color {
		return goapp_fyne.StrToColor("blue")
	})
	return this.SizeBorderColor
}

func (this DesignCellConfigData) GetMinSizeFillColor() color.Color {
	return color.Alpha16{0}
}

func (this DesignCellConfigData) GetMinSizeBorderColor() color.Color {
	this.MinSizeBorderColor = this.getColor(this.SizeBorderColor, "DesignCellSizeBorderColor", func() color.Color {
		return goapp_fyne.StrToColor("red")
	})
	return this.MinSizeBorderColor
}

func (this DesignCellConfigData) GetBorderWidth() float32 {
	v := this.BorderWidth
	if v <= 0 {
		v = float32(fyne.CurrentApp().Preferences().Float("DesignCellBorderWidth"))
		if v <= 0 {
			v = 1
		}
		this.BorderWidth = v
	}
	return v
}

var _ fyne.Widget = (*DesignCellWidget)(nil)

type DesignCellWidget struct {
	widget.BaseWidget

	Id       int
	Active   bool
	Content  fyne.CanvasObject
	OnTapped func()
}

func NewDesignCellWidget(id int, content fyne.CanvasObject, trapped func()) *DesignCellWidget {
	o := &DesignCellWidget{
		Id:       id,
		Active:   false,
		Content:  content,
		OnTapped: trapped,
	}
	o.ExtendBaseWidget(o)
	return o
}

func (this *DesignCellWidget) Tapped(*fyne.PointEvent) {
	if this.OnTapped != nil {
		this.OnTapped()
	}
}

func (this *DesignCellWidget) IsActive() bool {
	return this.Active
}

// CreateRenderer implements fyne.Widget.
func (this *DesignCellWidget) CreateRenderer() fyne.WidgetRenderer {
	this.ExtendBaseWidget(this)

	rcSize := canvas.NewRectangle(DesignCellConfig.GetSizeFillColor())
	rcMinSize := canvas.NewRectangle(DesignCellConfig.GetMinSizeFillColor())

	r := &designCellWidgetRenderer{
		self:        this,
		rectSize:    rcSize,
		rectMinSize: rcMinSize,
	}
	r.syncData()
	return r
}

type designCellWidgetRenderer struct {
	goapp_fyne.BaseRenderer
	self        *DesignCellWidget
	rectSize    *canvas.Rectangle
	rectMinSize *canvas.Rectangle
}

// Layout the components of the button widget
func (this *designCellWidgetRenderer) Layout(size fyne.Size) {
	sz := this.self.Content.Size()
	this.rectSize.Move(fyne.NewPos(0, 0))
	this.rectSize.Resize(sz)

	msz := this.self.Content.MinSize()
	this.rectMinSize.Move(fyne.NewPos(0, 0))
	this.rectMinSize.Resize(msz)

	this.self.Content.Move(fyne.NewPos(0, 0))
	this.self.Content.Resize(size)
}

func (this *designCellWidgetRenderer) MinSize() (size fyne.Size) {
	return this.self.Content.MinSize()
}

func (this *designCellWidgetRenderer) Refresh() {
	this.syncData()
	for _, o := range this.Objects() {
		o.Refresh()
	}
}

// applyTheme updates this button to match the current theme
func (this *designCellWidgetRenderer) syncData() {
	this.rectSize.FillColor = DesignCellConfig.GetSizeFillColor()
	this.rectSize.StrokeColor = DesignCellConfig.GetSizeBorderColor()
	this.rectSize.StrokeWidth = DesignCellConfig.GetBorderWidth()
	this.rectMinSize.FillColor = DesignCellConfig.GetMinSizeFillColor()
	this.rectMinSize.StrokeColor = DesignCellConfig.GetMinSizeBorderColor()
	this.rectMinSize.StrokeWidth = DesignCellConfig.GetBorderWidth()

	sz := this.self.Content.Size()
	msz := this.self.Content.MinSize()

	if this.self.IsActive() {
		this.rectSize.Show()
		this.rectMinSize.Show()
		if msz.Width > sz.Width || msz.Height > sz.Height {
			this.rectSize.StrokeColor = goapp_fyne.ColorWithAlpha(this.rectSize.StrokeColor, 0.5)
			objects := []fyne.CanvasObject{
				this.self.Content,
				this.rectMinSize,
				this.rectSize,
			}
			this.SetObjects(objects)
		} else {
			this.rectMinSize.StrokeColor = goapp_fyne.ColorWithAlpha(this.rectMinSize.StrokeColor, 0.5)
			objects := []fyne.CanvasObject{
				this.self.Content,
				this.rectSize,
				this.rectMinSize,
			}
			this.SetObjects(objects)
		}
	} else {
		this.rectSize.Hide()
		this.rectMinSize.Hide()
		this.SetObjects([]fyne.CanvasObject{this.self.Content})
	}
}
