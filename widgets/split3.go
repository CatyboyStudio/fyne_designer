package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.CanvasObject = (*Split3)(nil)

type Split3 struct {
	widget.BaseWidget
	OffsetL  float64
	OffsetT  float64
	Leading  fyne.CanvasObject
	Content  fyne.CanvasObject
	Trailing fyne.CanvasObject

	showLeading  bool
	showTrailing bool
}

func NewSplit3(leading, content, trailing fyne.CanvasObject) *Split3 {
	s := &Split3{
		OffsetL:  1.0 / 3,
		OffsetT:  1.0 / 3,
		Leading:  leading,
		Content:  content,
		Trailing: trailing,
	}
	s.BaseWidget.ExtendBaseWidget(s)
	return s
}

func (s *Split3) CreateRenderer() fyne.WidgetRenderer {
	s.BaseWidget.ExtendBaseWidget(s)
	dl := newDivider(s, false)
	dr := newDivider(s, true)
	return &Split3ContainerRenderer{
		Split3:   s,
		dividerL: dl,
		dividerT: dr,
		objects:  []fyne.CanvasObject{s.Leading, s.Content, s.Trailing, dl, dr},
	}
}

func (s *Split3) ExtendBaseWidget(wid fyne.Widget) {
	s.BaseWidget.ExtendBaseWidget(wid)
}

func (s *Split3) SetOffsetL(offset float64) {
	if s.OffsetL == offset {
		return
	}
	if offset > s.OffsetT+offset {
		offset = s.OffsetT
	}
	s.OffsetL = offset
	s.Refresh()
}

func (s *Split3) SetOffsetT(offset float64) {
	if s.OffsetT == offset {
		return
	}
	if offset < s.OffsetL {
		offset = s.OffsetL
	}
	s.OffsetT = offset
	s.Refresh()
}

func (s *Split3) SetOffset(offset float64, trailing bool) {
	if trailing {
		s.SetOffsetT(offset)
	} else {
		s.SetOffsetL(offset)
	}
}

func (s *Split3) GetOffset(trailing bool) float64 {
	if trailing {
		return s.OffsetT
	}
	return s.OffsetL
}

func (s *Split3) SetVisible(trailing bool, v bool) {
	if trailing {
		if s.showTrailing == v {
			return
		}
		s.showTrailing = v
	} else {
		if s.showLeading == v {
			return
		}
		s.showLeading = v
	}
	s.Refresh()
}

func (s *Split3) IsVisible(trailing bool) bool {
	if trailing {
		return s.showTrailing
	}
	return s.showLeading
}

var _ fyne.WidgetRenderer = (*Split3ContainerRenderer)(nil)

type Split3ContainerRenderer struct {
	Split3   *Split3
	dividerL *divider
	dividerT *divider
	objects  []fyne.CanvasObject
}

func (r *Split3ContainerRenderer) divider(trailing bool) *divider {
	if trailing {
		return r.dividerT
	}
	return r.dividerL
}

func (r *Split3ContainerRenderer) Destroy() {
}

func (r *Split3ContainerRenderer) Layout(size fyne.Size) {
	x := r.doLayout(size, false, 0)
	r.doLayout(size, true, x)
}

func (r *Split3ContainerRenderer) doLayout(size fyne.Size, trailing bool, cpos float32) float32 {
	visible := r.Split3.IsVisible(trailing)

	var dividerPos, leadingPos, trailingPos fyne.Position
	var dividerSize, leadingSize, trailingSize fyne.Size

	if true {
		lw, tw := r.computeSplit3Lengths(size.Width, r.minLeadingWidth(), r.minTrailingWidth(), trailing)
		if !visible {
			if trailing {
				lw = size.Width
				tw = 0
			} else {
				lw = 0
				tw = size.Width
			}
		}
		leadingPos.X = 0
		leadingSize.Width = lw
		leadingSize.Height = size.Height
		dividerPos.X = lw
		dividerSize.Width = dividerThickness()
		dividerSize.Height = size.Height
		trailingPos.X = lw + dividerSize.Width
		trailingSize.Width = tw
		trailingSize.Height = size.Height
	}

	r.divider(trailing).Move(dividerPos)
	r.divider(trailing).Resize(dividerSize)
	if trailing {
		pos := leadingPos
		pos.X = cpos
		sz := leadingSize
		sz.Width = sz.Width - cpos
		r.Split3.Content.Move(pos)
		r.Split3.Content.Resize(sz)
		if visible {
			r.Split3.Trailing.Show()
			r.divider(trailing).Show()
		} else {
			r.Split3.Trailing.Hide()
			r.divider(trailing).Hide()
		}
		r.Split3.Trailing.Move(trailingPos)
		r.Split3.Trailing.Resize(trailingSize)
	} else {
		if visible {
			r.Split3.Leading.Show()
			r.divider(trailing).Show()
		} else {
			r.Split3.Leading.Hide()
			r.divider(trailing).Hide()
		}
		r.Split3.Leading.Move(leadingPos)
		r.Split3.Leading.Resize(leadingSize)
		r.Split3.Content.Move(trailingPos)
		r.Split3.Content.Resize(trailingSize)
	}
	canvas.Refresh(r.divider(trailing))
	return trailingPos.X
}

func (r *Split3ContainerRenderer) MinSize() fyne.Size {
	s := fyne.NewSize(0, 0)
	for _, o := range r.objects {
		min := o.MinSize()
		if true {
			s.Width += min.Width
			s.Height = fyne.Max(s.Height, min.Height)
		}
	}
	return s
}

