// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Base() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"a web based yt-dlp frontend\"><link rel=\"stylesheet\" href=\"/static/css/pico.min.css\" type=\"text/css\"><link rel=\"stylesheet\" href=\"/static/css/base.css\" type=\"text/css\"><script src=\"/static/js/htmx.min.js\"></script><title>yt-dlp gui</title></head><body><header><progress value=\"0\" max=\"100\" class=\"big-progress\"></progress><p>Downloading</p></header><main class=\".container\"><form action=\"/\" method=\"POST\"><label for=\"url\">Youtube Url:</label><fieldset id=\"url\" role=\"group\"><input type=\"url\" name=\"url\" placeholder=\"https://www.youtube.com/watch?v=KxGRhd_iWuE\" required> <button type=\"submit\">Download</button></fieldset></form><progress value=\"0\" max=\"100\"></progress><article class=\"options\"><section><h2>Quality selection</h2><label>Video quality <select name=\"video_quality\" aria-label=\"Select video quality\"><option selected>Best</option> <option>1080p</option> <option>720</option></select></label> <label>Audio quality <select name=\"video_quality\" aria-label=\"Select video quality\"><option selected>Best</option> <option>1080p</option> <option>720</option></select></label></section><section><h2>Playlist selection</h2>dropdown this <label><input type=\"checkbox\" name=\"is_plalist\"> Is this a playlist</label> <ins>Starts at 1</ins> <label>Start index <input class=\"index-input\" type=\"number\" name=\"playlist_start_index\" placeholder=\"0\"></label> <label>End index <input class=\"index-input\" type=\"number\" name=\"playlist_end_index\" placeholder=\"-1\"></label></section><section><h2>Parts download selection</h2>dropdown this <label><input type=\"checkbox\" name=\"is_plalist\"> Do you want to download this in parts</label> requires ffmpeg download wider than desired <label>Start time <input class=\"time-input\" type=\"number\" name=\"start_hour\" placeholder=\"0\"> <input class=\"time-input\" type=\"number\" name=\"start_minute\" placeholder=\"0\"> <input class=\"time-input\" type=\"number\" name=\"start_second\" placeholder=\"0\"></label> <label>End time <input class=\"time-input\" type=\"number\" name=\"end_hour\" placeholder=\"0\"> <input class=\"time-input\" type=\"number\" name=\"end_minute\" placeholder=\"0\"> <input class=\"time-input\" type=\"number\" name=\"end_second\" placeholder=\"0\"></label></section><section><h2>File name builder</h2><label>File name: <input type=\"text\" name=\"file_name\" placeholder=\"%(title)s.%(ext)s\"></label> <label>Options: <button>title</button> <button>extension</button> <button>title</button> <button>title</button> <button>title</button></label></section><section><h2>Extra video data section</h2><label><input type=\"checkbox\" name=\"embed_thumbnail\"> Embedd thumbnail</label> <label><input type=\"checkbox\" name=\"embed_metadata\"> Embed metadata</label> <label><input type=\"checkbox\" name=\"embed_subtitles\"> Embed subtitles</label></section><section><h2>Others</h2>Uses current browser <label><input type=\"checkbox\" name=\"browser_cookies\"> Use browser cookies</label></section><section><label>Extra options <input type=\"text\" name=\"other_options\"></label></section></article></main></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
