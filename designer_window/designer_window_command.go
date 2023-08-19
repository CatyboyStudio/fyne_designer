package designer_window

func (this *DesignerWindow) syncToggleValue() {
	this.toggleLeftItem.Checked = this.toggleLeft
	this.toggleRightItem.Checked = this.toggleRight
	this.toggleItem.Checked = this.toggle
	this.toggleRefresh()
}

func (this *DesignerWindow) toggleValue(p *bool) {
	*p = !*p
	this.syncToggleValue()
	this.build_Designer()
}

func (this *DesignerWindow) toggleView() {
	this.toggleValue(&this.toggle)
}

func (this *DesignerWindow) toggleLeftPanel() {
	this.toggleValue(&this.toggleLeft)
}

func (this *DesignerWindow) toggleRightPanel() {
	this.toggleValue(&this.toggleRight)
}