func (r *Split3ContainerRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *Split3ContainerRenderer) Refresh() {
	r.objects[0] = r.Split3.Leading
	r.objects[1] = r.Split3.Content
	r.objects[2] = r.Split3.Trailing
	// [3,4] is divider which doesn't change
	r.Layout(r.Split3.Size())
	canvas.Refresh(r.Split3)
}

func (r *Split3ContainerRenderer) computeSplit3Lengths(total, lMin, tMin float32, trailing bool) (float32, float32) {
	available := float64(total - dividerThickness())
	if available <= 0 {
		return 0, 0
	}
	ld := float64(lMin)
	tr := float64(tMin)
	offset := r.Split3.GetOffset(trailing)

	min := ld / available
	max := 1 - tr/available
	if min <= max {
		if offset < min {
			offset = min
		}
		if offset > max {
			offset = max
		}
	} else {
		offset = ld / (ld + tr)
	}

	ld = offset * available
	tr = available - ld
	return float32(ld), float32(tr)
}

func (r *Split3ContainerRenderer) minLeadingWidth() float32 {
	if r.Split3.Leading.Visible() {
		return r.Split3.Leading.MinSize().Width
	}
	return 0
}

func (r *Split3ContainerRenderer) minLeadingHeight() float32 {
	if r.Split3.Leading.Visible() {
		return r.Split3.Leading.MinSize().Height
	}
	return 0
}

func (r *Split3ContainerRenderer) minTrailingWidth() float32 {
	if r.Split3.Trailing.Visible() {
		return r.Split3.Trailing.MinSize().Width
	}
	return 0
}

func (r *Split3ContainerRenderer) minTrailingHeight() float32 {
	if r.Split3.Trailing.Visible() {
		return r.Split3.Trailing.MinSize().Height
	}
	return 0
}

// Declare conformity with interfaces
var _ fyne.CanvasObject = (*divider)(nil)
var _ fyne.Draggable = (*divider)(nil)
var _ desktop.Cursorable = (*divider)(nil)
var _ desktop.Hoverable = (*divider)(nil)

type divider struct {
	widget.BaseWidget
	Split3   *Split3
	trailing bool
	hovered  bool
}

func newDivider(Split3 *Split3, trailing bool) *divider {
	d := &divider{
		Split3:   Split3,
		trailing: trailing,
	}
	d.ExtendBaseWidget(d)
	return d
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (d *divider) CreateRenderer() fyne.WidgetRenderer {
	d.ExtendBaseWidget(d)
	background := canvas.NewRectangle(theme.ShadowColor())
	foreground := canvas.NewRectangle(theme.ForegroundColor())
	return &dividerRenderer{
		divider:    d,
		background: background,
		foreground: foreground,
		objects:    []fyne.CanvasObject{background, foreground},
	}
}

func (d *divider) Cursor() desktop.Cursor {
	return desktop.HResizeCursor
}

func (d *divider) DragEnd() {
}

func (d *divider) Dragged(event *fyne.DragEvent) {
	offset := d.Split3.GetOffset(d.trailing)
	if true {
		if leadingRatio := float64(d.Split3.Leading.Size().Width) / float64(d.Split3.Size().Width); offset < leadingRatio {
			offset = leadingRatio
		}
		if trailingRatio := 1. - (float64(d.Split3.Trailing.Size().Width) / float64(d.Split3.Size().Width)); offset > trailingRatio {
			offset = trailingRatio
		}
		offset += float64(event.Dragged.DX) / float64(d.Split3.Size().Width)
	}
	d.Split3.SetOffset(offset, d.trailing)
}

func (d *divider) MouseIn(event *desktop.MouseEvent) {
	d.hovered = true
	d.Split3.Refresh()
}

func (d *divider) MouseMoved(event *desktop.MouseEvent) {}

func (d *divider) MouseOut() {
	d.hovered = false
	d.Split3.Refresh()
}

var _ fyne.WidgetRenderer = (*dividerRenderer)(nil)

type dividerRenderer struct {
	divider    *divider
	background *canvas.Rectangle
	foreground *canvas.Rectangle
	objects    []fyne.CanvasObject
}

func (r *dividerRenderer) Destroy() {
}

func (r *dividerRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	var x, y, w, h float32
	if true {
		x = (dividerThickness() - handleThickness()) / 2
		y = (size.Height - handleLength()) / 2
		w = handleThickness()
		h = handleLength()
	}
	r.foreground.Move(fyne.NewPos(x, y))
	r.foreground.Resize(fyne.NewSize(w, h))
}

func (r *dividerRenderer) MinSize() fyne.Size {
	return fyne.NewSize(dividerThickness(), dividerLength())
}

func (r *dividerRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *dividerRenderer) Refresh() {
	if r.divider.hovered {
		r.background.FillColor = theme.HoverColor()
	} else {
		r.background.FillColor = theme.ShadowColor()
	}
	r.background.Refresh()
	r.foreground.FillColor = theme.ForegroundColor()
	r.foreground.Refresh()
	r.Layout(r.divider.Size())
}

func dividerThickness() float32 {
	return theme.Padding() * 2
}

func dividerLength() float32 {
	return theme.Padding() * 6
}

func handleThickness() float32 {
	return theme.Padding() / 2
}

func handleLength() float32 {
	return theme.Padding() * 4
}
