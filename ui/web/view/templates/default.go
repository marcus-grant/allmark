// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package templates

import (
	"fmt"
)

var masterTemplate = fmt.Sprintf(`<!DOCTYPE HTML>
<html lang="{{.LanguageTag}}">
<head>
	<meta charset="utf-8">
	<base href="{{ .BaseUrl }}">

	<title>{{.Title}}</title>

	<link rel="schema.DC" href="http://purl.org/dc/terms/">
	<meta name="DC.date" content="{{.CreationDate}}">

	{{if .GeoLocation }}
	{{if .GeoLocation.Coordinates}}
	<meta name="geo.position" content="{{.GeoLocation.Coordinates}}">
	{{end}}
	
	{{if .GeoLocation.PlaceName}}
	<meta name="geo.placename" content="{{.GeoLocation.PlaceName}}">	
	{{end}}
	{{end}}

	<link rel="alternate" type="application/rss+xml" title="RSS" href="/rss.xml">

	<link rel="shortcut icon" href="/theme/favicon.ico" />

	<link rel="stylesheet" href="/theme/deck.css" media="screen">
	<link rel="stylesheet" href="/theme/screen.css" media="screen">
	<link rel="stylesheet" href="/theme/print.css" media="print">
	<link rel="stylesheet" href="/theme/codehighlighting/highlight.css" media="screen, print">

	<script src="/theme/modernizr.js"></script>
</head>
<body>

{{ if .ToplevelNavigation}}
<nav class="toplevel">
	<ul>
	{{range .ToplevelNavigation.Entries}}
	<li>
		<a href="{{.Path}}">{{.Title}}</a>
	</li>
	{{end}}
	</ul>
</nav>
{{end}}

{{ if .BreadcrumbNavigation}}
<nav class="breadcrumb">
	<ul>
	{{range .BreadcrumbNavigation.Entries}}
	<li>
		<a href="{{.Path}}">{{.Title}}</a>
	</li>
	{{end}}
	</ul>
</nav>
{{end}}

<article class="{{.Type}} level-{{.Level}}">
%s
</article>

<footer>
	<nav>
		<ul>
			<li><a href="/tags.html">Tags</a></li>
			<li><a href="/sitemap.html">Sitemap</a></li>
			<li><a href="/rss.xml">RSS Feed</a></li>
		</ul>
	</nav>
</footer>

<script src="/theme/jquery.js"></script>
<script src="/theme/autoupdate.js"></script>
<script src="/theme/pdf.js"></script>
<script src="/theme/pdf-preview.js"></script>
<script src="/theme/codehighlighting/highlight.js"></script>
<script>hljs.initHighlightingOnLoad();</script>
<script src="/theme/deck.js"></script>
<script src="/theme/presentation.js"></script>

</body>
</html>`, ChildTemplatePlaceholder)

const repositoryTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>

<div class="cleaner"></div>

{{ if .TopDocs }}
<section class="childs">
<ol class="list">
{{range .TopDocs}}
<li class="child">
	<a href="{{.Route}}" class="child-title child-link">{{.Title}}</a>
	<p class="child-description">{{.Description}}</p>
	<div>{{.Content}}</div>
</li>
{{end}}
</ol>
</section>
{{end}}

{{if .TagCloud}}
<section class="tagcloud">
	<h1>Tag Cloud</h1>
	<div class="tags">
	{{range .TagCloud}}
	<span class="level-{{.Level}}">
		<a href="{{.Route}}">{{.Name}}</a>
	</span>
	{{end}}
	</div>
