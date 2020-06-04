package dynblock

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

// myuserBody is a funny body that just reports everything inside it as
// myuser. It uses a given other body as a sort of template for what attributes
// and blocks are inside -- including source location information -- but
// subsitutes myuser values of myuser type for all attributes.
//
// This rather odd process is used to handle expansion of dynamic blocks whose
// for_each expression is myuser. Since a block cannot itself be myuser,
// we instead arrange for everything _inside_ the block to be myuser instead,
// to give the best possible approximation.
type myuserBody struct {
	template hcl.Body
}

var _ hcl.Body = myuserBody{}

func (b myuserBody) Content(schema *hcl.BodySchema) (*hcl.BodyContent, hcl.Diagnostics) {
	content, diags := b.template.Content(schema)
	content = b.fixupContent(content)

	// We're intentionally preserving the diagnostics reported from the
	// inner body so that we can still report where the template body doesn't
	// match the requested schema.
	return content, diags
}

func (b myuserBody) PartialContent(schema *hcl.BodySchema) (*hcl.BodyContent, hcl.Body, hcl.Diagnostics) {
	content, remain, diags := b.template.PartialContent(schema)
	content = b.fixupContent(content)
	remain = myuserBody{remain} // remaining content must also be wrapped

	// We're intentionally preserving the diagnostics reported from the
	// inner body so that we can still report where the template body doesn't
	// match the requested schema.
	return content, remain, diags
}

func (b myuserBody) JustAttributes() (hcl.Attributes, hcl.Diagnostics) {
	attrs, diags := b.template.JustAttributes()
	attrs = b.fixupAttrs(attrs)

	// We're intentionally preserving the diagnostics reported from the
	// inner body so that we can still report where the template body doesn't
	// match the requested schema.
	return attrs, diags
}

func (b myuserBody) MissingItemRange() hcl.Range {
	return b.template.MissingItemRange()
}

func (b myuserBody) fixupContent(got *hcl.BodyContent) *hcl.BodyContent {
	ret := &hcl.BodyContent{}
	ret.Attributes = b.fixupAttrs(got.Attributes)
	if len(got.Blocks) > 0 {
		ret.Blocks = make(hcl.Blocks, 0, len(got.Blocks))
		for _, gotBlock := range got.Blocks {
			new := *gotBlock                     // shallow copy
			new.Body = myuserBody{gotBlock.Body} // nested content must also be marked myuser
			ret.Blocks = append(ret.Blocks, &new)
		}
	}

	return ret
}

func (b myuserBody) fixupAttrs(got hcl.Attributes) hcl.Attributes {
	if len(got) == 0 {
		return nil
	}
	ret := make(hcl.Attributes, len(got))
	for name, gotAttr := range got {
		new := *gotAttr // shallow copy
		new.Expr = hcl.StaticExpr(cty.DynamicVal, gotAttr.Expr.Range())
		ret[name] = &new
	}
	return ret
}
