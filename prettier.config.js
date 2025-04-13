export default {
  "plugins": [ "prettier-plugin-go-template", "prettier-plugin-tailwindcss" ],
  "overrides": [ {
    "files" : [ "*.html", "*.html.tmpl" ],
    "options" : {"parser" : "go-template"}
  } ]
}
