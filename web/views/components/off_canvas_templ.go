// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.898
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

type OffCanvasProps struct {
	ID     string
	Button templ.Component
}

type OCButtonVariant int

const (
	SUCCESS_BTN OCButtonVariant = iota
	ERROR_BTN
	WARNING_BTN
	DEFAULT_BTN
)

type OCButtonProps struct {
	Label    string
	Disabled bool
	Variant  OCButtonVariant
}

func OCButton(props OCButtonProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		switch props.Variant {
		case SUCCESS_BTN:
			if props.Disabled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<button type=\"button\" class=\"bg-success text-white text-xs rounded-sm px-3 py-1.5 cursor-pointer hover:opacity-75 disabled:bg-gray focus:border-none focus:outline-none\" x-on:click=\"open()\" disabled>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/components/off_canvas.templ`, Line: 34, Col: 18}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<button type=\"button\" class=\"bg-success text-white text-xs rounded-sm px-3 py-1.5 cursor-pointer hover:opacity-75 disabled:bg-gray focus:border-none focus:outline-none\" x-on:click=\"open()\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/components/off_canvas.templ`, Line: 40, Col: 18}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		case ERROR_BTN:
			if props.Disabled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<button type=\"button\" class=\"bg-error text-white text-xs rounded-sm px-3 py-1.5 cursor-pointer hover:opacity-75 disabled:bg-gray focus:border-none focus:outline-none\" x-on:click=\"open()\" disabled>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/components/off_canvas.templ`, Line: 49, Col: 18}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<button type=\"button\" class=\"bg-error text-white text-xs rounded-sm px-3 py-1.5 cursor-pointer hover:opacity-75 disabled:bg-gray focus:border-none focus:outline-none\" x-on:click=\"open()\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/components/off_canvas.templ`, Line: 55, Col: 18}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		case WARNING_BTN:
			if props.Disabled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<button type=\"button\" class=\"bg-warning text-white text-xs rounded-sm px-3 py-1.5 cursor-pointer hover:opacity-75 disabled:bg-gray focus:border-none focus:outline-none\" x-on:click=\"open()\" disabled>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/components/off_canvas.templ`, Line: 64, Col: 18}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<button type=\"button\" class=\"bg-warning text-white text-xs rounded-sm px-3 py-1.5 cursor-pointer hover:opacity-75 disabled:bg-gray focus:border-none focus:outline-none\" x-on:click=\"open()\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/components/off_canvas.templ`, Line: 70, Col: 18}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		case DEFAULT_BTN:
			if props.Disabled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "<button type=\"button\" class=\"bg-primary text-white text-xs rounded-sm px-3 py-1.5 cursor-pointer hover:opacity-75 disabled:bg-gray focus:border-none focus:outline-none\" x-on:click=\"open()\" disabled>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/components/off_canvas.templ`, Line: 79, Col: 18}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "<button type=\"button\" class=\"bg-primary text-white text-xs rounded-sm px-3 py-1.5 cursor-pointer hover:opacity-75 disabled:bg-gray focus:border-none focus:outline-none\" x-on:click=\"open()\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 string
				templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/components/off_canvas.templ`, Line: 85, Col: 18}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		return nil
	})
}

func OffCanvas(props OffCanvasProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "<div><div x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("offCanvas('%s')", props.ID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/views/components/off_canvas.templ`, Line: 92, Col: 56}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "\" x-init=\"init()\" x-effect=\"syncStateWithURL()\"><!-- offCanvas Button -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = props.Button.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "<!-- offCanvas Backdrop --><div x-cloak x-show=\"offCanvasOpen\" x-on:keydown.esc.window=\"close()\" x-on:click.self=\"close()\" role=\"dialog\" class=\"fixed inset-0 z-90 bg-gray/50\"><!-- offCanvas Sidebar --><div x-cloak x-show=\"offCanvasOpen\" x-transition:enter=\"transition ease-out duration-200\" x-transition:enter-start=\"translate-x-full\" x-transition:enter-end=\"translate-x-0\" x-transition:leave=\"transition ease-in duration-200\" x-transition:leave-start=\"translate-x-0\" x-transition:leave-end=\"translate-x-full\" role=\"document\" class=\"fixed top-0 right-0 flex h-full w-lg flex-col overflow-y-scroll bg-white px-4 pt-4 pb-14 outline-1 outline-gray\"><!-- Content -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var10.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "</div></div></div><script>\n            (function() {\n                 document.addEventListener(\"alpine:init\", () => {\n                     Alpine.data(\"offCanvas\", (id) => ({\n                         offCanvasOpen: false,\n                         init() {\n                             const params = new URLSearchParams(window.location.search);\n                             this.offCanvasOpen = params.get('offCanvas') === id;\n                         },\n                         open() {\n                             this.offCanvasOpen = true;\n                             const params = new URLSearchParams(window.location.search);\n                             params.set('offCanvas', id);\n                             history.replaceState({}, '', `${window.location.pathname}?${params}`);\n                         },\n                         close() {\n                             this.offCanvasOpen = false;\n                             const params = new URLSearchParams(window.location.search);\n                             if (params.get('offCanvas') === id) {\n                                 params.delete('offCanvas');\n                                 history.replaceState({}, '', `${window.location.pathname}?${params}`);\n                             }\n                         },\n                         syncStateWithURL() {\n                             const params = new URLSearchParams(window.location.search);\n                             this.offCanvasOpen = params.get('offCanvas') === id;\n                         },\n                     }))\n                 })\n             })();\n        </script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
