const markdownpdf = require("markdown-pdf")
const fs = require("fs")

fs.createReadStream("test.md")
  .pipe(markdownpdf())
  .pipe(fs.createWriteStream("test.pdf"))