</section>
{{end}}
`

const documentTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>

<div class="cleaner"></div>

{{ if .Tags }}
<section class="tags">
	<header>
		Tags:
	</header>

	<ul class="tags">
	{{range .Tags}}
	<li class="tag">
		<a href="{{.Route}}" title="{{.Description}}">{{.Name}}</a>
	</li>
	{{end}}
	</ul>
</section>
{{end}}

{{ if .Locations }}
<section class="locations">
	<header>
		Locations:
	</header>

	<ol class="list">
	{{range .Locations}}
	<li class="location">
		<a href="{{.Route}}">{{.Title}}</a>
		{{ if .Description }}
		<p>{{.Description}}</p>
		{{end}}

		{{ if .GeoLocation }}

		{{ if .GeoLocation.Address }}
		<p class="address">{{ .GeoLocation.Address }}</p>
		{{end}}

		{{ if .GeoLocation.Coordinates }}
		<p class="geo">
			<span class="latitude">{{ .GeoLocation.Latitude }}</span>;
			<span class="longitude">{{ .GeoLocation.Longitude }}</span>
		</p>
		{{end}}

		{{ end }}
	</li>
	{{end}}
	</ol>
</section>
{{end}}

{{ if .Childs }}
<section class="childs">
<ol class="list">
{{range .Childs}}
<li class="child">
	<a href="/{{.Route}}" class="child-title child-link">{{.Title}}</a>
	<p class="child-description">{{.Description}}</p>
</li>
{{end}}
</ol>
</section>
{{end}}
`

const locationTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>

<div class="cleaner"></div>

{{ if .Tags }}
<section class="tags">
	<header>
		Tags:
	</header>

	<ul class="tags">
	{{range .Tags}}
	<li class="tag">
		<a href="{{.Route}}" title="{{.Description}}">{{.Name}}</a>
	</li>
	{{end}}
	</ul>
</section>
{{end}}

{{ if .RelatedItems }}
<section class="related-items">
<ol class="list">
{{range .RelatedItems}}
<li class="related-item">
	<a href="{{.Route}}">{{.Title}}</a>
	<p>{{.Description}}</p>
</li>
{{end}}
</ol>
</section>
{{end}}

{{ if .Childs }}
<section class="childs">
<ol class="list">
{{range .Childs}}
<li class="child">
	<a href="{{.Route}}" class="child-title child-link">{{.Title}}</a>
	<p class="child-description">{{.Description}}</p>
</li>
{{end}}
</ol>
</section>
{{end}}
`

const presentationTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<nav>
	<div class="nav-element pager deck-status">
		<span class="deck-status-current"></span> /	<span class="deck-status-total"></span>
	</div>

	<div class="nav-element controls">
		<button class="deck-prev-link" title="Previous">&#8592;</button>
		<button href="#" class="deck-next-link" title="Next">&#8594;</button>
	</div>

	<div class="nav-element jumper">
		<form action="." method="get" class="goto-form">
			<label for="goto-slide">Go to slide:</label>
			<input type="text" name="slidenum" id="goto-slide" list="goto-datalist">
			<datalist id="goto-datalist"></datalist>
			<input type="submit" value="Go">
		</form>
	</div>
</nav>

<section class="content">
{{.Content}}
</section>

<div class="cleaner"></div>
`

const messageTemplate = `
<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>
`

const errorTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>
`

var sitemapContentTemplate = fmt.Sprintf(`
<li>
	<a href="{{.Route}}" {{ if .Description }}title="{{.Description}}"{{ end }}>{{.Title}}</a>

	{{ if .Childs }}	
	<ol>
	%s
	</ol>
	{{ end }}
</li>`, ChildTemplatePlaceholder)

const sitemapTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
<ol>
{{.Content}}
</ol>
</section>
`

var tagmapContentTemplate = `
{{ if .Tags }}
<ul class="tags">
{{range .Tags}}
<li class="tag">
	<a name="{{.Name}}" title="{{.Description}}">{{.Name}}</a>
	<ol class="childs">
		{{range .Childs}}
		<li class="child">
			<a href="{{.Route}}" class="child-title child-link">{{.Title}}</a>
			<p class="child-description">{{.Description}}</p>
		</li>
		{{end}}
	</ol>
</li>
{{end}}
</ul>
{{else}}
There are currently no tags.
{{end}}
`

const tagmapTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>
`
